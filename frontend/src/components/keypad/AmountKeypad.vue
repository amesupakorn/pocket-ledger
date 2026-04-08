<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue'

const emit = defineEmits(['update'])

const value = ref('0')
const open = ref(false)

const update = () => {
  emit('update', value.value)
}

const append = (num: string) => {
  if (value.value === '0') value.value = num
  else value.value += num
  update()
}

const remove = () => {
  value.value = value.value.slice(0, -1) || '0'
  update()
}

const handleKey = (e: KeyboardEvent) => {
  if (!open.value) return

  if (/[0-9]/.test(e.key)) append(e.key)
  if (e.key === '.') append('.')
  if (e.key === 'Backspace') remove()
}

onMounted(() => {
  window.addEventListener('keydown', handleKey)
})

onBeforeUnmount(() => {
  window.removeEventListener('keydown', handleKey)
})
</script>

<template>
  <div class="relative">
    
    <!-- display (click เพื่อเปิด keypad) -->
    <div
      @click="open = !open"
      class="bg-gray-100 rounded-xl p-4 text-center text-3xl font-semibold cursor-pointer"
    >
      ฿ {{ value }}
    </div>

    <!-- keypad dropdown -->
    <transition name="fade">
      <div
        v-if="open"
        class="absolute left-0 right-0 mt-2 bg-white rounded-xl shadow-lg p-3 z-10"
      >
        <div class="grid grid-cols-3 gap-2">
          <button
            v-for="n in 9"
            :key="n"
            @click="append(String(n))"
            class="bg-white rounded-lg py-3 text-lg shadow-sm hover:bg-gray-100"
          >
            {{ n }}
          </button>

          <button @click="append('0')" class="bg-white rounded-lg py-3 text-lg shadow-sm hover:bg-gray-100">0</button>
          <button @click="append('.')" class="bg-white rounded-lg py-3 text-lg shadow-sm hover:bg-gray-100">.</button>
          <button @click="remove" class="bg-white rounded-lg py-3 text-lg shadow-sm hover:bg-gray-100">⌫</button>
        </div>
      </div>
    </transition>

  </div>
</template>

<style>

/* animation */
.fade-enter-active,
.fade-leave-active {
  transition: all 0.15s ease;
}
.fade-enter-from {
  opacity: 0;
  transform: translateY(-6px);
}
.fade-leave-to {
  opacity: 0;
  transform: translateY(-6px);
}
</style>