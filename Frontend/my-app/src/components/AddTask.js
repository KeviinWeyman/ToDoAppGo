import React, { useState } from "react";

const AddTask = ({ socket }) => {
  const [task, setTask] = useState("");

  const sendTask = () => {
    if (task.trim()) {
      socket.send(task);
      setTask("");
    }
  };

  return (
    <div>
      <input value={task} onChange={(e) => setTask(e.target.value)} />
      <button onClick={sendTask}>Agregar</button>
    </div>
  );
};

export default AddTask;
