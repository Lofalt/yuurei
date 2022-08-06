<template>
  <div class="panel" @click="archive(article.ID)">
    <!--        <div class="background" :style="{backgroundImage:background}">-->
    <div class="background" :style="{backgroundImage:`url(${article.ArticleCategory.HeaderPic})`}">
      <!--            <img src="~assets/img/backgroundimg/background1.png"/>-->
    </div>
    <div class="panelTitle">
      <a>{{ article.ArticleTitle }}</a>
    </div>
    <div class="panelSummary">
      {{ article.ArticleSummary }}
    </div>
    <div class="panelInfo">
      <span>{{ new Date(article.CreatedAt).toLocaleDateString() }}</span>
      <!--                        <i class="el-icon-view"/> <span>阅读 {{article.articleReadTimes}}</span>-->
      <!--                        <span>❤ 点赞 {{article.articleLoveTimes}}次</span>-->
      <i class="el-icon-document"/><span>分类: {{ article.ArticleCategory.ArticleCategoryName }}</span>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {useRouter} from "vue-router";

const router = useRouter()

const props = defineProps<{
  article: {},
}>()


function archive(articleId: number) {
  router.push('/archive/article/' + articleId)
}
</script>


<style lang="less" scoped>
.panel {
  border: 3px solid rgb(53, 53, 53);
  color: rgb(39, 39, 39);
  position: relative;
  margin-left: 10px;
  margin-right: 10px;
  margin-bottom: 50px;
  height: 250px;
  border-radius: 7px;
  box-shadow: 0 0 5px rgba(100, 100, 100, .5);
  text-overflow: ellipsis;
  overflow: hidden;
  white-space: nowrap;


  .panelJump {

    a:visited,
    a:link {
      color: white;
      text-decoration: none;
    }
  }

  .panelSummary {
    color: #252525;
    position: relative;
    text-overflow: ellipsis;
    text-align: center;
    overflow: hidden;
    white-space: nowrap;
    margin: 10px 30px;
    padding-bottom: 10px;
    /*border-bottom:1px white solid;*/
    transform: translate3d(0, -10px, 0);
    opacity: 0;
    transition: .3s ease all;
    -webkit-transition: .3s ease all;
  }

  &:hover {
    box-shadow: 0 0 10px rgba(110, 104, 104, 0.8);

    .panelTitle {
      transform: translate3d(0, -10px, 0);
      -webkit-transition: .2s ease all;
      font-size: 30px;
      white-space: nowrap;

    }
  }

  &:hover .background {
    transform: scale(1.2, 1.2);
  }

  &:hover .panelSummary {
    transform: translate3d(0, 0, 0);
    -webkit-transition: .3s ease all;
    opacity: 1;
  }

  &:hover .panelInfo {
    transform: translate3d(0, 0, 0);
    opacity: 1;
  }

  .panelInfo {
    position: absolute;
    left: 5px;
    padding-top: 10px;
    color: black;
    padding-right: 100px;
    bottom: 5px;
    margin-left: 2px;
    transform: translate3d(0, -10px, 0);
    opacity: 0;
    transition: .3s ease all;
    -webkit-transition: .3s ease all;

    /*border-top: 1px solid white;*/

    span {
      margin-right: 20px;
    }
  }

  .panelJump {
    position: relative;
  }

  .panelTitle {
    text-align: center;
    position: relative;
    margin-top: -180px;
    font-size: 25px;
    color: black;
    transition: .3s ease all;
    -webkit-transition: .3s ease all;
    cursor: pointer;

    /*span{*/
    /*    cursor: pointer;*/
    /*}*/

    a {
      display: block;
      text-overflow: ellipsis;
      white-space: nowrap;
      overflow: hidden;
    }
  }
}

.background {
  position: relative;
  /*background-image: url(~assets/img/backgroundimg/background1.png);*/
  background-size: cover;
  /*<!--background-position:  0 -200px;-->*/
  width: 100%;
  height: 30vh;
  min-height: 250px;
  transition: .3s all ease;
  //background-image: url(../assets/ec8ce499ly1h1hhm1o6m1j22pg1ww4qp.jpg);
  background-position: top;
  // background-size: cover;
}

@media screen and (max-width: 800px) {
  .panel {
    /*width:100%;*/
    height: 200px;
    font-size: 14px;

    .panelTitle {
      margin-top: -200px;
      opacity: 1;
      transform: translate3d(0, 0, 0);
    }

    .panelInfo {
      opacity: 1;
      transform: translate3d(0, 0, 0);
    }

    .panelSummary {
      opacity: 1;
      transform: translate3d(0, 0, 0);
    }
  }
}
</style>