<template>

  <div class="app">
    <div :class="[isActive?`activeClass`:`unActive`,`mask`]" @click="showNav"></div>
    <div :class="[isActive?`lefting`:``,`navButton`]" @click="showNav">
      <div class="icons"></div>
      <div class="icons"></div>
      <div class="icons"></div>
    </div>
    <transition>
      <n-icon size="40" color="#ffffff" @click="changePage(0)" class="fixedButton" v-show="pageData.pagedata.count!=0">
        <arrow-bar-to-up/>
      </n-icon>
    </transition>
    <n-dialog-provider>
      <left-bar :class="[isActive?`showing`:`notShowing`,`leftBar`]"/>
    </n-dialog-provider>
    <router-view></router-view>
  </div>
</template>
<script lang="ts" setup>
import {computed, getCurrentInstance, onMounted, provide, ref, watch} from 'vue';
import LeftBar from './components/LeftBar.vue';
import {usePageData} from './store/pageData';
import Date from "./components/Date.vue";
import {ArrowBarToUp} from "@vicons/tabler"
import {NIcon} from "naive-ui"
import {useUserInfo} from "@/store/UserInfo";
import axios from "@/request/index"
import {useRouter} from "vue-router";
import {useMessage, NDialogProvider} from "naive-ui";

const router = useRouter()
const isShow = ref(true)
const pageData = usePageData()
let isWheeling = false
const isAdmin = ref(true)
const userInfo = useUserInfo()
const isActive = ref(false)

router.beforeEach((to, from, next) => {
  if (to.meta.needAuth && !userInfo.user.IsAdmin && to.name != "login") {

    // alert("你没有权限")
    // next({
    //   name: "home"
    // })
    // return false
  }

  if (userInfo.user.Username === null) {
    axios.get("/yuurei/info", {}).then((res: any) => {
      if (res.code === 200) {
        userInfo.user = res.data.user
      }
    })
  }

  next()
})

function showNav() {
  // var childs = document.getElementsByClassName("leftBar") as HTMLSelectElement
  // childs[0].style.transform = `translateX(0)`
  isActive.value = !isActive.value
}

onMounted(() => {
  let user = userInfo.user
  if (user.Username === null) {
    axios.get("/yuurei/info", {}).then((res: any) => {
      if (res.code === 200) {
        userInfo.user = res.data.user
      }
    })
  }
})


function changePage(num: number) {
  pageData.pagedata.count = num

}

</script>

<style lang="less" scoped>
.lefting {
  transform: translateX(300px);
}

.showing {
  transform: translateX(0);
}

.activeClass {
  background-color: rgba(0, 0, 0, .3);
}


.unActive {
  display: none;
}


.leftBar {
  z-index: 5;
  @media (max-width: 1024px) {
    width: 300px;
  }
}

.mask {
  //transform: translateX(300px);
  //background-color: black;
  z-index: 1;
  width: 100vw;
  height: 100vh;
  position: absolute;
}

.fixedButton {
  position: fixed;
  top: 5vh;
  left: 18vw;
  z-index: 5;
  cursor: pointer;
  @media (max-width: 1024px) {
    height: 10vh;
    left: 1vw;
    // padding: 10px;
    //padding: 0;
    //width: 8vw;
    top: 1vh;
    //writing-mode: vertical-lr;
    //background-color: rgb(49, 49, 49);
    //color: white;
  }

  &:hover {
    //background-color: rgb(255, 83, 83) !important;
    //color: white !important;

  }

}

#rightbar {
  position: relative;
  width: 83.5%;
  height: 100vh;
  float: right;
  // overflow: hidden;
  transition: all .8s ease-in-out;

  @media (max-width: 1024px) {
    width: 100%;
  }
}

.middle {
  float: right;
  position: relative;
  overflow: hidden;
  width: 100%;
  min-height: 100vh;
}

.bottom,
.info {
  position: relative;
  float: right;
  width: 100%;
  height: 100vh;
  transition: all 2s;
}

iframe {
  width: 100%;
  height: 100%;
}

.bottomPage,
.infoPage {
  overflow: auto;
  // background-image: url("./assets/wallhaven-p262de.jpg");
  border: 5px solid rgb(49, 49, 49);
  width: 80%;
  height: 80%;
  position: relative;
  background-size: cover;
  margin: 10vh auto;
  border-radius: 5px;
  background-color: rgb(255, 255, 255);
  // padding: 100px;
  display: flex;
  justify-content: center;
  align-items: center;

  &::-webkit-scrollbar-thumb {
    width: 30px;
    background-color: rgba(73, 73, 73, .5);
    color: black;
    border-radius: 20px;
  }

}

.infoPage {
  // background-image: url("./assets/wallhaven-l3w9kr.jpg");
}

button {
  position: absolute;
  left: 10px;
  z-index: 10;

  // &:nth-child(2) {
  //   left: 20px;
  // }
}

#app {
  min-height: 100vh;
  // overflow: hidden;
}

.app {
  background-size: cover;
  background-repeat: repeat-y;
  // background-image: url("./assets/94029233_p0.jpg");
  position: relative;
  min-height: 100vh;
  overflow: hidden;
}

.middle {
  animation: move 1s ease;
}

/* .top {
  position: absolute;
  right: 1000px;
  animation: move 1s ease;
} */

@keyframes move {
  from {
    transform: translateY(1000px);
  }

  /* to {
    transform: translateX(-1000px);
  } */
}

.v-leave-to,
.v-enter-from {
  transform: translateY(-100vh);
}

.v-leave-active,
.v-enter-active {
  transition: all .8s ease-in-out;
}

.navButton {
  position: absolute;
  //transition: .5s all ease;
  transition: all .3s ease-in-out;

  background-color: white;
  padding: 1px;
  border-radius: 2px;
  z-index: 10;
  left: 0;
  top: 50px;

  &:hover {
    cursor: pointer;
  }

  @media screen and (min-width: 1024px) {
    display: none;
  }

  .icons {
    height: 2px;
    min-width: 20px;
    margin: 5px 8px;
    background-color: #3d3535;
    border-radius: 2px;

    &:nth-of-type(1) {
      margin-top: 8px;
    }

    &:nth-of-type(3) {
      margin-bottom: 8px;
    }
  }

}


</style>