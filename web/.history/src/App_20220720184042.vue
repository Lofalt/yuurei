<template>
  <div class="app">
    <button @click="change">dmwo </button>
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
  actualCard.style.transition = "all 1s ease"
  actualCard.style.transform = "translateY(-100vh)"
  // actualCard.style.transition = ""
  const card2 = document.getElementsByClassName("bottom") as HTMLSelectElement
  const actualCard2 = card2[0]
  actualCard2.style.transition = "all 1s ease"
  actualCard2.style.transform = "translateY(-100vh)"
}

</script>

<style scoped>
.rightbar {
  /* background-color: #fff; */
  width: 83.5%;
  min-height: 100vh;
  float: right;
  overflow: hidden;
}

.top {
  position: relative;
  float: right;
  top: -100vh;
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
  position: relative;
  float: right;
  /* bottom: -100vh; */
  overflow: hidden;
  width: 100%;
  min-height: 100vh;
  background-color: bisque;
  transition: all 2s;
}

button {
  position: absolute;
  left: 10px;
  z-index: 10;
}

#app {
  min-height: 100vh;
}

.app {
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