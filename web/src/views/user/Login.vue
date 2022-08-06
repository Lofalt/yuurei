<template>

  <div class="card">
    <!--      <n-space vertical>-->
    <div class="cardContainer">
        <span class="label">
        用户名:
          </span>
      <n-input class="input" type="text" @change.lazy="checkName" v-model:value="username"
               placeholder="输入用户名"/>
      <div v-show="!hasName" class="error">用户名不能为空</div>
    </div>
    <div class="cardContainer">
        <span class="label">
        密码:
          </span>
      <n-input @keyup.enter="login" @change.lazy="checkPass" type="password" v-model:value="password"
               placeholder="输入密码"/>
      <div v-show="!hasPass" class="error">密码不能为空</div>
    </div>

    <div class="footer">
      <button @click="login" class="btn" type="info">登陆</button>
    </div>
    <!--      </n-space>-->

  </div>


</template>

<script lang="ts" setup>
import {NCard, NSpace, NInput, NButton, useMessage} from "naive-ui"
import {ref} from "vue";
import axios from "@/request"
import qs from "qs"
import {useRouter} from "vue-router";

const username = ref("")
const password = ref("")
const router = useRouter()
const hasName = ref(true)
const hasPass = ref(true)
const message = useMessage()

function checkName() {
  if (username.value == "") {
    hasName.value = false
  } else {
    hasName.value = true
  }
}

function checkPass() {
  if (password.value == "") {
    hasPass.value = false
  } else {
    hasPass.value = true
  }
}

function login() {
  checkPass()
  checkName()
  if (!hasName.value || !hasPass.value) {
    return
  } else {
    axios.request("/yuurei/login", "post", qs.stringify({
      username: username.value,
      password: password.value,
    })).then((res: any) => {
      console.log(res)
      if (res.code == 200) {
        localStorage.setItem('token', res.data.token)
        message.info("登陆成功,马上跳转")
        setTimeout(() => {

          router.back()
        }, 3000)
      } else {
        message.error(res.msg)
      }
    })
  }
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
    position: relative;
    margin: 15px auto;
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
  background-color: #fff;
  color: rgba(49, 49, 49, .8);
  border: 1px solid rgba(49, 49, 49, .2);
  padding: 5px 15px;
  border-radius: 5px;
  cursor: pointer;

  &:hover {
    color: white;
    background-color: var(--button-color);
  }
}

.error {
  position: absolute;
  bottom: -25px;
  color: #ff4d4d;
}

//.n-input {
//  outline: none !important;
//  --n-border-hover: 1px solid black !important;
//  --n-border-focus: 1px solid black !important;
//  --n-loading-color: black !important;
//  --n-caret-color: black !important;
//  --n-border: 3px solid black !important;
//  --n-bar-color: black !important;
//}
//
//.input {
//  outline: none;
//  //border: 3px solid black;
//}
</style>