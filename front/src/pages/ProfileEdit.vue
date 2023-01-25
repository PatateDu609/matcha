<script lang="ts" setup>

import { ref } from 'vue'
import { useUserStore } from 'stores/user';
import {api} from 'boot/axios';

const userStore = useUserStore()

let usernameText = ref(userStore.username);
let firstNameText = ref(userStore.firstName);
let lastNameText = ref(userStore.lastName);
let bioText = ref(userStore.bio);
let emailText = ref(userStore.email);
let genreSelect = ref("Homme");
let genreOptions = ref(["Homme", "Femme", "Autre"]);
let preferenceSelect = ref("Hétérosexuel");
let preferenceOptions = ref(["Hétérosexuel", "Homosexuel", "Bisexuel"]);
let geoCheck = ref(false);

function onSubmit() {
  let data = {
    'username': usernameText,
    'first-name': firstNameText,
    'last-name': lastNameText,
    'bio': bioText,
    'email': emailText,
    'genre': genreSelect,
    'preference': preferenceSelect,
    'geolocalisation': geoCheck,
  }

  type EditUserPayload = {
    id: string
  }

  api.patch<EditUserPayload>('/edit-user', data).then(response => {
    router.push('/profile')
  })
  .catch(reason => console.error(reason))
}

</script>

<template>
  <template v-if="userStore.loading">
    uwu <!-- here use a loaded you prefer ok -->
  </template>

  <template v-else>
    <q-page class="row items-center justify-evenly">
      <q-card class="my-card">
        <q-card-section>
          <q-form class="column q-gutter-md" @submit="onSubmit">
            <div class="row no-wrap items-center">
              <q-item>
                  <q-input outlined v-model="usernameText" label="Username *" />
              </q-item>
              <q-item>
                  <q-input outlined v-model="emailText" label="Email *" />
              </q-item>
            </div>
            <div class="row no-wrap items-center">
              <q-item>
                  <q-input outlined v-model="firstNameText" label="FirstName *" />
              </q-item>
              <q-item>
                  <q-input outlined v-model="lastNameText" label="LastName *" />
              </q-item>
            </div>
            <div class="row no-wrap items-center">
              <q-item>
                  <q-input outlined v-model="bioText" label="Bio" style="width:371px;" />
              </q-item>
            </div>
            <div class="row no-wrap items-center">
              <q-item>
                  <q-select outlined v-model="genreSelect" :options="genreOptions" label="Genre *" style="width:195px;" />
              </q-item>
              <q-item>
                  <q-select outlined v-model="preferenceSelect" :options="preferenceOptions" label="Preference *" />
              </q-item>
            </div>
            <div style="padding-left:40px;">
              <q-uploader url="http://localhost:4444/upload" label="Profile Picture upload" 
              multiple batch accept=".png, jpeg, .jpg, image/*" max-files="5"/>
            </div> 
            <!-- <q-checkbox v-model="geoCheck" label="Autoriser l'app à me géolocaliser ?" /> -->
            <q-item>
              <q-btn color="green" label="Save" push type="submit"/>
            </q-item>
          </q-form>
        </q-card-section>
      </q-card>
    </q-page>
  </template>
</template>
