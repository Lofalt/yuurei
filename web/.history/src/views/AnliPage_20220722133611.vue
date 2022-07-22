<template>
    <div id="main" @wheel="listenScroll">
        <img class="test" v-for="(item, index) in picList" :src="item" :key="item" ref="waterFallItem"
            @click="zoom(index)" :style="{ animationDelay: (index % 4) * 0.1 + 's' }" />
    </div>
    <transition>
        <div @click="zoomout" v-show="hover" class="photoInfo"
            :style="{ backgroundImage: `url('` + showingPage + `')` }">

        </div>
    </transition>

</template>

<script lang="ts" setup>

import axios from "axios"
import qs from 'qs'
import { reactive, ref, onMounted, watch, nextTick, inject, Ref } from "vue";


const isShow = ref(true)
const picList = ref([])
const offset = ref(20)
const colRaw = ref(2)
const screenWidth = ref(0)
const waterFallItem = ref<HTMLSelectElement>()
function listenScroll(event: any) {
    const box = document.getElementById("main") as HTMLSelectElement
    // console.log(box.offsetHeight)
    console.log(box.scrollTop)
    console.log(box.scrollHeight)
    console.log(box.offsetHeight)
    if (box.scrollTop + box.offsetHeight > box.scrollHeight) {
        offset.value += 10
        getPic(offset.value)
    }

}
const hover = inject('hover') as Ref<boolean>
let widthTmp = ''
let leftTmp = ''
// let photoNum = 20
const showingPage = ref('')

function zoomout() {

    // childs[index].style.width = widthTmp
    // childs[index].style.left = leftTmp
    // childs[index].style.zIndex = "0"    
    hover.value = false
}
function zoom(event: number) {
    let childs = document.getElementsByClassName("test") as HTMLSelectElement
    showingPage.value = childs[event].getAttribute("src") as string
    hover.value = true
    // let childs = document.getElementsByClassName("test") as HTMLSelectElement
    // // console.log("waterfall被调用")
    // widthTmp = childs[event].style.width;
    // leftTmp = childs[event].style.left
    // childs[event].style.width = "100%"
    // childs[event].style.left = "0"
    // childs[event].style.zIndex = "+10"

}

function getPic(num: number) {
    axios.defaults.baseURL = "/api"
    axios.post("/api/img", qs.stringify({ num: num })).then((result) => {
        picList.value = result.data.data.list
        setTimeout(() => {
            waterFall()

        }, 100);
    })
}
function waterFall() {
    setTimeout(() => {
        // nextTick(() => {
        let childs = document.getElementsByClassName("test") as HTMLSelectElement
        console.log("waterfall被调用")
        // showingPage=childs[0]
        showingPage.value = childs[0].getAttribute("src") as string
        console.log(showingPage.value)
        var boxWidth = childs[0].offsetWidth
        var heightArr = [], boxHeight = 0, minBoxHeight = 0, minIndex = 0
        var col: number = colRaw.value
        for (let i = 0; i < childs.length; i++) {
            boxHeight = childs[i].offsetHeight
            // console.log(boxHeight)
            // console.log(childs[i])
            if (i < col) {
                heightArr.push(boxHeight + 10)
                childs[i].style.position = 'absolute'
                childs[i].style.left = i * boxWidth + 'px'
                childs[i].style.top = '0px'
            } else {
                // console.log(heightArr, boxHeight, minBoxHeight, minIndex)
                minBoxHeight = heightArr[minBox(heightArr)]
                minIndex = minBox(heightArr)
                childs[i].style.position = 'absolute'
                childs[i].style.left = minIndex * boxWidth + 'px'
                childs[i].style.top = minBoxHeight + 'px'
                heightArr[minIndex] += boxHeight
            }
        }
    }, 500)
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
                getPic(offset.value)
                offset.value += 10
            }
        }
        window.onresize = () => {
            return (() => {
                screenWidth.value = document.body.clientWidth
                // waterFall()
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
.v-enter-from,
.v-leave-to {
    opacity: 0;
}

.v-enter-active {
    transition: all .8s;
}

.v-leave-active {
    transition: all .1s;
}

.photoInfo {
    background-color: rgba(43, 43, 43, 0.738);
    width: 60vw;
    // width: auto;
    height: 100%;
    position: absolute;
    // position: fixed;
    // z-index: 10;
    background-size: contain;
    background-repeat: no-repeat;
    // animation: zoom 1s ease;
    cursor: pointer;
    background-position: center;
}

@keyframes zoom {
    from {
        opacity: 0;
    }
}

.test {
    cursor: pointer;
    left: 0px;
    width: 50%;
    // height: 100%;
    box-sizing: border-box;
    padding: 8px;
    // margin: 8px;
    /* margin:15px; */
    // float: left;
    // transition: all 0.5s;
    transition: all 1s;
    animation: enter 1s;
    animation-timing-function: ease-in-out;
    animation-fill-mode: backwards;
}

#main {
    // position: relative;
    // margin: 0 auto;
    height: 100%;
    // width: 80%;
    // margin: 0 auto;
    width: 90%;
    float: left;
    height: 100%;
    overflow: auto;
    position: relative;
    bottom: 0;
    margin: 0 3vw;
    padding: 10vh;

    &::-webkit-scrollbar {
        width: 0;
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
        transform: translateY(1000px);
        opacity: 0;
    }
}
</style>