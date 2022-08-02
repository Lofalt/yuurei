<template>
  <div class="container">
    <div class="inside">
      <n-tabs type="segment" animated style="margin: 0 auto">
        <n-tab-pane name="login" tab="登陆" display-directive="show">
          <login/>
        </n-tab-pane>
        <n-tab-pane name="signin" tab="注册" display-directive="show">
          <signin/>
        </n-tab-pane>

      </n-tabs>
    </div>
  </div>


</template>

<script lang="ts" setup>
import {NTabs, NTabPane} from "naive-ui"
import Login from "./Login.vue"
import Signin from "./Signup.vue"
</script>

<style lang="less" scoped>
.container {
  float: right;
  width: 83.5%;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;

  @media (max-width: 504px) {
    width: 100%;
  }

  .inside {
    //display: flex;
    //justify-content: center;
    //align-items: center;
    //width: 50%;
    background-color: #fff;
    padding: 10px;
    border-radius: 10px;
    height: 50%;
  }

}
</style>