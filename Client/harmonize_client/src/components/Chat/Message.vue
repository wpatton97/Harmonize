<template>
  <div id="message-container">
      <q-item >
        <q-item-section avatar>
          <q-avatar rounded >
              <img v-bind:src="icon" />
          </q-avatar>
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
