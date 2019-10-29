<template>
  <div id="message-container">
      <q-item >
        <q-item-section avatar>
          <q-avatar>
              <img v-bind:src="icon" />
          </q-avatar>
        </q-item-section>

        <q-item-section>
          <q-item-label class="text-weight-bolder">{{author}}</q-item-label>
          <q-item-label>{{message}}</q-item-label>
        </q-item-section>
      </q-item>
  </div>    
</template>

<style lang="sass">
</style>

<script>
export default {
  name: 'Message',
  props: ["icon", "message", "author"]
}
</script>
