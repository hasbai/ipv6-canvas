import { InjectionKey } from "vue";
import { createStore, useStore as baseUseStore, Store } from "vuex";
import createPersistedState from "vuex-persistedstate";

export interface State {}

export const store = createStore<State>({
  plugins: [
    createPersistedState({
      paths: [],
    }),
  ],
  state: {},
  mutations: {},
});

export const key: InjectionKey<Store<State>> = Symbol();

export function useStore() {
  return baseUseStore(key);
}
