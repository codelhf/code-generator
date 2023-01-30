<template>
  <el-icon>
    <component :is="iconName" :class="iconClass" />
  </el-icon>
</template>

<script>
import { defineComponent } from 'vue'
// element-plus icon
// 在 main.js 做如下配置
// import * as icons from '@element-plus/icons-vue/dist/index'
// for (const [key, component] of Object.entries(icons)) {
//   app.component(key, component)
// }

export default defineComponent({
  name: 'IconOffline',
  props: {
    name: {
      type: String,
      required: true
    },
    className: {
      type: String,
      default: ''
    }
  },
  setup(props) {
    const iconName = `${props.name}`
    const iconClass = props.className ? `icon ${props.className}` : `${props.className}`
    return {
      iconName,
      iconClass
    }
  }
})
</script>

<style scoped>

</style>
