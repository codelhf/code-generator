<template>
  <div class="code-container">
    <!--<el-input-->
    <!--  id="codeEditor"-->
    <!--  ref="codeEditor"-->
    <!--  type="textarea"-->
    <!--  v-model="code"-->
    <!--  :autosize="{ minRows: 12, maxRows: 20 }"-->
    <!--  :placeholder="placeholder"-->
    <!--/>-->
    <div id="codeEditor" ref="codeEditor"></div>
    <!--<el-select v-model="mode" class="code-mode-select" @change="changeMode">-->
    <!--  <el-option v-for="item in modes" :key="item.value" :label="item.label" :value="item.value" />-->
    <!--</el-select>-->
  </div>
</template>

<script>
export default {
  name: 'CodeEditor',
  props: {
    // 外部传入的内容，用于实现双向绑定
    value: {
      type: String,
      default: ''
    },
    // 外部传入的语法类型
    language: {
      type: String,
      default: null
    },
    placeholder: {
      type: String,
      required: false,
      default: ''
    }
  }
}
</script>

<script setup>
import { onMounted, ref, toRaw } from 'vue'
import * as monaco from "monaco-editor";
// 1. 引入monaco-editor中的language文件
import { language } from 'monaco-editor/esm/vs/basic-languages/handlebars/handlebars';

// 生命周期
onMounted(() => {
  _initialize()
})

// 默认的语法类型
const mode = ref('handlebars')
// 编辑器语法类型
const modes = ref([
  {
    value: 'handlebars',
    label: 'Handlebars'
  },
  {
    value: 'javascript',
    label: 'Javascript'
  }
])

const props = defineProps({
  // 外部传入的内容，用于实现双向绑定
  value: {
    type: String,
    default: ''
  },
  // 外部传入的语法类型
  language: {
    type: String,
    default: null
  },
  placeholder: {
    type: String,
    required: false,
    default: ''
  }
})
const emits = defineEmits(['input', 'language-change'])

// 组件内的代码值
const code = ref('')
// 编辑器实例
const codeEditor = ref(null)
// 编辑器配置
const option = ref({
  value: '',
  theme: 'vs-dark',
  language: mode.value,
  tabSize: 4, // 空格键大小
  folding: true, // 是否折叠
  foldingHighlight: true, // 折叠等高线
  foldingStrategy: "indentation", // 折叠方式  auto | indentation
  showFoldingControls: "always", // 是否一直显示折叠 always | mouseover
  disableLayerHinting: true, // 等宽优化
  emptySelectionClipboard: false, // 空选择剪切板
  selectionClipboard: false, // 选择剪切板
  automaticLayout: true, // 自动布局
  codeLens: false, // 代码镜头
  scrollBeyondLastLine: false, // 滚动完最后一行后再滚动一屏幕
  colorDecorators: true, // 颜色装饰器
  accessibilitySupport: "off", // 辅助功能支持  "auto" | "off" | "on"
  lineNumbers: "on", // 行号 取值： "on" | "off" | "relative" | "interval" | function
  lineNumbersMinChars: 5, // 行号最小字符   number
  enableSplitViewResizing: false,
  readOnly: false, //是否只读  取值 true | false
})
function _initialize() {
  // 初始化编辑器实例，传入需要被实例化的文本域对象和默认配置
  const container = document.getElementById('codeEditor')
  codeEditor.value = monaco.editor.create(container, option.value)
  // 代码提示
  monaco.languages.registerCompletionItemProvider(mode.value, {
    provideCompletionItems: function () {
      let suggestions = [];
      // 这个keywords就是python.js文件中有的
      language.keywords.forEach(item => {
        suggestions.push({
          label: item,
          kind: monaco.languages.CompletionItemKind.Keyword,
          insertText: item
        });
      })
      return {
        // 最后要返回一个数组
        suggestions
      };
    },
  })
  // 编辑器赋值
  toRaw(codeEditor.value).setValue(props.value || code.value)
  // 支持双向绑定
  codeEditor.value.onDidChangeModelContent(e => {
    const raw = toRaw(codeEditor.value)
    code.value = raw.textContent
    emits('input', code.value)
  })

  // 尝试从父容器获取语法类型
  if (props.language) {
    // 获取具体的语法类型对象
    const propsMode = _getLanguage(props.language)
    // 判断父容器传入的语法是否被支持
    if (propsMode) {
     mode.value = propsMode.label
    }
  }
}

// 获取当前语法类型
function _getLanguage(language) {
  // 在支持的语法类型列表中寻找传入的语法类型
  return modes.value.find((mode) => {
    // 所有的值都忽略大小写，方便比较
    const currentLanguage = language.toLowerCase()
    const currentLabel = mode.label.toLowerCase()
    const currentValue = mode.value.toLowerCase()
    // 由于真实值可能不规范，例如 java 的真实值是 x-java ，所以讲 value 和 label 同时和传入语法进行比较
    return currentLabel === currentLanguage || currentValue === currentLanguage
  })
}

// 自动格式化代码
function formatCode() {
  toRaw(codeEditor.value).trigger("anything", "editor.action.formatDocument");
  // 或者
  // this.codeEditor.getAction(['editor.action.formatDocument']).run()
}

// 更改模式
function changeMode(val) {
  // 修改编辑器的语法配置
  toRaw(codeEditor.value).setLanguageConfiguration(`${val}`)
  // 获取修改后的语法
  const label = _getLanguage(val).label.toLowerCase()
  // 允许父容器通过以下函数监听当前的语法值
  emits('language-change', label)
}
</script>

<style lang="scss" scoped>
.code-container {
  width: 100%;
  height: 50vh;
  #codeEditor {
    width: 100%;
    height: 100%;
  }
}
</style>
