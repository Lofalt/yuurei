<template>
    <div>这里是安利页</div>
    {{ pageNum }}
</template>

<script lang="ts" setup>

// import { defineProps } from 'vue';


defineProps<{
    pageNum: string,
}>()

// console.log(pageNum)
</script>