<template>
  <div style="padding: 10px;border: 1px solid #eee;border-radius: 4px;">
    <el-input v-model="filterText" @input="filterTable" />
    <el-card class="box-card" shadow="never">
      <div v-for="item in filterData" :key="item.name" :class="item.active ? 'item active' : 'item'" @click="clickItem(item)" @contextmenu="rightClick(item)">
        <svg-icon v-if="showIcon" class="icon" :icon-class="item.tableType === 1 ? 'table': 'view'" />
        <span class="name">{{ item.name }}</span>
      </div>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'DataSourceTable',
  props: {
    data: {
      type: Array,
      required: true
    },
    showIcon: {
      type: Boolean,
      required: false,
      default: false
    }
  },
  data() {
    return {
      filterText: '',
      filterData: []
    }
  },
  watch: {
    data: function() {
      this.filterTable()
    }
  },
  methods: {
    filterTable() {
      this.filterData = this.data.filter(item => {
        if (!this.filterText) return true
        if (this.filterText) {
          return item.name.toLowerCase().indexOf(this.filterText.toLowerCase()) > -1
        }
      })
    },
    clickItem(item) {
      this.$emit('item-click', item)
    },
    rightClick(item) {
      this.$emit('right-click', item)
    }
  }
}
</script>

<style>
  .box-card {
    width: 100%;
    height: 560px;
  }
  .box-card .el-card__body{
    padding: 0 20px;
    width: 100%;
    height: 100%;
    overflow: scroll;
  }
  .box-card .item {
    padding: 5px 0;
    width: max-content;
    min-width: 156px;
    height: 26px;
    line-height: 16px;
    font-size: 14px;
    cursor: pointer;
  }
  .box-card .item:hover {
    background-color: #eeeeee;
  }
  .box-card .item.active {
    background-color: #eeeeee;
  }
  .box-card .item .name {
    margin-left: 5px;
  }
</style>
