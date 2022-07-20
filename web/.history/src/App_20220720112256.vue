<template>
  <div class="app">
    <left-bar></left-bar>
    <right-bar></right-bar>
  </div>
</template>

<script lang="ts" setup>
import LeftBar from './components/LeftBar.vue';
import RightBar from './components/RightBar.vue'
</script>

<style >
#app {
  min-height: 100vh;
}

.app {
  position: relative;
  min-height: 100vh;
}
</style>