import vue from '@vitejs/plugin-vue';
import { resolve } from 'path';
import type { ConfigEnv, UserConfigExport } from 'vite';
import { loadEnv } from '/@/utils/viteBuild';
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'
import ElementPlus from 'unplugin-element-plus/vite'

const pathResolve = (dir: string): any => {
  return resolve(__dirname, '.', dir);
};

const { VITE_PORT, VITE_OPEN, VITE_API_URL, VITE_PUBLIC_PATH, VITE_DROP_CONSOLE } = loadEnv();

const alias: Record<string, string> = {
  '/@': pathResolve('./src/'),
  'vue-i18n': 'vue-i18n/dist/vue-i18n.cjs.js',
};

// https://vitejs.dev/config/
export default ({ command }: ConfigEnv): UserConfigExport => {
  return {
    root: process.cwd(),
    resolve: { alias },
    base: process.env.NODE_ENV === 'production' ? VITE_PUBLIC_PATH : './',
    server: {
      host: '0.0.0.0',
      port: VITE_PORT,
      open: VITE_OPEN,
      proxy: {
        '/api': {
          target: VITE_API_URL,
          ws: true,
          changeOrigin: true,
          rewrite: (path: string) => path.replace(/^\/api/, ''),
        },
      },
    },
    build: {
      outDir: 'dist',
      // minify: 'esbuild',
      minify: 'terser',
      sourcemap: false,
      target: 'es2015',
      terserOptions: {
        compress: {
          keep_infinity: true,
          // Used to delete console in production environment
          drop_console: VITE_DROP_CONSOLE,
        },
      },
      // Turning off brotliSize display can slightly reduce packaging time
      brotliSize: false,
      chunkSizeWarningLimit: 2000,
    },
    optimizeDeps: {
      include: [
        'element-plus/lib/locale/lang/zh-cn',
        'element-plus/lib/locale/lang/en',
      ],
    },
    define: {
      __VUE_I18N_LEGACY_API__: JSON.stringify(false),
      __VUE_I18N_FULL_INSTALL__: JSON.stringify(false),
      __INTLIFY_PROD_DEVTOOLS__: JSON.stringify(false),
    },
    plugins: [
      vue(),
      Components({
        resolvers: [ElementPlusResolver()],
      }),
      ElementPlus(),
    ],
  }
};
