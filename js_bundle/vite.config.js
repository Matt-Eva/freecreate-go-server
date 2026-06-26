import { defineConfig } from "vite";
import { resolve } from "path";

export default defineConfig({
  // Sets the root directory for Vite to look for source files
  root: resolve(__dirname, "."),

  build: {
    // Points the output to your Go project's static folder
    outDir: resolve(__dirname, "../static"),

    // CRITICAL: Prevents Vite from deleting your existing non-bundled css/js files
    emptyOutDir: false,

    rollupOptions: {
      input: {
        // Defines the entry point for your bundle
        bundle: resolve(__dirname, "main.js"),
      },
      output: {
        // Forces Vite to name the output file cleanly (e.g., static/bundle.js)
        // instead of adding a random hash if you want predictable Go templates
        entryFileNames: "[name].js",
        chunkFileNames: "[name].js",
        assetFileNames: "[name].[ext]",
      },
    },
  },
});
