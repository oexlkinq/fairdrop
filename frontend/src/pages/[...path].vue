<script setup lang="ts">
import { useRoute } from 'vue-router';

const route = useRoute('/[...path]')
</script>

<template>404. route <span>"{{ route.params.path }}"</span> not found</template>
