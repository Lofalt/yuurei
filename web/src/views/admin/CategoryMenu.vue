<template>
  <n-button @click="showNew=!showNew">新增</n-button>
  <n-data-table
      :columns="columns"
      :data="data"
      :pagination="pagination"
      :bordered="false"
  />
  <n-modal v-model:show="showModal" display-directive="show">
    <n-card style="width: 50vw">
      <n-input type="text" v-model:value="newName" @keyup.enter="send" placeholder="输入名称"/>
      <n-button @click="send">提交</n-button>
    </n-card>
  </n-modal>
  <n-modal v-model:show="showNew" display-directive="show">
    <n-card style="width: 50vw">
      <!--      <div :style="{backgroundImage:`url(${backgroundImg})`}" class="img">-->
      <!--      </div>-->
      <upload-pic :name="`上传图片`" :ratio="2.5" directory="headPic" quality="40" @send-pic="changeImg"/>
      <n-input type="text" v-model:value="newCat" @keyup.enter="send2" placeholder="输入名称"/>
      <n-button @click="send2">提交</n-button>
    </n-card>
  </n-modal>
</template>

<script lang="ts" setup>
import {h, Ref, ref} from 'vue'
import {NButton, NDataTable, NModal, NCard, NInput} from 'naive-ui'
import type {DataTableColumns} from 'naive-ui'
import axios from "@/request/index"
import UploadPic from "@/components/file/UploadPic.vue"

const backgroundImg = ref("")
const currentId = ref(0)
const newName = ref('')
type Category = {
  id: number
  title: string
  createdAt: string
  updatedAt: string
}
const showModal = ref(false)
const showNew = ref(false)
const newCat = ref('')
const actions = [{
  name: "删除",
  key: "delete",
  call: function (id: any) {
    axios.request("/yuurei/articleCategory/" + id, "delete", {}).then((res) => {
      getCat()
    })
  }
}, {
  name: "修改",
  key: "update",
  call: function (id: any) {
    axios.get(`/yuurei/articleCategory/${id}`, {}).then((res: any) => {
      currentId.value = id
      newName.value = res.data.data.ArticleCategoryName
      showModal.value = !showModal.value
    })
  }
}]

getCat()

function changeImg(pic: string) {
  backgroundImg.value = pic
}

function send2() {
  axios.request("/yuurei/articleCategory", "post", {
    ArticleCategoryName: newCat.value
  }).then((res) => {
    console.log(res)
    showNew.value = !showNew.value
    getCat()
  })
}

function send() {
  axios.request("/yuurei/articleCategory/" + currentId.value, "put", {
    ArticleCategoryName: newName.value
  }).then((res) => {
    console.log(res)
    showModal.value = !showModal.value
    getCat()
  })
}

const createColumns = ({play}: { play: (row: Song) => void }): DataTableColumns<Song> => {
  return [
    {
      title: 'ID',
      key: 'id'
    },
    {
      title: 'Title',
      key: 'title'
    },
    {
      title: '创建时间',
      key: 'createdAt'
    },
    {
      title: "更新时间",
      key: "updatedAt",
    },
    {
      title: '操作',
      key: 'update',
      render(row) {
        const action = actions.map((a) => {
          return h(
              NButton,
              {
                strong: true,
                tertiary: true,
                size: 'small',
                onClick: () => a.call(row.id)
              },
              {default: () => a.name}
          )
        })
        return action
      }
    },
  ]
}

const data: Ref<Category[]> = ref([])

function getCat() {
  let cats = []
  axios.get("/yuurei/articleCategory/all", {}).then((res: any) => {
    res.data.data.forEach((da: any) => {
      cats.push({
        id: da.ID,
        title: da.ArticleCategoryName,
        createdAt: new Date(da.CreatedAt).toLocaleDateString() + " " + new Date(da.CreatedAt).toLocaleTimeString(),
        updatedAt: new Date(da.UpdatedAt).toLocaleDateString() + " " + new Date(da.UpdatedAt).toLocaleTimeString(),
      })
    })
    data.value = cats
  })
}

const columns = createColumns({
  play(row: Category) {
  }
})
const pagination = false as const

</script>

<style lang="less" scoped>

.img {
  width: 100%;
  height: 20vh;
  background-color: #ffffff;
  background-size: cover;
  background-position: 50% 20%;
}

</style>