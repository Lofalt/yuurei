<template>
    <div class="anliPage">这里是安利页 {{ pageNum }}</div>

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
    background-color: rgb(255, 226, 198);
    border-radius: 20px;
}
</style>
