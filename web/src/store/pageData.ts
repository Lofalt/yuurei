import {defineStore} from "pinia";

export const usePageData = defineStore('pageData', {
    state: () => {
        return {
            pagedata: {
                data: "",
                count: 0,
            }
        }
    },
    actions: {
        countIncrease() {
            this.pagedata.count++
        }
    }
})

