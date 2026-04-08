<script setup lang="ts">
import { Search, Plus } from 'lucide-vue-next'

</script>

<template>
  <header class="h-16 bg-white flex items-center justify-between px-6 border-b border-gray-100">

    <!-- Left -->
    <div class="flex items-center gap-3">
      <h1 class="text-lg font-semibold text-gray-800">
        Transactions
      </h1>
    </div>

    <!-- Right -->
    <div class="flex items-center gap-3">

      <!-- Search -->
      <div class="relative">
        <Search class="w-4 h-4 text-gray-400 absolute left-3 top-1/2 -translate-y-1/2" />

        <input
          type="text"
          placeholder="Search transactions..."
          class="border border-gray-200 rounded-xl pl-9 pr-3 py-1.5 text-sm focus:outline-none focus:ring-2 focus:ring-gray-100 w-48 transition"
        />
      </div>

    </div>

  </header>
</template>