<template>
    <div>这里是安利页</div>
    {{ pageNum }}
</template>
<!-- <script lang="ts" setup>

const props = defineProps<{
    pageNum: string
}>()
</script> -->
<script >
export default {
    name: "AnliPage",
    props: ["pageNum"]
}
</script>