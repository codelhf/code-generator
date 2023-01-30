<template>
  <template v-if="!menu.hidden">
    <router-link
      v-if="!menu.children"
      :key="menu.path"
      :to="menu.path"
      @click="hideMenu"
    >
      <el-menu-item :index="menu.name || menu.path">
        <template #title>
          <component v-if="menu.meta && menu.meta.icon" :is="useRenderIcon(menu.meta.icon, 'svg')" />
          <span>{{ generateTitle(menu.meta && menu.meta.title) }}</span>
        </template>
      </el-menu-item>
    </router-link>
    <router-link
      v-else-if="hasOneShowingChildren(menu.children) && !menu.children[0].children && !menu.alwaysShow"
      :key="menu.children[0].path"
      :to="menu.path === '/' ? '/' + menu.children[0].path : menu.path + '/' + menu.children[0].path"
      @click="hideMenu"
    >
      <el-menu-item :index="menu.path + '/' + menu.children[0].path">
        <template #title>
          <component v-if="menu.meta && menu.meta.icon" :is="useRenderIcon(menu.meta.icon, 'svg')" />
          <span>{{ generateTitle(menu.meta && menu.meta.title) }}</span>
        </template>
      </el-menu-item>
    </router-link>
    <el-sub-menu v-else :index="menu.name || menu.path">
      <template #title>
        <component v-if="menu.meta && menu.meta.icon" :is="useRenderIcon(menu.meta.icon, 'svg')" />
        <span>{{ generateTitle(menu.meta && menu.meta.title) }}</span>
      </template>
      <menu-item v-for="item in menu.children" :key="item.path" :menu="item" :is-nest="true" />
    </el-sub-menu>
  </template>
</template>

<script>
import { defineComponent, computed } from 'vue'
import { useStore } from 'vuex'
import { generateTitle } from '@/i18n'
export default defineComponent({
  name: 'MenuItem',
  props: {
    menu: {
      type: Object,
      required: true
    },
    isNest: {
      type: Boolean,
      default: false
    }
  },
  setup() {
    const store = useStore()
    const isCollapse = computed(() => store.state.app.isCollapse)
    const hideMenu = () => {
      if (document.body.clientWidth <= 1000 && !isCollapse.value) {
        store.commit('app/isCollapseChange', true)
      }
    }
    // 判断是否只显示一个子节点
    function hasOneShowingChildren(children) {
      if (!children) {
        return false
      }
      const showingChildren = children.filter(item => {
        return !item.hidden
      })
      return showingChildren.length === 1
    }
    return {
      hideMenu,
      generateTitle,
      hasOneShowingChildren
    }
  }
})
</script>

<style lang="scss" scoped>
  .el-menu-item,
  .el-sub-menu {
    text-align: left;
    width: 250px;
  }
  .collapse .el-menu-item,
  .collapse .el-sub-menu {
    text-align: left;
    width: 60px;
  }
  .el-menu-item svg, .el-sub-menu__title svg {
    padding-right: 4px;
  }
  .el-menu-item i, .el-sub-menu__title i {
    padding-right: 4px;
  }
</style>
