<template>
  <div class="container"
       :style="`--entry-main-color:${Entry.MainColor};--entry-sec-color:${Entry.SecColor};--entry-third-color:${Entry.ThirdColor}`">
    <div class="actualHeader">
      <h3>{{ Entry.Quote }}
      </h3></div>
    <div class="header">
      <div class="profile">
        <br/>
        <h1>东京事变</h1>
        <br/>
        <h2>作品简介</h2>
        <div v-html="Entry.Detail"></div>
      </div>
      <div class="info">
        <div class="infoHeader">
          {{ Entry.Title }}
        </div>
        <img :src="Entry.HeaderPicture" class="infoPic"/>
        <div class="table">
          <table class="tb">
            <tr v-for="(item,index) in Entry.Infos" :key="item.ID">
              <td class="tbF">{{ item.Label }}</td>
              <td class="tbS">{{ item.Value }}</td>
            </tr>
          </table>
        </div>
      </div>
    </div>
    <div style=";width: 100%" class="body">
      <h1>我的话</h1>
      <div v-html="Entry.MyWord"></div>
    </div>
    <br/>
    <br/>
    <div style="min-height: 100vh;width: 100%" class="body">
      <h1>相关影音</h1>
      <hr/>
      <div v-html="Entry.About"></div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import axios from "../../request/index"
import {ref} from 'vue'

const Entry = ref({
  Title: "",
  Quote: "",
  Summary: "",
  Detail: "",
  MyWord: "",
  About: "",
  HeaderPicture: "",
  Infos: [],
  MainColor: "",
  SecColor: "",
  ThirdColor: "",
  EntryCategoryID: "",
})

axios.get("/yuurei/entry/6", {}).then((res: any) => {
  Entry.value = res.data.data
})
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
      //background-color: #20bbbb;
      background-color: Var(--entry-main-color);
    }
  }

  .actualHeader {
    font-size: 1.5em;
    background-color: var(--entry-third-color);
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
      background-color: var(--entry-sec-color);
    }

    h3 {
      line-height: 10vh;
      display: inline-block;
      padding-left: 1.7em;
      position: relative;

      @media (max-aspect-ratio: 9/16) {
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
        background-color: var(--entry-third-color);
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
              background-color: var(--entry-sec-color);
            }

            &:nth-child(2) {
              background-color: var(--entry-third-color);
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


  @media (max-aspect-ratio: 9/16) {
    //height: 50vw;
    min-height: 0;
    //height: 100px;
  }
}


</style>