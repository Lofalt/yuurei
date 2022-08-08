<template>
  <n-button @click="back">返回</n-button>
  <n-table :single-line="false">
    <thead>
    <tr>
      <th>项</th>
      <th>值</th>
    </tr>
    </thead>
    <tr>
      <td>标题</td>
      <td>
        <n-input type="text" v-model:value="Title"/>
      </td>
    </tr>
    <tr>
      <td>
        分类
      </td>
      <td>
        <n-space vertical>
          <n-select v-model:value="category" :options="options" placeholder="选择分类"/>
        </n-space>
      </td>
    </tr>
    <tr>
      <td>引用</td>
      <td>
        <n-input type="text" v-model:value="Quote"/>
      </td>
    </tr>
    <tr>
      <td>词条简介</td>
      <td>
        <n-input type="text" v-model:value="Summary"/>
      </td>
    </tr>

    <tr>
      <td>细节</td>
      <td>
        <wang-editor @upload-content="acceptDetail" :article="Detail"/>
      </td>
    </tr>
    <tr>
      <td>我的话</td>
      <td>
        <wang-editor @upload-content="acceptMyWord" :article="MyWord"/>
      </td>
    </tr>
    <tr>
      <td>相关影音</td>
      <td>
        <wang-editor @upload-content="acceptAbout" :article="About "/>
      </td>
    </tr>
    <tr>
      <td>头图</td>
      <td>
        <upload-pic name="上传头图" :raw-src="HeaderPicture" directory="" quality="" :ratio="0.6908" @confirm="acceptPic"/>
      </td>
    </tr>
    <tr>
      <td>信息
        <n-button @click="copy">+</n-button>
      </td>
      <td id="input">
        <div v-for="item in rawInfos" class="input">
          <button class="delBtn" @click="del">-</button>
          <input :value="item.ID" class="infos" style="display: none"/>
          <input :value="item.Label" @keyup.enter="copy" class="infos"/>
          <input :value="item.Value" @keyup.enter="copy" class="infos"/>
        </div>
        <div class="input">
          <button class="delBtn" @click="del">-</button>
          <input :value="0" class="infos" style="display: none"/>
          <input @keyup.enter="copy" class="infos"/>
          <input @keyup.enter="copy" class="infos"/>
        </div>
      </td>
    </tr>
    <tr>
      <td>主色</td>
      <td>
        <n-color-picker :modes="[`hex`]" v-model:value="MainColor"/>
      </td>
    </tr>
    <tr>
      <td>辅色</td>
      <td>
        <n-color-picker :modes="[`hex`]" v-model:value="SecColor"/>
      </td>
    </tr>
    <tr>
      <td>次色</td>
      <td>
        <n-color-picker :modes="[`hex`]" v-model:value="ThirdColor"/>
      </td>
    </tr>
  </n-table>
  <n-button @click="edit">修改</n-button>

</template>

<script lang="ts" setup>
import UploadPic from "../../components/file/UploadPic.vue"
import WangEditor from "../../components/wangeditor/EntryEditor.vue"
import {NColorPicker, NTable, NInput, NButton, NSpace, NSelect} from "naive-ui"
import {ref} from "vue";
import axios from "../../request/index"
import {useRouter} from "vue-router";

type info = {
  Label: string,
  Value: string,
}
const router = useRouter()
const Infos = ref([]) as any
const content = ref("") as any
const Title = ref("")
const Quote = ref("")
const Summary = ref("")
const Detail = ref("")
const MyWord = ref("")
const About = ref("")
const HeaderPicture = ref("")
const MainColor = ref("#000")
const SecColor = ref("#000")
const ThirdColor = ref("#000")
const options = ref([])
const category = ref(null)
const rawInfos = ref([])
const props = defineProps<{
  id:string
}>()

function back(){
  router.back()
}

axios.get("/yuurei/entryCategory/all", {}).then((res: any) => {
  for (let i = 0; i < res.data.data.length; i++) {
    options.value.push({
      value: res.data.data[i].ID,
      label: res.data.data[i].CategoryName
    })
  }
})

axios.get("/yuurei/entry/"+props.id,{}).then((res:any)=>{
  let Entry = res.data.data
  Title.value = Entry.Title
  Quote.value = Entry.Quote
  Detail.value = Entry.Detail
  Summary.value = Entry.Summary
  MyWord.value = Entry.MyWord
  About.value = Entry.About
  category.value = Entry.EntryCategoryID
  HeaderPicture.value = Entry.HeaderPicture
  MainColor.value = Entry.MainColor
  SecColor.value = Entry.SecColor
  ThirdColor.value = Entry.ThirdColor
  Title.value = Entry.Title
  rawInfos.value = Entry.Infos
})

function acceptDetail(as: string) {
  Detail.value = as
}

function acceptTitle(title: string) {
  Title.value = title
}

function acceptSummary(summary: string) {
  Summary.value = summary
}

function acceptAbout(about: string) {
  About.value = about
}

function acceptMyWord(myword: string) {
  MyWord.value = myword
}

function acceptPic(url: string) {
  HeaderPicture.value = url
}
function add(event){
  if(event.keyCode == 13){
    let node= document.getElementsByClassName("input")[0] as HTMLSelectElement
    let father = document.getElementById("input") as HTMLSelectElement
    let child = node.cloneNode(true)

    child.childNodes[0].addEventListener("click", del)
    child.childNodes[2].addEventListener("keyup", add)
    child.childNodes[3].addEventListener("keyup", add)

    child.childNodes[1].value=0
    child.childNodes[2].value = ""
    child.childNodes[3].value = ""
    // console.log(child.childNodes)
    father.appendChild(child)
  }
}

function copy() {
  let node = document.getElementsByClassName("input")[0] as HTMLSelectElement
  let father = document.getElementById("input") as HTMLSelectElement
  let child = node.cloneNode(true)
  child.childNodes[0].addEventListener("click", del)
  child.childNodes[2].addEventListener("keyup", add)
  child.childNodes[3].addEventListener("keyup", add)

  child.childNodes[1].value=0
  child.childNodes[2].value = ""
  child.childNodes[3].value = ""
  // console.log(child.childNodes)
  father.appendChild(child)
}

function del(event: any) {
  if(document.getElementsByClassName("input").length>1){
  let father = event.parentNode
  console.log(event)
  event.path[2].removeChild(event.path[1])

  }
}



function edit() {
  Infos.value =[]
  let infos = document.getElementsByClassName("infos") as HTMLCollection
  for (let i = 0; i < infos.length; i += 3) {
    if(infos[i+1].value!="") {
      Infos.value.push({
        ID: Number(infos[i].value),
        Label: infos[i + 1].value,
        Value: infos[i + 2].value
      })
    }
  }
  axios.request("/yuurei/entry/"+props.id, "put", {
    // ID:props.id,
    Title: Title.value,
    Quote: Quote.value,
    Summary: Summary.value,
    Detail: Detail.value,
    MyWord: MyWord.value,
    About: About.value,
    HeaderPicture: HeaderPicture.value,
    Infos: Infos.value,
    MainColor: MainColor.value,
    SecColor: SecColor.value,
    ThirdColor: ThirdColor.value,
    EntryCategoryID: category.value
  }).then((res:any)=>{
    if(res.code==200){
      router.push("/admin/entries")
    }
  })
}
</script>

<style lang="less" scoped>

.input {
  display: flex;
  justify-content: center;

  .infos {
    width: 50%;
    height: 5vh;
  }
}

.delBtn {
  width: 6vh;
  height: 5vh;
}
</style>