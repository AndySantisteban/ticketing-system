import react from "@vitejs/plugin-react";
import { defineConfig } from "vite";
import Pages from "vite-plugin-pages";

export default defineConfig({
  plugins: [
    react(),
    Pages({
      dirs: "src/pages",
    }),
  ],
  define: {
    "process.env": {},
  },
  build: {
    outDir: "../dist",
    emptyOutDir: true,
  },
});
