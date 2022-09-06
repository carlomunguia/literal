<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-5">Login</h1>
        <hr />
        <form-tag @event="submitHandler" name="vueForm" event="event">
          <text-input v-model="email" label="Email" type="email" name="email" required="true"></text-input>
          <text-input v-model="password" label="Password" type="password" name="password" required="true"></text-input>
          <hr />

          Email: {{ email }}

          <hr />
          <input type="submit" class="btn btn-primary" value="Login" />
        </form-tag>
      </div>
    </div>
  </div>
</template>

<script>
  import TextInput from './forms/TextInput.vue'
  import FormTag from './forms/FormTag.vue'
  import { store } from './store.js'
  export default {
    name: 'UserLogin',
    components: {
      TextInput,
      FormTag
    },
    data() {
      return {
        email: '',
        password: '',
        store,
      }
    },
    methods: {
      submitHandler() {
        console.log("submitHandler")

        const payload = {
          email: this.email,
          password: this.password
        }

        const requestOptions = {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(payload)
        }

        fetch("http://localhost:8081/users/login", requestOptions)
          .then((response) => response.json())
          .then((data) => {
            if (data.error) {
              console.log("Error: " + data.message)
            } else {
              console.log("Token: " + data.data.token.token)
              store.token = data.data.token.token
            }
          })
      }
    },
  }
</script>
