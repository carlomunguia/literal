<template>
  <form
    @submit.prevent="submit"
    :ref="name"
    :event="event"
    autocomplete="off"
    :method="method"
    :action="action"
    class="needs-validation"
    novalidate>
    <slot></slot>
  </form>
</template>

<script>
  export default {
    name: 'FormTag',
    props: ['method', 'action', 'event', 'name'],
    methods: {
      submit() {
        let vueForm = this.$refs[this.$props.name]

        if (vueForm.checkValidity()) {
          this.$emit(this.$props.event)
        } else {
          vueForm.classList.add('was-validated')
        }
      }
    }
  }
</script>
