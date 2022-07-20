<template>
    <div>这里是安利页</div>
    <!-- {{ pageNum }} -->
</template>

<script lang="ts" setup>

import { defineProps } from 'vue';

const props = defineProps({
    pageNum: {
        type: String,
        default: ""
    }
})

console.log(props.pageNum)
</script>