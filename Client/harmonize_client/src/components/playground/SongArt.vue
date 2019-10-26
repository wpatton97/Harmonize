<template>
  <div class="song-art-wrapper q-ma-lg" v-bind:style="{ width: width }">
    <q-img
      class="song-art shadow-23"
      v-bind:src="image"
      spinner-color="white"
      v-bind:style="{ width: width }"
    />
    <div class="text-white q-pt-md q-pa-sm text-center text-h4 text-weight-bolder">{{ title }}</div>
    <div class="text-white text-center text-h6">{{ album }} · {{ artist }}</div>
    <div class="text-white text-center text-h6">
        <q-avatar size="28px">
          <img v-bind:src="user.avatar">
        </q-avatar>
        {{ user.name }}
    </div>
  </div>
</template>

<style lang="sass">
</style>

<script>
export default {
  name: "song-art",
  props: ["width", "title", "artist", "album", "image", "user"]
};
</script>