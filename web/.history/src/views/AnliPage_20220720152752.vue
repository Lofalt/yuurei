<template>
    <div>这里是安利页</div>
    <!-- {{ pageNum }} -->
</template>

<script lang="ts" setup>

import { useRouter } from 'vue-router'
const router = useRouter()

console.log(router.currentRoute.value.query['pageNum'])
// interface Props {
//     pageNum: number
// }
// defineProps<Props>()

// console.log(pageNum)

</script>