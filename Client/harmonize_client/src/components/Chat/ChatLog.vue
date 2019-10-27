<template>
  <div id="chat-container" class="bg-black text-white shadow-7">
    <span></span>
    <div id="message-log">
      <Message v-for="n in 25" :key="n" icon="https://i.imgur.com/Ska1zn6.png" message="Yeet the props are doing it" />
    </div>
    <div id="send-message">
      <Send />
    </div>
  </div>
</template>

<style lang="sass">
#chat-container
  height: 100%
#send-message
  position: absolute
  bottom: 0
  width: 100% 
  padding: .75em
#message-log
  height: 90%
  overflow: scroll
</style>

<script>
import Message from "./Message";
import Send from "./Send";

function updateScroll() {
  var element = document.getElementById("message-log");
  element.scrollTop = element.scrollHeight;
}
setInterval(updateScroll, 5000);

export default {
  name: "chat-log",
  components: { Send, Message }
};
</script>
