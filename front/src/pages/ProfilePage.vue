<script lang="ts" setup>

import { ref } from 'vue'
import { useUserStore } from 'stores/user';
import AsyncProfileComponent from '../components/AsyncProfileComponent.vue'

const userStore = useUserStore()
</script>

<template>
  <template v-if="userStore.loading">
    uwu <!-- here use a loaded you prefer -->
  </template>

  <template v-else>
    <Suspense>
      <template #default>
        <AsyncProfileComponent />
      </template>
      <template #fallback>
        <!-- <q-page class="row items-center justify-evenly">
          <q-card class="my-card">
            <q-skeleton width="400px" height="400px" />
            <q-card-section>
              <div class="row no-wrap items-center">
                <div class="col text-h6 ellipsis">
                  <q-skeleton :type="text" />
                </div>
              </div>
              <q-item>
                <div class="text-caption text-grey">
                  <q-skeleton :type="text" style="margin:10px;padding-left:80px;"/>
                </div>
              </q-item>
              <q-item>
                <q-skeleton :type="QBtn" style="margin:10px;padding-left:40px;" />
                <q-skeleton :type="QBtn" style="margin:10px;padding-left:40px;" />
                <q-skeleton :type="QBtn" style="margin:10px;padding-left:40px;" />
              </q-item>
            </q-card-section>
          </q-card>
        </q-page> -->
      </template>
    </Suspense>
  </template>
</template>
