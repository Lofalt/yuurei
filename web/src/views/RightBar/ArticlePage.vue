<template>
  <div class="container">
    <div class="catList">
      <div class="cat" @click="getArtsByCat(`all`)" :class="{active:isActive('all')}">全部分类</div>
      <div class="cat" v-for="item in catList" @click="getArtsByCat(item.ID)" :class="{active:isActive(item.ID)}">
        {{ item.ArticleCategoryName }}
      </div>
    </div>
    <panel-com v-for="(item,index) in articleList"
               @click="changePage(1)"
               :style="`--i:${index}`"
               :article="item" :key="item.ID"></panel-com>
    <!--    <n-pagination v-model:page="pageNum" :page-count="pageTotal" show-quick-jumper class="page" v-show="needToShow">-->
    <n-pagination size="medium" v-model:page="pageNum" :page-count="pageTotal" class="page" v-show="needToShow">
    </n-pagination>
  </div>
</template>
<script lang="ts" setup>
import {NPagination} from 'naive-ui'
import {usePageData} from '@/store/pageData';
import PanelCom from '../../components/PanelCom.vue'
import {useRouter} from 'vue-router';
import {computed, ref, watch} from "vue";
import axios from '@/request'

const pageData = usePageData()
const router = useRouter()
const pageNum = ref(1)
const pageTotal = ref(0)
const pageSize = ref(5)
const currentCat = ref("all") as any

const needToShow = computed(() => {
  return pageTotal.value > 1
})

function isActive(num: any) {
  return num === currentCat.value
}

const isShow = computed(() => {
  return pageTotal.value > 8
})

function changePage(num: number) {
  pageData.pagedata.count = 1
  // router.push('/About')
}

const articleList = ref([]) as any
const catList = ref([]) as any

getArts(pageNum.value, pageSize.value)

function getArtsByCat(id: number) {
  pageNum.value = 1
  currentCat.value = id

}

function getArts(pagenum: number, pagesize: number) {
  if (currentCat.value === "all") {
    axios.get(`yuurei/article/all?pageNum=${pagenum}&pageSize=${pagesize}`, {}).then((result: any) => {
      articleList.value = result.data.data
      pageTotal.value = Math.ceil(result.data.total / pageSize.value)
      document.getElementsByClassName("container")[0].scrollTo(0, 0)
    })

  } else {
    axios.get(`yuurei/article/categoryId/${currentCat.value}?pageNum=${pagenum}&pageSize=${pagesize}`, {}).then((result: any) => {
      articleList.value = result.data.data
      pageTotal.value = Math.ceil(result.data.total / pageSize.value)
      document.getElementsByClassName("container")[0].scrollTo(0, 0)

    })
  }
}

axios.get("/yuurei/articleCategory/all", {}).then((result: any) => {
  result.data.data.forEach((data: any) => {
    catList.value.push(data)
  })
})

watch(pageNum, (newValue, oldValue) => {
  getArts(newValue, pageSize.value)
})
watch(currentCat, (newValue, oldValue) => {
  getArts(pageNum.value, pageSize.value)
})
</script>
<style lang="less" scoped>
.page {
  margin: 10px auto;
  justify-content: center;


}

.n-pagination-item--active {
  color: black !important;
  border: 1px black solid !important;
}

.n-pagination {
  --n-item-text-color-active: white !important;
  --n-item-color-active: var(--button-color) !important;

  --n-item-border-active: 0px solid black !important;
  --n-item-text-color-hover: white !important;
  --n-item-color-hover: var(--button-color) !important;
  --n-item-color-active-hover: var(--button-color) !important;


}

.active {
  //background-color: #FF5353FF !important;
  background-color: var(--button-color) !important;
  color: white;
}

.container {
  width: 98%;
  float: left;
  height: 100%;
  overflow: auto;
  position: relative;
  bottom: 0;
  margin: auto auto;

  .catList {
    display: flex;
    justify-content: space-evenly;
    flex-wrap: wrap;

    &::after,
    &::before {
      content: ""; //设置内容为空

      height: 0; //高度为0

      line-height: 0; //行高为0

      display: block; //将文本转为块级元素

      visibility: hidden; //将元素隐藏

      clear: both //清除浮动
    }
  }

  .cat {
    //display: flex;
    //justify-content: center;
    //align-items: center;
    cursor: pointer;
    margin-top: 3vh;
    padding: .2vh 3vh;
    //height:30px;
    //width: 80px;
    margin-left: .9vh;
    //margin-right: .6vh;
    //font-size: 1.5vh;
    line-height: 1.5em;
    font-weight: bold;
    font-family: 微软雅黑;
    background-color: white;
    //text-align: center;
    border: .3vh solid rgb(49, 49, 49);
    border-radius: .2vh;
    box-shadow: -.05vh .05vh 0px 0px rgb(49, 49, 49);
    @media (max-aspect-ratio: 1/1) {
      padding: 0px 3vh;
    }

    &:hover {
      background-color: var(--button-color);
      color: white;
    }
  }

  &::-webkit-scrollbar {
    width: 5px;
    // background-color: aquamarine;
    // color: black;
  }


  &::-webkit-scrollbar-thumb {
    width: 5px;
    background-color: rgba(73, 73, 73, .3);
    color: black;
    border-radius: 20px;
  }

  scrollbar-width: none;
}

::-webkit-scrollbar {
  width: 0px;
}

.panel {
  //background-color: rgba(109, 155, 255, 0.8);
  width: 80%;
  height: 25vh;
  margin: 3vh auto;
  border-radius: .5vh;
  text-align: center;
  cursor: pointer;

  @media (max-aspect-ratio: 9/16) {
    height: 35vw;
  }
}
</style>