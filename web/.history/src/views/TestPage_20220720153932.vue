<template>
    <div>{{ hello }}</div>
</template>

<script lang="ts" setup>

defineProps<{
    hello: string
}>()
console.log(name)


</script>

<script lang="ts">
export default {
    name: "TestPage"
}
</script>