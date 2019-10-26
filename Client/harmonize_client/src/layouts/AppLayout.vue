<template>
  <q-layout view="hHh lpR fFf">

    <q-header class="bg-black text-white">
      <channel-selector />
    </q-header>

    <q-drawer show-if-above v-model="right" side="right" bordered>
      <chat />
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>

  </q-layout>
</template>

<script>
import ChannelSelector from '../components/ChannelSelector';
import Chat from '../components/Chat/ChatLog';

export default {
  name: 'AppLayout',
  components: { ChannelSelector, Chat },
  data () {
    return {
      leftDrawerOpen: false
    }
  }
}
</script>
