<template>
  <div>
    <v-tooltip top>
      <template #activator="{ on }">
        <v-btn
            v-on="on"
            class="rounded"
            color="neutral"
            elevation="0"
            fab
            height="36"
            tile
            small
            :width="filterByStore.defaultFilters.length > 0 ? '36' : '30'"
            @click="showHideFilterMenu">
          <v-badge
              v-if="filterByStore.defaultFilters.length > 0"
              color="red"
              value="true"
              dot
              offset-x="5"
              offset-y="8">
            <v-icon>mdi-filter</v-icon>
          </v-badge>
        </v-btn>
      </template>
      <span>Set Filter</span>
    </v-tooltip>
  </div>
</template>

<script setup>
import {useFilterByStore} from '@/components/filterBy/useFilterByStore.js';
import {storeToRefs} from 'pinia';

const filterByStore = useFilterByStore();
const {showFilterMenu} = storeToRefs(filterByStore);

const showHideFilterMenu = () => {
  filterByStore.clearSelection();
  showFilterMenu.value = !showFilterMenu.value;
};
</script>

<style scoped>

</style>
