<template>
  <div id="chat-container" class="bg-black text-white">
    test
    <div> 
      <Message />
    </div>
    
    <hr />
    <div> 
      <Send />
    </div>
  </div>
</template>

<style lang="sass">
#chat-container
  height: 100%
</style>

<script>
import Message from './Message';
import Send from './Send';

export default {
  name: 'chat-log'
}
</script>
