<template>
  <div class="article-comments">
    <div>
      <h1>
        发表评论
      </h1>
    </div>
    <send-comment :username="`anonymous`" @send="send"/>
    <comment v-for="item in article.ArticleComments" :key="item.ID" :comment="item"/>
  </div>
</template>

<script lang="ts" setup>
import SendComment from "./SendComment.vue"
import axios from "../../request/index"
import Comment from "./Comment.vue"
import {useComment} from "@/store/commentData";
import {useUserInfo} from "@/store/UserInfo";

const userInfo = useUserInfo()
const props = defineProps<{
  article: any
}>()
const commentData = useComment()

function send(args: any) {
  axios.request("/yuurei/articleComment", "post", {
    ArticleCommentContent: args.msg,
    ArticleCommentUserName: commentData.commentData.username,
    ArticleID: props.article.ID,
    UserID: userInfo.user.ID,
    Icon: args.icon,
  }).then((res) => {
    console.log(res)
  })
}
</script>

<style lang="less" scoped>
.article-comments {

}
</style>