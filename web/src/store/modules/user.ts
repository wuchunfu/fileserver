import { Module } from "vuex";
import { RootStateTypes, UserInfosState } from '/@/store/interface/index';
import { Session } from "/@/utils/storage";

const userInfosModule: Module<UserInfosState, RootStateTypes> = {
  namespaced: true,
  state() {
    return {
      userInfos: {},
    }
  },
  mutations: {
    // 设置用户信息
    getUserInfos(state: UserInfosState, payload: object) {
      state.userInfos = payload;
    },
  },
  actions: {
    // 设置用户信息
    async setUserInfos({ commit }, payload: object) {
      if (payload) {
        commit('getUserInfos', payload);
      } else {
        if (Session.get('userInfo')) commit('getUserInfos', Session.get('userInfo'));
      }
    },
  },
  getters: {
    getUserInfos(state: UserInfosState) {
      return state.userInfos;
    },
  }
}

export default userInfosModule
