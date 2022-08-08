<template>
  <div class="header">
    <n-button @click="write">写文章</n-button>
    <n-space style="width: 200px" vertical>
      <n-select v-model:value="category" :options="options" default-value="all"/>
    </n-space>
  </div>
  <n-data-table
      :columns="columns"
      :data="data"
      :pagination="pagination"
      :bordered="false"
  />
</template>

<script lang="ts" setup>
import {h, onMounted, provide, Ref, ref, watch} from 'vue'
import {NButton, NDataTable, NCard, NModal, NSpace, NSelect} from 'naive-ui'
import type {DataTableColumns} from 'naive-ui'
import axios from "@/request/index"
import {useDialog} from 'naive-ui'
import ArticleEditor from "@/components/edit/ArticleEditor.vue"
import ArticleUpdateEditor from "@/components/edit/ArticleUpdateEditor.vue"
import {useRouter} from "vue-router";

const router = useRouter()
const article = ref("")
provide('article', article)
const showModal = ref(false)
const showUpdateModal = ref(false)
// const message = useMessage()
const dialog = useDialog()
const options = ref([
  {
    label: "全部分类",
    value: "all"
  }
])
const category = ref('all')
type Article = {
  id: number
  title: string
  updatedAt: string
  createdAt: string
  category: string
  summary: string
}

function write() {
  router.push("/admin/create_entry")
}

function handleSuccess(arg: any) {
  showModal.value = false
  showUpdateModal.value = false
  getArticle()

}

function getCategory() {
  axios.get("/yuurei/entryCategory/all", {}).then((res: any) => {
    for (let i = 0; i < res.data.data.length; i++) {
      options.value.push({
        value: res.data.data[i].ID,
        label: res.data.data[i].CategoryName
      })
    }
  })
}


const actions = [{
  name: "删除",
  key: "delete",
  type:"error",
  call: function (id: any) {
    dialog.warning({
      title: '警告',
      content: '确定删除吗？',
      positiveText: '确定',
      negativeText: '算了',
      onPositiveClick: () => {
        axios.request("/yuurei/entry/" + id, "delete", {}).then((result: any) => {
          console.log(result)
        }).then(() => {
          getArticle()
        })
      },
      onNegativeClick: () => {
      }
    })
  }
}, {
  name: "修改",
  key: "update",
  type:"info",
  call: function (id: any) {
      router.push("/admin/editEntries/"+id)
  }
}]


const createColumns = ({play}: { play: (row: Article) => void }): DataTableColumns<Article> => {
  return [
    {
      title: 'id',
      key: 'id'
    },
    {
      title: 'Title',
      key: 'title'
    },
    {
      title: '分类',
      key: 'category'
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
        const action = actions.map((a:any) => {
          return h(
              NButton,
              {
                style:{
                  marginRight:`6px`
                },
                strong: true,
                tertiary: true,
                size: 'small',
                type:a.type,
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

const data: Ref<Article[]> = ref([])

function getArticle() {
  if (category.value === "all") {
    axios.get("/yuurei/entry/all", {}).then((result: any) => {
      // console.log(result.data)
      let dataTmp: Article[] = []
      result.data.data.forEach((art: any) => {
        dataTmp.push({
          id: art.ID,
          title: art.Title,
          updatedAt: new Date(art.UpdatedAt).toLocaleDateString() + " " + new Date(art.UpdatedAt).toLocaleTimeString(),
          createdAt: new Date(art.CreatedAt).toLocaleDateString() + " " + new Date(art.CreatedAt).toLocaleTimeString(),
          category: art.EntryCategory.CategoryName,
          summary: art.Summary
        })
      })
      data.value = dataTmp
    })
  } else {
    axios.get("/yuurei/entry/categoryId/" + category.value, {}).then((result: any) => {
      // console.log(result.data)
      let dataTmp: Article[] = []
      result.data.data.forEach((art: any) => {
        dataTmp.push({
          id: art.ID,
          title: art.Title,
          updatedAt: new Date(art.UpdatedAt).toLocaleDateString() + " " + new Date(art.UpdatedAt).toLocaleTimeString(),
          createdAt: new Date(art.CreatedAt).toLocaleDateString() + " " + new Date(art.CreatedAt).toLocaleTimeString(),
          category: art.EntryCategory.CategoryName,
          summary: art.Summary
        })
      })
      data.value = dataTmp
    })
  }
}

onMounted(() => {
  getArticle()
  getCategory()
})

const columns = createColumns({
  play(row: Article) {
  }
})
const pagination = false as const

watch(category, () => {
  getArticle()
})


</script>

<style lang="less" scoped>
.header {
  display: flex;
  justify-content: flex-start;
}

</style>