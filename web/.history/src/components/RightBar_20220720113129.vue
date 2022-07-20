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

    .showpage {
        width: 80vw;
        height: 80vh;
        background-color: #fff;
    }
}
</style>