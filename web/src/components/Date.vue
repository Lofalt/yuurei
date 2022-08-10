<template>
  <div class="date">
    <div class="time">
      {{ time }}
    </div>
    <div class="day">
      {{ day + "\n\r" + youbi }}
    </div>
    <div class="word">
      立秋了啊。
    </div>
  </div>
</template>

<script lang="ts" setup>

import {ref} from "vue";

interface Date {
  Format(param: string): void;
}

Date.prototype.Format = function (fmt: string) {
  var o = {
    "M+": this.getMonth() + 1, //月份
    "d+": this.getDate(), //日
    "h+": this.getHours(), //小时
    "m+": this.getMinutes(), //分
    "s+": this.getSeconds(), //秒
    "q+": Math.floor((this.getMonth() + 3) / 3), //季度
    "S": this.getMilliseconds() //毫秒
  };
  if (/(y+)/.test(fmt)) fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
  for (var k in o)
    if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
  return fmt;
}
const time = ref((new Date()).Format("hh:mm:ss"))
const day = ref((new Date()).Format("yyyy/M/d"))
const youbi = ref("(" + "日月火水木金土".charAt(new Date().getDay()) + ")")


setInterval(() => {
  time.value = (new Date()).Format("hh:mm:ss")
  day.value = (new Date()).Format("yyyy/M/d")
}, 1000)

</script>

<style lang="less" scoped>

.date {
  user-select: none;
  width: 1000px;
  height: 1000px;
  //background-color: #fff;
  color: var(--date-color);
  position: absolute;
  top: 20vh;
  left: 5vw;

  @media (max-width: 800px) {
    //display: none;
  }

  .time {
    font-size: 5em;
    font-family: "等线 Light", "微软雅黑 Light", "锐字工房洪荒之力中黑简1.0", "锐字锐线俏皮简1.0", "方正精气神体 简 ExtraLight";
  }

  .day {
    font-size: 2em;
  }

  .word {
    font-size: 2em;
    margin-top: 2.5vh;
  }
}

</style>