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

                <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="1em"
                    height="1em" viewBox="0 0 24 24">
                    <path fill="currentColor"
                        d="M9.82 13.87c1.07 0 1.95.87 1.95 1.95a1.95 1.95 0 0 1-1.95 1.95c-1.08 0-1.95-.88-1.95-1.95c0-1.08.87-1.95 1.95-1.95M14.5 3.34l.68-.03c3.76 0 6.82 3.06 6.82 6.82l-.05.82l-1.19-.37l.02-.45c0-3.09-2.51-5.6-5.6-5.6l-.35.01l-.33-1.2m.82 2.89c2.06.07 3.73 1.77 3.76 3.83l-1.24-.38c-.19-1.12-1.06-2-2.17-2.18l-.35-1.27M2 15.41c-.03-.61.07-2.77 2.95-5.44c3.4-3.16 4.87-2.92 4.87-2.92s3.18-.3 1.24 3.41h.07c.47-.5 1.49-1.25 3.56-1.46c2.08-.21 2.08 1.5 1.81 2.7c1.88.94 3.06 2.33 3.06 3.88c0 2.82-3.93 5.11-8.78 5.11h-.28c-3.5 0-6.5-1.27-7.79-3.1c-.46-.62-.71-1.3-.71-2.01v-.17m7.82-3.49c-3.23 0-5.85 1.75-5.85 3.9s2.62 3.9 5.85 3.9c3.23 0 5.85-1.75 5.85-3.9s-2.62-3.9-5.85-3.9Z">
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
    overflow: hidden;
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
            color: aqua;
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