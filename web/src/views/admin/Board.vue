<template>
  <div class="board">
    <n-card title="后台管理" style="margin-bottom: 16px;min-height: 100vh;">
      <n-tabs type="line" animated default-value="article">
        <n-tab-pane name="article" tab="文章管理" display-directive="show">
          <n-dialog-provider>
            <article-menu/>
          </n-dialog-provider>
        </n-tab-pane>
        <n-tab-pane name="menu" tab="菜单管理" display-directive="show">
          <category-menu/>
        </n-tab-pane>
        <n-tab-pane name="pic" tab="图片管理 " display-directive="show">
          <pic-menu/>
        </n-tab-pane>
      </n-tabs>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
import {NCard, NTabs, NTabPane, NDialogProvider} from "naive-ui";
import articleMenu from "@/views/admin/ArticleMenu.vue";
import ArticleMenu from "@/views/admin/ArticleMenu.vue";
import CategoryMenu from "@/views/admin/CategoryMenu.vue"
import PicMenu from "@/views/admin/PicMenu.vue"
</script>

<style lang="less" scoped>

.board {
  width: 83.5%;
  height: 100vh;
  float: right;
  overflow: auto;

  @media (max-width: 800px) {
    width: 100%;
  }
}

</style>