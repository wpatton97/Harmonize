<template>
  <div class="text-white">
    <song-art :title="nowPlaying.Title" :artist="nowPlaying.Author" :album="nowPlaying.Album" 
       :image="nowPlaying.Art"
       v-bind:user="{name: nowPlaying.Addedby.Name, avatar: nowPlaying.Addedby.Avatar}"/>
  </div>
</template>

<script>
import { mapState } from 'vuex';
import SongArt from './SongArt';

export default {
  name: "current-player",
  watch: {
     $route (to, from){
       // When we change channels, we need to actually get the information about the current song.
        this.$store.dispatch('rest/fetchNowPlaying', { channelID: to.params.channelID });
     }
  },
  mounted () {
      this.$store.dispatch('rest/fetchNowPlaying', { channelID: this.$route.params.channelID });
  },
  computed: mapState('rest', {
      nowPlaying: state => state.nowPlaying
  }),
  components: { SongArt },
  props: [ 'channelID' ]
};
</script>
