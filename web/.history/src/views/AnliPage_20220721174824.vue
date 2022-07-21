<template>
    <div id="main">
        <img class="test" v-for="(item, index) in picList" :src="item" :key="item"
            :style="{ animationDelay: (index % 4) * 0.1 + 's' }" />
    </div>
</template>

<script lang="ts" setup>

import axios from "axios"
import qs from 'qs'
import { reactive, ref, onMounted, watch, nextTick } from "vue";


const isShow = ref(true)
let picList = reactive([])
const offset = ref(20)
const colRaw = ref(4)
const screenWidth = ref(0)

function getPic(num: number) {
    axios.defaults.baseURL = "/api"
    axios.post("/api/img", qs.stringify({ num: num })).then((result) => {
        console.log(result)
        picList = result.data.data.list
        console.log(picList)
        // nextTick(() => {
        setTimeout(() => {
            waterFall()

        }, 5000);

        // })
    })
}
function waterFall() {
    setTimeout(() => {

        let childs = document.getElementsByClassName("test") as HTMLSelectElement
        console.log(childs)
        var boxWidth = childs[0].offsetWidth
        var heightArr = [], boxHeight = 0, minBoxHeight = 0, minIndex = 0
        var col: number = colRaw.value
        for (let i = 0; i < childs.length; i++) {
            boxHeight = childs[i].offsetHeight
            if (i < col) {
                heightArr.push(boxHeight + 10)
                childs[i].style.position = 'absolute'
                childs[i].style.left = i * boxWidth + 'px'
                childs[i].style.top = '0px'
                console.log(heightArr, boxHeight, minBoxHeight, minIndex)
            } else {
                console.log(heightArr, boxHeight, minBoxHeight, minIndex)
                minBoxHeight = heightArr[minBox(heightArr)]
                minIndex = minBox(heightArr)
                childs[i].style.position = 'absolute'
                childs[i].style.left = minIndex * boxWidth + 'px'
                childs[i].style.top = minBoxHeight + 'px'
                heightArr[minIndex] += boxHeight
            }
            // console.log(childs[i][0].style)
        }
    }, 100);
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
        // window.onresize = () => {
        //   return (() => {
        //     this.screenWidth = document.body.clientWidth
        //     this.waterFall("main", "test")
        //   })()
        // }

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