<template>
  <div class="CommentContainer">
    <div class="left">
      <div class="icon" :style="{backgroundImage:`url(${comment.IsAdmin?config.Icon:comment.Icon})`}"></div>
    </div>
    <div class="right">
      <div class="header"><span style="font-weight: bold">{{comment.IsAdmin?config.MyName:comment.ArticleCommentUserName }}</span> <span style="color:darkgray;margin-left: 1vh">{{new Date(comment.CreatedAt).Format("yyyy/M/dd hh:mm:ss")}}</span></div>
      <div class="info" v-html="comment.ArticleCommentContent"></div>
      <div class="footer">
        <span @click="showReply=!showReply">{{ showReply ? `收起` : `回复` }}</span>&nbsp;&nbsp;<span v-show="userInfo.user.IsAdmin" @click="del">删除</span>
      </div>
    </div>
  </div>
  <n-message-provider>
  <send-comment @send="send" :comment="comment" v-show="showReply" class="comment" father-id="" article=""/>
  </n-message-provider>
  <div class="child">
    <child-comment @reload="reload" v-for="item in list" :key="item.ID" :comment="item"/>
  </div>
</template>

<script lang="ts" setup>

import SendComment from "./SendComment.vue"
import ChildComment from "./ChildComment.vue"
import {inject, ref, watch} from "vue";
import axios from "@/request";
import {useComment} from "@/store/commentData";
import {useUserInfo} from "@/store/UserInfo";
import {NMessageProvider} from "naive-ui";

const config = inject("globalConfig") as any

const commentData = useComment()
const userInfo = useUserInfo()
const showReply = ref(false)

const props = defineProps<{
  comment: any
}>()
const emit = defineEmits(['reloaded'])
const list = ref(props.comment.Comments)


watch(props,(newvalue,oldvalue)=>{
  list.value = newvalue.comment.Comments
})
function del(){
  axios.request("/yuurei/articleComment/"+props.comment.ID,"delete",{}).then((res:any)=>{
    console.log(res)
    if(res.code==200){
      emit('reloaded')
    }
  })
}

function reload(){
  axios.get("/yuurei/articleComment/"+props.comment.ID,{}).then((res:any)=>{
    list.value = res.data.data.Comments
  })
}
function send(args: any) {
  axios.request("/yuurei/articleComment", "post", {
    ArticleCommentContent: `<a>@${props.comment.IsAdmin?config.value.MyName:props.comment.ArticleCommentUserName}&nbsp;&nbsp;</a>`+ args.msg.replace(/\r/ig, '').replace(/\n/ig, '<br/>'),
    ArticleCommentUserName: commentData.commentData.username,
    ArticleID: props.comment.ArticleID,
    UserID: userInfo.user.ID,
    FatherID: props.comment.ID,
    Icon: args.icon,
    IsAdmin:userInfo.user.IsAdmin,

  }).then((res:any) => {
    if(res.code==200){
      showReply.value = false
      axios.get("/yuurei/articleComment/"+props.comment.ID,{}).then((res:any)=>{
        list.value = res.data.data.Comments
      })
    }
  })
}
</script>

<style lang="less" >
.CommentContainer {
  display: flex;
  justify-content: flex-start;
  width: 70%;
  margin-left: 8vh ;
  margin-bottom: 2vh;
  margin-top:5vh;
  //animation
  //margin-top: 5vh;
  animation:enter 1s ease;

  @media (max-aspect-ratio: 9/16) {
    width: 95%;
    margin:0 auto;
    margin-bottom: 2vh;
    margin-top:5vh;
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
      background-position: center;
      background-size: cover;

      @media (max-aspect-ratio: 9/16) {
        width: 8vh;
        height: 8vh;
      }
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

.child {
  width: 70%;
  //min-height: 40vh;
  margin-left: 6vh;
  //margin-bottom: 6vh;
  @media (max-aspect-ratio: 9/16) {
    float: right;
    width: 90%;
    margin: 0 auto;

  }
}

</style>