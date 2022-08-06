<template>
<!--    <div>没什么艺术细胞的人</div>-->
  <div id="main" @wheel="listenScroll" @touchstart.stop="touchStart" @touchend.stop="touchEnd">
    <!-- <button class="scrollToTop" @click="scrollToTop">dmwo</button> -->
    <img class=" test" v-for="(item, index) in picList" :src="item.Path.slice(1)" :key="item.ID" ref="waterFallItem"
         @click="zoom(index)" :style="{ animationDelay: (index % 2) * 0.1 + 's' }"/>
  </div>
  <!-- <loading-com class="loading"></loading-com> -->
  <loading-com class="loading" v-show="isLoading"></loading-com>
  <!-- <div class="loading" v-show="offset == 15">--</div> -->
  <transition>
    <div @click="zoomout" v-if="hover" class="photoInfo" :style="{ backgroundImage: `url('` + showingPage + `')` }">
      <div class="info" v-show="currentInfo">
        {{ currentInfo }}
      </div>
    </div>
  </transition>
</template>

<script lang="ts" setup>

import axios from "@/request"
import qs from 'qs'
import {reactive, ref, onMounted, watch, nextTick, inject, Ref} from "vue";
import LoadingCom from "../../components/LoadingCom.vue"

const isShow = ref(true)
const picList = ref<any>([])
const offset = ref(10)
const total = ref(0)
const pageSize = ref(6)
const pageNum = ref(1)
const colRaw = ref(2)
const screenWidth = ref(0)
const waterFallItem = ref<HTMLSelectElement>()
const showingPage = ref('')
const hover = inject('hover') as Ref<boolean>
const showVlog = inject('showVlog') as Ref<boolean>
const isLoading = ref(false)
const currentInfo = ref("")
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
    if (!isLoading.value && picList.value.length < total.value) {
      isLoading.value = true
      setTimeout(() => {
        getNext()

      }, 1000);
    }
  }

}


function scroll(scrollDuration: number, box: any) {
  var cosParameter = box.scrollY / 2,
      scrollCount = 0,
      oldTimestamp = performance.now();

  function step(newTimestamp: any) {
    scrollCount += Math.PI / (scrollDuration / (newTimestamp - oldTimestamp));
    if (scrollCount >= Math.PI) box.scrollTo(0, 0);
    if (box.scrollY === 0) return;
    box.scrollTo(0, Math.round(cosParameter + cosParameter * Math.cos(scrollCount)));
    oldTimestamp = newTimestamp;
    // box.requestAnimationFrame(step);
  }

  // box.requestAnimationFrame(step);q
}

function scrollToTop() {
  const box = document.getElementById("main") as HTMLSelectElement
  // box.scrollTo(0, 0)
  scroll(500, box)
}

function listenScroll(event: any) {
  console.log("？")
  const box = document.getElementById("main") as HTMLSelectElement
  if (box.scrollTop + box.offsetHeight + 100 > box.scrollHeight) {
    // offset.value += 5
    // console.log((pageNum.value ) * pageSize.value)
    if (!isLoading.value && picList.value.length < total.value) {
      console.log(".")
      isLoading.value = true
      setTimeout(() => {
        getNext()

      }, 1000);

    } else {
      return
    }
  }

}


function zoomout() {
  hover.value = false
}

function zoom(event: number) {
  let childs = document.getElementsByClassName("test") as HTMLSelectElement
  showingPage.value = childs[event].getAttribute("src") as string
  hover.value = true
  currentInfo.value = picList.value[event].Description

}

function getNext() {

  // if ((pageNum.value + 1) * pageSize.value > total.value) {
  //   return
  // }
  isLoading.value = true
  pageNum.value += 1
  axios.get(`/yuurei/gallery/all?pageNum=${pageNum.value}&pageSize=${pageSize.value}`, {}).then((result: any) => {
    for (let i = 0; i < result.data.data.length; i++) {
      picList.value.push(result.data.data[i])
    }
    waterFall()
  })
}

// qs.stringify({num: num})
function getPic() {
  axios.get(`/yuurei/gallery/all?pageNum=${pageNum.value}&pageSize=${pageSize.value}`, {}).then((result: any) => {
    picList.value = result.data.data
    total.value = result.data.total
    setTimeout(() => {
      waterFall()

    }, 500);
  })
}

function waterFall() {
  setTimeout(() => {
    let childs = document.getElementsByClassName("test") as HTMLSelectElement
    // showingPage.value = childs[0].getAttribute("src") as string
    var boxWidth = childs[0].offsetWidth
    var heightArr = [], boxHeight = 0, minBoxHeight = 0, minIndex = 0
    var col: number = colRaw.value
    for (let i = 0; i < childs.length; i++) {
      boxHeight = childs[i].offsetHeight
      if (i < col) {
        heightArr.push(boxHeight + 50)
        childs[i].style.position = 'absolute'
        childs[i].style.left = i * boxWidth + 'px'
        childs[i].style.top = 3 + '0px'
        childs[i].style.opacity = '1'
        childs[i].style.transform = "translateY(0)"
      } else {
        minBoxHeight = heightArr[minBox(heightArr)]
        minIndex = minBox(heightArr)
        childs[i].style.position = 'absolute'
        childs[i].style.left = minIndex * boxWidth + 'px'
        childs[i].style.top = minBoxHeight + 'px'
        childs[i].style.opacity = '1'
        childs[i].style.transform = "translateY(0)"
        heightArr[minIndex] += boxHeight
      }
    }
    isLoading.value = false
  }, 600)
}

function minBox(box: any) {
  var j = 0
  for (let i = 0; i < box.length; i++) {
    if (box[j] > box[i]) {
      j = i
    }
  }
  return j
}

onMounted(() => {
  getPic()
  nextTick(() => {
    // waterFall()

    window.onscroll = () => {
      let scrollTop = document.documentElement.scrollTop; //获取滚动距离
      let scrollHeigh = document.documentElement.scrollHeight; //获取整个页面的高度
      let clientHeigh = document.documentElement.clientHeight; //获取
      if (scrollTop + clientHeigh - scrollHeigh >= -50) {
        getNext()
      }
    }
    window.onresize = () => {
      return (() => {
        screenWidth.value = document.body.clientWidth
      })()
    }

  })
})
watch(screenWidth, () => {
  if (screenWidth.value >= 1000) {
    colRaw.value = 2
  }
  // if (screenWidth.value < 1000 && screenWidth.value >= 700) {
  //     colRaw.value = 2

  // }
  if (screenWidth.value < 800) {
    colRaw.value = 2
  }
  waterFall()
})
watch(colRaw, () => {
  // let childs = document.getElementsByClassName("test") as HTMLSelectElement
  // for (let i = 0; i < childs.length; i++) {
  //     childs[i].style.width = 100 / colRaw.value + "%"
  // }
})
</script>

<style lang="less" scoped>
.loading {
  // position: relative;
  bottom: 0;
  // height: 5%;
  // float: left;
  text-align: center;
}

.v-enter-from,
.v-leave-to {
  opacity: 0;
}

.v-enter-active {
  transition: all .8s;
}

.v-leave-active {
  transition: all .3s;
}

.photoInfo {
  background-color: rgba(43, 43, 43, 0.738);
  width: 100%;
  // width: auto;
  height: 100%;
  position: absolute;
  // position: fixed;
  top: 0;
  left: 0;
  // z-index: 10;
  background-size: contain;
  background-repeat: no-repeat;
  // animation: zoom 1s ease;
  cursor: pointer;
  background-position: center;
  // z-index: 10;

  .info {
    position: absolute;
    top: 15px;
    left: 2%;
    background-color: rgb(124, 108, 105);
    //background-color: rgb(124, 105, 105);
    color: white;
    padding: 10px 20px;
    border-radius: 5px;
    //border:2px solid rgb(49,49,49);
    //border:3px solid rgba(255,255,255,.1);
    box-shadow: -2px 2px 1px 1px rgba(49, 49, 49, .2);
  }
}

@keyframes zoom {
  from {
    opacity: 0;
  }
}

.scrollToTop {
  position: fixed;
  z-index: 10;
  top: 100px;
}

.test {
  cursor: pointer;
  left: 0px;
  width: 50%;
  // height: 100%;
  box-sizing: border-box;
  padding: 8px;
  // margin: 8px;
  /* margin:15px; */
  // float: left;
  transition: opacity 0.2s ease-in-out;
  // transition: all 1s;
  // animation: enter 1s;
  // animation-timing-function: ease-in-out;
  // animation-fill-mode: backwards;
  // box-shadow: 0 0 1px 1px rgba(43, 43, 43, 0.738);
  opacity: 0;
  // transform: translateY(1000px);

  &:hover {
    transform: translate(-5px, -5px) !important;
    box-shadow: -2px 2px 5px 1px rgba(0, 0, 0, 0.1);
    transition: all .2s ease;

  }

  @media (max-width: 800px) {
    padding: 5px;

  }
}

#main {
  // position: relative;
  // margin: 0 auto;
  height: 100%;
  // width: 80%;
  // margin: 0 auto;
  width: 70%;
  // float: left;
  // height: 100%;
  overflow: auto;
  overflow-x: hidden;
  position: relative;
  // bottom: 0;
  margin: 0 auto;
  padding-top: 5vh;

  &::-webkit-scrollbar {
    width: 0px;
    // transition: all 1s;
  }

  &::-webkit-scrollbar-thumb {
    width: 5px;
    background-color: rgba(110, 110, 110, 0.6);
    // transition: all 1s;
    border-radius: 10px;
  }

  &::after,
  &::before {
    content: "";
    // overflow: hidden;
    display: block;
    height: 0;
    clear: both;
    visibility: hidden;
  }

  @media (max-width: 800px) {
    // padding-top: 2vh;
    width: 90%
  }


}


button {
  position: absolute
}

.box {
  width: 100%;
  background-color: yellow;
  height: 100px;
  transition: all 1s ease-in-out;
}


@keyframes enter {
  from {
    // transform: translateY(1000px);
    opacity: 0;
  }
}
</style>