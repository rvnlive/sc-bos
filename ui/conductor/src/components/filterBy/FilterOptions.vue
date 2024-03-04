<template>
  <v-list class="px-0 mx-0 mt-5 overflow-y-auto" dense max-height="300">
    <v-subheader class="mt-n4 mb-1 ml-n2 text-caption neutral sticky-wrapper">
      <span v-if="selectedOption !== null">
        {{ selectedOption.title }}
      </span>
      <v-btn
          :block="selectedOption === null"
          class="ml-auto mr-n2 pt-1 px-1 rounded"
          :disabled="filterByStore.clearDisabled"
          small
          text
          width="auto"
          @click="filterByStore.doRemoveFilter">
        {{ clearResetButtonLabels }}
      </v-btn>
    </v-subheader>

    <FilterList v-if="!selectedOption"/>
    <FilterList v-else-if="selectedOption?.type === 'list'" list-type="option"/>
    <FilterRangeSlider v-if="selectedOption && selectedOption?.type === 'range'"/>
    <FilterRadioGroup v-else-if="selectedOption?.type === 'radio'"/>
  </v-list>
</template>
<script setup>
import FilterList from '@/components/filterBy/filterTypes/FilterList.vue';
import FilterRadioGroup from '@/components/filterBy/filterTypes/FilterRadioGroup.vue';
import FilterRangeSlider from '@/components/filterBy/filterTypes/FilterRangeSlider.vue';
import {useFilterByStore} from '@/components/filterBy/useFilterByStore';
import {storeToRefs} from 'pinia';


const filterByStore = useFilterByStore();
const {selectedOption, clearResetButtonLabels} = storeToRefs(filterByStore);
</script>

<style scoped>
.sticky-wrapper {
  position: sticky;
  top: -16px; /* Adjust this value based on your needs */
  z-index: 1000; /* Ensure the sticky element stays above other content */
}
</style>

