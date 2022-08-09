<template>
  <div class="container">
    <div class="catList">
      <div class="cat" @click="getEntsByCat(`all`)" :class="{active:isActive('all')}">ALL</div>
      <div class="cat" v-for="item in catList" @click="getEntsByCat(item.ID)" :class="{active:isActive(item.ID)}">
        {{ item.CategoryName }}
      </div>
    </div>
    <!--    <div class="nav">navbar</div>-->
    <div class="inside">
      <entry v-for="item in entries" :id="item.ID" :headPic="item.HeaderPicture" :title="item.Title"
             :summary="item.Summary" :key="item.ID"/>
    </div>
    <n-pagination size="medium" v-model:page="pageNum" :page-count="pageTotal" class="page" v-show="needToShow">
    </n-pagination>
  </div>
</template>

<script lang="ts" setup>
import Entry from "../../components/Entry.vue"
import axios from "../../request/index"
import {computed, ref, watch} from "vue";
import {usePageData} from "@/store/pageData";
import {useRouter} from "vue-router";
import {NPagination} from "naive-ui"

const entries = ref([])


const pageData = usePageData()
const router = useRouter()
const pageNum = ref(1)
const pageTotal = ref(8)
const pageSize = ref(5)
const currentCat = ref("all") as any
const catList = ref([]) as any

getEnts(pageNum.value, pageSize.value)
// axios.get("/yuurei/entry/all",{}).then((res:any)=>{
//   entries.value = res.data.data
// })
getEnts(pageNum.value, pageSize.value)

axios.get("/yuurei/entryCategory/all", {}).then((result: any) => {
  result.data.data.forEach((data: any) => {
    catList.value.push(data)
  })
})

watch(pageNum, (newValue, oldValue) => {
  getEnts(newValue, pageSize.value)
})
watch(currentCat, (newValue, oldValue) => {
  getEnts(pageNum.value, pageSize.value)
})

function getEnts(pagenum: number, pagesize: number) {
  if (currentCat.value === "all") {
    axios.get(`yuurei/entry/all?pageNum=${pagenum}&pageSize=${pagesize}`, {}).then((result: any) => {
      entries.value = result.data.data
      pageTotal.value = Math.ceil(result.data.total / pageSize.value)
      document.getElementsByClassName("container")[0].scrollTo(0, 0)
    })

  } else {
    axios.get(`yuurei/entry/categoryId/${currentCat.value}?pageNum=${pagenum}&pageSize=${pagesize}`, {}).then((result: any) => {
      entries.value = result.data.data
      pageTotal.value = Math.ceil(result.data.total / pageSize.value)
      document.getElementsByClassName("container")[0].scrollTo(0, 0)

    })
  }
}

function getEntsByCat(id: number) {
  pageNum.value = 1
  currentCat.value = id

}

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


</script>

<style lang="less" scoped>
.container {
  width: 100%;
  padding: 1vh 1vh;
  overflow: auto;
  //min-height: 50vh;
  position: relative;
  bottom: 0;
  height: 100%;
  margin: 0 auto;
  //overflow-x: hidden;

  .catList {
    display: flex;
    justify-content: space-evenly;
    flex-wrap: wrap;
    margin-bottom: 2vh;

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
    margin-top: 3vh;
    padding: .1vh 3vh;
    //height:30px;
    //width: 80px;
    margin-left: .9vh;
    //margin-right: .6vh;
    font-size: 1.5vh;
    font-weight: bold;
    font-family: 微软雅黑;
    background-color: white;
    text-align: center;
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

  .nav {
    height: 40px;
    text-align: center;
    background-color: #ffb0d1;
  }

  .inside {
    text-align: center;
    margin: 0 auto;
    width: 100%;
    //min-height: 100%;
    padding-left: 70px;
    display: flex;
    justify-content: center;
    align-content: flex-start;
    flex-wrap: wrap;
    @media (max-width: 800px) {
      padding-left: 0;
    }
  }

  &::-webkit-scrollbar {
    width: 5px;
  }

  &::-webkit-scrollbar-thumb {
    width: 5px;
    background-color: rgba(49, 49, 49, .2);
  }

  @media (max-width: 800px) {
    width: 100%;
  }

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
}


</style>