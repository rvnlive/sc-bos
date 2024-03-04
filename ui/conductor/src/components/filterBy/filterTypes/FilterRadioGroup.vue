<template>
  <v-radio-group v-if="findSearchedSourceOptions.length" class="py-0 mb-n6 mt-1 ml-2" :value="findActiveRadioFilter">
    <v-radio
        v-for="(value, index) in findSearchedSourceOptions"
        class="mb-2"
        :key="selectedOption.title + '-' + index"
        :label="value"
        :value="value"
        @click="filterByStore.updateFilter(selectedOption.title, value)">
      <template #label>
        <v-list-item class="text-truncate rounded">
          <v-list-item-content class="ml-1 pr-0 mr-0 text-truncate">
            <v-list-item-title class="text-body-2 d-flex flex-row justify-space-between align-center">
              {{ value }}
            </v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </template>
    </v-radio>
  </v-radio-group>
  <div v-else>
    <v-list-item class="ml-0 text-truncate rounded">
      <v-list-item-content class="ml-n1">
        <v-list-item-title class="text-body-2 d-flex flex-row align-center">
          No options available
        </v-list-item-title>
      </v-list-item-content>
    </v-list-item>
  </div>
</template>

<script setup>
import {useFilterByStore} from '@/components/filterBy/useFilterByStore';
import {computed} from 'vue';
import {storeToRefs} from 'pinia';
import {stringsMatch} from '@/util/string';


const filterByStore = useFilterByStore();
const {
  availableFilters,
  selectedOption,
  findSearchedSourceOptions
} = storeToRefs(filterByStore);

const findActiveRadioFilter = computed(() => {
  if (!selectedOption.value.value) return 'All';

  const radioValue = availableFilters.value.find((filter) => {
    if (stringsMatch(filter.title, selectedOption.value.title)) {
      return filter;
    }
  });

  if (radioValue.value === undefined) {
    return 'All';
  } else if (typeof radioValue.value === 'boolean') {
    return radioValue.value === false ? 'No' : 'Yes';
  } else {
    return radioValue.value;
  }
});
</script>
