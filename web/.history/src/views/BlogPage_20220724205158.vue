<template>
    <div class="container">
        <div class="head" @click.capture="showModal = false">
            <!-- <joke></joke>
            <joke></joke>
            <joke></joke>
            <joke></joke>
            <joke></joke> -->
        </div>
        <div class="left">
            <n-icon size="30" color="#000000" class="messageIcon" @click="changeModal">
                <message-circle />
            </n-icon>
            <transition>
                <div class="messageDiv" v-show="showModal">
                    <div class="inputUsername"><input type="text" placeholder="输入署名,为空匿名" /></div>
                    <wang-editor-mini ref="wang"></wang-editor-mini>
                    <!-- <transition name="button"> -->
                    <button class="sendButton">send</button>
                    <!-- </transition> -->
                </div>
            </transition>
        </div>
        <!-- <n-modal v-model:show="showTheModal" :display-directive="'if'" :transform-origin="'mouse'">
            <n-card size="huge" aria-modal="true" style="width: 600px; height:60vh;">
                <!-- <button @click="post">点我发送</button> -->
        <div class="messageDiv" style="height:80%">
            <div class="inputUsernameModal"><input type="text" placeholder="署名,为空匿名" /></div>
            <wang-editor-mini></wang-editor-mini>
            <!-- <transition name="button"> -->
            <button class="sendButton">send</button>
            <!-- </transition> -->
        </div>

        <!-- <template #footer> -->
        <!-- </template> -->
        </n-card>
        </n-modal> -->
    </div>
</template>

<script lang="ts" setup>
import { computed, provide, ref } from 'vue';
import WangEditorMini from "../components/wangeditor/WangEditorMini.vue"
import { useRouter } from 'vue-router'
import TinyMce from "../components/tinymce/TinyMce.vue"
import '@wangeditor/editor/dist/css/style.css' // 引入 css
import { NModal, NCard } from 'naive-ui';
import Joke from "../components/comments/Joke.vue"
import { onBeforeUnmount, shallowRef, onMounted } from 'vue'
import { Editor, Toolbar } from '@wangeditor/editor-for-vue'
import { MessageCircle } from "@vicons/tabler"
import { NIcon } from "naive-ui"
import PanelCom from "../components/PanelCom.vue"
const router = useRouter()
const showModal = ref(false)
const container = document.getElementsByClassName("container")[0] as HTMLElement
const showTheModal = ref(false)
const focus = computed(() => {
    return showModal.value || showTheModal.value
})

provide('focus', focus)
function post() {
    alert("post!")
}

const wang = ref<InstanceType<typeof WangEditorMini>>();

// const wang = ref(null)

function changeModal() {
    console.log(wang)
    if (document.body.clientWidth > 1024) {
        showModal.value = !showModal.value
    } else {
        showTheModal.value = !showTheModal.value
    }
}



</script>


<style lang="less" scoped>
.head {
    width: 100%;
    height: 100%;
}

.disapper {
    display: none;
}

.left {
    position: absolute;
    width: 100%;
    // height: 100%;
    // background-color: rgb(255, 128, 128);
    bottom: 0;
    height: 40vh;
    // display: flex;
    // align-items: flex-end;
    // justify-content: start;

    // @media (max-width:900px) {
    // width: 15%
    // z-index: 10;

    .messageIcon {
        position: absolute;
        // float: left;
        // margin: 10px;
        // vertical-align: bottom;
        bottom: 15px;
        left: 15px;
        cursor: pointer;
    }

    .messageDiv {
        background-color: #fff;
        width: 80%;
        float: left;
        left: 60px;
        position: absolute;
        bottom: 25px;
        height: 100%;
        // left: 60px;
        // height: 23vh;
        // flex: 10;
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

    .inputUsername {
        position: absolute;
        // left: 10px;
        top: -40px;
        z-index: 1;
        width: 40%;

        input {
            // width: 10%;
            padding-left: 10px;
            height: 30px;
            outline: none;
            border: 3px solid rgb(49, 49, 49)
        }
    }

}

.inputUsernameModal {
    position: relative;
    // left: 10px;
    top: -13px;
    z-index: 1;
    width: 20vw;

    input {
        // width: 10%;
        padding-left: 10px;
        height: 30px;
        outline: none;
        border: 3px solid rgb(49, 49, 49)
    }
}



.sendButton {
    position: absolute;
    right: 40px;
    top: 15px;
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

    // }
}

.container {
    width: 100%;
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

.button-leave-to,
.button-enter-from {
    opacity: 0;
    transform: scale(0, 0)
}

.button-leave-active,
.button-enter-active {
    transition: all .1s ease-in-out
}
</style>