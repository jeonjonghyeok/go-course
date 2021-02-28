<template>
<div>
  <div v-for="(item, idx) in items" :key="idx">
    {{item}}
  </div>
  </div>
</template>

<script>
import{ ref } from 'vue'
export default {
  name: 'App',
  setup() {
      const items=  ref(['one', 'two', 'three'])
    return {
      items
    }
  }
}
</script>

<style>
</style>
