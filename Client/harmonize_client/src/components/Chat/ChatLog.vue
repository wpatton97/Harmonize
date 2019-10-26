<template>
  <div id="chat-container">
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
</style>

<script>
import Message from './Message';
import Send from './Send';


export default {
  name: 'Chat'
}
</script>
