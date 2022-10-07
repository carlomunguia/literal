<template>
  <div class="mb-3">
    <label :for="name" class="form-label">{{ label }}</label>
    <select
      :id="name"
      :required="required"
      :name="name"
      class="form-select"
      :multiple="multiple"
      :value="modelValue"
      @input="$emit('update:modelValue', $event.target.value)">
      <option disabled value="">Choose...</option>
      <option v-for="option in items" :value="option.value" :key="option.value">
        {{ option.text }}
      </option>
    </select>
  </div>
</template>

<script>
  export default {
    name: 'SelectInput',
    props: ['items', 'name', 'requred', 'label', 'modelValue', 'multiple'],
    emits: ['update:modelValue']
  }
</script>
