<template>

    <div class="f-card" @click="open(link)">
      <div class="f-icon" :style="{backgroundImage:`url(${icon})`}">
      </div>
      <div class="f-right">
        <div class="f-name">{{ name }}</div>
        <div class="f-summary">{{ summary }}</div>
      </div>
    </div>
</template>

<script lang="ts" setup>


const props = defineProps<{
  name:string,
  summary:string,
  link:string,
  icon:string,
}>()

function open(link:string){
  window.open(link)
}
</script>

<style lang="less" scoped>


.f-container {
  display: flex;
  flex-direction: row;

  .f-card {
    user-select: none;
    cursor: pointer;
    display: flex;
    flex-direction: row;
    margin-bottom: 4vh;
    //margin-right: 12vh;
    padding:1vh;
    border-radius: .5vh;
    width: 50%;

    @media (max-aspect-ratio: 12/16){
      margin-right: 2vh;
      width: 90%;
    }

    &:hover{
      box-shadow:0 0 1vh 0vh rgba(0, 0, 0, 0.09);
    }

    .f-icon {
      border:.1vh solid #dadada;
      width: 8vh;
      height: 8vh;
      background-color: black;
      border-radius: 50%;
      margin-right: 1vh;
      background-position: center;
      background-size: cover;
      background-repeat: no-repeat;
    }

    .f-right {
      display: flex;
      flex-direction: column;

      .f-name {
        font-size: 1.1em;
        font-weight: bold;
        color: var(--button-color)
      }
      .f-summary{
        color: #625f5f;
        margin-top:1vh;
      }
    }
  }

}

@media screen and (max-aspect-ratio: 12/16) {
  width: 95vw;
  height: 50vh;
}

</style>