<template>
  <div id="message-container">
      <q-item >
        <q-item-section avatar>
          <q-avatar rounded color="purple" text-color="white" v-bind:icon="icon" />
        </q-item-section>
        <q-item-section>{{message}}</q-item-section>
      </q-item>
  </div>    
</template>

<style lang="sass">
#Message

</style>

<script>
export default {
  name: 'Message',
  props: ["icon", "message"]
}
</script>
