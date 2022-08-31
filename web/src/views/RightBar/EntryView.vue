<template>
  <div class="container">
    <div class="actualHeader">
      <h3>听JPOP不听东京事变，就像听JPOP不听东京事变。
      </h3></div>
    <div class="header">
      <div class="profile">
        <br/>
        <h1>东京事变</h1>
        <br/>
        <h2>作品简介</h2>
        <br/>
        简介: 东京事变是一支日本的乐团，成立于2003年，乐团于2012年解散，2020年宣布复出。 [1]
        <br/>
        <br/>
        乐团成员现有五人，分别为椎名林檎（主唱）、龟田诚治（贝斯手）、浮云（吉他手）、刃田缀色（鼓手）和伊泽一叶（钢琴手·吉他手）。
        <br/>
        <br/>
        虽然其理念上大体以椎名林檎为中心，但每位乐手们依然各自闪耀，并也各自衬托，他们的风格多变，这也是长久受到观众喜爱和支持的原因。
        <br/>
        <br/>
        2021年发行专辑《Music》，并获得亚洲流行音乐大奖2021年度最佳专辑（海外）
        <br/>
        <br/>
        出生:不知道
        <br/>
        地区:日本
      </div>
      <div class="info">
        <div class="infoHeader">
          东京事变!
        </div>
        <img src="../../assets/s33712391.jpg" class="infoPic"/>
        <div class="table">
          <table class="tb">
            <tr>
              <td class="tbF">原作</td>
              <td class="tbS">石森章太郎<br>是吗</td>
            </tr>
            <tr>
              <td>王四月</td>
              <td><a href="www.baidu.com">不知道捏</a></td>
            </tr>
          </table>
        </div>
      </div>
    </div>
    <div style=";width: 100%" class="body">
      <h1>我的话</h1>
      <hr/>
      比末世代更末世代
      <br>
      <br>
      谁来作你偶像？
      <br>
      <br>
      供仰神，唾弃神，美化神，憎恶神
      <br>
      <br>
      比神更高，站在高处戴着礼帽
      眼露邪恶
      眼角带笑
      <br>
      <br>
      嚣张的吉他
      高敏的歌声
      <br>
      <br>
      比云端更飘渺 站在神的背后狂笑
      <br>
      <br>
      女王。女王！
      <br>
      <br>
      我崇拜你比任何人更多！

    </div>
    <br/>
    <br/>
    <div style="min-height: 100vh;width: 100%" class="body">
      <h1>相关影音</h1>
      <hr/>
      <iframe src="//player.bilibili.com/player.html?aid=87459531&bvid=BV1q741187Lw&cid=149438521&page=1" scrolling="no"
              border="0" frameborder="no" framespacing="0" allowfullscreen="true"></iframe>
    </div>
  </div>
</template>

<script lang="ts" setup>
</script>

<style lang="less" scoped>

.container {
  font-size: 1.8vh;
  padding-left: 5vh;
  display: flex;
  flex-direction: column;
  width: 80%;
  justify-content: flex-start;
  min-height: 100vh;
  //overflow: auto;
  //background-color: ghostwhite;
  background-attachment: unset;

  hr {
    border-collapse: collapse;
    border: 0px solid black;
    height: 1px;
    color: black;
    background-color: black;
    margin: 1vh auto;
  }

  @media (max-aspect-ratio: 11/16) {
    width: 95%;
    padding-left: 0vh;
    align-self: center;
  }

  h1 {
    &::before {
      content: '';
      display: inline-block;
      margin-right: 1vh;
      width: 0.25em;
      height: 1em;
      transform: translateY(15%);
      //padding-top: .5vh;
      border-radius: 10px;
      background-color: #20bbbb;
    }
  }

  .actualHeader {
    font-size: 1.5em;
    background-color: #fdf6ff;
    //height: 10vh;
    padding-left: 3vh;
    margin-top: 1vh;
    //max-width: 30vw;
    position: relative;

    &::before {
      content: '';
      display: block;
      position: absolute;
      width: 2vh;
      height: 100%;
      left: 0;
      background-color: #f3a4a4;
    }

    h3 {
      line-height: 10vh;
      display: inline-block;
      padding-left: 1.7em;
      position: relative;

      @media (max-aspect-ratio: 12/16) {
        line-height: 7vh;
      }

      &::before {
        content: '"';
        font-size: 3em;
        //display: inline-block;
        position: absolute;
        left: 1vh;
        //width: 3em;
        height: 1vh;
        transform: translateY(50%);
        margin-right: 1vh;
        color: gray;
      }
    }
  }

  .header {
    margin-top: 5vh;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    //flex-wrap: wrap;
    width: 100%;

    @media (max-aspect-ratio: 1/1) AND (min-aspect-ratio: 11/16) {
      flex-wrap: nowrap;
    }

    @media (max-aspect-ratio: 11/16) {
      flex-direction: column-reverse;

    }

    .profile {
      //margin-top: 3vh;
      //margin-left: 3vh;
      margin: 3vh 0;
      //margin-right: 0;
      //height: 50vh;
      width: 50vh;
      //background-color: #b1f0f1;

      @media (max-aspect-ratio: 11/16) {
        width: 90%;
      }
    }

    .info {
      //margin-top: 3vh;
      //margin-right: 3vh;
      //margin: 3vh;
      //min-height: 50vh;
      width: 40vh;
      //background-color: #b1f0f1;
      display: flex;
      flex-direction: column;

      @media (max-aspect-ratio: 11/16) {
        width: 100%
      }

      .infoHeader {
        text-align: center;
        height: 8vh;
        background-color: whitesmoke;
        line-height: 8vh;
        font-size: 1.5em;
        font-weight: bold;
      }

      .infoPic {
        //min-height: 50vh;
        //background-image: url("../../assets/s33712391.jpg");
        width: 100%;
        background-position: top;
        background-size: cover;
        background-repeat: no-repeat;

      }

      .table {
        width: 100%;

        .tb {
          width: 100%;
          border-spacing: 0;
          border: 0;
          border-collapse: collapse;
          text-align: center;

          tr td {
            height: 4vh;
            borer: 0;

            &:first-child {
              width: 40%;
              background-color: pink;
            }

            &:nth-child(2) {
              background-color: #b1f0f1;
            }
          }

        }
      }
    }
  }
}

iframe {
  width: 100%;
  min-height: 80vh;


  @media (max-aspect-ratio: 12/16) {
    //height: 50vw;
    min-height: 0;
    //height: 100px;
  }
}


</style>