<template>
  <div class="header">
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

const article = ref("")
provide('article', article)
const showModal = ref(false)
const showUpdateModal = ref(false)
// const message = useMessage()
const dialog = useDialog()

type Visitor = {
  id: number
  IPAddress: string
  IPAddressLoc: string
  User: object
  FirstVisited: string
  LastVisited: string

}

function write() {
  showModal.value = !showModal.value

}

function handleSuccess(arg: any) {
  showModal.value = false
  showUpdateModal.value = false
  getVisitors()

}


const actions = [{
  name: "删除",
  key: "delete",
  type: "error",
  call: function (id: any) {
    dialog.warning({
      title: '警告',
      content: '确定删除吗？',
      positiveText: '确定',
      negativeText: '算了',
      onPositiveClick: () => {
        axios.request("/yuurei/article/" + id, "delete", {}).then((result: any) => {
          console.log(result)
        }).then(() => {
          getVisitors()
        })
      },
      onNegativeClick: () => {
      }
    })
  }
}, {
  name: "修改",
  key: "update",
  type: "info",
  call: function (id: any) {
    axios.get("/yuurei/article/" + id, {}).then((res: any) => {
      article.value = res.data.data
      showUpdateModal.value = !showUpdateModal.value
    })
  }
}]


const createColumns = ({play}: { play: (row: Visitor) => void }): DataTableColumns<Visitor> => {
  return [
    // {
    //   title: 'id',
    //   key: 'id'
    // },
    {
      title: 'IP地址',
      key: 'IPAddress'
    },
    {
      title: '用户名',
      key: 'User.Name'
    },
    {
      title: '初次访问',
      key: 'FirstVisited'
    },
    {
      title: "最后访问",
      key: "LastVisited",
    },
    {
      title: "归属地",
      key: "IPAddressLoc",
    },
    // {
    //   title: '操作',
    //   key: 'update',
    //   render(row) {
    //     const action = actions.map((a: any) => {
    //       return h(
    //           NButton,
    //           {
    //             style: {
    //               marginRight: `6px`
    //             },
    //             strong: true,
    //             tertiary: true,
    //             size: 'small',
    //             type: a.type,
    //             onClick: () => a.call(row.id)
    //           },
    //           {default: () => a.name}
    //       )
    //     })
    //     return action
    //   }
    // },
  ]
}

const data: Ref<Visitor[]> = ref([])

function getVisitors() {
  axios.get("/yuurei/visitors", {}).then((result: any) => {
    // console.log(result.data)
    let dataTmp: Visitor[] = []
    result.data.data.forEach((vis: any) => {
      dataTmp.push({
        id: vis.ID,
        IPAddress: vis.IPAddress,
        IPAddressLoc: vis.IPAddressLoc,
        User: vis.User,
        FirstVisited: new Date(vis.CreatedAt).toLocaleDateString() + " " + new Date(vis.CreatedAt).toLocaleTimeString(),
        LastVisited: new Date(vis.UpdatedAt).toLocaleDateString() + " " + new Date(vis.UpdatedAt).toLocaleTimeString(),
      })
    })
    data.value = dataTmp
  })
}

onMounted(() => {
  getVisitors()
})

const columns = createColumns({
  play(row: Visitor) {
  }
})
const pagination = false as const


</script>

<style lang="less" scoped>
.header {
  display: flex;
  justify-content: flex-start;
}

</style>