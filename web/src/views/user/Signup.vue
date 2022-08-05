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
      <div v-show="!hasName" class="error">用户名不能为空</div>
      <div v-show="!wrongName" class="error">用户名包含非法字符</div>
      <div v-show="!shortName" class="error">用户名过短(少于6位)</div>
    </div>
    <div class="cardContainer">
        <span class="label">
        密码:
          </span>
      <n-input @keyup.enter="signup" @change.lazy="checkPass" type="password" v-model:value="password"
               placeholder="输入密码"/>
      <div v-show="!hasPass" class="error">密码不能为空</div>
      <div v-show="!wrongPass" class="error">密码包含非法字符</div>
      <div v-show="!shortPass" class="error">密码过短(少于6位)</div>
    </div>
    <div class="footer">
      <n-button @click="signup" class="btn" type="info">注册</n-button>
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
  if (username.value == "") {
    hasName.value = false
  } else {
    hasName.value = true
    if (username.value.indexOf(" ") >= 0 || username.value.indexOf("?") >= 0) {
      wrongName.value = false
    } else {
      wrongName.value = true
      if (username.value.length < 6) {
        shortName.value = false
      } else {
        shortName.value = true
      }
    }
  }
  return hasName.value && wrongName.value && shortName.value
}

function checkPass() {
  if (password.value == "") {
    hasPass.value = false
  } else {
    hasPass.value = true
    if (password.value.indexOf(" ") >= 0 || password.value.indexOf("?") >= 0) {
      wrongPass.value = false
    } else {
      wrongPass.value = true
      if (password.value.length < 6) {
        shortPass.value = false
      } else {
        shortPass.value = true
      }
    }
  }

  return hasPass.value && wrongPass.value && shortPass.value
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
}

.error {
  position: absolute;
  bottom: -25px;
  color: #ff4d4d;
}
</style>