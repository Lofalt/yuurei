<template>

  <div class="card">
    <!--      <n-space vertical>-->
      <div>其实注册了也没什么用</div>
    <div class="cardContainer">
        <span class="label">
        昵称:
          </span>
      <n-input type="text" v-model:value="name" placeholder="输入昵称"/>
    </div>
    <div class="cardContainer">
        <span class="label">
        用户名:
          </span>
      <n-input type="text" v-model:value="username" placeholder="输入用户名"/>
    </div>
    <div class="cardContainer">
        <span class="label">
        密码:
          </span>
      <n-input type="password" v-model:value="password" placeholder="输入密码"/>
    </div>
    <div class="footer">
      <n-button @click="signup" class="btn" type="info">注册</n-button>
    </div>
    <!--      </n-space>-->

  </div>


</template>

<script lang="ts" setup>
import {NCard, NSpace, NInput, NButton} from "naive-ui"
import {ref} from "vue";
import axios from "@/request"
import qs from 'qs'

const name = ref("")
const username = ref("")
const password = ref("")

function signup() {
  axios.request("/yuurei/sign_up", "post", qs.stringify({
    username: username.value,
    name: name,
    password: password.value,
  })).then((res) => {
    console.log(res)
  })
}
</script>

<style lang="less" scoped>


.card {
  width: 600px;
  height: 240px;
  background-color: #fff;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  //margin-bottom: 20px;
  @media (max-width: 600px) {
    width: 100vw;
    height: 250px;
  }

  .cardContainer {
    margin: 10px auto;
    //margin-bottom: 20px;
    display: flex;
    justify-content: space-around;
    width: 80%;
    text-wrap: nowrap;

    &:nth-child(4) {
      margin-bottom: 30px;
    }

    .label {
      display: inline-block;
      width: 80px;
    }
  }
}

.btn {
  margin: 0px 20px;
}

</style>