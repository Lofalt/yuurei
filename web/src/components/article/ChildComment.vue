<template>
  <div class="childCommentContainer">
    <div class="left">
      <div class="icon" :style="{backgroundImage:`url(${comment.IsAdmin?config.Icon:comment.Icon})`}"></div>
    </div>
    <div class="right">
<!--      <div class="header">{{comment.ArticleCommentUserName}} 发表于 2022/8/16 22:22:22</div>-->
      <div class="header"><span style="font-weight: bold">{{ comment.IsAdmin?config.MyName:comment.ArticleCommentUserName }}</span> <span style="color:darkgray;margin-left: 1vh">{{new Date(comment.CreatedAt).Format("yyyy/M/dd hh:mm:ss")}}</span></div>
      <div class="info" v-html="comment.ArticleCommentContent"></div>
      <div class="footer">
        <span @click="showReply=!showReply">{{ showReply ? `收起` : `回复` }}</span>&nbsp;&nbsp;<span v-show="userInfo.user.IsAdmin" @click="del">删除</span>
      </div>
    </div>
  </div>
  <n-message-provider>
  <send-comment @send="send" :comment="comment" v-if="showReply" class="comment" father-id="" article=""/>
  </n-message-provider>

</template>

<script lang="ts" setup>

import SendComment from "./SendComment.vue"
import {inject, ref} from "vue";
import axios from "@/request";
import {useComment} from "@/store/commentData";
import {useUserInfo} from "@/store/UserInfo";
import {NMessageProvider} from "naive-ui";

const config = inject("globalConfig") as any
const commentData = useComment()
const userInfo = useUserInfo()

const showReply = ref(false)
const emit = defineEmits(['reload'])
const props = defineProps<{
  comment: any
}>()
function del(){
  axios.request("/yuurei/articleComment/"+props.comment.ID,"delete",{}).then((res)=>{
    console.log(res)
    emit('reload')

  })
}

function send(args: any) {
  axios.request("/yuurei/articleComment", "post", {
    ArticleCommentContent: `<a>@${props.comment.IsAdmin?config.value.MyName:props.comment.ArticleCommentUserName}&nbsp;&nbsp;</a>`+ args.msg.replace(/\r/ig, '').replace(/\n/ig, '<br/>'),
    ArticleCommentUserName: commentData.commentData.username,
    ArticleID: props.comment.ArticleID,
    UserID: userInfo.user.ID,
    FatherID: props.comment.FatherID,
    Icon: args.icon,
    IsAdmin:userInfo.user.IsAdmin,
  }).then((res:any) => {
    if(res.code==200){
      showReply.value = false
      emit('reload')
    }
  })
}
</script>

<style lang="less" >
.childCommentContainer {
  display: flex;
  justify-content: flex-start;
  width: 90%;
  margin: 2vh auto;
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
      width: 7vh;
      height: 7vh;
      background-color: black;
      border-radius: 50%;
      background-position: center;
      background-size: cover;

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
      //padding-left: 1vh;
      //border-bottom: .1vh dashed darkgray;
      //padding-bottom:1vh;
    }

    .info {
      //border-top: .2vh solid var(--sec-color);
      margin-top: 2vh;
      //padding-top: 3vh;
      //padding-left: 3vh;
      //padding-right: 2vh;
      //background-color: var(--third-color);
      color: #343434;
      //min-height: 20vh;
      border-radius: 1vh;
      //font-weight: bold;
      a{
        color:var(--button-color);
      }
    }

    .footer {
      span {
        //padding: 1vh 2vh;
        border: none;
        //color:white;
        //outline: none;
        //border-radius: .5vh;
        //background-color: var(--sec-color);
        margin-top: 2vh;
        color:var(--sec-color);
        display: inline-block;
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