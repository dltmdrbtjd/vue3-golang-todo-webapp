import { defineStore } from "pinia";

type State = {
  testNumber: number;
};

export const useTestStore = defineStore("test", {
  state: (): State => ({
    testNumber: 0,
  }),
  actions: {
    Increse(num: number) {
      this.testNumber += num;
    },
  },
  getters: {
    getCurrentNumber(state): number {
      return state.testNumber + 6;
    },
  },
});
