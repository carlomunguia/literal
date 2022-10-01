<template>
  <div class="container">
    <div class="col">
      <h1 class="mt-3">Book Edit</h1>
    </div>
  </div>
</template>

<script>
  import Security from './security.js'

  export default {
    name: 'BookEdit',
    beforeMount() {
      Security.requireToken()
    }
  }
</script>
