<template>
  <q-layout view="hHh lpR fFf">

    <q-header class="bg-black text-white">
      <q-toolbar>
        <q-btn flat color="white" label="#weird" />
        <q-btn flat color="white" label="#hiphop" />
        <q-btn flat color="white" label="#pop" />
      </q-toolbar>
    </q-header>

    <q-drawer show-if-above v-model="right" side="right" bordered>
      <!-- drawer content -->
    </q-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>

  </q-layout>
</template>

<script>
export default {
  name: 'AppLayout',
  data () {
    return {
      leftDrawerOpen: false
    }
  }
}
</script>
