<template>
    <div class="messageBox">
        <div class="icon"></div>
        <div class="messageInfo">
            <div></div>
        </div>
    </div>

</template>

<style lang="less" scoped>
.messageBox {
    // float: left;
    width: 90%;
    margin: 5vh auto;

    // height: 100%;
    &::after,
    &::before {
        content: '';
        clear: both;
        overflow: hidden;
    }

    .icon {
        float: left;
        width: 5vw;
        height: 5vw;
        background-color: rgb(49, 49, 49);
        margin-right: 2vw;
    }

    .messageInfo {
        float: left;
        width: 80%;
        height: 10vw;
        background-color: aquamarine;
    }
}
</style>