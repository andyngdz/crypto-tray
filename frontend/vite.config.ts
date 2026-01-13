import react from '@vitejs/plugin-react'
import wails from '@wailsio/runtime/plugins/vite'
import path from 'path'
import { defineConfig } from 'vite'

// Build target for all configurations
const BUILD_TARGET = 'esnext'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react(), wails('./bindings')],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
      '@bindings': path.resolve(__dirname, './bindings/crypto-tray'),
    },
  },
  server: {
    port: 9245,
    strictPort: true,
  },
  build: {
    target: BUILD_TARGET,
  },
  esbuild: {
    target: BUILD_TARGET,
  },
  optimizeDeps: {
    esbuildOptions: {
      target: BUILD_TARGET,
    },
  },
})
