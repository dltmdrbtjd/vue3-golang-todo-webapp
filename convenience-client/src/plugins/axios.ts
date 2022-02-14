import axios from "axios";
import { App } from "vue";
import getEnv from "@/scripts/env";

axios.defaults.baseURL = getEnv("VITE_BASE_URL");

export default {
    install: (app: App) => {
        axios.interceptors.request.use(
            (response) => {
                return Promise.resolve(response)
            },
            (error) => {
                return Promise.reject(error)
            }
        )
        
        app.config.globalProperties.$http = axios;
    }
}