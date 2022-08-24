<template>
  <div class="replyContainer">
    <div class="replyIcon" :style="{backgroundImage:`url(${backgroundImage})`}">
      <div :class="child.IsAdmin?`admin`:``" class="name">{{ name }}</div>
      <!--      :class="msg.IsAdmin?`admin`:``"-->
    </div>
    <div class="reply">
      <!--      {{child.MessageContent}}-->
      <div v-show="userInfo.user && userInfo.user.IsAdmin " class="delete" @click="deleteMsg">删除</div>
      <div class="replys" @click="showReply=!showReply">{{ !showReply ? `回复` : `收起` }}</div>
      <div class="date">
        {{ new Date(child.CreatedAt).Format("yyyy/M/dd") + " \n\n\n" + (new Date(child.CreatedAt)).Format("hh:mm:ss") }}
      </div>
      <div class="content" v-html="child.MessageContent"></div>
    </div>
  </div>
  <n-message-provider>
    <send-reply @send="accept" v-show="showReply" :UserName="name" :ID="child.FatherID"/>
  </n-message-provider>
</template>

<script lang="ts" setup>
import {useUserInfo} from "@/store/UserInfo";
import axios from "@/request";
import {computed, inject, ref} from "vue";
import SendReply from "./SendReply.vue"

const config = inject('globalConfig') as any
const userInfo = useUserInfo()
const showReply = ref(false)
const emit = defineEmits(['reload'])
import {NMessageProvider} from "naive-ui"

const backgroundImage = computed(() => {
  return props.child.IsAnonymous ? config.value.MessageDefaultIcon : (props.child.IsAdmin ? config.value.Icon : props.child.Icon)
})

const name = computed(() => {
  return props.child.IsAnonymous ? config.value.AnonymousName : (props.child.IsAdmin ? config.value.MyName : props.child.UserName)
})

function deleteMsg() {
  axios.request("/yuurei/msg/" + props.child.ID, "delete", {}).then((result) => {
    emit('reload')
  })
}

function accept() {
  emit('reload')
  showReply.value = false
}

const props = defineProps<{
  child: any
}>()

</script>

<style lang="less">

.replyContainer {
  font-family: "微软雅黑";
  --replyColor: rgb(63, 61, 61);
  color: var(--replyColor);

  //float: right;
  width: 95%;
  min-height: 15vh;
  display: flex;
  justify-content: flex-end;
  margin: 4vh auto;

  .replyIcon {
    width: 10vh;
    height: 10vh;
    background-color: #499ba9;
    background-position: center;
    background-size: cover;
    float: right;
    border-radius: 50%;
    margin-right: 2vh;
    position: relative;
    //border: .3vh solid var(--replyColor);
    border: .4vh solid rgb(49, 49, 49);

    @media (max-aspect-ratio: 9/16) {
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

  .replys {
    cursor: pointer;
    position: absolute;
    bottom: -4vh;
    left: 0vh;
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
    //text-shadow: -1px 1px 5px rgba(164, 74, 74, 0.4);
    text-align: right;
    //color: rgb(49, 49, 49);
    color: white;
    padding: 0px 5px;
    border-radius: 2px;
    background-color: var(--sec-color);
    //box-shadow: .4vh .4vh .1vh 0 rgba(49, 49, 49, 0.2);
    white-space: nowrap;

  }

  .reply {
    padding: 2vh;
    height: 100%;
    float: right;
    width: 80%;
    min-height: 13vh;
    background-color: #fff;
    word-break: break-word;
    white-space: pre-wrap;
    border: .4vh solid var(--replyColor);
    border-left: .5vh solid var(--replyColor);
    border-bottom: .5vh solid var(--replyColor);
    border-radius: 1vh;
    position: relative;


    @media (max-aspect-ratio: 9/16) {
      width: 66%;
      min-height: 11vh;
      border: .6vw solid var(--replyColor);
      border-left: .8vw solid var(--replyColor);
      border-bottom: .8vw solid var(--replyColor);
    }
    //&::after {
    //  position: absolute;
    //  display: block;
    //  left: -1.1vh;
    //  top: 0vh;
    //  transform: rotate(225deg);
    //  content: '';
    //  width: 0;
    //  height: 0;
    //  border: 1.2vh solid transparent;
    //  border-top-color: var(--replyColor);
    //  border-right-color: var(--replyColor);
    //  //border-radius: 2px;
    //  z-index: -1;
    //  border-radius: .4vh;
    //
    //  @media (max-aspect-ratio: 1/1) {
    //    border: 1.2vh solid transparent ;
    //    border-top-color: var(--replyColor);
    //    border-right-color: var(--replyColor);
    //    border-radius: .4vh;
    //    left: -1.1vh;
    //    top: .1vh;
    //  }
    //
    //}
    .delete {
      cursor: pointer;
      position: absolute;
      bottom: -4vh;
      right: 25vh;
      //text-shadow: -1px 1px 2px rgb(49,49,49);
      text-shadow: -1px 1px 5px rgba(49, 49, 49, .4);

      text-align: right;
      color: rgb(49, 49, 49)
    }

    .content {
      color: var(--replyColor);
      font-weight: bold;

      a {
        color: var(--sec-color)
      }
    }
  }

}

</style>