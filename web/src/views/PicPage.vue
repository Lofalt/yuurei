<template>
    <div class="container">
      <panel-com v-for="item in articleList"
                 @click="changePage(1)"
                 :article="item" :key="item.ID"></panel-com>
    </div>
</template>
<script lang="ts" setup>
import { usePageData } from '@/store/pageData';
import PanelCom from '../components/PanelCom.vue'
import { useRouter } from 'vue-router';
import {ref} from "vue";
import axios from 'axios'
const pageData = usePageData()
const router = useRouter()

function changePage(num: number) {
    pageData.pagedata = { data: "fuckyou~" }
    // router.push('/About')
}

const articleList =  ref([])
axios.defaults.baseURL="/api"
axios.get("/yuurei/articles").then((result)=>{
  articleList.value = result.data.data.data
})



</script>
<style lang="less" scoped>
.container {
    width: 98%;
    float: left;
    height: 100%;
    overflow: auto;
    position: relative;
    bottom: 0;
    margin: auto auto;

    &::-webkit-scrollbar {
        width: 10px;
        // background-color: aquamarine;
        // color: black;
    }



    &::-webkit-scrollbar-thumb {
        width: 30px;
        background-color: rgba(73, 73, 73, .5);
        color: black;
        border-radius: 20px;
    }

    scrollbar-width: none;
}

::-webkit-scrollbar {
    width: 0px;
}

.panel {
    background-color: rgba(109, 155, 255, 0.8);
    width: 80%;
    height: 25vh;
    margin: 30px auto;
    border-radius: 5px;
    text-align: center;
    cursor: pointer;
}
</style>