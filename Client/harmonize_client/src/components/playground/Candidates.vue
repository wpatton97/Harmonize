<template>
  <div>
    <span class="text-h5 text-white">Up next...</span> 
        <div>
            <song-art class="candidate-art float-left" v-for="candidate in candidates" :key="candidate.Song.ID" :title="candidate.Song.Title" :artist="candidate.Song.Author" :album="candidate.Song.Album" v-bind:user="{name: candidate.Addedby.Name, avatar: candidate.Addedby.Avatar}" 
                 width="200px"
                 :image="candidate.Song.Art"/>
        </div>
   </div>
</template>

<style lang="sass">
 .candidate-art:hover
    border: 15x solid red;
</style>

<script>
import { mapState } from 'vuex';
import SongArt from './SongArt';

export default {
  name: "candidates",
  watch: {
     $route (to, from){
       // TODO: since we'll be doing a repeat either way, this doesn't matter as much
       this.$store.dispatch('rest/fetchCandidates', { channelID: this.$route.params.channelID });
     }
  },
  computed: mapState('rest', {
      candidates: state => state.candidates
  }),
  mounted () {
      this.$store.dispatch('rest/fetchCandidates', { channelID: this.$route.params.channelID });
  },
  components: { SongArt },
  props: [ 'channelID' ]
};
</script>
