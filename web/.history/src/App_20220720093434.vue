<template>
  <div class="app">
    <left-bar></left-bar>
  </div>
</template>

<script lang="ts" setup>
import LeftBar from './components/LeftBar.vue';
</script>

<style scoped>
.app {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
}
</style>