import tailwindcss from "@tailwindcss/vite";
import react from "@vitejs/plugin-react";
import vike from "vike/plugin";
import { defineConfig } from "vite";

export default defineConfig({
  ssr: {
    noExternal: ['radar-sdk-js']
  },
  plugins: [vike({}), react({}), tailwindcss()],
  build: {
    target: "es2022",
    chunkSizeWarningLimit: 1000,
    rollupOptions: {
      output: {
        manualChunks(id) {
          if (id.includes('node_modules')) {
            return id.toString().split('node_modules/')[1].split('/')[0].toString();
          }
        }
      }
    }
  },
  server: {
    proxy: {
      "/radar_pub_key.txt": {
        target: "https://ddx16zqbfs90u.cloudfront.net",
        changeOrigin: true,
      },
      "/api": {
        target: "https://ddx16zqbfs90u.cloudfront.net",
        changeOrigin: true,
      },
    },
  },
});
