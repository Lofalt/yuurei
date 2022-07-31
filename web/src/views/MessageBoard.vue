<template>
    <div class="container">
        <div class="head" @click.capture="showModal = false">
          <transition-group tag="p">
          <joke v-for="(item,index) in msgList" :msg="item" :key="item.ID" @reload="getMsgs"/>
          </transition-group>
        </div>
        <div class="left">

            <n-icon size="30" color="#000000" class="messageIcon" @click="changeModal">
                <message-circle />
            </n-icon>
            <transition>
                <div class="messageDiv" v-show="showModal">
                    <div class="inputUsername"><input type="text" v-model="userName" placeholder="输入署名,为空匿名" /></div>
                    <wang-editor-mini ref="wang" @change-content="changeContent"></wang-editor-mini>

                    <!-- <transition name="button"> -->
                    <button class="sendButton" @click="sendMsg">send</button>
                    <!-- </transition> -->
                  <input type="file" accept="image/jpeg,image/png,image/bmp,image/gif" ref="inputFile" @change="getFile" style="display: none" >
                  <div class="uploadButton" :style="{backgroundImage:`url(${backgroundImg})`}" @click=upload></div>
                </div>
            </transition>
        </div>
        <n-modal v-model:show="showTheModal" :display-directive="'if'" :transform-origin="'mouse'">
            <n-card size="huge" aria-modal="true" style="width: 600px; height:60vh;">
                <!-- <button @click="post">点我发送</button> -->
                <div class="messageDiv" style="height:80%">
                    <div class="inputUsernameModal"><input type="text" v-model="userName" placeholder="署名,为空匿名" /></div>
                    <wang-editor-mini @change-content="changeContent"></wang-editor-mini>
                    <!-- <transition name="button"> -->
                  <button class="sendButton" @click="sendMsg">send</button>

                  <div class="uploadButton inside" :style="{backgroundImage:`url(${backgroundImg})`}" @click=upload></div>


                  <!-- </transition> -->
                </div>
                <!-- <template #footer> -->
                <!-- </template> -->
            </n-card>
        </n-modal>
    </div>
</template>

<script lang="ts" setup>
import { computed, provide, ref } from 'vue';
import WangEditorMini from "../components/wangeditor/WangEditorMini.vue"
import { useRouter } from 'vue-router'
// import TinyMce from "../components/tinymce/TinyMce.vue"
import '@wangeditor/editor/dist/css/style.css' // 引入 css
import { NModal, NCard } from 'naive-ui';
import Joke from "../components/comments/Joke.vue"
// import { onBeforeUnmount, shallowRef, onMounted } from 'vue'
// import { Editor, Toolbar } from '@wangeditor/editor-for-vue'
import { MessageCircle } from "@vicons/tabler"
import { NIcon } from "naive-ui"
import axios from "@/request/index"
const router = useRouter()
const showModal = ref(false)
const container = document.getElementsByClassName("container")[0] as HTMLElement
const showTheModal = ref(false)
const msgContent = ref("")
const focus = computed(() => {
    return showModal.value || showTheModal.value
})
const backgroundImg =ref("/img/-421408862.jpg")
const userName = ref("")
const msgList = ref([])


getMsgs()
function getMsgs(){
  axios.get("/yuurei/msg/all",{}).then((res)=>{
    console.log(res)
    let msgTmp:any[] = []
    res.data.data.forEach((msg)=>{
      msgTmp.push(msg)
    })
    msgList.value = msgTmp
  })
}
provide('focus', focus)
function post() {
    alert("post!")
}
function changeContent(content:string){
  msgContent.value= content
}
function sendMsg(){
  axios.request("/yuurei/msg","post",{
    MessageContent:msgContent.value,
    Icon:backgroundImg.value,
    UserName : userName.value,
  }).then((res)=>{
    console.log(res)
    getMsgs()
    showModal.value = false
    showTheModal.value = false
  })
}
const wang = ref<InstanceType<typeof WangEditorMini>>();

// const wang = ref(null)

function changeModal() {
    if (document.body.clientWidth > 1024) {
        showModal.value = !showModal.value
    } else {
        showTheModal.value = !showTheModal.value
    }
}
const inputFile = ref(null)
function upload(){
  console.log(inputFile.value.click())
}
function RndNum(num:number){
  return Math.floor(Math.random() * Math.pow(10,num)).toString()
}
function getFile(event:any){
  let formFile = new FormData()
  formFile.append("file", event.target.files[0]);
  formFile.append("apply_info_id", RndNum(12));
  formFile.append("file_type", '');
  axios.file_upload("/yuurei/uploadImg",formFile).then((res)=>{
    console.log(res)
    backgroundImg.value = res.data.fileName
  })
}

</script>


<style lang="less" scoped>

.v-leave-to,
.v-enter-from{
  opacity: 0;
  transform: translateX(1000px);
}

.v-leave-active,
.v-enter-active{
  transition: all 1.5s ease;
}
.uploadButton{
  width:100px;
  height: 100px;
  background-color: black;
  background-size: cover;
  background-position: center;
  position: absolute;
  top:-110px;
  left:200px;
  cursor: pointer;
  border:3px solid rgb(49,49,49);

  &::after{
    content: "修改头像";
    //z-index: -1 ;
    box-sizing: border-box;
    display: flex;
    justify-content: center;
    align-items: center;
    color:white;
    width: 100px;
    height: 100px;
    //display: none;
    visibility: hidden;
    background: rgba(49,49,49,.2);
    position:absolute;
    left:-3px;
    top:-3px;
    border:3px solid rgba(49,49,49,0.2);
  @media (max-width: 1024px){
    visibility:visible!important;

  }

  }
  &:hover::after{
    visibility:visible;

  }
}
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
     //height: 100%;
    // background-color: rgb(255, 128, 128);
    bottom: 0;
    //height: 40vh;
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
        height: 40vh;
        // left: 60px;
        // height: 23vh;
        // flex: 10;
        // overflow: hidden;
        // white-space: nowrap;

    @media (max-height: 800px){
      height:40vh;
    }
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
        //width: 40%;

        input {
            // width: 10%;
            padding-left: 10px;
            height: 30px;
            outline: none;
            border: 3px solid rgb(49, 49, 49)
        }
    }

}
.inside{
     left:10% !important;
     top:-100px !important;
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