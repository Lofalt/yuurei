<template>

  <div class="card">
    <!--      <n-space vertical>-->
    <div style="color:rgba(49,49,49,.3)">其实注册了也没什么用</div>
    <div class="cardContainer">
        <span class="label">
        昵称:
          </span>
      <input class="input" type="text" v-model="name" placeholder="输入昵称"/>
    </div>
    <div class="cardContainer">
        <span class="label">
        邮箱地址:
          </span>
      <input class="input" @change.lazy="checkName" type="text" v-model="username" placeholder="输入用户名"/>
      <div class="error">{{ nameErrorMessage }}</div>
    </div>
    <div class="cardContainer">
        <span class="label">
        密码:
          </span>
      <input class="input" @keyup.enter="signup" @change.lazy="checkPass" type="password" v-model="password"
             placeholder="输入密码"/>
      <div class="error">{{ passErrorMessage }}</div>

    </div>
    <div class="cardContainer">
      <span class="label">
        验证码:
          </span>
      <input class="input" @keyup.enter="signup" @change.lazy="checkPass" type="password" v-model="validCode"
             placeholder="输入验证码"/>
      <button class="validBtn" @click="sendEmail" v-show="!sendedEmail">发送验证码</button>
      <span class="validMessage" v-show="sendedEmail">{{ timeout }}秒后重新发送</span>
      <div class="error">{{ codeErrorMessage }}</div>

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
const validCode = ref("")

function signup() {
  checkName()
  checkPass()
  checkCode()
  if (!checkName() || !checkPass() || !checkCode()) {
    return
  }
  axios.request("/yuurei/sign_up", "post", qs.stringify({
    username: username.value,
    name: name,
    password: password.value,
    code: validCode.value
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

const sendedEmail = ref(false)

function sendEmail() {
  if (!checkName()) {
    return
  }
  let timer = reset()
  axios.request("/yuurei/validEmail", "post", qs.stringify({
    Email: username.value
  })).then((res: any) => {
    if (res.code != 200) {
      sendedEmail.value = false
      clearTimeout(timer)
      nameErrorMessage.value += res.msg

    }
  }).catch(() => {
    clearTimeout(timer)
    sendedEmail.value = false
  })
}

const timeout = ref(10)

function reset() {
  sendedEmail.value = true
  timeout.value = 60
  let timer = setInterval(() => {
    timeout.value -= 1
  }, 1000)
  setTimeout(() => {
    sendedEmail.value = false
    clearInterval(timer)
  }, 60000)
  return timer
}

const codeErrorMessage = ref("")

function checkCode() {
  codeErrorMessage.value = ""
  if (validCode.value == "") {
    codeErrorMessage.value += "验证码不能为空"
  }
  return codeErrorMessage.value == ""
}

function checkName() {
  nameErrorMessage.value = ''
  if (username.value == "") {
    nameErrorMessage.value += " 用户名不能为空"
  } else if (username.value.length < 6) {
    nameErrorMessage.value += "用户名过短"
  } else if (!/^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/.test(username.value)) {
    nameErrorMessage.value += "请输入正确的邮箱"
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

<!--<style lang="less" scoped>-->


<!--.card {-->
<!--  width: 600px;-->
<!--  height: 280px;-->
<!--  background-color: #fff;-->
<!--  border-radius: 10px;-->
<!--  display: flex;-->
<!--  align-items: center;-->
<!--  justify-content: center;-->
<!--  flex-direction: column;-->
<!--  //margin-bottom: 20px;-->
<!--  @media (max-aspect-ratio: 9/16) {-->
<!--    width: 100vw;-->
<!--    height: 70vw;-->
<!--  }-->
<!--  @media (max-height: 500px){-->
<!--  width: 100vw;-->
<!--  height: 70vw;-->
<!--}-->


<!--  .cardContainer {-->
<!--    margin: 15px auto;-->
<!--    //margin-bottom: 20px;-->
<!--    display: flex;-->
<!--    position: relative;-->
<!--    justify-content: space-around;-->
<!--    width: 80%;-->
<!--    text-wrap: nowrap;-->

<!--    &:nth-child(4) {-->
<!--      margin-bottom: 30px;-->
<!--    }-->

<!--    .label {-->
<!--      display: inline-block;-->
<!--      width: 80px;-->
<!--    }-->
<!--  }-->
<!--}-->

<!--.btn {-->
<!--  margin: 0px 20px;-->
<!--  background-color: #fff;-->
<!--  color: rgba(49, 49, 49, .8);-->
<!--  border: 1px solid rgba(49, 49, 49, .2);-->
<!--  padding: 5px 15px;-->
<!--  border-radius: 5px;-->
<!--  cursor: pointer;-->

<!--  &:hover {-->
<!--    color: white;-->
<!--    background-color: var(&#45;&#45;button-color);-->
<!--  }-->
<!--}-->

<!--.error {-->
<!--  position: absolute;-->
<!--  bottom: -25px;-->
<!--  color: #ff4d4d;-->
<!--}-->
<!--</style>-->


<style lang="less" scoped>


.card {
  width: 90%;
  //height: 1;
  background-color: #fff;
  border-radius: 10px;
  display: flex;
  align-items: center;
  margin-top: 1%;
  justify-content: flex-start;
  flex-direction: column;
  //margin-bottom: 20px;
  @media (max-aspect-ratio: 9/16) {
    width: 95vw;
    //height: 50vw;
  }

  @media (max-height: 500px) {
    width: 100vw;
    //height: 50vw;
  }

  .cardContainer {
    position: relative;
    margin: 1.5vh auto;
    //margin-bottom: 20px;
    display: flex;
    justify-content: center;
    align-items: center;
    width: 90%;
    text-wrap: nowrap;

    @media (max-aspect-ratio: 9/16) {
      width: 100%;
    }

    .input {
      border: .3vh solid black;
      width: 70%;
      padding: 2% 3%;
      margin-right: 1vh;
    }

    &:nth-child(5) {
      margin-bottom: 2.5vh;
    }

    .label {
      display: inline-block;
      width: 15%;
    }
  }
}

.footer {
  width: 100%;
  height: 20%;
  //display: flex;
  //justify-content: ;
  align-items: center;
  text-align: center;
}

.btn {
  margin: 0px auto;
  font-size: 1.1em;
  background-color: #fff;
  color: rgba(49, 49, 49, .8);
  border: .3vh solid black;
  padding: 2% 4%;
  border-radius: 5px;
  cursor: pointer;


  &:hover {
    color: white;
    background-color: var(--button-color);
  }
}

.error {
  width: 100%;
  position: absolute;
  bottom: -2.5vh;
  text-align: center;
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
.validMessage {
  position: absolute;
  display: inline-block;
  right: 6vh;
}

.validBtn {
  position: absolute;
  right: 6vh;
  font-size: 1em;
  padding: .5vh 2vh;
  border: none;
  background-color: #fff;
  border: .3vh solid black;
  border-radius: .5vh;
  cursor: pointer;

  &:hover {
    color: white;
    background-color: var(--button-color);
  }
}
</style>