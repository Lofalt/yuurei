<template>
  <div style="font-size: 3vh;font-weight: bold;border-bottom: 3px solid black;margin:4vh;text-align: center">
    这不是我的故事,而是它的故事
  </div>
    <div class="container">
      <div id="img1">
        <img style="width: 100%" src="../../assets/Snipaste_2022-08-06_11-36-10.png"/>
        <div style="color:rgba(49,49,49,.4);text-align: center">(甚至还写了图形界面)</div>
      </div>
      <p>之前闲着无聊学日语的时候，会习惯把记不住的单词记到一个excel文件上.后来记录的词多了，就想着有什么方法可以随机抽查一下。</p>
      <p>然后想起自己学过一点点编程。于是就用尽“毕生所学”写了一个导入excel背单词的小程序出来(右图👉)</p>
      <p>有听写有翻译，虽然很简单但也够用了.但事实上我也并没有用过几次就是了。</p>
      <br/>
      <p> 在写这个东西的过程中，又开始回想起当初自学编程时的那一点点乐趣，虽然自己的脑子不太好使，但学习新知识的感觉还是非常不错的。</p>
      <p>然后在写完这个小程序后，又勾起了我的一个执念。</p>
      <br/>
      <p>很多年前，在我还很喜欢写东写西的时候，在浏览过一个网友的个人网站之后，便对拥有属于自己的个人网络空间这件事产生了憧憬，后来就成了一个执念。</p>
      <p>如今我已经不再喜欢写有的没的了，但这个执念却一直没有消散</p>
      <p>虽然在两年前也曾经尝试写过一次，但写了一半被我丢上服务器运行之后就再也没有写下去的动力了，最后也就不了了之。</p>
      <img style="width: 100%;" src="../../assets/72222614d313d458.png"/>
      <div style="text-align: center;line-height: 3vh;color:rgba(49,49,49,.3)">(前不久在本地试着跑了一下，没想到还能动起来)</div>
      <div class="hr"></div>
      <p>至于为什么会失去动力，大概是那时的我觉得做这种事情没有意义吧，就放弃了.</p>
      <p>至于为什么我这次又重新开始做这种事情，大概是因为我已经不在乎意义了吧.</p>
      <br/>
      <p>在动手之前，在PS画了一下设计图:</p>
      <img style="width: 100%" src="../../assets/Snipaste_2022-07-19_22-50-42.png"/>
      <p>在多次的修修改改之后，最后变成了现在这个样子</p>
      <p>反正，我自己还是挺喜欢现在这个样式的</p>
    </div>
</template>

<script lang="ts" setup>


</script>

<style lang="less" scoped>

.container {
  font-family: 微软雅黑;
  color: rgb(69, 69, 69,);
  width: 90%;
  height: 90%;
  line-height: 3vh;
  margin-top: 3vh;
  font-size: 1.8vh;

  @media (min-width: 1600px) {
    //font-size: 1.2em;
    line-height: 3vh;

  }

  p {
    margin-top: 10px;
  }
}

#img1 {
  margin-left: 30px;
  padding-top: 1vh;
  line-height: 3vh;
  display: block;
  display: inline-block;
  float: right;
  //height: 60vh;
  max-width: 10vw;
  margin-bottom: 4vh;


  @media (max-aspect-ratio: 9/16) {
    max-width: 30vw;
  }
}

.hr {
  margin: 20px auto;
  margin-top: 20px;
  width: 100%;
  height: 1px;
  border-bottom: 1px solid rgba(49, 49, 49, .1)
}
</style>