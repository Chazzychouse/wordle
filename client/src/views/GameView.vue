<script setup lang="ts">
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useGameStore } from '../stores/game'

const router = useRouter()
const gameStore = useGameStore()

onMounted(async () => {
  try {
    
  } catch (error) {
    console.error('Failed to initialize game:', error)
    router.push('/')
  }
})

</script>

<template>
  <div class="game-container">
    <div class="game-header">
      <h1>Definitely Not Wordle</h1>
      <div v-if="gameStore.isLoading" class="loading-indicator">
        Loading...
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>