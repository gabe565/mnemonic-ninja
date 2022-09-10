import { VuetifyResolver } from 'unplugin-vue-components/resolvers';
import { createVuePlugin as vue } from 'vite-plugin-vue2';
import Components from 'unplugin-vue-components/vite';
import { defineConfig } from 'vite';
import { VitePWA } from 'vite-plugin-pwa';

export default defineConfig({
  plugins: [
    vue(),
    Components({
      directives: false,
      resolvers: [
        VuetifyResolver(),
      ],
      types: [
        {
          from: 'vue-router',
          names: ['RouterLink', 'RouterView'],
        },
      ],
    }),
    VitePWA({
      filename: "service-worker.js",
      devOptions: {
        enabled: true,
      },
      includeAssets: [
        'favicon.ico',
        'img/icons/*',
      ],
      manifest: {
        name: 'Mnemonic Ninja',
        short_name: 'Mnemonic Ninja',
        id: '/',
        description: 'Convert a number into a phrase to aid in memorization.',
        theme_color: '#607D8B',
        background_color: '#607D8B',
        icons: [
          {
            src: 'img/icons/android-chrome-192x192.png',
            sizes: '192x192',
            type: 'image/png',
          },
          {
            src: 'img/icons/android-chrome-512x512.png',
            sizes: '512x512',
            type: 'image/png',
          },
          {
            src: 'img/icons/android-chrome-maskable-192x192.png',
            sizes: '192x192',
            type: 'image/png',
            purpose: 'maskable',
          },
          {
            src: 'img/icons/android-chrome-maskable-512x512.png',
            sizes: '512x512',
            type: 'image/png',
            purpose: 'maskable',
          },
        ],
      },
      workbox: {
        clientsClaim: true,
        globPatterns: ['**/*{js,css,html,woff2}'],
      },
    }),
  ],
  resolve: {
    alias: {
      "@": "/src",
    },
  },
});
