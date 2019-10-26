<template>
  <q-page class="bg-black">
    <div class="background-image"></div>
    <div class="content">
        
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
  background-image: url('https://cdn.quasar.dev/img/mountains.jpg');
  height: 800px;
  filter: blur(20px);
.content
  position: fixed;
  left: 0;
  right: 0;
  z-index: 9999;
  margin-left: 20px;
  margin-right: 20px;
</style>

<script>
import CurrentPlayer from '../components/playground/CurrentPlayer';
import Visualizer from '../components/playground/Visualizer';
import Voting from '../components/playground/Voting';

export default {
  name: 'Player',
  components: {
    CurrentPlayer,
    Visualizer,
    Voting
  }
}
</script>
