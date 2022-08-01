<template>
  <div style="border: 1px solid #ccc;">
    <Toolbar
        style="border-bottom: 1px solid #ccc"
        :editor="editorRef"
        :defaultConfig="toolbarConfig"
        :mode="mode"
    />
    <Editor
        style="height:500px;"
        v-model="valueHtml"
        :defaultConfig="editorConfig"
        :mode="mode"
        @onCreated="handleCreated"
        @onChange="changeContent"
    />
  </div>
</template>

<script lang="ts" setup>
import {inject, nextTick, ref, watch} from 'vue';
import {useRouter} from 'vue-router'
import '@wangeditor/editor/dist/css/style.css' // 引入 css

import {onBeforeUnmount, shallowRef, onMounted} from 'vue'
import {Editor, Toolbar} from '@wangeditor/editor-for-vue'

const router = useRouter()

// const content = ref('')
const editorRef = shallowRef()
// 内容 HTML
const valueHtml = ref('')
const emit = defineEmits(['uploadContent'])
const props = defineProps<{
  article: object
}>()

if (props.article != undefined) {
  valueHtml.value = props.article.ArticleContent
}

watch(props, (newValue: any) => {
  valueHtml.value = props.article.ArticleContent
})

function changeContent() {
  emit('uploadContent', valueHtml.value)
}

// const focus = inject('focus', true)
// 模拟 ajax 异步获取内容
onMounted(() => {
  setTimeout(() => {

  }, 1500)

})


const toolbarConfig = {}
const editorConfig = {placeholder: '说点啥吧', focus: true}

function calling() {
}


// 组件销毁时，也及时销毁编辑器
onBeforeUnmount(() => {
  const editor = editorRef.value
  if (editor == null) return
  editor.destroy()

})

const handleCreated = (editor: any) => {
  editorRef.value = editor // 记录 editor 实例，重要！
  // editor.value.focus(true)
}
</script>
<style lang="less" scoped>
.w-e-text-container {
  height: 500px !important;
}

</style>