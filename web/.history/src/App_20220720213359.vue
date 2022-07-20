<template>
  <div class="app">
    <button @click="change">dmwo </button>
    <button @click="change2">dmwo </button>
    <left-bar></left-bar>
    <div class="rightbar">

      <div class="middle">
        <router-view></router-view>
      </div>
      <div class="bottom">
        bottom
      </div>
    </div>

  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import LeftBar from './components/LeftBar.vue';
import RightBar from './components/RightBar.vue'

const isShow = ref(true)

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
</script>

<style lang="less" scoped>
.rightbar {
  /* background-color: #fff; */
  width: 83.5%;
  min-height: 100vh;
  float: right;
  overflow: hidden;
}

.top {
  position: absolute;
  // float: right;
  top: 100vh;
  transition: all 2s;
  width: 100%;
  min-height: 100vh;
  background-color: rgb(196, 255, 229);
}

.middle {
  float: right;
  position: relative;
  overflow: hidden;
  width: 100%;
  min-height: 100vh;
}

.bottom {
  position: absolute;
  // float: right;
  /* bottom: -100vh; */
  overflow: hidden;
  width: 100%;
  min-height: 100vh;
  background-color: bisque;
  transition: all 2s;
  bottom: -100vh;
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
}

.app {
  background-size: cover;
  background-repeat: repeat-y;
  // background-image: url("./assets/wallhaven-l338ky.jpg");
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