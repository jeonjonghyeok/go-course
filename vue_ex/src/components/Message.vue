<template>
<div>
   {{props.message}} 
   </div>
</template>

<script>
export default {
    name: 'Message',
    props: {
        message: String
    },
    setup(props){
        return { 
            props
        }
    }

}
</script>

<style scoped>
</style>