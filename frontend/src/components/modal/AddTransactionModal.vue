<script setup lang="ts">
import { ref, computed } from 'vue'
import { expenseCategories, incomeCategories } from '@/constants/categories'
import AmountKeypad from '../keypad/AmountKeypad.vue'
import CategoryDropdown from '../dropdown/CategoryDropdown.vue'
import DatePicker from '../picker/DatePicker.vue'

defineProps<{
  open: boolean
}>()

const emit = defineEmits(['close'])

const type = ref<'expense' | 'income'>('expense')

const categories = computed(() =>
  type.value === 'expense' ? expenseCategories : incomeCategories
)

</script>


<template>
  <div v-if="open" class="fixed inset-0 z-50 flex items-center justify-center">
    
    <!-- overlay -->
    <div
      class="absolute inset-0 bg-black/30 backdrop-blur-sm"
      @click="emit('close')"
    />

    <!-- modal -->
    <div class="relative bg-white w-full max-w-md rounded-2xl p-6 shadow-xl">
      
      <!-- header -->
      <div class="flex justify-between items-center mb-4">
        <h2 class="text-lg font-semibold">Add Transaction</h2>
        <button @click="emit('close')" class="text-gray-400 hover:text-gray-600">
          ✕
        </button>
      </div>

      <!-- amount -->
      <div class="bg-gray-100 rounded-xl p-4 text-center mb-4">
        <p class="text-sm text-gray-500">TRANSACTION AMOUNT</p>
         <AmountKeypad @update="(v) => console.log(v)" />

      </div>

      <!-- type -->
     <div class="flex bg-gray-100 rounded-lg mb-4 p-1">
        <button
            @click="type = 'expense'"
            :class="[
            'flex-1 py-2 text-sm rounded-lg transition',
            type === 'expense'
                ? 'bg-white shadow text-green-900'
                : 'text-gray-500'
            ]"
        >
            Expense
        </button>

        <button
            @click="type = 'income'"
            :class="[
            'flex-1 py-2 text-sm rounded-lg transition',
            type === 'income'
                ? 'bg-white shadow text-green-900'
                : 'text-gray-500'
            ]"
        >
            Income
        </button>
     </div>

      <!-- category -->
       <CategoryDropdown
            :categories="categories"
            @select="(c) => console.log(c)"
        />


      <!-- date -->
       <DatePicker />


      <!-- note -->
      <div class="mb-4">
        <label class="text-xs text-gray-500">NOTE</label>
        <input
          type="text"
          placeholder="What was this for?"
          class="w-full border border-gray-300 rounded-lg px-3 py-2 mt-1"
        />
      </div>

      <!-- actions -->
      <div class="flex gap-2">
        <button class="flex-1 bg-green-600 text-white py-2 rounded-lg">
          Save Transaction
        </button>
        <button
          @click="emit('close')"
          class="flex-1 bg-gray-200 py-2 rounded-lg"
        >
          Cancel
        </button>
      </div>

    </div>
  </div>
</template>

