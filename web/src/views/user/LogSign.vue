<template>
  <div class="container">
    <div class="inside">
      <div class="header">
        <button :class="isActive('Login')?'activeCard':''" class="btn" @click="change(`Login`)">登陆</button>
        <button :class="isActive('Signup')?'activeCard':''" class="btn" @click="change(`Signup`)">注册</button>
      </div>
      <n-message-provider>
        <component :is="component"/>
      </n-message-provider>

      <!--      <n-message-provider>-->
      <!--        <login/>-->
      <!--      </n-message-provider>-->
      <!--      <n-message-provider>-->
      <!--        <signin/>-->
      <!--      </n-message-provider>-->
    </div>
  </div>


</template>

<script lang="ts" setup>
import {NTabs, NTabPane, NMessageProvider} from "naive-ui"
import Login from "./Login.vue"
import Signup from "./Signup.vue"
import {ref} from 'vue'

const activeCard = ref("Login")
const component = ref(Login) as any

function change(target: string) {
  if (target == "Login") {
    component.value = Login
    activeCard.value = "Login"
  } else {
    component.value = Signup
    activeCard.value = "Signup"
  }
}

function isActive(arg: string) {
  return arg == activeCard.value
}

</script>

<style lang="less" scoped>
.container {
  font-size: 1.5vh;
  float: right;
  width: 83.5%;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;

  @media (max-aspect-ratio: 1/1) {
    width: 100%;
  }
  @media (max-height: 500px) {
    width: 100%;

  }


  .inside {
    //border: 4px solid rgb(49, 49, 49);

    display: flex;
    justify-content: flex-start;
    align-items: center;
    flex-direction: column;
    //width: 50%;
    background-color: #fff;
    padding: 10px;
    border-radius: 10px;
    //min-height: 50vh;
    width: 80vh;
    border: .4vh solid black;

    .header {
      display: flex;
      justify-content: center;
      //height: 23%;
      width: 100%;

      .btn {
        padding: 1% 2%;
        margin: 2% 4%;
        margin-top: 5%;
        border: none;
        font-size: 1.1em;
        border-radius: 10%;
        cursor: pointer;
        background: none;
        border: .3vh solid black;

        &:hover {
          color: white;
          background-color: var(--button-color);
        }
      }
    }

    @media (max-aspect-ratio: 1/1) {
      width: 80vw;
      min-height: 50vw;
    }

    @media (max-aspect-ratio: 9/16) {
      width: 100vw;
      min-height: 62.5vw;
    }
  }

  .activeCard {
    color: white;
    background-color: var(--button-color) !important;
  }
}
</style>