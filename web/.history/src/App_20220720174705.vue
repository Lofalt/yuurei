<template>
  <div class="app">
    <left-bar></left-bar>
    <div>
      <router-view></router-view>
    </div>
    <div>
      <router-view></router-view>
    </div>
    <div>
      <router-view></router-view>
    </div>
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

.card-enter-from,
.card-leave-to {
  /* opacity: 0; */
  transform: translateY(1000px);
}

.card-enter-active,
.card-leave-active {
  transition: transform 1.5s ease
}

.card-leave-to {
  display: none;
}
</style>