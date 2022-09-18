<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-3">All Users</h1>
      </div>

      <hr />

      <table class="table table-compact table-striped">
        <thead>
          <tr>
            <th>User</th>
            <th>Email</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="u in this.users" v-bind:key="u.id">
            <td>
              <router-link :to="`/admin/users/${u.id}`"
                >{{ u.last_name }}, {{ u.first_name }}</router-link
              >
            </td>
            <td>{{ u.email }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
  import Security from './security.js'

  export default {
    data() {
      return {
        users: []
      }
    },
    beforeMount() {
      Security.requireToken()

      fetch(process.env.VUE_APP_LITERAL_API_URL + '/admin/users', Security.requestOptions(''))
        .then((response) => response.json())
        .then((response) => {
          if (response.error) {
            this.$emit('error', response.message)
          } else {
            this.users = response.data.users
          }
        })
        .catch((error) => {
          this.$emit('error', error)
        })
    }
  }
</script>
