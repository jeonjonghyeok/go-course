<template>
  <div class="message" :class="{'my-message': mine}"> 
      <div class="sender" v-if="displaySender">
          {{props.message.sender}}
          </div>
    <div class="text" >
       {{props.message.text}}
        </div>
    </div>
    
</template>

<script>
export default {
    name: 'Message',
    props: {
        message: Object,
        mine: Boolean,
        displaySender: Boolean
    },
    setup(props) {
        return {
            props
        }
    }
};
</script>

<style scoped>
.sender {
    font-size: 12px;
}
.my-message .text{
    background-color: #ffed4e;
}
.message {
    font-family: sans-serif;
    margin: 5px;
}
.text {
    display: inline-block;
    font-size: 14px;
    padding: 10px;
    background-color: #fff;
    border-radius: 10px;
    box-shadow: 0 2px 2px rgba(0, 0, 0, 0.15);
}

</style>
