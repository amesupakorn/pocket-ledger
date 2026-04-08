<script setup lang="ts">
import MainLayout from '@/layouts/MainLayout.vue'
import FilterBar from '@/components/FilterBar.vue'
import TransactionCard from '@/components/TransactionCard.vue'
import AddTransactionModal from '@/components/modal/AddTransactionModal.vue'
import { ref } from 'vue'
import { useTransactions } from '@/composables/useTransactions'
import type { TransactionwithCKey } from '@/types/transaction'

const { transactions } = useTransactions()

const showAdd = ref(false)

const handleAdd = (newItem: TransactionwithCKey) => {
  transactions.value.unshift(newItem)
  showAdd.value = false
}

const handleDelete = (id: number) => {
  transactions.value = transactions.value.filter(t => t.id !== id)
}

</script>

<template>
  <MainLayout>
    
    <FilterBar @add="showAdd = true"/>

    <div class="space-y-4 mt-4">
      <TransactionCard
        v-for="t in transactions"
        :key="t.id"
        :transaction="t"
        @delete="handleDelete"
      />
    </div>

    <AddTransactionModal
      :open="showAdd"
      @close="showAdd = false"
      @create="handleAdd"
    />

  </MainLayout>
</template>