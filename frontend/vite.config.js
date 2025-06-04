import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      '/glow-zaar': {
        target: 'http://localhost:8989',
        changeOrigin: true,
        secure: false,
      },
    },
  },
})
