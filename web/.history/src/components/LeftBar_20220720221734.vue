<template>
    <div class="leftbar">
        <div class="profile"></div>
        <div class="name">LOFALT</div>
        <div class="icon">
            <div>
                <svg class="icon" aria-hidden="false">
                    <use xlink:href="#icon-weibo"></use>
                </svg>
            </div>
            <div>


                <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="2em"
                    height="2em" viewBox="0 0 24 24">
                    <path fill="#888888"
                        d="M17.525 11.378c1.263.392 2.669 1.336 2.669 3.004c0 2.763-3.98 6.239-9.964 6.239c-4.565 0-9.23-2.213-9.23-5.852c0-1.902 1.204-4.102 3.277-6.177c2.773-2.77 6.004-4.033 7.219-2.816c.537.537.588 1.464.244 2.572c-.178.557.525.25.525.25c2.24-.938 4.196-.994 4.909.027c.38.543.343 1.306-.008 2.19c-.163.407.048.471.36.563zm-7.282 7.939c3.641-.362 6.401-2.592 6.167-4.983c-.237-2.391-3.382-4.038-7.023-3.677c-3.64.36-6.403 2.59-6.167 4.98c.237 2.394 3.382 4.039 7.023 3.68zM6.16 14.438c.754-1.527 2.712-2.39 4.446-1.94c1.793.463 2.707 2.154 1.976 3.8c-.744 1.682-2.882 2.578-4.695 1.993c-1.752-.566-2.493-2.294-1.727-3.853zm1.446 2.587c.568.257 1.325.013 1.676-.55c.346-.568.163-1.217-.407-1.459c-.563-.237-1.291.008-1.64.553c-.354.547-.189 1.202.371 1.456zm2.206-1.808c.219.092.501-.012.628-.231c.123-.22.044-.466-.178-.548c-.216-.084-.486.018-.613.232c-.123.214-.054.458.163.547zM19.873 9.5a.725.725 0 1 1-1.378-.451a1.38 1.38 0 0 0-.288-1.357a1.395 1.395 0 0 0-1.321-.425a.723.723 0 1 1-.303-1.416a2.836 2.836 0 0 1 3.29 3.649zm-3.916-6.575A5.831 5.831 0 0 1 21.5 4.72a5.836 5.836 0 0 1 1.22 5.704a.838.838 0 0 1-1.06.54a.844.844 0 0 1-.542-1.062a4.143 4.143 0 0 0-4.807-5.327a.845.845 0 0 1-.354-1.65z">
                    </path>
                </svg>



            </div>
            <div>

            </div>
        </div>
        <div class="footer">
            <button class="footerbox" @click="about">HOME</button>
            <button class="footerbox" @click="home">ABOUT</button>
            <button class="footerbox">INFO</button>
        </div>
    </div>
</template>

<script lang="ts">

export default {
    name: "LeftBar",
};
</script>

<script lang="ts" setup>
import { useRouter } from 'vue-router';

const router = useRouter()

// function about() {
//     router.push("/About")
// }
// function home() {
//     router.push("/home")
// }
</script>

<style lang="less" scoped>
.icon {
    width: 1em;
    height: 1em;
    vertical-align: -0.15em;
    fill: currentColor;
    // overflow: hidden;
}

.leftbar {
    border-right: 1px solid rgba(0, 0, 0, 0.2);
    color: white;
    position: fixed;
    float: left;
    width: 16.5%;
    min-height: 100vh;
    // background-color: rgb(131, 224, 255);
    display: block;
    box-shadow: 0 0 40px rgba(0, 0, 0, 0.2);
    backdrop-filter: blur(3px);

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
        background-image: url("../assets/ec8ce499ly1h1hhm1o6m1j22pg1ww4qp.jpg");
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
        font-size: 1.75em;
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
            // background-color: rgb(20, 21, 22);
            color: rgb(241, 91, 91);
            font-size: 1em;
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
            cursor: pointer;
            width: 80%;
            background-color: #fcfae0;
            height: 30px;
            border-radius: 10px;
        }
    }

}
</style>