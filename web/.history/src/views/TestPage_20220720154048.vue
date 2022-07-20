<template>
    <div>{{ hello }}</div>
</template>

<script lang="ts" setup>

const props = defineProps<{
    hello: string
}>()
console.log(props.hello)


</script>

<script lang="ts">
export default {
    name: "TestPage"
}
</script>