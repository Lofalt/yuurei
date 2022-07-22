<template>
    <div id="main" @wheel="listenScroll">
        <img class="test" v-for="(item, index) in picList" :src="item" :key="item" ref="waterFallItem"
            :style="{ animationDelay: (index % 4) * 0.1 + 's' }" />
    </div>
</template>

<script lang="ts" setup>

import axios from "axios"
import qs from 'qs'
import { reactive, ref, onMounted, watch, nextTick } from "vue";


const isShow = ref(true)
const picList = ref([])
const offset = ref(20)
const colRaw = ref(2)
const screenWidth = ref(0)
const waterFallItem = ref<HTMLSelectElement>()
function listenScroll(event: any) {
    console.log(event)
    console.log(document.getElementById("main"))
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
    // setTimeout(() => {
    nextTick(() => {
        let childs = document.getElementsByClassName("test") as HTMLSelectElement
        console.log("waterfall被调用")
        var boxWidth = childs[0].offsetWidth
        var heightArr = [], boxHeight = 0, minBoxHeight = 0, minIndex = 0
        var col: number = colRaw.value
        for (let i = 0; i < childs.length; i++) {
            boxHeight = childs[i].offsetHeight
            console.log(boxHeight)
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
    })
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
                waterFall()
            })()
        }

    })
})
watch(screenWidth, () => {
    if (screenWidth.value >= 1000) {
        colRaw.value = 4
    }
    else if (screenWidth.value < 1000 && screenWidth.value >= 700) {
        colRaw.value = 3

    } else if (screenWidth.value < 700) {
        colRaw.value = 2
    }
})
watch(colRaw, () => {
    let childs = document.getElementsByClassName("test") as HTMLSelectElement
    for (let i = 0; i < childs.length; i++) {
        childs[i].style.width = 100 / colRaw.value + "%"
    }
    waterFall()
})
</script>

<style lang="less" scoped>
.test {

    left: 0px;
    width: 40%;
    height: auto;
    // display: block;
    box-sizing: border-box;
    /* padding: 5px; */
    padding: 8px;
    /* margin:15px; */
    // float: left;
    /* transition: all 0.5s; */
    // animation: enter 1s;
    // animation-timing-function: ease-in-out;
    // animation-fill-mode: backwards;
}

#main {
    // position: relative;
    // margin: 0 auto;
    // height: 100%;
    // width: 80%;
    width: 80%;
    float: left;
    height: 100%;
    overflow: auto;
    position: relative;
    bottom: 0;
    margin: auto auto;
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