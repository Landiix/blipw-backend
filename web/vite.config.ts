import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// Dev: браузер → Vite (:5173) → proxy → Go API (:8081).
// Так можно жить без CORS на бэкенде, пока учишься.
export default defineConfig({
  plugins: [react()],
  server: {
    port: 5173,
    proxy: {
      "/api": {
        target: "http://localhost:8081",
        changeOrigin: true,
      },
    },
  },
});
