<template>
  <AppHeader />
  <AppBody />
  <AppFooter />
</template>

<script>
import AppHeader from './components/AppHeader.vue'
import AppFooter from './components/AppFooter.vue'
import AppBody from './components/AppBody.vue'

export default {
    name: 'App',
    components: {
        AppHeader,
        AppFooter,
        AppBody
    }
}
</script>

<style>

</style>
