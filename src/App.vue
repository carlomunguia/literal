<template>
  <AppHeader />
  <div>
    <router-view />
  </div>
  <AppFooter />
</template>

<script>
  import AppHeader from './components/AppHeader.vue'
  import AppFooter from './components/AppFooter.vue'
  import { store } from './components/store.js'

  export default {
    name: 'App',
    components: {
      AppHeader,
      AppFooter
    },
    data() {
      return {
        store,
      }
    },
  }

</script>

<style></style>
