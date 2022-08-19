import { createRouter, createWebHistory } from 'vue-router'
import AppBody from './../components/AppBody.vue'
import UserLogin from './../components/UserLogin.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: AppBody
  },
  {
    path: '/login',
    name: 'Login',
    component: UserLogin
  }
]

const router = createRouter({ history: createWebHistory(), routes })
export default router
