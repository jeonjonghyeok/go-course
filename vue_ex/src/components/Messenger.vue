<template>
    <div v-for="(message, idx) in messages" 
    :key="idx">
        {{message}}
    </div>
    <MessageInput @send="sendMessage" />
</template>

<script>
import MessageInput from '@/components/MessageInput.vue'
import { ref } from 'vue'
export default {
    name: 'Messenger',
    components: {
        MessageInput
    },
    setup() {
        const socket = new WebSocket('ws://localhost:5000');
        const message = ref("")
        const messages = ref([])
        const sendMessage = () => {
            socket.send(message)
        }
        socket.onmessage = (message) => {
            messages.value.push(message.data)
        }
        
        

        return {
            message,
            messages,
            sendMessage

        }
    }
}
</script>

<style scoped>
</style>