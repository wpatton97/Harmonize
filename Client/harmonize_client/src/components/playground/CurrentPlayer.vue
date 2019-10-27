<template>
  <div class="text-white">
    <song-art :title="nowPlaying.Title" :artist="nowPlaying.Author" :album="nowPlaying.Album" 
       :image="base_url + nowPlaying.Art"
       v-bind:user="{name: nowPlaying.AddedBy.Name, avatar: nowPlaying.AddedBy.Avatar}"/>
     <div class="text-center">
       <q-btn-group outline>
          <q-btn :ripple="{ color: 'red' }" flat color="white" icon="thumb_down" size="23px"/>
          <q-btn v-if="playing === true" flat color="white" icon="music_off" @click="mute" size="45px"/>
          <q-btn v-else flat color="white" icon="music_note" @click="mute" size="45px"/>
          <q-btn :ripple="{ color: 'green' }" flat class="relative-position" color="white" icon="thumb_up" size="23px"/>
       </q-btn-group>
      </div>
      <div class="audio-component">
          <av-bars
            :canv-width="500"
            :canv-height="300"
            class="viz"
            caps-color="#FFF"
            :bar-color="['#FFF']"
            canv-fill-color="transparent"
            :caps-height="5"
            ref-link="player"
            controls="false"
          ></av-bars>
          <audio ref="player" controls="false" v-bind:src="base_url + nowPlaying.URL" preload="auto" style="display:hidden;"></audio>
      </div>
    </div>
  </div>
</template>

<style lang="sass">
.viz
  position: fixed;
  bottom: 0;
  opacity: 0.3;
  left: 0;
</style>

<script>
import { mapState } from 'vuex';
import SongArt from './SongArt';
import AudioVisual from 'vue-audio-visual'

export default {
  name: "current-player",
  components: { AudioVisual },
  methods: {
      mute(){
        this.$store.dispatch('rest/setPlaying', { playing: !this.$store.state.rest.playing  });
      },
      audioPlay() {
        console.log("audio data is loaded");
        alert("audio loaded and can be played!");
      }
  },
  watch: {
     $route (to, from){
       // When we change channels, we need to actually get the information about the current song.
        this.$store.dispatch('rest/fetchNowPlaying', { channelID: to.params.channelID });
        let time = this.$store.state.rest.nowPlaying.Time.Current;

        var player = this.$refs.player;

        player.crossOrigin = 'anonymous';

        setTimeout(function() {
          //your code to be executed after 1 second
          player.currentTime = time + 1.5;
          player.play();
        }, 1500);
     }
  },
  mounted () {
      this.$store.dispatch('rest/fetchNowPlaying', { channelID: this.$route.params.channelID });
      let time = this.$store.state.rest.nowPlaying.Time.Current;

      var player = this.$refs.player;

      player.crossOrigin = 'anonymous';

      setTimeout(function() {
        //your code to be executed after 1 second
        player.currentTime = time + 1.5;
        player.play();
      }, 1500);

      /*this.$refs.player.addEventListener("ended", function(){
          alert("Song ended!");

          this.$store.dispatch('rest/fetchNowPlaying', { channelID: this.$route.params.channelID });
      });*/
  },
  computed: mapState('rest', {
      nowPlaying: state => state.nowPlaying,
      playing: state => state.playing,
      base_url: state => state.base_url
  }),
  components: { SongArt },
  props: [ 'channelID' ]
};
</script>
