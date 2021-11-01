// 接口类型声明

// 路由列表
export interface RoutesListState {
  routesList: Array<object>;
}

// 路由缓存列表
export interface KeepAliveNamesState {
  keepAliveNames: Array<string>;
}

// 用户信息
export interface UserInfosState {
  userInfos: object;
}

// 主接口(顶级类型声明)
export interface RootStateTypes {
  routesList: RoutesListState;
  keepAliveNames: KeepAliveNamesState;
  userInfos: UserInfosState;
}
