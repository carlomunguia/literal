<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-5">Login</h1>
        <hr />
        <form-tag @event="submitHandler" name="vueForm" event="event">
          <text-input
            v-model="email"
            label="Email"
            type="email"
            name="email"
            required="true"></text-input>
          <text-input
            v-model="password"
            label="Password"
            type="password"
            name="password"
            required="true"></text-input>
          <hr />
          <input type="submit" class="btn btn-primary" value="Login" />
        </form-tag>
      </div>
    </div>
  </div>
</template>

<script>
  import { ref, onMounted } from 'vue'
  import TextInput from './forms/TextInput.vue'
  import { store } from './store.js'
  import FormTag from './forms/FormTag.vue'
  import router from './../router/index.js'
  import Security from './security.js'

  export default {
    name: 'LoginComp',
    props: {},
    emits: ['error'],
    components: {
      'text-input': TextInput,
      'form-tag':  FormTag,
    },

    setup(props, context) {
      let email = ref('')
      let password = ref('')

      onMounted(() => {
        console.log("Using new component")
      })

      function submitHandler() {
        const payload = {
          email: email.value,
          password: password.value,
        }

        fetch(
          process.env.VUE_APP_LITERAL_API_URL + '/users/login',
          Security.requestOptions(payload)
        )
          .then((response) => response.json())
          .then((response) => {
            if (response.error) {
              context.emit('error', response.message)
            } else {
              store.token = response.data.token.token

              store.user = {
                id: response.data.user.id,
                first_name: response.data.user.first_name,
                last_name: response.data.user.last_name,
                email: response.data.user.email
              }

              let date = new Date()
              let expDays = 1
              date.setTime(date.getTime() + expDays * 24 * 60 * 60 * 1000)
              const expires = 'expires=' + date.toUTCString()

              document.cookie =
                '_site_data=' +
                JSON.stringify(response.data) +
                '; ' +
                expires +
                '; path=/; SameSite=strict; Secure;'
              router.push('/')
            }
          })
      }
    }
  }

</script>
