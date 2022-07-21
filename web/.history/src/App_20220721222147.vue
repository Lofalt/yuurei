<template>
  <div class="app">
    <button class="fixedButton" @click="changePage(0)">回到顶端</button>
    <!-- <button @click="change">dmwo </button> -->
    <!-- <button @click="change2">dmwo </button> -->
    <left-bar @change-page="changePage"></left-bar>
    <div id="rightbar" @wheel="wheeling">

      <div class="middle">
        <router-view></router-view>
      </div>
      <div class="bottom">
        <div class="bottomPage">
          <!-- <anli-page></anli-page> -->
        </div>
      </div>
      <div class="info">
        <div class="infoPage">
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { provide, ref, watch } from 'vue';
import LeftBar from './components/LeftBar.vue';
import RightBar from './components/RightBar.vue'
import AnliPage from './views/AnliPage.vue';
import { usePageData } from './store/pageData';

const isShow = ref(true)
const pageCount = ref(0)
const pageData = usePageData()
watch(pageData, (newValue, oldValue) => {
  console.log(newValue.pagedata.data)
  pageCount.value = 2
})

function wheeling(event: any) {
  console.log(event.wheelDelta)
  if (event.wheelDelta == 150 || event.wheelDelta == 120) {
    console.log(event.wheelDelta)
    pageCount.value--
    if (pageCount.value < 0) {
      pageCount.value = 0
    }
  }
  if (event.wheelDelta == -150 || event.wheelDelta == -120) {
    pageCount.value++
    if (pageCount.value > 2) {
      pageCount.value = 2
    }
  }
}


watch(pageCount, (newValue, oldValue) => {
  const rightPage = document.getElementById("rightbar") as HTMLSelectElement
  rightPage.style.transform = `translateY(-${newValue}00vh)`
  console.log(newValue)
})
function changePage(num: number) {
  // const rightPage = document.getElementById("rightbar") as HTMLSelectElement
  pageCount.value = num
  // rightPage.style.transform = `translateY(-${num}00vh)`

}

function change() {
  const card = document.getElementsByClassName("middle") as HTMLSelectElement
  const actualCard = card[0]
  actualCard.style.transition = "all 1s cubic-bezier(0.075, 0.82, 0.165, 1)"
  actualCard.style.transform = "translateY(-100vh)"
  // actualCard.style.transition = ""
  const card2 = document.getElementsByClassName("bottom") as HTMLSelectElement
  const actualCard2 = card2[0]
  actualCard2.style.transition = "all 1s cubic-bezier(0.075, 0.82, 0.165, 1)"
  actualCard2.style.transform = "translateY(-100vh)"
}
function change2() {
  const card = document.getElementsByClassName("middle") as HTMLSelectElement
  const actualCard = card[0]
  actualCard.style.transition = "all 1s cubic-bezier(0.075, 0.82, 0.165, 1)"
  actualCard.style.transform = "translateY(0)"
  // actualCard.style.transition = ""
  const card2 = document.getElementsByClassName("bottom") as HTMLSelectElement
  const actualCard2 = card2[0]
  actualCard2.style.transition = "all 1s cubic-bezier(0.075, 0.82, 0.165, 1)"
  actualCard2.style.transform = "translateY(0)"
}

provide('pageCount', pageCount)
</script>

<style lang="less" scoped>
.fixedButton {
  position: fixed;
  top: 10vh;
  left: 18vw;
}

#rightbar {
  position: relative;
  width: 83.5%;
  height: 100vh;
  float: right;
  // overflow: hidden;
  transition: all 1.2s cubic-bezier(0.075, 0.82, 0.165, 1);
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

.bottomPage,
.infoPage {
  // background-image: url("./assets/wallhaven-p262de.jpg");
  width: 80%;
  min-height: 80%;
  position: relative;
  background-size: cover;
  margin: 10vh auto;
  border-radius: 20px;
  background-color: rgb(255, 170, 170);
  // padding: 100px;
}

.infoPage {
  // background-image: url("./assets/wallhaven-l3w9kr.jpg");
}

button {
  position: absolute;
  left: 10px;
  z-index: 10;

  &:nth-child(2) {
    left: 20px;
  }
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
</style>