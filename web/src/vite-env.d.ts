/// <reference types="vite/client" />
/// <reference types="vue/ref-macros" />
/// <reference types="element-plus/global" />

interface ImportMetaEnv extends Readonly<Record<string, string>> {
  readonly ENV: string
  readonly VITE_PORT: string
  readonly VITE_OPEN: string
  readonly VITE_PUBLIC_PATH: string
  readonly VITE_API_URL: string
  readonly VITE_DROP_CONSOLE: string
  readonly VITE_APP_TITLE: string
  // 更多环境变量...
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
