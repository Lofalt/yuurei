<template>
  <div class="container">
    <crop-img :count="count" :ratio="ratio" :imgSrc="src" @send-detail="acceptDetail"/>
    <!--    <img v-show="!cropping" :src="src" alt="sad"/>-->
    <div class="inside">
      <div class="btn" @click="upload">
        <loading-com v-show="isUploading"/>
        <span v-show="!isUploading">  {{ name }}</span>
      </div>
      <div class="btn" v-if="src" @click="cropPic">
        <loading-com v-show="isLoading"/>
        <span v-show="!isLoading">裁剪</span>
      </div>
      <div class="btn" v-if="src" @click="confirm">提交</div>
    </div>
  </div>
  <input type="file" accept="image/jpeg" ref="inputFile" @change="getFile"
         style="display: none">
  <input type="crop" accept="image/jpeg" ref="cropFile" @change="cropPic"
         style="display: none">
  <!--  </div>-->
</template>

<script lang="ts" setup>
import CropImg from "./CropImg.vue"
import {ref} from "vue";
import axios from "@/request";
import {NButton} from 'naive-ui'
import LoadingCom from "@/components/LoadingCom.vue";

const isUploading = ref(false)
const isLoading = ref(false)
const count = ref(0)
const cropping = ref(true)
const emit = defineEmits(['sendPic', 'confirm'])
const src = ref(null)
const props = defineProps<{
  quality: number,
  directory: string,
  ratio: number,
  name: string
}>()

const inputFile = ref(null) as any
const cropFile = ref(null) as any
const detail = ref(null) as any

function acceptDetail(details: any) {
  detail.value = details
}

function confirm() {
  isLoading.value = true
  axios.request("/yuurei/cropPic", "post", qs.stringify({
    imgPath: src.value,
    x: detail.value.x,
    y: detail.value.y,
    width: detail.value.width,
    height: detail.value.height
  })).then((res: any) => {
    console.log(res)
    if (res.code === 200) {
      src.value = res.data.data
      count.value += 1
      emit('confirm', src.value)
    }
    isLoading.value = false

  })
}

function crop() {
  cropFile.value.click()
}

function upload() {
  inputFile.value.click()
}

function getFile(event: any) {
  console.log("upload")
  isUploading.value = true
  src.value = "1"
  let formFile = new FormData()
  formFile.append("file", event.target.files[0]);
  formFile.append("apply_info_id", RndNum(12));
  formFile.append("file_type", '');
  axios.file_upload(`/yuurei/uploadImg?qua=${props.quality}&dir=${props.directory}`, formFile).then((res: any) => {
    src.value = res.data.fileName
    isUploading.value = false
  })
}

import qs from 'qs'

function cropPic(event: any) {
  isLoading.value = true
  axios.request("/yuurei/cropPic", "post", qs.stringify({
    imgPath: src.value,
    x: detail.value.x,
    y: detail.value.y,
    width: detail.value.width,
    height: detail.value.height
  })).then((res: any) => {
    if (res.code === 200) {
      src.value = res.data.data
      count.value += 1
    }
    isLoading.value = false
  })
}


function RndNum(num: number) {
  return Math.floor(Math.random() * Math.pow(10, num)).toString()
}
</script>

<style lang="less" scoped>
.container {
  margin: 0 auto;
  display: flex;
  //height: 50vh;
  max-width: 100%;
  flex-direction: row;
  justify-content: space-around;

  @media (max-width: 800px) {
    flex-direction: column;
  }
}

.inside {

  @media (max-width: 800px) {
    display: flex;
    justify-content: flex-start;

  }
}

.btn {
  position: relative;
  background-color: #fff;
  border: 3px solid rgb(49, 49, 49);
  //padding: 5px 10px;
  width: 20px;
  height: 35px;
  border-radius: 5px;
  cursor: pointer;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 80px;
  margin: 5px auto;
  text-align: center;


  &:hover {
    background-color: #ff4f4d;
    color: white;
  }
}

.loading {
  top: 40%;
  position: absolute;
  background-color: rgba(0, 0, 0, .2);
  @media (max-width: 800px) {
    left: 50%
  }
}
</style>