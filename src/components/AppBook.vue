<template>
  <div class="container">
    <div class="row">
      <div class="col-md-2">
        <img
          v-if="this.ready"
          class="img-fluid img-thumbnail"
          :src="`${imgPath}/covers/${book.slug}.jpg`"
          alt="cover" />
      </div>
      <div class="col-md-10">
        <template v-if="this.ready">
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
  export default {
    data() {
      return {
        book: {},
        imgPath: process.env.VUE_APP_LITERAL_IMAGE_URL,
        ready: false
      }
    },
    activated() {
      fetch(process.env.VUE_APP_LITERAL_API_URL + '/books/' + this.$route.params.bookName)
        .then((response) => response.json())
        .then((data) => {
          if (data.error) {
            this.$emit('error', data.message)
          } else {
            this.ready = true
            this.book = data.data.book
          }
        })
    },
    deactivated() {}
  }
</script>
