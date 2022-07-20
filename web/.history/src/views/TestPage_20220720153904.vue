<template>
    <div>{{ name }}</div>
</template>

<script lang="ts" setup>

defineProps<{
    name: string
}>()
console.log(name.value)


</script>

<script lang="ts">
export default {
    name: "TestPage"
}
</script>