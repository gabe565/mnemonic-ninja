import vue from "@vitejs/plugin-vue";
import autoprefixer from "autoprefixer";
import { FileSystemIconLoader } from "unplugin-icons/loaders";
import Icons from "unplugin-icons/vite";
import { defineConfig } from "vite";
import { VitePWA } from "vite-plugin-pwa";
import vuetify from "vite-plugin-vuetify";
import { CompileCmudict } from "./src/plugins/compile-cmudict/CompileCmudict";

export default defineConfig({
  plugins: [
    vue(),
    Icons({
      compiler: "vue3",
      autoInstall: true,
      customCollections: {
        "mnemonic-ninja": FileSystemIconLoader("./src/assets/icons", (svg) =>
          svg.replace(/^<svg /, '<svg fill="currentColor" '),
        ),
      },
    }),
    vuetify({
      styles: {
        configFile: "src/scss/variables.scss",
      },
    }),
    VitePWA({
      filename: "service-worker.js",
      devOptions: {
        enabled: true,
      },
      includeAssets: ["favicon.ico", "img/icons/*"],
      manifest: {
        name: "Mnemonic Ninja",
        short_name: "Mnemonic Ninja",
        id: "/",
        description: "Convert a number into a phrase to aid in memorization.",
        theme_color: "#607D8B",
        background_color: "#607D8B",
        icons: [
          {
            src: "img/icons/android-chrome-192x192.png",
            sizes: "192x192",
            type: "image/png",
          },
          {
            src: "img/icons/android-chrome-512x512.png",
            sizes: "512x512",
            type: "image/png",
          },
          {
            src: "img/icons/android-chrome-maskable-192x192.png",
            sizes: "192x192",
            type: "image/png",
            purpose: "maskable",
          },
          {
            src: "img/icons/android-chrome-maskable-512x512.png",
            sizes: "512x512",
            type: "image/png",
            purpose: "maskable",
          },
        ],
      },
      workbox: {
        clientsClaim: true,
        globPatterns: ["**/*{js,css,html,woff2,svg}"],
      },
    }),
    CompileCmudict(),
  ],
  css: {
    postcss: {
      plugins: [autoprefixer({})],
    },
  },
});
