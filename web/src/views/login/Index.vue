<template>
  <div class="home-container">
    <div>{{ msg }}</div>
  </div>
</template>

<script lang="ts">

import { reactive, toRefs } from "vue";

export default {
  name: 'login',
  setup() {
    const state = reactive({
      msg: "Login Page"
    });
    return {
      ...toRefs(state),
    };
  },
};
</script>

<style lang="scss" scoped>
.home-container {
  overflow-x: hidden;
}

body > .el-container {
  margin-bottom: 40px;
}

.el-header,
.el-footer {
  background-color: #b3c0d1;
  color: var(--el-text-color-primary);
  text-align: center;
  line-height: 60px;
}

.el-main {
  background-color: #e9eef3;
  color: var(--el-text-color-primary);
  text-align: center;
  height: 100%;
  //line-height: 160px;
}
</style>
