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
          <q-route-tab to="/chats" label="Chats" />
        </q-tabs>

        <q-space />

        <q-tabs align="right">
          <q-route-tab to="/sign-up" label="Sign up" />
          <q-route-tab to="/log-in" label="Log in" />
        </q-tabs>
      </q-toolbar>
    </q-header>

    <q-page-container>
      <router-view />
    </q-page-container>

    <q-footer class="bg-grey-8 text-white">
      <!-- We need to find what to put in the footer... -->
    </q-footer>
  </q-layout>
</template>

<script setup lang="ts">
</script>
