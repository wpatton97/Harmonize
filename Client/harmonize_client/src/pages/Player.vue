<template>
  <q-page>
    <div class="q-pa-md">
        <div class="row">
          <div class="col">
            <channel-selector />
          </div>
          <div class="col">
            <chat />
          </div>
        </div>
    </div>
  </q-page>
</template>

<script>
import ChannelSelector from '../components/ChannelSelector';
import Chat from '../components/Chat';

export default {
  name: 'Player',
  components: {
    ChannelSelector,
    Chat
  }
}
</script>
