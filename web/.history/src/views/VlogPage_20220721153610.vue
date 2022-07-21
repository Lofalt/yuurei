<script setup lang="ts">
import axios from 'axios'
import qs from 'qs'
import { ref } from 'vue'
import { onMounted } from 'vue'

const list = ref<unknown[]>([])
const loading = ref(false)
const over = ref(false)
let num = 20
axios.defaults.baseURL = "/api"
const isMounted = ref(false)
function getNext() {
    console.log("next")
    loading.value = true
    if (num > 50) {
        loading.value = false
        return
    }
    num += 20
    axios.post("/api/img", qs.stringify({ num: num })).then((result) => {
        if (list.value.length === result.data.data.list.length) {
            loading.value = false
            return
        }
        const newList = ref<unknown[]>([])
        for (let i = num - 20; i < num; i++) {
            console.log("mounted")
            let obj = { cover: '', title: "hello", outline: "?", count: 0 }
            obj.cover = result.data.data.list[i]
            obj.count = i % 4
            newList.value.push(obj)
        }
        list.value = [...list.value, ...newList.value]
        loading.value = false
        console.log(list.value)
    })
}
onMounted(() => {
    loading.value = true
    axios.post("/api/img", qs.stringify({ num: 20 })).then((result) => {
        const newList = ref<unknown[]>([])
        for (let i = 0; i < 20; i++) {
            console.log("mounted")
            let obj = { cover: '', title: "hello", outline: "?", count: 0 }
            obj.cover = result.data.data.list[i]
            obj.count = i % 4
            newList.value.push(obj)
        }
        list.value = newList.value
        loading.value = false
        // console.log(result)
        // list.value = list.value.concat(result.data.data.list)
        // console.log(list.value)
    }
    )
    isMounted.value = true
}
)
</script>

<template>
    <div class="father-box">
        <v3-waterfall class="waterfall" scrollBodySelector=".father-box" :isMounted="isMounted" :list="list"
            srcKey="cover" :gap="10" :colWidth="200" :distanceToScroll="100" :isLoading="loading" :isOver="over"
            @scrollReachBottom="getNext" :dotsNum="3">
            <template v-slot:default="slotProp">
                <div class="list-item" :style="`--i:` + slotProp.item.count">
                    <a :href="'https://gkshi.com/blog/' + slotProp.item._id">
                        <div class="cover-wrapper">
                            <img v-if="slotProp.item.cover" :src="slotProp.item.cover" class="cover" />
                        </div>
                        <div class="brief">
                            <!-- <h3>{{ slotProp.item.title }}</h3> -->
                            <!-- <p>{{ slotProp.item.outline }}</p> -->
                        </div>
                    </a>
                </div>
            </template>
        </v3-waterfall>
    </div>
</template>

<style lang="scss" scoped>
.father-box {
    width: 80vw;
    height: 300px;
    /* 父元素一定要指定高度 */
    overflow-y: scroll;
    /* 一定要制定父元素超出滚动 */
}
</style>
