<template>
  <div id="send-container">
    <q-input class="bg-grey" id="Send" label="Message" counter maxlength="12" :dense="dense">
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
  name: "Send"
};
</script>
