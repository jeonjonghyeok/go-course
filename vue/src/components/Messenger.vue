<template>
  <div class="messages">
    <Message
      style="margin: 5px"
      v-for="(message, idx) in messages"
      :key="idx"
      :message="message"
    />
  </div>

  <div class="message-input-bar">
    <MessageInput @send="sendMessage" />
  </div>
</template>

<script>
import { ref } from 'vue'
import Message from '@/components/Message.vue'
import MessageInput from '@/components/MessageInput.vue'

export default {
  name: 'Messenger',
  components: {
    Message,
    MessageInput
  },
  setup() {
    const socket = new WebSocket('ws://localhost:8081')

    const messages = ref([])
    const sendMessage = message => {
      socket.send(message)
    }

    socket.onmessage = (message) => {
        messages.value.push(message.data)
    }

    return {
      messages,
      sendMessage
    };
  }
};
</script>

<style scoped>
.messages {
  width: 500px;
  height: 100%;
  margin: auto;

  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: stretch;

  flex-grow: 1;
  overflow: auto;
  padding: 10px;
  display: flex;
  flex-direction: column;
  align-items: stretch;
}

.message-input-bar {
  height: 60px;
  width: 100%;
  position: fixed;
  bottom: 0;
  background-color: #bbc9e0;
}
</style>