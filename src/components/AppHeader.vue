<template>
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
  <div class="container-fluid">
    <a class="navbar-brand" href="#">Navbar</a>
    <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarNav">
      <ul class="navbar-nav">
        <li class="nav-item">
          <router-link class="nav-link active" aria-current="page" to="/">Home</router-link>
        </li>
        <li class="nav-item">
          <router-link v-if="store.token == ''" class="nav-link" to="/login">Login</router-link>
          <router-link v-else class="nav-link" to="/logout">Logout</router-link>
        </li>
      </ul>
    </div>
  </div>
</nav>
</template>

<script>

import { store } from './store.js'

export default {
  data() {
    return {
      store
    }
  }
}
</script>
