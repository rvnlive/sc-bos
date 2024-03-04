<template>
  <ContentCard class="pb-0 pb-1 mb-0">
    <v-row class="px-2 pt-3 pb-0 align-center sticky-wrapper">
      <v-tooltip v-if="selectedOption && !selectedFilterChip" left>
        <template #activator="{ on }">
          <v-btn
              class="px-4 mr-1 rounded"
              height="40"
              icon
              v-on="on"
              width="22"
              @click="filterByStore.clearSelection">
            <v-icon>
              mdi-arrow-left
            </v-icon>
          </v-btn>
        </template>
        <span>Back</span>
      </v-tooltip>
      <FilterInput v-show="!['range'].includes(selectedOption?.type)"/>
    </v-row>
    <FilterOptions/>
  </ContentCard>
</template>

<script setup>
import ContentCard from '@/components/ContentCard.vue';
import FilterInput from '@/components/filterBy/FilterInput.vue';
import FilterOptions from '@/components/filterBy/FilterOptions.vue';
import {useFilterByStore} from '@/components/filterBy/useFilterByStore.js';
import {storeToRefs} from 'pinia';

const filterByStore = useFilterByStore();
const {selectedFilterChip, selectedOption} = storeToRefs(filterByStore);
</script>

<style scoped>
.sticky-wrapper {
  position: sticky;
  top: 0; /* Adjust this value based on your needs */
  z-index: 1000; /* Ensure the sticky element stays above other content */
}
</style>
