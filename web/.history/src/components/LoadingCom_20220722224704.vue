<template>
    <div class="father" style="--i:1">
        <div class="circle" style="--i:1"></div>
        <div class="circle" style="--i:2"></div>
        <div class="circle" style="--i:3"></div>
        <div class="circle" style="--i:4"></div>
        <div class="circle" style="--i:5"></div>
        <div class="circle" style="--i:6"></div>
        <div class="circle" style="--i:7"></div>
        <div class="circle" style="--i:8"></div>
        <div class="circle" style="--i:9"></div>
        <div class="circle" style="--i:10"></div>
        <div class="circle" style="--i:11"></div>
        <div class="circle" style="--i:12"></div>
        <div class="circle" style="--i:13"></div>
        <div class="circle" style="--i:14"></div>
        <div class="circle" style="--i:15"></div>
        <div class="circle" style="--i:16"></div>
        <div class="circle" style="--i:17"></div>
        <div class="circle" style="--i:18"></div>
        <div class="circle" style="--i:19"></div>
        <div class="circle" style="--i:20"></div>
    </div>

</template>

<script>
export default {
    name: "LoadingCom"
}

</script>

<style>
.father {
    width: 25px;
    height: 25px;
    /* position: relative; */

}

@keyframes spin {
    to {
        transform: rotate(1turn);
    }
}

.circle {
    position: absolute;
    left: 50%;
    width: 2px;
    height: 5px;
    border-radius: 5px;
    transform: rotate(calc(18deg*var(--i)));
    background-color: #fcaf68;
    animation: move 1.2s backwards infinite ease-in-out alternate-reverse;
    animation-delay: calc(Var(--i)*0.06s);
    animation-fill-mode: backwards;
    transform-origin: 0 25px;
    box-shadow: 0px 0px 15px 2px rgba(255, 255, 255, 0.6);
    filter: blur(10);
}

@keyframes move {
    to {
        background-color: #ffffff;
        transform: scale(0);
    }
}
</style>
