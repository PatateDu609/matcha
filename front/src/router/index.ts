import { route } from 'quasar/wrappers';
import {
  createMemoryHistory,
  createRouter,
  createWebHashHistory,
  createWebHistory,
  RouteLocationNormalized,
} from 'vue-router';

import routes from './routes';
import { useUserStore } from 'stores/user';
import { useRouterStore } from 'stores/router';
import { storeToRefs } from 'pinia';

/*
 * If not building with SSR mode, you can
 * directly export the Router instantiation;
 *
 * The function below can be async too; either use
 * async/await or return a Promise which resolves
 * with the Router instance.
 */

export default route(function (/* { store, ssrContext } */) {
  const createHistory = process.env.SERVER
    ? createMemoryHistory
    : process.env.VUE_ROUTER_MODE === 'history'
    ? createWebHistory
    : createWebHashHistory;

  const router = createRouter({
    scrollBehavior: () => ({ left: 0, top: 0 }),
    routes,

    // Leave this as is and make changes in quasar.conf.js instead!
    // quasar.conf.js -> build -> vueRouterMode
    // quasar.conf.js -> build -> publicPath
    history: createHistory(process.env.VUE_ROUTER_BASE),
  });

  router.beforeEach(async (to: RouteLocationNormalized) => {
    if (to.path == '/log-in' || to.path == '/sign-up') return true;
    let authorized = false;
    const userStore = useUserStore();

    if (userStore.isLogged) return true;

    await userStore
      .fetchCurrentUser()
      .then(() => {
        authorized = true;
      })
      .catch((reason) => {
        console.error(reason);
      });

    if (authorized) return true;

    const routerStore = useRouterStore();
    const { redirectURL } = storeToRefs(routerStore);
    redirectURL.value = to.path;

    return { path: '/log-in' };
  });

  return router;
});
