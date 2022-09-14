<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-3">User</h1>
        <hr />

        <form-tag @userEditEvent="submitHandler" name="userForm" event="userEditEvent" :user="user">
          <text-input
            v-model="user.first_name"
            type="text"
            required="true"
            label="First Name"
            name="first-name" />
          <text-input
            v-model="user.last_name"
            type="text"
            required="true"
            label="Last Name"
            name="last-name" />
          <text-input
            v-model="user.email"
            type="email"
            required="true"
            label="Email"
            name="email" />
          <text-input
            v-if="this.user.id === 0"
            v-model="user.password"
            type="password"
            required="true"
            label="Password"
            :value="user.password"
            name="password" />
          <text-input
            v-else
            v-model="user.password"
            type="password"
            label="Password"
            :value="user.password"
            name="password" />
          <hr />
        </form-tag>
      </div>
    </div>
  </div>
</template>

<script>
  import Security from './security.js'
  import FormTag from './forms/FormTag.vue'
  import TextInput from './forms/TextInput.vue'
  // import notie from 'notie'

  export default {
    beforeMount() {
      Security.requireToken()

      if (parseInt(String(this.$route.params.userId), 10) > 0) {
        //todo
      }
    },
    data() {
      return {
        user: {
          id: 0,
          first_name: '',
          last_name: '',
          email: '',
          password: ''
        }
      }
    },
    components: {
      'form-tag': FormTag,
      'text-input': TextInput
    },
    methods: {
      submitHandler() {}
    }
  }
</script>
