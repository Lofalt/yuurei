<template>
    <div class="rightbar" @click.self="pushCard">
        <div class="showpage" @wheel.stop="">
            <div class="showtag tag1" @click="pullCard(0)">
                鲤鱼姐
            </div>
            <div class="showtag tag1" @click="pullCard(0)">
                鲤鱼姐
            </div>
            <pic-page></pic-page>
            <!-- <router-view v-show="showPic"></router-view> -->
        </div>
        <div class="showpage" @wheel.stop="">
            <div class="showtag tag2" @click="pullCard(1)">
                没有你
            </div>
            <blog-page></blog-page>
            <!-- <router-view name="BlogPage" v-show="showBlog"></router-view> -->
        </div>
        <div class="showpage" @wheel.stop="">
            <div class="showtag tag3" @click="pullCard(2)">
                我怎么
            </div>
            <!-- <router-view name="LovePage" v-show="showLove"></router-view> -->
            <love-page></love-page>
        </div>
        <div class="showpage" @wheel.stop="">
            <div class="showtag tag4" @click="pullCard(3)">
                活啊
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
import AnliPage from '../../../src/views/RightBar/PicPage.vue';
import PicPage from "../../../src/views/RightBar/ArticlePage.vue";
import VlogPage from "@/views/VlogPage.vue";
import LovePage from "../../../src/views/RightBar/LovePage.vue";
import BlogPage from "../../../src/views/RightBar/MessageBoard.vue";
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
        card.style.right = "0"
    }
    for (let i = num; i < cards.length; i++) {
        let card = cards[i] as HTMLElement
        card.style.transitionDelay = (i - 4) * -0.02 + "s"
        card.style.right = "-55vw"
    }

    let element = cards[num] as HTMLElement
    element.style.right = "0"

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
        card.style.right = "-55vw"
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

    .showpage {
        // background-attachment: fixed;
        // overflow: ;
        position: absolute;
        right: 0;
        top: 10vh;
        box-shadow: -2px 2px 5px 1px rgba(0, 0, 0, 0.1);
        width: 60vw;
        height: 80vh;
        background-color: #fff;
        // backdrop-filter: blur(3px);
        transition: all 1s cubic-bezier(0.075, 0.82, 0.165, 1);
        border: 4px solid rgb(73, 73, 73);
        border-right: none;
        border-radius: 15px;
        right: -55vw;
        background-size: cover;

        &:nth-child(4) {
            z-index: 4;
            width: 60vw;
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
            transform: translateX(-5vw);
            width: 60vw;
            // background-image: url(../assets/89025157_p0.jpg);
        }

        &:nth-child(2) {
            z-index: 2;
            // transform: translateX(10vw);
            transform: translateX(-10vw);
            width: 60vw;
            // background-image: url(../assets/wallhaven-k71327.jpg);
        }

        &:nth-child(1) {
            z-index: 1;
            transform: translateX(-15vw);
            width: 60vw;
            // background-image: url(../assets/ec8ce499ly1h1hhm1o6m1j22pg1ww4qp.jpg);
        }

        .showtag {
            // border: 2px solid rgba(0, 0, 0, 1);
            // box-sizing: content-box;
            z-index: 10;
            position: absolute;
            width: 35px;
            height: 110px;
            background-color: #fff;
            left: -5px;
            top: 25px;
            cursor: pointer;
            box-shadow: 0 0 4px 1px rgba(0, 0, 0, 0.1);
            border-radius: 3px;
            writing-mode: vertical-lr;
            text-align: start;
            color: white;
            font-size: 1.3em;
            padding: 8px 5px;
        }

        .tag1 {
            background-color: #9dd498;

            &:nth-child(2) {
                border: 10px solid rgba(73, 73, 73);
                left: -10px;
                z-index: -4;
                transform: translateZ(-5px) scaleY(1.05);
                // position: absolute;
                position: fixed;
            }
        }

        .tag2 {
            background-color: #f1c388;
            top: 20vh
        }

        .tag3 {
            background-color: #f18888;
            top: 40vh
        }

        .tag4 {
            background-color: #78d2d8;
            top: 60vh;
        }
    }

}

@keyframes move {
    to {
        right: 0;
    }
}
</style>