<script lang="ts" setup>

import {ref} from 'vue';
import {api} from 'boot/axios';
import {useUserStore} from 'stores/user';
import { useRouter } from 'vue-router';

let username = ref('')
let password = ref('')
let email = ref('')
let firstName = ref('')
let lastName = ref('')

const userStore = useUserStore()

const router = useRouter()

function onSubmit() {
  let data = {
    'username': username.value,
    'password': password.value,
    'email': email.value,
    'first-name': firstName.value,
    'last-name': lastName.value,
  }

  type SignUpPayload = {
    id: string
  }

  console.log(data)

  api.post<SignUpPayload>('/sign-up', data)
    .then(response => {
      let userID = response.data.id

      console.log(response)
      console.log(userID)

      userStore.fetchUser(userID).then( () => router.push('/'))
    })
    .catch(reason => console.error(reason))
}
</script>

<template>
  <q-page class="row justify-center items-center" padding>
    <div class="row col justify-center">
      <div class="col" style="max-width: 400px">
        <q-form class="column q-gutter-md" @submit="onSubmit">
          <q-input
            v-model="username"
            label="Username"
            name="username"
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

          <q-input
            v-model="email"
            label="Email"
            name="email"
            outlined
            type="text"
          />

          <q-input
            v-model="firstName"
            label="First Name"
            name="first-name"
            outlined
            type="text"
          />

          <q-input
            v-model="lastName"
            label="Last Name"
            name="last-name"
            outlined
            type="text"
          />

          <q-btn color="secondary" type="submit">Sign up</q-btn>
        </q-form>
      </div>
    </div>
  </q-page>
</template>

<style scoped>
</style>
