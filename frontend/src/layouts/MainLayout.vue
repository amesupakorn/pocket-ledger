<script setup lang="ts">
import { ref } from 'vue'
import Sidebar from '@/components/Sidebar.vue'
import Topbar from '@/components/Topbar.vue'

</script>

<template>
  <div class="flex h-screen">

    <!-- Sidebar -->
    <Sidebar />

    <div class="flex-1 flex flex-col">
      <Topbar/>
      <main class="p-6 flex-1 overflow-y-auto">
        <slot />
      </main>
    </div>

  </div>
</template>