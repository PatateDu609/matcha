<template>
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
               {{ data.username }}
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
          <template v-if="profileInfo">
            <q-rating v-model="data.score" color="red-7" icon="favorite_border" size="2em"/>
            <span class="text-caption text-grey q-ml-sm">{{ data.score }} (551)</span>
          </template>
          <template v-else>
            <q-rating v-model="userStore.fameRating" color="red-7" icon="favorite_border" size="2em"/>
            <span class="text-caption text-grey q-ml-sm">{{ userStore.fameRating }} (239)</span>
          </template>
        </div>
        <q-item>
          <div class="text-caption text-grey">
            <template v-if="profileOwner">
              {{ data.bio }}
            </template>
            <template v-else>
              {{ userStore.bio }}
            </template>
            <q-badge color="blue" label="#fun" outline style="margin-right:5px;"/>
            <q-badge color="secondary" label="#code" outline style="margin-right:5px;"/>
            <q-badge color="orange" label="#bubbler" outline style="margin-right:5px;"/>
          </div>
        </q-item>
        <q-item>
          <template v-if="profileOwner">
            <q-btn color="primary" label="Like" push style="margin:10px;"/>
            <q-btn color="red" label="Block" push style="margin:10px;">
              <q-tooltip>This user will no longer be able to talk to you.</q-tooltip>
            </q-btn>
            <q-btn color="orange" label="Report" push style="margin:10px;">
              <q-tooltip>This user will be reported to Bubbler staff</q-tooltip>
            </q-btn>
          </template>
          <template v-else>
            <q-btn color="grey" label="Edit Profile" push style="margin:10px;" href="http://localhost:9000/profile/edit">
              <q-tooltip>Edit your information here.</q-tooltip>
            </q-btn>
            <q-btn color="red" label="Likers" push style="margin:10px;">
              <q-tooltip>Check who liked you.</q-tooltip>
            </q-btn>
            <q-btn color="blue" label="Viewers" push style="margin:10px;">
              <q-tooltip>Check who visited your profile.</q-tooltip>
            </q-btn>
          </template>
        </q-item>
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { api } from '../boot/axios';
import { useUserStore } from 'stores/user';

let profilePicture = ref(1)

const userStore = useUserStore()

let uri = window.location.search.substring(1); 
let params = new URLSearchParams(uri);
let profileOwner = params.get("user");

export type UserPayload = {
  'first-name': string;
  'last-name': string;
  username: string;
  'full-name': string;
  email: string;
  bio: string | undefined | null;
  score: number;
};

let data = ref(null)

await new Promise((r) => setTimeout(r, 1000)) // define here the loading time

if (profileOwner != null) {
  try {
    let res = await api
      .get<UserPayload>(`/user/${profileOwner}`, {
        responseType: 'json',
      })
      .then((response) => {
        const d = response.data;
        return (d);
      })
      .catch((reason) => {
        console.error(reason);
      });
    console.log(res)
    data = res;
  } catch (e) {
    console.error(e)
  }
}
console.log(data)
</script>