<template>
    <div class="rightbar" @click.self="pushCard">
        <div class="showpage activeCard" @wheel.stop="">
            <div class="showtag tag1" @click="pullCard(0)">
                人間話
            </div>
            <!-- <div class="showtag tag1" @click="pullCard(0)"> -->
            <!-- 鲤鱼姐 -->
            <!-- </div> -->
            <pic-page></pic-page>
            <!-- <router-view v-show="showPic"></router-view> -->
        </div>
        <div class="showpage activeCard" @wheel.stop="">
            <div class="showtag tag2" @click="pullCard(1)">
                碎碎念
            </div>
            <blog-page></blog-page>
            <!-- <router-view name="BlogPage" v-show="showBlog"></router-view> -->
        </div>
        <div class="showpage activeCard" @wheel.stop="">
            <div class="showtag tag3" @click="pullCard(2)">
                好物
            </div>
            <!-- <router-view name="LovePage" v-show="showLove"></router-view> -->
            <love-page></love-page>
        </div>
        <div class="showpage activeCard" @wheel.stop="">
            <div class="showtag tag4" @click="pullCard(3)">
                人间
            </div>
            <!-- <router-view name="VlogPage" v-show="showVlog"></router-view> -->
            <!-- <div class="father"> -->
            <anli-page></anli-page>
            <!-- </div> -->
            <!-- <water-page></water-page> -->
        </div>
    </div>
</template>

<script lang="ts">
export default {
    name: "RightBar"
}
</script>
<script lang="ts" setup>
import WaterPage from "../views/WaterPage.vue"
import { onDeactivated, onUnmounted, provide, ref } from 'vue';
import { useRouter } from 'vue-router'
import AnliPage from '../../../src/views/PicPage.vue';
import PicPage from "../../../src/views/ArticlePage.vue";
import VlogPage from "@/views/VlogPage.vue";
import LovePage from "@/views/LovePage.vue";
import BlogPage from "../../../src/views/MessageBoard.vue";
// import LovePage from "@/views/LovePage.vue";
const router = useRouter();
const showBlog = ref(false)
const showVlog = ref(false)
const showLove = ref(false)
const showPic = ref(false)
const hover = ref(false)
provide('hover', hover)
provide('showVlog', showVlog)

function pullCard(num: number): void {
    // if (router.currentRoute.value.fullPath == "/home") {
    //     router.push("/home/pages")
    // }
    const cards = document.getElementsByClassName("showpage") as HTMLCollection
    for (let i = num; i >= 0; i--) {
        let card = cards[i] as HTMLElement
        card.style.transitionDelay = (i - 4) * -0.02 + "s"
        // card.style.right = "0"
        card.className = "showpage"
    }
    for (let i = num; i < cards.length; i++) {
        let card = cards[i] as HTMLElement
        card.style.transitionDelay = (i - 4) * -0.02 + "s"
        // card.style.right = "-65vw"
        // card.style.right = "0"
        card.className = "showpage activeCard"
    }

    let element = cards[num] as HTMLElement
    element.className = "showpage"

    switch (num) {
        case 0:
            showPic.value = true;
            break;
        case 1:
            showBlog.value = true
            break;
        case 2:
            showLove.value = true
            break;
        case 3:
            showVlog.value = true
            break;
    }
}

function pushCard() {
    hover.value = false
    const cards = document.getElementsByClassName("showpage") as HTMLCollection

    for (let i = 0; i < cards.length; i++) {
        let card = cards[i] as HTMLElement
        card.style.transitionDelay = "0"
        card.style.right = "0"
        card.className = "showpage activeCard"
    }

    setTimeout(() => {

        showBlog.value = false
        showLove.value = false
        showPic.value = false
        showVlog.value = false
    }, 1000);
}


</script>

<style lang="less" scoped>
.father {
    overflow: auto;
    width: 80%;
    height: 100vh;
}

.rightbar {
    float: right;
    position: relative;
    overflow: hidden;
    width: 100%;
    min-height: 100vh;

    // min-height: 100vh;
    // width: 83.5vw;
    // overflow: hidden;
    .activeCard {
        right: -70vw !important;

        @media (max-width:800px) {
            right: -90vw !important;
        }
    }

    .showpage {
        // background-attachment: fixed;
        // overflow: ;
        position: absolute;
        right: 0;
        top: 10vh;
        box-shadow: -2px 2px 5px 1px rgba(0, 0, 0, 0.1);
        width: 70vw;
        height: 80vh;
        background-color: #fff;
        // backdrop-filter: blur(3px);
        transition: all .5s ease-in-out; // transition: all 1.5s cubic-bezier(0.19, 1, 0.22, 1);
        border: 4px solid rgb(49, 49, 49);
        border-right: none;
        border-radius: 5px;
        // right: -65vw;
        background-size: cover;

        @media (max-width:800px) {
            width: 90vw;
            // right: -85vw;
        }

        &:nth-child(4) {
            z-index: 4;
            // width: 60vw;
            // transform: translateX(20vw);
            // overflow: auto;
            // background-image: url(../assets/91375949_p0.jpg);
            // padding-left: 10vw;
            // padding-top: 5vh;
            display: flex;
            justify-content: space-around;
            flex-direction: column;
        }

        &:nth-child(3) {
            z-index: 3;
            // transform: translateX(15vw);
            // transform: translateX(-5vw);
            // width: 60vw;
            // background-image: url(../assets/89025157_p0.jpg);
        }

        &:nth-child(2) {
            z-index: 2;
            // transform: translateX(10vw);
            // transform: translateX(-10vw);
            // width: 60vw;
            // background-image: url(../assets/wallhaven-k71327.jpg);
        }

        &:nth-child(1) {
            z-index: 1;
            // transform: translateX(-15vw);
            // width: 60vw;
            // background-image: url(../assets/ec8ce499ly1h1hhm1o6m1j22pg1ww4qp.jpg);
        }

        .showtag {
            // box-sizing: content-box;
            display: flex;
            justify-content: flex-start;
            align-items: center;
            // border: 4px solid rgba(73, 73, 73, 1);
            // box-sizing: content-box;
            z-index: 10;
            position: absolute;
            width: 35px;
            height: 10vhpx;
            background-color: #fff;
            left: -35px;
            top: 5vh;
            cursor: pointer;
            box-shadow: -3px 0px 0px 2px rgb(54, 54, 54);
            border-radius: 3px;
            writing-mode: vertical-lr;
            text-align: 1px 1 px;
            color: rgb(87, 87, 87);
            font-size: 1.3em;
            padding: 8px 5px;
            border-top: 3px solid rgb(73, 73, 73);
            border-bottom: 3px solid rgb(73, 73, 73);


            &:hover {
                background-color: rgb(255, 83, 83);
                color: white;
            }
        }

        .tag1 {
            // background-color: #9dd498;

            &:nth-child(2) {
                // box-sizing: content-box;
                border: 4px solid rgb(47, 47, 47);
                left: -10px;
                top: 20px;
                width: 0.1px;
                z-index: -4;
                height: 120px;
                transform-origin: center;
                transform: translateZ(-5px);
                padding: 0;
                // position: absolute;
                // position: fixed;
            }
        }

        .tag2 {
            background-color: rgb(255, 83, 83);
            // background: linear-gradient(to right, rgb(255, 83, 83), #ffffff);
            top: 23vh;
            color: white;
        }

        .tag3 {
            // background-color: #f18888;
            top: 41vh
        }

        .tag4 {
            // background-color: #78d2d8;
            top: 59vh;
        }
    }

}

@keyframes move {
    to {
        right: 0;
    }
}
</style>