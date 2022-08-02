import {defineStore} from "pinia";

export const useUserInfo = defineStore('userInfo', {
    state: () => {
        return {
            user: {
                Name: null,
                Username: null,
                IsAdmin: false
            }
        }
    },
    actions: {
        fuc() {
        }
    }
})

