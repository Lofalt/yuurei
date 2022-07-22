<template>
    <waterfall :list="list" :cols="4" :footerHeight="60" :nomore="noMore" @preLoaded="loadedEnd"
        @scrollReachBottom="loadMore">
        <!-- 自定义图片 -->
        <template v-slot:default="{ item }">
            <div class="imgbox">
                <img :src="item.imgSrc" :alt="item.title">
            </div>
        </template>
        <!-- 自定义底部信息块 -->
        <template v-slot:footer="{ item }">
            <span>{{ item.title }}</span>
        </template>
        <!-- 自定义加载状态标示 -->
        <template v-slot:loading>
            <span>正在拼命加载中...</span>
        </template>
        <template v-slot:nomore>
            <span>没有更多数据了...</span>
        </template>
    </waterfall>
</template>

<script>
import { defineComponent } from 'vue'
import VueWaterfall from 'vue3-waterfall'

export default defineComponent({
    data() {
        return {
            list: [],
            noMore: false
        }
    },
    components: {
        VueWaterfall
    },
    mounted() {
        axios.post("/api/img", qs.stringify({ num: 20 })).then((result) => {
            let newList = []
            for (let i = 0; i < 20; i++) {
                console.log("mounted")
                let obj = { imgSrc: '', title: "hello", outline: "?", count: 0 }
                obj.cover = result.data.data.list[i]
                obj.count = i % 4
                newList.value.push(obj)
            }
            this.list = newList
            loading.value = false
            // console.log(result)
            // list.value = list.value.concat(result.data.data.list)
            // console.log(list.value)
        }
        )
    }
    methods: {
        loadMore() {
            console.log('加载更多')
            if (this.list.length >= 30) {
                this.noMore = true
            }
        },
        loadedEnd() {
            console.log('图片预加载完成')
        }
    },
})
</script>