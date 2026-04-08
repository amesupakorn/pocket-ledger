<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { categoryMap } from '@/constants/categories'
import AmountKeypad from '../keypad/AmountKeypad.vue'
import CategoryDropdown from '../dropdown/CategoryDropdown.vue'
import DatePicker from '../picker/DatePicker.vue'
import { api } from '@/services/api'
import { useCategories } from '@/composables/useCategories'

defineProps<{
  open: boolean
}>()

const emit = defineEmits(['close', 'create'])

const type = ref<'expense' | 'income'>('expense')

const loading = ref(false)

const { mapCategories, loadingCate } = useCategories()

const categories = computed(() =>
  mapCategories(type.value)
)
const amount = ref(0)
const selectedCategory = ref<any>(null)
const note = ref('')
const walletId = 1
const selectedDate = ref<Date | null>(new Date())
//call api

const handleAmountUpdate = (v: string) => {
  amount.value = Number(v)
}

const handleDateUpdate = (d: Date) => {
  selectedDate.value = d
}

const handleCategorySelect = (c: any) => {
  selectedCategory.value = c
}

const  createTransaction = async () => {
    try{
        if(!selectedCategory.value) {
            alert('Please select catrgory')
            return
        }

        if(amount.value <= 0){
            alert("Please input your amount")
            return
        }

        loading.value = true

        const res = await api.post('/create/transaction', {
            wallet_id: walletId,
            category_id: selectedCategory.value.id,
            amount: amount.value,
            type: type.value,
            note: note.value,
            created_at: selectedDate.value?.toISOString(),
        })

        emit('create', res.data.data)
        
        amount.value = 0
        note.value = ''
        selectedCategory.value = null
        emit('close')

    } catch (err) {
        console.error(err)
        alert('Error saving transactions')
    } finally {
        loading.value = false
    }
}

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
         <AmountKeypad @update="handleAmountUpdate" />
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
            @select="handleCategorySelect"
        />


      <!-- date -->
       <DatePicker @update="handleDateUpdate" />


      <!-- note -->
      <div class="mb-4 mt-1">
        <label class="text-xs text-gray-500">NOTE</label>
        <input
          v-model="note"
          type="text"
          placeholder="What was this for?"
          class="w-full border border-gray-200 bg-white rounded-xl px-4 py-2.5 flex items-center justify-between shadow-sm hover:border-gray-300 transition"

        />
      </div>

      <!-- actions -->
      <div class="flex gap-2">
        <button 
          @click="createTransaction"
          :disabled="loading"
          class="flex-1 bg-green-600 text-white py-2 rounded-lg">
          {{ loading ? 'Create....' : 'Create Transaction' }}
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

