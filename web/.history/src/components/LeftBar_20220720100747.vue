<template>
    <div class="leftbar">
        <div class="profile"></div>
        <div class="name">这里是名字</div>
        <div class="icon">这里是图标</div>
        <div class="footer">这里是底部</div>
    </div>
</template>

<script>
export default {
    name: "LeftBar",
};
</script>

<script lang="ts" setup>

</script>

<style lang="less" scoped>
.leftbar {
    color: white;
    position: absolute;
    width: 16.5%;
    min-height: 100%;
    background-color: rgb(131, 224, 255);
    display: block;

    .profileF {
        flex: .1;
    }

    div {
        margin: 0 auto;
    }

    .profile {
        margin: 0 auto;
        display: flex;
        width: 200px;
        height: 200px;
        border-radius: 50%;
        background: white;
        text-align: center;
        // justify-content: center;
        // align-items: center;
        color: black;
    }

    .name {
        background-color: #fff;
        width: 50%;
    }

}
</style>