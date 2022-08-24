<template>
  <div class="container">
    <div class="icon" :style="{backgroundImage:`url(${src})`}" @click="upload"></div>
    <div class="right">
      <textarea class="inputText" draggable="false" :placeholder="`回复:@${UserName}`" v-model="msgContent"></textarea>
      <div class="foot">
        <input v-model="username" type="text" placeholder="输入昵称" class="input" @change="changeName"/>
        <button class="send" @click="send">回复</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {inject, ref, watch} from 'vue'
import axios from "@/request/index"
import {useMessage} from "naive-ui"
import {useUserInfo} from "@/store/UserInfo";

const userInfo = useUserInfo()
const message = useMessage()
const emit = defineEmits(['send', 'upload','change-name'])

const props = defineProps<{
  ID: number,
  UserName: string,
}>()

const src = inject('src')
const msgContent = ref("")
const username = ref("")

function upload() {
  console.log("send")
  emit('upload')
}

watch(username,()=>{
  emit('change-name',username.value)

})
function changeName(){
  emit('change-name',username.value)
}

function send() {
  if(userInfo.user.Sended){
    message.info("发言间隔为30秒")
    return
  }
  if(msgContent.value==""){
    message.info("先说点啥吧")
    return
  }
  axios.request("/yuurei/msg", "post", {
    MessageContent:`<a>@${props.UserName}&nbsp;&nbsp;</a>`+ msgContent.value.replace(/\r/ig, '').replace(/\n/ig, '<br/>'),
    FatherID: props.ID,
    Icon: src.value,
    UserName:username.value,
    IsAnonymous:username.value=="" && !userInfo.user.IsAdmin,
    IsAdmin:userInfo.user.IsAdmin,
    UserID:userInfo.user.ID,
  }).then((res) => {
    console.log(res)
    emit('send')
    username.value=""
    msgContent.value=""
    userInfo.user.Sended = true
    setTimeout(()=>{
      userInfo.user.Sended = false
    },30000)
  })
}

</script>

<style scoped lang="less">

.container {
  width: 95%;
  display: flex;
  flex-direction: row;
  align-items: start;
  justify-content: flex-end;
  margin: 0 auto;

  .icon {
    //flex:2;
    width: 10vh;
    height: 10vh;
    background-position: center;
    background-size: cover;
    border-radius: 50%;
    margin-right: 2vh;
  }

  .right {
    display: flex;
    flex-direction: column;
    align-items: flex-end;
    width: 80%;

    @media (max-aspect-ratio: 9/16){
      width: 70%;

    }
    .foot{
      width: 100%;
      display: flex;
      margin-top:.5vh;
      flex-direction: row;
      justify-content: space-between;
      align-items: flex-end;

      .input{
        height: 4vh;
        outline: none;
        width: 20vh;
        border-radius: .5vh;
        border:.4vh solid black;
        padding-left:2vh;

        @media (max-aspect-ratio: 9/16){
          width: 15vh;
        }

      }
      .send {
        //flex:2;
        font-size: 1em;
        cursor: pointer;
        border: none;
        outline: none;
        padding: 0.5vh 1vh;
        color: white;
        background-color: var(--button-color);
        border-radius: .5vh;
        padding-left: 2vh;
        padding-right: 2vh;
        margin: 1vh;
        margin-bottom: 0;
      }
    }
    .inputText {
      height: 16vh;
      width: 100%;
      resize: none;
      outline: none;
      //float: right;
      border: 0;
      padding: 2vh;
      border-radius: 1vh;
      font-family: 微软雅黑;
      font-size: 1em;
      //position: absolute;
      //bottom: 0;
      //left: 0;
      border: .4vh solid rgba(49, 49, 49, 1);
      box-sizing: border-box;

      @media (max-aspect-ratio: 9/16) {
        //height: 57vw;
        width: 100%;

        //border-top: .4vw solid rgba(49, 49, 49, .1);
      }

    }


  }

}

</style>