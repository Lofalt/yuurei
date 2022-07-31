<template>
  <div class="container">
    <div class="nav">navbar</div>
    <div class="inside">
      <entry/>
      <entry/>
      <entry/>
      <entry/>
      <entry/>
      <entry/>
      <entry/>
      <entry/>
      <entry/>
    </div>
  </div>
</template>

<script lang="ts" setup>
import Entry from "../components/Entry.vue"


</script>

<style lang="less" scoped>
.container {
  width: 100%;
  padding: 20px;
  overflow: auto;
  height: 100%;
  margin: 0 auto;
  overflow-x: hidden;

  .nav {
    height: 40px;
    text-align: center;
    background-color: #ffb0d1;
  }

  .inside {
    text-align: center;
    margin: 0 auto;
    width: 100%;
    height: 100%;
    padding-left: 70px;
    display: flex;
    justify-content: space-around;
    flex-wrap: wrap;
    @media (max-width: 800px) {
      padding-left: 0;
    }
  }

  &::-webkit-scrollbar {
    width: 5px;
  }

  &::-webkit-scrollbar-thumb {
    width: 5px;
    background-color: rgba(49, 49, 49, .2);
  }

  @media (max-width: 800px) {
    width: 100%;
  }

}


</style>