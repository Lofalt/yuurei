<template>
    <div class="leftbar">
        <div class="icon">这里是头像</div>
        <span>这里是名字</span>
        <span>这里是图标</span>
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

<style scoped>
.leftbar {
    color: white;
    position: absolute;
    width: 16.5%;
    min-height: 100%;
    background-color: rgb(131, 224, 255);
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}
</style>