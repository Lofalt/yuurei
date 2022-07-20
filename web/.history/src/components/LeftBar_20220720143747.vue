<template>
    <div class="leftbar">
        <div class="profile"></div>
        <div class="name">LOFALT</div>
        <div class="icon">
            <div></div>
            <div></div>
            <div></div>
        </div>
        <div class="footer">
            <div class="footerbox"></div>
            <div class="footerbox"></div>
            <div class="footerbox"></div>
        </div>
    </div>
</template>

<script>
export default {
    name: "LeftBar",
};
</script>

<script lang="ts" setup>

</script>

<style lang="less" scoped>
.leftbar {
    color: white;
    position: relative;
    float: left;
    width: 16.5%;
    min-height: 100vh;
    // background-color: rgb(131, 224, 255);
    display: block;
    // box-shadow: 0 0 40px rgba(0, 0, 0, 0.2);

    @media (max-width:500px) {
        display: none;
    }

    div {
        margin: 0 auto;
        text-align: center;
    }

    .profile {
        border: 3px rgb(79, 79, 79) solid;
        position: relative;
        margin: 0 auto;
        // display: flex;
        width: 12.2vw;
        height: 12.2vw;
        border-radius: 50%;
        background: white;
        text-align: center;
        margin-top: 30%;
        // justify-content: center;
        // align-items: center;
        color: black;
        background-image: url("../assets/-421408862.jpg");
        background-size: cover;
        background-position: center;
        background-repeat: no-repeat;
    }

    .name {
        margin-top: 55px;
        // background-color: #fff;
        width: 16.5vw;
        color: rgb(59, 53, 53);
        text-align: center;
        font-size: 1.5em;
        text-justify: inter-word;
        letter-spacing: 5px;
        text-indent: 5px;
        word-wrap: break-word;
        text-shadow: 0px 0px 2px rgba(235, 235, 235, 0.5);
    }

    .icon {
        display: flex;
        justify-content: space-between;
        align-items: center;
        width: 8vw;
        margin: 20px auto;

        div {
            background-color: rgb(20, 21, 22);
            width: 20px;
            height: 20px;
        }
    }

    .footer {
        display: flex;
        margin: 50px 20%;
        min-height: 25vh;
        // background-color: aliceblue;
        justify-content: space-around;
        align-items: center;
        flex-direction: column;

        .footerbox {
            width: 80%;
            background-color: #fcfae0;
            height: 30px;
            border-radius: 10px;
        }
    }

}
</style>