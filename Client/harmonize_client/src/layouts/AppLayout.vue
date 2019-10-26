<template>
  <q-layout view="lHh Lpr lFf">
    <ChannelSelector />
    <q-page-container>
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script>
import { ChannelSelector } from '../components/ChannelSelector'

export default {
  name: 'AppLayout',
  components: {
    ChannelSelector
  },
  data () {
    return {
      leftDrawerOpen: false
    }
  }
}
</script>
