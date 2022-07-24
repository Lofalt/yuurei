<template>
    <div class="messageBox">
        <div class="icon">
            <div class="name">Lofalt</div>
        </div>
        <div class="messageInfo">
            <div class="date">2022.7.24 21:16</div>
            <span></span>
            <div class="info">
                鲤鱼姐没有你我怎么活啊鲤鱼姐没有你我怎么活啊鲤鱼姐没有你我怎么活啊鲤鱼姐没有你我怎么活啊鲤鱼姐没有你我怎么活啊鲤鱼姐没有你我怎么活啊鲤鱼姐没有你我怎么活啊鲤鱼姐没有你我怎么活啊鲤鱼姐没有你我怎么活啊鲤鱼姐没有你我怎么活啊鲤鱼姐没有你我怎么活啊鲤鱼姐没有你我怎么活啊鲤鱼姐没有你我怎么活啊鲤鱼姐没有你我怎么活啊鲤鱼姐没有你我怎么活啊鲤鱼姐没有你我怎么活啊鲤鱼姐没有你我怎么活啊鲤鱼姐没有你我怎么活啊
            </div>
            <div class="info">鲤鱼姐没有你我怎么活啊</div>
            <div class="info">鲤鱼姐没有你我怎么活啊</div>
            <div class="info">鲤鱼姐没有你我怎么活啊</div>
            <div class="info">鲤鱼姐没有你我怎么活啊</div>
            <div class="info">鲤鱼姐没有你我怎么活啊</div>
            <div class="info">鲤鱼姐没有你我怎么活啊</div>
            <div class="info">鲤鱼姐没有你我怎么活啊</div>
            <div class="info">鲤鱼姐没有你我怎么活啊</div>
            <div class="info">鲤鱼姐没有你我怎么活啊</div>
            <div class="info">鲤鱼姐没有你我怎么活啊</div>
            <div class="info">鲤鱼姐没有你我怎么活啊</div>
            <div class="info">鲤鱼姐没有你我怎么活啊</div>
            <div class="info">鲤鱼姐没有你我怎么活啊</div>
            <div class="info">鲤鱼姐没有你我怎么活啊</div>
            <div class="info">鲤鱼姐没有你我怎么活啊</div>
            <div class="info">鲤鱼姐没有你我怎么活啊</div>
        </div>
    </div>

</template>

<style lang="less" scoped>
.messageBox {
    // float: left;
    width: 90%;
    margin: 5vh auto;
    position: relative;

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
        background-image: url("../../assets/ec8ce499ly1h1hhm1o6m1j22pg1ww4qp.jpg");
        background-size: cover;
        border: 3px solid rgb(49, 49, 49);
        float: left;
        width: 6vw;
        height: 6vw;
        background-color: rgb(49, 49, 49);
        margin-right: 20px;
        position: relative;

        @media (max-width:1024px) {
            width: 20vw;
            height: 20vw;
        }

        .name {
            position: absolute;
            bottom: -30px;
            width: 100%;
            text-align: center;
            text-shadow: -1px 1px 2px black;

        }
    }

    .messageInfo {
        border: 4px solid rgb(49, 49, 49);
        border-left: 6px solid rgb(49, 49, 49);
        border-bottom: 6px solid rgb(49, 49, 49);
        padding: 1vw;
        float: left;
        width: 80%;
        border-radius: 3px;
        min-height: 15vh;
        // background-color: aquamarine;
        position: relative;

        .date {
            position: absolute;
            bottom: -30px;
            right: 2px;
            text-shadow: -1px 1px 2px black;
            text-align: right;
            color: rgb(49, 49, 49)
        }

        @media (max-width:1024px) {
            min-height: 10vh;
            width: 60%;

        }

        &::after {
            position: absolute;
            display: block;
            left: -15px;
            // top: 5px;
            top: 1px;
            transform: rotate(225deg);
            content: '';
            width: 0;
            height: 0;
            border: 10px solid transparent;
            border-top-color: rgb(49, 49, 49);
            border-right-color: rgb(49, 49, 49);
            border-radius: 2px;
        }

        &::before {
            position: absolute;
            display: block;
            left: -8px;
            // top: 5px;
            z-index: 10;
            top: 4px;
            transform: rotate(225deg);
            content: '';
            width: 0;
            height: 0;
            border: 8px solid transparent;
            border-top-color: white;
            border-right-color: white;
        }

        .info {
            font-family: "华文细黑";
            font-size: 1.2em;
        }
    }
}
</style>