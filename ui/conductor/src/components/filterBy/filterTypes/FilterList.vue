<template>
  <span v-if="listItems.length">
    <v-list-item
        v-for="(source) in listItems.filter((item) => item !== 'All')"
        active-class="rounded"
        class="ml-0 text-truncate rounded"
        :key="source.title ? source.title : source"
        :title="source.title ? source.title : source"
        @click="doSelect(source.title ? source.title : source)"
        @keydown.enter="doSelect(source.title ? source.title : source)">
      <v-list-item-icon v-if="isSource" class="ml-n2">
        <v-icon size="22">
          {{ source.icon }}
        </v-icon>
      </v-list-item-icon>
      <v-list-item-content class="ml-n1 pr-0 mr-0">
        <v-list-item-title class="text-body-2 d-flex flex-row justify-space-between align-center">
          {{ source.title ? source.title : source }}
          <v-icon
              v-if="activeFilter(isSource, source.title ? source.title : source)"
              :class="{'ml-auto mr-0': isSource, 'ml-auto': !isSource}"
              color="primary"
              size="16">
            mdi-tag-outline
          </v-icon>
          <v-icon v-if="isSource" class="mr-n1">
            mdi-chevron-right
          </v-icon>
        </v-list-item-title>
      </v-list-item-content>
    </v-list-item>
  </span>

  <span v-else>
    <v-list-item class="ml-0 text-truncate rounded">
      <v-list-item-content class="ml-n1">
        <v-list-item-title class="text-body-2 d-flex flex-row align-center">
          No option available
        </v-list-item-title>
      </v-list-item-content>
    </v-list-item>
  </span>
</template>

<script setup>
import {useFilterByStore} from '@/components/filterBy/useFilterByStore.js';
import {stringsMatch} from '@/util/string.js';
import {storeToRefs} from 'pinia';
import {computed} from 'vue';

const props = defineProps({
  listType: {
    type: String,
    default: 'source'
  }
});

const filterByStore = useFilterByStore();
const {
  activeFilter,
  filterInput,
  findSearchedSources,
  findSearchedSourceOptions,
  selectedOption
} = storeToRefs(filterByStore);

const isSource = computed(() => props.listType === 'source'); // Source or Option
const listItems = computed(() => isSource.value ?
    findSearchedSources.value.filter((item) => item !== 'All') :
    findSearchedSourceOptions.value.filter((item) => item !== 'All')
);

const doSelect = (title) => {
  if (isSource.value) {
    selectedOption.value = listItems.value.find(
        (source) => stringsMatch(source.title, title) && source !== 'All');
    filterInput.value = '';
  } else {
    filterByStore.updateFilter(selectedOption.value.title, title);
  }
};
</script>

<style scoped>

</style>
