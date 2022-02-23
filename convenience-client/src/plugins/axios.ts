import axios from "axios";
import { App } from "vue";
import getEnv from "@/scripts/env";
import { setLocalStorage } from "@/scripts/localStorage";

axios.defaults.baseURL = getEnv("VITE_APP_BASE_URL");

axios.interceptors.response.use(
    (response) => {
        if(response.data.data?.name) {
            setLocalStorage("username" , response.data.data?.name)
        }
        return Promise.resolve(response)
    },
    (error) => {
        return Promise.reject(error)
    }
)

export default {
    install: (app: App) => {
        app.config.globalProperties.$http = axios;
    }
}