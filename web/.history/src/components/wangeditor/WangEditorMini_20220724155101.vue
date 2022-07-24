<template>
    <div class="son">
        <!-- <button>点我</button> -->
        <Toolbar class="toolbar" :editor="editorRef" :defaultConfig="toolbarConfig" :mode="mode" />
        <Editor class="editor" v-model="content" :defaultConfig="editorConfig" :mode="mode"
            @onCreated="handleCreated" />
    </div>
</template>

<script lang="ts" setup>
import { inject, ref, watch } from 'vue';
import { useRouter } from 'vue-router'
import '@wangeditor/editor/dist/css/style.css' // 引入 css

import { onBeforeUnmount, shallowRef, onMounted } from 'vue'
import { Editor, Toolbar } from '@wangeditor/editor-for-vue'
const router = useRouter()

const content = ref('')

const editorRef = shallowRef()

const mode = ref('default')
// 内容 HTML
const valueHtml = ref('<p>hello</p>')

// 模拟 ajax 异步获取内容
onMounted(() => {
    setTimeout(() => {
        valueHtml.value = '<p>模拟 Ajax 异步设置内容</p>'
        console.log(editorRef.value.getAllMenuKeys())
    }, 1500)

})

const toolbarConfig = {
    toolbarKeys: [
        // 菜单 key
        'emotion',

        // 分割线
        // '|',

        // 菜单 key
        'color',

        // 菜单组，包含多个菜单
        // {
        //     key: 'group-more-style', // 必填，要以 group 开头
        //     title: '更多样式', // 必填
        //     iconSvg: '<svg>....</svg>', // 可选
        //     menuKeys: ["through", "code", "clearStyle"] // 下级菜单 key ，必填
        // },
        // 继续配置其他菜单...
    ]
}
// console.log(Toolbar)
const editorConfig = { placeholder: '说点啥吧', focus: true }


const focus = inject('focus', true)

watch(focus, (newValue, oldValue) => {
    console.log(newValue)
    // if (newValue) {
    editorRef.value.focus()
    // }
})
// 组件销毁时，也及时销毁编辑器
onBeforeUnmount(() => {
    const editor = editorRef.value
    if (editor == null) return
    editor.destroy()

})

const handleCreated = (editor: any) => {
    editorRef.value = editor // 记录 editor 实例，重要！
    editor.value.focus(true)
}
</script>

<style lang="less" scoped>
.editor {
    height: 80% !important;
    // height: auto !important;

    overflow: auto;
    overflow-y: auto;

}

.toolbar {
    border-bottom: 3px solid rgb(49, 49, 49);
    margin: 5px;
}

.son {
    border: 4px solid rgb(49, 49, 49);
    height: 100%;
    border-radius: 2px;
    // overflow: auto;

    @media (max-width:1024px) {
        // height: 100%;    
        height: 40vh;
    }
}

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