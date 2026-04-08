<script setup lang="ts">
import { Calendar, Filter, Plus } from 'lucide-vue-next'

const emit = defineEmits(['add'])
</script>

<template>
  <div class="bg-white p-4 rounded-2xl shadow-sm flex flex-wrap gap-3 items-center">

    <!-- Month -->
    <div class="relative">
      <Calendar class="w-4 h-4 text-gray-400 absolute left-3 top-1/2 -translate-y-1/2" />

      <select class="border border-gray-200 rounded-xl pl-9 pr-3 py-2 text-sm bg-white focus:outline-none focus:ring-2 focus:ring-gray-100">
        <option>September 2023</option>
      </select>
    </div>

    <!-- Category -->
    <div class="relative">
      <Filter class="w-4 h-4 text-gray-400 absolute left-3 top-1/2 -translate-y-1/2" />

      <select class="border border-gray-200 rounded-xl pl-9 pr-3 py-2 text-sm bg-white focus:outline-none focus:ring-2 focus:ring-gray-100">
        <option>All Categories</option>
      </select>
    </div>

    <!-- Type Filter -->
    <div class="flex gap-2 ml-auto">

      <button
        class="px-3 py-1.5 rounded-xl text-sm font-medium bg-gray-900 text-white shadow-sm"
      >
        All
      </button>

      <button
        class="px-3 py-1.5 rounded-xl text-sm font-medium bg-gray-100 text-gray-600 hover:bg-gray-200 transition"
      >
        Income
      </button>

      <button
        class="px-3 py-1.5 rounded-xl text-sm font-medium bg-gray-100 text-gray-600 hover:bg-gray-200 transition"
      >
        Expense
      </button>

    </div>

      <!-- Add button -->
    <button 
        @click="$emit('add')"
        class="flex items-center gap-1 px-3 py-1.5 rounded-xl bg-green-600 text-white text-sm font-medium hover:opacity-90 transition">

        <Plus class="w-4 h-4" />
        Add
    </button>


  </div>
</template>