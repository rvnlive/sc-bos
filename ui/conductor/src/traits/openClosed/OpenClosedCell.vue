<template>
  <StatusAlert v-if="error" icon="mdi-cancel" :resource="error"/>

  <v-tooltip v-else left>
    <template #activator="{ on }">
      <v-icon :class="openClosedDoorState.class" right size="20" v-on="on">
        {{ openClosedDoorState.icon }}
      </v-icon>
    </template>
    <span class="text-capitalize">{{ openClosedDoorState?.text }}</span>
  </v-tooltip>
</template>

<script setup>
import StatusAlert from '@/components/StatusAlert.vue';
import useOpenClosedTrait from '@/traits/openClosed/useOpenClosedTrait.js';

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

const {openClosedDoorState, error} = useOpenClosedTrait(() => props.name, () => props.paused);
</script>

<style scoped>
.open, .moving {
  color: var(--v-success-base);
}

.closed {
  color: var(--v-warning-base);
}

.unknown {
  color: grey;
}
</style>
