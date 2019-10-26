<template>
  <div id="chat-container" class="bg-black text-white">
    <span></span>
    <div id="message-log">
      <Message icon="https://i.imgur.com/Ska1zn6.png" message="Yeet the props are doing it" />
      <Message icon="https://i.imgur.com/Ska1zn6.png" message="Why are you like this" />

      <Message icon="https://i.imgur.com/Ska1zn6.png" message="yep this is fine" />

      <Message icon="https://i.imgur.com/Ska1zn6.png" message="is is dinner yet" />

      <Message icon="https://i.imgur.com/Ska1zn6.png" message="overcomplicated is fine" />

      <Message icon="https://i.imgur.com/Ska1zn6.png" message="my job is hard" />

      <Message icon="https://i.imgur.com/Ska1zn6.png" message="did we win" />

      <Message icon="https://i.imgur.com/Ska1zn6.png" message="get me all the songs" />

      <Message icon="https://i.imgur.com/Ska1zn6.png" message="Kappa" />

      <Message icon="https://i.imgur.com/Ska1zn6.png" message="Kappa" />

      <Message icon="https://i.imgur.com/Ska1zn6.png" message="Kappa" />

      <Message icon="https://i.imgur.com/Ska1zn6.png" message="Kappa" />

      <Message icon="https://i.imgur.com/Ska1zn6.png" message="Kappa" />

      <Message icon="https://i.imgur.com/Ska1zn6.png" message="Kappa" />

      <Message icon="https://i.imgur.com/Ska1zn6.png" message="heKapparld" />
      <Message icon="https://i.imgur.com/Ska1zn6.png" message="heKapparld" />
      <Message icon="https://i.imgur.com/Ska1zn6.png" message="heKapparld" />

      <Message icon="https://i.imgur.com/Ska1zn6.png" message="heKapparld" />

      <Message
        icon="https://i.imgur.com/Ska1zn6.png"
        message="this is going tor e bethe most misspleted piece of shit ever so get over it"
      />

      <Message icon="https://i.imgur.com/Ska1zn6.png" message="heKapparld" />

      <Message icon="https://i.imgur.com/Ska1zn6.png" message="heKapparld" />

      <Message icon="https://i.imgur.com/Ska1zn6.png" message="hellow world" />
    </div>
    <div id="send-message">
      <hr />
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
  color: white
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
setInterval(updateScroll, 1000);

export default {
  name: "chat-log",
  components: { Send, Message }
};
</script>
