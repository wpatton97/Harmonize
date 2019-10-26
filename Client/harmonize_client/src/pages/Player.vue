<template>
  <q-page class="bg-black">
    <channel-selector />
    <div class="background-image"></div>
    <div class="content">
      <div class="q-pa-md">
          <div class="row">
            <div class="col-3">         
              <current-player />
            </div>
            <div class="col q-ml-lg">
               <span class="text-h5 text-white">Up next...</span>
               <div class="q-gutter-sm q-mt-md">
                 <song-art title="Goodie Bag" artist="Still Woozy" album="Goodie Bag" width="200px" image="http://localhost:8080/statics/goodie-bag.jpg"/>
                 <song-art title="Goodie Bag" artist="Still Woozy" album="Goodie Bag" width="200px" image="http://localhost:8080/statics/goodie-bag.jpg"/>
               </div>
            </div>
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
  z-index: 1;
  display: block;
  background: url('http://localhost:8080/statics/lately-ep.jpg');
  background-position: center;
  background-repeat: no-repeat;
  background-size: cover;
  height: 100%;
  filter: blur(20px) grayscale(100%);
  opacity: 0.7;
.content
  position: fixed;
  left: 0;
  right: 0;
  z-index: 9999;
  top: 20%;
  left: 50px;
</style>

<script>
import ChannelSelector from '../components/ChannelSelector';
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
    SongArt
  }
}
</script>
