<template>
  <el-config-provider
    :locale="locale"
    :size="size"
  >
    <router-view />
  </el-config-provider>
</template>

<script>
import { defineComponent, computed } from 'vue'
import { useStore } from 'vuex'
import { zhCn } from 'element-plus/lib/locales'
export default defineComponent({
  name: 'App',
  setup() {
    const store = useStore()
    const locale = zhCn
    const size = computed(() => store.state.app.elementSize)
    return {
      locale,
      size
    }
  }
})
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  width: 100%;
  height: 100vh;
}
</style>
