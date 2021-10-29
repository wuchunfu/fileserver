package logx

import (
	"fileserver/middleware/configx"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

// ZapLogger 接收gin框架默认的日志
func ZapLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 请求地址
		requestPath := ctx.Request.URL.Path
		// 请求参数
		requestQuery := ctx.Request.URL.RawQuery
		// 请求客户端 ip
		clientIP := ctx.ClientIP()
		// 开始时间
		startTime := time.Now()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求时间
		requestTime := startTime.Format("2006-01-02 15:04:05")
		// 请求方式
		requestMethod := ctx.Request.Method
		// 请求协议
		requestProto := ctx.Request.Proto
		// 请求路由
		requestUri := ctx.Request.RequestURI
		// 状态码
		statusCode := ctx.Writer.Status()
		// 用户代理
		userAgent := ctx.Request.UserAgent()
		// 错误
		errors := ctx.Errors.ByType(gin.ErrorTypePrivate).String()
		// 处理请求
		ctx.Next()

		logger.Info(requestPath,
			zap.String("requestPath", requestPath),
			zap.String("requestQuery", requestQuery),
			zap.String("requestUri", requestUri),
			zap.String("requestProto", requestProto),
			zap.String("requestMethod", requestMethod),
			zap.Int("requestStatus", statusCode),
			zap.String("requestTime", requestTime),
			zap.String("clientIP", clientIP),
			zap.String("userAgent", userAgent),
			zap.String("errors", errors),
			zap.Duration("cost", latencyTime),
		)
	}
}

// ZapRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func ZapRecovery(stack bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
				if brokenPipe {
					logger.Error(ctx.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					ctx.Error(err.(error)) // nolint: errcheck
					ctx.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		ctx.Next()
	}
}

type Options struct {
	Config        zap.Config
	AppName       string        //日志文件前缀
	Development   bool          //是否是开发模式
	Level         zapcore.Level //日志等级
	LogFileDir    string        //文件保存地方
	ErrorFileName string        //error日志文件后缀
	WarnFileName  string        //warn日志文件后缀
	InfoFileName  string        //info日志文件后缀
	DebugFileName string        //debug日志文件后缀
	MaxAge        int           //保存的最大天数
	MaxBackups    int           //最多存在多少个切片文件
	MaxSize       int           //日志文件小大（M）
}

type ModOptions func(options *Options)

var (
	l                                *Logger
	sp                               = string(filepath.Separator)
	errorWS, warnWS, infoWS, debugWS zapcore.WriteSyncer       // IO输出
	debugConsoleWS                   = zapcore.Lock(os.Stdout) // 控制台标准输出
	errorConsoleWS                   = zapcore.Lock(os.Stderr)
)

type Logger struct {
	Logger      *zap.Logger
	ZapSugar    *zap.SugaredLogger
	ZapConfig   zap.Config
	Opts        *Options `json:"opts"`
	Initialized bool
	Mux         sync.RWMutex
}

func NewLogger(mod ...ModOptions) *zap.Logger {
	l = &Logger{}
	l.Mux.Lock()
	defer l.Mux.Unlock()
	if l.Initialized {
		l.Info("[NewLogger] logger initEd")
		return nil
	}

	l.Opts = &Options{
		AppName:       "app",
		Development:   true,
		Level:         zapcore.DebugLevel,
		LogFileDir:    "",
		ErrorFileName: "error.log",
		WarnFileName:  "warn.log",
		InfoFileName:  "info.log",
		DebugFileName: "debug.log",
		MaxAge:        30,  // 日志文件保留天数
		MaxBackups:    60,  // 最大保留日志文件数量
		MaxSize:       100, // 文件大小限制,单位MB
	}

	if l.Opts.Development {
		l.ZapConfig = zap.NewDevelopmentConfig()
	} else {
		l.ZapConfig = zap.NewProductionConfig()
	}

	if l.Opts.LogFileDir == "" {
		l.Opts.LogFileDir, _ = filepath.Abs(filepath.Dir(filepath.Join(".")))
		l.Opts.LogFileDir += sp + "logs" + sp
	}

	if l.Opts.Config.OutputPaths == nil || len(l.Opts.Config.OutputPaths) == 0 {
		l.ZapConfig.OutputPaths = []string{"stdout"}
	}

	if l.Opts.Config.ErrorOutputPaths == nil || len(l.Opts.Config.ErrorOutputPaths) == 0 {
		l.ZapConfig.OutputPaths = []string{"stderr"}
	}

	for _, fn := range mod {
		fn(l.Opts)
	}

	l.ZapConfig.Level.SetLevel(l.Opts.Level)
	l.Init()
	l.Initialized = true
	return l.Logger
}

func (l *Logger) Init() {
	l.setSyncs()
	var err error
	l.Logger, err = l.ZapConfig.Build(l.WithConfig())
	if err != nil {
		panic(err)
	}
	defer l.Logger.Sync()
}

func (l *Logger) setSyncs() {
	f := func(fN string) zapcore.WriteSyncer {
		appName := strings.TrimSpace(l.Opts.AppName)
		if l.Opts.AppName != "" {
			appName = l.Opts.AppName + "-"
		}
		fileName := l.Opts.LogFileDir + sp + appName + fN
		if l.Opts.Development {
			fileName = l.Opts.LogFileDir + sp + appName + "dev-" + fN
		}
		return zapcore.AddSync(&lumberjack.Logger{
			Filename:   fileName,
			MaxSize:    l.Opts.MaxSize,
			MaxBackups: l.Opts.MaxBackups,
			MaxAge:     l.Opts.MaxAge,
			Compress:   true,
			LocalTime:  true,
		})
	}
	errorWS = f(l.Opts.ErrorFileName)
	warnWS = f(l.Opts.WarnFileName)
	infoWS = f(l.Opts.InfoFileName)
	debugWS = f(l.Opts.DebugFileName)
	return
}

func (l *Logger) generateEncoderConfig() zapcore.EncoderConfig {
	customTimeEncoder := timeEncoder
	if l.Opts.Development {
		l.ZapConfig = zap.NewDevelopmentConfig()
		// 自定义时间输出格式
		customTimeEncoder = timeEncoder
	} else {
		l.ZapConfig = zap.NewProductionConfig()
		customTimeEncoder = timeUnixNano
	}

	// 自定义日志级别显示
	customLevelEncoder := func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(level.CapitalString())
	}

	// 自定义文件：行号输出项
	encodeCaller := zapcore.FullCallerEncoder
	if !l.Opts.Development {
		encodeCaller = zapcore.ShortCallerEncoder // 只显示 package/file.go:line
	}

	// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
	return zapcore.EncoderConfig{
		MessageKey:    "msg",
		LevelKey:      "level",
		TimeKey:       "ts",
		CallerKey:     "file",
		StacktraceKey: "trace",
		EncodeTime:    customTimeEncoder,
		EncodeLevel:   customLevelEncoder,
		EncodeCaller:  encodeCaller,
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
		EncodeName: zapcore.FullNameEncoder,
	}
}

// WithConfig 根据配置文件更新 logger
func (l *Logger) WithConfig() zap.Option {
	// 生成 encoderConfig
	encoderConfig := l.generateEncoderConfig()
	var encoder zapcore.Encoder
	if l.Opts.Development {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	// 实现四个判断日志等级的interface
	errorPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.ErrorLevel && zapcore.ErrorLevel-l.ZapConfig.Level.Level() > -1
	})
	warnPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel && zapcore.WarnLevel-l.ZapConfig.Level.Level() > -1
	})
	infoPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel && zapcore.InfoLevel-l.ZapConfig.Level.Level() > -1
	})
	debugPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DebugLevel && zapcore.DebugLevel-l.ZapConfig.Level.Level() > -1
	})

	cores := []zapcore.Core{
		zapcore.NewCore(encoder, errorWS, errorPriority),
		zapcore.NewCore(encoder, warnWS, warnPriority),
		zapcore.NewCore(encoder, infoWS, infoPriority),
		zapcore.NewCore(encoder, debugWS, debugPriority),
	}

	if l.Opts.Development {
		cores = append(cores, []zapcore.Core{
			zapcore.NewCore(encoder, errorConsoleWS, errorPriority),
			zapcore.NewCore(encoder, debugConsoleWS, warnPriority),
			zapcore.NewCore(encoder, debugConsoleWS, infoPriority),
			zapcore.NewCore(encoder, debugConsoleWS, debugPriority),
		}...)
	}

	return zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return zapcore.NewTee(cores...)
	})
}

var logger *zap.Logger

func GetLogger() *zap.Logger {
	return logger
}

// InitLog log instance init
func InitLog(setting *configx.Log) {
	// 设置日志级别
	level := strings.ToLower(strings.TrimSpace(setting.Level))
	logLevel := zap.DebugLevel
	switch level {
	case "error":
		logLevel = zap.ErrorLevel
	case "warn":
		logLevel = zap.WarnLevel
	case "info":
		logLevel = zap.InfoLevel
	case "debug":
		logLevel = zap.DebugLevel
	default:
		logLevel = zap.InfoLevel
	}

	logger = NewLogger(
		SetAppName(setting.AppName),
		SetDevelopment(setting.Development),
		SetLevel(logLevel),
		SetLogFileDir(setting.LogFileDir),
		SetErrorFileName(setting.ErrorFileName),
		SetWarnFileName(setting.WarnFileName),
		SetInfoFileName(setting.InfoFileName),
		SetDebugFileName(setting.DebugFileName),
		SetMaxAge(setting.MaxAge),
		SetMaxBackups(setting.MaxBackups),
		SetMaxSize(setting.MaxSize),
	)
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func timeUnixNano(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt64(t.UnixNano() / 1e6)
}

func SetAppName(AppName string) ModOptions {
	return func(option *Options) {
		option.AppName = AppName
	}
}

func SetDevelopment(Development bool) ModOptions {
	return func(option *Options) {
		option.Development = Development
	}
}

func SetLevel(Level zapcore.Level) ModOptions {
	return func(option *Options) {
		option.Level = Level
	}
}

func SetLogFileDir(LogFileDir string) ModOptions {
	return func(option *Options) {
		option.LogFileDir = LogFileDir
	}
}

func SetErrorFileName(ErrorFileName string) ModOptions {
	return func(option *Options) {
		option.ErrorFileName = ErrorFileName
	}
}

func SetWarnFileName(WarnFileName string) ModOptions {
	return func(option *Options) {
		option.WarnFileName = WarnFileName
	}
}

func SetInfoFileName(InfoFileName string) ModOptions {
	return func(option *Options) {
		option.InfoFileName = InfoFileName
	}
}

func SetDebugFileName(DebugFileName string) ModOptions {
	return func(option *Options) {
		option.DebugFileName = DebugFileName
	}
}

func SetMaxAge(MaxAge int) ModOptions {
	return func(option *Options) {
		option.MaxAge = MaxAge
	}
}

func SetMaxSize(MaxSize int) ModOptions {
	return func(option *Options) {
		option.MaxSize = MaxSize
	}
}

func SetMaxBackups(MaxBackups int) ModOptions {
	return func(option *Options) {
		option.MaxBackups = MaxBackups
	}
}

// With adds a variadic number of fields to the logging context.
// see https://github.com/uber-go/zap/blob/v1.10.0/sugar.go#L91
func (l *Logger) With(args ...interface{}) *Logger {

	l.ZapSugar = l.ZapSugar.With(args...)
	return l
}

// Debug package sugar of zap
func (l *Logger) Debug(args ...interface{}) {
	l.ZapSugar.Debug(args...)
}

// Debugf package sugar of zap
func (l *Logger) Debugf(template string, args ...interface{}) {
	l.ZapSugar.Debugf(template, args...)
}

// Info package sugar of zap
func (l *Logger) Info(args ...interface{}) {
	l.ZapSugar.Info(args...)
}

// Infof package sugar of zap
func (l *Logger) Infof(template string, args ...interface{}) {
	l.ZapSugar.Infof(template, args...)
}

// Warn package sugar of zap
func (l *Logger) Warn(args ...interface{}) {
	l.ZapSugar.Warn(args...)
}

// Warnf package sugar of zap
func (l *Logger) Warnf(template string, args ...interface{}) {
	l.ZapSugar.Warnf(template, args...)
}

// Error package sugar of zap
func (l *Logger) Error(args ...interface{}) {
	l.ZapSugar.Error(args...)
}

// Errorf package sugar of zap
func (l *Logger) Errorf(template string, args ...interface{}) {
	l.ZapSugar.Errorf(template, args...)
}

// Fatal package sugar of zap
func (l *Logger) Fatal(args ...interface{}) {
	l.ZapSugar.Fatal(args...)
}

// Fatalf package sugar of zap
func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.ZapSugar.Fatalf(template, args...)
}

// Panic package sugar of zap
func (l *Logger) Panic(args ...interface{}) {
	l.ZapSugar.Panic(args...)
}

// Panicf package sugar of zap
func (l *Logger) Panicf(template string, args ...interface{}) {
	l.ZapSugar.Panicf(template, args...)
}
