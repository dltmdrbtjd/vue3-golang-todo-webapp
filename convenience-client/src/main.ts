import { createApp } from "vue";
import App from "@/App.vue";
import router from "@/routes";
import { createPinia } from "pinia";
import "@/styles/tailwind.scss";
import "@/styles/main.scss";
import "@/plugins/axios"

createApp(App).use(createPinia()).use(router).mount("#app");
