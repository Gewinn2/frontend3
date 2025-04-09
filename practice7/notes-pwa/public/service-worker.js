const CACHE_NAME = 'notes-pwa-v3';
const urlsToCache = [
    './',
    './index.html',
    './static/js/main.[hash].js',
    './static/css/main.[hash].css',
    './manifest.json'
];

self.addEventListener('install', (event) => {
    event.waitUntil(
        caches.open(CACHE_NAME).then((cache) => cache.addAll(urlsToCache))
    );
})

self.addEventListener('activate', (event) => {
    event.waitUntil(
        caches.keys().then((keys) =>
            Promise.all(keys.map((key) =>
                    key !== CACHE_NAME ? caches.delete(key) : null
                )
            )
        )
    );
});