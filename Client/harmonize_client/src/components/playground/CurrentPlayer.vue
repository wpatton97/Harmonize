<template>
  <div class="text-white">
    <song-art title="Lava" artist="Still Woozy" album="Lately EP" image="http://localhost:8080/statics/lately-ep.jpg"/>
  </div>
</template>

<script>
import SongArt from './SongArt';

export default {
  name: "current-player",
  components: { SongArt }
};
</script>
