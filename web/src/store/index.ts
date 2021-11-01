import { createStore, Store, useStore as baseUseStore } from "vuex"
import { InjectionKey } from "vue";
import { RootStateTypes } from "/@/store/interface";

export const key: InjectionKey<Store<RootStateTypes>> = Symbol();

export function useStore() {
  return baseUseStore(key);
}

// Vite supports importing multiple modules from the file system using the special import.meta.glob function
// see https://cn.vitejs.dev/guide/features.html#glob-import
const modulesFiles = import.meta.globEager('./modules/*.ts');
const pathList: string[] = [];

for (const path in modulesFiles) {
  pathList.push(path);
}

const modules = pathList.reduce((modules: { [x: string]: any }, modulePath: string) => {
  const moduleName = modulePath.replace(/^\.\/modules\/(.*)\.\w+$/, '$1');
  const value = modulesFiles[modulePath];
  modules[moduleName] = value.default;
  return modules;
}, {});

export const store = createStore<RootStateTypes>({ modules });

export default store;
