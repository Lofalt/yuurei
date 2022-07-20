<template>
    <div class="rightbar" @click.self="pushCard">
        <div class="showpage">
            <div class="showtag tag1" @click="pullCard(0)"></div>
            <router-view></router-view>
        </div>
        <div class="showpage">
            <div class="showtag tag2" @click="pullCard(1)"></div>
            <router-view></router-view>
        </div>
        <div class="showpage">
            <div class="showtag tag3" @click="pullCard(2)"></div>
            <router-view></router-view>
        </div>
        <div class="showpage">
            <div class="showtag tag4" @click="pullCard(3)"></div>
            <router-view></router-view>
        </div>
    </div>
</template>

<script lang="ts">
export default {
    name: "RightBar"
}
</script>
<script lang="ts" setup>


function pullCard(num: number): void {
    const cards = document.getElementsByClassName("showpage") as HTMLCollection
    for (let i = num; i >= 0; i--) {
        let card = cards[i] as HTMLElement
        card.style.right = "0"
    }
    for (let i = num; i < cards.length; i++) {
        let card = cards[i] as HTMLElement
        card.style.right = "-55vw"
    }

    let element = cards[num] as HTMLElement
    element.style.right = "0"
}

function pushCard() {
    const cards = document.getElementsByClassName("showpage") as HTMLCollection

    for (let i = 0; i < cards.length; i++) {
        let card = cards[i] as HTMLElement
        card.style.right = "-55vw"
    }
}
</script>

<style lang="less" scoped>
.rightbar {
    float: right;
    position: relative;
    overflow: hidden;
    width: 83.5%;
    min-height: 100vh;
    // min-height: 100vh;
    // width: 83.5vw;
    // overflow: hidden;

    .showpage {
        position: absolute;
        right: 0;
        top: 10vh;
        box-shadow: -2px 2px 5px 1px rgba(0, 0, 0, 0.1);
        width: 80vw;
        height: 80vh;
        background-color: #f8f7ec;
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
            width: 50px;
            height: 100px;
            background-color: #fff;
            left: -25px;
            top: 25px;
            cursor: pointer;
            box-shadow: 0 0 4px 1px rgba(0, 0, 0, 0.1);
            border-radius: 3px;
            // border: 1px solid rgba(0, 0, 0, 0.3)
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
</style>