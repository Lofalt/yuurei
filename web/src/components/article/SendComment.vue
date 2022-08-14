<template>
  <div class="commentContainer">
    <div class="header">
      <input v-model="commentData.commentData.username" class="inputName" type="text" placeholder="*昵称"/>
      <div @click="showModal=true" class="icon" :style="{backgroundImage:`url(${bgi})`}"></div>
    </div>
    <textarea class="text-area" v-model="msgContent" placeholder="说点啥吧"></textarea>
    <div class="footer">
      <!--      😳-->
      <button @click="send">发送</button>
    </div>
    <n-modal v-model:show="showModal" display-directive="show">
      <n-card class="card">
        <upload-pic @confirm="accept" name="上传头像" raw-src="" directory="commentIcon" quality="30" ratio="1"/>
      </n-card>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import {computed, inject, ref, watch} from 'vue'
import {useComment} from "@/store/commentData";
import {NModal, NCard} from "naive-ui"
import UploadPic from "../file/UploadPic.vue"
import {useMessage} from "naive-ui"
import {useUserInfo} from "@/store/UserInfo";

const userInfo = useUserInfo()
const message = useMessage()
const commentData = useComment()
const username = ref('')
const msgContent = ref('')
const icon = ref('')
const showModal = ref(false)
const config = inject('globalConfig') as any
const props = defineProps<{
  article: object;
  fatherId: number;
  comment: object,
}>()
const emit = defineEmits(['send'])
const bgi = computed(() => {
  return commentData.commentData.icon === "" ? config.value.MessageDefaultIcon : commentData.commentData.icon
})

function accept(url: string) {
  commentData.commentData.icon = url
  showModal.value = false
}

watch(username, (newValue: any, oldValue: any) => {
  commentData.commentData.username = newValue
})

function send() {
  if(userInfo.user.Sended  && !userInfo.user.IsAdmin ){
    message.warning("发言过快,30秒后重试")
    return
  }
  if(commentData.commentData.username=="" && !userInfo.user.IsAdmin ){
    message.warning("昵称不能为空")
    return
  }
  if(msgContent.value==""){
    message.warning("先说点啥吧")
    return
  }
  emit('send', {
    msg: msgContent.value,
    icon: bgi.value
  })
  msgContent.value=""
  userInfo.user.Sended = true
  setTimeout(()=>{
    userInfo.user.Sended = false
  },30000)
}
</script>

<style lang="less" scoped>
.commentContainer {

  display: flex;
  flex-direction: column;
  width: 70%;
  margin-left: 8vh ;
  margin-bottom: 2vh;

  @media (max-aspect-ratio: 9/16) {
    width: 95%;
    margin: 0 auto;
    margin-bottom:2vh;


  }

  .header {
    display: flex;
    justify-content: space-between;
    align-items: flex-end;

    .inputName {
      padding-left: 2vh;
      height: 5vh;
      border-radius: .5vh;
      outline: none;
      border: .3vh solid var(--sec-color);
      //outline: .2vh solid var(--sec-color);
    }

    .icon {
      width: 10vh;
      height: 10vh;
      background-color: black;
      border-radius: 50%;
      background-position: center;
      background-size: cover;
      cursor: pointer;
    }
  }

  .text-area {
    height: 20vh;
    resize: none;
    margin-top: 1vh;
    outline: none;
    padding: 2vh;
    font-size: 1.1em;
    //outline: .2vh solid var(--sec-color);
    border-radius: .5vh;
    border: .3vh solid var(--sec-color);
    //border-color: darkgray;
  }

  .footer {
    display: flex;
    margin-top: 1vh;

    button {
      background-color: var(--button-color);
      color: white;
      padding: 1vh 2vh;
      outline: none;
      border: none;
      border-radius: .4vh;
      //margin-bottom: 1vh;
    }
  }


}

.card {
  width: 50%;

  @media (max-aspect-ratio: 1/1) {
    width: 90%;
  }
}
</style>