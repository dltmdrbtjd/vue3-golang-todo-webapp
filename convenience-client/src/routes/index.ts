import { createWebHistory, createRouter } from "vue-router";

const routes = [
  {
    path: "/",
    name: "Todo",
    component: () => import("@/pages/Todo.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
