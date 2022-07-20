<template>
    <div>
        <router-link to="/home/pages/anli?pageNum=1" :pageNum=0>click me</router-link>
        <router-link to="/home/pages/anli?pageNum=2" :pageNum=0>click me</router-link>
        <router-view></router-view>
    </div>
</template>

<script lang="ts" setup>
import TestPage from './TestPage.vue';
</script>