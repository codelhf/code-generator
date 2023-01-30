/** url链接正则 */
export function isUrl(value) {
  const reg = /(((^https?:(?:\/\/)?)(?:[-;:&=\+\$,\w]+@)?[A-Za-z0-9.-]+(?::\d+)?|(?:www.|[-;:&=\+\$,\w]+@)[A-Za-z0-9.-]+)((?:\/[\+~%\/.\w-_]*)?\??(?:[-\+=&;%@.\w_]*)#?(?:[\w]*))?)$/
  // @ts-expect-error
  return reg.test(value)
}

/** 手机号码正则 */
export function isPhone(value) {
  const reg = /^[1](([3][0-9])|([4][0,1,4-9])|([5][0-3,5-9])|([6][2,5,6,7])|([7][0-8])|([8][0-9])|([9][0-3,5-9]))[0-9]{8}$/
  // @ts-expect-error
  return reg.test(value)
}

/** 邮箱正则 */
export function isEmail(value) {
  const reg = /^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/
  // @ts-expect-error
  return reg.test(value)
}

/** 密码正则（密码格式应为8-18位数字、字母、符号的任意两种组合） */
export function isPassword(value) {
  const reg = /^(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9])(?=.*[^a-zA-Z0-9])\S{8,}$/
  // @ts-expect-error
  return reg.test(value)
}

/** 6位数字验证码正则 */
export function isValidSmsCode(value) {
  const reg = /^\d{6}$/
  // @ts-expect-error
  return reg.test(value)
}

/** 6位图片验证码正则 */
export function isValidImgCode(value) {
  const reg = /^[a-zA-Z0-9]{6}$/
  // @ts-expect-error
  return reg.test(value)
}

export function is(val, type) {
  return Object.prototype.toString.call(val) === `[object ${type}]`
}

export function isObject(val) {
  return val !== null && is(val, 'Object')
}

export function isString(val) {
  return val !== null && is(val, 'String')
}

export function isNumber(val) {
  return val !== null && is(val, 'Number')
}

export function isBoolean(val) {
  return val !== null && is(val, 'Boolean')
}

export function isArray(val) {
  return val && Array.isArray(val)
}

export function isRegExp(val) {
  return val !== null && is(val, 'RegExp')
}

export function isDate(val) {
  return val !== null && is(val, 'Date')
}

export function isFunction(val) {
  return typeof val === 'function'
}

export function isWindow(val) {
  return typeof window !== 'undefined' && is(val, 'Window')
}

export function isNull(val) {
  return val === null || val === undefined
}

export function isEmpty(val) {
  if (isArray(val) || isString(val)) {
    return val.length === 0
  }

  if (val instanceof Map || val instanceof Set) {
    return val.size === 0
  }

  if (isObject(val)) {
    return Object.keys(val).length === 0
  }

  return false
}

export function isPromise(val) {
  return (
    is(val, 'Promise') &&
    isObject(val) &&
    isFunction(val.then) &&
    isFunction(val.catch)
  )
}

