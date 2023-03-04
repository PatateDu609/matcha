<template>
  <q-btn icon="bi-google" id="google-btn" @click="onClick()" />
</template>

<script setup lang="ts">
import { api } from 'boot/axios';

async function onClick() {
  await api
    .get<string>('/auth/google')
    .then((resp) => (window.location.href = resp.data))
    .catch((reason) => console.error(reason));
}
</script>

<style scoped lang="scss">
@import "../../css/quasar.variables";

#google-btn {
  background-color: $google;
}
</style>
