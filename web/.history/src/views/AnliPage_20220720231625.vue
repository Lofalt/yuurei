<template>
    <div class="anliPage">
        <!-- {{ pageNum }} -->
    </div>

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
    background-color: rgba(255, 255, 255, 0.644);
    border-radius: 5px;
    box-shadow: 0 0 5px 2px rgba(0, 0, 0, 0.2);
    overflow: auto;
}
</style>
