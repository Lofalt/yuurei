<template>
  <n-button @click="showNew=!showNew">新增</n-button>
  <n-data-table
      :columns="columns"
      :data="data"
      :pagination="pagination"
      :bordered="false"
  />
  <n-modal v-model:show="showModal" display-directive="show">
    <div> 请输入 <input type="text" v-model="newName" @keyup.enter="send"/>
      <button @click="send">提交</button>
    </div>
  </n-modal>
  <n-modal v-model:show="showNew" display-directive="show">
    <div> 请输入 <input type="text" v-model="newCat" @keyup.enter="send2"/>
      <button @click="send2">提交</button>
    </div>
  </n-modal>
</template>

<script lang="ts" setup>
import {h, Ref, ref} from 'vue'
import {NButton, NDataTable, NModal} from 'naive-ui'
import type {DataTableColumns} from 'naive-ui'
import axios from "@/request/index"

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
    showModal.value = !showModal.value
    currentId.value = id
  }
}]

getCat()

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