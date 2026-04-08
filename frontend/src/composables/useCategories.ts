import { ref, computed, onMounted } from 'vue'
import { api } from '@/services/api'
import { categoryMap } from '@/constants/categories'

export const useCategories = () => {
  const rawCategories = ref<
    Array<{
      id: number
      key: keyof typeof categoryMap
      type: 'expense' | 'income'
    }>
  >([])
  const loading = ref(false)

  const fetchCategories = async () => {
    loading.value = true
    try {
      const res = await api.get('/categories')
      rawCategories.value = res.data
    } catch (err) {
      console.error('fetch categories error', err)
    } finally {
      loading.value = false
    }
  }

  onMounted(fetchCategories)

  const mapCategories = (type: 'expense' | 'income') => {
    return rawCategories.value
      .filter(c => c.type === type)
      .map(c => {
        const map = categoryMap[c.key] || categoryMap.other

        return {
          id: c.id,
          value: c.key,
          label: map.label,
          icon: map.icon,
          color: map.color,
        }
      })
  }

  return {
    rawCategories,
    loading,
    mapCategories,
  }
}