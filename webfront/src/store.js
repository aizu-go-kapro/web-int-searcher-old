import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const Links = {
  namespaced: true,
  state: {
    keyword: '',
    links: []
  },
  mutations: {
    searchKeyword(state, keyword, links) {
      state.keyword = keyword
      state.links = links
    }
  },
  actions: {
    buttonAction({ commit }, keyword) {
      const params = encodeURIComponent(keyword).replace(/%20/g, '+')
      return fetch('http://localhost:3000/?q=' + params)
        .then(res => {
          const links = res.json()
          commit(keyword, links)
        })
    }
  },
  getters: {
    getKeyword(state) {
      return state.keyword
    },
    getlinks(state) {
      return state.links
    }
  }
}

export default new Vuex.Store({
  modules: {
    Links
  }
})
