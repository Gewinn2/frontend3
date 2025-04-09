// serviceWorkerRegistration.js
export function register(config) {
    if ('serviceWorker' in navigator) {
        window.addEventListener('load', () => {
            const swUrl = `${process.env.PUBLIC_URL}/service-worker.js`;

            navigator.serviceWorker.register(swUrl)
                .then(registration => {
                    console.log('ServiceWorker registration successful');

                    if (config && config.onUpdate) {
                        registration.onupdatefound = () => {
                            const installingWorker = registration.installing;
                            if (installingWorker == null) return;

                            installingWorker.onstatechange = () => {
                                if (installingWorker.state === 'installed') {
                                    if (navigator.serviceWorker.controller) {
                                        config.onUpdate(registration);
                                    }
                                }
                            };
                        };
                    }
                })
                .catch(error => {
                    console.error('ServiceWorker registration failed:', error);
                });
        });
    }
}

export function unregister() {
    if ('serviceWorker' in navigator) {
        navigator.serviceWorker.ready.then(registration => {
            registration.unregister();
        });
    }
}