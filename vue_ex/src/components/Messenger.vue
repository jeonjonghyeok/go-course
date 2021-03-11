<template>
    <div class="myname-input-bar">
        <input 
            type="text"
            v-model="myname"
            placeholder ="name input"
        />
    </div>
    <div class="messages">
    <Message
    v-for="(message, idx) in messages"
    :key="idx"
    :message="message"
    :displaySender="(myname !== message.sender) && (messages[idx-1] ?? {sender: null}).sender !== message.sender"
    :mine="myname === message.sender"
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
        const socket = new WebSocket('ws://localhost:5000')
        const messages = ref([])
        const myname = ref("jjh");
        const sendMessage = (message) => {
            socket.send(JSON.stringify({
                sender : myname.value,
                text: message 
            })
            )
        }
        socket.onmessage = (message) => {
            messages.value.push(JSON.parse(message.data))
        }
        return { 
            messages,
            sendMessage,
            myname
        };
    }
};
</script>

<style scoped>
.myname-input-bar {
    box-sizing: border-box;
    width: 500px;
    font-size: 20px;
    margin: auto;

}
.messages {
    width: 500px;
    height: 100%;
    margin:auto;

    display: flex;
    flex-direction: column;
    justify-content: space-between;
    padding: 10px;

    padding-top: 70px;
    padding-bottom: 70px;
}
.message-input-bar{
    height: 60px;
    width: 100%;
    bottom: 0;
    position: fixed;
    background-color: #bbc9e0;
}
</style>