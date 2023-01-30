import { defineComponent, h } from 'vue'
import IconFont from '@/components/IconFont/index.vue'
import IconOffline from '@/components/IconOffline/index.vue'

/**
 * 支持fontawesome4、5+、iconfont、自定义svg、element-plus的icons
 * @param icon     必传 string   图标
 * @param iconType 可选 string   图标类型
 * @param attrs    可选 object   属性
 * @returns Component
 */
export default function useRenderIcon(icon, iconType, attrs) {
  return defineComponent({
    name: 'Icon',
    render() {
      if (iconType) {
        return h(IconFont, {
          name: icon,
          type: iconType,
          ...attrs
        })
      } else {
        return h(IconOffline, {
          name: icon,
          ...attrs
        })
      }
    }
  })
}
