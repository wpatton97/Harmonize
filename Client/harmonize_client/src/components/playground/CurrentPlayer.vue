<template>
  <div class="text-white">
    <song-art title="Lava" artist="Still Woozy" album="Lately EP" 
       image="http://localhost:8080/statics/lately-ep.jpg"
       v-bind:user="{name: 'Goose', avatar: 'https://i.imgur.com/Ska1zn6.png'}"/>
  </div>
</template>

<script>
import { mapState } from 'vuex';
import SongArt from './SongArt';

export default {
  name: "current-player",
  mounted () {
      this.$store.dispatch('rest/fetchNowPlaying', { channelID: 1 });
  },
  computed: mapState('rest', {
      nowPlaying: state => state.nowPlaying
  }),
  components: { SongArt },
  props: [ 'channelID' ]
};
</script>
