<script setup lang="ts">
import { api } from 'boot/axios';
import { AxiosResponse } from 'axios';

interface Param {
  authuser: string;
  code: string;
  prompt: string;
  scope: string;
  state: string;
}

const props = defineProps<Param>();

console.log('Props are: ', props);

api
  .put<void, AxiosResponse<void>, Param>('/auth/google/redirect', props)
  .then(() => console.log('request ended successfully'))
  .catch((reason) => console.error(reason));
</script>

<template>
  <div></div>
</template>
