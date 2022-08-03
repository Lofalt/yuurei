<template>

  <div class="card">
    <!--      <n-space vertical>-->
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
      <n-input @keyup.enter="login" type="password" v-model:value="password" placeholder="输入密码"/>
    </div>
    <div class="footer">
      <n-button @click="login" class="btn" type="info">登陆</n-button>
    </div>
    <!--      </n-space>-->

  </div>


</template>

<script lang="ts" setup>
import {NCard, NSpace, NInput, NButton} from "naive-ui"
import {ref} from "vue";
import axios from "@/request"
import qs from "qs"
import {useRouter} from "vue-router";

const username = ref("")
const password = ref("")
const router = useRouter()

function login() {
  axios.request("/yuurei/login", "post", qs.stringify({
    username: username.value,
    password: password.value,
  })).then((res: any) => {
    console.log(res)
    localStorage.setItem('token', res.data.token)
    router.back()
  })
}
</script>

<style lang="less" scoped>


.card {
  width: 500px;
  height: 200px;
  background-color: #fff;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  //margin-bottom: 20px;
  @media (max-width: 600px) {
    width: 100vw;
    height: 200px;
  }

  .cardContainer {
    margin: 10px auto;
    //margin-bottom: 20px;
    display: flex;
    justify-content: space-around;
    width: 80%;
    text-wrap: nowrap;

    &:nth-child(2) {
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