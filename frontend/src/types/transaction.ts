export type TransactionType = 'income' | 'expense'

export type Category =
  | 'food'
  | 'groceries'
  | 'transport'
  | 'shopping'
  | 'salary'
  | 'other'

export interface Transaction {
  id: string
  title: string
  amount: number
  category: Category
  type: TransactionType
  createdAt: string
}

export interface TransactionwithCKey {
  id: number
  amount: number
  type: string
  note: string
  created_at: string
  category_key: string
}