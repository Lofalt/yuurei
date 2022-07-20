<template>
    <div class="leftbar">
        <div class="profile"></div>
        <div class="name">LOFALT</div>
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
    box-shadow: 0 0 40px rgba(0, 0, 0, 0.1);

    .profileF {
        flex: .1;
    }

    div {
        margin: 0 auto;
    }

    .profile {
        position: relative;
        margin: 0 auto;
        // display: flex;
        width: 12vw;
        height: 12vw;
        border-radius: 50%;
        background: white;
        text-align: center;
        margin-top: 30%;
        // justify-content: center;
        // align-items: center;
        color: black;
    }

    .name {
        margin-top: 50px;
        // background-color: #fff;
        width: 50%;
        color: rgb(59, 53, 53);
        text-align: center;
        font-size: 2em;
        // text-justify: inter-word;
        text-indent: 10px;
        letter-spacing: 10px;
    }

}
</style>