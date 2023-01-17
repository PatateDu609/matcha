<script setup>
import { useRoute, useRouter } from 'vue-router';
import { api } from 'boot/axios';

const router = useRouter();
const route = useRoute();
console.log(route.query);
console.log(route);

if (!('token' in route.query)) {
  router.push('/');
}

api.put('/user/verify', route.query['token']).then();

router.push('/');
</script>

<template>
  <div></div>
</template>
