<template>
  <div class="system-table-box">
    <el-table
      ref="tableDom"
      v-bind="$attrs"
      class="system-table"
      :data="data"
      :loading="loading"
      highlight-current-row
      border
      height="100%"
    >
      <el-table-column type="selection" align="center" width="50" v-if="showSelection" />
      <el-table-column label="序号" width="60" align="center" v-if="showIndex">
        <template #default="scope">
          {{ (pageNum - 1) * pageSize + scope.$index + 1 }}
        </template>
      </el-table-column>
      <slot></slot>
    </el-table>
  </div>
  <div class="system-page-box">
    <el-pagination
      v-show="showPage"
      class="system-page"
      :background="pageBackground"
      :page-sizes="pageSizes"
      v-model:current-page="currentPage"
      v-model:page-size="pageLimit"
      :layout="pageLayout"
      :total="pageTotal"
      v-bind="$attrs"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
    />
  </div>
</template>

<script>
import { defineComponent, ref, onMounted, computed } from 'vue'
import { scrollTo } from '@/utils/components/scroll-to'

export default defineComponent({
  name: 'TableList',
  props: {
    data: { type: Array, default: () => [] }, // 数据源
    loading: { type: Boolean, default: false }, // 是否展示loading，默认否
    selected: { type: Array, default: () => [] }, // 已选择的数据，与selection结合使用
    showIndex: { type: Boolean, default: false }, // 是否展示index选择，默认否
    showSelection: { type: Boolean, default: false }, // 是否展示选择框，默认否
    showPage: { type: Boolean, default: true }, // 是否展示页级组件，默认是
    pageTotal: { required: true, type: Number },
    pageNum: { type: Number, default: 1 },
    pageSize: { type: Number, default: 20 },
    pageSizes: { type: Array, default: () => [10, 20, 50, 100] },
    pageLayout: { type: String, default: 'total, sizes, prev, pager, next, jumper' }, // 分页需要显示的东西，默认全部
    pageBackground: { type: Boolean, default: true },
    pageAutoScroll: { type: Boolean, default: true }
  },
  setup(props, ctx) {
    // 加载时
    const tableDom = ref(null)
    onMounted(() => {
      const selected = props.selected && props.selected.length > 0
      tableDom.value.toggleRowSelection(props.selected, selected)
    })
    // 获取选中行keys
    const getSelectedRowKeys = () => {
      const selectionRows = tableDom.value.getSelectionRows()
      return selectionRows.map(row => {
        return row.id
      })
    }
    // 分页部分
    const currentPage = computed({
      get: () => {
        return props.pageNum
      },
      set: (val) => {
        ctx.emit('update:pageNum', val)
      }
    })

    const pageLimit = computed({
      get: () => {
        return props.pageSize
      },
      set: (val) => {
        ctx.emit('update:pageSize', val)
      }
    })

    function handleCurrentChange(val) {
      ctx.emit('pagination', { pageNum: val, pageSize: pageLimit })
      if (props.autoScroll) {
        scrollTo(0, 800, null)
      }
    }

    function handleSizeChange(val) {
      ctx.emit('pagination', { pageNum: currentPage, pageSize: val })
      if (props.autoScroll) {
        scrollTo(0, 800, null)
      }
    }

    return {
      tableDom,
      getSelectedRowKeys,
      currentPage,
      pageLimit,
      handleCurrentChange,
      handleSizeChange
    }
  }
})
</script>

<style lang="scss" scoped>
.system-table-box {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: flex-start;
  width: 100%;
  height: calc(100% - 56px);
  background: var(--system-page-background);
  color: var(--system-page-color);
  border-color: var(--system-page-border-color);
}

.system-page-box {
  padding-top: 20px;
  background: var(--system-page-background);
  color: var(--system-page-color);
  border-color: var(--system-page-border-color);
}
</style>

<style lang="scss">
.el-table, .el-table th.el-table__cell, .el-table td.el-table__cell{
  background-color: var(--system-page-background);
  color: var(--system-page-color);
  border-color: var(--system-page-border-color);
}
</style>
