<template>
  <q-page class="bg-black">
    <channel-selector />
    <div class="background-image"></div>
    <div class="content">
        <current-player />
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
  background-image: url('http://localhost:8080/statics/lately-ep.jpg');
  background-position: center;
  background-repeat: no-repeat;
  background-size: cover;
  height: 100%;
  filter: blur(20px)
  opacity: 0.4;
.content
  position: fixed;
  left: 0;
  right: 0;
  z-index: 9999;
  top: 50%;
</style>

<script>
import ChannelSelector from '../components/ChannelSelector';
import CurrentPlayer from '../components/playground/CurrentPlayer';
import Visualizer from '../components/playground/Visualizer';
import Voting from '../components/playground/Voting';

export default {
  name: 'Player',
  components: {
    CurrentPlayer,
    Visualizer,
    Voting,
    ChannelSelector
  }
}
</script>
