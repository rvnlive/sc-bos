<template>
  <v-text-field
      autofocus
      clearable
      class="elevation-0"
      color="primary"
      dense
      flat
      :label="filterInputLabel"
      hide-details
      outlined
      :placeholder="filterInputLabel"
      single-line
      solo
      v-model="filterInput"/>
</template>

<script setup>
import {computed} from 'vue';
import {useFilterByStore} from '@/components/filterBy/useFilterByStore.js';
import {storeToRefs} from 'pinia';

const filterByStore = useFilterByStore();
const {filterInput, selectedOption} = storeToRefs(filterByStore);

const filterInputLabel = computed(() => {
  if (!selectedOption.value?.title) return 'Filter by...';
  else return 'Filter by ' + selectedOption.value.title.toLowerCase();
});
</script>

<style scoped>

</style>
