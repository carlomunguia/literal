<template>
  <AppHeader />
  <div>
    <router-view
      v-slot="{ Component }"
      :key="componentKey"
      @success="success"
      @error="error"
      @warning="warning"
      @forceUpdate="forceUpdate">
      <keep-alive>
        <component :is="Component" />
      </keep-alive>
    </router-view>
  </div>
  <AppFooter />
</template>

<script>
  import AppHeader from './components/AppHeader.vue'
  import AppFooter from './components/AppFooter.vue'
  import { store } from './components/store.js'
  import notie from 'notie'

  const getCookie = (name) => {
    return document.cookie.split('; ').reduce((r, v) => {
      const parts = v.split('=')
      return parts[0] === name ? decodeURIComponent(parts[1]) : r
    }, '')
  }

  export default {
    name: 'App',
    components: {
      AppHeader,
      AppFooter
    },
    data() {
      return {
        store,
        componentKey: 0
      }
    },
    beforeMount() {
      let data = getCookie('_site_data')

      if (data !== '') {
        let cookieData = JSON.parse(data)
        store.token = cookieData.token.token
        store.user = {
          id: cookieData.user.id,
          first_name: cookieData.user.first_name,
          last_name: cookieData.user.last_name,
          email: cookieData.user.email
        }
      }
    },
    methods: {
      success(message) {
        notie.alert({ type: 'success', text: message })
      },
      error(message) {
        notie.alert({ type: 'error', text: message })
      },
      warning(message) {
        notie.alert({ type: 'warning', text: message })
      },
      forceUpdate() {
        this.componentKey += 1
      }
    }
  }
</script>

<style></style>
