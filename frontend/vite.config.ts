import path from 'path'
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// Build target for all configurations
const BUILD_TARGET = 'esnext'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
      '@wailsjs': path.resolve(__dirname, './wailsjs'),
    },
  },
  build: {
    target: BUILD_TARGET
  },
  esbuild: {
    target: BUILD_TARGET
  },
  optimizeDeps: {
    esbuildOptions: {
      target: BUILD_TARGET
    }
  }
})
