import { RouteLocationNormalized } from 'vue-router';
import { useUserStore } from 'stores/user';
import { useRouterStore } from 'stores/router';
import { storeToRefs } from 'pinia';
import { boot } from 'quasar/wrappers';

export default boot(({ router, store }) => {
  router.beforeEach(async (to: RouteLocationNormalized) => {
    if (to.path == '/log-in' || to.path == '/sign-up') return true;
    let authorized = false;
    const userStore = useUserStore(store);

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

    const routerStore = useRouterStore(store);
    const { redirectURL } = storeToRefs(routerStore);
    redirectURL.value = to.path;

    return { path: '/log-in' };
  });
});
