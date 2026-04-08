<script setup lang="ts">
import { ref, watch } from 'vue'
import { VueDatePicker } from '@vuepic/vue-datepicker'
import '@vuepic/vue-datepicker/dist/main.css'

const emit = defineEmits<{
    (e: 'update', value:Date): void
}>()

const date = ref(new Date())

watch(date, (val) => {
    emit('update', val)
})

</script>

<template>
  <div class="w-full pt-4">
    <label class="text-xs text-gray-500 mb-1 block">DATE</label>

    <div class="relative">
      <VueDatePicker
        v-model="date"
        :enable-time-picker="false"
        auto-apply
        format="dd MMM yyyy"
        input-class-name="custom-input py-2.5"
      />

    </div>
  </div>
</template>

<style>
/* input style */
.custom-input {
  width: 100%;
  border-radius: 0.75rem;
  border: 1px solid #e5e7eb;
  padding: 10px 40px 10px 12px;
  font-size: 12px;
  transition: all 0.2s ease;
  background: white;
}

.custom-input:focus {
  outline: none;
  border-color: #9ca3af;
  box-shadow: 0 0 0 2px rgba(0,0,0,0.05);
}

/* popup calendar */
.dp__theme_light {
  --dp-border-radius: 12px;
  --dp-font-family: inherit;
}

/* hover / selected */
.dp__cell_inner:hover {
  background: #f3f4f6;
}

.dp__active_date {
  background: #111827;
  color: white;
  border-radius: 8px;
}
</style>