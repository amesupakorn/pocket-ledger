<script setup lang="ts">
import { computed, ref } from 'vue'
import { categoryMap } from '@/constants/categories'
import { Pencil, Clock, Trash2, ArrowDownLeft, ArrowUpRight } from 'lucide-vue-next'
import type { TransactionwithCKey } from '@/types/transaction'
import DeleteModal from '@/components/modal/DeleteTransactionModal.vue'
import { api } from '@/services/api'

const { transaction } = defineProps<{
  transaction: TransactionwithCKey
}>()

const emit = defineEmits(['edit', 'delete', 'click'])

const category = computed(() =>
  categoryMap[transaction.category_key as keyof typeof categoryMap] || categoryMap.other
)

const date = computed(() => new Date(transaction.created_at))

const isExpense = computed(() => transaction.type === 'expense')

const badgeClass = computed(() =>
  isExpense.value
    ? 'bg-red-50 text-red-600 border-red-200'
    : 'bg-green-50 text-green-600 border-green-200'
)

const Icon = computed(() =>
  isExpense.value ? ArrowDownLeft : ArrowUpRight
)

const formattedAmount = computed(() => {
  const sign = isExpense.value ? '-' : '+'
  return `${sign}${Math.abs(transaction.amount).toFixed(2)}`
})


const MAX_SWIPE = 128 // ระยะปุ่ม 2 ปุ่ม (64px * 2)
const startX = ref(0)
const currentOffset = ref(0)
const isSwiping = ref(false)
const isOpen = ref(false)

// เริ่มแตะ
const touchStart = (e: TouchEvent) => {
  startX.value = e.touches[0].clientX
  isSwiping.value = true
}

// ขณะลาก
const touchMove = (e: TouchEvent) => {
  const touchX = e.touches[0].clientX
  let deltaX = touchX - startX.value

  // ถ้าเมนูเปิดอยู่ ให้เริ่มนับจากระยะที่เปิดไว้
  if (isOpen.value) deltaX -= MAX_SWIPE

  // จำกัดการลาก: ลากไปทางซ้ายได้สูงสุด MAX_SWIPE และห้ามลากไปทางขวาเกิน 0
  if (deltaX > 0) deltaX = 0
  if (deltaX < -MAX_SWIPE - 20) deltaX = -MAX_SWIPE - 20

  currentOffset.value = deltaX
}

// ปล่อยนิ้ว
const touchEnd = () => {
  isSwiping.value = false
  // ถ้าลากไปเกินครึ่งทางของปุ่ม ให้เปิดค้างไว้
  if (currentOffset.value < -MAX_SWIPE / 2) {
    currentOffset.value = -MAX_SWIPE
    isOpen.value = true
  } else {
    currentOffset.value = 0
    isOpen.value = false
  }
}

// ฟังก์ชันปิดเมนูเมื่อกดปุ่ม Action
const closeMenu = () => {
  currentOffset.value = 0
  isOpen.value = false
}

// call api

const showDelete = ref(false)

const openDelete = () => {
  showDelete.value = true
}

const closeDelete = () => {
  showDelete.value = false
}

const confirmDelete = async () => {
  await api.delete(`transaction/${transaction.id}`, {
    method: 'DELETE'
  })

  emit('delete', transaction.id)
  closeDelete()
}

</script>
<template>
  
  <div
    class="relative bg-white border border-gray-200 rounded-2xl overflow-hidden group cursor-pointer h-[76px] touch-pan-y"
    @touchstart="touchStart"
    @touchmove="touchMove"
    @touchend="touchEnd"
  >
    <div class="absolute inset-y-0 right-0 flex z-0 border-l border-gray-100">
      <div 
        class="w-[64px] bg-slate-50 flex flex-col items-center justify-center text-slate-500 hover:bg-blue-50 hover:text-blue-600 transition-colors duration-200 cursor-pointer"
        @click.stop="$emit('edit', transaction); closeMenu()"
      >
        <Pencil class="w-4 h-4 mb-1" />
        <span class="text-[10px] font-medium tracking-wide">Edit</span>
      </div>

      <div 
        class="w-[64px] bg-rose-50 flex flex-col items-center justify-center text-rose-500 hover:bg-rose-100 transition-colors duration-200 cursor-pointer"
        @click.stop="openDelete()"
      >
        <Trash2 class="w-4 h-4 mb-1" />
        <span class="text-[10px] font-medium tracking-wide">Delete</span>
      </div>
    </div>
    <div
      @click="isOpen ? closeMenu() : $emit('click', transaction)"
      :style="isSwiping || isOpen ? { marginRight: `${Math.abs(currentOffset)}px` } : {}"
      :class="[
        'relative z-10 bg-white h-full flex items-center justify-between px-5',
        'transition-all duration-300 ease-out flex-nowrap',
        'group-hover:mr-[128px]'
      ]"
    >
      <div class="flex items-center gap-4 min-w-0 flex-1"> 
        
        <div :class="['w-11 h-11 rounded-xl flex items-center justify-center flex-shrink-0', category.color]">
          <component :is="category.icon" class="w-5 h-5" />
        </div>

        <div class="flex flex-col min-w-0 flex-1">
          <p class="font-semibold text-gray-800 truncate leading-tight">
            {{ category.label }}
          </p>
          <p class="text-xs text-gray-500 truncate mt-0.5">
            {{ transaction.note }}
          </p>
        </div>
      </div>
      <div class="flex flex-col items-end gap-1 ml-3 flex-shrink-0">
        <div :class="['flex gap-1 px-2 py-0.5 rounded-full border text-[9px] font-bold uppercase', badgeClass]">
          <component :is="Icon" class="w-3 h-3 flex-shrink-0" />
          <span>{{ isExpense ? 'Expense' : 'Income' }}</span>
        </div>
        <p :class="['font-bold text-base leading-none', isExpense ? 'text-red-600' : 'text-green-600']">
          {{ formattedAmount }}
        </p>
      </div>

  
    </div>
  </div>


  <DeleteModal
    :open="showDelete"
    title="Delete Transaction"
    description="This action cannot be undone."
    @close="closeDelete"
    @confirm="confirmDelete"
  />

  
</template>