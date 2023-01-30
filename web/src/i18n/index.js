import { createI18n } from 'vue-i18n'
import elementEnLocale from 'element-plus/lib/locale/lang/en' // element-ui lang
import elementZhLocale from 'element-plus/lib/locale/lang/zh-CN' // element-ui lang
import enLocale from './system/en'
import zhLocale from './system/zh'
import enPages from './language/en'
import zhPages from './language/zh'

const messages = {
  en: {
    ...elementEnLocale,
    ...enLocale,
    ...enPages
  },
  zh: {
    ...elementZhLocale,
    ...zhLocale,
    ...zhPages
  }
}

export function getLanguage() {
  const chooseLanguage = localStorage.getItem('language')
  if (chooseLanguage) {
    return chooseLanguage
  }
  // if has not choose language
  const language = (navigator.language || navigator.browserLanguage).toLowerCase()
  const locales = Object.keys(messages)
  for (const locale of locales) {
    if (language.indexOf(locale) > -1) {
      return locale
    }
  }
  return 'en'
}

// 重新封装方法
export function $t(args) {
  return i18n.global.t(args) || args
}

// translate router.meta.title, be used in breadcrumb sidebar tags view
export function generateTitle(title) {
  const hasKey = i18n.global.te('route.' + title)
  if (hasKey) {
    // $t :this method from vue-i18n, inject in @/lang/index.js
    return i18n.global.t('route.' + title)
  }
  return title
}

const i18n = createI18n({
  // legacy: false, // 使用CompetitionAPI必须添加这条.
  locale: getLanguage(), // 首选语言
  // fallbackLocale: 'en', // 备选语言
  // globalInjection: true, // 加上这一句
  // 将其设为true,则所有前缀为$的属性和方法(只是vue-i18n所携带的)将被注入到所有vue组件中.
  // 即在所有组件中都可以使用 $i18n $t $rt $d $n $tm
  // set locale messages
  messages
})

export default i18n
