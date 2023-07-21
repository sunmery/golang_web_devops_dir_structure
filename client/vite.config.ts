import { defineConfig, UserConfig } from "vite"
import electron from 'vite-plugin-electron'
import renderer from 'vite-plugin-electron-renderer'
import react from '@vitejs/plugin-react'
import {resolve} from 'node:path'

// https://vitejs.dev/config/
export default defineConfig(({mode}): UserConfig => {
  switch (mode) {
    case 'production':
      return {...publicConfig, ...productionConfig}
    case 'development':
      return {...publicConfig, ...developmentConfig}
    case 'test':
      return {...publicConfig, ...testConfig}
    default:
      throw new Error('请设置环境变量')
  }
})

// 公共配置
const publicConfig: UserConfig = {
  plugins:[
    electron([
      {
        // Main-Process entry file of the Electron App.
        entry: 'electron/main.ts',
      },
      {
        entry: 'electron/preload.ts',
        onstart(options) {
          // Notify the Renderer-Process to reload the page when the Preload-Scripts build is complete,
          // instead of restarting the entire Electron App.
          options.reload()
        },
      },
    ]),
    renderer(),
  ],

  resolve: {
    alias: {
      // 别名配置
      '@': resolve(__dirname, 'src'),
    },
  },

  // 打包目录
  base: './',
}

// 开发模式配置
const developmentConfig: UserConfig = {
  plugins: [
    electron([
      {
        // Main-Process entry file of the Electron App.
        entry: 'electron/main.ts',
      },
      {
        entry: 'electron/preload.ts',
        onstart(options) {
          // Notify the Renderer-Process to reload the page when the Preload-Scripts build is complete,
          // instead of restarting the entire Electron App.
          options.reload()
        },
      },
    ]),
    renderer(),
  ],

  server: {
    port: 3000,
    strictPort: true,
  },

  build: {
    target: ['es2021', 'chrome105', 'safari13'],

    // 为调试构建生成源代码映射 (sourcemap)
    sourcemap: !!process.env.TAURI_DEBUG,

    outDir: 'dist-dev',
  },
}

// 测试环境配置
const testConfig: UserConfig = {
  plugins: [
    electron([
      {
        // Main-Process entry file of the Electron App.
        entry: 'electron/main.ts',
      },
      {
        entry: 'electron/preload.ts',
        onstart(options) {
          // Notify the Renderer-Process to reload the page when the Preload-Scripts build is complete,
          // instead of restarting the entire Electron App.
          options.reload()
        },
      },
    ]),
    renderer(),
  ],

  server: {
    port: 443,
    strictPort: true,
  },

  build: {
    target: ['es2021', 'chrome105', 'safari13'],

    // don't minify for debug builds
    // minify: !process.env.TAURI_DEBUG ? 'esbuild' : false,
    minify: process.env.VITE_NODE_ENV === 'production' ? 'esbuild' : false,
    // 为调试构建生成源代码映射 (sourcemap)
    sourcemap: !!process.env.TAURI_DEBUG,

    outDir: 'dist-test',
  },
}

// 生产模式配置
const productionConfig: UserConfig = {
  plugins: [
    electron([
      {
        // Main-Process entry file of the Electron App.
        entry: 'electron/main.ts',
      },
      {
        entry: 'electron/preload.ts',
        onstart(options) {
          // Notify the Renderer-Process to reload the page when the Preload-Scripts build is complete,
          // instead of restarting the entire Electron App.
          options.reload()
        },
      },
    ]),
    renderer(),
  ],

  server: {
    host: '0.0.0.0',
    port: 443,
    strictPort: true,
  },

  build: {
    // 为调试构建生成源代码映射 (sourcemap)
    sourcemap: !!process.env.TAURI_DEBUG,

    outDir: 'dist-production',
  },
}
