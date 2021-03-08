<template>
  <div class="message">
    <div class="text">
       {{props.message}}
        </div>
    </div>
</template>

<script>
export default {
    name: 'Message',
    props: {
        message: String
    },
    setup(props) {
        return {
            props
        }
    }
};
</script>

<style scoped>
.message {
    font-family: sans-serif;
    margin: 5px;
}
.text {
    display: inline-block;
    font-size: 14px;
    padding: 10px;
    background-color: #ffed4e;
    border-radius: 10px;
    box-shadow: 0 2px 2px rgba(0, 0, 0, 0.15);
}

</style>
