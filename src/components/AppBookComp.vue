<template>
  <div class="container">
    <div class="row">
      <div class="col-md-2">
        <img
          v-if="ready"
          class="img-fluid img-thumbnail"
          :src="`${imgPath}/covers/${book.slug}.jpg`"
          alt="cover" />
      </div>
      <div class="col-md-10">
        <template v-if="ready">
          <h1>{{ book.title }}</h1>
          <p>{{ book.description }}</p>
          <p>
            <a :href="book.url" target="_blank">Read more</a>
          </p>
          <h3 class="mt-3">{{ book.title }}</h3>
          <hr />
          <p>
            <strong>Author:</strong> {{ book.author.author_name }}<br />
            <strong>Published:</strong> {{ book.publication_year }}<br />
          </p>
          <p>
            {{ book.description }}
          </p>
        </template>
        <p v-else>Loading...</p>
      </div>
    </div>
  </div>
</template>

<script>
  import { ref, onMounted } from 'vue'
  import { useRoute } from 'vue-router'

  export default {
    name: 'AppBookComp',
    emits: ['error'],
    props: {},

    setup(props, context) {
      let ready = ref(false)
      const imgPath = ref(process.env.VUE_APP_LITERAL_IMAGE_URL)
      let book = ref({})
      const route = useRoute()

      onMounted(() => {
      fetch(process.env.VUE_LITERAL_API_URL + '/books' + route.params.bookName)
        .then((response) => {
          if (response.ok) {
            return response.json()
          } else {
            throw new Error('Network response was not ok.')
          }
        })
        .then((jsonResponse) => {
          books.value = jsonResponse
          ready.value = true
        })
        .catch((error) => {
          console.error('There has been a problem with your fetch operation:', error)
          context.emit('error', error)
        })
    })

      return {
        ready,
        imgPath,
        book
      }
    }
  }

</script>
