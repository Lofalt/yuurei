<template>
    <div>这里是安利页</div>
    {{ props.pageNum }}
</template>
<script lang="ts" setup>

const props = defineProps<{
    pageNum: string
}>()
</script>
