<template>
  <MainLayout>
    
    <FilterBar />

    <div class="space-y-4 mt-4">
      <TransactionCard />
      <TransactionCard />
    </div>

  </MainLayout>
</template>

<script setup lang="ts">
import MainLayout from '@/layouts/MainLayout.vue'
import FilterBar from '@/components/FilterBar.vue'
import TransactionCard from '@/components/TransactionCard.vue'
</script>