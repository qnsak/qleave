
import { precacheAndRoute } from 'workbox-precaching'

precacheAndRoute(self.__WB_MANIFEST)

// const CACHE_NAME = 'PWA-v1';

// const PRECACHE_ASSETS  = [
//   '/icons/icon-32x32.png',
//   '/icons/icon-128x128.png',
//   '/icons/icon-192x192.png',
//   '/icons/icon-512x512.png',
// ];

// self.addEventListener('install', event => {
//   event.waitUntil((async () => {
//     const cache = await caches.open(CACHE_NAME);
//     cache.addAll(PRECACHE_ASSETS);
//   })());
// });

// self.addEventListener('activate', event => {
//   event.waitUntil(clients.claim());
// });

// self.addEventListener('fetch', function(event) {
//   event.respondWith(async function() {
//     try {
//       var res = await fetch(event.request);
//       var cache = await caches.open('cache');
//       cache.put(event.request.url, res.clone());
//       return res;
//     } catch(error) {
//       return caches.match(event.request);
//     }
//   }());
// });