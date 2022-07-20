<template>
  <div class="app">
    <left-bar></left-bar>
    <div class="top" v-show="isShow">
      <router-view></router-view>
    </div>
    <div>
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
import LeftBar from './components/LeftBar.vue';
import RightBar from './components/RightBar.vue'
</script>

<style scoped>
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
} */
</style>