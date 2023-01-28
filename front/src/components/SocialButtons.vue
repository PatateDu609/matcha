<template>
  <div class="row justify-center items-center text-center">
    <div class="col-12">{{ message }}</div>
    <q-separator spaced="15px"/>
    <div class="col-12">
      <div class="flex-xl">
        <q-btn-group spread rounded>
          <google-component />
          <discord-component />
          <school42-component />
          <github-component />
        </q-btn-group>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import GoogleComponent from 'components/social-buttons/GoogleComponent.vue';
import DiscordComponent from 'components/social-buttons/DiscordComponent.vue';
import School42Component from 'components/social-buttons/School42Component.vue';
import GithubComponent from 'components/social-buttons/GithubComponent.vue';
import { ContextType } from 'components/models';

const props = defineProps<{
  contextType: ContextType;
}>();

const message =
  props.contextType == ContextType.LOG_IN ? 'Log in with' : 'Sign up with';
</script>

<style scoped></style>
