import { createRouter, createWebHistory } from 'vue-router'
import AppBody from './../components/AppBody.vue'
import UserLogin from './../components/UserLogin.vue'
import BooksComp from './../components/BooksComp.vue'
// import BooksList from './../components/BooksList.vue'
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
    name: 'UserLogin',
    component: UserLogin
  },
  {
    path: '/books',
    name: 'BooksComp',
    component: BooksComp
  },
  {
    path: '/books/:bookName',
    name: 'AppBook',
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
    name: 'AppUsers',
    component: AppUsers
  },
  {
    path: '/admin/users/:userId',
    name: 'UserEdit',
    component: UserEdit
  }
]

const router = createRouter({ history: createWebHistory(), routes })

router.beforeEach(() => {
  Security.checkToken()
})

export default router
