<template>
    <div>这里是安利页</div>
    {{ pageNum }}
</template>

<script lang="ts" setup>

interface Props {
    pageNum: number
}
defineProps<Props>()


</script>