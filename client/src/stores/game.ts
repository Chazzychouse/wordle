import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useGameStore = defineStore('game', () => {
  const isLoading = ref(false)

  return {
    isLoading,
  }
})
