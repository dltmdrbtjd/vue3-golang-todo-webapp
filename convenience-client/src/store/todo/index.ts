import { defineStore } from "pinia";
import { Todo } from "./type";
import $http, { AxiosError } from "axios";

interface State {
  todoList: Todo[];
}

export const useTodoStore = defineStore("todo", {
  state: (): State => ({
    todoList: [],
  }),
  actions: {
    isEditTodo(todoId: string, edit: boolean) {
      const todoItemIndex = this.todoList.findIndex((x) => x._id === todoId);
      this.todoList[todoItemIndex].isEdit = edit;
    },
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
    async editTodo(todoId: string, content: string, index: number) {
      try {
        this.todoList[index].content = content;
        await $http.put(`/todo/${todoId}`, { content });
        this.todoList[index].isEdit = false;
      } catch (error) {
        console.error(error);
      }
    },
  },
  getters: {
    getTodoList: (state): Todo[] => [...state.todoList],
  },
});
