<template>
  <div class="app">
    <transition>
      <n-icon size="40" color="#ffffff" @click="changePage(0)" class="fixedButton" v-show="pageData.pagedata.count!=0">
        <arrow-bar-to-up/>
      </n-icon>
    </transition>
    <left-bar></left-bar>
    <!--    <transition>-->

    <!--    <Date v-show="pageData.pagedata.count==0"></Date>-->
    <!--    </transition>-->

    <router-view></router-view>
    <!--    <div id="rightbar" @touchstart="touchStart" @touchend="touchEnd">-->

    <!--      <div class="middle">-->
    <!--        &lt;!&ndash; <router-view></router-view> &ndash;&gt;-->
    <!--        <right-bar @wheel.self="wheeling"></right-bar>-->
    <!--      </div>-->
    <!--      <div class="info" @wheel.self="wheeling">-->
    <!--        <div class="infoPage" @touchend.stop="">-->

    <!--          <router-view></router-view>-->
    <!--        </div>-->
    <!--      </div>-->
    <!--      <div class=" bottom" @wheel.self="wheeling">-->
    <!--        <div class="bottomPage">-->
    <!--          &lt;!&ndash; <anli-page></anli-page> &ndash;&gt;-->
    <!--          &lt;!&ndash;          <iframe src="//player.bilibili.com/player.html?aid=4539251&bvid=BV1Ds411q7Xp&cid=7360965&page=1"&ndash;&gt;-->
    <!--          &lt;!&ndash;                  scrolling="no" border="0" frameborder="no" framespacing="0" allowfullscreen="true"></iframe>&ndash;&gt;-->
    <!--          <iframe src="//player.bilibili.com/player.html?aid=290694593&bvid=BV1yf4y1Y7V1&cid=341769102&page=1"-->
    <!--                  scrolling="no" border="0" frameborder="no" framespacing="0" allowfullscreen="true"></iframe>-->

    <!--        </div>-->
    <!--      </div>-->
    <!--      <div class=" bottom" @wheel.self="wheeling" v-show="isAdmin">-->
    <!--        &lt;!&ndash;        <div class="bottomPage">&ndash;&gt;-->
    <!--        <article-editor/>-->
    <!--        &lt;!&ndash;        </div>&ndash;&gt;-->
    <!--        &lt;!&ndash;        <home/>&ndash;&gt;-->
    <!--      </div>-->
    <!--    </div>-->
  </div>
</template>

<script lang="ts" setup>
import {computed, getCurrentInstance, provide, ref, watch} from 'vue';
import LeftBar from './components/LeftBar.vue';
import {usePageData} from './store/pageData';
import Date from "./components/Date.vue";
import {ArrowBarToUp} from "@vicons/tabler"
import {NIcon} from "naive-ui"

const isShow = ref(true)
const pageData = usePageData()
let isWheeling = false
const isAdmin = ref(true)
const showHome = computed(() => {
  return pageData.pagedata.count == 0
})


function changePage(num: number) {
  pageData.pagedata.count = num

}

</script>

<style lang="less" scoped>
.fixedButton {
  position: fixed;
  top: 5vh;
  left: 18vw;
  z-index: 5;
  cursor: pointer;
  @media (max-width: 1024px) {
    height: 10vh;
    left: 1vw;
    // padding: 10px;
    //padding: 0;
    //width: 8vw;
    top: 1vh;
    //writing-mode: vertical-lr;
    //background-color: rgb(49, 49, 49);
    //color: white;
  }

  &:hover {
    //background-color: rgb(255, 83, 83) !important;
    //color: white !important;

  }

}

#rightbar {
  position: relative;
  width: 83.5%;
  height: 100vh;
  float: right;
  // overflow: hidden;
  transition: all .8s ease-in-out;

  @media (max-width: 1024px) {
    width: 100%;
  }
}

.middle {
  float: right;
  position: relative;
  overflow: hidden;
  width: 100%;
  min-height: 100vh;
}

.bottom,
.info {
  position: relative;
  float: right;
  width: 100%;
  height: 100vh;
  transition: all 2s;
}

iframe {
  width: 100%;
  height: 100%;
}

.bottomPage,
.infoPage {
  overflow: auto;
  // background-image: url("./assets/wallhaven-p262de.jpg");
  border: 5px solid rgb(49, 49, 49);
  width: 80%;
  height: 80%;
  position: relative;
  background-size: cover;
  margin: 10vh auto;
  border-radius: 5px;
  background-color: rgb(255, 255, 255);
  // padding: 100px;
  display: flex;
  justify-content: center;
  align-items: center;

  &::-webkit-scrollbar-thumb {
    width: 30px;
    background-color: rgba(73, 73, 73, .5);
    color: black;
    border-radius: 20px;
  }

}

.infoPage {
  // background-image: url("./assets/wallhaven-l3w9kr.jpg");
}

button {
  position: absolute;
  left: 10px;
  z-index: 10;

  // &:nth-child(2) {
  //   left: 20px;
  // }
}

#app {
  min-height: 100vh;
  // overflow: hidden;
}

.app {
  background-size: cover;
  background-repeat: repeat-y;
  // background-image: url("./assets/94029233_p0.jpg");
  position: relative;
  min-height: 100vh;
  overflow: hidden;
}

.middle {
  animation: move 1s ease;
}

/* .top {
  position: absolute;
  right: 1000px;
  animation: move 1s ease;
} */

@keyframes move {
  from {
    transform: translateY(1000px);
  }

  /* to {
    transform: translateX(-1000px);
  } */
}

.v-leave-to,
.v-enter-from {
  transform: translateY(-100vh);
}

.v-leave-active,
.v-enter-active {
  transition: all .8s ease-in-out;
}
</style>