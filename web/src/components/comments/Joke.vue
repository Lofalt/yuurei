<template>
  <div>
    <div class="messageBox">
      <div class="icon" :style="{backgroundImage:`url(${backgroundImage})`}">
        <div class="name" :class="msg.IsAdmin?`admin`:``">{{ userName }}</div>
      </div>
      <div class="messageInfo">
        <div v-show="userInfo.user && userInfo.user.IsAdmin " class="delete" @click="deleteMsg">删除</div>
        <div class="reply" @click="showReply=!showReply">{{ !showReply ? `回复` : `收起` }}</div>
        <div class="date">{{
            new Date(msg.CreatedAt).Format("yyyy/M/dd") + " \n\n\n" + (new Date(msg.CreatedAt)).Format("hh:mm:ss")
          }}
        </div>
        <div :id="`info${msg.ID}`" class="info" v-html="msg.MessageContent">
        </div>
      </div>
    </div>
    <n-message-provider>
      <send-reply v-show="showReply" :ID="msg.ID" :UserName="userName" @send="accept" @upload="upload"
                  @change-name="changeName"/>
    </n-message-provider>
    <replies v-for="item in list" :key="item.ID" :child="item" @reload="accept"/>
  </div>
</template>
<!--    <div class="messageInfo">as </div>-->


<script lang="ts" setup>
import {computed, inject, onMounted, ref} from 'vue'
import axios from '@/request/index'
import {useUserInfo} from "@/store/UserInfo";
import {isMobile} from "@/util/utils";
import Replies from "./Replies.vue";
import SendReply from "./SendReply.vue"
import {NMessageProvider} from "naive-ui"

const showReply = ref(false)
const flag = ref(isMobile())
const backgroundImg = ref("")
const userInfo = useUserInfo()
const emit = defineEmits(['reload', 'emitPic', 'upload', 'change-name'])
const content = ref("")
const list = ref([]) as any
const props = defineProps<{
  msg: any
}>()

function upload() {
  console.log("joke")
  emit('upload')
}

list.value = props.msg.Replies
let pics = props.msg.Pics.split(",")

function changeName(name: string) {
  console.log(name)
  emit('change-name', name)
}

function accept() {
  axios.get("/yuurei/msg/" + props.msg.ID, {}).then((res: any) => {
    list.value = res.data.data.Replies
    showReply.value = false
  })
}

function open(url: string) {
  window.open(url)
}

const config = inject('globalConfig') as any
backgroundImg.value = props.msg.Icon

const userName = computed(() => {
  if (props.msg.IsAdmin) {
    return config.value.MyName
  }
  if (props.msg.IsAnonymous) {
    return config.value.AnonymousName
  }
  return props.msg.UserName

})

const backgroundImage = computed(() => {
  if (props.msg.IsAdmin) {
    return config.value.Icon
  }
  if (props.msg.IsAnonymous) {
    return config.value.MessageDefaultIcon
  }
  return backgroundImg.value
})

onMounted(() => {
  let info = document.getElementById(`info${props.msg.ID}`) as HTMLSelectElement
  let br = document.createElement("br")
  info.appendChild(br)
  let div = document.createElement("div")
  div.classList.add("box")
  for (let i = 0; i < pics.length; i++) {
    let img = document.createElement("img")
    img.setAttribute("src", pics[i])
    img.addEventListener("click", () => {
      emit('emitPic', pics[i])
    })
    img.classList.add("insertPic")
    div.appendChild(img)
  }
  info.appendChild(div)
})

function deleteMsg() {
  axios.request("/yuurei/msg/" + props.msg.ID, "delete", {}).then((result) => {
    emit('reload')
  })
}
</script>


<style lang="less" scoped>

.messageBox {
  // float: left;
  width: 95%;
  margin: 8vh auto;
  margin-top: 3vh;
  position: relative;
  display: flex;
  //font-size: 1.7vh;
  justify-content: flex-end;

  // height: 100%;
  &::after,
  &::before {
    content: ""; //设置内容为空

    height: 0; //高度为0

    line-height: 0; //行高为0

    display: block; //将文本转为块级元素

    visibility: hidden; //将元素隐藏

    clear: both //清除浮动
  }


  .icon {
    //background-image: url("../../assets/ec8ce499ly1h1hhm1o6m1j22pg1ww4qp.jpg");
    background-size: cover;
    border-radius: 50%;
    border: .4vh solid rgb(49, 49, 49);
    float: left;
    width: 10vh;
    height: 10vh;
    background-color: rgb(49, 49, 49);
    margin-right: 4vh;
    position: relative;
    background-position: center;
    box-shadow: .4vh .4vh 1px 0 rgba(49, 49, 49, 0.1);

    @media (max-aspect-ratio: 1/1) {
      width: 15vw;
      height: 15vw;
      margin-right: 2vh;
      border: .6vw solid rgb(49, 49, 49);

    }

    .name {
      position: absolute;
      bottom: -3vh;
      width: 100%;
      font-size: 1.8vh;
      font-family: "幼圆";
      text-align: center;
      display: inline-block;
      //transform-origin: top;
      //text-wrap: avoid;
      height: 1.7vh;
      //word-break: keep-all;
      //line-height: 100%;
      text-shadow: -.15vh .15vh .1vh rgba(124, 101, 101, 0.2);
      font-weight: bold;
    }
  }

  .messageInfo {
    border: .5vh solid rgb(49, 49, 49);
    border-left: .6vh solid rgb(49, 49, 49);
    border-bottom: .6vh solid rgb(49, 49, 49);
    padding: 2vh;
    width: 90%;
    border-radius: 1vh;
    min-height: 15vh;
    background-color: #fff;
    word-break: break-word;
    white-space: pre-wrap;
    position: relative;
    top: 1vw;
    //white-space: pre-wrap;

    @media (max-aspect-ratio: 1/1) {
      min-height: 10vh;
      width: 75%;
      border: .5vh solid rgb(49, 49, 49);
      border-left: .6vh solid rgb(49, 49, 49);
      border-bottom: .6vh solid rgb(49, 49, 49);
    }
    @media (max-aspect-ratio: 12/16) {
      top: 1.5vh;
      border: .7vw solid rgb(49, 49, 49);
      border-left: 1vw solid rgb(49, 49, 49);
      border-bottom: 1vw solid rgb(49, 49, 49);
    }

    &::after {
      position: absolute;
      display: block;
      left: -1.1vh;
      top: 0vh;
      transform: rotate(225deg);
      content: '';
      width: 0;
      height: 0;
      border: 1.2vh solid transparent;
      border-top-color: rgb(49, 49, 49);
      border-right-color: rgb(49, 49, 49);
      //border-radius: 2px;
      z-index: -1;
      border-radius: .4vh;

      @media (max-aspect-ratio: 1/1) {
        border: 1.8vw solid transparent ;
        border-top-color: rgb(49, 49, 49);
        border-right-color: rgb(49, 49, 49);
        border-radius: .4vh;
        left: -1.6vw;
        //top: .1vh;
      }

    }

    .delete {
      cursor: pointer;
      position: absolute;
      bottom: -4vh;
      right: 23vh;
      //text-shadow: -1px 1px 2px rgb(49,49,49);
      text-shadow: -1px 1px 5px rgba(49, 49, 49, .4);

      text-align: right;
      color: rgb(49, 49, 49)
    }

    .reply {
      cursor: pointer;
      position: absolute;
      bottom: -4vh;
      left: 0;
      //background-color: var(--sec-color);
      background-color: var(--fourth-color);
      //text-shadow: -1px 1px 2px rgb(49,49,49);
      //text-shadow: -1px 1px 10px rgba(49, 49, 49, .4);
      //padding:1vh;
      padding-left: 1vh;
      padding-right: 1vh;
      border-radius: .5vh;
      font-size: .9em;
      text-align: right;
      //color: rgb(49, 49, 49)
      color: white;
    }

    .date {
      position: absolute;
      bottom: -4vh;
      right: 2px;
      text-shadow: -1px 1px 5px rgba(164, 74, 74, 0.4);
      text-align: right;
      //color: rgb(49, 49, 49);
      color: white;
      padding: 0px 5px;
      border-radius: 2px;
      //background-color: rgba(49, 49, 49, 0.76);
      background-color: var(--sec-color);
      box-shadow: .4vh .4vh .1vh 0 rgba(49, 49, 49, 0.2);
      white-space: nowrap;

    }


    .info {
      font-family: "微软雅黑";
      font-size: 1em;
      font-weight: 600;
      color: rgb(40, 40, 40);
      padding: 5px;


      pre {
        font-family: "微软雅黑";
        white-space: pre-line;
      }
    }
  }
}

</style>

<style lang="less">
.insertPic {
  user-select: none;
  max-width: 50vw;
  max-height: 40vh;
  margin-top: 20px;
  margin-right: 20px;
  align-self: flex-start;
  cursor: pointer;

  @media (max-aspect-ratio: 1/1) {
    max-width: 90%;
    max-height: 80vh;

  }
}

.box {
  display: flex;
  flex-wrap: wrap;
}

.admin {
  color: var(--button-color)
}

</style>