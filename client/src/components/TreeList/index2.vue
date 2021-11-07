<template>
  <el-row>
    <el-col :span="4">
      <div class="tree-list" @contextmenu.prevent @click="closeRightMenu">
        <el-input v-if="showSearch" v-model="filterText" class="filter-input" />
        <el-card
          :class="{'filter-tree': true, 'has-input': showSearch, 'has-menu': showMenu}"
          shadow="never"
        >
          <div
            v-for="item in filterData"
            :key="item.name"
            :class="item.active ? 'item active' : 'item'"
            @click="itemClick(item)"
            @contextmenu="rightClick(item, $event)"
          >
            <svg-icon v-if="showIcon" class="icon" :icon-class="item.table ? 'table': 'view'" />
            <span class="name">{{ item.name }}</span>
          </div>
        </el-card>
        <div
          v-show="!showMenu && showRightMenu && clickedItem.rightMenu"
          class="right-menu"
          :style="{left:left+'px',top:top+'px'}"
          @click.stop
        >
          <el-button-group>
            <el-button class="menu" size="mini" icon="el-icon-plus" @click="addItem">
              {{ $t('components.treeList.addItem') }}
            </el-button>
            <el-button class="menu" size="mini" icon="el-icon-edit" @click="updateItem">
              {{ $t('components.treeList.updateItem') }}
            </el-button>
            <el-button class="menu" size="mini" icon="el-icon-delete" @click="deleteItem">
              {{ $t('components.treeList.deleteItem') }}
            </el-button>
          </el-button-group>
        </div>
        <div v-show="showMenu" class="bottom-menu">
          <el-button-group>
            <el-button class="menu" size="mini" icon="el-icon-plus" @click="addItem">
            </el-button>
            <el-button class="menu" size="mini" icon="el-icon-edit" @click="updateItem">
            </el-button>
            <el-button class="menu" size="mini" icon="el-icon-delete" @click="deleteItem">
            </el-button>
          </el-button-group>
        </div>
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
  name: 'TreeList',
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
    showMenu: {
      type: Boolean,
      required: false,
      default: false
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
      showRightMenu: false,
      top: 0,
      left: 0,
      clickedItem: ''
    }
  },
  watch: {
    data() {
      this.filterTable()
    },
    filterText(val) {
      this.filterTable(val)
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
      if (this.showMenu) return
      this.closeRightMenu()
      this.clickedItem = item
      this.$emit('itemClick', item)
    },
    rightClick(item, e) {
      if (this.showMenu) return
      this.top = e.clientY
      this.left = e.clientX
      this.clickedItem = item
      this.showRightMenu = true
      this.$emit('rightClick', item)
    },
    closeRightMenu() {
      this.showRightMenu = false
    },
    addItem() {
      this.$emit('addItem')
    },
    updateItem() {
      this.$emit('updateItem', this.clickedItem)
    },
    deleteItem() {
      this.$emit('deleteItem', this.clickedItem)
    }
  }
}
</script>

<style lang="scss">
  .tree-list {
    width: 100%;
    height: calc(100vh - 124px);
    border: 1px solid #dddddd;
    border-radius: 4px;
    .filter-tree {
      width: 100%;
      height: 100%;
      border: 1px solid #dddddd;
      border-radius: 4px;
      overflow: scroll;
      .el-card__body {
        padding: 0 20px;
      }
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
      }
      .item:hover {
        background-color: #eeeeee;
      }
      .item.active {
        background-color: #eeeeee;
      }
    }
    .filter-tree.has-input.has-menu {
      height: calc(100% - 68px);
    }
    .filter-tree.has-input {
      height: calc(100% - 40px);
    }
    .filter-tree.has-menu {
      height: calc(100% - 28px);
    }
    .right-menu {
      position: fixed;
      width: 80px;
      height: 90px;
      z-index: 999;
      .menu {
        border-radius: 0;
      }
    }
    .bottom-menu {
      text-align: center;
      .menu {
        width: 60px;
      }
    }
  }
  .tree-right {
    padding-left: 20px;
    height: calc(100vh - 124px);
  }
</style>
