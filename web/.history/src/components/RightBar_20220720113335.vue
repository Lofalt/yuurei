<template>
    <div class="rightbar">这里是右边
        <div class="showpage">
            <router-view></router-view>
        </div>
        <div class="showpage">
            <router-view></router-view>
        </div>
        <div class="showpage">
            <router-view></router-view>
        </div>
        <div class="showpage">
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
    min-width: 80vw;

    .showpage {
        position: absolute;
        width: 70vw;
        height: 70vh;
        background-color: #fff;
    }
}
</style>