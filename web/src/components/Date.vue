<template>
  <div class="date">
    <div class="time">
      {{time}}
    </div>
    <div class="day">
    {{day +"\n\r"+ youbi }}
    </div>
  </div>
</template>

<script lang="ts" setup>

import {ref} from "vue";

const time = ref(new Date().toLocaleTimeString())
const day  = ref(new Date().toLocaleDateString())
const youbi = ref("("+ "日月火水木金土".charAt(new Date().getDay())+")")

setInterval(()=>{
  time.value = new Date().toLocaleTimeString()
  day.value = new Date().toLocaleDateString()
},1000)

</script>

<style lang="less" scoped>

.date{
  width: 1000px;
  height: 1000px;
  //background-color: #fff;
  position: absolute;
  top:20vh;
  left:21vw;

  @media (max-width: 800px){
    display: none;
  }
  .time{
    font-size: 5em;
    font-family: "等线 Light","微软雅黑 Light","锐字工房洪荒之力中黑简1.0","锐字锐线俏皮简1.0","方正精气神体 简 ExtraLight";
  }

  .day{
    font-size: 2em;
  }
}

</style>