import { createRouter, createWebHistory } from 'vue-router'
import AppBody from './../components/AppBody.vue'
import UserLogin from './../components/UserLogin.vue'
import BooksList from './../components/BooksList.vue'
import AppBook from './../components/AppBook.vue'
import BooksAdmin from './../components/BooksAdmin.vue'
import BookEdit from './../components/BookEdit.vue'
import AppUsers from './../components/AppUsers.vue'
import UserEdit from './../components/UserEdit.vue'
import Security from './../components/security.js'

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
  },
  {
    path: '/books',
    name: 'Books',
    component: BooksList
  },
  {
    path: '/books/:bookName',
    name: 'Book',
    component: AppBook
  },
  {
    path: '/admin/books',
    name: 'BooksAdmin',
    component: BooksAdmin
  },
  {
    path: '/admin/books/:bookId',
    name: 'BookEdit',
    component: BookEdit
  },
  {
    path: '/admin/users',
    name: 'Users',
    component: AppUsers
  },
  {
    path: '/admin/users/:userId',
    name: 'User',
    component: UserEdit
  }
]

const router = createRouter({ history: createWebHistory(), routes })

router.beforeEach(() => {
  Security.checkToken()
})

export default router
