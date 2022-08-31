<template>
  <div class="rightbar" @click.self="pushCard">
    <div class="showpage activeCard" @wheel.stop="" @touchend.stop="">
      <div @mouseenter.stop="touchCard(0)" class="showtag tag1" @click.stop="pullCard(0)">
        废话
      </div>
      <!-- <div class="showtag tag1" @click="pullCard(0)"> -->
      <!-- 鲤鱼姐 -->
      <!-- </div> -->
      <article-page v-show="showBlog"></article-page>
      <!-- <router-view v-show="showBlog"></router-view> -->
    </div>
    <div class="showpage activeCard" @wheel.stop="" @touchend.stop="">
      <div @mouseenter.stop="touchCard(1)" class="showtag tag2" @click.stop="pullCard(1)">
        蒐集
      </div>
<!--      <n-message-provider>-->
<!--        <blog-page v-if="showMessage"></blog-page>-->
<!--      </n-message-provider>-->
      <love-page v-show="showMessage"></love-page>

      <!-- <router-view name="BlogPage" v-show="showMessage"></router-view> -->
    </div>
    <div class="showpage activeCard" @wheel.stop="" @touchend.stop="">
      <div @mouseenter.stop="touchCard(2)" class="showtag tag3" @click.stop="pullCard(2)">
        留言
      </div>
      <!-- <router-view name="LovePage" v-show="showLove"></router-view> -->
      <n-message-provider>
        <blog-page v-if="showLove"></blog-page>
      </n-message-provider>
<!--      <love-page v-show="showLove"></love-page>-->
    </div>
    <div class="showpage activeCard" @wheel.stop="" @touchend.stop="">
      <div @mouseenter.stop="touchCard(3)" class="showtag tag4" @click.stop="pullCard(3)">
        关于
      </div>
      <!-- <router-view name="VlogPage" v-show="showVlog"></router-view> -->
      <!-- <div class="father"> -->
<!--      <anli-page v-if="showVlog"></anli-page>-->
      <div class="aboutContainer">
      <about v-if="showVlog"></about>
      </div>

      <!-- </div> -->
      <!-- <water-page></water-page> -->
    </div>
  </div>
</template>

<script lang="ts">
export default {
  name: "RightBar"
}
</script>
<script lang="ts" setup>
import {provide, ref} from 'vue';
import {useRouter} from 'vue-router'
import AnliPage from '@/views/RightBar/PicPage.vue';
import ArticlePage from "@/views/RightBar/ArticlePage.vue";
import LovePage from "@/views/RightBar/LovePage.vue";
import BlogPage from "@/views/RightBar/MessageBoard.vue";
import {NMessageProvider} from 'naive-ui'
import About from "../../views/RightBar/About.vue"
const router = useRouter();
const showMessage = ref(false)
const showVlog = ref(false)
const showLove = ref(false)
const showBlog = ref(false)
const hover = ref(false)
const isPushing = ref(false)
provide('hover', hover)
provide('showVlog', showVlog)

function touchCard(num: number) {
  if (document.body.clientWidth / document.body.clientHeight < 12/16) {
    return
  }
  pullCard(num)
}


function pullCard(num: number): void {

  if (isPushing.value) {
    return
  }
  router.push("/home")
  switch (num) {
    case 0:
      showBlog.value = true;
      break;
    case 1:
      showMessage.value = true
      break;
    case 2:
      showLove.value = true
      break;
    case 3:
      showVlog.value = true
      break;
  }

  const tags = document.getElementsByClassName("showtag") as HTMLSelectElement
  const cards = document.getElementsByClassName("showpage") as HTMLCollection
  for (let i = num; i >= 0; i--) {
    let card = cards[i] as HTMLElement
    card.style.transitionDelay = (i - 4) * -0.02 + "s"
    // card.style.right = "0"
    card.className = "showpage"
    let tag = tags[i] as HTMLElement
    tag.classList.remove("activeTag")
  }
  for (let i = num; i < cards.length; i++) {
    let card = cards[i] as HTMLElement
    let tag = tags[i] as HTMLElement

    card.style.transitionDelay = (i - 4) * -0.02 + "s"
    // card.style.right = "-65vw"
    // card.style.right = "0"
    card.className = "showpage activeCard"
    tag.classList.remove("activeTag")

  }

  let element = cards[num] as HTMLElement
  let tag = tags[num] as HTMLElement
  element.className = "showpage"
  tag.classList.add("activeTag")


}

function pushCard() {
  hover.value = false
  isPushing.value = true
  const cards = document.getElementsByClassName("showpage") as HTMLCollection
  const tags = document.getElementsByClassName("showtag") as HTMLSelectElement
  for (let i = 0; i < cards.length; i++) {
    let tag = tags[i] as HTMLElement
    let card = cards[i] as HTMLElement
    card.style.transitionDelay = "0"
    card.style.right = "0"
    card.className = "showpage activeCard"
    tag.classList.remove("activeTag")
  }
  if (!showMessage.value && !showVlog.value && !showBlog.value && !showLove.value) {
    isPushing.value = false
    return
  }

  setTimeout(() => {
    isPushing.value = false

    showMessage.value = false
    showLove.value = false
    showBlog.value = false
    showVlog.value = false
  }, 800);
}


</script>

<style lang="less" scoped>
.father {
  overflow: auto;
  width: 80%;
  height: 100vh;
}

.activeTag {
  background-color: var(--fourth-color) !important;
  color: rgb(255, 255, 255) !important;
}

.rightbar {
  float: right;
  position: relative;
  overflow: hidden;
  width: 100%;
  min-height: 100vh;

  // min-height: 100vh;
  // width: 83.5vw;
  // overflow: hidden;
  .activeCard {
    right: -74vw !important;

    @media (max-aspect-ratio: 1/1) {
      right: -89vw !important;
    }
  }

  .showpage {
    user-select: none;
    // background-attachment: fixed;
    // overflow: ;
    position: absolute;
    right: 0;
    top: 10vh;
    box-shadow: -2px 2px 5px 1px rgba(0, 0, 0, 0.1);
    width: 75vw;
    height: 80vh;
    background-color: #fff;
    // backdrop-filter: blur(3px);
    transition: all .5s ease-in-out; // transition: all 1.5s cubic-bezier(0.19, 1, 0.22, 1);
    border: .5vh solid rgb(49, 49, 49);
    border-right: none;
    border-radius: 5px;
    // right: -65vw;
    background-size: cover;

    &::-webkit-scrollbar {
      width: 0px;
      // transition: all 1s;
    }

    &::-webkit-scrollbar-thumb {
      width: 5px;
      background-color: var(--button-color);
      // transition: all 1s;
      border-radius: 10px;
    }

    @media (max-aspect-ratio: 1/1) AND (min-aspect-ratio: 12/16)  {
      width: 90vw;
      min-height: 80vh;
      //max-height: ;
      // right: -85vw;
    }

    @media (max-aspect-ratio: 12/16) {
      width: 90vw;
      min-height: 85vh;
    }


    &:nth-child(4) {
      z-index: 4;
      // width: 60vw;
      // transform: translateX(20vw);
      // overflow: auto;
      // background-image: url(../assets/91375949_p0.jpg);
      // padding-left: 10vw;
      // padding-top: 5vh;
      display: flex;
      justify-content: space-around;
      flex-direction: column;
      //overflow: auto;

    .aboutContainer{
      height: 100%;
      overflow: auto;
      position: absolute;
      background-size: cover;
      //margin: 10vh auto;
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
        // transition: all 1s;
      }

      &::-webkit-scrollbar-thumb {
        width: 5px;
        background-color: var(--button-color);
        // transition: all 1s;
        border-radius: 10px;
      }
    }
      // background-image: url("./assets/wallhaven-p262de.jpg");
      //border: .5vh solid rgb(49, 49, 49);
      //width: 90%;
      //height: 95%;

    }

    &:nth-child(3) {
      z-index: 3;
      // transform: translateX(15vw);
      // transform: translateX(-5vw);
      // width: 60vw;
      // background-image: url(../assets/89025157_p0.jpg);
      //overflow: auto;
    }

    &:nth-child(2) {
      z-index: 2;
      // transform: translateX(10vw);
      // transform: translateX(-10vw);
      // width: 60vw;
      // background-image: url(../assets/wallhaven-k71327.jpg);
      background-color: white;
    }

    &:nth-child(1) {
      z-index: 1;
      // transform: translateX(-15vw);
      // width: 60vw;
      // background-image: url(../assets/ec8ce499ly1h1hhm1o6m1j22pg1ww4qp.jpg);
    }

    .showtag {

      display: flex;
      z-index: 10;
      position: absolute;
      width: 35px;
      height: 110px;
      background-color: var(--sec-color);
      left: -35px;
      top: 5vh;
      cursor: pointer;
      //box-shadow: -3px 0px 0px 2px rgb(49, 49, 49);
      border-radius: .3vh;
      writing-mode: vertical-lr;
      color: rgb(255, 255, 255);
      font-size: 1.3em;
      padding-top: 1vh;
      border-top: 3px solid rgb(0, 0, 0);
      border-bottom: 3px solid rgb(0, 0, 0);
      box-shadow: -.3vh 0px 0px 2px rgb(0, 0, 0);
      left: -4.1vh;
      width: 4vh;
      height: 11vh;
      font-size: 2vh;
      border-top: .3vh solid rgb(0, 0, 0);
      border-bottom: .3vh solid rgb(0, 0, 0);
      border-right: 0;


      @media (min-width: 1600px) {
        box-shadow: -.3vh 0px 0px 2px rgb(49, 49, 49);
        left: -4.1vh;
        width: 4vh;
        height: 10vh;
        font-size: 2vh;
        border-top: .3vh solid rgb(49, 49, 49);
        border-bottom: .3vh solid rgb(49, 49, 49);
      }

      &:hover {
        background-color: var(--third-color);
        color: rgb(255, 255, 255);
      }
    }

    .tag1 {
      //background-color: #9dd498;

      &:nth-child(2) {
        // box-sizing: content-box;
        border: 4px solid rgb(47, 47, 47);
        left: -10px;
        top: 20px;
        width: 0.1px;
        z-index: -4;
        height: 120px;
        transform-origin: center;
        transform: translateZ(-5px);
        padding: 0;
        // position: absolute;
        // position: fixed;
      }
    }

    .tag2 {
      //background-color: rgb(252, 211, 104);
      // background: linear-gradient(to right, rgb(255, 83, 83), #ffffff);
      top: 23vh;
      // color: white;
    }

    .tag3 {
      //background-color: #f18888;
      top: 41vh
    }

    .tag4 {
      //background-color: #78d2d8;
      top: 59vh;
    }
  }

}

@keyframes move {
  to {
    right: 0;
  }
}
</style>
