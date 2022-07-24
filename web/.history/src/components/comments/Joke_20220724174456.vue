<template>
    <div class="messageBox">
        <div class="icon"></div>
        <div class="messageInfo">
            <span></span>
            <div class="info">鲤鱼姐没有你我怎么活啊</div>
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
        content: ""; //设置内容为空

        height: 0; //高度为0

        line-height: 0; //行高为0

        display: block; //将文本转为块级元素

        visibility: hidden; //将元素隐藏

        clear: both //清除浮动
    }

    .icon {
        background-image: url("../../assets/91375949_p0.jpg");
        background-size: cover;
        border: 3px solid rgb(49, 49, 49);
        float: left;
        width: 5vw;
        height: 5vw;
        background-color: rgb(49, 49, 49);
        margin-right: 2vw;
    }

    .messageInfo {
        border: 4px solid rgb(49, 49, 49);
        padding: 1vw;
        float: left;
        width: 80%;
        border-radius: 3px;
        height: 15vh;
        // background-color: aquamarine;
        position: relative;

        @media (max-width:1024px) {
            height: 10vh;
        }

        &::after {
            position: absolute;
            display: block;
            left: -12px;
            // top: 5px;
            top: 1px;
            transform: rotate(225deg);
            content: '';
            width: 0;
            height: 0;
            border: 8px solid transparent;
            border-top-color: black;
            border-right-color: black;
        }

        &::before {
            position: absolute;
            display: block;
            left: -6px;
            // top: 5px;
            z-index: 10;
            top: 1px;
            transform: rotate(225deg);
            content: '';
            width: 0;
            height: 0;
            border: 8px solid transparent;
            border-top-color: white;
            border-right-color: white;
        }

        .info {}
    }
}
</style>