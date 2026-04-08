import { createRouter, createWebHistory } from 'vue-router'
import TransactionsView from '@/views/TransactionsView.vue'

const routes = [
  { path: '/', component: TransactionsView },
  { path: '/transactions', component: TransactionsView },
  { path: '/analytics', component: TransactionsView },
]

export const router = createRouter({
  history: createWebHistory(),
  routes,
})