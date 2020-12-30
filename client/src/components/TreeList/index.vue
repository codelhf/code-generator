<template>
  <el-row>
    <el-col :span="4">
      <div class="tree-list" @contextmenu.prevent @click="closeRightMenu">
        <el-input v-if="showSearch" v-model="filterText" @input="filterTable" />
        <el-card class="box-card" shadow="never">
          <div
            v-for="item in filterData"
            :key="item.name"
            :class="item.active ? 'item active' : 'item'"
            @click="itemClick(item)"
            @contextmenu="rightClick(item, $event)"
          >
            <svg-icon v-if="showIcon" class="icon" :icon-class="item.table ? 'table': 'view'" />
            <span class="name">{{ item.name }}</span>
            <div
              v-show="rightMenu && item.rightMenu && item.name === clickedItem"
              class="right-menu"
              :style="{left:left+'px',top:top+'px'}"
              @click.stop
            >
              <span class="menu plus" @click="addItem()">{{ $t('components.treeList.addItem') }}</span>
              <span class="menu edit" @click="updateItem(item)">{{ $t('components.treeList.updateItem') }}</span>
              <span class="menu delete" @click="deleteItem(item)">{{ $t('components.treeList.deleteItem') }}</span>
            </div>
          </div>
        </el-card>
      </div>
    </el-col>
    <el-col :span="20">
      <div class="tree-right">
        <slot />
      </div>
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: 'DataSourceTable',
  props: {
    data: {
      type: Array,
      required: true
    },
    showSearch: {
      type: Boolean,
      required: false,
      default: true
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
      filterData: [],
      rightMenu: false,
      top: 0,
      left: 0,
      clickedItem: ''
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
    itemClick(item) {
      this.closeRightMenu()
      this.$emit('itemClick', item)
    },
    rightClick(item, e) {
      this.top = e.clientY
      this.left = e.clientX
      this.clickedItem = item.name
      this.rightMenu = true
      this.$emit('rightClick', item)
    },
    closeRightMenu() {
      this.rightMenu = false
    },
    addItem() {
      this.$emit('addItem')
    },
    updateItem(item) {
      this.$emit('updateItem', item)
    },
    deleteItem(item) {
      this.$emit('deleteItem', item)
    }
  }
}
</script>

<style lang="scss">
  .tree-list {
    width: 100%;
    height: calc(100vh - 124px);
    border: 1px solid #eeeeee;
    border-radius: 4px;
    .el-card {
      width: 100%;
      height: 100%;
    }
    .el-input+.el-card{
      height: calc(100% - 40px);
    }
    .el-card__body{
      padding: 0 20px;
      width: 100%;
      height: 100%;
      overflow: scroll;
      .item {
        padding: 5px 0;
        width: max-content;
        min-width: 156px;
        height: 26px;
        line-height: 16px;
        font-size: 14px;
        cursor: pointer;
        .name {
          margin-left: 5px;
        }
        .right-menu {
          position: fixed;
          width: 80px;
          height: 92px;
          border: 1px solid #bfcbd9;
          z-index: 999;
          .menu {
            display: inline-block;
            padding: 0 10px;
            width: 100%;
            height: 30px;
            line-height: 30px;
            background-color: #ffffff;
          }
          .menu:hover {
            background-color: #eeeeee;
          }
        }
      }
      .item:hover {
        background-color: #eeeeee;
      }
      .item.active {
        background-color: #eeeeee;
      }
    }
  }
  .tree-right {
    padding-left: 20px;
    height: calc(100vh - 124px);
  }
</style>
