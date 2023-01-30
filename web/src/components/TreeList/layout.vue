<template>
  <div class="tree-left">
    <slot></slot>
  </div>
</template>

<script>
export default {
  name: 'TreeLayout'
}
</script>

<style scoped>
.tree-left {
  display: inline-block;
  float: left;
  width: 250px;
  height: calc(100vh - 130px);
  max-height: calc(100vh - 130px);
}
</style>

<style>
.tree-left + .layout-container {
  display: inline-block;
  float: left;
  width: calc(100% - 270px);
  margin-left: 20px;
}
</style>
