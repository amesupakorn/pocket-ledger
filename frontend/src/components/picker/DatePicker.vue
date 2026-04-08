<script setup lang="ts">
import { ref } from 'vue'
import { VueDatePicker } from '@vuepic/vue-datepicker'
import '@vuepic/vue-datepicker/dist/main.css'

const date = ref(new Date())
</script>

<template>
  <div>
    <label class="text-xs text-gray-500">DATE</label>

    <VueDatePicker v-model="date" />

  </div>
</template>