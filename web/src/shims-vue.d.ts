// 声明文件，*.vue 后缀的文件交给 vue 模块来处理
declare module '*.vue' {
    import { DefineComponent } from 'vue'
    const component: DefineComponent<{}, {}, any>
    export default component
}

declare module '*.svg'
declare module '*.png'
declare module '*.jpg'
declare module '*.jpeg'
declare module '*.gif'
declare module '*.bmp'
declare module '*.tiff'
declare module '*.json'
declare module '*.scss'
declare module '*.js'
declare module '*.ts'
declare module '*.d.ts'

interface AnyObject {
    [key: string]: any
}

// 声明文件，定义全局变量。其它 app.config.globalProperties.xxx，使用 getCurrentInstance() 来获取
interface Window {
  nextLoading: boolean;
}
