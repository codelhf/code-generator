<template>
  <div class="in-coder-panel">
    <el-input
      id="codeEditor"
      ref="codeEditor"
      v-model="code"
      type="textarea"
      :placeholder="placeholder"
    />
    <!--<el-select v-model="mode" class="code-mode-select" @change="changeMode">-->
    <!--  <el-option v-for="item in modes" :key="item.value" :label="item.label" :value="item.value" />-->
    <!--</el-select>-->
  </div>
</template>

<script>
import 'codemirror/lib/codemirror.css'
import 'codemirror/theme/mbo.css'
import 'codemirror/addon/hint/show-hint.css'
import CodeMirror from 'codemirror/lib/codemirror'
import 'codemirror/addon/edit/matchbrackets'
import 'codemirror/addon/selection/active-line'
import 'codemirror/addon/hint/show-hint'
import 'codemirror/mode/handlebars/handlebars'
import 'codemirror/mode/javascript/javascript'
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
  },
  data() {
    return {
      // 内部真实的内容
      code: '',
      // 默认的语法类型
      mode: 'x-javascript',
      modes: [
        {
          value: 'x-handlebars-template',
          label: 'Handlebars'
        },
        {
          value: 'javascript',
          label: 'Javascript'
        }
      ],
      // 编辑器实例
      coder: null,
      option: {
        // 缩进单位为4
        indentUnit: 4,
        // 使用tab键缩进
        indentWithTabs: true,
        // 智能缩进
        smartIndent: true,
        // 显示行数
        lineNumbers: true,
        // 括号匹配
        matchBrackets: true,
        // 自动换行
        lineWrapping: true,
        // 当前行背景高亮
        styleActiveLine: true,
        // 在选择时是否显示光标，默认为false
        showCursorWhenSelecting: true,
        // 编辑器主题
        theme: 'mbo',
        mode: 'text/x-handlebars-template',
        extraKeys: {
          // 智能提示
          'Alt-/': 'autocomplete',
          // Tab键换成4个空格
          'Tab': function(cm) {
            var spaces = Array(cm.getOption('indentUnit') + 1).join(' ')
            cm.replaceSelection(spaces)
          }
        }
      }
    }
  },
  mounted() {
    // 初始化
    console.log(this.value)
    this._initialize()
  },
  methods: {
    // 初始化
    _initialize() {
      // 初始化编辑器实例，传入需要被实例化的文本域对象和默认配置
      const textarea = document.getElementById('codeEditor')
      this.coder = CodeMirror.fromTextArea(textarea, this.option)
      // 编辑器赋值
      console.log(this.value)
      this.coder.setValue(this.value || this.code)

      // 支持双向绑定
      this.coder.on('change', (coder) => {
        this.code = coder.getValue()
        if (this.$emit) {
          this.$emit('input', this.code)
        }
      })

      // 尝试从父容器获取语法类型
      if (this.language) {
        // 获取具体的语法类型对象
        const mode = this._getLanguage(this.language)
        // 判断父容器传入的语法是否被支持
        if (mode) {
          this.mode = mode.label
        }
      }
    },
    // 获取当前语法类型
    _getLanguage(language) {
      // 在支持的语法类型列表中寻找传入的语法类型
      return this.modes.find((mode) => {
        // 所有的值都忽略大小写，方便比较
        const currentLanguage = language.toLowerCase()
        const currentLabel = mode.label.toLowerCase()
        const currentValue = mode.value.toLowerCase()
        // 由于真实值可能不规范，例如 java 的真实值是 x-java ，所以讲 value 和 label 同时和传入语法进行比较
        return currentLabel === currentLanguage || currentValue === currentLanguage
      })
    },
    // 更改模式
    changeMode(val) {
      // 修改编辑器的语法配置
      this.coder.setOption('mode', `text/${val}`)
      // 获取修改后的语法
      const label = this._getLanguage(val).label.toLowerCase()
      // 允许父容器通过以下函数监听当前的语法值
      this.$emit('language-change', label)
    }
  }
}
</script>

<style>
  .in-coder-panel{
    flex-grow: 1;
    display: flex;
    position: relative;
  }
  .in-coder-panel .CodeMirror{
    flex-grow: 1;
    z-index: 1;
    height: 400px;
  }
  .in-coder-panel .CodeMirror .CodeMirror-code{
    line-height: 20px;
  }
  .in-coder-panel .code-mode-select{
    position: absolute;
    z-index: 2;
    top: 10px;
    right: 10px;
    max-width: 130px;
  }
</style>
