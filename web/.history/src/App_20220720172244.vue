<template>
  <div class="app">
    <left-bar></left-bar>
    <!-- <router-view></router-view> -->
    <router-view v-slot="{ Component }">
      <transition name="card">
        <component :is="Component" />
      </transition>
    </router-view>
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
}

.card-enter-from,
.card-leave-to {
  transform: translateY(1000px);
}

.card-enter-active,
.card-leave-active {
  transition: transform 1.2s cubic-bezier(0.075, 0.82, 0.165, 1)
}
</style>