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

const editorRef = shallowRef()
// 内容 HTML
const valueHtml = ref('')
const emit = defineEmits(['uploadContent'])
const props = defineProps<{
  article: string
}>()


if (props.article != undefined) {
  valueHtml.value = props.article
}

watch(props, (newValue: any) => {
  valueHtml.value = props.article
})

function changeContent() {
  emit('uploadContent', valueHtml.value)
}

// const focus = inject('focus', true)
// 模拟 ajax 异步获取内容

const toolbarConfig = {}
const editorConfig = {
  placeholder: '说点啥吧', focus: true,
  MENU_CONF: {
    uploadImage: {
      server: '/api/yuurei/uploadImgArticle',

      fieldName: 'your-custom-name',

      // 单个文件的最大体积限制，默认为 2M
      maxFileSize: 20 * 1024 * 1024, // 1M

      // 最多可上传几个文件，默认为 100
      maxNumberOfFiles: 10,

      // 选择文件时的类型限制，默认为 ['image/*'] 。如不想限制，则设置为 []
      allowedFileTypes: ['image/*'],

      // 自定义上传参数，例如传递验证的 token 等。参数会被添加到 formData 中，一起上传到服务端。
      meta: {
        token: 'xxx',
        otherKey: 'yyy'
      },

      // 将 meta 拼接到 url 参数中，默认 false
      metaWithUrl: false,

      // 自定义增加 http  header
      headers: {
        Accept: 'text/x-json',
        otherKey: 'xxx'
      },

      // 跨域是否传递 cookie ，默认为 false
      withCredentials: true,

      // 超时时间，默认为 10 秒
      timeout: 5 * 1000, // 5 秒
    }
  }

}


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