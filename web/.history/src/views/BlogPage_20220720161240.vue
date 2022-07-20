<template>
    <div class="btnHeader">
        <button class="button">王四月</button>
        <button class="button">椎名林檎</button>
        <button class="button">唐可可</button>
        <button class="button">BUMP OF CHICKEN</button>
    </div>
    <div>
        <router-view></router-view>
    </div>
</template>

<script lang="ts" setup>


</script>

<style>
.btnHeader {
    /* padding: 10px; */
    /* display: flex; */
    margin: 0 auto;
}

.button {
    margin: 0 auto;
}
</style>