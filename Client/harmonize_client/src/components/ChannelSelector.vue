<template>
  <q-toolbar class="selector-wrapper">
     <q-btn flat color="white" label="#weird" class="text-lowercase text-h4 text-weight-bolder" />
     <q-btn v-for="n in 5" :key="n" flat color="white" label="#hiphop" class="text-lowercase text-h5 inactive-channel" />
  </q-toolbar>
</template>

<style lang="sass">
.selector-wrapper
  position: fixed;
  left: 20px;
  z-index: 50;
.inactive-channel
  opacity: 0.6;
</style>

<script>
export default {
  name: 'channel-selector'
}
</script>
