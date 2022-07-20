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
        box-shadow: 0 0 10px 1px rgba(0, 0, 0, 0.1);
        width: 70vw;
        height: 70vh;
        background-color: #fff;

        &:nth-child(1) {
            z-index: 4;
            background-color: aquamarine;
            transform: translateX(70vw);
        }
    }


}
</style>