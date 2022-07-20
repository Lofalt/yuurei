<template>
    <div>这里是安利页</div>
    {{ pageNum }}
</template>

<script lang="ts" setup>

interface Props {
    pageNum: number
}
const props = defineProps<Props>()

console.log(props.pageNum)
</script>