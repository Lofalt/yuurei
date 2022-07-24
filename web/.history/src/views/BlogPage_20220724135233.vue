<template>
    <div class="container" @click.self="showModal = false">
        <div class="left">
            <transition>
                <div class="messageDiv" v-show="showModal">
                    <div class="inputUsername">署名:<input type="text" /></div>
                    <wang-editor-mini></wang-editor-mini>
                    <!-- <transition name="button"> -->
                    <button class="sendButton">send</button>
                    <!-- </transition> -->
                </div>
            </transition>
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
const showModal = ref(false)
const container = document.getElementsByClassName("container")[0] as HTMLElement
function post() {
    alert("post!")
}

</script>

<style lang="less" scoped>
.inputUsername {
    position: absolute;
    left: 10px;
    top: -40px;
    z-index: 1;
    width: 40%;

    input {
        // width: 10%;
        height: 30px;
        outline: none;
        border: 3px solid rgb(49, 49, 49)
    }
}

.messageDiv {
    width: 80%;
    position: absolute;
    bottom: 25px;
    left: 60px;
    height: 23vh;

    // overflow: hidden;
    // white-space: nowrap;


    &::after {
        content: '';
        width: 0px;
        height: 0px;
        border: 6px solid transparent;
        position: absolute;
        left: -12px;
        bottom: 0;
        border-right-color: rgb(49, 49, 49);
        border-bottom-color: rgba(49, 49, 49)
    }
}

.sendButton {
    position: absolute;
    right: 10px;
    top: 10px;
    outline: none;
    border: none;
    padding: 5px 8px;
    background-color: #fff;
    cursor: pointer;

    border: 3px solid rgb(49, 49, 49); // transition: all 1s;

    &:hover {
        background-color: rgb(49, 49, 49);
        color: white
    }
}

.button-leave-to,
.button-enter-from {
    opacity: 0;
    transform: scale(0, 0)
}

.button-leave-active,
.button-enter-active {
    transition: all .1s ease-in-out
}

.messageIcon {
    // position: absolute;
    bottom: 15px;
    left: 5px;
    cursor: pointer;
}

.left {
    position: absolute;
    width: 100%;
    // height: 100%;
    // background-color: rgb(255, 128, 128);
    bottom: 0;

    @media (max-width:900px) {
        width: 15%
    }
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
    transform: translateY(100px; ); // opacity: 0 !important;
    opacity: 0;
}

.v-leave-active,
.v-enter-active {
    transition: all .4s ease-in-out;
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