<template>
  <el-row>
    <el-col :span="4">
      <div class="tree-list" @contextmenu.prevent @click="closeRightMenu">
        <el-input v-if="showSearch" v-model="filterText" class="filter-input" />
        <div :class="{'filter-tree': true, 'has-input': showSearch, 'has-menu': showMenu}">
          <el-tree
            ref="tree"
            class="tree"
            :data="data"
            :props="defaultProps"
            :show-checkbox="true"
            :check-strictly="true"
            :filter-node-method="filterNode"
            @node-click="nodeClick"
            @node-contextmenu="rightClick"
          />
        </div>
        <div
          v-show="!showMenu && showRightMenu"
          class="right-menu"
          :style="{top:top+'px',left:left+'px'}"
          @click.stop
        >
          <el-button-group>
            <el-button class="menu" size="mini" icon="el-icon-plus" @click="addItem">添加</el-button>
            <el-button class="menu" size="mini" icon="el-icon-edit" @click="updateItem">修改</el-button>
            <el-button class="menu" size="mini" icon="el-icon-delete" @click="deleteItem">删除</el-button>
          </el-button-group>
        </div>
        <div v-show="showMenu" class="bottom-menu">
          <el-button-group>
            <el-button class="menu" size="mini" icon="el-icon-plus" @click="addItem" />
            <el-button class="menu" size="mini" icon="el-icon-edit" @click="updateItem" />
            <el-button class="menu" size="mini" icon="el-icon-delete" @click="deleteItem" />
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
    },
    defaultProps: {
      type: Object,
      required: false,
      default: () => {
        return {
          label: (data, node) => {
            return data.label || data.name
          }
        }
      }
    }
  },
  data() {
    return {
      filterText: '',
      filterData: [],
      showRightMenu: false,
      top: 0,
      left: 0,
      clickedItem: null
    }
  },
  watch: {
    filterText(val) {
      this.$refs.tree.filter(val)
    }
  },
  methods: {
    filterNode(value, data) {
      console.log(value, data)
      if (!value) return true
      if (data.label) {
        return data.label.indexOf(value) !== -1
      }
      if (data.name) {
        return data.name.indexOf(value) !== -1
      }
      return false
    },
    nodeClick(data, node, co) {
      console.log(data, node, co)
      this.closeRightMenu()
      this.clickedItem = data
      this.$emit('node-click', data)
    },
    rightClick(e, data, node, co) {
      console.log(data, node, co)
      this.top = e.clientY
      this.left = e.clientX
      this.clickedItem = data
      this.showRightMenu = true
      this.$emit('right-click', data)
    },
    closeRightMenu() {
      this.showRightMenu = false
    },
    addItem() {
      if (this.showMenu) {
        const checkedNodes = this.$refs.tree.getCheckedNodes()
        if (checkedNodes.length !== 1) {
          this.$message.warning('请选择一个节点')
          return
        }
        this.clickedItem = checkedNodes[0]
      }
      this.$emit('add-item', this.clickedItem)
    },
    updateItem() {
      if (this.showMenu) {
        const checkedNodes = this.$refs.tree.getCheckedNodes()
        if (checkedNodes.length !== 1) {
          this.$message.warning('请选择一个节点')
          return
        }
        this.clickedItem = checkedNodes[0]
      }
      this.$emit('update-item', this.clickedItem)
    },
    deleteItem() {
      if (this.showMenu) {
        const checkedNodes = this.$refs.tree.getCheckedNodes()
        if (checkedNodes.length !== 1) {
          this.$message.warning('请选择一个节点')
          return
        }
        this.clickedItem = checkedNodes[0]
      }
      this.$emit('delete-item', this.clickedItem)
    }
  }
}
</script>

<style lang="scss">
.tree-list {
  padding: 10px;
  width: 100%;
  height: calc(100vh - 20px);
  border: 1px solid #dddddd;
  border-radius: 4px;
  .filter-tree {
    width: 100%;
    height: 100%;
    border: 1px solid #dddddd;
    border-radius: 4px;
    overflow: scroll;
  }
  .filter-tree.has-input.has-menu {
    height: calc(100% - 64px);
  }
  .filter-tree.has-input {
    height: calc(100% - 36px);
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
  }
}
.tree-right {
  padding-left: 20px;
  height: calc(100vh - 20px);
  overflow: auto;
}
</style>

<style>
.tree {
  height: 100%;
}
.el-tree {
  min-width: 100%;
  display:inline-block !important;
}
</style>
