<template>

  <div class="card">
    <!--      <n-space vertical>-->
    <div class="cardContainer">
        <span class="label">
        用户名:
          </span>
      <input @keyup.enter="login" class="input" type="text" @change.lazy="checkName" v-model="username"
             placeholder="输入邮箱地址"/>
      <div v-show="!hasName" class="error">用户名不能为空</div>
    </div>
    <div class="cardContainer">
        <span class="label">
        密码:
          </span>
      <input class="input" @keyup.enter="login" @change.lazy="checkPass" type="password" v-model="password"
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
  width: 90%;
  //height: 1;
  background-color: #fff;
  border-radius: 10px;
  display: flex;
  align-items: center;
  margin-top: 2%;
  justify-content: flex-start;
  flex-direction: column;
  //margin-bottom: 20px;
  @media (max-aspect-ratio: 9/16) {
    width: 95vw;
    height: 50vw;
  }

  @media (max-height: 500px) {
    width: 100vw;
    height: 50vw;
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

    &:nth-child(2) {
      margin-bottom: 30px;
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
  padding: 1% 3%;
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
</style>