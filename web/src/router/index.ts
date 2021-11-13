import { createRouter, createWebHashHistory } from "vue-router"
import NProgress from 'nprogress';
import 'nprogress/nprogress.css';
import { NextLoading } from '/@/utils/loading';
import { Session } from "/@/utils/storage";

const routes = [
  {
    path: '/',
    name: 'home',
    meta: {
      title: 'message.router.home',
    },
    component: () => import('/@/views/home/index.vue'),
  },
  {
    path: '/login',
    name: 'login',
    meta: {
      title: 'message.router.login',
    },
    component: () => import('/@/views/login/index.vue'),
  },
  {
    path: '/404',
    name: 'notFound',
    meta: {
      title: 'message.staticRoutes.notFound',
    },
    component: () => import('/@/views/error/404.vue'),
  },
  {
    path: '/401',
    name: 'noPower',
    meta: {
      title: 'message.staticRoutes.noPower',
    },
    component: () => import('/@/views/error/401.vue'),
  },
]

/**
 * 创建一个可以被 Vue 应用程序使用的路由实例
 * @method createRouter(options: RouterOptions): Router
 * @link 参考：https://next.router.vuejs.org/zh/api/#createrouter
 */
const router = createRouter({
  history: createWebHashHistory(),
  routes
})

// 路由加载前
router.beforeEach(async (to, from, next) => {
  // 存储 token 到浏览器缓存
  Session.set('token', Math.random().toString(36).substr(0));
  NProgress.configure({ showSpinner: false });
  if (to.meta.title) NProgress.start();
  const token = Session.get('token');
  if (to.path === '/login' && !token) {
    next();
    NProgress.done();
  } else {
    if (!token) {
      next(`/login?redirect=${ to.path }&params=${ JSON.stringify(to.query ? to.query : to.params) }`);
      Session.clear();
      NProgress.done();
    } else if (token && to.path === '/login') {
      next('/home');
      NProgress.done();
    } else {
      next();
    }
  }
});

// 路由加载后
router.afterEach(() => {
  NProgress.done();
  NextLoading.done();
});

export default router;
