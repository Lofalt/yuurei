<template>
  <div class="container">
    <div class="left">
      <div class="icon"></div>
    </div>
    <div class="right">
      <div class="header">匿名具足虫 发表于 2022/8/16 22:22:22</div>
      <div class="info">阿迪阿塞啊岁大赛阿迪阿塞啊岁大赛阿迪阿塞啊岁大赛阿迪阿塞啊岁大赛阿迪阿塞啊岁大赛<br>阿迪阿塞啊岁大赛阿迪阿塞啊岁大赛阿迪阿塞啊岁大赛</div>
      <div class="footer">
        <button @click="showReply=!showReply">{{ showReply ? `收起` : `回复` }}</button>
      </div>
    </div>
  </div>
  <send-comment :comment="comment" v-show="showReply" class="comment" father-id="" article=""/>

</template>

<script lang="ts" setup>

import SendComment from "./SendComment.vue"
import {ref} from "vue";

const showReply = ref(false)

const props = defineProps<{
  comment: object
}>()
</script>

<style lang="less" scoped>
.container {
  display: flex;
  justify-content: flex-start;
  width: 90%;
  margin: 5vh auto;
  float: right;
  margin-bottom: 2vh;
  //margin-top: 5vh;
  @media (max-aspect-ratio: 9/16) {
    width: 95%;
    margin-right: 0;
  }

  .left {
    display: flex;
    flex-direction: column;
    justify-content: flex-start;

    .icon {
      width: 8vh;
      height: 8vh;
      background-color: black;
      border-radius: 50%;

    }
  }

  .right {
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    margin-left: 2vh;
    width: 80%;

    .header {
      display: block;
      padding-left: 1vh;
    }

    .info {
      border-top: .2vh solid var(--third-color);
      margin-top: 2vh;
      padding-top: 3vh;
      padding-left: 3vh;
      padding-right: 2vh;
      background-color: var(--third-color);
      min-height: 20vh;
      border-radius: 2vh;
    }

    .footer {
      button {
        padding: 1vh 2vh;
        border: none;
        outline: none;
        border-radius: 1vh;
        background-color: var(--third-color);
        margin-top: 2vh;
        cursor: pointer;
      }
    }
  }

}

.comment {
  width: 60%;
  @media (max-aspect-ratio: 9/16) {
    width: 80%;

  }
}

</style>