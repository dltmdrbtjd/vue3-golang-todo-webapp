import { defineStore } from "pinia";
import { Todo } from "./type";
import $http from "axios";

interface State {
  todoList: Todo[];
}

export const useTodoStore = defineStore("todo", {
  state: (): State => ({
    todoList: [],
  }),
  actions: {
    async addTodo(newTodo: Todo) {
      try {
        const resp = await $http.post("/todo", newTodo);
        this.todoList = [...this.todoList, resp.data.data];
      } catch (error) {
        console.error(error);
      }
    },
    async getTodo(username: string) {
      try {
        const resp = await $http.get(`/todo/${username}`);
        this.todoList = resp.data.data || [];
      } catch (error) {
        console.error(error);
      }
    },
    async deleteTodo(todoId: string) {
      try {
        this.todoList = this.todoList.filter((x) => x._id !== todoId);
        await $http.delete(`/todo/${todoId}`);
      } catch (error) {
        console.error(error);
      }
    },
  },
  getters: {
    getTodoList: (state): Todo[] => [...state.todoList],
  },
});
