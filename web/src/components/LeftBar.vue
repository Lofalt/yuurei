<template>
  <div class="leftbar">
    <div class="header">{{ WebsiteName }}</div>
    <div class="profile" :style="{backgroundImage:backgroundImage}"></div>
    <div class="name">{{ MyName }}</div>
    <div class="icon1">
      <a href="javascript:">
        <n-icon size="20" color="#484848">
          <grin-tongue/>
        </n-icon>
      </a>
      <a href="javascript:">
        <n-icon size="20" color="#484848">
          <weibo/>
        </n-icon>
      </a>
      <!-- </div> -->
      <!-- <div> -->
      <a href="javascript:">
        <n-icon size="20" color="#484848    ">
          <twitter/>
        </n-icon>
      </a>
      <!-- </div> -->
    </div>
    <div class="footer">
      <button class="footerbox" @click="changePage(0)" :class="isThisPage(0) ? 'active' : ``">HOME</button>
      <button class="footerbox" @click="changePage(1)" :class="isThisPage(1) ? 'active' : ``">INFO</button>
      <button class="footerbox" @click="changePage(2)" :class="isThisPage(2) ? 'active' : ``">ABOUT</button>
      <!--      <button class="footerbox" @click="changePage(3)" :class="isThisPage(3) ? 'active' : ``">WRITE</button>-->
    </div>
    <n-icon v-show="isAdmin" class="settings" size="20" color="#484848" @click="jumpTo(`/admin`)">
      <settings/>
    </n-icon>
    <n-icon class="home" size="20" color="#484848" @click="jumpTo(`/home`)">
      <home/>
    </n-icon>
    <n-icon v-show="!isLogin" class="login" size="20" color="#484848" @click="logIn">
      <login/>
    </n-icon>
    <n-icon v-show="isLogin" class="logout" size="20" color="#484848" @click="logOut">
      <logout/>
    </n-icon>

  </div>
</template>

<script lang="ts">

export default {
  name: "LeftBar",
};
</script>

<script lang="ts" setup>
import {computed, inject, onMounted, reactive, ref} from 'vue';
import {useRouter} from 'vue-router';
import {Weibo, Twitter, GrinTongue} from '@vicons/fa'
import {NIcon} from "naive-ui"
import {usePageData} from '@/store/pageData';
import {Settings, Home, Login, Logout} from '@vicons/tabler'
import {useUserInfo} from "@/store/UserInfo";
import {useDialog} from "naive-ui";

const router = useRouter()
const dialog = useDialog()
const pageData = usePageData()
const userInfo = useUserInfo()
const emit = defineEmits(['toggleNav'])
const config = inject("globalConfig") as any

const backgroundImage = computed(() => {
  if (config.value == null) {
    return ``
  }
  return `url(${config.value.Icon})`
})
const WebsiteName = computed(() => {
  if (config.value == null) {
    return 'NULL'
  }
  return config.value.WebsiteName
})

const MyName = computed(() => {
  if (config.value == null) {
    return 'NULL'
  }
  return config.value.MyName
})

const isAdmin = computed(() => {
  if (userInfo.user.Username === null) {
    return false
  } else if (!userInfo.user.IsAdmin) {
    return false
  }
  return true
})

const isLogin = computed(() => {
  if (userInfo.user.Username === null) {
    return false
  }
  return true
})

function isThisPage(num: number) {
  return num == pageData.pagedata.count;

}

function jumpTo(path: string) {
  router.push(path)
  if (path === '/home') {
    pageData.pagedata.count = 0
    toggle()
  }
}

function toggle() {
  emit('toggleNav')
}

function logIn() {
  router.push("/login")
  toggle()

}

function logOut() {
  dialog.warning({
    title: '登出确认',
    content: '确定登出吗？',
    positiveText: '确定',
    negativeText: '算了',
    onPositiveClick: () => {
      localStorage.removeItem('token')
      userInfo.user = {
        Name: null,
        Username: null,
        IsAdmin: false
      }
      router.replace(router.currentRoute.value)
      toggle()
    },
    onNegativeClick: () => {
      return
    }
  })

}

function changePage(num: number) {
  if (router.currentRoute.value.path.startsWith("/login")) {
    router.push("/home")
    emit('toggleNav')
  }
  pageData.pagedata.count = num;

}
</script>

<style lang="less" scoped>
.n-icon:hover {
  cursor: not-allowed;
}

a {
  cursor: not-allowed;
  // pointer-events: none;
}

.hover:hover {
  cursor: not-allowed;
}

.settings {
  position: absolute;
  left: 30px;
  bottom: 5px;
  cursor: pointer !important;
}

.login {
  position: absolute;
  right: 5px;
  bottom: 5px;
  cursor: pointer !important;
}

.logout {
  position: absolute;
  right: 5px;
  bottom: 5px;
  cursor: pointer !important;
}

.home {
  position: absolute;
  left: 5px;
  bottom: 5px;
  cursor: pointer !important;
}

.icon {
  width: 1em;


  height: 1em;
  vertical-align: -0.15em;
  fill: currentColor;

  &:hover {
    cursor: not-allowed;
  }

  // overflow: hidden;
}

.leftbar {
  user-select: none;
  border-right: 4px solid rgb(49, 49, 49);
  height: 100vh;
  color: white;
  position: fixed;
  float: left;
  width: 16.5%;
  min-height: 100vh;
  // background: linear-gradient(#ffff, rgb(255, 244, 232));
  background-color: rgb(255, 255, 255);
  display: block;
  box-shadow: 0 0 40px rgba(0, 0, 0, 0.2);
  //backdrop-filter: blur(3px);
  background-size: cover;
  background-position: center;
  // background-image: url("../assets/69167110_p0.png");
  transition: all .3s ease-in-out;
  overflow: hidden;

  @media (max-width: 1024px) {
    //left: -16.5%
    transform: translateX(-300px);
    //display: none;
  }

  .header {
    position: relative;
    color: black;
    top: 20px;
    font-size: 2em;
    //border:3px solid rgb(49,49,49);
    //border-radius: 2px;
    box-shadow: -1px 1px 4px 1px rgba(49, 49, 49, .5);
    width: 12vw;
    background-color: rgb(55, 166, 187);
    color: white;

    @media (max-width: 1024px) {
      width: 200px;
    }
  }

  div {
    margin: 0 auto;
    text-align: center;
  }

  .profile {

    border: 4px rgb(49, 49, 49) solid;
    position: relative;
    margin: 0 auto;
    // display: flex;
    width: 12vw;
    height: 15vw;
    // font-weight: 5000;
    // font-family: Verdana, Geneva, Tahoma, sans-serif;
    // border-radius: 50%;

    background: white;
    text-align: center;
    margin-top: 20%;
    // justify-content: center;
    // align-items: center;
    color: black;
    //background-image: url("../assets/e342ee78a065e6c608d9d31959dd8b24.jpeg");
    // background-image: url("../assets/QQ图片20220723102142.jpg");
    background-size: cover;
    background-position: center;
    background-repeat: no-repeat;

    @media (max-width: 1024px) {
      width: 200px;
      height: 250px;
    }
  }

  .name {
    margin-top: 55px;
    width: 80%;
    color: rgb(78, 78, 78);
    text-align: center;
    // font-weight: 600;
    // font-family: 'Times New Roman', Times, serif;
    font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
    font-size: 1.75em;
    text-justify: inter-word;
    letter-spacing: 2px;
    text-indent: 5px;
    word-wrap: break-word;
    text-shadow: 0px 0px 2px rgba(235, 235, 235, 0.5);

  }

  .icon1 {
    display: flex;
    justify-content: space-around;
    align-items: center;
    width: 50%;
    margin: 20px auto;

    div {
      color: rgb(241, 91, 91);
      font-size: 1em;
      width: 20px;
      height: 20px;
    }
  }

  .footer {
    display: flex;
    margin: 50px 10%;
    min-height: 25vh;
    justify-content: space-around;
    align-items: center;
    flex-direction: column;
    position: absolute;
    width: 80%;
    bottom: 0vh;

    .footerbox {
      cursor: pointer;
      width: 80%;
      background-color: #ffffff;
      height: 30px;
      border-radius: 2px;
      outline-style: none;
      border: 3px solid rgb(51, 51, 51);
      box-shadow: -1px 1px 0 1px rgb(51, 51, 51);

      &:hover {
        //background-color: rgb(255, 83, 83) !important;
        background-color: var(--button-color) !important;

        color: white !important;
      }
    }
  }

}

.active {
  //background-color: rgb(255, 83, 83) !important;
  background-color: var(--button-color) !important;
  color: white !important;;

}
</style>