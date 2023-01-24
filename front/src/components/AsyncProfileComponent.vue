<template>
  {{ data }}
</template>

<script lang="ts">
import { ref } from 'vue'
import { api } from '../boot/axios';

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

export default {
  name: 'MyAsyncComponent',
  async setup() {
    let data = ref(null)

    await new Promise((r) => setTimeout(r, 2000))

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
      data = res.username;
    } catch (e) {
      console.error(e)
    }
    return { data }
  },
}
</script>