<template>
  <router-view />
</template>

<script lang="ts" setup>
import { useQuasar } from 'quasar';

const customIcons: { [Key: string]: string } = {
  'app:42school': 'img:src/assets/42_Logo.svg',
};

const $q = useQuasar();

$q.iconMapFn = (iconName: string) => {
  const icon = customIcons[iconName];

  if (icon !== void 0) {
    return { icon: icon };
  }
};
</script>
