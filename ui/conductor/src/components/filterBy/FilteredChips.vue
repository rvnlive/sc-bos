<template>
  <v-row class="mr-0 d-flex flex-row align-center">
    <v-sheet class="px-1" elevation="0" width="100%" max-width="1400">
      <span v-for="filter in availableFilters" :key="filter.title">
        <FilterChipMenu
            v-if="filter.showChip === true"
            :filter="filter"/>
      </span>
    </v-sheet>
  </v-row>
</template>

<script setup>
import FilterChipMenu from '@/components/filterBy/FilterChipMenu.vue';
import {storeToRefs} from 'pinia';
import {useFilterByStore} from '@/components/filterBy/useFilterByStore.js';

const {availableFilters} = storeToRefs(useFilterByStore());
</script>
