import { loginApi, getInfoApi, logoutApi } from '@/api/user'
import { loginPhoneApi } from '@/api/user-phone'
import { loginQRApi } from '@/api/user-qrcode'

const state = () => ({
  token: '', // 登录token
  info: {} // 用户信息
})

// getters
const getters = {
  token(state) {
    return state.token
  },
  info(state) {
    return state.info
  }
}

// mutations
const mutations = {
  tokenChange(state, token) {
    state.token = token
  },
  infoChange(state, info) {
    state.info = info
  }
}

// actions
const actions = {
  // login by login.vue
  login({ commit, dispatch }, params) {
    return new Promise((resolve, reject) => {
      loginApi(params)
        .then(res => {
          commit('tokenChange', res.data.token)
          dispatch('getInfo', { token: res.data.token })
            .then(infoRes => {
              resolve(res.data.token)
            })
        })
    })
  },
  // get user info after user logined
  getInfo({ commit }, params) {
    return new Promise((resolve, reject) => {
      getInfoApi(params)
        .then(res => {
          commit('infoChange', res.data)
          resolve(res.data)
        })
    })
  },

  // login out the system after user click the loginOut button
  logout({ commit }) {
    logoutApi()
      .then(() => {

      })
      .catch(() => {

      })
      .finally(() => {
        localStorage.removeItem('tabs')
        localStorage.removeItem('vuex')
        location.reload()
      })
  },

  // login by login.vue
  loginByPhone({ commit, dispatch }, params) {
    return new Promise((resolve, reject) => {
      loginPhoneApi(params)
        .then(res => {
          commit('tokenChange', res.data.token)
          dispatch('getInfo', { token: res.data.token })
            .then(infoRes => {
              resolve(res.data.token)
            })
        })
    })
  },

  // login by login.vue
  loginByQRCode({ commit, dispatch }, params) {
    return new Promise((resolve, reject) => {
      loginQRApi(params)
        .then(res => {
          commit('tokenChange', res.data.token)
          dispatch('getInfo', { token: res.data.token })
            .then(infoRes => {
              resolve(res.data.token)
            })
        })
    })
  },
}

export default {
  namespaced: true,
  state,
  actions,
  getters,
  mutations
}
