<script setup lang="ts">
import { ref } from 'vue'
import Sidebar from '@/components/Sidebar.vue'
import Topbar from '@/components/Topbar.vue'
import AddTransactionModal from '@/components/modal/AddTransactionModal.vue'

const isOpen = ref(false)
</script>

<template>
  <div class="flex h-screen">

    <!-- Sidebar -->
    <Sidebar @add="isOpen = true" />

    <div class="flex-1 flex flex-col">
      <Topbar />
      <main class="p-6 flex-1 overflow-y-auto">
        <slot />
      </main>
    </div>

    <!-- Modal -->
    <AddTransactionModal
      :open="isOpen"
      @close="isOpen = false"
    />

  </div>
</template>