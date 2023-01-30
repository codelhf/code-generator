<template>
  <div class="box">
    {{ title }}
  </div>
</template>

<script>
import { defineComponent, ref } from 'vue'
import { HomeApi } from '@/api/pages/dashboard'
export default defineComponent({
  name: 'Dashboard',
  setup() {
    const title = ref('')
    getTitle()
    function getTitle() {
      HomeApi({}).then(res => {
        title.value = res.data.title
      })
    }
    return {
      title
    }
  }
})
</script>

<style lang="scss" scoped>
  .box {
    padding: 15px;
  }
</style>
