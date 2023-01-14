<script lang="ts" setup>

import {api} from 'boot/axios';

type payload = {
  FirstName: string
  LastName: string
  Username: string
  FullName: string
  Email: string
  Biography: string
  FameRating: number
}

let data: payload;
let loading = true;

api.get<payload>('/user/780a6250-b31e-4eb9-9e12-6b80ba743df8')
  .then(value => {
    data = value;
    loading = false;
    console.log(data)
  })
  .catch(reason => console.error(reason))

</script>

<template>
  <template v-if="loading">
    uwu <!-- here use a loaded you prefer -->
  </template>

  <template v-else>
    <q-page class="row items-center justify-evenly">
      <q-card class="my-card">
        <q-carousel v-model="slide" animated arrows infinite navigation>
          <q-carousel-slide :name="1"
                            img-src="https://i.pinimg.com/originals/61/2f/d9/612fd974285f8057189b37657542e1ff.jpg"/>
          <q-carousel-slide :name="2" img-src="https://cdn.quasar.dev/img/parallax1.jpg"/>
          <q-carousel-slide :name="3" img-src="https://cdn.quasar.dev/img/parallax2.jpg"/>
          <q-carousel-slide :name="4" img-src="https://cdn.quasar.dev/img/quasar.jpg"/>
          <q-carousel-slide :name="5" img-src="https://cdn.quasar.dev/img/parallax1.jpg"/>
        </q-carousel>
        <q-card-section>
          <div class="row no-wrap items-center">
            <q-badge color="green" floating transparent>Online</q-badge>
            <div class="col text-h6 ellipsis">
              {{ data.fullname }}
            </div>
            <div class="col-auto text-grey text-caption q-pt-md row no-wrap items-center">
              <q-icon name="place"/>
              250 Km
            </div>
          </div>
          <div class="row no-wrap items-center">
            <q-rating v-model="ratingModel" color="red-7" icon="favorite_border" size="2em"/>
            <span class="text-caption text-grey q-ml-sm">3.2 (551)</span>
          </div>
          <q-item>
            <div class="text-caption text-grey">
              Small description about me
              <q-badge color="blue" label="#fun" outline/>
              <q-badge color="secondary" label="#code" outline/>
              <q-badge color="orange" label="#bubbler" outline/>
            </div>
          </q-item>
          <q-item>
            <q-btn color="primary" label="Like" push style="margin:10px;"/>
            <q-btn color="red" label="Block" push style="margin:10px;">
              <q-tooltip>This user will no longer be able to talk to you.</q-tooltip>
            </q-btn>
            <q-btn color="orange" label="Report" push style="margin:10px;">
              <q-tooltip>This user will be reported to Bubbler staff</q-tooltip>
            </q-btn>
          </q-item>
        </q-card-section>
      </q-card>
    </q-page>
  </template>
</template>
