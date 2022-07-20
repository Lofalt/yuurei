<template>
  <div class="app">
    <left-bar></left-bar>
    <div class="rightbar">
      <div class="top">
        top
      </div>
      <div v-show="isShow" class="middle">
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
  console.log(isShow.value)
  isShow.value = !isShow.value
}
</script>

<style scoped>
.rightbar {
  float: right;
  position: relative;
  overflow: hidden;
  width: 100%;
  min-height: 100vh;
  min-height: 100vh;
  width: 83.5vw;
  overflow: hidden;
}

.rightbar {
  /* background-color: #fff; */
  width: 83.5%;
  min-height: 100vh;
  float: right;
}

.top {
  position: relative;
  float: right;
  top: -100vh;
}

.bottom {
  position: relative;
  float: right;
  /* bottom: -100vh; */
  overflow: hidden;
  width: 100%;
  min-height: 100vh;
  background-color: bisque;
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
  overflow: scroll;
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