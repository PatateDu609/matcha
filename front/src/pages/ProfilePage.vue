<script lang="ts" setup>

import { ref } from 'vue'
import { useUserStore } from 'stores/user';
import AsyncProfileComponent from '../components/AsyncProfileComponent.vue'

const userStore = useUserStore()

let profilePicture = ref(1)

let uri = window.location.search.substring(1); 
let params = new URLSearchParams(uri);
let profileOwner = params.get("user");

</script>

<template>
  <template v-if="userStore.loading">
    uwu <!-- here use a loaded you prefer -->
  </template>

  <template v-else>
    <q-page class="row items-center justify-evenly">
      <q-card class="my-card">
        <q-carousel v-model="profilePicture" animated arrows infinite navigation>
          <q-carousel-slide :name="1" img-src="https://i.pinimg.com/originals/61/2f/d9/612fd974285f8057189b37657542e1ff.jpg"/>
          <q-carousel-slide :name="2" img-src="https://cdn.quasar.dev/img/parallax1.jpg"/>
          <q-carousel-slide :name="3" img-src="https://cdn.quasar.dev/img/parallax2.jpg"/>
          <q-carousel-slide :name="4" img-src="https://cdn.quasar.dev/img/quasar.jpg"/>
          <q-carousel-slide :name="5" img-src="https://cdn.quasar.dev/img/parallax1.jpg"/>
        </q-carousel>
        <q-card-section>
          <div class="row no-wrap items-center">
            <q-badge color="green" floating transparent>Online</q-badge>
            <div class="col text-h6 ellipsis">
              <template v-if="profileOwner">
                <Suspense>
                  <template #default>
                    <AsyncProfileComponent />
                  </template>
                  <template #fallback>
                    <q-skeleton :type="text" />
                  </template>
                </Suspense>
              </template>
              <template v-else>
                {{ userStore.fullName }}
              </template>
            </div>
            <div class="col-auto text-grey text-caption q-pt-md row no-wrap items-center">
              <q-icon name="place"/>
              250 Km
            </div>
          </div>
          <div class="row no-wrap items-center">
            <q-rating v-model="userStore.fameRating" color="red-7" icon="favorite_border" size="2em"/>
            <span class="text-caption text-grey q-ml-sm"><template v-if="profileInfo">{{ userStore.fameRating }} (551)</template></span>
          </div>
          <q-item>
            <div class="text-caption text-grey">
              <template v-if="profileInfo">
                {{ userStore.bio }}
              </template>
              <q-badge color="blue" label="#fun" outline style="margin-right:5px;"/>
              <q-badge color="secondary" label="#code" outline style="margin-right:5px;"/>
              <q-badge color="orange" label="#bubbler" outline style="margin-right:5px;"/>
            </div>
          </q-item>
          <q-item>
            <template v-if="!userStore.username">
              <q-btn color="primary" label="Like" push style="margin:10px;"/>
              <q-btn color="red" label="Block" push style="margin:10px;">
                <q-tooltip>This user will no longer be able to talk to you.</q-tooltip>
              </q-btn>
              <q-btn color="orange" label="Report" push style="margin:10px;">
                <q-tooltip>This user will be reported to Bubbler staff</q-tooltip>
              </q-btn>
            </template>
            <template v-if="userStore.username">
              <q-btn color="grey" label="Edit Profile" push style="margin:10px;" href="http://localhost:9000/profile/edit"/>
              <q-btn color="red" label="Likers" push style="margin:10px;"/>
              <q-btn color="blue" label="Viewers" push style="margin:10px;"/>
            </template>
          </q-item>
        </q-card-section>
      </q-card>
    </q-page>
  </template>
</template>
