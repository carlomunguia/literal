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
            <th>Active</th>
            <th>Status</th>
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

            <td v-if="u.active === 1">
              <span class="badge bg-success">Active</span>
            </td>
            <td v-else>
              <span class="badge bg-danger">Inactive</span>
            </td>
            <td v-if="u.token.id > 0">
              <span class="badge bg-success" @click="logUserOut(u.id)">Logged In</span>
            </td>
            <td v-else>
              <span class="badge bg-danger">So Not Logged In</span>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
  import Security from './security.js'
  import { store } from './store.js'
  import notie from 'notie'

  export default {
    data() {
      return {
        users: [],
        store
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
            this.ready = true
          }
        })
        .catch((error) => {
          this.$emit('error', error)
        })
    },
    methods: {
      logUserOut(id) {
        if (id !== store.user.id) {
          notie.confirm({
            text: 'Are you sure you want to log this user out?',
            submitText: 'Yes',
            cancelText: 'No',
            submitCallback: () => {
              fetch(
                process.env.VUE_APP_LITERAL_API_URL + '/admin/users/logout/' + id,
                Security.requestOptions('')
              )
                .then((response) => response.json())
                .then((response) => {
                  if (response.error) {
                    this.$emit('error', response.message)
                  } else {
                    this.$emit('success', response.message)
                    this.$router.go()
                  }
                })
                .catch((error) => {
                  this.$emit('error', error)
                })
            }
          })
        }
      }
    }
  }
</script>
