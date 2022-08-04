<template>
  <div class="container">
    <img id="image" ref="image" :src="imgSrc" alt=""/>
  </div>
</template>

<script lang="ts" setup>
import 'cropperjs/dist/cropper.css'
import Cropper from "cropperjs"
import {onMounted, ref, watch} from "vue";
import {pProps} from "naive-ui";

const image = ref(null) as any
const emit = defineEmits(['send-detail'])
let cropper = null as any

const props = defineProps<{
  imgSrc: string,
  ratio: number,
  count: number
}>()
onMounted(() => {
  cropper = new Cropper(image.value, {
    aspectRatio: props.ratio,
    crop(event: any) {
      emit('send-detail', event.detail)
    },
    autoCrop: true,
    autoCropArea: 1,
    viewMode: 1,
  })
})
watch(props, (newValue: any, oldValue: any) => {
  cropper.replace(props.imgSrc)

})
</script>

<style lang="less" scoped>
.container {
  height: 50vh;
  max-width: 80%;

}


#image {
  display: block;

  /* This rule is very important, please don't ignore this */
  max-height: 100%;
  max-width: 80%;
}

</style>