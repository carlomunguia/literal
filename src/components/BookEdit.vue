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
      </form-tag>
    </div>
  </div>
</template>

<script>
  import Security from './security.js'
  import FormTag from '@/components/forms/FormTag.vue'
  import TextInput from '@/components/forms/TextInput.vue'
  import SelectInput from '@/components/forms/SelectInput.vue'
  import notie from 'notie'

  export default {
    name: 'BookEdit',
    beforeMount() {
      Security.requireToken()
    },
    methods: {
      submitHandler: function (data) {
        console.log(data)
      }
    }
  }
</script>
