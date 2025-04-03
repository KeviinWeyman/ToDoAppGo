import React, { useEffect } from "react";
import useStore from "../store";

const WebSocketClient = () => {
    const addTask = useStore((state) => state.addTask);
  
    useEffect(() => {
      const socket = new WebSocket("ws://localhost:8080/ws");
  
      socket.onmessage = (event) => {
        addTask(event.data);
      };
  
      return () => socket.close();
    }, [addTask]);
  
    return <></>;
  };
  
  export default WebSocketClient;