import { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [{ path: '', component: () => import('pages/IndexPage.vue') }],
  },
  {
    path: '/log-in',
    component: () => import('layouts/MainLayout.vue'),
    children: [{ path: '', component: () => import('pages/LogInPage.vue') }],
  },
  {
    path: '/sign-up',
    component: () => import('layouts/MainLayout.vue'),
    children: [{ path: '', component: () => import('pages/SignUpPage.vue') }],
  },
  {
    path: '/browse',
    component: () => import('layouts/MainLayout.vue'),
    children: [{ path: '', component: () => import('pages/BrowseUsers.vue') }],
  },
  {
    path: '/profile',
    component: () => import('layouts/MainLayout.vue'),
    children: [{ path: '', component: () => import('pages/ProfilePage.vue') }],
  },
  {
    path: '/profile/edit',
    component: () => import('layouts/MainLayout.vue'),
    children: [{ path: '', component: () => import('pages/ProfileEdit.vue') }],
  },
  {
    path: '/validate/email',
    component: () => import('pages/VerifyPage.vue'),
  },
  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue'),
  },
];

export default routes;
