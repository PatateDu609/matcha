<script lang="ts" setup>
import { ref } from 'vue';
import { api } from 'boot/axios';
import { CurrentUserPayload, useUserStore } from 'stores/user';
import { useRouter } from 'vue-router';
import { useRouterStore } from 'stores/router';
import { storeToRefs } from 'pinia';
import SocialButtons from 'components/SocialButtons.vue';
import { ContextType } from 'components/models';

let identifier = ref('');
let password = ref('');

const router = useRouter();

const userStore = useUserStore();
const routerStore = useRouterStore();
const { redirectURL } = storeToRefs(routerStore);

function onSubmit() {
  let data = {
    identifier: identifier.value,
    password: password.value,
  };

  console.log(data);

  api
    .post<CurrentUserPayload>('/log-in', data)
    .then((response) => {
      userStore.fillWithPayload(response.data);

      if (redirectURL.value != undefined) router.push(redirectURL.value);
      else router.push('/');
    })
    .catch((reason) => console.error(reason));
}
</script>

<template>
  <q-page class="row justify-center items-center" padding>
    <div class="col">
      <div class="row justify-center">
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

            <q-btn color="secondary" type="submit">Log in</q-btn>
          </q-form>
        </div>
      </div>

      <q-separator spaced="6px" style="visibility: hidden" />

      <div class="row justify-center">
        <social-buttons :context-type="ContextType.LOG_IN" />
      </div>
    </div>
  </q-page>
</template>

<style scoped></style>
