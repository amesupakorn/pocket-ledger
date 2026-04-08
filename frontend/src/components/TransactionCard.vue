<script setup lang="ts">
import { categoryMap } from '@/constants/categories'
import type { Category } from '@/types/transaction'

const transaction = {
  title: 'Artisan Coffee',
  amount: -42.5,
  category: 'food' as Category
}
</script>

<template>
  <div class="bg-white border border-gray-300 rounded-xl px-6 py-4 flex justify-between items-center">

    <!-- Left -->
    <div class="flex items-center gap-4">
      
      <div
        :class="[
          'w-10 h-10 rounded-lg flex items-center justify-center',
          categoryMap[transaction.category].color
        ]"
      >
        <component
          :is="categoryMap[transaction.category].icon"
          class="w-5 h-5"
        />
      </div>

      <div>
        <p class="font-medium">{{ transaction.title }}</p>
        <p class="text-sm text-gray-500">
          {{ categoryMap[transaction.category].label }}
        </p>
      </div>

    </div>

    <!-- Right -->
    <div class="text-right">
      <p class="text-red-500 font-semibold">
        {{ transaction.amount }}
      </p>
    </div>

  </div>
</template>