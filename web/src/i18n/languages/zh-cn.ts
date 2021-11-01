import zhLocale from "element-plus/lib/locale/lang/zh-cn";

export const lang = {
  el: zhLocale.el,
  language: '简体中文',
  router: {
    home: '首页',
    system: '系统设置',
    systemUser: '用户管理',
  },
  staticRoutes: {
    signIn: '登录',
    notFound: '找不到此页面',
    noPower: '没有权限',
  },
  noAccess: {
    accessTitle: '您未被授权，没有操作权限~',
    accessMsg: '联系方式：加QQ群探讨 665452019',
    accessBtn: '重新授权',
  },
  notFound: {
    foundTitle: '地址输入错误，请重新输入地址~',
    foundMsg: '您可以先检查网址，然后重新输入或给我们反馈问题。',
    foundBtn: '返回首页',
  },
  buttons: {
    changeLanguage: "切换语言"
  }
}
