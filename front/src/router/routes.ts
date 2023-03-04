import { RouteRecordRaw } from 'vue-router';
import GooglePage from 'pages/auth/GooglePage.vue';

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
    path: '/chat',
    component: () => import('layouts/MainLayout.vue'),
    children: [{ path: '', component: () => import('pages/ChatPage.vue') }],
  },
  {
    path: '/notifications',
    component: () => import('layouts/MainLayout.vue'),
    children: [{ path: '', component: () => import('pages/NotificationsPage.vue') }],
  },
  {
    path: '/validate/email',
    component: () => import('pages/VerifyPage.vue'),
  },
  {
    path: '/auth/redirect',
    component: () => import('layouts/AuthLayout.vue'),
    children: [
      {
        path: 'google',
        component: () => import('pages/auth/GooglePage.vue'),
        beforeEnter: (to) => {
          return (
            to.query.authuser != null &&
            to.query.code != null &&
            to.query.prompt != null &&
            to.query.scope != null &&
            to.query.state != null
          );
        },
        props: (to) => ({
          authuser: to.query.authuser,
          code: to.query.code,
          prompt: to.query.prompt,
          scope: to.query.scope,
          state: to.query.state,
        }),
      },
    ],
  },
  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue'),
  },
];

export default routes;
