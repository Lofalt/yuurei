<template>
  <div id="rightbar" @touchstart="touchStart" @touchend="touchEnd">
    <transition>

      <Date v-show="pageData.pagedata.count===0"></Date>
    </transition>
    <div class="middle">
      <right-bar @wheel.self="wheeling"></right-bar>
    </div>
    <div class="info" @wheel.self="wheeling">
      <!--                <span v-show="showSpan"-->
      <!--                      style="display:block;text-align:center;width:80%;margin:0 auto;z-index:1;position: absolute;font-size: 1.5em;color:rgba(49,49,49,.4)">这里空空如也,就像我的脑袋瓜一样.</span>-->
      <div class="infoPage" @touchend.stop="">
        <archive :id="AID" v-show="showSpan"/>
        <router-view></router-view>
        <!--        <entry-view-info/>-->
      </div>
    </div>
    <div class=" bottom" @wheel.self="wheeling">
<!--      <div class="bottomPage" @touchend.stop="">-->
        <div class="bottomPage">
<!--        <about v-if="pageData.pagedata.count==2"/>-->
    <div class="bottomTitle">
      <div>记忆</div>
      <div>一些碎片</div>
    </div>
        <pic-page v-if="pageData.pagedata.count==2"/>
      </div>
    </div>
    <div class=" bottom" @wheel.self="wheeling" v-show="isAdmin">
    </div>
  </div>

</template>

<script lang="ts" setup>
import {computed, onMounted, provide, ref, watch} from 'vue';
import RightBar from '@/components/root/RightBar.vue'
import {usePageData} from '@/store/pageData';
import Date from "@/components/util/Date.vue";
import About from "@/views/RightBar/About.vue";
import {useRouter} from "vue-router";
import EntryInfo from "@/views/RightBar/EntryView.vue"
import EntryViewInfo from "../../views/RightBar/EntryViewInfo.vue"
import Archive from "../../views/RightBar/Archive.vue"
import axios from "../../request/index"
import PicPage from "../../views/RightBar/PicPage.vue"

const isShow = ref(true)
const pageData = usePageData()
const AID = ref(0)
let isWheeling = false
var touchX = 0
var touchY = 0
const router = useRouter()
const hover = ref(false)
const isPushing = ref(false)
provide('hover', hover)
// provide('showVlog', showVlog)

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

axios.get("yuurei/article/all?pageNum=1&pageSize=1", {}).then((res) => {
  AID.value = res.data.data[0].ID
})
onMounted(() => {
  const rightPage = document.getElementById("rightbar") as HTMLSelectElement
  rightPage.style.transform = `translateY(-${pageData.pagedata.count}00%)`
})

watch(pageData.pagedata, (newValue, oldValue) => {
  const rightPage = document.getElementById("rightbar") as HTMLSelectElement
  rightPage.style.transform = `translateY(-${newValue.count}00%)`
})

function changePage(num: number) {
  pageData.pagedata.count = num

}
onMounted(()=>{
  // let rightbar = document.getElementById("rightbar") as any
  // rightbar.style.height = window.innerHeight + "px"

  // window.onresize = () => {
    // return (() => {
    //   let rightbar = document.getElementById("rightbar") as any
    //   rightbar.style.height = window.innerHeight + "px"
    // })()
  // }
})


</script>

<style lang="less" scoped>

* {
  font-size: 1.8vh;

  @media (max-aspect-ratio: 9/16) {
    font-size: 3vw;
  }
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
  height: 100%;
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
  height: 100%;
}

.bottom,
.info {
  position: relative;
  float: right;
  width: 100%;
  height: 100%;
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
    background-color: var(--button-color);
    color: black;
    border-radius: 20px;
  }

}

.bottomPage {
  //background-color: #fffdf7;
  //border: 3px;
  padding:4vh 0;
  padding-bottom: 0;
  flex-direction: row;
  position: relative;
  .bottomTitle{
    display: flex;
    flex-direction: column;
    position: relative;
    justify-content: flex-start;
    height: 100%;
    width: 10vh;
    padding-left: 2vh;
    flex-shrink: 0;
    align-content: center;
    @media (max-aspect-ratio: 9/16){
      margin-bottom: 1vh;
      height: 3.5%;
      padding-left: 0;
      margin-left: 6%;
      flex-direction: row;
      width: 88%;
      margin:.5vh auto;
      justify-content: space-between;
      //padding-left: 5%;
      //align-content: flex-start;
      //justify-content: flex-start;
    }
    div:nth-child(1){
    min-height: 1vh;
      position: relative;
      padding-left:1.5vh;
      font-family: -apple-system, BlinkMacSystemFont, Tahoma, Arial, "Hiragino Sans GB", "Microsoft YaHei", sans-serif;
      font-size: 1.5em;
      font-weight: bold;
      color: #343333;

      @media (max-aspect-ratio: 9/16){
        font-size: 1.4em;
        //padding-left: 5%;
        //align-content: flex-start;
        //justify-content: flex-start;
      }
      &::before {
      left: 0;
      border-radius: 1vh;
      content: " ";
      width: .8vh;
      height: 80%;
        top:10%;
      display: inline-block;
      position: absolute;
      background-color: var(--button-color);
    }
    }
    div:nth-child(2) {
      margin-top:1vh;
      color:grey;
      font-size: .9em;
    }
    }
  &::-webkit-scrollbar {
    width: 0px;
  }

  &::-webkit-scrollbar-thumb {
    width: 5px;
    background-color: var(--button-color);
    color: black;
    border-radius: 20px;
  }
  @media (max-aspect-ratio: 9/16){
    flex-direction: column;
    padding-top:1vh;
    //align-content: flex-end;
    align-items: flex-start;
    //justify-content: flex-start;
    //padding-left: 5%;

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