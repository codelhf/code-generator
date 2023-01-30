<template>
  <div class="layout-container">
    <div class="layout-container-tree">
      <div v-if="hasTree" class="tree-left">
        <slot name="tree"></slot>
      </div>
    </div>
    <div class="layout-container-right" :class="{ 'has-tree': hasTree }">
      <div class="layout-container-form" ref="formRef" :class="{ 'has-form': hasForm }">
        <el-row>
          <el-col :span="24">
            <slot name="form"></slot>
          </el-col>
        </el-row>
      </div>
      <div class="layout-container-table" ref="tableRef" :class="{ 'has-form': hasForm }" :style="formHeight">
        <el-row class="title">
          <el-col :span="12">
            <h3>{{ title }}</h3>
          </el-col>
          <el-col :span="12">
            <el-button-group>
              <slot name="button"></slot>
            </el-button-group>
          </el-col>
        </el-row>
        <el-row class="table">
          <el-col :span="24">
            <slot name="table"></slot>
          </el-col>
        </el-row>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent, useSlots, computed, ref, onMounted } from 'vue'

export default defineComponent({
  name: 'TableLayout',
  props: {
    title: {
      type: String,
      default: '列表'
    }
  },
  setup(props, ctx) {
    const slot = useSlots()
    const hasTree = computed(() => !!slot.tree)
    const hasForm = computed(() => !!slot.form)
    // 动态高度
    const formRef = ref(null)
    const formHeight = ref('--formHeight: 68px')
    onMounted(() => {
      formHeight.value = `--formHeight: ${formRef.value.offsetHeight + 48}px`
    })
    return {
      hasTree,
      hasForm,
      formRef,
      formHeight
    }
  }
})
</script>

<style lang="scss" scoped>
.layout-container {
  display: inline-block;
  width: 100%;
  height: 100%;

  &-tree {
    display: inline-block;
    float: left;
    height: 100%;

    .tree-left {
      width: 250px;
      height: calc(100vh - 130px);
      background-color: var(--system-container-main-background);
    }
  }

  &-right.has-tree {
    margin-left: 20px;
    width: calc(100% - 270px);
  }

  &-right {
    display: inline-block;
    float: left;
    width: 100%;
    height: 100%;

    .layout-container-form.has-form {
      padding-top: 18px;
      margin-bottom: 15px;
    }

    .layout-container-form {
      background-color: var(--system-container-main-background);
    }

    .layout-container-table.has-form {
      height: calc(100% - var(--formHeight));
    }

    .layout-container-table {
      padding: 15px;
      height: calc(100% - 30px);
      background-color: var(--system-container-main-background);

      .title {
        margin-bottom: 15px;

        h3 {
          margin: 0;
          text-align: left;
          height: 32px;
          line-height: 32px;
        }

        .el-button-group {
          float: right;
          height: 32px;
          line-height: 32px;
        }
      }

      .table {
        height: calc(100% - 47px);
        .el-col {
          height: 100%;
        }
      }
    }
  }
}
</style>

<style>

</style>
