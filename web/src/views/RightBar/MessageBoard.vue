<template>
  <div class="msgContainer" @wheel="listenScroll" @touchstart.stop="touchStart" @touchend.stop="touchEnd">
    <transition name="scale">
      <div :style="{backgroundImage:`url(${bgi})`}" class="scalePic" @click="showScale=false" v-show="showScale"></div>
    </transition>
    <div class="head" @click.capture="showModal = false">
      <transition-group tag="p" name="jokes">
        <joke @emit-pic="acceptScale" v-for="(item,index) in msgList" :style="`--i:${index%3}`" :msg="item"
              :key="item.ID"
              @reload="getMsgs"/>
      </transition-group>
    </div>
    <div class="loading">
      <loading-com v-show="isLoading"/>
    </div>
    <div class="loading end" v-show="msgList.length===total &&total>5">———— 时间的尽头 ————</div>
    <div class="left">

      <n-icon size="3vh" color="#000000" class="messageIcon" @click="changeModal">
        <message-circle/>
      </n-icon>
      <transition>
        <div class="messageDiv" v-show="showModal">
          <div class="insertPic">
            <div @click="deletePic(index)" v-for="(item,index) in picList" :style="{backgroundImage:`url(${item})`}"
                 class="insidePic"></div>
          </div>
          <div class="inputUsername"><input type="text" v-model="userName" placeholder="输入昵称,为空匿名"/></div>
          <!--                    <wang-editor-mini ref="wang" @change-content="changeContent"></wang-editor-mini>-->
          <emoji @emit-emoji="insertEmoji" v-show="showEmoji" class="emojiCard"/>
          <div class="messageHeader">
            <n-icon @click="upload" class="messageHeaderIcon" color="black" size="3vh">
              <picture-outlined/>
            </n-icon>
            <n-icon class="messageHeaderIcon" @click="toggleEmoji" color="black" size="3vh">
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
      <n-card size="huge" aria-modal="true" style="width: 100%; height:100vw;position: fixed;top: 0;">
        <!-- <button @click="post">点我发送</button> -->
        <div class="messageDiv2">
          <div class="insertPic">
            <div @click="deletePic(index)" v-for="(item,index) in picList" :style="{backgroundImage:`url(${item})`}"
                 class="insidePic"></div>
          </div>

          <div class="inputUsernameModal"><input type="text" v-model="userName" placeholder="输入昵称,为空匿名"/></div>
          <!--                    <wang-editor-mini @change-content="changeContent"></wang-editor-mini>-->
          <!-- <transition name="button"> -->
          <!--          <button class="sendButton" @click="sendMsg">send</button>-->
          <div class="messageHeader">
            <n-icon @click="upload" class="messageHeaderIcon" color="black" size="7vw">
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
import WangEditorMini from "../../components/wangeditor/WangEditorMini.vue"
import {useRouter} from 'vue-router'
import '@wangeditor/editor/dist/css/style.css' // 引入 css
import {NModal, NCard, NInput, useMessage} from 'naive-ui';
import Joke from "../../components/comments/Joke.vue"
import {MessageCircle} from "@vicons/tabler"
import {NIcon} from "naive-ui"
import axios from "@/request"
import LoadingCom from "@/components/LoadingCom.vue";
import UploadPic from "@/components/file/UploadPic.vue"
import {PictureOutlined, SmileOutlined} from "@vicons/antd"
import Emoji from "@/components/comments/Emoji.vue"
import {useUserInfo} from "@/store/UserInfo";

const userInfo = useUserInfo()
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
const msgContainer = document.getElementsByClassName("msgContainer")[0] as HTMLElement
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
const total = ref(0)
const pageSize = ref(6)
const pageNum = ref(1)
// const isLoading = ref(false)
const config = inject("globalConfig") as any

const backgroundImage = computed(() => {
  if (config.value == null) {
    return ``
  }
  if ((config.value != null && backgroundImg.value == '') || userName.value == '') {
    return `url(${config.value.MessageDefaultIcon})`
  }
  return `url(${backgroundImg.value})`
})

const uploadImg = computed(() => {
  if (backgroundImg.value == '' || userName.value == '') {
    return config.value.MessageDefaultIcon
  }
  return backgroundImg.value
})

function editPic() {
  if (userName.value == '') {
    message.warning("请先输入昵称")
    return
  } else {
    iconEdit.value = true
  }
}


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
  axios.get(`/yuurei/msg/all?pageNum=${pageNum.value}&pageSize=${pageSize.value}`, {}).then((res: any) => {
    let msgTmp: any[] = []
    total.value = res.data.total
    res.data.data.forEach((msg: any) => {
      msgTmp.push(msg)
    })
    msgList.value = msgTmp
  })
}

// 懒加载相关
var touchX = 0
var touchY = 0

function touchStart(event: any) {
  touchX = event.targetTouches[0].pageX;
  touchY = event.targetTouches[0].pageY;
}

function touchEnd(event: any) {
  // alert("hello")
  // alert(touchY)
  let touchYEnd = event.changedTouches[0].pageY
  let touches = touchYEnd - touchY
  if (touches < -90) {
    if (!isLoading.value && msgList.value.length < total.value) {
      isLoading.value = true
      setTimeout(() => {
        getNext()

      }, 1000);
    }
  }
}

// provide('focus', focus)
function scrollToTop() {
  const box = document.getElementsByClassName("msgContainer")[0] as HTMLSelectElement
  // box.scrollTo(0, 0)
  scroll(500, box)
}

function listenScroll(event: any) {
  const box = document.getElementsByClassName("msgContainer")[0] as HTMLSelectElement
  if (box.scrollTop + box.offsetHeight + 100 > box.scrollHeight) {
    if (!isLoading.value && msgList.value.length < total.value) {
      console.log("我直接疯狂加载")
      isLoading.value = true
      setTimeout(() => {
        getNext()

      }, 1000);

    } else {
      return
    }
  }
}

function getNext() {

  // if ((pageNum.value + 1) * pageSize.value > total.value) {
  //   return
  // }
  isLoading.value = true
  pageNum.value += 1
  axios.get(`/yuurei/msg/all?pageNum=${pageNum.value}&pageSize=${pageSize.value}`, {}).then((result: any) => {
    for (let i = 0; i < result.data.data.length; i++) {
      msgList.value.push(result.data.data[i])
    }
    total.value = result.data.total
    isLoading.value = false
  })
}

onMounted(() => {
  getMsgs()
  window.onscroll = () => {
    let scrollTop = document.documentElement.scrollTop; //获取滚动距离
    let scrollHeigh = document.documentElement.scrollHeight; //获取整个页面的高度
    let clientHeigh = document.documentElement.clientHeight; //获取
    if (scrollTop + clientHeigh - scrollHeigh >= -50) {
      getNext()
    }
  }
})

function changeContent(content: string) {
  msgContent.value = content
}

const fnlContent = computed(() => {
  if (picList.value[0] != undefined && msgContent.value === "") {
    return "分享图片"
  }
  return msgContent.value
})

function sendMsg() {
  console.log(fnlContent.value)
  if (userInfo.user.Sended) {
    message.warning("30秒内只能留言一次")
    // return
  }
  if (fnlContent.value == "") {
    message.warning("先说点啥吧")
    return
  } else if (msgContent.value.length > 140) {
    message.warning("字数超限(<=140)")
    return
  }
  axios.request("/yuurei/msg", "post", {
    MessageContent: fnlContent.value.replace(/\r/ig, '').replace(/\n/ig, '<br/>'),
    Icon: uploadImg.value,
    UserName: userName.value,
    IsAdmin: userInfo.user.IsAdmin,
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
    axios.get(`/yuurei/msg/all?pageNum=1&pageSize=1}`, {}).then((res: any) => {
      msgList.value.unshift(res.data.data[0])
    })
    userInfo.user.Sended = true
    showModal.value = false
    showTheModal.value = false
    msgContent.value = ""
    picList.value = []
    setTimeout(() => {
      userInfo.user.Sended = false
    }, 30000)
  })
}

const wang = ref<InstanceType<typeof WangEditorMini>>();

// const wang = ref(null)

function changeModal() {
  if ((document.body.clientWidth / document.body.clientHeight) > 9 / 16) {
    showModal.value = !showModal.value
  } else {
    showTheModal.value = !showTheModal.value
  }
}

function deletePic(index: number) {
  axios.request(`/yuurei/file?fileName=${picList.value[index]}`, "delete", {}).then((res: any) => {
    if (res.code == 200) {
      picList.value.splice(index, 1)

    }
  })
}

const inputFile = ref(null)

function upload() {
  if (picList.value.length >= 3) {
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
  axios.file_upload("/yuurei/uploadImg?qua=40&dir=messagePics", formFile).then((res: any) => {
    isLoading.value = false
    if (res.code == 200) {
      // backgroundImg.value = res.data.fileName
      picList.value.push(res.data.fileName)
      inputFile.value.value = ""
    } else {
      inputFile.value.value = ""
      message.error(res.msg)
    }

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
  @media (max-aspect-ratio: 9/16) {
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
  width: 10vh;
  height: 10vh;
  background-color: #ffffff;
  background-size: cover;
  background-position: center;
  position: absolute;
  top: -11vh;
  //left: 200px;
  right: 0;
  cursor: pointer;
  border: .4vh solid rgb(49, 49, 49);

  @media (max-aspect-ratio: 9/16) {
    width: 23vw;
    height: 23vw;
    top: -25vw !important;

  }

  &::after {
    content: "修改头像";
    //z-index: -1 ;
    box-sizing: border-box;
    font-size: 1.5vh;
    display: flex;
    justify-content: center;
    align-items: center;
    color: white;
    width: 101%;
    height: 101%;
    //display: none;
    visibility: hidden;
    background: rgba(49, 49, 49, .2);
    position: absolute;
    //left: -3px;
    //top: -3px;
    //border: 3px solid rgba(49, 49, 49, 0.2);
    @media (max-aspect-ratio: 9/16) {
      visibility: visible !important;
      border: .6vw solid rgba(49, 49, 49, 0.2);
      font-size: 4vw;
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
    bottom: 1.5vh;
    left: 2vh;
    cursor: pointer;
  }


  .inputUsername {
    position: absolute;
    left: -3px;
    top: -4.5vh;
    z-index: 1;
    //width: 40%;
    input {
      // width: 10%;
      padding-left: 10px;
      height: 3.5vh;
      outline: none;
      border: .35vh solid rgb(49, 49, 49)
    }
  }

}

.messageDiv {
  background-color: #fff;
  width: 90%;
  //float: left;
  left: 6.5vh;
  position: absolute;
  bottom: 2.5vh;
  height: 22vh;
  border: .35vh solid black;
  z-index: 10;
  // left: 60px;
  // height: 23vh;
  // flex: 10;
  // overflow: hidden;
  // white-space: nowrap;

  @media (max-aspect-ratio: 1/1) AND (min-aspect-ratio: 9/16) AND (min-height: 800px) {
    width: 80%;
    //max-height: ;
    // right: -85vw;
  }

  .insertPic {
    position: absolute;
    width: 100%;
    bottom: -12vh;
    height: 10vh;

    .insidePic {
      cursor: pointer;
      background-size: cover;
      background-position: center;
      background-color: #fff;
      border: 3px solid rgb(49, 49, 49);
      width: 8vh;
      height: 10vh;
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

        @media (max-aspect-ratio: 9/16) {
          visibility: visible;
        }

      }


      @media (max-aspect-ratio: 9/16) {
        width: 100px  ;
        height: 150px;
        &:nth-child(3) {
          margin-right: 0;
        }

      }
    }
  }

  @media (max-height: 800px) {
    //height: 180px;
  }
  @media (max-aspect-ratio: 9/16) {
    border: .7vw solid black;
    height: 70vw;
    margin: 0 auto;
    left: 0;
    width: 100%;
  }

  &::after {
    content: '';
    width: 0px;
    height: 0px;
    border: .7vh solid transparent;
    position: absolute;
    left: -1.5vh;
    bottom: -.3vh;
    border-right-color: rgb(49, 49, 49);
    border-bottom-color: rgba(49, 49, 49);

    @media (max-aspect-ratio: 9/16) {
      display: none;
    }
  }

  .inputTextarea {
    height: 16vh;
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
    border-top: .2vh solid rgba(49, 49, 49, .1);
    box-sizing: border-box;

    @media (max-aspect-ratio: 9/16) {
      height: 57vw;
      border-top: .4vw solid rgba(49, 49, 49, .1);
    }
  }

  .emojiCard {
    position: absolute;
    top: -21vh;
    left: 0;
    z-index: 10;

  }

  .messageHeader {
    display: flex;
    justify-content: flex-start;
    align-items: center;
    width: 100%;
    //padding: 10px;
    height: 5vh;

    @media (max-aspect-ratio: 9/16) {
      //height: 100px;
      height: 12vw;
    }

    .messageHeaderIcon {
      cursor: pointer;
      margin-left: 1vh;

    }

    .sendButton {
      //position: absolute;
      //right: 40px;
      //top: 15px;
      outline: none;
      //padding: .6vh 1vh;
      height: 4vh;
      width: 6.5vh;
      background-color: #fff;
      cursor: pointer;
      justify-self: flex-end;
      place-self: flex-end;
      font-size: 1em;
      margin-right: 1vh;
      border: .4vh solid rgb(49, 49, 49); // transition: all 1s;

      @media (max-aspect-ratio: 9/16) {
        width: 15vw;
        border: .8vw solid rgb(49, 49, 49); // transition: all 1s;
        height: 8vw;
        align-self: center;
        font-size: 4vw;
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
  //top: -120px !important;
  z-index: 10;
  @media (max-aspect-ratio: 9/16) {
    right: 0px !important;
  }
}

.inputUsernameModal {
  position: absolute;
  left: 0px;
  top: -80vw;
  z-index: 1;
  width: 40vw;
  top: -10vw;

  input {
    width: 100%;
    padding-left: 2vw;
    height: 7vw;
    outline: none;
    border: .3vh solid rgb(49, 49, 49);

    font-size: 3vw;
    //top: -50px;
    border: .5vw solid rgb(49, 49, 49);

  }
}


.msgContainer {
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

.jokes-leave-to,
.jokes-enter-from {
  //transform: translateX(1000px);
  //transform: scale(0, 0)
  opacity: 0;
}

.jokes-leave-active,
.jokes-enter-active {
  transition: all .5s ease;
  transition-delay: calc(Var(--i) * .1s);
}

.messageDiv2 {
  background-color: #fff;
  width: 90%;
  //float: left;
  left: 6.5vh;
  position: absolute;
  bottom: 2.5vh;
  height: 22vh;
  border: .35vh solid black;
  z-index: 10;
  // left: 60px;
  // height: 23vh;
  // flex: 10;
  // overflow: hidden;
  // white-space: nowrap;
  .uploadButton {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 10vh;
    height: 10vh;
    background-color: #ffffff;
    background-size: cover;
    background-position: center;
    position: absolute;
    top: -11vh;
    //left: 200px;
    right: 0;
    cursor: pointer;
    border: .4vh solid rgb(49, 49, 49);

    width: 23vw;
    height: 23vw;
    top: -25vw !important;


    &::after {
      content: "修改头像";
      //z-index: -1 ;
      box-sizing: border-box;
      font-size: 1.5vh;
      display: flex;
      justify-content: center;
      align-items: center;
      color: white;
      width: 101%;
      height: 101%;
      //display: none;
      visibility: hidden;
      background: rgba(49, 49, 49, .2);
      position: absolute;
      //left: -3px;
      //top: -3px;
      //border: 3px solid rgba(49, 49, 49, 0.2);
      visibility: visible !important;
      border: .6vw solid rgba(49, 49, 49, 0.2);
      font-size: 4vw;

    }

    &:hover::after {
      visibility: visible;

    }
  }

  .insertPic {
    position: absolute;
    width: 100%;
    bottom: -12vh;
    height: 10vh;

    .insidePic {
      cursor: pointer;
      background-size: cover;
      background-position: center;
      background-color: #fff;
      border: .3vh solid rgb(49, 49, 49);
      width: 25vw;
      height: 35vw;
      float: left;
      overflow: hidden;
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
        //border: .3vh solid black;
        visibility: hidden;
        left: -1vh;
        top: -1vh;
        width: 110%;
        height: 110%;
        position: absolute;
        background: rgba(0, 0, 0, .3);
        font-size: 1.5vh;
        visibility: visible;

      }

      &:nth-child(3) {
        margin-right: 0;
      }

    }
  }


  border: .7vw solid black;
  height: 70vw;
  margin: 0 auto;
  left: 0;
  width: 100%;

  &::after {
    content: '';
    width: 0px;
    height: 0px;
    border: .7vh solid transparent;
    position: absolute;
    left: -1.5vh;
    bottom: -.3vh;
    border-right-color: rgb(49, 49, 49);
    border-bottom-color: rgba(49, 49, 49);

    display: none;
  }

  .inputTextarea {
    height: 16vh;
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
    border-top: .2vh solid rgba(49, 49, 49, .1);
    box-sizing: border-box;

    height: 57vw;
    border-top: .4vw solid rgba(49, 49, 49, .1);
  }

  .emojiCard {
    position: absolute;
    top: -21vh;
    left: 0;
    z-index: 10;

  }

  .messageHeader {
    display: flex;
    justify-content: flex-start;
    align-items: center;
    width: 100%;
    //padding: 10px;
    height: 5vh;

    //height: 100px;
    height: 12vw;

    .messageHeaderIcon {
      cursor: pointer;
      margin-left: 1vh;

    }

    .sendButton {
      //position: absolute;
      //right: 40px;
      //top: 15px;
      outline: none;
      //padding: .6vh 1vh;
      height: 4vh;
      width: 6.5vh;
      background-color: #fff;
      cursor: pointer;
      justify-self: flex-end;
      place-self: flex-end;
      font-size: 1em;
      margin-right: 1vh;
      border: .4vh solid rgb(49, 49, 49); // transition: all 1s;

      width: 15vw;
      border: .8vw solid rgb(49, 49, 49); // transition: all 1s;
      height: 8vw;
      align-self: center;
      font-size: 4vw;

      &:hover {
        background-color: rgb(49, 49, 49);
        color: white
      }

      // }
    }
  }
}

.loading {
  position: relative;
  width: 100%;
  text-align: center;
  display: flex;
  justify-content: center;
}

.end {
  font-weight: bold;
  color: rgba(49, 49, 49, .2);
  font-size: 1.3em;
  line-height: 1.8em;
  height: 1.8em;
}
</style>