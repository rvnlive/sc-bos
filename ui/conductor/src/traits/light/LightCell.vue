<template>
  <status-alert v-if="error" icon="mdi-lightbulb-outline" :resource="error"/>

  <span v-else class="d-flex flex-row flex-nowrap">
    <v-tooltip bottom>
      <template #activator="{ on, attrs }">
        <span v-on="on" v-bind="attrs" class="d-flex flex-row">
          <span class="text-caption" style="min-width: 4ex">{{ levelStr }}</span>
          <v-icon right :color="level > 0 ? 'yellow' : 'white' " size="20">
            {{ icon }}
          </v-icon>
        </span>
      </template>
      <span>Lighting</span>
    </v-tooltip>
  </span>
</template>

<script setup>
import StatusAlert from '@/components/StatusAlert.vue';
import {useBrightness, usePullBrightness} from '@/traits/light/light.js';

const props = defineProps({
  name: {
    type: String,
    default: ''
  },
  paused: {
    type: Boolean,
    default: false
  }
});
const {value, streamError: error} = usePullBrightness(() => props.name, () => props.paused);
const {levelStr, level, icon} = useBrightness(value);
</script>
