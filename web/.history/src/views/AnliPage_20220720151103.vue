<template>
    <div>这里是安利页</div>
    {{ pageNum }}
</template>

<script lang="ts" setup>

// import { defineProps } from 'vue';
interface Props {
    pageNum: number
}

defineProps<Props>()

// console.log(pageNum)
</script>