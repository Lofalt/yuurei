<script lang="ts" setup>
import { LazyImg, Waterfall } from 'vue-waterfall-plugin-next'
import 'vue-waterfall-plugin-next/style.css'
import axios from 'axios'
import { onMounted, ref } from 'vue';
import qs from 'qs';
const list = ref<unknown[]>([])

onMounted(() => {
    // loading.value = true
    axios.defaults.baseURL = "/api"
    axios.post("/api/img", qs.stringify({ num: 20 })).then((result) => {
        const newList = ref<unknown[]>([])
        for (let i = 0; i < 20; i++) {
            console.log("mounted")
            let obj = { src: '', title: "hello", outline: "?", count: 0 }
            obj.src = result.data.data.list[i]
            obj.count = i % 4
            newList.value.push(obj)
        }
        list.value = newList.value
        // loading.value = false
        // console.log(result)
        // list.value = list.value.concat(result.data.data.list)
        // console.log(list.value)
    }
    )
    // isMounted.value = true
}
)

</script>

<template>
    <div class="father">
        <Waterfall :list="list">
            <template #item="{ item, url, index }">
                <div class="card">
                    <LazyImg :url="url" />
                    <p class="text">这是内容</p>
                </div>
            </template>
        </Waterfall>
    </div>
</template>

<style lang="less" scoped>
// .father {
// overflow: auto;
// }
</style>