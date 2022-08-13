import {defineStore} from "pinia";

export const useComment = defineStore('commentData', {
    state: () => {
        return {
            commentData: {
                username: "",
                icon: "",
            }
        }
    },
    actions: {
        countIncrease() {
            // this.pagedata.count++
        }
    }
})

