import axios, { AxiosRequestConfig, AxiosRequestHeaders, AxiosResponse } from 'axios';
import { ElMessage, ElMessageBox } from 'element-plus';
import { Session } from '/@/utils/storage';

// 配置新建一个 axios 实例
const service = axios.create({
  // axios中请求配置有baseURL选项，表示请求URL公共部分
  baseURL: import.meta.env.VITE_API_URL as any,
  // 超时
  timeout: 50000,
  headers: { 'Content-Type': 'application/json; charset=utf-8' },
});

// 添加请求拦截器
service.interceptors.request.use((config: AxiosRequestConfig) => {
    let headers: AxiosRequestHeaders
    let common: any
    if (config.headers) {
      headers = config.headers;
      common = headers.common;
    }
    let token = Session.get('token');
    // 在发送请求之前做些什么 token
    if (token) {
      // 让每个请求携带自定义token 请根据实际情况自行修改
      common["Authorization"] = `Bearer ${ token }`;
    }
    // get请求映射params参数
    if (config.method === 'get' && config.params) {
      let url = config.url + '?';
      for (const propName of Object.keys(config.params)) {
        const value = config.params[propName];
        const part = encodeURIComponent(propName) + "=";
        if (value !== null && typeof (value) !== "undefined") {
          if (typeof value === 'object') {
            for (const key of Object.keys(value)) {
              if (value[key] !== null && typeof (value[key]) !== 'undefined') {
                let params = propName + '[' + key + ']';
                let subPart = encodeURIComponent(params) + '=';
                url += subPart + encodeURIComponent(value[key]) + '&';
              }
            }
          } else {
            url += part + encodeURIComponent(value) + "&";
          }
        }
      }
      url = url.slice(0, -1);
      config.params = {};
      config.url = url;
    }
    return config;
  }, (error) => {
    console.log(error)
    // 对请求错误做些什么
    return Promise.reject(error);
  }
);

// 添加响应拦截器
service.interceptors.response.use((response: AxiosResponse) => {
    // 对响应数据做点什么
    const res = response.data;
    if (res.code && res.code !== 200) {
      // token 过期或者账号已在别处登录
      if (res.code === 401 || res.code === 4001) {
        // 清除浏览器全部临时缓存
        Session.clear();
        ElMessageBox.confirm('登录状态已过期，您可以继续留在该页面，或者重新登录', '系统提示', {
          confirmButtonText: '重新登录',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          // 去登录页
          window.location.href = '/';
        }).catch(() => {
        });
        return Promise.reject("认证失败，无法访问系统资源");
      } else if (res.code === 500) {
        ElMessage({
          message: "系统未知错误，请反馈给管理员",
          type: 'error'
        })
        return Promise.reject(new Error("系统未知错误，请反馈给管理员"))
      } else if (res.code === 403) {
        ElMessage({
          message: "当前操作没有权限",
          type: 'error'
        })
        return Promise.reject(new Error("当前操作没有权限"))
      } else if (res.code === 404) {
        ElMessage({
          message: "访问资源不存在",
          type: 'error'
        })
        return Promise.reject(new Error("访问资源不存在"))
      }
    } else {
      return response;
    }
  }, (error) => {
    console.log(error)
    // 对响应错误做点什么
    let { message, response } = error;
    if (message.indexOf('timeout') != -1) {
      message = '系统接口请求超时';
    } else if (message == 'Network Error') {
      message = '后端接口连接异常';
    } else if (message.includes("Request failed with status code")) {
      message = "系统接口" + message.substr(message.length - 3) + "异常";
    } else {
      if (response) {
        message = response.statusText;
      } else {
        message = '接口路径找不到';
      }
    }
    ElMessage.error({
      type: 'error',
      message: message
    });
    return Promise.reject(error);
  }
);

// 导出 axios 实例
export default service;
