<template>
  <div :class="{'hidden': hidden}" class="pagination-container">
    <el-pagination
      :background="background"
      :page-sizes="pageSizes"
      v-model:current-page="currentPage"
      v-model:page-size="pageSize"
      :layout="layout"
      :total="total"
      v-bind="$attrs"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
    />
  </div>
</template>

<script>
import { defineComponent, computed } from 'vue'
import { scrollTo } from '@/utils/components/scroll-to'

export default defineComponent({
  name: 'Pagination',
  props: {
    total: {
      required: true,
      type: Number
    },
    page: {
      type: Number,
      default: 1
    },
    limit: {
      type: Number,
      default: 20
    },
    pageSizes: {
      type: Array,
      default: () => {
        return [10, 20, 30, 50]
      }
    },
    layout: {
      type: String,
      default: 'total, sizes, prev, pager, next, jumper'
    },
    background: {
      type: Boolean,
      default: true
    },
    autoScroll: {
      type: Boolean,
      default: true
    },
    hidden: {
      type: Boolean,
      default: false
    }
  },
  setup(props, ctx) {

    const currentPage = computed({
      get: () => {
        return props.page
      },
      set: (val) => {
        ctx.emit('update:page', val)
      }
    })

    const pageSize = computed({
      get: () => {
        return props.limit
      },
      set: (val) => {
        ctx.emit('update:limit', val)
      }
    })

    function handleCurrentChange(val) {
      ctx.emit('pagination', { page: val, limit: pageSize })
      if (props.autoScroll) {
        scrollTo(0, 800, null)
      }
    }

    function handleSizeChange(val) {
      ctx.emit('pagination', { page: currentPage, limit: val })
      if (props.autoScroll) {
        scrollTo(0, 800, null)
      }
    }

    return {
      currentPage,
      pageSize,
      handleCurrentChange,
      handleSizeChange
    }
  }
})
</script>

<style scoped>
.pagination-container {
  padding: 32px 16px;
  background: var(--system-page-background);
  color: var(--system-page-color);
  border-color: var(--system-page-border-color);
}

.pagination-container.hidden {
  display: none;
}
</style>
