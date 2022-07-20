<template>
    <div class="rightbar">
        <div class="showpage">
            <div class="tag"></div>
            <router-view></router-view>
        </div>
        <div class="showpage">
            <div class="tag"></div>
            <router-view></router-view>
        </div>
        <div class="showpage">
            <div class="tag"></div>
            <router-view></router-view>
        </div>
        <div class="showpage">
            <div class="tag"></div>
            <router-view></router-view>
        </div>
    </div>
</template>

<script>
export default {
    name: "RightBar"
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
        box-shadow: 1px 0 10px 1px rgba(0, 0, 0, 0.1);
        width: 80vw;
        height: 80vh;
        background-color: rgb(250, 242, 234);
        transition: all 1s cubic-bezier(0.075, 0.82, 0.165, 1);



        &:nth-child(1) {
            z-index: 4;
            // background-color: aquamarine;
            transform: translateX(75vw);
        }

        &:nth-child(2) {
            z-index: 3;
            // background-color: aquamarine;
            transform: translateX(70vw);
        }

        &:nth-child(3) {
            z-index: 2;
            // background-color: aquamarine;
            transform: translateX(65vw);
        }

        &:nth-child(4) {
            z-index: 1;
            // background-color: aquamarine;
            transform: translateX(60vw);
        }

        &:hover {
            transform: translateX(0);
        }
    }

    .tag {

        position: absolute;
        width: 50px;
        height: 100px;
        background-color: #fff;
        left: -25px;
        top: 25px;

        // &:nth-of-type(1) {
        // top: 15vh;
        // }

        &:nth-child(2) {
            top: 25vh;
        }

    }
}
</style>