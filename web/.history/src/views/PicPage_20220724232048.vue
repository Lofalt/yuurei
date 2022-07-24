<template>
    <div class="container">
        <!-- <div class="panel" @click="changePage(2)">
            panel
        </div>
        <div class="panel">
            panel

        </div>
        <div class="panel">
            panel

        </div>
        <div class="panel">
            panel

        </div>
        <div class="panel">
            panel

        </div>
        <div class="panel">
            panel

        </div> -->
        <panel-com @click="changePage(2)"></panel-com>
        <panel-com @click="changePage(2)"></panel-com>
        <panel-com @click="changePage(2)"></panel-com>
        <panel-com @click="changePage(2)"></panel-com>
    </div>
</template>
<script lang="ts" setup>
import { usePageData } from '@/store/pageData';
import PanelCom from '../components/PanelCom.vue'
import { useRouter } from 'vue-router';
const pageData = usePageData()
const router = useRouter()

function changePage(num: number) {
    pageData.pagedata = { data: "fuckyou~" }
    router.push('/About')
}


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