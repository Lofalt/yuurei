<template>
  <div class="emojiContainer">
    <!--    sadas-->
  </div>
</template>

<script lang="ts" setup>
import {onMounted} from "vue";

const emoji = ['😀', '😁', '😂', '🤣', '😃', '😄', '😅', '😆', '😉', '😊', '😋', '😎', '😍', '😘',
  '🥰', '😗', '😙', '😚', '🙂', '🤗', '🤩', '🤔', '🤨', '😐', '😑', '😶', '🙄', '😏', '😣', '😥', '😮', '🤐',
  '😯', '😪', '😫', '🥱', '😴', '😌', '😛', '😜', '😝', '🤤', '😒', '😓', '😔', '😕', '🙃', '🤑', '😲', '🙁', '😖',
  '😞', '😟', '😤', '😢', '😭', '😦', '😧', '😨', '😩', '🤯', '😬', '😰', '😱', '🥵', '🥶', '😳', '🤪', '😡', '😠', '🤬', '🤮', '❤', '💔', '💪', '🦵', '🦶',
  '👂', '👧', '👈', '👉', '☝', '👆', '👇', '✌', '🤞', '🖖', '🤘', '🤙', '🖐', '✋', '👌', '👍', '👎', '✊', '👊', '🤛', '👋', '🤟', '🤲', '👏', '🙏']

const emit = defineEmits(['emit-emoji'])

onMounted(() => {
  let container = document.getElementsByClassName("emojiContainer")[0]

  for (let i = 0; i < emoji.length; i++) {
    container.innerHTML = container.innerHTML + `<span class="emoji">${emoji[i]}</span>`
    // console.log(container.innerHTML)
  }

  let emojis = document.getElementsByClassName("emoji") as HTMLSelectElement

  for (let i = 0; i < emojis.length; i++) {
    emojis[i].addEventListener("click", () => {
      emit('emit-emoji', emoji[i])
    })
  }

})

</script>

<style lang="less">

.emojiContainer {
  border: 3px solid rgba(49, 49, 49, 1);
  width: 280px;
  background-color: #fff;
  border-radius: 10px;
  padding: 15px;
  height: 200px;
  display: flex;
  flex-wrap: wrap;
  overflow: auto;
  justify-content: flex-start;
  animation: enter .5s ease;

  &::-webkit-scrollbar {
    width: 5px;
    //right: 10px;
  }

  &::-webkit-scrollbar-thumb {
    width: 5px;
    border-radius: 10px;
    background-color: rgba(0, 0, 0, .1);
  }
}

.emoji {
  cursor: pointer;
  user-select: none;
  text-align: center;
  font-size: 1.2em;
  margin: 2px;
  width: 30px;
  height: 30px;
  padding: 1px;
}

@keyframes enter {
  from {
    opacity: 0;
  }
}
</style>