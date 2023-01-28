<script lang="ts" setup>

import { ref } from 'vue'
import { useUserStore } from 'stores/user';
import ChatMessageComponent from '../components/ChatMessageComponent.vue'

const userStore = useUserStore()
</script>

<template>
<div class="q-pa-md flex">
  <div class="q-pa-md flex justify-left">
    <div style="max-width: 90%; width: 300px;">
      <q-intersection
        v-for="index in 60"
        :key="index"
        transition="flip-right"
        class="example-item"
      >
        <q-item clickable v-ripple>
          <q-item-section avatar>
            <q-avatar color="primary" text-color="white">
              M
            </q-avatar>
          </q-item-section>

          <q-item-section>
            <q-item-label>Match #{{ index }}</q-item-label>
            <q-item-label caption lines="1">Connected 2min ago</q-item-label>
          </q-item-section>

          <q-item-section side>
            <q-icon name="chat_bubble" color="green" />
          </q-item-section>
        </q-item>
      </q-intersection>
    </div>
  </div>
  <div style="width: 600px;">
    <ChatMessageComponent/>
    <q-input outlined v-model="text" label="Chat here..." />
  </div>
</div>
</template>
