<script setup lang="ts">
import { ref } from 'vue'

const props = defineProps<{
  categories: any[]
}>()

const emit = defineEmits(['select'])

const open = ref(false)
const selected = ref<any>(null)

const selectItem = (item: any) => {
  selected.value = item
  emit('select', item)
  open.value = false
}
</script>

<template>
  <div class="relative w-full">
    
    <!-- trigger -->
    <button
      @click="open = !open"
      class="w-full border border-gray-200 bg-white rounded-xl px-4 py-2.5 flex items-center justify-between shadow-sm hover:border-gray-300 transition"
    >
      <div v-if="selected" class="flex items-center gap-2">
        <div
          class="p-1.5 rounded-md"
          :class="selected.color"
        >
          <component :is="selected.icon" class="w-4 h-4" />
        </div>
        <span class="font-medium text-gray-800">
          {{ selected.label }}
        </span>
      </div>

      <span v-else class="text-gray-400">
        Select category
      </span>

      <!-- arrow -->
      <svg
        class="w-4 h-4 text-gray-400 ml-2"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        viewBox="0 0 24 24"
      >
        <path d="M19 9l-7 7-7-7" />
      </svg>
    </button>

    <!-- dropdown -->
    <transition name="fade">
      <div
        v-if="open"
        class="absolute mt-2 w-full bg-white border border-gray-200 rounded-xl shadow-lg z-10 overflow-hidden"
      >
        <div
          v-for="item in categories"
          :key="item.value"
          @click="selectItem(item)"
          class="flex items-center gap-3 px-4 py-2.5 cursor-pointer transition hover:bg-gray-50"
        >
          <div
            class="p-1.5 rounded-md"
            :class="item.color"
          >
            <component :is="item.icon" class="w-4 h-4" />
          </div>

          <span class="text-gray-700 font-medium">
            {{ item.label }}
          </span>
        </div>
      </div>
    </transition>

  </div>
</template>

<style>
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