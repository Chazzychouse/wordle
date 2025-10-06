<script setup lang="ts">
import { useRouter } from 'vue-router'

const router = useRouter()

const startNewGame = () => {
  router.push('/game')
}
</script>

<template>
  <div class="home-container">
    <div class="welcome-section">
      <h1 class="welcome-title">Definitely Not Wordle</h1>
      <p class="welcome-description">
        Guess the 5-letter word in 6 tries. Each guess must be a valid word. 
        The color of the tiles will change to show how close your guess was to the word.
      </p>
      
      <button @click="startNewGame" class="start-game-btn">
        Start New Game
      </button>
    </div>
  </div>
</template>

<style scoped>
.home-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 60vh;
  padding: 2rem;
}

.welcome-section {
  text-align: center;
  max-width: 600px;
  background: white;
  padding: 3rem;
  border-radius: 1rem;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.welcome-title {
  font-size: 3rem;
  font-weight: bold;
  color: #212529;
  margin-bottom: 1.5rem;
  letter-spacing: 0.1em;
}

.welcome-description {
  font-size: 1.2rem;
  color: #6c757d;
  margin-bottom: 2rem;
  line-height: 1.6;
}

.start-game-btn {
  background-color: #28a745;
  color: white;
  border: none;
  padding: 1rem 2rem;
  font-size: 1.2rem;
  font-weight: bold;
  border-radius: 0.5rem;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.start-game-btn:hover {
  background-color: #218838;
}

@media (max-width: 768px) {
  .welcome-section {
    padding: 2rem;
    margin: 1rem;
  }
  
  .welcome-title {
    font-size: 2.5rem;
  }
  
  .welcome-description {
    font-size: 1.1rem;
  }
}
</style>
