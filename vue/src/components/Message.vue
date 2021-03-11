<template>
  <div class="message">
      <div class="text">
          {{message}}
      </div>
  </div>
</template>

<script>
export default {
  name: 'Message',
  props: {
    message: String
  }
};
</script>

<style scoped>
.message {
  font-family: sans-serif;
}

.text {
  word-wrap: break-word;
  display: inline-block;
  margin: 0;
  flex-grow: 0;
  max-width: 300px;
  background-color: #fff;
  font-size: 14px;
  padding: 10px;
  border-radius: 10px;
  box-shadow: 0 2px 2px rgba(0, 0, 0, 0.15);
  background-color: #ffed4e;
}
</style>