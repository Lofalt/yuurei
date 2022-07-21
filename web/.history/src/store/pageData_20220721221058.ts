import { defineStore } from "pinia";

export const usePageData = defineStore('pagedata', {
    state: () => {
        return {
            pagedata: {
                data: ""
            }
        }
    }
})

