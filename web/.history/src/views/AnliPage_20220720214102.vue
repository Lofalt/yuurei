<template>
    <div class="anliPage">{{ pageNum }}</div>

</template>
<script lang="ts" setup>

const props = defineProps<{
    pageNum: string
}>()
</script>

<style lang="less">
.anliPage {
    margin: 0 auto;
    text-align: center;
    width: 80%;
    background-color: rgb(255, 247, 240);
    border-radius: 20px;
}
</style>
