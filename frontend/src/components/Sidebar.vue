<script setup lang="ts">
import { useRoute } from 'vue-router';
import {
  CreditCard,
  LayoutDashboard,
  BarChart3,
  Settings
} from 'lucide-vue-next'

const route = useRoute()

const menus = [
  {name: 'Transactions', path: '/transactions', icon: CreditCard},
  {name: 'Dashboard', path: '/dashboard', icon: LayoutDashboard},
  {name: 'Analytics', path: '/analytics', icon: BarChart3},
  {name: 'Setting', path: 'setting', icon: Settings}
]

const isActive = (path: string) => route.path === path

</script>

<template>
  <aside class="w-64 h-screen bg-white border-r flex flex-col justify-between">

    <!-- Top -->
    <div>
      <!-- Logo -->
      <div class="px-6 py-5 text-xl font-semibold tracking-tight">
        The Ledger
      </div>

      <!-- Menu -->
      <nav class="px-3 space-y-1">
        <router-link
          v-for="item in menus"
          :key="item.path"
          :to="item.path"
          class="flex items-center gap-3 px-4 py-2.5 rounded-xl text-sm font-medium transition"
          :class="isActive(item.path)
            ? 'bg-green-100 text-green-700'
            : 'text-gray-500 hover:bg-gray-100 hover:text-gray-900'"
        >
          <component :is="item.icon" class="w-5 h-5" />

          <span>{{ item.name }}</span>
        </router-link>
      </nav>
    </div>


  </aside>
</template>