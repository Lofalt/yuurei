<template>
  <div class="entry">
    <div class="entry-pic">

    </div>
    <div class="entry-right">
      <div class="entry-title">
        椎名林檎
      </div>
      <div class="entry-summary">

        &nbsp;&nbsp;
        &nbsp;&nbsp;
        这里是简介,我要在这里写很多很多很多这里是简介,我要在这里写很多很多很多这里是简介,我要在这里写很多很多很多这里是简介,我要在这里写很多很多很多这里是简介,我要在这里写很多很多很多这里是简介,我要在这里写很多很多很多这里是简介,我要在这里写很多很多很多这里是简介,我要在这里写很多很多很多这里是简介,我要在这里写很多很多很多这里是简介,我要在这里写很多很多很多这里是简介,我要在这里写很多很多很多
      </div>
    </div>
  </div>
</template>

<style lang="less" scoped>

.entry {
  position: relative;
  //background-color: #2dc26b;
  //border: 5px solid rgba(49, 49, 49, .8);
  //margin: 0px 100px 100px 100px;
  margin-bottom: 20px;
  margin-right: 10px;
  width: 450px;
  height: 350px;
  float: left;
  cursor: pointer;
  transition: all .2s ease;

  &:hover {
    transform: translate(-10px, -10px);
    box-shadow: 3px 3px 10px 4px rgba(49, 49, 49, .2);

  }

  @media (max-width: 800px) {
    width: 300px;
    height: 200px;
    overflow-x: hidden;
  }

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
    width: 190px;
    height: 274px;
    background-color: #fff;
    position: relative;
    top: 20px;
    border: 3px solid rgb(49, 49, 49);

    left: 15px;
    float: left;
    background-image: url("../assets/s33712391.jpg");
    background-size: cover;

    @media (max-width: 800px) {
      width: 115px;
      height: 165px;
    }
  }

  .entry-right {
    position: relative;
    right: 15px;
    top: 20px;
    float: right;
    width: 200px;
    background-color: #fff;
    height: 300px;
    text-align: center;

    @media (max-width: 800px) {
      width: 150px;
      height: 165px;
    }

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
        left: 38px;
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

      @media (max-width: 800px) {
        -webkit-line-clamp: 5;

      }
    }
  }

}
</style>