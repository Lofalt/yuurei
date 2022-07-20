<template>
    <div>这里是安利页</div>
    {{ pageNum }}
</template>
<script lang="ts" setup>

import { ref } from 'vue';
// import { useRouter } from 'vue-router'
// const router = useRouter()
// const pageNum = ref(router.currentRoute.value.query['pageNum'])
// interface Props {
    //     pageNum: number
// }
// defineProps<Props>()

// console.log(pageNum)

</script>
<scrip lang="ts">
export default {
props: ["pageNum"]
}
< /script>