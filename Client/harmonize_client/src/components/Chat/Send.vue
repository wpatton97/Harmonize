<template>
  <div id="send-container">
    <q-input id="Send" label="Message" counter maxlength="120" dark color="white">
      <template v-slot:after>
        <q-btn round dense flat icon="send" color="white" />
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
