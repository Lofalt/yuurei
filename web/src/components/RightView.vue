<template>
  <div id="rightbar" @touchstart="touchStart" @touchend="touchEnd">

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
        <iframe src="//player.bilibili.com/player.html?aid=290694593&bvid=BV1yf4y1Y7V1&cid=341769102&page=1"
                scrolling="no" border="0" frameborder="no" framespacing="0" allowfullscreen="true"></iframe>

      </div>
    </div>
    <div class=" bottom" @wheel.self="wheeling" v-show="isAdmin">
      <!--        <div class="bottomPage">-->
      <article-editor/>
      <!--        </div>-->
      <!--        <home/>-->
    </div>
  </div>

</template>

<script lang="ts" setup>
import {computed, getCurrentInstance, provide, ref, watch} from 'vue';
import RightBar from '@/components/RightBar.vue'
import {usePageData} from '@/store/pageData';
import ArticleEditor from "@/components/edit/ArticleEditor.vue"

const isShow = ref(true)
const pageData = usePageData()
// const pageCount = pageData.pagedata.count
let isWheeling = false
var touchX = 0
var touchY = 0


function touchStart(event: any) {
  touchX = event.targetTouches[0].pageX;
  touchY = event.targetTouches[0].pageY;
}

function touchEnd(event: any) {
  // alert("hello")
  // alert(touchY)
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


// watch(pageData, (newValue, oldValue) => {
//   pageData.pagedata.count = 1
// })

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
      if (pageData.pagedata.count > 3) {
        pageData.pagedata.count = 3
      }
    }
  }
}


watch(pageData.pagedata, (newValue, oldValue) => {
  const rightPage = document.getElementById("rightbar") as HTMLSelectElement
  rightPage.style.transform = `translateY(-${newValue.count}00vh)`
})

function changePage(num: number) {
  // const rightPage = document.getElementById("rightbar") as HTMLSelectElement
  pageData.pagedata.count = num
  // rightPage.style.transform = `translateY(-${num}00vh)`

}

// function change() {
//   const card = document.getElementsByClassName("middle") as HTMLSelectElement
//   const actualCard = card[0]
//   actualCard.style.transition = "all 1s cubic-bezier(0.075, 0.82, 0.165, 1)"
//   actualCard.style.transform = "translateY(-100vh)"
//   // actualCard.style.transition = ""
//   const card2 = document.getElementsByClassName("bottom") as HTMLSelectElement
//   const actualCard2 = card2[0]
//   actualCard2.style.transition = "all 1s cubic-bezier(0.075, 0.82, 0.165, 1)"
//   actualCard2.style.transform = "translateY(-100vh)"
// }
// function change2() {
//   const card = document.getElementsByClassName("middle") as HTMLSelectElement
//   const actualCard = card[0]
//   actualCard.style.transition = "all 1s cubic-bezier(0.075, 0.82, 0.165, 1)"
//   actualCard.style.transform = "translateY(0)"
//   // actualCard.style.transition = ""
//   const card2 = document.getElementsByClassName("bottom") as HTMLSelectElement
//   const actualCard2 = card2[0]
//   actualCard2.style.transition = "all 1s cubic-bezier(0.075, 0.82, 0.165, 1)"
//   actualCard2.style.transform = "translateY(0)"
// }

// provide('pageCount', pageCount)
</script>

<style lang="less" scoped>
.fixedButton {
  position: fixed;
  top: 10vh;
  left: 18vw;
  // outline: none;
  // padding: 10px;
  // border: none;
  // background-color: rgba(255, 255, 255, 0.384);
  // border-radius: 10px;
  // cursor: pointer;
  // color: rgba(37, 37, 37, 0.733);
  cursor: pointer;
  // width: 5vw;
  padding: 5px 20px;
  background-color: #ffffff;
  // height: 4vh;
  border-radius: 2px;
  outline-style: none;
  border: 3px solid rgb(51, 51, 51);
  box-shadow: -1px 1px 0 0px rgb(51, 51, 51);

  @media (max-width: 1024px) {
    height: 10vh;
    left: 2vw;
    // padding: 10px;
    padding: 0;
    width: 8vw;
    top: 0;
    writing-mode: vertical-lr;
    background-color: rgb(49, 49, 49);
    color: white;
  }

  &:hover {
    background-color: rgb(255, 83, 83) !important;
    color: white !important;

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