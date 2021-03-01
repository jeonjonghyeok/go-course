<template>
<form
   @submit.prevent="sendMessage"
   >
        <input type="text" v-model="message" placeholder="아무말이나 입력하세요"/>
        <input type="submit" value="보내기"/>
    </form>
</template>

<script>
import { ref } from "vue"
export default {
    name: 'MessageInput',
    setup(props,{ emit }) {
        const message = ref("")
        const sendMessage = () => {
            emit("send",message.value.trim())
        }
        return { 
            message,
            sendMessage
        }
    }
}
</script>

<style scoped>
</style>