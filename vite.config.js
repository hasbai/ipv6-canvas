import { defineConfig } from "vite";
import Vue from "@vitejs/plugin-vue";
import path from "path";

export default defineConfig({
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "src"),
      "@c": path.resolve(__dirname, "src/components"),
    },
  },
  base: "/",
  plugins: [Vue()],
  define: {
    "process.env": {},
  },
  server: {
    proxy: {
      // "/image": {
      //   target: "https://dev.x.hath.top:8443/",
      //   changeOrigin: true,
      //   ws: true,
      //   secure: false,
      // },
    },
  },
});
