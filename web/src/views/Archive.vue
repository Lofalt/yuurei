<template>
  <div class="articlePage">
    <div v-if="article" class="article">
      <div class="articleHeader">
        <div class="articleTitle">
          {{ article.ArticleTitle }}
        </div>
        <div class="articleInfo">
          <span>{{ new Date(article.CreatedAt).toLocaleString() }}</span>
          <!--                <span>阅读数:{{article.ArticleReadTimes}}</span>-->
          <!--                <span>被分享:{{article.articleShareTimes}}</span>-->
        </div>
        <div class="articleSummary">
          <div class="summary">
            {{ article.ArticleSummary }}
          </div>
        </div>
      </div>
      <div class="articleContent" v-html="article.ArticleContent"/>
    </div>
    <div class="preAndNext">
      <div class="preAndNextBox" v-if="hasPre">
        <n-icon color="#000000" size="15">
          <arrow-left/>
        </n-icon>
        &nbsp;<a href="javascript:" @click="jump(pre.ID)">{{ pre.ArticleTitle }}</a>
      </div>
      <div v-else disabled>

        <!--        该分类下没有上一篇啦!-->
      </div>
      <div class="preAndNextBox" v-if="hasNext">
        <a href="javascript:" @click="jump(next.ID)">{{ next.ArticleTitle }}</a> &nbsp;
        <n-icon color="#000000" size="15">
          <arrow-right/>
        </n-icon>
      </div>
      <div v-else disabled>
        <!--        该分类下没有下一篇啦!-->

      </div>
    </div>
    <!--    <comments :article="article" v-cloak/>-->
  </div>
</template>


<script lang="ts" setup>


import {computed, onMounted, reactive, ref, watch} from "vue";
// import axios from "axios"
import axios from "@/request/index"
import qs from "qs"
import {useRouter} from "vue-router"
import Comments from "../components/article/Comments.vue"
import {NIcon} from 'naive-ui'
import {ArrowLeft, ArrowRight} from "@vicons/fa";

const router = useRouter()
const article = ref({
  ID: '',
  ArticleTitle: '',
  ArticleContent: '',
  ArticleSummary: '',
  CreatedAt: ''
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

function jump(id: number) {
  router.push("/archive/article/" + id)
}
</script>

<!--<script>-->
<!--    import {request} from "../requests/requests";-->
<!--    import scrollToTop from "../util/scrollToTop";-->
<!--    import Comments from "../components/article/Comments";-->

<!--    export default {-->
<!--    name: "Archive",-->
<!--      components:{-->
<!--      Comments-->
<!--      },-->
<!--    data(){-->
<!--      return {-->
<!--        article:null,-->
<!--        next:null,-->
<!--        pre:null,-->
<!--        id:null,-->
<!--      }-->
<!--    },-->
<!--    created() {-->
<!--      request({-->
<!--        url:'/archive/'+this.$route.params.articleId-->
<!--      }).then((result)=>{-->
<!--        if(!result.data){-->
<!--          location.href='/404'-->
<!--        }-->
<!--        this.article = result.data-->
<!--        request({-->
<!--          url: '/archive/preAndNext',-->
<!--          params:{-->
<!--            'articleCreateTime':result.data.articleCreateTime,-->
<!--            'categoryId':result.data.articleCategories[0].categoryId-->
<!--          }-->
<!--          }-->
<!--        ).then((result2)=> {-->
<!--          this.next = result2.data.next-->
<!--          this.pre = result2.data.pre-->
<!--        })-->
<!--      })-->
<!--    },-->
<!--      beforeRouteUpdate(to,from,next){-->
<!--        request({-->
<!--          url:'/archive/'+to.params.articleId-->
<!--        }).then((result)=>{-->
<!--          this.article = result.data-->
<!--          request({-->
<!--              url: '/archive/preAndNext',-->
<!--              params:{-->
<!--                'articleCreateTime':result.data.articleCreateTime,-->
<!--                'categoryId':result.data.articleCategories[0].categoryId-->
<!--              }-->
<!--            }-->
<!--          ).then((result2)=> {-->
<!--            this.next = result2.data.next-->
<!--            this.pre = result2.data.pre-->
<!--          })-->
<!--        })-->
<!--        next()-->
<!--      },-->
<!--      methods:{-->
<!--          jump(id){-->
<!--            scrollToTop(100)-->
<!--            this.$router.push({-->
<!--              path:'/archive/'+id,-->
<!--            }).catch(err=>{})-->
<!--          },-->
<!--      },-->
<!--      watch:{-->
<!--      '$route'(to,from){-->
<!--        if(this.$route.params.articleId){-->
<!--          request({-->
<!--            url:'/archive/'+this.$route.params.articleId-->
<!--          }).then((result)=>{-->
<!--            this.article = result.data-->
<!--            request({-->
<!--                url: '/archive/preAndNext',-->
<!--                params:{-->
<!--                  'articleCreateTime':result.data.articleCreateTime,-->
<!--                  'categoryId':result.data.articleCategories[0].categoryId-->
<!--                }-->
<!--              }-->
<!--            ).then((result2)=> {-->
<!--              this.next = result2.data.next-->
<!--              this.pre = result2.data.pre-->
<!--            })-->
<!--          })-->
<!--        }-->
<!--      }-->
<!--      },-->
<!--      beforeRouteLeave(to,from,next){-->
<!--        this.article = null-->
<!--        next();-->
<!--      },-->

<!--  }-->
<!--</script>-->

<style lang="less">

.preAndNextBox {
  display: flex;
  justify-content: center;
  align-items: center;

}

.articlePage {
  /*color: white;*/
  height: 100%;
  width: 100%;
  //padding-top:20px;
  background-color: #ffffff;
  font-family: 微软雅黑;
  font-size: 20px;
  overflow: auto;

  &::-webkit-scrollbar {
    width: 5px;
  }

  &::-webkit-scrollbar-thumb {
    width: 5px;
    background-color: rgba(49, 49, 49, .3);
    color: black;
    border-radius: 20px;
  }

}

.article {
  .articleContent {

    * {
      /*margin-top:15px;*/
      font-size: 16px;
      padding-bottom: 15px;
      color: rgb(49, 49, 49)
      /*color: #005cc5;*/
      /*color: #009486;*/
    }

    p:nth-child(1) {
      padding-top: 50px;
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

  .articleSummary {
    font-size: 17px;
    text-align: center;
    background-color: #524444;
    /*width: 60%;*/
    color: white;
    margin: 0 auto;
    min-height: 80px;
    /*padding-top:30px;*/
    box-sizing: border-box;
    margin-top: -30px;
    display: flex;
    padding: 20px;
    justify-content: center;
    align-items: center;
    border-radius: 2px;
    border: 3px solid rgb(49, 49, 49);
    /*padding-bottom: 40px;*/
    /*border-bottom: 1px solid black ;*/
    box-shadow: -1px 1px 0px .5px rgb(49, 49, 49);

    .summary {
    }
  }

  margin: 50px 50px 20px 50px;
  /*margin-bottom: 20px;*/
  /*text-overflow: ellipsis;*/
  /*overflow: hidden;*/
  /*max-height: 190px;*/
  /*background-color: #ffffff;*/
  padding: 30px;
  /*border:5px #ffe4b5 dashed;*/
  /*box-shadow: 0 0 5px rgba(100,100,100,.5);*/
  border-bottom: 1px solid rgba(58, 58, 58, 0.1);
  /*border-radius: 10px;*/

  .articleHeader {
    .articleTitle {
      margin-top: 10px;
      margin-bottom: 20px;
      text-align: center;
      font-size: 30px;
      font-weight: 500;
      color: #494a4c;

    }

    .articleInfo {
      text-align: center;
      padding-bottom: 10px;
      margin-bottom: 50px;
      border-bottom: 1px solid rgba(65, 65, 65, 0.1);

      span {
        margin-left: 10px;
        font-size: 15px;
        color: rgb(52, 40, 40);
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
  font-size: 16px;
  margin-bottom: 20px;

  a {
    text-decoration: none;
    color: #ff4d4d;
    font-weight: bold;
    text-shadow: 0 0 3px rgba(255, 245, 231, 0.5);

    &:hover {
      color: #ff707f;
    }
  }
}

@media screen and (max-width: 800px) {
  .preAndNext {
    margin: 10px 5px;
  }

  .article {
    //height: 100%;
    padding: 10px;
    margin: 0 5px;
  }

  .articlePage {
    /*padding:0 10px;*/
    margin: 0;
    /*max-width: 100%;*/
  }
}

@media screen and (min-width: 800px) {

  .articleSummary {
    width: 80%;
  }


}
</style>