<template>
  <div class="entry" @click="jump">
    <div class="entry-pic" :style="{backgroundImage:`url(${headPic})`}">
      <!--      <n-rate :size="size" color="var(&#45;&#45;button-color)" class="star" readonly :default-value="4" allow-half/>-->
    </div>
    <div class="entry-right">
      <div class="entry-title">
        {{ title }}
      </div>
      <div class="entry-summary">
        {{ summary }}
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {NRate} from 'naive-ui'
import {computed} from "vue";
import {useRouter} from "vue-router";
import {usePageData} from "@/store/pageData";

const pageData = usePageData()
const router = useRouter()
const size = computed(() => {
  if (document.body.clientWidth < 800) {
    return `20`
  } else {
    return `normal`

  }
})

function jump() {
  pageData.pagedata.count = 1
  router.push("/archive/entry/" + props.id)
}


const props = defineProps<{
  headPic: string,
  title: string,
  summary: string,
  id: number

}>()

</script>

<style lang="less" scoped>

.entry {
  //display: flex;
  //justify-content: space-evenly;
  position: relative;
  //background-color: #2dc26b;
  //border-bottom: .1vh solid rgba(49, 49, 49, .1);
  //margin: 0px 100px 100px 100px;
  margin-bottom: 2vh;
  //margin-top: 4vh;
  margin-right: 1vh;
  width: 60vh;
  height: 37vh;
  float: left;
  cursor: pointer;
  transition: all .2s ease;
  //background-color: black;
  //
  //@media (max-aspect-ratio: 16/12){
  //  width: 80vw;
  //  height:40vw;
  //}
  @media (max-aspect-ratio: 1/1) {
    width: 80vw;
    height: 50vw;
  }


  &:hover {
    transform: translate(-10px, -10px);
    //box-shadow: 3px 3px 10px 4px rgba(49, 49, 49, .2);

  }

  //@media (max-width: 800px) {
  //  width: 300px;
  //  height: 200px;
  //  overflow-x: hidden;
  //}

  &:nth-child(2n+0) {
  }

  &::after,
  &::before {
    content: ""; //设置内容为空

    height: 0; //高度为0

    line-height: 0; //行高为0

    display: block; //将文本转为块级元素

    visibility: hidden; //将元素隐藏

    clear: both //清除浮动
  }

  .entry-pic {
    width: 20.7vh;
    height: 30vh;
    background-color: #fff;
    position: relative;
    top: 2vh;
    //border: .3vh solid rgb(49, 49, 49);
    box-shadow: 1vh 1vh .5vh .0vh rgba(49, 49, 49, .05);


    left: 2vh;
    float: left;
    background-image: url("../../assets/p2874590316.webp");
    background-size: cover;

    //@media (max-width: 800px) {
    //  width: 115px;
    //  height: 165px;
    //}
    @media (max-aspect-ratio: 1/1) {
      //left:3vw;
      width: 27.06vw;
      height: 39.23vw;
    }

    .star {
      position: absolute;
      width: 100%;
      text-align: center;
      bottom: -3vh;
      left: 0;
      display: flex;
      justify-content: center;
    }
  }

  .entry-right {
    position: relative;
    //right: 2vh;
    //margin-left:;
    top: 2vh;
    float: right;
    width: 33vh;
    background-color: #fff;
    height: 30vh;
    //text-align: center;
    @media (max-aspect-ratio: 1/1) {
      width: 40vw;
      height: 39.23vw;
    }


    //@media (max-width: 800px) {
    //  width: 140px;
    //  height: 165px;
    //}

    .entry-title {
      margin-top: 10px;
      font-size: 1.2em;
      font-family: '黑体', serif;
      font-weight: bold;
      position: relative;
      text-align: left;
      //display: inline-block;
      //padding-left: 3vh;
      &::after {
        content: '';
        //visibility: hidden;
        width: 100%;
        height: 100%;
        border-bottom: 1px solid rgba(49, 49, 49, .1);
        //position: relative;
        position: absolute;
        //left: 6vh;
        left: 0;
        //background-color: #000;
        top: 3px;

        @media (max-width: 800px) {
          //left: 28px;
        }


      }

      @media (max-width: 800px) {
        font-size: 1.1em;
        margin-top: 0;
      }

    }

    .entry-summary {
      color: #646363;
      text-overflow: ellipsis;
      overflow: hidden;
      //height: 80%;
      width: 100%;
      padding-top: 10px;
      //padding-left: 10px;
      padding-right: 10px;
      display: -webkit-box;
      -webkit-line-clamp: 10;
      -webkit-box-orient: vertical;
      //white-space: nowrap;
      text-align: left;

      @media (max-aspect-ratio: 6/10) {
        -webkit-line-clamp: 7;

      }
      @media (max-aspect-ratio: 9/16) {
        -webkit-line-clamp: 5;

      }
    }
  }

}
</style>