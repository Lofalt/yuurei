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
            srcKey="cover" :gap="10" :colWidth="300" :distanceToScroll="100" :isLoading="loading" :isOver="over"
            @scrollReachBottom="getNext" :dotsNum="3">
            <template v-slot:default="slotProp">
                <div class="list-item" :style="`--i:` + slotProp.item.count">
                    <!-- <a :href="'https://gkshi.com/blog/' + slotProp.item._id"> -->
                    <div class="cover-wrapper">
                        <img v-if="slotProp.item.cover" :src="slotProp.item.cover" class="cover" />
                    </div>
                    <div class="brief">
                        <!-- <h3>{{ slotProp.item.title }}</h3> -->
                        <!-- <p>{{ slotProp.item.outline }}</p> -->
                    </div>
                    <!-- </a> -->
                </div>
            </template>
        </v3-waterfall>
    </div>
</template>

<style lang="scss" scoped>
.father-box {
    // width: 70vw;
    height: 80vh;
    /* 父元素一定要指定高度 */
    overflow-y: scroll;
    /* 一定要制定父元素超出滚动 */
    padding-top: 50px;
}

.menu {
    position: fixed;
    left: 0;
    padding-top: 300px;
    width: 200px;
    height: 100vh;
    border-right: 1px solid #ccc;

    p {
        margin: 20px;
        cursor: pointer;
    }
}

.active {
    color: #20c180;
}

.content {
    padding-top: 30px;
    margin-left: 210px;
    background-color: #dee3e7;
}

.limit-box {
    height: 300px;
    overflow-y: scroll;
}

.waterfall,
.tags {
    width: 80%;
    margin: 0 auto;
    min-width: 1200px;
}

.tags-wrapper {
    padding: 10px 0;
    background-color: #fff;
    border-bottom: 1px solid #c4cdd4;
    margin-bottom: 15px;

    .tags {
        .tag {
            display: inline-block;
            width: 100px;
            font-size: 14px;
            color: #576575;
            padding: 5px 0;
            text-align: center;
            border: 1px solid #e8eaee;
            border-radius: 4px;
            margin: 0 10px 8px 0;

            &:hover {
                border-color: #8599ad;
            }

            &:not(.active) {
                cursor: pointer;
            }

            &.active {
                color: #ffffff;
                background-color: #576575;
                border-color: #576575;
            }
        }
    }
}

.brief {
    width: 100%;
    box-sizing: border-box;
    background-color: #fff;

    h3,
    >p {
        word-wrap: break-word;
        word-break: break-all;
    }

    h3 {
        text-align: center;
        padding: 20px 20px 12px;
        font-weight: 400;
        color: #22252a;
    }

    >p {
        padding: 0 20px 12px;
        color: #8a98a8;
        font-size: 13px;
    }
}

.outline-bottom {
    border-top: 1px solid #e7ebef;
    background-color: #f9fafb;
    padding: 10px;
    display: flex;
    justify-content: space-between;

    .article-tags,
    time {
        font-size: 12px;
        color: #73828c;
    }

    .article-tags {
        span {
            margin-right: 6px;

            &.tag:hover {
                text-decoration: underline;
                color: #000000;
                cursor: pointer;
            }
        }
    }
}

@keyframes enter {
    from {
        transform: translateY(500px);
        opacity: 0;
    }

    to {
        opacity: 1;
    }
}

.list-item {
    box-sizing: border-box;
    animation: enter .8s cubic-bezier(0.075, 0.82, 0.165, 1) backwards;
    animation-delay: calc((var(--i)/4)*0.1s);

    .cover-wrapper {
        overflow: hidden;
        background-color: #fff;
    }

    &:hover {
        box-shadow: 5px 5px 5px #ccc;

        @keyframes scaleImg {
            0% {
                transform: scale(1);
            }

            100% {
                transform: scale(1.5);
            }
        }
    }
}

.cover {
    width: 100%;
    vertical-align: middle;
}

/* ipad pro */
@media screen and(max-width: 1024px) {
    header {
        height: 40vh;
        background-size: 100% 100%;
    }
}

@media screen and(max-width: 1000px) {

    .waterfall,
    .tags {
        width: 100%;
        padding: 0 30px;
        box-sizing: border-box;
        min-width: auto;
    }
}

@media screen and(max-width: 576px) {
    header {
        height: 30vh;
        min-height: 300px;

        background-size: 150% 100%;
        background-position: 40% 0;
    }

    .waterfall,
    .tags {
        width: 100%;
        padding: 0;
        box-sizing: border-box;
        min-width: auto;
    }

    .tags {
        width: 90%;
    }
}
</style>
