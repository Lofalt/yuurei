<template>
  <div class="messageBox">
    <div class="icon" :style="{backgroundImage:`url(${backgroundImg})`}">
      <div class="name">{{ msg.UserName }}</div>
    </div>
    <div class="messageInfo">
      <div v-show="userInfo.user && userInfo.user.IsAdmin " class="delete" @click="deleteMsg">删除</div>
      <div class="date">{{
          new Date(msg.CreatedAt).toLocaleDateString() + " \n\n\n" + new Date(msg.CreatedAt).toLocaleTimeString()
        }}
      </div>
      <span></span>
      <div class="info" v-html="msg.MessageContent">
        </div>
<!--        <div class="info">-->
<!--          <pre>{{msg.MessageContent}}</pre>-->

    </div>
  </div>

</template>

<script lang="ts" setup>
import {ref} from 'vue'
import axios from '@/request/index'
import {useUserInfo} from "@/store/UserInfo";

const backgroundImg = ref("")
const userInfo = useUserInfo()
const emit = defineEmits(['reload'])

const props = defineProps<{
  msg: object
}>()

backgroundImg.value = props.msg.Icon

function deleteMsg() {
  axios.request("/yuurei/msg/" + props.msg.ID, "delete", {}).then((result) => {
    console.log(result)
    emit('reload')
  })
}
</script>


<style lang="less" scoped>

.messageBox {
  // float: left;
  width: 90%;
  margin: 8vh auto;
  position: relative;

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
    background-image: url("../../assets/ec8ce499ly1h1hhm1o6m1j22pg1ww4qp.jpg");
    background-size: cover;
    border: 3px solid rgb(49, 49, 49);
    float: left;
    width: 6vw;
    height: 6vw;
    background-color: rgb(49, 49, 49);
    margin-right: 20px;
    position: relative;
    background-position: center;
    box-shadow: -4px 4px 1px 0 rgba(49, 49, 49, 0.2);

    @media (max-width: 1024px) {
      width: 8vh;
      height: 8vh;
    }

    .name {
      position: absolute;
      bottom: -30px;
      width: 100%;
      text-align: center;
      display: inline-block;
      //transform-origin: top;
      //text-wrap: avoid;
      height: 15px;
      //word-break: keep-all;
      //line-height: 100%;
      text-shadow: -1px 1px 1px rgb(180, 157, 157);

    }
  }

  .messageInfo {
    border: 4px solid rgb(49, 49, 49);
    border-left: 6px solid rgb(49, 49, 49);
    border-bottom: 6px solid rgb(49, 49, 49);
    padding: 1vw;
    float: left;
    width: 80%;
    border-radius: 3px;
    min-height: 15vh;
    // background-color: aquamarine;
    position: relative;

    .delete {
      cursor: pointer;
      position: absolute;
      bottom: -40px;
      right: 160px;
      //text-shadow: -1px 1px 2px rgb(49,49,49);
      text-shadow: -1px 1px 5px rgba(49, 49, 49, .4);

      text-align: right;
      color: rgb(49, 49, 49)
    }

    .date {
      position: absolute;
      bottom: -40px;
      right: 2px;
      text-shadow: -1px 1px 5px rgba(164, 74, 74, 0.4);
      text-align: right;
      //color: rgb(49, 49, 49);
      color: white;
      padding: 0px 5px;
      border-radius: 2px;
      background-color: rgba(49, 49, 49, 0.76);
      box-shadow: -4px 4px 1px 0 rgba(49, 49, 49, 0.2);
    }

    @media (max-width: 1024px) {
      min-height: 10vh;
      width: 70%;
      border: 3px solid rgb(49, 49, 49);
      border-left: 5px solid rgb(49, 49, 49);
      border-bottom: 5px solid rgb(49, 49, 49);

    }

    &::after {
      position: absolute;
      display: block;
      left: -15.5px;
      top: 4px;
      transform: rotate(225deg);
      content: '';
      width: 0;
      height: 0;
      border: 10px solid transparent;
      border-top-color: rgb(49, 49, 49);
      border-right-color: rgb(49, 49, 49);
      border-radius: 2px;

      @media (max-width: 1024px) {
        border: 8px solid transparent ;
        border-top-color: rgb(49, 49, 49);
        border-right-color: rgb(49, 49, 49);
        border-radius: 2px;
        left: -12.5px;
      }
    }

    &::before {
      position: absolute;
      display: block;
      left: -8px;
      // top: 5px;
      z-index: 10;
      top: 6px;
      transform: rotate(225deg);
      content: '';
      width: 0;
      height: 0;
      border: 8px solid transparent;
      border-top-color: white;
      border-right-color: white;
      @media (max-width: 1024px) {

        left: -7px;
        top: 4px;

      }

    }

    .info {
      font-family: "微软雅黑";
      font-size: 1.2em;
      font-weight: 600;
      color: rgb(98, 95, 95);
      padding: 5px;

      pre{
        font-family: "微软雅黑";
        white-space: pre-line;
      }
    }
  }
}
</style>