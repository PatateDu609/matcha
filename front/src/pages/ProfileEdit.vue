<script lang="ts" setup>

import { ref } from 'vue'
import { useUserStore } from 'stores/user';

const userStore = useUserStore()

let firstNameText = userStore.firstName;
let lastNameText = userStore.lastName;
let bioText = userStore.bio;
let emailText = userStore.email;

</script>

<template>
  <template v-if="userStore.loading">
    uwu <!-- here use a loaded you prefer -->
  </template>

  <template v-else>
    <q-page class="row items-center justify-evenly">
      <q-card class="my-card">
        <q-card-section>
          <div class="row no-wrap items-center">
            <q-item>
                <q-input outlined v-model="firstNameText" label="FirstName" />
            </q-item>
            <q-item>
                <q-input outlined v-model="lastNameText" label="LastName" />
            </q-item>
          </div>
          <div class="row no-wrap items-center">
            <q-item>
                <q-input outlined v-model="emailText" label="Email" />
            </q-item>
            <q-item>
                <q-input outlined v-model="bioText" label="Bio" />
            </q-item>
          </div>
          <q-item>
            <q-btn color="green" label="Save" push href="http://localhost:9000/profile"/>
          </q-item>
        </q-card-section>
      </q-card>
    </q-page>
  </template>
</template>
