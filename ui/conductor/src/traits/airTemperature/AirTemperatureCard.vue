<template>
  <v-card elevation="0" tile>
    <v-list tile class="ma-0 pa-0">
      <v-subheader class="text-title-caps-large neutral--text text--lighten-3">Temperature</v-subheader>
      <v-list-item v-for="(val, key) of airTempData" :key="key" class="py-1">
        <v-list-item-title class="text-body-small text-capitalize">{{ camelToSentence(key) }}</v-list-item-title>
        <v-list-item-subtitle class="text-capitalize font-weight-medium">{{ val }}</v-list-item-subtitle>
      </v-list-item>
    </v-list>
    <v-progress-linear
        height="34"
        class="mx-4 my-2"
        :value="tempProgress"
        background-color="neutral lighten-1"
        color="accent"/>
    <v-card-actions class="px-4">
      <v-spacer/>
      <v-btn
          small
          color="neutral lighten-1"
          elevation="0"
          @click="changeSetPoint(-0.5)"
          :disabled="blockActions || (props.value?.temperatureSetPoint === undefined)">
        Down
      </v-btn>
      <v-btn
          small
          color="neutral lighten-1"
          elevation="0"
          @click="changeSetPoint(0.5)"
          :disabled="blockActions || (props.value?.temperatureSetPoint === undefined)">
        Up
      </v-btn>
    </v-card-actions>
    <v-progress-linear color="primary" indeterminate :active="updateValue.loading"/>
  </v-card>
</template>

<script setup>
import {newActionTracker} from '@/api/resource';
import useAuthSetup from '@/composables/useAuthSetup';
import {useAirTemperature} from '@/traits/airTemperature/airTemperature.js';
import {camelToSentence} from '@/util/string';
import {reactive} from 'vue';

const {blockActions} = useAuthSetup();

const props = defineProps({
  value: {
    type: Object, // of AirTemperature.AsObject
    default: () => ({})
  },
  loading: {
    type: Boolean,
    default: false
  }
});
const emit = defineEmits([
  'updateAirTemperature' // of number | AirTemperature.AsObject | UpdateAirTemperatureRequest.AsObject
]);

const {
  tempProgress,
  airTempData
} = useAirTemperature(() => props.value);

const updateValue = reactive(newActionTracker());

/**
 * @param {number} value
 */
function changeSetPoint(value) {
  if (props.value?.temperatureSetPoint?.valueCelsius !== undefined) {
    emit('updateAirTemperature', props.value.temperatureSetPoint.valueCelsius + value);
  }
}

</script>

<style scoped>
.v-list-item {
  min-height: auto;
}

.v-progress-linear {
  width: auto;
}
</style>
