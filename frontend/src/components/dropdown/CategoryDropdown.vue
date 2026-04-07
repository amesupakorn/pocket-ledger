<script setup lang="ts">
import { ref, computed } from 'vue'

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
  <div class="relative">
    
    <!-- trigger -->
    <button
      @click="open = !open"
      class="w-full border rounded-lg border-gray-300 px-3 py-2 flex items-center justify-between"
    >
      <div v-if="selected" class="flex items-center gap-2">
        <component :is="selected.icon" class="w-4 h-4" />
        <span>{{ selected.label }}</span>
      </div>
      <span v-else class="text-gray-400">Select category</span>
    </button>

    <!-- dropdown -->
    <transition name="fade">
      <div
        v-if="open"
        class="absolute mt-1 w-full bg-white border rounded-lg shadow z-10"
      >
        <div
          v-for="item in categories"
          :key="item.value"
          @click="selectItem(item)"
          class="flex items-center gap-3 px-3 py-2 hover:bg-gray-100 cursor-pointer"
        >
          <component :is="item.icon" class="w-4 h-4" />
          <span>{{ item.label }}</span>
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
  transform: translateY(-4px);
}
.fade-leave-to {
  opacity: 0;
  transform: translateY(-4px);
}
</style>