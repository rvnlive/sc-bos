<template>
  <div>
    <v-col class="ma-0 pa-0" cols="12">
      <v-subheader class="text-h6 justify-center mt-n3">
        {{ activeRangeString }}
      </v-subheader>
    </v-col>

    <v-range-slider
        class="mt-2 mb-n2 mx-7"
        :max="identifyRangeLength - 1"
        :min="0"
        :thumb-label="false"
        thumb-size="0"
        tick-size="4"
        ticks="always"
        :step="1"
        v-model="identifyRange"/>
  </div>
</template>

<script setup>
import {storeToRefs} from 'pinia';
import {useFilterByStore} from '@/components/filterBy/useFilterByStore.js';

const filterByStore = useFilterByStore();
const {identifyRange, identifyRangeLength, activeRangeString} = storeToRefs(filterByStore);
</script>
