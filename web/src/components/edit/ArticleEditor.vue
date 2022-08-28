<template>
  <div class="container">
    <upload-pic @confirm="acceptPic" name="上传头图" raw-src="" directory="articleHeaderPicture" :quality="30"
                :ratio="3.5"/>
    <div class="header">
      <input v-model="title" placeholder="输入标题">
    </div>
    <div class="summary">
      <input v-model="summary" placeholder="输入简介">
    </div>
    <div class="neck">
      <n-space vertical>
        <n-select v-model:value="category" :options="options" placeholder="选择分类"/>
      </n-space>
      <n-dynamic-tags v-model:value="tags" @create="onCreate"/>
    </div>
    <div>
    </div>
    <div style="height: 50px;width: 80%;">
      <n-color-picker :modes="[`hex`]"
                      :swatches="[
      '#FFFFFF',
      '#18A058',
      '#2080F0',
      '#F0A020',
      'rgba(208, 48, 80, 1)'
    ]"
                      size="small" v-model:value="tagColor"/>
    </div>

    <wang-editor class="editor" @upload-content="acceptContent"/>
    {{
      content.split(/<[^>]*>/).join("").length
    }}字
    <button class="sendBtn" @click="send">点我</button>
  </div>
</template>

<script lang="ts" setup>
import WangEditor from "../wangeditor/WangEditor.vue"
import {inject, ref, Ref, watch} from "vue";
import {NSpace, NSelect, NDynamicTags, NColorPicker} from "naive-ui";
// import axios from "../../request/index"
// import axios from "axios"
import axios from "@/request/index"
import UploadPic from "../file/UploadPic.vue"

const category = ref(null)
const title = ref('')
const summary = ref('')
const content: Ref<string> = ref('')
const tagColor = ref('#000000')
const options = ref([])
const headerPic = ref("")

function acceptPic(url: string) {
  headerPic.value = url
}

axios.get("/yuurei/articleCategory/all", {}).then((res: any) => {
  for (let i = 0; i < res.data.data.length; i++) {
    options.value.push({
      value: res.data.data[i].ID,
      label: res.data.data[i].ArticleCategoryName
    })
  }
})

const emit = defineEmits(['success'])
const tags = ref([])

function send() {
  var Tags = []
  for (let i = 0; i < tags.value.length; i++) {
    let tag = {
      TagName: tags.value[i].label,
      ID: tags.value[i].value
    }
    Tags.push(tag)
  }
  var ArticleCategory = {
    ID: category.value
  }
  var Article = {
    ArticleTitle: title.value,
    ArticleCategory: ArticleCategory,
    AriticleContent: content.value,
    Tags: Tags,
  }
  axios.request("/yuurei/article", "post", {
    ArticleTitle: title.value,
    ArticleCategory: ArticleCategory,
    ArticleSummary: summary.value,
    ArticleContent: content.value,
    Tags: Tags,
    HeaderPicture: headerPic.value
  }).then((res) => {
    emit('success', res)
  })
}

function acceptContent(cont: string) {
  content.value = cont
}

function onCreate(label: string) {
  //先发送请求 验证是否存在tag
  let value: any;
  let name = label
  axios.get("/yuurei/tag/" + label, {}).then((res: any) => {
    if (res.data.data != null) {
      value = res.data.data.ID
      if (tags.value[tags.value.length - 1].label == label) {
        tags.value[tags.value.length - 1].value = res.data.data.ID
      }
    }
  }).then(() => {
    if (value === undefined) {

      axios.request("/yuurei/tag", "post",
          {
            TagName: label,
            TagColor: tagColor.value,
          }).then((res: any) => {
        value = res.data.data.ID
        if (tags.value[tags.value.length - 1].label == label) {
          tags.value[tags.value.length - 1].value = res.data.data.ID
        }
      })
    }

  })
  //若不存在 ,即创建新数据行 返回新ID
  //若存在 返回标签ID
  return {label: name, value: value}
}
</script>
<style lang="less" scoped>


.sendBtn {
  position: absolute;
  right: 20px;
  top: 220px;
  padding: 10px 20px;
  outline: none;
  cursor: pointer;
  background-color: white;
  border: 4px solid rgb(49, 49, 49)

}

.container {
  overflow: auto;
  // background-image: url("./assets/wallhaven-p262de.jpg");
  border: 5px solid rgb(49, 49, 49);
  width: 80%;
  height: 80%;
  position: relative;
  background-size: cover;
  margin: 10vh auto;
  border-radius: 5px;
  background-color: rgb(255, 255, 255);
  // padding: 100px;
  //display: flex;
  //justify-content: space-between;
  //flex-direction: column;
  //align-items: center;
  padding: 20px;

  &::-webkit-scrollbar {
    width: 5px;
  }

  &::-webkit-scrollbar-thumb {
    width: 10px;
    background-color: rgba(73, 73, 73, .5);
    color: black;
    border-radius: 20px;
  }
}

.header {
  height: 5vh;
  margin: 20px 0;
  width: 100%;

  input {
    height: 100%;
    width: 80%;
    margin: 0 auto;
    display: block;
    text-align: center;
  }
}

.summary {
  width: 100%;
  margin: 10px auto;

  input {
    height: 50px;
    width: 80%;
    margin: 0 auto;
    display: block;
    text-align: center;

  }
}

.neck {
  min-height: 50px;
  display: block;
}

.editor {
  width: 100%;
}
</style>