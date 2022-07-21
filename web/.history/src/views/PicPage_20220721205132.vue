<template>
    <div class="container">
        <div class="panel">
            panel
        </div>
        <div class="panel">
            panel

        </div>
        <div class="panel">
            panel

        </div>
        <div class="panel">
            panel

        </div>
        <div class="panel">
            panel

        </div>
        <div class="panel">
            panel

        </div>
    </div>
</template>

<style lang="less" scoped>
.container {
    width: 80%;
    float: left;
    height: 100vh;
    overflow: auto;
    position: relative;
    bottom: 0;
    margin: auto auto;
}

::-webkit-scrollbar {
    width: 0px;
}

.panel {
    background-color: rgba(109, 155, 255, 0.8);
    width: 80%;
    height: 25vh;
    margin: 30px auto;
    border-radius: 5px;
    text-align: center;
}
</style>