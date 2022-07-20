<template>
    <div class="leftbar">
        <div class="profile">这里是头像</div>
        <span class="name">这里是名字</span>
        <span class="icon">这里是图标</span>
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
    display: flex;
    flex-direction: column;
    justify-content: space-around;
    align-items: center;

    .profile {
        display: flex;
        width: 200px;
        height: 200px;
        border-radius: 50%;
        background: white;
        text-align: center;
        justify-content: center;
        align-items: center;
        color: black;
        // flex: 1
    }
}
</style>