<template>
    <form
    class="message-input-form"
    @submit.prevent="sendMessage"
    >
        <input 
        type="text"
        v-model="message" 
        placeholder="메세지를 입력하세요"
        />
        <input
        type="submit"
        value="send"
        
        />
    </form>
</template>

<script>
import { ref } from 'vue'

export default {
    name: 'MessageInput',
    setup(props, { emit }){
        const message = ref("");
        const sendMessage = () => {
            emit("send",message.value.trim())
            message.value=''
        }
        return { 
            message,
            sendMessage
        }
    }
}
</script>

<style scoped>
.message-input-form {
    width: 500px;
    height: 50px;
    background-color: white;
    display: flex;
    border-radius: 10px;
    margin: auto;
    box-shadow: 0 2px 2px rgba(0,0,0,0.15);

}
.message-input-form input[type="text"]{
    box-sizing: border-box;
    padding: 0px 20px;
    width: 80%;
    font-size: 16px;
    border: none;
}
.message-input-form input[type="text"]:focus{
    outline:none;
}
.message-input-form input[type="text"]::placeholder{
    color: #aaa;
}

.message-input-form input[type="submit"] {
    box-sizing: border-box;
    width: 20%;
    border: none;
    font-size: 20px;
    border-radius: 10px;
    background-color: #ffed4e;
}
.message-input-form input[type="submit"]:focus{
    outline: none;
}
.message-input-form input[type="submit"]:hover {
    background-color: #e0d041;
}
.message-input-form input[type="submit"]:disabled {
    background-color: #ddd;
}


</style>