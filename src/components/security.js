import { store } from './store.js'
import { router } from './../router.js'

let Security = {
  requireToken: function () {
    if (store.token === '') {
      router.push('/login')
      return false
    }
  },

  requestOptions: function (payload) {
    const headers = new Headers()
    headers.append('Content-Type', 'application/json')
    headers.append('Authorization', 'Bearer ' + store.token)
    return {
      method: 'POST',
      headers: headers,
      body: JSON.stringify(payload)
    }
  }
}

export default Security
