<template>
  <el-dropdown @command="handleCommand" size="default">
    <span class="el-dropdown-link">
      <component :is="useRenderIcon('font-size', 'svg')" />
    </span>
    <template #dropdown>
      <el-dropdown-menu>
        <el-dropdown-item
          v-for="d in list"
          :key="d.size"
          :command="d.size"
          :disabled=" elementSize === d.size "
        >
          {{ d.name }}
        </el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script>
import { defineComponent, computed, unref } from 'vue'
import { useRoute } from 'vue-router'
import { useStore } from 'vuex'
import { $t } from '@/i18n'
export default defineComponent({
  name: 'FontSize',
  setup() {
    const store = useStore()
    const route = useRoute()
    const elementSize = computed(() => store.state.app.elementSize)
    const list = [
      { size: 'large', name: $t('navbar.fontSize.large') },
      { size: 'default', name: $t('navbar.fontSize.default') },
      { size: 'small', name: $t('navbar.fontSize.small') }
    ]
    const { fullPath } = unref(route)
    function handleCommand(command) {
      store.commit('app/stateChange', {
        name: 'elementSize',
        value: command
      })
      // window.location.reload()
    }
    return {
      list,
      elementSize,
      fullPath,
      handleCommand
    }
  }
})
</script>

<style lang="scss" scoped>

</style>
