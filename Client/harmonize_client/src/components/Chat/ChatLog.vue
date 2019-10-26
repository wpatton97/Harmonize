<template>
  <div id="chat-container" class="bg-black text-white">
    <div>
      <Message />
    </div>

    
    <div id="SendMessage"> 
      <hr />
      <Send />
    </div>
  </div>
</template>

<style lang="sass">
#chat-container
  height: 100%
#SendMessage
  position: absolute
  bottom: 0
  width: 100%
  border-radius: 
</style>

<script>
import Message from "./Message";
import Send from "./Send";

export default {
  name: "chat-log",
  components: { Send, Message }
};
</script>
