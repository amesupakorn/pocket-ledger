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

export const categoryMap = {
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
    icon: Briefcase,
    color: 'bg-emerald-100 text-emerald-600'
  },
  freelance: {
    label: 'Freelance',
    icon: DollarSign,
    color: 'bg-green-100 text-green-600'
  },
  investment: {
    label: 'Investment',
    icon: TrendingUp,
    color: 'bg-blue-100 text-blue-600'
  },
  other: {
    label: 'Other',
    icon: CircleHelp,
    color: 'bg-gray-100 text-gray-600'
  }
} as const