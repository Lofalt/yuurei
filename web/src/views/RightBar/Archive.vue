<template>
  <div class="articlePage">
    <img class="headerPic" :src="article.HeaderPicture"/>
    <div v-if="article" class="article">
      <div class="articleHeader">
        <div class="archive-head">
          <!--        {{article.ArticleSummary}}-->
          {{
            new Date(article.CreatedAt).Format("yyyy/M/dd") + "   " + ` (${"日月火水木金土".charAt(new Date(article.CreatedAt).getDay())})`
          }}
        </div>
        <div class="articleTitle">
          {{ article.ArticleTitle }}
        </div>
        <div class="articleInfo">
          <span>{{ article.ArticleSummary }}</span>
          <!--          <span>分类:{{ article.ArticleCategory.ArticleCategoryName }}</span>-->
          <!--          <span>阅读数:{{ article.ArticleReadTimes }}</span>-->
          <!--          <span v-show="article.CreatedAt!=article.UpdatedAt">修改于 {{new Date(article.UpdatedAt).Format("yyyy/M/dd  hh:mm:ss") }}</span>-->
          <!--                <span>被分享:{{article.articleShareTimes}}</span>-->
        </div>
        <!--        <div class="articleSummary" :style="{backgroundImage:`url(${article.HeaderPicture})`}">-->
        <!--          <div class="summary">-->
        <!--            {{ article.ArticleSummary }}-->
        <!--          </div>-->
        <!--        </div>-->
        <!--        <blockquote>{{article.ArticleSummary}}</blockquote>-->
      </div>
      <div class="articleContent" v-html="article.ArticleContent" v-highlight/>
      <div class="footer">
        文章标题:&nbsp;&nbsp;{{ article.ArticleTitle }}<br>
        发布于&nbsp;&nbsp; {{ new Date(article.CreatedAt).Format("yyyy/M/dd  hh:mm:ss") }}<br>
        分类:&nbsp;&nbsp;{{ article.ArticleCategory.ArticleCategoryName }}<br>
        修改于&nbsp;&nbsp; {{ new Date(article.UpdatedAt).Format("yyyy/M/dd  hh:mm:ss") }}
      </div>
    </div>
    <div class="preAndNext">
      <div class="preAndNextBox" v-if="hasPre">
        <n-icon color="#000000" size="1.5vh">
          <arrow-left/>
        </n-icon>
        &nbsp;<a href="javascript:" @click="jump(pre.ID)">{{ pre.ArticleTitle }}</a>
      </div>
      <div v-else disabled>

        <!--        该分类下没有上一篇啦!-->
      </div>
      <div class="preAndNextBox" v-if="hasNext">
        <a href="javascript:" @click="jump(next.ID)">{{ next.ArticleTitle }}</a> &nbsp;
        <n-icon color="#000000" size="1.5vh">
          <arrow-right/>
        </n-icon>
      </div>
      <div v-else disabled>
        <!--        该分类下没有下一篇啦!-->

      </div>
    </div>
    <comments :article="article"/>
    <!--    <comments :article="article" v-cloak/>-->
  </div>
</template>


<script lang="ts" setup>
// import '@wangeditor/editor/dist/css/style.css'

import {computed, onMounted, reactive, ref, watch} from "vue";
// import axios from "axios"
import axios from "@/request"
import qs from "qs"
import {useRouter} from "vue-router"
// import Comments from "../components/article/Comments.vue"
import {NIcon} from 'naive-ui'
import {ArrowLeft, ArrowRight} from "@vicons/fa";
import Comments from "../../components/article/Comments.vue"

const router = useRouter()
const article = ref({
  ID: '',
  ArticleTitle: '',
  ArticleContent: '',
  ArticleSummary: '',
  CreatedAt: '',
  ArticleCategory: {
    ArticleCategoryName: ""
  }
})

const props = defineProps<{
  id: any
}>()


const pre = ref(null)
const next = ref(null)

const hasPre = computed(() => {
  if (pre.value === null) {
    return false
  }
  if (pre.value.ID === 0) {
    return false
  }
  return true

})


const hasNext = computed(() => {
  if (next.value === null) {
    return false
  }
  if (next.value.ID === 0) {
    return false
  }
  return true

})


onMounted(() => {
  axios.get(
      "/yuurei/article/" + props.id,
      {}
  ).then((result: any) => {
    article.value = result.data.data
    pre.value = result.data.pre
    next.value = result.data.next
  })
})
watch(router.currentRoute, (newValue, oldValue) => {
  axios.get(
      "/yuurei/article/" + newValue.params.id, {}
  ).then((result: any) => {
    article.value = result.data.data
    pre.value = result.data.pre
    next.value = result.data.next
    document.getElementsByClassName("articlePage")[0].scrollTo(0, 0)

  })
})
watch(props, (newValue, oldValue) => {
  axios.get(
      "/yuurei/article/" + newValue.id, {}
  ).then((result: any) => {
    article.value = result.data.data
    pre.value = result.data.pre
    next.value = result.data.next
    document.getElementsByClassName("articlePage")[0].scrollTo(0, 0)

  })
})

function jump(id: number) {
  router.push("/archive/article/" + id)
}
</script>


<style lang="less">

  //@media (max-aspect-ratio: 9/16) {
  button{
    cursor: pointer;
  }
  textarea{
    font-family: 微软雅黑;
  }
* {
  //font-size: 2vh;
  ////line-height: 3vh;
  //
  //  font-size: 3.5vw;
  //}
}


.hljs {
  background: #f8feff !important;
  font-family: Consolas;
}

iframe {
  width: 100%;
  min-height: 80vh;


  @media (max-aspect-ratio: 9/16) {
    //height: 50vw;
    min-height: 0;
    //height: 100px;
  }
}

blockquote {
  background: #f9f9f9;
  border-left: 1vh solid #ccc;
  margin: 1.5em 1vh;
  padding: 0.5em 1vh;
  quotes: "\201C""\201D""\2018""\2019";
}

blockquote:before {
  color: #ccc;
  content: open-quote;
  font-size: 4em;
  line-height: 0.1em;
  margin-right: 0.25em;
  vertical-align: -0.4em;
}

blockquote p {
  display: inline;
}

.preAndNextBox {
  display: flex;
  justify-content: center;
  align-items: center;

}

.articlePage {
  font-size: 1.8vh;
  height: 100%;
  width: 100%;
  //padding-top:20px;
  background-color: #ffffff;
  font-family: 微软雅黑;
  //font-size: 20px;
  overflow: auto;

  &::-webkit-scrollbar {
    width: 5px;
  }

  &::-webkit-scrollbar-thumb {
    width: 5px;
    background-color: var(--button-color);
    color: black;
    border-radius: 20px;
  }

  .headerPic {
    width: 100%;
  }

}

.article {
  .articleContent {
    line-height: 4vh;
    font-size: 1.7vh;
    @media (max-aspect-ratio: 9/16) {
      font-size: 3.5vw;
    }

    pre {
      code {
        background-color: #fdf1f1 !important;
      }

      &::before {
        border-top-left-radius: 2vh;
        border-top-right-radius: 2vh;
        display: block;
        content: "";
        width: 100%;
        height: 4vh;
        //background-color: var(--button-color);
        background-color: #ffb1b1;
      }
    }

    * {
      /*margin-top:15px;*/
      //font-size: 1px;

      padding-top: 1.5vh;
      //color: rgb(49, 49, 49)
      /*color: #005cc5;*/
      /*color: #009486;*/
    }

    p {
      //border-bottom: .1vh solid var(--button-color);
      color: rgb(49, 48, 48);
      //display: inline;
      //font-family: "思源宋体","Noto Serif CJK SC","Noto Serif SC","Source Han Serif SC","Source Han Serif",source-han-serif-sc,SongTi SC,SimSum,"Hiragino Sans GB",system-ui,-apple-system,Segoe UI,Roboto,Helvetica,"Microsoft YaHei","WenQuanYi Micro Hei",sans-serif;;
      //font-family: "思源宋体","微软雅黑 Light";
      letter-spacing: .03vh;
    }

    &:nth-child(1) {
      margin-top: 50px;
    }

    a img {
      max-width: 100%;
    }

    img {
      max-width: 100%;
    }

    margin-top: 20px;
    //border-top: 1px solid rgba(58, 58, 58, 0.1);
  }

  .footer {
    font-size: .9em;
    margin-top: 10vh;
    color: darkgray;
  }

  .articleSummary {
    //font-size: 17px;
    font-size: 1.05em;
    text-align: center;
    background-color: #524444;
    width: 70%;
    color: rgb(70, 48, 48);
    margin: 0 auto;
    min-height: 80px;
    /*padding-top:30px;*/
    background-size: cover;
    box-sizing: border-box;
    margin-top: -3vh;
    display: flex;
    padding: 2vh;
    justify-content: center;
    align-items: center;
    font-family: 幼圆;
    border-radius: 2px;
    //border: 3px solid rgb(49, 49, 49);
    /*padding-bottom: 40px;*/
    /*border-bottom: 1px solid black ;*/
    //box-shadow: -1px 1px 0px .5px rgb(49, 49, 49);

    @media (max-aspect-ratio: 1/1) {
      width: 100%;
    }

  }

  margin: 5vh 5vh 2vh 5vh;
  padding: 3vh;
  padding-top: 0;
  border-bottom: 1px solid rgba(58, 58, 58, 0.1);

  .articleHeader {
    position: relative;

    .articleTitle {
      margin-top: 1vh;
      margin-bottom: 2vh;
      //text-align: center;
      font-size: 1.5em;
      font-weight: bold;
      color: #494a4c;

    }

    .articleInfo {
      //text-align: center;
      padding-bottom: 1vh;
      margin-bottom: 5vh;
      border-bottom: 1px solid rgba(65, 65, 65, 0.1);

      span {
        margin-right: 1vh;
        font-size: .9em;
        //color: rgb(52, 40, 40);
        color: rgba(49, 49, 49, .4);
      }
    }
  }
}

.preAndNext {
  position: relative;
  //float: left;
  bottom: 0;
  display: flex;
  justify-content: space-between;
  margin: 0 50px;
  font-size: 1em;
  margin-bottom: 20px;

  a {
    text-decoration: none;
    color: var(--button-color);
    font-weight: bold;
    text-shadow: 0 0 3px rgba(255, 245, 231, 0.5);

    &:hover {
      color: var(--button-color);
    }
  }
}

@media screen and (max-aspect-ratio: 1/1) {
  .preAndNext {
    margin: 1vh .5vh;
  }

  .article {
    //height: 100%;
    padding: 1vh;
    margin: 3vh .6vh;
    //margin: 5vh 5vh 2vh 5vh;
  }

  .articlePage {
    /*padding:0 10px;
    margin: 0;
    /*max-width: 100%;*/
  }
}

@media screen and (max-aspect-ratio: 10/16) {

  .articleSummary {
    width: 100%;
  }


}

.archive-head {
  position: relative;
  padding-left: 1.4vh;
  //height:2vh;
  font-family: "微软雅黑";
  color: rgba(49, 49, 49, .5);
  margin-bottom: 6vh;

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
</style>