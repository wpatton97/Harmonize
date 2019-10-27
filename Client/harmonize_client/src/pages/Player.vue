<template>
  <q-page class="bg-black">
    <channel-selector />
    <div ref="bg" class="background-image"></div>
    <div class="content q-pa-md">
          <div class="row">
          
            <div class="col-3">         
              <current-player />
            </div>

            <div class="col q-ml-xl">
              <candidates />
            </div>
          </div>
    </div>
  </q-page>
</template>

<style lang="sass">
.background-image
  position: fixed;
  left: 0;
  right: 0;
  bottom: 0;
  z-index: 0;
  display: block;
  background-position: center;
  background-repeat: no-repeat;
  background-size: cover;
  height: 100%;
  filter: blur(20px) grayscale(100%);
  opacity: 0.7;
.content
  position: absolute;
  left: 0;
  right: 0;
  bottom: 0;
  z-index: 9999;
  top: 20%;
  padding-left: 50px;
  right: 10%;
</style>

<script>
import { mapState } from 'vuex';

import ChannelSelector from '../components/ChannelSelector';
import Candidates from '../components/playground/Candidates';
import CurrentPlayer from '../components/playground/CurrentPlayer';
import Visualizer from '../components/playground/Visualizer';
import Voting from '../components/playground/Voting';
import SongArt from '../components/playground/SongArt';

export default {
  name: 'Player',
  components: {
    CurrentPlayer,
    Visualizer,
    Voting,
    ChannelSelector,
    SongArt,
    Candidates
  }, 
  created() {
    this.$store.watch(
      (state, getters) => getters.status,
      (newValue, oldValue) => {
        console.log(`Updating from ${oldValue} to ${newValue}`);

        
      },
    );
  },
  mounted () {
    console.log("url('" + this.$store.state.rest.base_url + this.$store.state.rest.nowPlaying.Art + "')");

    this.$refs.bg.style.backgroundImage = "url('" + this.$store.state.rest.base_url + this.$store.state.rest.nowPlaying.Art + "')";
  },
  data() {
    return { test: "shit" }
  },
  computed: mapState('rest', {
      nowPlaying: state => state.nowPlaying,
      base_url: state => state.base_url
  })
}
</script>
