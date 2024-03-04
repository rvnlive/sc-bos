<template>
  <v-col cols="auto" class="d-flex flex-row align-center py-0">
    <FilteredChips/>
    <v-menu
        content-class="mt-2"
        :close-on-content-click="false"
        left
        max-width="290"
        min-width="290"
        nudge-bottom="0"
        nudge-left="0"
        nudge-right="0"
        nudge-top="0"
        offset-y
        return-value="true"
        v-model="showFilterMenu">
      <template #activator="{ on }">
        <FilterButton v-on="on"/>
      </template>
      <FilterMenu/>
    </v-menu>
  </v-col>
</template>
<script setup>
import FilterButton from '@/components/filterBy/FilterButton.vue';
import FilteredChips from '@/components/filterBy/FilteredChips.vue';
import FilterMenu from '@/components/filterBy/FilterMenu.vue';
import {useFilterByStore} from '@/components/filterBy/useFilterByStore.js';
import {storeToRefs} from 'pinia';
import {computed, watch} from 'vue';

const props = defineProps({
  notification: {
    type: Boolean,
    default: false
  }
});
const filterByStore = useFilterByStore();
const {showFilterMenu} = storeToRefs(filterByStore);
const activeStoreType = computed(() => {
  // eslint-disable-next-line no-unused-vars
  return Object.entries(props).find(([_, value]) => value === true)[0];
});

// Watch the active store type and update the filterByStore storeType value, so we can import the correct store
watch(activeStoreType, (newType) => {
  filterByStore.storeType = newType;
}, {immediate: true});
</script>
