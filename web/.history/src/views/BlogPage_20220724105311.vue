<template>
    <div class="container">
        <div class="btnHeader">
            {{ content }}
        </div>
        <div>
            <!-- <keep-alive> -->
            <!-- <router-view class="info"></router-view> -->
            <!-- </keep-alive> -->
        </div>
        <!-- <tiny-mce v-model="content" @change="change2"></tiny-mce> -->
        <div style="border: 1px solid #ccc">
            <Toolbar style="border-bottom: 1px solid #ccc" :editor="editorRef" :defaultConfig="toolbarConfig"
                :mode="mode" />
            <Editor style="height: 100px; overflow-y: auto;" v-model="valueHtml" :defaultConfig="editorConfig"
                :mode="mode" @onCreated="handleCreated" />
        </div>
    </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router'
import TinyMce from "../components/tinymce/TinyMce.vue"
import '@wangeditor/editor/dist/css/style.css' // 引入 css

import { onBeforeUnmount, shallowRef, onMounted } from 'vue'
import { Editor, Toolbar } from '@wangeditor/editor-for-vue'
const router = useRouter()

function change() {
    router.push("/home/pages/anli/1")
}
function change2() {

}
const content = ref('')

const editorRef = shallowRef()

const mode = ref('default')
// 内容 HTML
const valueHtml = ref('<p>hello</p>')

// 模拟 ajax 异步获取内容
onMounted(() => {
    setTimeout(() => {
        valueHtml.value = '<p>模拟 Ajax 异步设置内容</p>'
    }, 1500)
    console.log(editorRef.value)
})

const toolbarConfig = {
    // toolbarKeys: [
    //     // 菜单 key
    //     'headerSelect',

    //     // 分割线
    //     '|',

    //     // 菜单 key
    //     'bold', 'italic',

    //     // 菜单组，包含多个菜单
    //     {
    //         key: 'group-more-style', // 必填，要以 group 开头
    //         title: '更多样式', // 必填
    //         iconSvg: '<svg>....</svg>', // 可选
    //         menuKeys: ["through", "code", "clearStyle"] // 下级菜单 key ，必填
    //     },
    //     // 继续配置其他菜单...
    // ]
}
console.log(Toolbar)
const editorConfig = { placeholder: '请输入内容...' }

// 组件销毁时，也及时销毁编辑器
onBeforeUnmount(() => {
    const editor = editorRef.value
    if (editor == null) return
    editor.destroy()
})

// editorRef.value.getAllMenuKeys()
const handleCreated = (editor: any) => {
    editorRef.value = editor // 记录 editor 实例，重要！
}
</script>

<style lang="less" scoped>
.container {
    width: 80%;
    height: 100%;
    /* background-color: #fff; */
    /* margin: 0 auto; */
    overflow: auto;
    // scrollbar-width: none;
}

::-webkit-scrollbar {
    width: 3px;
    // background: rgba(255, 255, 255, 0.5);
}

.btnHeader {
    /* width: 100%; */
    /* height: 100%; */
    width: 80%;
    margin: 8vh auto;
    text-align: center;
    border-bottom: 3px solid rgba(255, 255, 255, 0.1);
    padding-bottom: 5vh;
}

.button {
    margin: 0 auto;
    background-color: rgb(255, 255, 255);
    border: none;
    padding: 15px 25px;
    margin: 20px;
    border-radius: 10px;
    color: rgb(43, 43, 43);
    cursor: pointer;

    &:hover {
        background-color: rgb(252, 125, 125);
    }
}

.info {
    // width: 100%;
    height: 100vh;
    // background-color: aliceblue;
    overflow: auto;
}
</style>