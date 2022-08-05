<template>
  <div class="container">
    <transition name="scale">
      <div :style="{backgroundImage:`url(${bgi})`}" class="scalePic" @click="showScale=false" v-show="showScale"></div>
    </transition>
    <div class="head" @click.capture="showModal = false">
      <transition-group tag="p">
        <joke @emit-pic="acceptScale" v-for="(item,index) in msgList" :msg="item" :key="item.ID" @reload="getMsgs"/>
      </transition-group>
    </div>
    <div class="left">

      <n-icon size="30" color="#000000" class="messageIcon" @click="changeModal">
        <message-circle/>
      </n-icon>
      <transition>
        <div class="messageDiv" v-show="showModal">
          <div class="insertPic">
            <div @click="deletePic(index)" v-for="(item,index) in picList" :style="{backgroundImage:`url(${item})`}"
                 class="insidePic"></div>
          </div>
          <div class="inputUsername"><input type="text" v-model="userName" placeholder="输入署名,为空匿名"/></div>
          <!--                    <wang-editor-mini ref="wang" @change-content="changeContent"></wang-editor-mini>-->
          <emoji @emit-emoji="insertEmoji" v-show="showEmoji" class="emojiCard"/>
          <div class="messageHeader">
            <n-icon @click="upload" class="messageHeaderIcon" color="black" size="30">
              <picture-outlined/>
            </n-icon>
            <n-icon class="messageHeaderIcon" @click="toggleEmoji" color="black" size="30">
              <smile-outlined/>
            </n-icon>
            <div style="width: 90%"></div>
            <button class="sendButton" @click="sendMsg">send</button>
          </div>
          <textarea class="inputTextarea" @click="showEmoji=false" ref="inputTextarea" type="textarea" draggable="false"
                    v-model="msgContent"> </textarea>

          <!-- <transition name="button"> -->
          <!-- </transition> -->
          <input type="file" accept="image/jpeg,image/png" ref="inputFile" @change="getFile" style="display: none">
          <div class="uploadButton" :style="{backgroundImage:backgroundImage}" @click="editPic">
            <loading-com v-show="isLoading"/>
          </div>
        </div>
      </transition>
    </div>
    <n-modal v-model:show="iconEdit" display-directive="if">
      <n-card class="iconCard">
        <upload-pic @confirm="check" :name="`上传图片`" :ratio="1" directory="profileIcon" quality="40"
                    @send-pic="changeImg"/>
      </n-card>
    </n-modal>
    <n-modal v-model:show="showTheModal" :display-directive="'if'" :transform-origin="'mouse'">
      <n-card size="huge" aria-modal="true" style="width: 100%; height:450px;position: fixed;top: 0;">
        <!-- <button @click="post">点我发送</button> -->
        <div class="messageDiv">
          <div class="insertPic">
            <div @click="deletePic(index)" v-for="(item,index) in picList" :style="{backgroundImage:`url(${item})`}"
                 class="insidePic"></div>
          </div>

          <div class="inputUsernameModal"><input type="text" v-model="userName" placeholder="署名,为空匿名"/></div>
          <!--                    <wang-editor-mini @change-content="changeContent"></wang-editor-mini>-->
          <!-- <transition name="button"> -->
          <!--          <button class="sendButton" @click="sendMsg">send</button>-->
          <div class="messageHeader">
            <n-icon @click="upload" class="messageHeaderIcon" color="black" size="30">
              <picture-outlined/>
            </n-icon>
            <div style="width: 90%"></div>
            <button class="sendButton" @click="sendMsg">send</button>
          </div>
          <textarea class="inputTextarea" @click="showEmoji=false" ref="inputTextarea2" draggable="false"
                    v-model="msgContent"/>
          <div class="uploadButton inside" :style="{backgroundImage:backgroundImage}"
               @click="editPic"></div>


          <!-- </transition> -->
        </div>
        <!-- <template #footer> -->
        <!-- </template> -->
      </n-card>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import {computed, inject, nextTick, onMounted, onUpdated, provide, ref} from 'vue';
import WangEditorMini from "../components/wangeditor/WangEditorMini.vue"
import {useRouter} from 'vue-router'
import '@wangeditor/editor/dist/css/style.css' // 引入 css
import {NModal, NCard, NInput,useMessage} from 'naive-ui';
import Joke from "../components/comments/Joke.vue"
import {MessageCircle} from "@vicons/tabler"
import {NIcon} from "naive-ui"
import axios from "@/request/index"
import LoadingCom from "@/components/LoadingCom.vue";
import UploadPic from "@/components/file/UploadPic.vue"
import {PictureOutlined, SmileOutlined} from "@vicons/antd"
import Emoji from "@/components/comments/Emoji.vue"

const message = useMessage()
const showScale = ref(false)
const picList = ref([]) as any
const showEmoji = ref(false)
var EmojiConvertor = require('emoji-js');
var emoji = new EmojiConvertor()
const iconEdit = ref(false)
const isLoading = ref(false)
const router = useRouter()
const showModal = ref(false)
const container = document.getElementsByClassName("container")[0] as HTMLElement
const showTheModal = ref(false)
const msgContent = ref("")
const focus = computed(() => {
  return showModal.value || showTheModal.value
})
const backgroundImg = ref("")
const userName = ref("")
const msgList = ref([])
const inputTextarea = ref(null) as any
const bgi = ref("") as any

getMsgs()
const config = inject("globalConfig") as any

const backgroundImage = computed(()=>{
  if(config.value==null){
    return ``
  }
  if((config.value!=null && backgroundImg.value=='')||userName.value==''){
  return `url(${config.value.MessageDefaultIcon})`
  }
  return `url(${backgroundImg.value})`
})

const uploadImg = computed(()=>{
  if(backgroundImg.value==''||userName.value==''){
    return config.value.MessageDefaultIcon
  }
  return backgroundImg.value
})

function editPic(){
  if(userName.value==''){
    message.warning("请先输入署名")
    return
  }else{
    iconEdit.value = true
  }
}

onMounted(() => {
  let pics = document.getElementsByTagName("img")
  nextTick(() => {
    for (let i = 0; i < pics.length; i++) {
      pics[i].addEventListener("click", () => {
        alert(pics[i].getAttribute("src"))
      })
    }

  })

})

function acceptScale(img: string) {
  bgi.value = img;
  showScale.value = true
}

function insertEmoji(emoji: string) {
  inputTextarea.value.focus()
  inputTextarea.value.setRangeText(emoji)
  inputTextarea.value.selectionStart += emoji.length
}

function toggleEmoji() {
  showEmoji.value = !showEmoji.value
}

function check(arg: any) {
  backgroundImg.value = arg
  iconEdit.value = false
}

function getMsgs() {
  axios.get("/yuurei/msg/all", {}).then((res: any) => {
    let msgTmp: any[] = []
    res.data.data.forEach((msg: any) => {
      msgTmp.push(msg)
    })
    msgList.value = msgTmp
  })
}

provide('focus', focus)

function post() {
  alert("post!")
}

function changeContent(content: string) {
  msgContent.value = content
}

function sendMsg() {
  if(msgContent.value==null||msgContent.value==""){
    message.warning("先说点啥吧")
    return
  }
  axios.request("/yuurei/msg", "post", {
    MessageContent: msgContent.value.replace(/\r/ig, '').replace(/\n/ig, '<br/>'),
    Icon: uploadImg.value,
    UserName: userName.value,
    Pics: (() => {
      let pics: string = ""
      for (let i = 0; i < picList.value.length; i++) {
        pics = pics + picList.value[i]
        if (i != picList.value.length - 1) {
          pics = pics + ","
        }
      }
      return pics
    })()
  }).then((res) => {
    getMsgs()
    showModal.value = false
    showTheModal.value = false
    msgContent.value = ""
    picList.value = []
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

function deletePic(index: number) {
  picList.value.splice(index, 1)
}

const inputFile = ref(null)

function upload() {
  if(picList.value.length>=3){
    message.warning("超过上传上限")
    return
  }
  console.log(inputFile.value.click())
}

function RndNum(num: number) {
  return Math.floor(Math.random() * Math.pow(10, num)).toString()
}

function getFile(event: any) {

  let formFile = new FormData()
  formFile.append("file", event.target.files[0]);
  formFile.append("apply_info_id", RndNum(12));
  formFile.append("file_type", '');
  isLoading.value = true
  // backgroundImg.value = ''
  axios.file_upload("/yuurei/uploadImg?qua=60", formFile).then((res: any) => {
    isLoading.value = false
    // backgroundImg.value = res.data.fileName
    picList.value.push(res.data.fileName)
    inputFile.value.value = ""

  }).catch(() => {
    isLoading.value = false
    backgroundImg.value = "/img/-421408862.jpg"
  })
}

</script>


<style lang="less" scoped>

.scalePic {
  background-color: rgba(0, 0, 0, 0.5);
  width: 100%;
  height: 100%;
  position: absolute;
  z-index: 12;
  background-size: contain;
  background-repeat: no-repeat;
  background-position: center;
  cursor: pointer;
}

.iconCard {
  //height:80vh;
  width: 50vw;
  //min-width: 50vw;
  border: 3px solid black;
  @media (max-width: 800px) {
    width: 100vw;
  }

}

.v-leave-to,
.v-enter-from {
  opacity: 0;
  transform: translateX(1000px);
}

.v-leave-active,
.v-enter-active {
  transition: all 1.5s ease;
}

.uploadButton {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100px;
  height: 100px;
  background-color: #ffffff;
  background-size: cover;
  background-position: center;
  position: absolute;
  top: -110px;
  //left: 200px;
  right: 0;
  cursor: pointer;
  border: 3px solid rgb(49, 49, 49);

  &::after {
    content: "修改头像";
    //z-index: -1 ;
    box-sizing: border-box;
    display: flex;
    justify-content: center;
    align-items: center;
    color: white;
    width: 100px;
    height: 100px;
    //display: none;
    visibility: hidden;
    background: rgba(49, 49, 49, .2);
    position: absolute;
    left: -3px;
    top: -3px;
    border: 3px solid rgba(49, 49, 49, 0.2);
    @media (max-width: 1024px) {
      visibility: visible !important;

    }

  }

  &:hover::after {
    visibility: visible;

  }
}

.head {
  width: 100%;
  min-height: 50%;
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


  .inputUsername {
    position: absolute;
    left: -3px;
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

.messageDiv {
  background-color: #fff;
  width: 80%;
  //float: left;
  left: 60px;
  position: absolute;
  bottom: 25px;
  height: 15vh;
  border: 3px solid black;
  z-index: 10;
  // left: 60px;
  // height: 23vh;
  // flex: 10;
  // overflow: hidden;
  // white-space: nowrap;

  .insertPic {
    position: absolute;
    width: 100%;
    bottom: -120px;
    height: 100px;

    .insidePic {
      cursor: pointer;
      background-size: cover;
      background-position: center;
      background-color: #fff;
      border: 3px solid rgb(49, 49, 49);
      width: 70px;
      height: 80px;
      float: left;
      margin-right: 20px;
      position: relative;

      &:hover::after {
        visibility: visible;
      }

      &::after {
        content: '删除';
        color: white;
        display: flex;
        justify-content: center;
        align-items: center;
        border: 3px solid black;
        visibility: hidden;
        left: -3px;
        top: -3px;
        width: 100%;
        height: 100%;
        position: absolute;
        background: rgba(0, 0, 0, .3);


      }

      @media (max-width: 1024px) {
        width: 100px  ;
        height: 150px;
        &:nth-child(3) {
          margin-right: 0;
        }

      }
    }
  }

  @media (max-height: 800px) {
    height: 180px;
  }
  @media (max-width: 800px) {
    height: 300px;
    margin: 0 auto;
    left: 0;
    width: 100%;
  }

  &::after {
    content: '';
    width: 0px;
    height: 0px;
    border: 6px solid transparent;
    position: absolute;
    left: -14px;
    bottom: -3px;
    border-right-color: rgb(49, 49, 49);
    border-bottom-color: rgba(49, 49, 49);

    @media (max-width: 800px) {
      display: none;
    }
  }

  .inputTextarea {
    height: 70%;
    width: 100%;
    resize: none;
    outline: none;
    border: 0;
    padding: 10px;
    font-family: 微软雅黑;
    font-size: 1em;
    position: absolute;
    bottom: 0;
    left: 0;
    border-top: 1px solid rgba(49, 49, 49, .1);
    box-sizing: border-box;

    @media (max-width: 1024px) {
      height: 80%;
    }
  }

  .emojiCard {
    position: absolute;
    top: -210px;
    left: 0;
    z-index: 10;

  }

  .messageHeader {
    display: flex;
    justify-content: flex-start;
    width: 100%;
    padding: 10px;

    @media (max-width: 800px) {
      //height: 100px;
    }

    .messageHeaderIcon {
      cursor: pointer;
      margin-left: 5px;

    }

    .sendButton {
      //position: absolute;
      //right: 40px;
      //top: 15px;
      outline: none;
      border: none;
      padding: 5px 8px;
      background-color: #fff;
      cursor: pointer;
      justify-self: flex-end;
      place-self: flex-end;

      border: 3px solid rgb(49, 49, 49); // transition: all 1s;

      @media (max-width: 800px) {
        width: 100px;
      }

      &:hover {
        background-color: rgb(49, 49, 49);
        color: white
      }

      // }
    }
  }
}

.inside {
  //left: 10% !important;
  top: -120px !important;
  z-index: 10;
  @media (max-width: 1024px) {
    right: 0px !important;
  }
}

.inputUsernameModal {
  position: absolute;
  left: 0px;
  top: -13px;
  z-index: 1;
  width: 20vw;
  @media (max-width: 800px) {
    top: -50px;
  }

  input {
    // width: 10%;
    padding-left: 10px;
    height: 30px;
    outline: none;
    border: 3px solid rgb(49, 49, 49)
  }
}


.container {
  width: 98%;
  height: 100%;
  /* background-color: #fff; */
  /* margin: 0 auto; */
  overflow: auto;
  // scrollbar-width: none;
  &::-webkit-scrollbar {
    width: 5px;
    // background-color: aquamarine;
    // color: black;
  }


  &::-webkit-scrollbar-thumb {
    width: 5px;
    background-color: rgba(73, 73, 73, .3);
    color: black;
    border-radius: 20px;
  }
}

::-webkit-scrollbar {
  width: 3px;
  // background: rgba(255, 255, 255, 0.5);
}

.v-leave-to,
.v-enter-from {
  transform: translateY(100px); // opacity: 0 !important;
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

.scale-leave-to,
.scale-enter-from {
  opacity: 0;
  //transform: scale(0, 0)
}

.scale-leave-active,
.scale-enter-active {
  transition: all .5s ease-in-out
}

</style>