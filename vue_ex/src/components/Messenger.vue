<template>
    <Message v-for="(message, idx) in messages" 
    :key="idx"
    :message="message" 
    >
    </Message>
    <MessageInput @send="sendMessage" />
</template>

<script>
import MessageInput from '@/components/MessageInput.vue'
import Message from '@/components/Message.vue'
import { ref } from 'vue'
export default {
    name: 'Messenger',
    components: {
        MessageInput,
        Message
    },
    setup() {
        const socket = new WebSocket('ws://localhost:5000');
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

        }
    }
}
</script>

<style scoped>
</style>