<template>
  <div class="logo-container">
    <img class="logo" src="@/assets/logo.png" alt="" />
    <h1 class="title" v-if="!isCollapse">{{ systemTitle }}</h1>
  </div>
</template>

<script>
import { defineComponent, computed } from 'vue'
import { useStore } from 'vuex'
import setting from '@/setting'
export default defineComponent({
  name: 'Logo',
  setup() {
    const store = useStore()
    const isCollapse = computed(() => store.state.app.isCollapse)
    const systemTitle = setting.systemTitle
    return {
      isCollapse,
      systemTitle
    }
  }
})
</script>

<style lang="scss" scoped>
  .logo-container {
    display: flex;
    width: 100%;
    height: 60px;
    align-items: center;
    //justify-content: center;
    padding-left: 10px;
    background-color: var(--system-logo-background);
    .logo {
      width: 40px;
      height: auto;
    }
    .title {
      margin-left: 18px;
      font-size: 18px;
      white-space: nowrap;
      color: var(--system-logo-color);
    }
  }
</style>
