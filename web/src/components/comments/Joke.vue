<template>
  <div class="messageBox">
    <div class="icon" :style="{backgroundImage:`url(${backgroundImage})`}">
      <div class="name" :class="msg.IsAdmin?`admin`:``">{{ userName }}</div>
    </div>
    <div class="messageInfo">
      <div v-show="userInfo.user && userInfo.user.IsAdmin " class="delete" @click="deleteMsg">删除</div>
      <div class="date">{{
          new Date(msg.CreatedAt).Format("yyyy/M/dd") + " \n\n\n" + (new Date(msg.CreatedAt)).Format("hh:MM:ss")
        }}
      </div>
      <span></span>
      <!--      <div class="info" v-html="msg.MessageContent">-->
      <div :id="`info${msg.ID}`" class="info" v-html="msg.MessageContent">
      </div>
      <!--        <div class="info">-->
      <!--          <pre>{{msg.MessageContent}}</pre>-->

    </div>
  </div>

</template>

<script lang="ts" setup>
import {computed, inject, onMounted, ref} from 'vue'
import axios from '@/request/index'
import {useUserInfo} from "@/store/UserInfo";
import {isMobile} from "@/util/utils";

const flag = ref(isMobile())
const backgroundImg = ref("")
const userInfo = useUserInfo()
const emit = defineEmits(['reload', 'emitPic'])
const content = ref("")
const props = defineProps<{
  msg: any
}>()
let pics = props.msg.Pics.split(",")
// let imgs = ""
// for (let i = 0; i < pics.length; i++) {
//   imgs = imgs + `<img src="${pics[i]}" class="insertPic" />`
// }
// content.value = props.msg.MessageContent + `<br/>` + `<div style="display: flex;flex-wrap: wrap">` + imgs + `</div>`

function open(url: string) {
  window.open(url)
}

const config = inject('globalConfig') as any
backgroundImg.value = props.msg.Icon

const userName = computed(() => {
  if (props.msg.IsAnonymous) {
    console.log(config.value)
    return config.value.AnonymousName
  }
  return props.msg.UserName

})

const backgroundImage = computed(() => {
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
  font-size: 1.7vh;

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
    width: 6vw;
    height: 6vw;
    background-color: rgb(49, 49, 49);
    margin-right: 4vh;
    position: relative;
    background-position: center;
    box-shadow: .4vh .4vh 1px 0 rgba(49, 49, 49, 0.1);

    @media (max-aspect-ratio: 1/1) {
      width: 8vh;
      height: 8vh;
      margin-right: 2vh;
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
    float: left;
    width: 90%;
    border-radius: 1vh;
    min-height: 15vh;
    //white-space: pre-line;
    // background-color: aquamarine;
    word-break: break-word;
    white-space: pre-wrap;
    position: relative;
    top: 2vw;
    //white-space: pre-wrap;
    @media (max-aspect-ratio: 9/16) {
      top: 3vh
    }

    .delete {
      cursor: pointer;
      position: absolute;
      bottom: -4vh;
      right: 17vh;
      //text-shadow: -1px 1px 2px rgb(49,49,49);
      text-shadow: -1px 1px 5px rgba(49, 49, 49, .4);

      text-align: right;
      color: rgb(49, 49, 49)
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
      background-color: rgba(49, 49, 49, 0.76);
      box-shadow: .4vh .4vh .1vh 0 rgba(49, 49, 49, 0.2);
      white-space: nowrap;

    }

    @media (max-aspect-ratio: 1/1) {
      min-height: 10vh;
      width: 75%;
      border: .3vh solid rgb(49, 49, 49);
      border-left: .5vh solid rgb(49, 49, 49);
      border-bottom: .5vh solid rgb(49, 49, 49);

    }

    &::after {
      position: absolute;
      display: block;
      left: -1.3vh;
      top: .2vh;
      transform: rotate(225deg);
      content: '';
      width: 0;
      height: 0;
      border: 1.2vh solid transparent;
      border-top-color: rgb(49, 49, 49);
      border-right-color: rgb(49, 49, 49);
      //border-radius: 2px;
      border-radius: .4vh;

      @media (max-aspect-ratio: 1/1) {
        border: .8vh solid transparent ;
        border-top-color: rgb(49, 49, 49);
        border-right-color: rgb(49, 49, 49);
        border-radius: .3vh;
        left: -1vh;
      }
    }

    //&::before {
    //  position: absolute;
    //  display: block;
    //  left: -1vh;
    //  // top: 5px;
    //  z-index: 10;
    //  top: 6px;
    //  transform: rotate(225deg);
    //  content: '';
    //  width: 0;
    //  height: 0;
    //  border: .9vh solid transparent;
    //  border-top-color: white;
    //  border-right-color: white;
    //  @media (max-aspect-ratio: 1/1) {
    //
    //    left: -7px;
    //    top: 4px;
    //
    //  }
    //  @media (min-width: 2200px) {
    //    left: -7px;
    //  }
    //
    //}

    .info {
      font-family: "微软雅黑";
      font-size: 1.2em;
      font-weight: 600;
      color: rgb(98, 95, 95);
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
  max-height: 60vh;
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