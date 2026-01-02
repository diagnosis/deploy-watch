import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from "@tailwindcss/vite";
import tanstackRouter from "@tanstack/router-plugin/vite";

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    tanstackRouter({
      target: 'react',
      autoCodeSplitting: true,
    }),
    react(),
    tailwindcss()
  ],
  server: {
    proxy: {
      '/api': {
        target: 'https://backend-cold-bush-2228.fly.dev',
        changeOrigin: true,
        secure: true,
      },
      '/auth': {
        target: 'https://backend-cold-bush-2228.fly.dev',
        changeOrigin: true,
        secure: true,
      },
    },
  },
})

