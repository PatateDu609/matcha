<script lang="ts" setup>

import { ref } from 'vue';
import { api } from 'boot/axios';
import { CurrentUserPayload, useUserStore } from 'stores/user';
import { useRouter } from 'vue-router';

let identifier = ref('');
let password = ref('');

const router = useRouter();

const userStore = useUserStore();

function onSubmit() {
  let data = {
    'identifier': identifier.value,
    'password': password.value
  };

  console.log(data);

  api.post<CurrentUserPayload>('/log-in', data)
    .then(response => {
      userStore.fillWithPayload(response.data);

      router.push('/');
    })
    .catch(reason => console.error(reason));
}
</script>

<template>
  <q-page class="row justify-center items-center" padding>
    <div class="row col justify-center">
      <div class="col" style="max-width: 400px">
        <q-form class="column q-gutter-md" @submit="onSubmit">
          <q-input
            v-model="identifier"
            label="Username or Email"
            name="identifier"
            outlined
            type="text"
          />

          <q-input
            v-model="password"
            label="Password"
            name="password"
            outlined
            type="password"
          />

          <q-btn color="secondary" type="submit">Sign up</q-btn>
        </q-form>
      </div>
    </div>
  </q-page>
</template>

<style scoped>
</style>
