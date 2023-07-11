import { fileURLToPath, URL } from 'node:url'

import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import { VitePWA } from "vite-plugin-pwa"
import * as dotenv from 'dotenv'
import * as fs from 'fs'
export default ({ mode }) => {
  const NODE_ENV = process.env.NODE_ENV || 'development'
  const envFiles = [
    `.env.${NODE_ENV}`
  ]
  for (const file of envFiles) {
    const envConfig = dotenv.parse(fs.readFileSync(file))
    for (const ec in envConfig) {
      process.env[ec] = envConfig[ec]
    }
  }

  return defineConfig({
    plugins: [
      vue(),
      VitePWA({
        mode: "development",
        injectRegister: 'script',
        strategies: 'injectManifest',
        devOptions: {
          enabled: true,
          type: 'module',
        },
        manifest: {
          name: process.env.APP_NAME,
          short_name: process.env.APP_SHORT_NAME,
          description: process.env.APP_DESCRIPTION,
          theme_color: '#fdd835',
          icons: [
            {
              src: "icons/icon-144x144.png",
              sizes: "144x144",
              type: "image/png",
              purpose: "any maskable"
            },
            {
              src: "icons/icon-192x192.png",
              sizes: "192x192",
              type: "image/png",
              purpose: "any maskable"
            },
            {
              src: "icons/icon-512x512.png",
              sizes: "512x512",
              type: "image/png",
              purpose: "any maskable"
            }
          ]
        },
        workbox: {
          manifest: false,
          registerType: 'autoUpdate',
          globPatterns: ["**/*.{js,css,html,ico,png,svg}"],
          workbox: {
            runtimeCaching: [
              {
                urlPattern: /(.*?)\.(js|css|ts)/,
                handler: 'CacheFirst',
                options: {
                  cacheName: 'js-css-cache',
                },
              },
              {
                urlPattern: /(.*?)\.(png|jpe?g|svg|gif|bmp|psd|tiff|tga|eps)/,
                handler: 'CacheFirst',
                options: {
                  cacheName: 'image-cache',
                },
              },
            ],
          },
        }
      }),
    ],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url))
      }
    },
    server: {
      proxy: {
        '/api': {
          target: process.env.VITE_SERVER_URL,
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/api/, ''),
        }
      }
    }
  });
}
