import defaultSettings from '@/settings'

const { sidebarLogo, fixedHeader, showTagsView } = defaultSettings

const state = {
  sidebarLogo: sidebarLogo,
  fixedHeader: fixedHeader,
  showTagsView: showTagsView
}

const mutations = {
  CHANGE_SETTING: (state, { key, value }) => {
    if (state[key]) {
      state[key] = value
    }
  }
}

const actions = {
  changeSetting({ commit }, data) {
    commit('CHANGE_SETTING', data)
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

