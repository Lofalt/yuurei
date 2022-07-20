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
    float: left;
    position: relative;
    min-height: 100vh;
    min-width: 83.5vw;

    .showpage {
        position: absolute;
        right: 0;
        top: 10vh;
        box-shadow: 0 0 10px 10px rgba(0, 0, 0, 0.5);
        width: 70vw;
        height: 70vh;
        background-color: #fff;
    }
}
</style>