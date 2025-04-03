const socket = new WebSocket("ws://localhost:8080/ws");

socket.onopen = () => {
  console.log("Conexión WebSocket abierta");
  // Enviar un mensaje de prueba al servidor
  socket.send("¡Hola desde el cliente!");
};

socket.onmessage = (event) => {
  console.log("Mensaje recibido:", event.data);
};

socket.onerror = (error) => {
  console.error("Error WebSocket:", error);
};

socket.onclose = () => {
  console.log("Conexión WebSocket cerrada");
};
