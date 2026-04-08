import {
  Utensils,
  ShoppingCart,
  Car,
  ShoppingBag,
  DollarSign,
  Briefcase,
  TrendingUp,
  CircleHelp
} from 'lucide-vue-next'

import type { Category } from '@/types/transaction'

export const categoryMap: Record<
  Category,
  {
    label: string
    icon: any
    color: string
  }
> = {
  food: {
    label: 'Food & Dining',
    icon: Utensils,
    color: 'bg-orange-100 text-orange-600'
  },
  groceries: {
    label: 'Groceries',
    icon: ShoppingCart,
    color: 'bg-green-100 text-green-600'
  },
  transport: {
    label: 'Transport',
    icon: Car,
    color: 'bg-blue-100 text-blue-600'
  },
  shopping: {
    label: 'Shopping',
    icon: ShoppingBag,
    color: 'bg-purple-100 text-purple-600'
  },
  salary: {
    label: 'Salary',
    icon: DollarSign,
    color: 'bg-emerald-100 text-emerald-600'
  },
  other: {
    label: 'Other',
    icon: CircleHelp,
    color: 'bg-gray-100 text-gray-600'
  }
}

export const expenseCategories = [
  {
    value: 'food',
    label: 'Food & Dining',
    icon: Utensils,
    color: 'bg-orange-100 text-orange-600'
  },
  {
    value: 'groceries',
    label: 'Groceries',
    icon: ShoppingCart,
    color: 'bg-green-100 text-green-600'
  },
  {
    value: 'transport',
    label: 'Transport',
    icon: Car,
    color: 'bg-blue-100 text-blue-600'
  },
  {
    value: 'shopping',
    label: 'Shopping',
    icon: ShoppingBag,
    color: 'bg-purple-100 text-purple-600'
  },
  {
    value: 'other',
    label: 'other',
    icon: CircleHelp,
    color: 'bg-gray-100 text-gray-600'
  }
]

export const incomeCategories = [
  {
    value: 'salary',
    label: 'Salary',
    icon: Briefcase,
    color: 'bg-emerald-100 text-emerald-600'
  },
  {
    value: 'freelance',
    label: 'Freelance',
    icon: DollarSign,
    color: 'bg-green-100 text-green-600'
  },
  {
    value: 'investment',
    label: 'Investment',
    icon: TrendingUp,
    color: 'bg-blue-100 text-blue-600'
  }
]