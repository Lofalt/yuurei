import { defineStore } from "pinia";

export const usePageData = defineStore('pageData', {
    state: () => {
        return {
            pagedata: {
                data: ""
            }
        }
    }
})

