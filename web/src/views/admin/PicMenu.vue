<template>
  <div class="pic">
    <n-button @click="showModal=!showModal">上传文件</n-button>
    <n-modal v-model:show="showModal" display-directive="show">
      <n-card style="height:80vh;width:50vw">
        <div class="container">
          <div class="up">
            <img style="max-height:50vh;max-width: 30vw" :src="backgroundImg"/>
          </div>
          <input class="inputC" type="text" v-model="des" placeholder="输入图片描述:">
          <input type="file" accept="image/jpeg,image/png,image/bmp,image/gif" ref="inputFile" @change="getFile"
                 style="display: none">
          <n-button @click="upload">上传图片</n-button>
          <n-button @click="newCard">新增卡片</n-button>
        </div>
      </n-card>
    </n-modal>
    <div v-if="picList.value!=[]">
    <n-card  v-for="item in picList" :key="item.ID" style="width: 20vw;float:left">
      <template #cover>
        <img :src="item.Path" style="max-width: 20vw">
      </template>
      <div style="margin-top:20px">
      {{item.Description}}
      </div>
      <n-button type="warning" @click="deletePic(item.ID)" style="margin-top: 20px">删除</n-button>
      <n-button type="info" @click="updatePic(item.ID)" style="margin-left: 20px;margin-top:20px">修改描述</n-button>

    </n-card>
    </div>
    <n-modal v-model:show="showPic">
      <n-card style="width: 400px;height: 200px">
        <input @keyup.enter="send" style="height: 80%;width: 90%;display: block;margin: 0 auto;outline: none;border:1px solid rgb(49,49,49);text-align: center"  type="text" v-model="currentPic.Description">
        <n-button style="margin:10px" type="info" @click="send">提交</n-button>
      </n-card>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>

import {NButton, NModal, NCard} from 'naive-ui'
import axios from "@/request";
import {onMounted, ref} from "vue";

const inputFile = ref(null) as any
const backgroundImg = ref(null)
const showModal = ref(false)
const des =ref('')
const picList = ref([])
const currentPic = ref(null) as any
const showPic = ref(false)

function updatePic(id:number){

  picList.value.forEach((pic:any)=>{
    if(pic.ID===id){
      currentPic.value = pic
    }
  })
  showPic.value = !showPic.value
}

function send(){
  axios.request("/yuurei/gallery/"+currentPic.value.ID,"put",{
    Description:currentPic.value.Description
  }).then((res)=>{
    console.log(res)
    getPic()
    showPic.value = !showPic.value
  })
}

function deletePic(id:number){
  axios.request("/yuurei/gallery/"+id,"delete",{}).then((res)=>{
    console.log(res)
    getPic()
  })
}
onMounted(()=>{
  getPic()
})
function upload() {
  console.log(inputFile.value.click())
}

function getPic() {
  axios.get(`/yuurei/gallery/all`,{}).then((result:any) => {
    picList.value = result.data.data
  })
}
function newCard() {
  axios.request("/yuurei/gallery","post",{
    Description:des.value,
    Path:backgroundImg.value
  }).then((res)=>{
    console.log(res)
    showModal.value =!showModal.value
    getPic()
  })
}

function getFile(event: any) {
  let formFile = new FormData()
  formFile.append("file", event.target.files[0]);
  formFile.append("apply_info_id", RndNum(12));
  formFile.append("file_type", '');
  axios.file_upload("/yuurei/uploadGallery", formFile).then((res: any) => {
    console.log(res)
    backgroundImg.value = res.data.fileName
  })
}

function RndNum(num: number) {
  return Math.floor(Math.random() * Math.pow(10, num)).toString()
}
</script>

<style lang="less" scoped>

.container {
  display: flex;
  justify-content: space-evenly;
  flex-direction: column;

  .inputC {
    height: 10vh;
    padding:10px
  }

  .up {
    height: 50vh;
    width: 100%;
  }

}
</style>