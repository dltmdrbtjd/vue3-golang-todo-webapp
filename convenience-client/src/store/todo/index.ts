import { defineStore } from "pinia";
import { Todo } from "./type";
import $http from "axios";

interface State {
    todoList: Todo[]
}

export const useTodoStore = defineStore('todo', {
    state: (): State => ({
        todoList: []
    }),
    actions: {
        async addTodo(newTodo: Todo) {
            const resp = await $http.post("/todo", newTodo)
            console.log(resp)
        }
    }
})