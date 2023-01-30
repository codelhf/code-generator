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
      resolve()
    })
  }
}

export default {
  namespaced: true,
  state,
  actions,
  getters,
  mutations
}
