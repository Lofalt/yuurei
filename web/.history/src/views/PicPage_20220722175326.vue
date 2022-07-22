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
        <panel-com></panel-com>
        <panel-com></panel-com>
        <panel-com></panel-com>
        <panel-com></panel-com>
    </div>
</template>
        <script lang=" ts" setup>
import { usePageData } from '@/store/pageData';
import PanelCom from '../components/PanelCom.vue'

const pageData = usePageData()

function changePage(num: number) {
    pageData.pagedata = { data: "fuckyou~" }
}


</script>
            <style lang="less" scoped>
            .container {
                width: 100%;
                float: left;
                height: 100%;
                overflow: auto;
                position: relative;
                bottom: 0;
                margin: auto auto;
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
