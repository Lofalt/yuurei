<template>
    <div class="rightbar" @click.self="pushCard">
        <div class="showpage">
            <div class="showtag tag1" @click="pullCard(0)">
                <!-- 休職 -->
            </div>
            <router-view v-show="showPic"></router-view>
        </div>
        <div class="showpage">
            <div class="showtag tag2" @click="pullCard(1)">
                <!-- 休職 -->
            </div>
            <router-view name="BlogPage" v-show="showBlog"></router-view>
        </div>
        <div class="showpage">
            <div class="showtag tag3" @click="pullCard(2)">
                <!-- 休職 -->
            </div>
            <router-view name="VlogPage" v-show="showVlog"></router-view>
        </div>
        <div class="showpage">
            <div class="showtag tag4" @click="pullCard(3)">
                <!-- 休職 -->
            </div>
            <router-view name="LovePage" v-show="showLove"></router-view>
        </div>
    </div>
</template>

<script lang="ts">
export default {
    name: "RightBar"
}
</script>
<script lang="ts" setup>
import { onDeactivated, onUnmounted, ref } from 'vue';
import { useRouter } from 'vue-router'

const router = useRouter();
const showBlog = ref(false)
const showVlog = ref(false)
const showLove = ref(false)
const showPic = ref(false)


function pullCard(num: number): void {
    if (router.currentRoute.value.fullPath == "/home") {
        router.push("/home/pages")
    }
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
            showVlog.value = true
            break;
        case 3:
            showLove.value = true
            break;
    }
}

function pushCard() {
    const cards = document.getElementsByClassName("showpage") as HTMLCollection

    for (let i = 0; i < cards.length; i++) {
        let card = cards[i] as HTMLElement
        card.style.transitionDelay = "0"
        card.style.right = "-55vw"
    }
    showBlog.value = false
    showLove.value = false
    showPic.value = false
    showVlog.value = false
}


</script>

<style lang="less" scoped>
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
        overflow: hidden;
        position: absolute;
        right: 0;
        top: 10vh;
        box-shadow: -2px 2px 5px 1px rgba(0, 0, 0, 0.1);
        width: 80vw;
        height: 80vh;
        // background-color: #ffffff;
        backdrop-filter: blur(3px);
        transition: all 1s cubic-bezier(0.075, 0.82, 0.165, 1);
        // border: 1px solid rgba(0, 0, 0, 0.3);
        // transform: translateX(70vw);
        border-radius: 5px;
        right: -55vw;


        &:nth-child(4) {
            z-index: 4;
            // background-color: aquamarine;
            // transform: translateX(75vw);
            transform: translateX(20vw);
        }

        &:nth-child(3) {
            z-index: 3;
            // background-color: aquamarine;
            // transform: translateX(70vw);
            transform: translateX(15vw);
        }

        &:nth-child(2) {
            z-index: 2;
            // background-color: aquamarine;
            // transform: translateX(65vw);
            transform: translateX(10vw);
        }

        &:nth-child(1) {
            z-index: 1;
            // background-color: aquamarine;
            // transform: translateX(60vw);
            transform: translateX(5vw);
        }

        // transform: translateX(70vw);

        // &:hover {
        //     right: 0;
        // }

        .showtag {

            position: absolute;
            width: 35px;
            height: 110px;
            background-color: #fff;
            left: 0px;
            top: 25px;
            cursor: pointer;
            box-shadow: 0 0 4px 1px rgba(0, 0, 0, 0.1);
            border-radius: 3px;
            // border: 1px solid rgba(0, 0, 0, 0.3)
            writing-mode: vertical-lr;
            text-align: center;
            color: white;
            font-size: 1.3em;
            padding: 8px 5px;
        }

        .tag1 {
            background-color: #ff5454;
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
            background-color: #16cedb;
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