<script setup lang="ts">
import { useClipboard, useMousePressed } from '@vueuse/core';
import { useTemplateRef } from 'vue';

const { copy } = useClipboard()

const badge = useTemplateRef('badge')
const { pressed } = useMousePressed({ target: badge })

function copyValue() {
  const text = badge.value?.textContent
  if (!text) {
    return
  }

  copy(text)
}
</script>

<template>
  <span class="badge rounded-pill px-2" :class="{ 'text-bg-secondary': !pressed, 'text-bg-success': pressed }"
    @click="copyValue" ref="badge">
    <slot>placeholder</slot>
  </span>
</template>

<style scoped>
.badge {
  cursor: pointer;
  transition: background-color .2s;
  user-select: none;
}
</style>
