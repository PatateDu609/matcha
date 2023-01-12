<script setup lang="ts">

import {ref} from 'vue';
import {api} from 'boot/axios';

let username = ref('')
let password = ref('')
let email = ref('')
let firstName = ref('')
let lastName = ref('')

function onSubmit() {
  let data = {
    'username': username.value,
    'password': password.value,
    'email': email.value,
    'first-name': firstName.value,
    'last-name': lastName.value,
  }

  console.log(data)

  api.post('/sign-up', data)
    .then(value => console.log(value))
    .catch(reason => console.error(reason))
}
</script>

<template>
  <q-page class="row justify-center items-center" padding>
    <div class="row col justify-center">
      <div class="col" style="max-width: 400px">
        <q-form class="column q-gutter-md" @submit="onSubmit">
          <q-input
            outlined
            type="text"
            v-model="username"
            name="username"
            label="Username"
          />

          <q-input
            outlined
            type="password"
            v-model="password"
            name="password"
            label="Password"
          />

          <q-input
            outlined
            type="text"
            v-model="email"
            name="email"
            label="Email"
          />

          <q-input
            outlined
            type="text"
            v-model="firstName"
            name="first-name"
            label="First Name"
          />

          <q-input
            outlined
            type="text"
            v-model="lastName"
            name="last-name"
            label="Last Name"
          />

          <q-btn color="secondary" type="submit">Sign up</q-btn>
        </q-form>
      </div>
    </div>
  </q-page>
</template>

<style scoped>
</style>
