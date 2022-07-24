<template>
    <div class="container">
        <!-- <div class="btnHeader"> -->
        <!-- {{ content }} -->
        <!-- </div> -->
        <!-- <div> -->
        <!-- <keep-alive> -->
        <!-- <router-view class="info"></router-view> -->
        <!-- </keep-alive> -->
        <!-- </div> -->
        <!-- <tiny-mce v-model="content" @change="change2"></tiny-mce> -->
        <button @click="showModal = true">点我</button>
        <n-modal v-model:show="showModal" :to="container">
            <wang-editor-mini></wang-editor-mini>
        </n-modal>
    </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import WangEditorMini from "../components/wangeditor/WangEditorMini.vue"
import { useRouter } from 'vue-router'
import TinyMce from "../components/tinymce/TinyMce.vue"
import '@wangeditor/editor/dist/css/style.css' // 引入 css
import { NModal } from 'naive-ui';

import { onBeforeUnmount, shallowRef, onMounted } from 'vue'
import { Editor, Toolbar } from '@wangeditor/editor-for-vue'
const router = useRouter()
const showModal = ref(false)
const container = document.getElementsByClassName("container")[0] as HTMLElement

</script>

<style lang="less" scoped>
.container {
    width: 80%;
    height: 100%;
    /* background-color: #fff; */
    /* margin: 0 auto; */
    overflow: auto;
    // scrollbar-width: none;
}

::-webkit-scrollbar {
    width: 3px;
    // background: rgba(255, 255, 255, 0.5);
}

.btnHeader {
    /* width: 100%; */
    /* height: 100%; */
    width: 80%;
    margin: 8vh auto;
    text-align: center;
    border-bottom: 3px solid rgba(255, 255, 255, 0.1);
    padding-bottom: 5vh;
}

.button {
    margin: 0 auto;
    background-color: rgb(255, 255, 255);
    border: none;
    padding: 15px 25px;
    margin: 20px;
    border-radius: 10px;
    color: rgb(43, 43, 43);
    cursor: pointer;

    &:hover {
        background-color: rgb(252, 125, 125);
    }
}

.info {
    // width: 100%;
    height: 100vh;
    // background-color: aliceblue;
    overflow: auto;
}
</style>