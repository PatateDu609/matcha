<template>
  <q-layout view="hHh LpR fff">
    <q-header reveal class="bg-primary text-white" height-hint="98">
      <q-toolbar>
        <router-link to="/">
          <q-avatar>
            <img alt="Logo" src="../assets/logo.svg" />
          </q-avatar>
        </router-link>

        <q-tabs align="left">
          <q-route-tab to="/browse"> 
            <q-icon name="groups" size="2rem" />
          </q-route-tab>
        </q-tabs>

        <q-space />
        <q-tabs align="right">
          <q-route-tab to="/chat"> 
            <q-icon name="question_answer" size="2rem" />
          </q-route-tab>
          <q-route-tab to="/profile"> 
            <q-icon name="account_circle" size="2rem" />
          </q-route-tab>
          <q-route-tab to="/notifications"> 
            <q-icon name="notifications" size="2rem" />
          </q-route-tab>
          <q-route-tab to="/sign-up" label="Sign up" />
          <q-route-tab to="/log-in" label="Log in" />
        </q-tabs>
      </q-toolbar>
    </q-header>

    <q-page-container>
      <router-view />
    </q-page-container>

    <q-footer class="row items-center justify-evenly">
      <a>Made by gboucett & rbourgea</a>
    </q-footer>
  </q-layout>
</template>

<script>

</script>
