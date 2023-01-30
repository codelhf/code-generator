<template>
  <div :title="isFullscreen ? $t('navbar.fullScreen.exit') : $t('navbar.fullScreen.full')">
    <component :is="useRenderIcon(isFullscreen ? 'fullscreen-exit' : 'fullscreen', 'svg')" @click="toggle" />
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import { useFullscreen } from '@vueuse/core'

export default defineComponent({
  name: 'Fullscreen',
  setup() {
    const { isFullscreen, toggle } = useFullscreen()
    return {
      isFullscreen,
      toggle
    }
  }
})
</script>

<style lang="scss" scoped>
  i {
    cursor: pointer;
    font-size: 18px;
    &:focus {
      outline: none;
    }
  }
</style>
