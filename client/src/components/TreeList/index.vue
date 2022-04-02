<template>
  <div class="tree-list" @contextmenu.prevent @click="closeRightMenu">
    <el-input v-if="showSearch" v-model="filterText" prefix-icon="el-icon-search" class="filter-input" />
    <div :class="{'filter-tree': true, 'has-input': showSearch}">
      <el-tree
        ref="tree"
        class="tree"
        :data="data"
        :node-key="nodeKey"
        :props="defaultProps"
        :show-checkbox="showCheckBox"
        :check-strictly="!enableParentCheck"
        :check-on-click-node="enableClickCheck"
        :filter-node-method="filterNode"
        :highlight-current="true"
        @check-change="nodeCheck"
        @node-click="nodeClick"
        @node-contextmenu="rightClick"
      >
        <span slot-scope="{ node }" class="custom-tree-node">
          <span v-if="showIcon">
            <i
              v-if="node.data.child && node.data.child.length > 0"
              :class="node.expanded ? 'el-icon-folder-opened' : 'el-icon-folder'"
            />
            <i
              v-else-if="node.data.children && node.data.children.length > 0"
              :class="node.expanded ? 'el-icon-folder-opened' : 'el-icon-folder'"
            />
            <svg-icon
              v-else-if="node.data.table === true || node.data.table === false"
              class="icon"
              :icon-class="node.data.table ? 'table': 'view'"
            />
            <i v-else class="el-icon-document" />
            {{ node.data.label ? node.data.label : node.data.name }}
          </span>
          <span v-else>
            {{ node.data.label ? node.data.label : node.data.name }}
          </span>
        </span>
      </el-tree>
    </div>
    <div v-show="showRightMenu" class="right-menu" :style="{top:top+'px',left:left+'px'}" @click.stop>
      <el-button-group>
        <el-button class="menu" size="mini" icon="el-icon-plus" @click="appendNode">
          {{ $t('components.treeList.addItem') || '添加' }}
        </el-button>
        <el-button class="menu" size="mini" icon="el-icon-edit" @click="updateNode">
          {{ $t('components.treeList.updateItem') || '修改' }}
        </el-button>
        <el-button class="menu" size="mini" icon="el-icon-delete" @click="deleteNode">
          {{ $t('components.treeList.deleteItem') || '删除' }}
        </el-button>
      </el-button-group>
    </div>
  </div>
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
    showIcon: {
      type: Boolean,
      required: false,
      default: true
    },
    showCheckBox: {
      type: Boolean,
      required: false,
      default: false
    },
    enableParentCheck: {
      type: Boolean,
      required: false,
      default: false
    },
    enableClickCheck: {
      type: Boolean,
      required: false,
      default: false
    },
    disableRightClick: {
      type: Boolean,
      required: false,
      default: false
    },
    disableParentNode: {
      type: Boolean,
      required: false,
      default: false
    },
    nodeKey: {
      type: String,
      required: false,
      default: 'id'
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
      clickedNode: null
    }
  },
  watch: {
    filterText(val) {
      this.$refs.tree.filter(val)
    }
  },
  created() {
    this.defaultProps.disabled = (data, node) => {
      if (this.disableParentNode && data.child) {
        return data.child.length > 0
      }
      if (this.disableParentNode && data.children) {
        return data.children.length > 0
      }
      return false
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
    nodeCheck(data, node, co) {
      console.log(data, node, co)
      this.clickedNode = data
      this.$emit('node-check', data, node)
    },
    nodeClick(data, node, co) {
      console.log(data, node, co)
      this.closeRightMenu()
      this.clickedNode = data
      this.$emit('node-click', data, node)
    },
    rightClick(e, data, node, co) {
      if (this.disableRightClick) {
        return
      }
      console.log(data, node, co)
      this.top = e.clientY
      this.left = e.clientX
      this.showRightMenu = true
      this.clickedNode = data
      this.$emit('right-click', data)
    },
    closeRightMenu() {
      this.showRightMenu = false
    },
    appendNode() {
      this.$emit('append-node', this.clickedNode)
      this.closeRightMenu()
    },
    updateNode() {
      this.$emit('update-node', this.clickedNode)
      this.closeRightMenu()
    },
    deleteNode() {
      this.$emit('delete-node', this.clickedNode)
      this.closeRightMenu()
    },
    getCheckedKeys() {
      return this.$refs.tree.getCheckedKeys()
    },
    setCheckedKeys(keys) {
      this.$refs.tree.setCheckedKeys(keys)
    },
    getCheckedNodes() {
      return this.$refs.tree.getCheckedNodes()
    },
    getHalfCheckedKeys() {
      return this.$refs.tree.getHalfCheckedKeys()
    }
  }
}
</script>

<style lang="scss">
.tree-list {
  padding: 10px;
  width: 100%;
  height: 100%;
  border: 1px solid #dddddd;
  border-radius: 4px;

  .filter-input {
    margin-bottom: 10px;

    .el-input__inner {
      //border-radius: 20px;
    }
  }

  .filter-tree {
    width: 100%;
    height: 100%;
    border: 1px solid #dddddd;
    border-radius: 4px;
    overflow: scroll;
  }

  .filter-tree.has-input {
    height: calc(100% - 46px);
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
</style>

<style lang="scss">
.tree {
  height: 100%;
}
.el-tree {
  min-width: 100%;
  display: inline-block !important;
}
.el-tree--highlight-current .el-tree-node .el-tree-node__content:hover {
  background-color: #dddddd;
}
.el-tree--highlight-current .el-tree-node.is-current .el-tree-node__content {
  background-color: #dddddd;
}
</style>
