import { User } from '@/models/user.model';
import { getLocalStorage } from '@/scripts/localStorage';
import { defineStore } from 'pinia';
import $http from 'axios';

interface State {
    userInfo: User
    oauthVerification: boolean
}

export const useUserStore = defineStore("user", {
    state: (): State => ({
        userInfo: {
            name: "",
            picture: ""
        },
        oauthVerification: false
    }),
    actions: {
        async getUserInfo() {
            const googleAccessToken = getLocalStorage("google-access-token");
            try {
                const resp = await $http.get(`/google-userinfo/${googleAccessToken}`)
                this.userInfo = resp.data.data
            } catch(error) {
                console.error(error)
            }
        },
        googleTokenVerification(status: boolean) {
            this.oauthVerification = status
        }
    },
    getters: {
        getUserDetail: (state): User => state.userInfo,
        getTokenVerification: (state): boolean => state.oauthVerification,
    }
})