<template>
  <v-card elevation="0" tile>
    <v-list tile class="ma-0 pa-0">
      <v-subheader class="text-title-caps-large neutral--text text--lighten-3">Air Quality</v-subheader>
      <v-list-item v-for="(val, key) of tableData" :key="key" class="py-1">
        <v-list-item-title class="text-body-small">{{ key }}</v-list-item-title>
        <v-list-item-subtitle class="font-weight-medium">{{ val }}</v-list-item-subtitle>
      </v-list-item>
    </v-list>
    <v-progress-linear
        v-if="hasScore"
        :value="score"
        height="34"
        class="mx-4 my-2"
        background-color="neutral lighten-1"
        :color="scoreColor"/>
  </v-card>
</template>

<script setup>
import {useAirQuality} from '@/traits/airQuality/airQuality.js';

const props = defineProps({
  value: {
    type: Object, // of AirQuality.AsObject
    default: () => ({})
  },
  loading: {
    type: Boolean,
    default: false
  }
});

const {hasScore, score, scoreColor, tableData} = useAirQuality(() => props.value);
</script>

<style scoped>
.v-list-item {
  min-height: auto;
}

.v-progress-linear {
  width: auto;
}

</style>
