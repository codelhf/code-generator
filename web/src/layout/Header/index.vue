<template>
  <header>
    <div class="left-box">
      <!-- 收缩按钮 -->
      <div class="menu-icon" @click="openedStateChange">
        <component :is="useRenderIcon(isCollapse ? 's-unfold' : 's-fold', 'svg')" />
      </div>
      <Breadcrumb />
    </div>
    <div class="right-box">
      <!-- 快捷功能按钮 -->
      <div class="function-list">
        <div class="function-list-item">
          <font-size />
        </div>
        <div class="function-list-item hidden-sm-and-down">
          <full-screen />
        </div>
        <div class="function-list-item hidden-sm-and-down">
          <theme />
        </div>
      </div>
      <!-- 用户信息 -->
      <div class="user-info-box">
        <el-dropdown>
          <div class="user-info">
            <img class="avatar" v-if="user.avatar" :src="user.avatar" alt="" />
            <span class="username">{{ user.nickName || user.username }}</span>
          </div>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item @click="userRoleChange(item.id)" v-for="(item, index) in user.roles" :key="index">
                <el-radio-group v-model="user.defaultRole">
                  <el-radio :label="item.id">{{ item.name }}</el-radio>
                </el-radio-group>
              </el-dropdown-item>
              <el-dropdown-item @click="toUserInfo" :icon="useRenderIcon('settings', 'svg')" divided>{{ $t('navbar.account') }}</el-dropdown-item>
              <el-dropdown-item @click="toPassword" :icon="useRenderIcon('edit', 'svg')">{{ $t('navbar.password') }}</el-dropdown-item>
              <el-dropdown-item @click="logout" :icon="useRenderIcon('quit', 'svg')" divided>{{ $t('navbar.logout') }}</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </div>
  </header>
</template>

<script>
import { defineComponent, computed } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import Breadcrumb from './Breadcrumb.vue'
import FontSize from './functionList/font-size.vue'
import FullScreen from './functionList/fullscreen.vue'
import Theme from './functionList/theme.vue'
export default defineComponent({
  name: 'Header',
  components: {
    Breadcrumb,
    FontSize,
    FullScreen,
    Theme
  },
  setup() {
    const store = useStore()
    const router = useRouter()

    // user info
    const user = computed(() => store.state.user.info) || {}

    const isCollapse = computed(() => store.state.app.isCollapse)
    // isCollapse change to hide/show the sidebar
    const openedStateChange = () => {
      store.commit('app/isCollapseChange', !isCollapse.value)
    }

    const userRoleChange = (id) => {
      store.dispatch('user/getInfo', { roleId: id })
        .then(res => {
          // router.push('/')
          location.reload()
        })
    }

    // to user info
    const toUserInfo = () => {
      router.push('/user_info')
    }

    // to password update
    const toPassword = () => {
      router.push('/update_password')
    }

    // logout the system
    const logout = () => {
      store.dispatch('user/logout')
    }

    return {
      user,
      isCollapse,
      openedStateChange,
      userRoleChange,
      toUserInfo,
      toPassword,
      logout
    }
  }
})
</script>

<style lang="scss" scoped>
  header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    height: 60px;
    background-color: var(--system-header-background);
    padding-right: 22px;
  }
  .left-box {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
    .menu-icon {
      width: 60px;
      height: 100%;
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 20px;
      font-weight: 100;
      cursor: pointer;
      margin-right: 10px;
      &:hover {
        background-color: var(--system-header-item-hover-color);
      }
      svg {
        color: var(--system-header-text-color);
      }
    }
  }
  .right-box {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
    .function-list{
      display: flex;
      height: 100%;
      .function-list-item {
        width: 30px;
        display: flex;
        justify-content: center;
        align-items: center;
        :deep(i) {
          color: var(--system-header-text-color);
        }
      }
    }
    .user-info-box {
      display: flex;
      height: 100%;
      margin-left: 20px;
      .user-info {
        display: flex;
        width: 100%;
        justify-content: center;
        align-items: center;
        .avatar {
          display: flex;
          width: 30px;
          height: 30px;
          border-radius: 30px;
          margin-right: 10px;
        }
        .username {
          display: flex;
          width: fit-content;
          color: var(--system-header-breadcrumb-text-color);
        }
      }
    }
  }
</style>
