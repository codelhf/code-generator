<template>
  <div class="error-container flex center 403">
    <div class="img">
      <img src="@/assets/images/403.png" alt="403" />
    </div>
    <div class="txt">
      <h1 class="title">403</h1>
      <p class="message">抱歉，你无权访问该页面</p>
      <el-button type="primary" @click="goHome">返回首页</el-button>
    </div>
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import { useRouter } from 'vue-router'
export default defineComponent({
  name: '403',
  setup(props, ctx) {
    const router = useRouter()
    function goHome() {
      router.push('/')
    }
    return {
      goHome
    }
  }
})
</script>

<style lang="scss">
.error-container {
  width: 100%;
  height: 100%;
  line-height: 100%;
  .img {
    width: 251px;
    height: 294px;
  }
  .txt {
    padding-top: 115px;
    padding-left: 20px;
    width: 251px;
    height: 294px;
    text-align: left;
  }
}
</style>
