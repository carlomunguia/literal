<template>
  <div class="container">
    <div class="col">
      <h1 class="mt-3">Add / Edit Book</h1>
      <hr />

      <form-tag @bookEditEvent="submitHandler" name="bookForm" event="bookEditEvent">
        <div v-if="this.book.slug != ''" class="mb-3">
          <img
            :src="`${this.imgPath}/covers/${this.book.slug}.jpg`"
            class="img-fluid img-thumbnail book-cover"
            alt="cover" />
        </div>

        <div class="mb-3">
          <label for="formFile" class="form-label">Cover Image</label>
          <input
            v-if="this.book.id === 0"
            ref="coverInput"
            class="form-control"
            type="file"
            id="formFile"
            @change="loadCoverImage" />
          <input
            v-else
            ref="coverInput"
            class="form-control"
            type="file"
            id="formFile"
            @change="loadCoverImage" />
        </div>

        <text-input
          v-model="book.title"
          type="text"
          required="true"
          label="Title"
          :value="book.title"
          name="title"></text-input>

        <select-input
          name="author-id"
          v-model="this.book.author_id"
          :items="this.authors"
          required="required"
          label="Author"></select-input>

        <text-input
          v-model="book.publication_year"
          type="text"
          required="true"
          label="Publication Year"
          :value="book.publication_year"
          name="publication-year"></text-input>

        <div class="mb-3">
          <label for="description" class="form-label">Description</label>
          <textarea
            required
            v-model="book.description"
            class="form-control"
            id="description"
            rows="3"></textarea>
        </div>

        <div class="mb-3">
          <label for="genres" class="form-label">Genres</label>
          <select
            ref="genres"
            id="genres"
            class="form-select"
            required
            size="7"
            v-model="this.book.genre_ids"
            multiple>
            <option v-for="g in this.genres" :value="g.value" :key="g.value">
              {{ g.text }}
            </option>
          </select>
        </div>

        <hr />

        <div class="float-start">
          <input type="submit" class="btn btn-primary me-2" value="Save" />
          <router-link to="/admin/books" class="btn btn-outline-secondary">Cancel</router-link>
        </div>
        <div class="float-end mt-2">
          <a
            v-if="this.book.id > 0"
            class="btn btn-danger"
            href="javascript:void(0);"
            @click="confirmDelete(this.book.id)">
            Delete
          </a>
        </div>
      </form-tag>
    </div>
  </div>
</template>

<script>
  import Security from './security.js'
  import FormTag from '@/components/forms/FormTag.vue'
  import TextInput from '@/components/forms/TextInput.vue'
  import SelectInput from '@/components/forms/SelectInput.vue'
  // import notie from 'notie'

  export default {
    name: 'BookEdit',
    beforeMount() {
      Security.requireToken()
    },
    components: {
      'form-tag': FormTag,
      'text-input': TextInput,
      'select-input': SelectInput
    },
    data() {
      return {
        book: {
          id: 0,
          title: '',
          author_id: 0,
          publication_year: 0,
          description: '',
          cover: '',
          slug: '',
          genres: [],
          genre_ids: []
        },
        authors: [],
        imgPath: process.env.VUE_APP_LITERAL_IMAGE_URL,
        genres: [
          { value: 1, text: 'Science Fiction' },
          { value: 2, text: 'Fantasy' },
          { value: 3, text: 'Romance' },
          { value: 4, text: 'Thriller' },
          { value: 5, text: 'Mystery' },
          { value: 6, text: 'Horror' },
          { value: 7, text: 'Classic' }
        ]
      }
    },
    methods: {
      submitHandler: function (data) {
        console.log(data)
      },
      loadCoverImage: function (event) {
        const file = event.target.files[0]
        const reader = new FileReader()
        reader.readAsDataURL(file)
        reader.onload = () => {
          this.book.cover = reader.result
        }
      },
      confirmDelete: {}
    }
  }
</script>
