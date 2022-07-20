<template>
  <div class="app">
    <left-bar></left-bar>
    <button @click.self="change">点我</button>
    <!-- <div class="top" v-show="!isShow">
      <router-view></router-view>
    </div> -->
    <div v-show="isShow">
      <router-view></router-view>
    </div>
    <!-- <div class="bottom">
      <router-view></router-view>
    </div> -->
    <!-- <router-view v-slot="{ Component }">
      <transition name="card">
        <component :is="Component" />
      </transition>
    </router-view> -->
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

/* .top {
  position: absolute;
  right: 1000px;
  animation: move 1s ease;
} */

@keyframes move {
  to {
    transform: translateX(-1000px);
  }
}
</style>