<template>
  <div id="chat-container" class="bg-black text-white shadow-7">
    <span id="header"> Chat</span>
    <div id="message-log">
      <Message icon="https://i.imgur.com/Ska1zn6.png" author="Goose" message="Good song!" />
      <Message icon="http://100.79.1.39:8080/userprofiles/xinx.png" author="Xinx" message="I agree that's pretty good" />
      <Message icon="https://i.imgur.com/Ska1zn6.png" author="Goose" message="vote for mine next please" />
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
  overflow-y: scroll
#header
  display: inline-block
  height: 42px
  font-size: x-large
  margin-left: .5em
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
