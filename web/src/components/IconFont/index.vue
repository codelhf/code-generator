<template>
  <svg v-if="type === 'svg'" :class="iconClass" aria-hidden="true">
    <use :xlink:href="iconName" />
  </svg>
  <i v-else :class="`${iconClass} ${iconName}`"></i>
</template>

<script>
import { defineComponent } from 'vue'
export default defineComponent({
  name: 'IconFont',
  props: {
    type: {
      type: String,
      default: 'svg'
    },
    name: {
      type: String,
      required: true
    },
    className: {
      type: String,
      default: ''
    }
  },
  setup(props) {
    let iconName = ''
    let iconClass = ''
    if (props.type === 'svg') {
      iconName = `#icon-${props.name}`
      iconClass = props.className ? `icon-svg ${props.className}` : 'icon-svg'
    } else {
      iconName = `${props.name}`
      iconClass = props.className ? `iconfont ${props.className}` : 'iconfont'
    }
    return {
      iconName,
      iconClass
    }
  }
})
</script>

<style scoped>
.icon-svg {
  width: 1em;
  height: 1em;
  fill: currentColor;
  vertical-align: middle;
  color: inherit;
  outline: none;
}
.iconfont {
  width: 1em;
  height: 1em;
  fill: currentColor;
  vertical-align: middle;
  color: inherit;
  outline: none;
}
</style>
