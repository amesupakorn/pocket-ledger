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