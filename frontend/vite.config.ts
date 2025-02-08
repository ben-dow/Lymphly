import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), tailwindcss()],
  server: {
    proxy: {
      "/radar_pub_key.txt":{
        target: "https://ddx16zqbfs90u.cloudfront.net",
        changeOrigin: true,
      },
      "/api":{
        target: "https://ddx16zqbfs90u.cloudfront.net",
        changeOrigin: true,
      }
    }
  }
})
