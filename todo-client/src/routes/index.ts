import { createWebHistory, createRouter } from 'vue-router';

const routes = [
  {
    path: '/',
    name: 'Todo',
    component: () => import('@/pages/Todo.vue'),
  },
  {
    path: '/google-login/callback',
    name: 'GoogleAuth',
    component: () => import('@/pages/GoogleAuth.vue'),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
