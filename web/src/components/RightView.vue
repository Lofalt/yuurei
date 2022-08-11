<template>
  <div id="rightbar" @touchstart="touchStart" @touchend="touchEnd">
    <transition>

      <Date v-show="pageData.pagedata.count===0"></Date>
    </transition>
    <div class="middle">
      <right-bar @wheel.self="wheeling"></right-bar>
    </div>
    <div class="info" @wheel.self="wheeling">
                <span v-show="showSpan"
                      style="display:block;text-align:center;width:80%;margin:0 auto;z-index:1;position: absolute;font-size: 1.5em;color:rgba(49,49,49,.4)">这里空空如也,就像我的脑袋瓜一样.</span>
      <div class="infoPage" @touchend.stop="">
        <router-view></router-view>
        <!--        <entry-view-info/>-->
      </div>
    </div>
    <div class=" bottom" @wheel.self="wheeling">
      <div class="bottomPage" @touchend.stop="">
        <about v-if="pageData.pagedata.count==2"/>

      </div>
    </div>
    <div class=" bottom" @wheel.self="wheeling" v-show="isAdmin">
    </div>
  </div>

</template>

<script lang="ts" setup>
import {computed, onMounted, ref, watch} from 'vue';
import RightBar from '@/components/RightBar.vue'
import {usePageData} from '@/store/pageData';
import Date from "@/components/Date.vue";
import About from "@/views/RightBar/About.vue";
import {useRouter} from "vue-router";
import EntryInfo from "@/views/RightBar/EntryView.vue"
import EntryViewInfo from "../views/RightBar/EntryViewInfo.vue"

const isShow = ref(true)
const pageData = usePageData()
let isWheeling = false
var touchX = 0
var touchY = 0
const router = useRouter()

function check(arg: any) {
  alert(arg)
}

const showSpan = computed(() => {
  return !router.currentRoute.value.path.startsWith("/archive")
})

function touchStart(event: any) {
  touchX = event.targetTouches[0].pageX;
  touchY = event.targetTouches[0].pageY;
}

function touchEnd(event: any) {
  let touchYEnd = event.changedTouches[0].pageY
  let touches = touchYEnd - touchY
  if (touches < -90) {
    if (pageData.pagedata.count < 2) {
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

* {
  font-size: 1.8vh;
}

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

  @media (max-aspect-ratio: 1/1) {
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
  //border: .5vh solid rgb(49, 49, 49);
  width: 90%;
  height: 95%;
  position: relative;
  background-size: cover;
  margin: 10vh auto;
  border-radius: 5px;
  background-color: rgb(255, 255, 255);
  // padding: 100px;
  display: flex;
  justify-content: flex-start;
  align-items: center;
  //font-size: 2vh;
  flex-direction: column;

  &::-webkit-scrollbar {
    width: 5px;
  }

  &::-webkit-scrollbar-thumb {
    width: 5px;
    background-color: rgba(73, 73, 73, .2);
    color: black;
    border-radius: 20px;
  }

}

.bottomPage {
  //background-color: #fffdf7;
  //border: 3px;
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