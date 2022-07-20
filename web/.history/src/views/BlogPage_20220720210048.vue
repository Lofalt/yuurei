<template>
    <div class="container">
        <div class="btnHeader">
            <button class="button" @click="change">王四月</button>
            <button class="button">椎名林檎</button>
            <button class="button">唐可可</button>
            <button class="button">BUMP OF CHICKEN</button>
        </div>
        <div>
            <keep-alive>
                <router-view class="info"></router-view>
            </keep-alive>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { useRouter } from 'vue-router'

const router = useRouter()

function change() {
    router.push("/home/pages/anli/1")
}
</script>

<style lang="less">
.container {
    width: 80%;
    height: 100%;
    /* background-color: #fff; */
    /* margin: 0 auto; */
    overflow: auto;
}

.btnHeader {
    /* width: 100%; */
    /* height: 100%; */
    width: 80%;
    margin: 8vh auto;
    text-align: center;
    border-bottom: 1px solid rgba(0, 0, 0, 0.2);
    padding-bottom: 5vh;
}

.button {
    margin: 0 auto;
    background-color: blanchedalmond;
    border: none;
    padding: 10px 20px;
    margin: 20px;
    border-radius: 10px;
    color: rgb(89, 88, 88);
    cursor: pointer;

    &:hover {
        background-color: rgb(252, 125, 125);
    }
}

.info {
    width: 100%;
    height: 100vh;
    background-color: aliceblue;
}
</style>