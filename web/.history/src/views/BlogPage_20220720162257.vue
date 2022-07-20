<template>
    <div class="container">
        <div class="btnHeader">
            <button class="button">王四月</button>
            <button class="button">椎名林檎</button>
            <button class="button">唐可可</button>
            <button class="button">BUMP OF CHICKEN</button>
        </div>
        <div>
            <router-view></router-view>
        </div>
    </div>
</template>

<script lang="ts" setup>


</script>

<style>
.container {
    width: 80%;
    height: 100%;
    background-color: #fff;
    /* margin: 0 auto; */
}

.btnHeader {
    /* width: 100%; */
    /* height: 100%; */
    margin: 0 auto;
}

.button {
    margin: 0 auto;
    background-color: blanchedalmond;
    border: none;
    padding: 10px 20px;
    margin: 20px;
}
</style>