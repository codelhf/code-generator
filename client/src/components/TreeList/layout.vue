<template>
  <el-row>
    <el-col :span="leftSpan">
      <div class="tree-left">
        <slot name="left" />
      </div>
    </el-col>
    <el-col :span="rightSpan">
      <div class="tree-right">
        <slot name="right" />
      </div>
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: 'TreeLayout',
  props: {
    leftSpan: {
      type: Number,
      required: false,
      default: 4
    },
    rightSpan: {
      type: Number,
      required: false,
      default: 20
    }
  }
}
</script>

<style scoped>
.tree-left {
  width: 100%;
  height: calc(100vh - 124px);
  max-height: calc(100vh);
}
.tree-right {
  padding-left: 20px;
  width: 100%;
  height: calc(100vh - 124px);
  overflow-y: auto;
}
</style>
