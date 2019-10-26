<template>
  <div class="song-art-wrapper">
    <div>
      <q-img
        class="song-art shadow-23"
        src="http://localhost:8080/statics/lately-ep.jpg"
        spinner-color="white"
      />
    </div>
    <div class="q-pt-md q-pa-sm text-center text-h4 text-weight-bolder">{{ title }}</div>
    <div class="text-center text-h6">{{ album }} · {{ artist }}</div>
  </div>
</template>

<style lang="sass">
.song-art-wrapper
    width: 250px;
.song-art
    width: 250px;
</style>

<script>
export default {
  name: "song-art",
  props: ["title", "artist", "album"]
};
</script>