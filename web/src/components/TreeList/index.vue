<template>
  <div class="tree-list" @click="closeRightMenu">
    <div class="tree-title">
      <h3 class="title">{{ title }}</h3>
    </div>
    <div class="tree-box" @contextmenu.prevent>
      <el-input v-if="showSearch" v-model="filterText" :prefix-icon="useRenderIcon('search')" class="filter-input" />
      <div class="filter-tree" :class="{'has-input': showSearch}">
        <el-tree
          ref="tree"
          class="tree"
          :data="data"
          :node-key="nodeKey"
          :props="propsConf"
          :show-checkbox="showCheckBox"
          :check-strictly="!enableParentCheck"
          :check-on-click-node="enableClickCheck"
          :filter-node-method="filterNode"
          :highlight-current="true"
          @check-change="nodeCheck"
          @node-click="nodeClick"
          @node-contextmenu="rightClick"
        >
          <template #default="{ node }">
            <div class="custom-tree-node">
              <div class="custom-tree-node-icon" v-if="showIcon">
                <component
                  v-if="node.data[propsConf.children] && node.data[propsConf.children].length > 0"
                  :is="useRenderIcon(node.expanded ? 'folder-opened' : 'folder')"
                />
                <component
                  v-else-if="node.data.table === true || node.data.table === false"
                  :is="useRenderIcon(node.data.table ? 'table': 'view', 'svg')"
                />
                <component v-else :is="useRenderIcon('document')" />
              </div>
              <div class="custom-tree-node-label" :class="{'has-icon': showIcon}">
                <span>{{ node.data[propsConf.label] }}</span>
              </div>
            </div>
          </template>
        </el-tree>
      </div>
    </div>
    <div v-show="showRightMenu" class="right-menu" :style="{top:top+'px',left:left+'px'}" @click.stop>
      <el-button-group>
        <el-button class="menu" size="small" :icon="useRenderIcon('plus')" @click="appendNode">添加</el-button>
        <el-button class="menu" size="small" :icon="useRenderIcon('edit')" @click="updateNode">修改</el-button>
        <el-button class="menu" size="small" :icon="useRenderIcon('delete')" @click="deleteNode">删除</el-button>
      </el-button-group>
    </div>
  </div>
</template>

<script>
import { defineComponent, ref, watch } from 'vue'

export default defineComponent({
  name: 'TreeList',
  props: {
    title: {
      type: String,
      required: true
    },
    data: {
      type: Array,
      required: true
    },
    showSearch: {
      type: Boolean,
      default: true
    },
    showIcon: {
      type: Boolean,
      default: true
    },
    showCheckBox: {
      type: Boolean,
      default: false
    },
    enableParentCheck: {
      type: Boolean,
      default: false
    },
    enableClickCheck: {
      type: Boolean,
      default: false
    },
    disableRightClick: {
      type: Boolean,
      default: false
    },
    nodeKey: {
      type: String,
      default: 'id'
    },
    propsConf: {
      type: Object,
      default: () => {
        return {
          label: 'name',
          children: 'children'
        }
      }
    }
  },
  setup(props, ctx) {
    const tree = ref(null)
    const filterText = ref('')
    const filterData = ref([])
    const showRightMenu = ref(false)
    const top = ref(0)
    const left = ref(0)
    const clickedNode = ref(null)
    // 监听输入框值
    watch(filterText, (val) => {
      tree.value.filter(val)
    })

    function filterNode(value, data) {
      console.log(value, data)
      if (!value) return true
      if (data[props.propsConf.label]) {
        return data[props.propsConf.label].toLowerCase().indexOf(value.toLowerCase()) !== -1
      }
      return false
    }

    function nodeCheck(data, node, co) {
      console.log(data, node, co)
      clickedNode.value = data
      ctx.emit('node-check', data, node)
    }

    function nodeClick(data, node, co) {
      console.log(data, node, co)
      closeRightMenu()
      clickedNode.value = data
      // 不展示复选框时单选
      if (!props.showCheckBox) {
        setCheckedKeys([data[props.nodeKey]])
      }
      ctx.emit('node-click', data, node)
    }

    function rightClick(e, data, node, co) {
      if (props.disableRightClick) {
        return
      }
      console.log(data, node, co)
      top.value = e.clientY
      left.value = e.clientX
      showRightMenu.value = true
      clickedNode.value = data
      ctx.emit('right-click', data)
    }

    function closeRightMenu() {
      showRightMenu.value = false
    }

    function appendNode() {
      ctx.emit('append-node', clickedNode.value)
      closeRightMenu()
    }

    function updateNode() {
      ctx.emit('update-node', clickedNode.value)
      closeRightMenu()
    }

    function deleteNode() {
      ctx.emit('delete-node', clickedNode.value)
      closeRightMenu()
    }

    function getCheckedKeys() {
      return tree.value.getCheckedKeys()
    }

    function setCheckedKeys(keys) {
      tree.value.setCheckedKeys(keys)
    }

    function getHalfCheckedKeys() {
      return tree.value.getHalfCheckedKeys().concat(tree.value.getCheckedKeys())
    }

    function setHalfCheckedKeys(keys) {
      tree.value.setCheckedKeys([])
      keys.map(key => {
        const node = tree.value.getNode(key)
        if (node.isLeaf) {
          tree.value.setChecked(node, true)
        }
      })
    }

    return {
      tree,
      filterText,
      filterData,
      showRightMenu,
      top,
      left,
      filterNode,
      nodeCheck,
      nodeClick,
      rightClick,
      closeRightMenu,
      appendNode,
      updateNode,
      deleteNode,
      getCheckedKeys,
      setCheckedKeys,
      getHalfCheckedKeys,
      setHalfCheckedKeys
    }
  }
})
</script>

<style lang="scss">
.tree-list {
  width: calc(100% - 2px);
  height: calc(100% - 2px);
  border: 1px solid var(--system-page-border-color);
  border-radius: 4px;
  background-color: var(--system-page-background);
  color: var(--system-page-color);

  .tree-title {
    padding: 0 10px;
    width: calc(100% - 20px);
    height: 50px;
    line-height: 50px;
    border-bottom: 1px solid var(--system-page-border-color);

    .title {
      margin: 0;
      display: inline-block;
      float: left;
      height: 100%;
    }

    .title-btn {
      margin-top: 9px;
      display: inline-block;
      float: right;
      height: 100%;
    }
  }

  .tree-box {
    padding: 10px;
    width: calc(100% - 20px);
    height: calc(100% - 70px);

    .filter-input {
      width: 100%;
      height: 32px;

      .el-input__inner {
        //border-radius: 20px;
      }
    }

    .filter-tree {
      width: 100%;
      height: 100%;
      border-bottom-left-radius: 4px;
      border-bottom-right-radius: 4px;
      overflow: auto;
      .custom-tree-node {
        display: flex;
        height: 100%;
        line-height: 100%;
        text-align: left;
        &-icon {
          display: inline-block;
          height: 100%;
          line-height: 100%;
          .el-icon {
            display: inline-block;
            margin: 5px 0;
          }
        }
        &-label {
          display: inline-block;
          height: 100%;
          line-height: 100%;
          span {
            display: inline-block;
            margin: 5px;
          }
        }
      }
    }

    .filter-tree.has-input {
      height: calc(100% - 32px);
    }
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
}
</style>

<!--组件外的样式-->
<style lang="scss">

.el-tree {
  display: inline-block !important;
  min-width: 100%;
  height: fit-content;
}

.el-tree--highlight-current .el-tree-node.is-current>.el-tree-node__content:hover {
  background-color: var(--system-primary-color);
  color: var(--system-primary-text-color);
}

.el-tree--highlight-current .el-tree-node.is-current>.el-tree-node__content {
  background-color: var(--system-primary-color);
  color: var(--system-primary-text-color);
}
</style>
