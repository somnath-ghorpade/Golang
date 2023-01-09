import { createStore } from "vuex";

export default createStore({
  state: {
    count: 0,
    basicInfo: {}
  },

  mutations: {
    increment(state, payload) {
      return (state.count = state.count + payload);
    },
    decrement(state, payload) {
      return (state.count = state.count - payload);
    },
    mutateBasicInfo(state, infoData) {
      return (state.basicInfo = infoData);
    }
  },
  actions: {
    increment(context, payload) {
      context.commit("increment", payload);
    },
    decrement(context, payload) {
      context.commit("decrement", payload);
    },
    actionBasicInfo(store, actionData) {
      store.commit("mutateBasicInfo", actionData);
    }
  },
  getters: {},
  modules: {}
});
