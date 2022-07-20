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
        width: 70vw;
        height: 70vh;
        background-color: #fff;
    }
}
</style>