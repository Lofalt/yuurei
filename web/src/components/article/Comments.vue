<template>
  <div class="articleComments">
    <div class="header">
      <h3>评论区</h3>
    </div>
    <n-message-provider>
      <send-comment :username="`anonymous`" @send="send"/>
    </n-message-provider>
<!--    <transition-group  tag="p" name="list" >-->
      <comment @reloaded="reload" v-for="(item,index) in list" :key="item.ID" :comment="item"/>
<!--    </transition-group>-->
  </div>
</template>

<script lang="ts" setup>
import SendComment from "./SendComment.vue"
import axios from "../../request/index"
import Comment from "./Comment.vue"
import {useComment} from "@/store/commentData";
import {useUserInfo} from "@/store/UserInfo";
import {ref, watch} from "vue";
import {NMessageProvider} from "naive-ui";


const userInfo = useUserInfo()
const props = defineProps<{
  article: any
}>()
const commentData = useComment()
const list = ref([]) as any


watch(props, (newvalue, oldvalue) => {
  list.value = newvalue.article.ArticleComments
})
function reload(){
  axios.get("/yuurei/article/" + props.article.ID, {}).then((res: any) => {
    list.value = res.data.data.ArticleComments
  })
}

function send(args: any) {
  axios.request("/yuurei/articleComment", "post", {
    ArticleCommentContent:  args.msg.replace(/\r/ig, '').replace(/\n/ig, '<br/>'),
    ArticleCommentUserName: commentData.commentData.username,
    ArticleID: props.article.ID,
    UserID: userInfo.user.ID,
    Icon: args.icon,
    IsAdmin:userInfo.user.IsAdmin,

  }).then((res: any) => {
    if (res.code == 200) {
      axios.get("/yuurei/article/" + props.article.ID, {}).then((res: any) => {
        list.value = res.data.data.ArticleComments
      })
    }
  })
}
</script>

<style lang="less" scoped>
.articleComments {
  font-size: 1.6vh;
  font-family: "微软雅黑";

  .header{
    width: 70%;
    margin-left:8vh;
    position: relative;
    padding-left:2vh;
    margin-top:5vh;
    @media (max-aspect-ratio: 12/16){
      width: 95%;
      margin:0 auto;

    }
    &::before {
      left: 0;
      border-radius: 1vh;
      content: " ";
      width: .8vh;
      height: 100%;
      display: inline-block;
      position: absolute;
      background-color: var(--button-color);
    }
  }

}


.list-enter-active,
.list-leave-active {
  transition: all 0.5s ease;
}
.list-enter-from,
.list-leave-to {
  opacity: 0;
  transform: translateX(30px);
}


</style>