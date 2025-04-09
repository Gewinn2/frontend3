// service-worker.js
const CACHE_NAME = 'notes-app-v1';
const STATIC_CACHE_URLS = [
    '/',
    '/index.html',
    '/index.css',
    '/App.css',
    '/App.js',
    '/Note.js',
    '/NoteForm.js',
    '/NoteList.js'
];

self.addEventListener('install', event => {
    event.waitUntil(
        caches.open(CACHE_NAME)
            .then(cache => cache.addAll(STATIC_CACHE_URLS)))
});

self.addEventListener('activate', event => {
    event.waitUntil(
        caches.keys().then(cacheNames => {
            return Promise.all(
                cacheNames.map(cache => {
                    if (cache !== CACHE_NAME) {
                        return caches.delete(cache);
                    }
                })
            );
        })
    );
});

self.addEventListener('fetch', event => {
    // Для API запросов или динамических данных используем стратегию "сеть, затем кэш"
    if (event.request.url.includes('/api/')) {
        event.respondWith(
            fetch(event.request)
                .then(response => {
                    // Клонируем ответ, потому что он может быть использован только один раз
                    const responseClone = response.clone();
                    caches.open(CACHE_NAME).then(cache => {
                        cache.put(event.request, responseClone);
                    });
                    return response;
                })
                .catch(() => {
                    return caches.match(event.request);
                })
        );
    } else {
        // Для статических ресурсов используем стратегию "кэш, затем сеть"
        event.respondWith(
            caches.match(event.request)
                .then(cachedResponse => {
                    return cachedResponse || fetch(event.request);
                })
        );
    }
});