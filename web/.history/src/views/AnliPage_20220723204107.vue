<template>
    <div id="main" @wheel="listenScroll">
        <!-- <button class="scrollToTop" @click="scrollToTop">dmwo</button> -->
        <img class=" test" v-for="(item, index) in picList" :src="item" :key="item" ref="waterFallItem"
            @click="zoom(index)" :style="{ animationDelay: (index % 2) * 0.1 + 's' }" />
    </div>
    <!-- <loading-com class="loading"></loading-com> -->
    <loading-com class="loading" v-show="isLoading"></loading-com>
    <!-- <div class="loading" v-show="offset == 15">--</div> -->
    <transition>
        <div @click="zoomout" v-if="hover" class="photoInfo" :style="{ backgroundImage: `url('` + showingPage + `')` }">

        </div>
    </transition>
</template>

<script lang="ts" setup>

import axios from "axios"
import qs from 'qs'
import { reactive, ref, onMounted, watch, nextTick, inject, Ref } from "vue";
import LoadingCom from "../components/LoadingCom.vue"

const isShow = ref(true)
const picList = ref<any>([])
const offset = ref(5)
const colRaw = ref(3)
const screenWidth = ref(0)
const waterFallItem = ref<HTMLSelectElement>()
const showingPage = ref('')
const hover = inject('hover') as Ref<boolean>
const showVlog = inject('showVlog') as Ref<boolean>
const isLoading = ref(false)

function scroll(scrollDuration: number, box: any) {
    var cosParameter = box.scrollY / 2,
        scrollCount = 0,
        oldTimestamp = performance.now();
    function step(newTimestamp: any) {
        scrollCount += Math.PI / (scrollDuration / (newTimestamp - oldTimestamp));
        if (scrollCount >= Math.PI) box.scrollTo(0, 0);
        if (box.scrollY === 0) return;
        box.scrollTo(0, Math.round(cosParameter + cosParameter * Math.cos(scrollCount)));
        oldTimestamp = newTimestamp;
        // box.requestAnimationFrame(step);
    }
    // box.requestAnimationFrame(step);q
}
function scrollToTop() {
    const box = document.getElementById("main") as HTMLSelectElement
    // box.scrollTo(0, 0)
    scroll(500, box)
}

function listenScroll(event: any) {
    const box = document.getElementById("main") as HTMLSelectElement
    if (box.scrollTop + box.offsetHeight + 50 > box.scrollHeight) {
        // offset.value += 5
        if (!isLoading.value && offset.value < 15) {
            isLoading.value = true
            setTimeout(() => {
                getNext(offset.value)

            }, 3000);

        } else {
            return
        }
    }

}


function zoomout() {
    hover.value = false
}
function zoom(event: number) {
    let childs = document.getElementsByClassName("test") as HTMLSelectElement
    showingPage.value = childs[event].getAttribute("src") as string
    hover.value = true

}
function getNext(num: number) {

    axios.defaults.baseURL = "/api"
    isLoading.value = true

    axios.post("/api/img", qs.stringify({ num: num + 5 })).then((result) => {
        for (let i = num; i < result.data.data.list.length; i++) {
            picList.value.push(result.data.data.list[i])
        }
        offset.value += 5
        waterFall()
    })
}
function getPic(num: number) {
    axios.defaults.baseURL = "/api"
    axios.post("/api/img", qs.stringify({ num: num })).then((result) => {
        picList.value = result.data.data.list
        setTimeout(() => {
            waterFall()

        }, 500);
    })
}
function waterFall() {
    setTimeout(() => {
        let childs = document.getElementsByClassName("test") as HTMLSelectElement
        // showingPage.value = childs[0].getAttribute("src") as string
        var boxWidth = childs[0].offsetWidth
        var heightArr = [], boxHeight = 0, minBoxHeight = 0, minIndex = 0
        var col: number = colRaw.value
        for (let i = 0; i < childs.length; i++) {
            boxHeight = childs[i].offsetHeight
            if (i < col) {
                heightArr.push(boxHeight + 50)
                childs[i].style.position = 'absolute'
                childs[i].style.left = i * boxWidth + screenWidth.value / 10 + 'px'
                childs[i].style.top = 5 + '0px'
                childs[i].style.opacity = '1'
                childs[i].style.transform = "translateY(0)"
            } else {
                minBoxHeight = heightArr[minBox(heightArr)]
                minIndex = minBox(heightArr)
                childs[i].style.position = 'absolute'
                childs[i].style.left = minIndex * boxWidth + screenWidth.value / 10 + 'px'
                childs[i].style.top = minBoxHeight + 'px'
                childs[i].style.opacity = '1'
                childs[i].style.transform = "translateY(0)"
                heightArr[minIndex] += boxHeight
            }
        }
        isLoading.value = false
    }, 600)
}

function minBox(box: any) {
    var j = 0
    for (let i = 0; i < box.length; i++) {
        if (box[j] > box[i]) {
            j = i
        }
    }
    return j
}
onMounted(() => {
    getPic(offset.value)
    nextTick(() => {
        // waterFall()

        window.onscroll = () => {
            let scrollTop = document.documentElement.scrollTop; //获取滚动距离
            let scrollHeigh = document.documentElement.scrollHeight; //获取整个页面的高度
            let clientHeigh = document.documentElement.clientHeight; //获取
            if (scrollTop + clientHeigh - scrollHeigh >= -50) {
                getNext(offset.value)
            }
        }
        window.onresize = () => {
            return (() => {
                screenWidth.value = document.body.clientWidth
            })()
        }

    })
})
watch(screenWidth, () => {
    // if (screenWidth.value >= 1000) {
    //     colRaw.value = 2
    // }
    // else if (screenWidth.value < 1000 && screenWidth.value >= 700) {
    //     colRaw.value = 2

    // }
    waterFall()
    // else if (screenWidth.value < 700) {
    // colRaw.value = 2
    // }
})
watch(colRaw, () => {
    let childs = document.getElementsByClassName("test") as HTMLSelectElement
    // for (let i = 0; i < childs.length; i++) {
    //     childs[i].style.width = 60 / colRaw.value + "%"
    // }
})
</script>

<style lang="less" scoped>
.loading {
    // position: relative;
    bottom: 0;
    // height: 5%;
    // float: left;
    text-align: center;
}

.v-enter-from,
.v-leave-to {
    opacity: 0;
}

.v-enter-active {
    transition: all .8s;
}

.v-leave-active {
    transition: all .3s;
}

.photoInfo {
    background-color: rgba(43, 43, 43, 0.738);
    width: 100%;
    // width: auto;
    height: 100%;
    position: absolute;
    // position: fixed;
    top: 0;
    left: 0;
    // z-index: 10;
    background-size: contain;
    background-repeat: no-repeat;
    // animation: zoom 1s ease;
    cursor: pointer;
    background-position: center;
    // z-index: 10;
}

@keyframes zoom {
    from {
        opacity: 0;
    }
}

.scrollToTop {
    position: fixed;
    z-index: 10;
    top: 100px;
}

.test {
    cursor: pointer;
    left: 0px;
    width: 33.3%;
    // height: 100%;
    box-sizing: border-box;
    padding: 8px;
    // margin: 8px;
    /* margin:15px; */
    // float: left;
    transition: all 0.5s ease-in-out;
    // transition: all 1s;
    // animation: enter 1s;
    // animation-timing-function: ease-in-out;
    // animation-fill-mode: backwards;
    // box-shadow: 0 0 1px 1px rgba(43, 43, 43, 0.738);
    opacity: 0;
    transform: translateY(1000px);

    &:hover {
        transform: translate(-10px, -10px) !important;
        box-shadow: -2px 2px 5px 1px rgba(0, 0, 0, 0.1);
        transition: all .2s ease;

    }
}

#main {
    // position: relative;
    // margin: 0 auto;
    height: 100%;
    // width: 80%;
    // margin: 0 auto;
    width: 80%;
    // float: left;
    // height: 100%;
    overflow: auto;
    position: relative;
    // bottom: 0;
    margin: 0 auto;
    padding-top: 10vh;

    &::-webkit-scrollbar {
        width: 0;
        // transition: all 1s;
    }

    &::after,
    &::before {
        content: "";
        // overflow: hidden;
        display: block;
        height: 0;
        clear: both;
        visibility: hidden;
    }


}


button {
    position: absolute
}

.box {
    width: 100%;
    background-color: yellow;
    height: 100px;
    transition: all 1s ease-in-out;
}


@keyframes enter {
    from {
        // transform: translateY(1000px);
        opacity: 0;
    }
}
</style>