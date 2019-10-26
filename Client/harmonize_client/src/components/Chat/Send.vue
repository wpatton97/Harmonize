<template>
  <div id="send-container">
      <q-input class="bg-black darken-1" color="white" id="Send" label="Message" counter maxlength="120" :dense="dense">
        <template v-slot:before>
          <q-avatar>
            <img src="https://imgur.com/Ska1zn6.png">
          </q-avatar>
        </template>
        <template v-slot:after>
          <q-btn round dense flat icon="send" />
        </template>
      </q-input>
  </div>
</template>

<style lang="sass">
#Send
  border-radius: 2em

</style>

<script>
export default {
  name: 'Send'
}
</script>
