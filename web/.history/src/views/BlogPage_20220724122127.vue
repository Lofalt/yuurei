<template>
    <div class="container" @click.self="showModal = false">
        <transition>
            <div class="messageDiv" v-show="showModal">
                <wang-editor-mini></wang-editor-mini>
            </div>
        </transition>
        <div class="left">
            <n-icon size="30" color="#000000" class="messageIcon" @click="showModal = !showModal">
                <message-circle />
            </n-icon>
        </div>
        <n-modal :to="container" :display-directive="'show'" :transform-origin="'mouse'">
            <n-card size="huge" aria-modal="true" style="width: 600px;">
                <template #footer>
                    <button @click="post">点我发送</button>
                </template>
            </n-card>
        </n-modal>
    </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import WangEditorMini from "../components/wangeditor/WangEditorMini.vue"
import { useRouter } from 'vue-router'
import TinyMce from "../components/tinymce/TinyMce.vue"
import '@wangeditor/editor/dist/css/style.css' // 引入 css
import { NModal, NCard } from 'naive-ui';

import { onBeforeUnmount, shallowRef, onMounted } from 'vue'
import { Editor, Toolbar } from '@wangeditor/editor-for-vue'
import { MessageCircle } from "@vicons/tabler"
import { NIcon } from "naive-ui"
const router = useRouter()
const showModal = ref(true)
const container = document.getElementsByClassName("container")[0] as HTMLElement
function post() {
    alert("post!")
}

</script>

<style lang="less" scoped>
.messageDiv {
    width: 80%;
    position: absolute;
    bottom: 25px;
    left: 50px;
    height: 23vh;
    overflow: hidden;
    white-space: nowrap;

    &::after {
        content: '';
        width: 0px;
        height: 0px;
        border: 3px solid white;
        position: absolute;
        left: 0;
        bottom: 0;

    }
}

.messageIcon {
    position: absolute;
    bottom: 15px;
    left: 5px;
    cursor: pointer;
}

.left {
    width: 5%;
    height: 100%;
    // background-color: rgb(255, 128, 128);
}

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

.v-leave-to,
.v-enter-from {
    width: 0;
    // opacity: 0 !important;
}

.v-leave-active,
.v-enter-active {
    transition: all .5s ease-in-out;
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