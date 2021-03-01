<template>
<form
   @submit.prevent="sendMessage"
   >
        <input type="text" v-model="message" placeholder="아무말이나 입력하세요"/>
        <input 
        type="submit" 
        value="보내기"
        :disabled="!messageIsValid"
        />
    </form>
</template>

<script>
import { computed ,ref } from "vue"
export default {
    name: 'MessageInput',
    setup(props,{ emit }) {
        const message = ref("")
        const sendMessage = () => {
            emit("send",message.value.trim())
            message.value = ''
        }
        const messageIsValid = computed(() =>message.value.trim() !== "");
        return { 
            message,
            sendMessage,
            messageIsValid
        }
    }
}
</script>

<style scoped>
</style>