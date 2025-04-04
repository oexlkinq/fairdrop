<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <main class="island">
    <div class="grid-vstack earth">
      <uv class="nav nav-pills nav-justified">
        <li class="nav-item">
          <RouterLink class="nav-link" active-class="active" to="/upload">Создать</RouterLink>
        </li>
        <li class="nav-item">
          <RouterLink class="nav-link" active-class="active" to="/">Открыть</RouterLink>
        </li>
      </uv>

      <Suspense>
        <RouterView />
      </Suspense>
    </div>
  </main>
</template>

<style scoped>
.earth {
  width: 25rem;
  min-height: 15rem;
  align-content: start;
}
</style>
