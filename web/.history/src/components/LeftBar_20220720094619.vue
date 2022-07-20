<template>
    <div class="leftbar">
        <span>这里是侧边栏</span>
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
    width: 20%;
    min-height: 100%;
    background-color: rgba(0, 0, 0, 1);
}
</style>