import { ref, onMounted } from 'vue'
import { api } from '@/services/api'
import type { TransactionwithCKey } from '@/types/transaction'
 
const transactions = ref<TransactionwithCKey[]>([])
const loading = ref(false)
const error = ref<string | null>(null)

export const useTransactions = () => {

    const fetchTransactions = async () => {
        loading.value = true
        error.value = null

        try {
            const res = await api.get('/transactions')
            transactions.value = res.data
        } catch (err: any) {
            console.error(err)
            error.value = 'Failed to load'
        } finally {
            loading.value = false
        }
    }

    onMounted(fetchTransactions)

    return {
        transactions,
        loading,
        error,
        fetchTransactions,
    }
}