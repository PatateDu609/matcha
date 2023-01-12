<script>
import { ref } from 'vue'

export default {
  setup () {
    const url = ref('https://placeimg.com/500/300/nature')

    return {
      url,
      refresh () {
        url.value = 'https://placeimg.com/500/300/nature?t=' + Math.random()
      }
    }
  }
}
</script>

<template>
  <q-page class="row items-center justify-evenly">
    <div class="q-pa-md q-gutter-sm">
      <q-img :src="url" spinner-color="white" style="height: 500px; max-width: 300px; width: 300px; border-radius: 20px;" /><br>
      <q-btn push color="red" style="margin-right: 100px;" label="Pass ❌" @click="refresh" />
      <q-btn push color="green" label="Match ❤️" @click="refresh" />
    </div>
  </q-page>
</template>
