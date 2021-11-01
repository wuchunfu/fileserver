<template>
  <el-config-provider :locale="locale">
    <router-view/>
  </el-config-provider>
</template>

<script lang="ts">
import { computed, defineComponent } from "vue";
import { i18n } from "./i18n";

export default defineComponent({
  components: {},
  setup() {
    let locale = computed(() => (i18n.global.messages as any)![i18n.global.locale]);
    return {
      locale,
    }
  },
})

</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
