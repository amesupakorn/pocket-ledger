<script setup lang="ts">
import { ref } from 'vue'

const props = defineProps<{
  open: boolean
  title?: string
  description?: string
}>()

const emit = defineEmits(['close', 'confirm'])

const loading = ref(false)

const handleClose = () => {
  if (loading.value) return
  emit('close')
}

const handleConfirm = async () => {
  if (loading.value) return

  try {
    loading.value = true
    await emit('confirm')
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div v-if="open" class="fixed inset-0 z-50 flex items-center justify-center">

    <!-- backdrop (blur อย่างเดียว) -->
    <div
      class="absolute inset-0 bg-black/20 backdrop-blur-sm"
      @click="handleClose"
    />

    <!-- modal -->
    <div class="relative bg-white rounded-2xl p-6 w-[320px] shadow-lg">

      <h3 class="text-lg font-semibold text-gray-800 mb-2">
        {{ title || 'Delete item' }}
      </h3>

      <p class="text-sm text-gray-500 mb-5">
        {{ description || 'Are you sure?' }}
      </p>

      <div class="flex justify-end gap-2">
        <button
          @click="handleClose"
          :disabled="loading"
          class="px-3 py-1.5 text-sm rounded-lg bg-gray-100 hover:bg-gray-200 disabled:opacity-50"
        >
          Cancel
        </button>

        <button
          @click="handleConfirm"
          :disabled="loading"
          class="px-3 py-1.5 text-sm rounded-lg bg-red-500 text-white hover:bg-red-600 disabled:opacity-50"
        >
          {{ loading ? 'Deleting...' : 'Delete' }}
        </button>
      </div>

    </div>

  </div>
</template>