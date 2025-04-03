

export default {
  server: {
    proxy: {
      '/ws': 'http://localhost:8080', // Proxy para WebSocket
    },
  },
};


