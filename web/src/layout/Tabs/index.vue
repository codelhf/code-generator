<template>
  <div class="tabs">
    <el-scrollbar
      class="scroll-container tags-view-container"
      ref="scrollbarDom"
      @wheel.passive="handleWillScroll"
      @scroll="handleScroll"
    >
      <Item
        v-for="menu in menuList"
        :key="menu.meta.title"
        :menu="menu"
        :active="activeMenu.path === menu.path"
        @close="delMenu(menu)"
        @reload="pageReload"
      />
    </el-scrollbar>
    <div class="handle">
      <el-dropdown placement="bottom">
        <div class="el-dropdown-link">
          <component :is="useRenderIcon('arrow-down', 'svg')" />
        </div>
        <template #dropdown>
          <el-dropdown-menu>
            <el-dropdown-item class="tab-ddropdown-item" @click="pageReload">
              <component :is="useRenderIcon('refresh-right', 'svg')" />{{ $t('tabs.refresh') }}
            </el-dropdown-item>
            <el-dropdown-item class="tab-ddropdown-item" :disabled="currentDisabled" @click="closeCurrentRoute">
              <component :is="useRenderIcon('close-circle', 'svg')" />{{ $t('tabs.close') }}
            </el-dropdown-item>
            <el-dropdown-item class="tab-ddropdown-item" :disabled="menuList.length < 3" @click="closeOtherRoute">
              <component :is="useRenderIcon('close-circle', 'svg')" />{{ $t('tabs.closeOthers') }}
            </el-dropdown-item>
            <el-dropdown-item class="tab-ddropdown-item" :disabled="menuList.length <= 1" @click="closeAllRoute">
              <component :is="useRenderIcon('close-circle', 'svg')" />{{ $t('tabs.closeAll') }}
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown>
      <el-tooltip class="item" effect="dark" :content="contentFullScreen ? '退出全屏':'内容全屏'" placement="bottom">
        <component :is="useRenderIcon(contentFullScreen ? 'fullscreen-exit' : 'fullscreen', 'svg')" @click="onFullscreen" />
      </el-tooltip>
    </div>
  </div>
</template>

<script>
import { defineComponent, computed, watch, reactive, ref, nextTick } from 'vue'
import { useStore } from 'vuex'
import { useRoute, useRouter } from 'vue-router'
import Item from './item.vue'
import tabsHook from './tabsHook'
export default defineComponent({
  components: {
    Item
  },
  setup() {
    const store = useStore()
    const route = useRoute()
    const router = useRouter()
    const scrollbarDom = ref(null)
    const scrollLeft = ref(0)
    const defaultMenu = {
      path: '/home',
      meta: { title: 'dashboard', hideClose: true }
    }
    const contentFullScreen = computed(() => store.state.app.contentFullScreen)
    const currentDisabled = computed(() => route.path === defaultMenu.path)

    let activeMenu = reactive({ path: '' })
    const menuList = ref(tabsHook.getItem())
    if (menuList.value.length === 0) { // 判断之前有没有调用过
      addMenu(defaultMenu)
    }
    watch(menuList.value, (newVal) => {
      tabsHook.setItem(newVal)
    })
    watch(menuList, (newVal) => {
      tabsHook.setItem(newVal)
    })
    router.afterEach(() => {
      addMenu(route)
      initMenu(route)
    })

    // 全屏
    function onFullscreen() {
      store.commit('app/contentFullScreenChange', !contentFullScreen.value)
    }
    // 当前页面组件重新加载
    function pageReload() {
      const self = route.matched[route.matched.length - 1].instances.default

      self.handleReload()
    }

    // 关闭当前标签，首页不关闭
    function closeCurrentRoute() {
      if (route.path !== defaultMenu.path) {
        delMenu(route)
      }
    }
    // 关闭除了当前标签之外的所有标签
    function closeOtherRoute() {
      menuList.value = [defaultMenu]
      if (route.path !== defaultMenu.path) {
        addMenu(route)
      }
      setKeepAliveData()
    }

    // 关闭所有的标签，除了首页
    function closeAllRoute() {
      menuList.value = [defaultMenu]
      setKeepAliveData()
      router.push(defaultMenu.path)
    }

    // 添加新的菜单项
    function addMenu(menu) {
      const { path, meta, name } = menu
      if (meta.hideTabs) {
        return
      }
      const hasMenu = menuList.value.some((obj) => {
        return obj.path === path
      })
      if (!hasMenu) {
        menuList.value.push({
          path,
          meta,
          name
        })
      }
    }

    // 删除菜单项
    function delMenu(menu) {
      let index = 0
      if (!menu.meta.hideClose) {
        if (menu.meta.cache && menu.name) {
          store.commit('keepAlive/delKeepAliveComponentsName', menu.name)
        }
        index = menuList.value.findIndex((item) => item.path === menu.path)
        menuList.value.splice(index, 1)
      }
      if (menu.path === activeMenu.path) {
        index - 1 > 0 ? router.push(menuList.value[index - 1].path) : router.push(defaultMenu.path)
      }
    }

    // 初始化activeMenu
    function initMenu(menu) {
      activeMenu = menu
      nextTick(() => {
        setPosition()
      })
    }
    // 设置当前滚动条应该在的位置
    function setPosition() {
      if (scrollbarDom.value) {
        const domBox = {
          scrollbar: scrollbarDom.value.wrap$.querySelector('.el-scrollbar__wrap '),
          activeDom: scrollbarDom.value.wrap$.querySelector('.active'),
          activeFather: scrollbarDom.value.wrap$.querySelector('.el-scrollbar__view')
        }
        let hasDoms = true
        Object.keys(domBox).forEach((dom) => {
          if (!dom) {
            hasDoms = false
          }
        })
        if (!hasDoms) {
          return
        }
        if (domBox.scrollbar && domBox.activeDom && domBox.activeFather) {
          const domData = {
            scrollbar: domBox.scrollbar.getBoundingClientRect(),
            activeDom: domBox.activeDom.getBoundingClientRect(),
            activeFather: domBox.activeFather.getBoundingClientRect()
          }
          domBox.scrollbar.scrollLeft = domData.activeDom.x - domData.activeFather.x + 1 / 2 * domData.activeDom.width - 1 / 2 * domData.scrollbar.width
        }
      }
    }

    // 配置需要缓存的数据
    function setKeepAliveData() {
      const keepAliveNames = []
      menuList.value.forEach((menu) => {
        menu.meta && menu.meta.cache && menu.name && keepAliveNames.push(menu.name)
      })
      store.commit('keepAlive/setKeepAliveComponentsName', keepAliveNames)
    }

    /** 监听鼠标滚动事件 */
    function handleWillScroll(e) {
      let distance = 0
      const speed = 5
      if (e.wheelDelta > 0) {
        distance = -10
      } else {
        distance = 10
      }
      // console.log(scrollLeft.value + eventDelta / 4)
      scrollbarDom.value?.setScrollLeft(scrollLeft.value + distance * speed)
    }

    /** 监听滚动事件 */
    function handleScroll({ scrollLeft: left }) {
      scrollLeft.value = left
    }

    // 初始化时调用：1. 新增菜单 2. 初始化activeMenu
    addMenu(route)
    initMenu(route)
    return {
      contentFullScreen,
      scrollbarDom,
      // 菜单相关
      menuList,
      activeMenu,
      currentDisabled,
      onFullscreen,
      pageReload,
      delMenu,
      closeCurrentRoute,
      closeOtherRoute,
      closeAllRoute,
      handleScroll,
      handleWillScroll
    }
  }
})
</script>

<style lang="scss" scoped>
  .tabs {
    display: flex;
    justify-content: space-between;
    align-items: center;
    height: 40px;
    background: var(--system-header-background);
    border-bottom: 1px solid var(--system-header-border-color);
    border-top: 1px solid var(--system-header-border-color);
    box-shadow: 0 1px 4px 0 rgba(0, 0, 0, .1);
    .handle {
      min-width: 95px;
      height: 100%;
      display: flex;
      align-items: center;
      .el-dropdown-link {
        margin-top: 5px;
        border-left: 1px solid var(--system-header-border-color);
        height: 25px;
        width: 40px;
        display: flex;
        justify-content: center;
        align-items: center;
      }
      i {
        color: var(--system-header-text-color);
      }
    }
  }
  .scroll-container {
    white-space: nowrap;
    position: relative;
    overflow: hidden;
    width: 100%;
    :deep {
      .el-scrollbar__bar {
        bottom: 0px;
      }
      .el-scrollbar__wrap {
        height: 49px;
      }
    }
  }
  .tags-view-container {
    height: 34px;
    flex: 1;
    width: 100%;
    display: flex;
  }
  .el-icon-full-screen {
    cursor: pointer;
    &:hover {
      background: rgba(0,0,0,.025);
    }
    &:focus {
      outline: none;
    }
  }
</style>
