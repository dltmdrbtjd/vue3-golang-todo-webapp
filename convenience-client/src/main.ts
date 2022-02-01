import { createApp } from "vue";
import App from "@/App.vue";
import router from "@/routes";
import { createPinia } from "pinia";
import "./styles/tailwind.scss";
import "./styles/main.scss";

createApp(App).use(createPinia()).use(router).mount("#app");
