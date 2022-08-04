<template>
  <div id="rightbar" @touchstart="touchStart" @touchend="touchEnd">
    <transition>

      <Date v-show="pageData.pagedata.count===0"></Date>
    </transition>
    <div class="middle">
      <!-- <router-view></router-view> -->
      <right-bar @wheel.self="wheeling"></right-bar>
    </div>
    <div class="info" @wheel.self="wheeling">
      <div class="infoPage" @touchend.stop="">

        <router-view></router-view>
      </div>
    </div>
    <div class=" bottom" @wheel.self="wheeling">
      <div class="bottomPage">
        <!-- <anli-page></anli-page> -->
        <!--          <iframe src="//player.bilibili.com/player.html?aid=4539251&bvid=BV1Ds411q7Xp&cid=7360965&page=1"-->
        <!--                  scrolling="no" border="0" frameborder="no" framespacing="0" allowfullscreen="true"></iframe>-->
        <!--        <iframe class="player" src="//player.bilibili.com/player.html?aid=290694593&bvid=BV1yf4y1Y7V1&cid=341769102&page=1"-->
        <!--                scrolling="no" border="0" frameborder="no" framespacing="0" allowfullscreen="true"></iframe>-->
        <!--        <video class="player" src="../assets/单向箭头finn_x264.mp4" width="80vw" height="80vh" controls="controls" loop>s?-->
        <!--        </video>-->

      </div>
    </div>
    <div class=" bottom" @wheel.self="wheeling" v-show="isAdmin">
      <!--        <div class="bottomPage">-->
      <!--      <article-editor/>-->
      <!--        </div>-->
      <!--        <home/>-->
      <!--      <crop-img/>-->

    </div>
  </div>

</template>

<script lang="ts" setup>
import {computed, onMounted, ref, watch} from 'vue';
import RightBar from '@/components/RightBar.vue'
import {usePageData} from '@/store/pageData';
import Date from "@/components/Date.vue";

const isShow = ref(true)
const pageData = usePageData()
let isWheeling = false
var touchX = 0
var touchY = 0

function check(arg: any) {
  alert(arg)
}

function touchStart(event: any) {
  touchX = event.targetTouches[0].pageX;
  touchY = event.targetTouches[0].pageY;
}

function touchEnd(event: any) {
  let touchYEnd = event.changedTouches[0].pageY
  let touches = touchYEnd - touchY
  if (touches < -90) {
    if (pageData.pagedata.count < 3) {
      pageData.pagedata.count++
    }
  }
  if (touches > 90) {
    if (pageData.pagedata.count > 0) {
      pageData.pagedata.count--
    }
  }
}


function touchChange(event: any) {
}

const isAdmin = ref(true)
const showHome = computed(() => {
  return pageData.pagedata.count == 0
})

function wheeling(event: any) {
  if (!isWheeling) {
    if (event.wheelDelta == 150 || event.wheelDelta == 120 || event.wheelDelta > 100) {
      pageData.pagedata.count--
      if (pageData.pagedata.count < 0) {
        pageData.pagedata.count = 0
      }
    }
    if (event.wheelDelta == -150 || event.wheelDelta == -120 || event.wheelDelta < -100) {
      pageData.countIncrease()
      if (pageData.pagedata.count > 2) {
        pageData.pagedata.count = 2
      }
    }
  }
}

onMounted(() => {
  const rightPage = document.getElementById("rightbar") as HTMLSelectElement
  rightPage.style.transform = `translateY(-${pageData.pagedata.count}00vh)`
})

watch(pageData.pagedata, (newValue, oldValue) => {
  const rightPage = document.getElementById("rightbar") as HTMLSelectElement
  rightPage.style.transform = `translateY(-${newValue.count}00vh)`
})

function changePage(num: number) {
  pageData.pagedata.count = num

}


</script>

<style lang="less" scoped>

.player {
  height: 100% !important;
  width: 100% !important;
}

.fixedButton {
  position: fixed;
  top: 1vh;
  right: 1vw;
  z-index: 5;
  cursor: pointer;
  @media (max-width: 1024px) {
    height: 10vh;
    left: 1vw;
    top: 1vh;
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
  display: flex;
  justify-content: center;
  align-items: center;
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
  width: 85%;
  height: 90%;
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

.bottomPage {
  //background-color: rgba(0, 0, 0, 1);
  border: 0px;
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
    transform: translateX(200px);
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