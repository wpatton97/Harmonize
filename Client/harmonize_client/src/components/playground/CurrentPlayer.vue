<template>
  <div class="text-white">
    <song-art :title="nowPlaying.Title" :artist="nowPlaying.Author" :album="nowPlaying.Album" 
       :image="nowPlaying.Art"
       v-bind:user="{name: nowPlaying.Addedby.Name, avatar: nowPlaying.Addedby.Avatar}"/>
     <div class="text-center">
       <q-btn-group outline>
          <q-btn :ripple="{ color: 'red' }" flat color="white" icon="thumb_down" size="23px"/>
          <q-btn v-if="playing === true" flat color="white" icon="music_off" @click="mute" size="45px"/>
          <q-btn v-else flat color="white" icon="music_note" @click="mute" size="45px"/>
          <q-btn :ripple="{ color: 'green' }" flat class="relative-position" color="white" icon="thumb_up" size="23px"/>
       </q-btn-group>
      </div>
      <div class="audio-component">
          <audio controls="true" src="https://www.mfiles.co.uk/mp3-downloads/grieg-holberg-suite-3-gavotte.mp3"></audio>
      </div>
  </div>
</template>

<script>
import { mapState } from 'vuex';
import SongArt from './SongArt';
import AudioVisual from 'vue-audio-visual'

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
