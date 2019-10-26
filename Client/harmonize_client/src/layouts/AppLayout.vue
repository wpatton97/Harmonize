<template>
  <q-layout view="hHh lpR fFf">

    <q-drawer show-if-above v-model="right" side="right">
      <chat-log />
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>

  </q-layout>
</template>

<script>
import ChatLog from '../components/Chat/ChatLog';

export default {
  name: 'AppLayout',
  components: { ChatLog },
  data () {
    return {
      leftDrawerOpen: false
    }
  }
}
</script>
