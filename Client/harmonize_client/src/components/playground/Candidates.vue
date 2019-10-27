<template>
  <div>
    <span class="text-h5 text-white">Vote for...</span> 
        <q-btn flat color="grey" icon="add_circle_outline" size="23px" id="AddSong" @click="add" />
        <div>
            <song-art class="candidate-art float-left" v-for="candidate in candidates" :key="candidate.ID" :title="candidate.Song.Title" :artist="candidate.Song.Author" :album="candidate.Song.Album" v-bind:user="{name: candidate.Addedby.Name, avatar: candidate.Addedby.Avatar}" 
                 width="200px"
                 :image="base_url + candidate.Song.Art"/>
        </div>
   </div>
</template>

<style lang="sass">
.candidate-art:hover .song-art
    outline: 4px solid #87D37C !important
#AddSong
  float: right
#AddSong:hover
  color: white !important
</style>

<script>
import { mapState } from 'vuex';
import SongArt from './SongArt';

export default {
  name: "candidates",
  methods: {
    add() {
      alert('Pick a song')
    }
  },
  watch: {
     $route (to, from){
       // TODO: since we'll be doing a repeat either way, this doesn't matter as much
       this.$store.dispatch('rest/fetchCandidates', { channelID: this.$route.params.channelID });
     }
  },
  computed: mapState('rest', {
      candidates: state => state.candidates,
      base_url: state => state.base_url
  }),
  mounted () {
      this.$store.dispatch('rest/fetchCandidates', { channelID: this.$route.params.channelID });
  },
  components: { SongArt },
  props: [ 'channelID' ]
};
</script>
