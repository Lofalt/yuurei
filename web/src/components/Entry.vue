<template>
  <div class="entry">
    <div class="entry-pic">
      <n-rate color="var(--button-color)" class="star" readonly :default-value="4" allow-half/>
    </div>
    <div class="entry-right">
      <div class="entry-title">
        捡来的男人
      </div>
      <div class="entry-summary">

        &nbsp;&nbsp;
        &nbsp;&nbsp;
        松户谕（仲野太贺饰）虽然有“想成为演员”的梦想，但没有门路和金钱。
        <br/>
        他离开故乡兵库西宫，投靠青梅竹马杉田（大东骏介饰），住进东京表参道的公寓。
        <br/>
        有一天，他在自动贩卖机下面捡到一张机票，发现机票的主人是模特事务所社长山村（药师丸博子饰）。因为厚脸皮和昭和风格的样貌被山村看中，谕的演员人生突然拉开帷幕。就在这时，得知音信全无的哥哥武志（草彅刚饰）在美国病倒，他踏上了去美国接哥哥的旅程
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {NRate} from 'naive-ui'



</script>

<style lang="less" scoped>

.entry {
  position: relative;
  //background-color: #2dc26b;
  //border: 5px solid rgba(49, 49, 49, .8);
  //margin: 0px 100px 100px 100px;
  margin-bottom: 4vh;
  //margin-top: 4vh;
  margin-right: 1vh;
  width: 60vh;
  height: 35vh;
  float: left;
  cursor: pointer;
  transition: all .2s ease;
  //background-color: black;
  //
  //@media (max-aspect-ratio: 16/12){
  //  width: 80vw;
  //  height:40vw;
  //}
  @media (max-aspect-ratio: 1/1){
    width:80vw;
    height:46vw;
  }


  &:hover {
    transform: translate(-10px, -10px);
    box-shadow: 3px 3px 10px 4px rgba(49, 49, 49, .2);

  }

  //@media (max-width: 800px) {
  //  width: 300px;
  //  height: 200px;
  //  overflow-x: hidden;
  //}

  &:nth-child(2n+0) {
  }

  &::after,
  &::before {
    content: ""; //设置内容为空

    height: 0; //高度为0

    line-height: 0; //行高为0

    display: block; //将文本转为块级元素

    visibility: hidden; //将元素隐藏

    clear: both //清除浮动
  }

  .entry-pic {
    width: 20.9vh;
    height: 30vh;
    background-color: #fff;
    position: relative;
    top: 2vh;
    //border: .3vh solid rgb(49, 49, 49);

    left: 2vh;
    float: left;
    background-image: url("../assets/p2874590316.webp");
    background-size: cover;

    //@media (max-width: 800px) {
    //  width: 115px;
    //  height: 165px;
    //}
    @media (max-aspect-ratio: 1/1){
      width:27vw;
      height:39.23vw;
    }

    .star{
      position: absolute;
      width: 100%;
      text-align: center;
      bottom:-3vh;
      left:0;
      display: flex;
      justify-content: center;
    }
  }

  .entry-right {
    position: relative;
    right: 15px;
    top: 20px;
    float: right;
    width: 33vh;
    background-color: #fff;
    height: 30vh;
    text-align: center;
    @media (max-aspect-ratio: 1/1){
      width:40vw;
      height:39.23vw;
    }


    //@media (max-width: 800px) {
    //  width: 140px;
    //  height: 165px;
    //}

    .entry-title {
      margin-top: 10px;
      font-size: 1.2em;
      font-family: '黑体', serif;
      font-weight: bold;
      position: relative;

      &::after {
        content: '';
        //visibility: hidden;
        width: 60%;
        height: 100%;
        border-bottom: 1px solid rgba(49, 49, 49, .1);
        //position: relative;
        position: absolute;
        left: 6vh;
        //background-color: #000;
        top: 3px;

        @media (max-width: 800px) {
          left: 28px;
        }


      }

      @media (max-width: 800px) {
        font-size: 1.1em;
        margin-top: 0;
      }

    }

    .entry-summary {
      color: #094a70;
      text-overflow: ellipsis;
      overflow: hidden;
      //height: 80%;
      width: 100%;
      padding-top: 10px;
      padding-left: 10px;
      padding-right: 10px;
      display: -webkit-box;
      -webkit-line-clamp: 10;
      -webkit-box-orient: vertical;
      //white-space: nowrap;

      @media (max-aspect-ratio: 6/10) {
        -webkit-line-clamp: 7;

      }
      @media (max-aspect-ratio: 9/16) {
        -webkit-line-clamp: 5;

      }
    }
  }

}
</style>