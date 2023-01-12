<script setup lang="ts">
</script>

<template>
  <q-page class="row justify-center items-center" padding>
    <div class="row col justify-center">
      <div class="col" style="max-width: 400px">
        <q-form class="column q-gutter-md">
          <q-input type="text" outlined name="username" label="Username" />
          <q-input type="password" outlined name="password" label="Password" />
          <q-input type="text" outlined name="email" label="Email" />
          <q-input type="text" outlined name="first-name" label="First Name" />
          <q-input type="text" outlined name="last-name" label="Last Name" />

          <q-btn color="secondary" type="submit">Sign up</q-btn>
        </q-form>
      </div>
    </div>
  </q-page>
</template>

<style scoped>

</style>
