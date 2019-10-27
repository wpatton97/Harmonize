<template>
  <q-toolbar class="selector-wrapper">
     <q-btn v-for="channel in channels" :key="channel.ID" v-bind:to="'channel/' + channel.ID" flat color="white" v-bind:label="'#' + channel.Name" class="text-lowercase text-h5" v-bind:class="{ inactive: channel.ID !== $route.params.channelID }"/>
  </q-toolbar>
</template>

<style lang="sass">
.selector-wrapper
  position: fixed;
  left: 20px;
  z-index: 50;
.inactive
  opacity: 0.6;
</style>

<script>
import { mapState } from 'vuex'

export default {
  name: 'channel-selector',
  mounted () {
       this.$store.dispatch('fetchChannels')
  },
  computed: mapState('rest', {
      channels: state => state.channels
  })
}
</script>
