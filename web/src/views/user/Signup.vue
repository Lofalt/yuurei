<template>

  <div class="card">
    <!--      <n-space vertical>-->
    <div style="color:rgba(49,49,49,.3)">其实注册了也没什么用</div>
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
      <n-input @change.lazy="checkName" type="text" v-model:value="username" placeholder="输入用户名"/>
      <div class="error">{{ nameErrorMessage }}</div>
    </div>
    <div class="cardContainer">
        <span class="label">
        密码:
          </span>
      <n-input @keyup.enter="signup" @change.lazy="checkPass" type="password" v-model:value="password"
               placeholder="输入密码"/>
      <div class="error">{{ passErrorMessage }}</div>
    </div>
    <div class="footer">
      <button @click="signup" class="btn" type="info">注册</button>
    </div>
    <!--      </n-space>-->

  </div>


</template>

<script lang="ts" setup>
import {NCard, NSpace, NInput, NButton, useMessage} from "naive-ui"
import {ref} from "vue";
import axios from "@/request"
import qs from 'qs'
import {useRouter} from "vue-router";

const router = useRouter()
const message = useMessage()
const name = ref("")
const username = ref("")
const password = ref("")
const hasPass = ref(true)
const hasName = ref(true)
const wrongName = ref(true)
const wrongPass = ref(true)
const shortName = ref(true)
const shortPass = ref(true)
const passErrorMessage = ref("")
const nameErrorMessage = ref("")

function signup() {
  checkName()
  checkPass()
  if (!checkName() || !checkPass()) {
    return
  }
  axios.request("/yuurei/sign_up", "post", qs.stringify({
    username: username.value,
    name: name,
    password: password.value,
  })).then((res: any) => {
    if (res.code == 200) {
      message.info(res.msg + "马上跳转登陆")
      setTimeout(() => {
        router.go(0)
      }, 1800)
    } else {
      message.error(res.msg)
    }
  })
}

function checkName() {
  nameErrorMessage.value = ''
  if (username.value == "") {
    nameErrorMessage.value += " 用户名不能为空"
  } else if (username.value.length < 6) {
    nameErrorMessage.value += "用户名过短"
  }
  if (username.value.indexOf(" ") >= 0 || username.value.indexOf("?") >= 0) {
    nameErrorMessage.value += " 包含非法字符"
  }
  return nameErrorMessage.value == ""
}

function checkPass() {
  passErrorMessage.value = ''
  if (password.value == "") {
    passErrorMessage.value += " 密码不能为空"
  } else if (password.value.length < 6) {
    passErrorMessage.value += "密码过短"
  }
  if (password.value.indexOf(" ") >= 0 || password.value.indexOf("?") >= 0) {
    passErrorMessage.value += " 包含非法字符"
  }

  return passErrorMessage.value == ""
}
</script>

<style lang="less" scoped>


.card {
  width: 600px;
  height: 280px;
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
    margin: 15px auto;
    //margin-bottom: 20px;
    display: flex;
    position: relative;
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
</style>