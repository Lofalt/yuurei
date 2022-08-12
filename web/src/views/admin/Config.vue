<template>
  <n-table :single-line="false">
    <thead>
    <tr>
      <th>项</th>
      <th>值</th>
    </tr>
    </thead>
    <tbody>
    <tr>
      <td>用户名</td>
      <td>
        <n-input type="text" v-model:value="MyName"/>
      </td>
    </tr>
    <tr>
      <td>网站名</td>
      <td>
        <n-input type="text" v-model:value="WebsiteName"/>
      </td>
    </tr>
    <tr>
      <td>匿名名字</td>
      <td>
        <n-input type="text" v-model:value="AnonymousName"/>
      </td>
    </tr>
    <tr>
      <td>标语</td>
      <td>
        <n-input type="text" v-model:value="Slogan"/>
      </td>
    </tr>
    <tr>
      <td>留言框默认头像</td>
      <td>
        <upload-pic @confirm="acceptMessageIcon" name="上传头像" :raw-src="MessageIcon" directory="messageIcon"
                    :quality="40" :ratio="1"/>
      </td>
    </tr>
    <tr>
      <td>头像</td>
      <td>
        <upload-pic @confirm="acceptMyIcon" name="上传头像" :raw-src="MyIcon" directory="MyIcon" :quality="40"
                    :ratio="0.8"/>
      </td>
    </tr>
    <tr>
      <td>背景图</td>
      <td>
        <upload-pic @confirm="acceptBackground" name="上传背景" :raw-src="BackgroundImage" directory="background"
                    quality="40" :ratio="16/8"/>
      </td>
    </tr>
    <tr>
      <td>手机端</td>
      <td>
        <upload-pic @confirm="acceptMobileBackground" name="上传背景" :raw-src="MobileBackgroundImage"
                    directory="mobileBackground"
                    quality="40" :ratio="0.5"/>
      </td>
    </tr>
    <tr>
      <td>日期颜色</td>
      <td>
        <n-color-picker :modes="[`hex`]" v-model:value="DateColor"/>
      </td>
    </tr>
    <tr>
      <td>按钮颜色</td>
      <td>
        <n-color-picker :modes="[`hex`]" v-model:value="ButtonColor"/>
      </td>
    </tr>
    </tbody>

  </n-table>
  <n-button type="info" @click="edit" v-show="HasInited">修改</n-button>
  <n-button type="default" @click="create" v-show="!HasInited">初始化</n-button>
</template>

<script lang="ts" setup>
import {NTable, NInput, NColorPicker, NButton} from "naive-ui"
import UploadPic from "@/components/file/UploadPic.vue"
import {onMounted, ref} from "vue";
import axios from "@/request/index"
import {useRouter} from "vue-router";


const MyName = ref("")
const WebsiteName = ref("")
const MessageIcon = ref("")
const AnonymousName = ref("")
const MyIcon = ref('')
const BackgroundImage = ref('')
const MobileBackgroundImage = ref("")
const DateColor = ref('')
const ButtonColor = ref('')
const router = useRouter()
const Slogan = ref("")

const HasInited = ref(false)

function acceptMobileBackground(url: string) {
  MobileBackgroundImage.value = url
}

function acceptMessageIcon(url: string) {
  MessageIcon.value = url
}

function acceptMyIcon(url: string) {
  MyIcon.value = url
}

function acceptBackground(url: string) {
  BackgroundImage.value = url
}

getConfig()

function create() {
  axios.request("/yuurei/settings", "post", {
    MyName: MyName.value,
    WebsiteName: WebsiteName.value,
    MessageDefaultIcon: MessageIcon.value,
    Icon: MyIcon.value,
    BackgroundImage: BackgroundImage.value,
    DateColor: DateColor.value,
    ButtonColor: ButtonColor.value,
    AnonymousName: AnonymousName.value,
    Slogan:Slogan.value
  }).then((res: any) => {
    if (res.code == 200) {
      router.go(0)
      // router.replace("/home")
    } else {
      alert(res.msg)
    }
  })
}


function getConfig() {
  axios.get("/yuurei/settings/1", {}).then((res: any) => {
    if (res.code !== 200) {
      return
    }
    let set = res.data.data
    MyName.value = set.MyName
    WebsiteName.value = set.WebsiteName
    MessageIcon.value = set.MessageDefaultIcon
    MyIcon.value = set.Icon
    BackgroundImage.value = set.BackgroundImage
    DateColor.value = set.DateColor
    ButtonColor.value = set.ButtonColor
    MobileBackgroundImage.value = set.MobileBackgroundImage
    AnonymousName.value = set.AnonymousName
    Slogan.value = set.Slogan
    HasInited.value = true
  })
}

onMounted(() => {
  getConfig()
})

function edit() {
  axios.request("/yuurei/settings/1", "put", {
    MyName: MyName.value,
    WebsiteName: WebsiteName.value,
    MessageDefaultIcon: MessageIcon.value,
    Icon: MyIcon.value,
    BackgroundImage: BackgroundImage.value,
    DateColor: DateColor.value,
    ButtonColor: ButtonColor.value,
    MobileBackgroundImage: MobileBackgroundImage.value,
    AnonymousName: AnonymousName.value,
    Slogan:Slogan.value

  }).then((res: any) => {
    if (res.code == 200) {
      router.go(0)
      // router.replace("/home")
    } else {
      alert(res.msg)
    }
  })
}

</script>

<style lang="less" scoped>

</style>