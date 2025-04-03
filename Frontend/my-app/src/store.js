import { create } from "zustand";


const useStore = create((set) => ({
    tasks: [],
    addTask: (task) => set((state) => ({tasks: [...state.tasks, task] })),
    setTasks: (tasks) => set({ tasks }),

}));

export default useStore;