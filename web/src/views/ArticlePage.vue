<template>
  <div class="container">
    <div class="catList">
      <div class="cat" @click="getArtsByCat(`all`)" :class="{active:isActive('all')}">ALL</div>
      <div class="cat" v-for="item in catList" @click="getArtsByCat(item.ID)" :class="{active:isActive(item.ID)}">
        {{ item.ArticleCategoryName }}
      </div>
    </div>
    <panel-com v-for="item in articleList"
               @click="changePage(1)"
               :article="item" :key="item.ID"></panel-com>
    <!--    <n-pagination v-model:page="pageNum" :page-count="pageTotal" show-quick-jumper class="page" v-show="needToShow">-->
    <n-pagination v-model:page="pageNum" :page-count="pageTotal" class="page" v-show="needToShow">
      <!--      <template #goto v-show="isShow">-->
      <!--        跳转-->
      <!--      </template>-->
    </n-pagination>
  </div>
</template>
<script lang="ts" setup>
import {NPagination} from 'naive-ui'
import {usePageData} from '@/store/pageData';
import PanelCom from '../components/PanelCom.vue'
import {useRouter} from 'vue-router';
import {computed, ref, watch} from "vue";
import axios from '@/request/index'

const pageData = usePageData()
const router = useRouter()
const pageNum = ref(1)
const pageTotal = ref(8)
const pageSize = ref(8)
const currentCat = ref("all")

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

const articleList = ref([])
const catList = ref([])

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

axios.get("/yuurei/articleCategory/all").then((result) => {
  result.data.data.forEach((data) => {
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
  --n-item-color-active: #FF5353FF !important;

  --n-item-border-active: 2px solid black !important;
  --n-item-text-color-hover: white !important;
  --n-item-color-hover: #FF5353FF !important;
  --n-item-color-active-hover: #FF5353FF !important;


}

.active {
  background-color: #FF5353FF !important;
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
    justify-content: space-around;
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
    cursor: pointer;
    margin-top: 20px;
    padding: 3px 20px;
    background-color: white;
    border: 3px solid rgb(49, 49, 49);
    border-radius: 2px;
    box-shadow: -1px 1px 0px 0px rgb(49, 49, 49);
    @media (max-width: 1024px) {
      padding: 0px 40px;
    }

    &:hover {
      background-color: #FF5353FF;
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
  background-color: rgba(109, 155, 255, 0.8);
  width: 80%;
  height: 25vh;
  margin: 30px auto;
  border-radius: 5px;
  text-align: center;
  cursor: pointer;
}
</style>