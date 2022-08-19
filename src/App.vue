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

  export default {
    name: 'App',
    components: {
      AppHeader,
      AppFooter
    }
  }
</script>

<style></style>
