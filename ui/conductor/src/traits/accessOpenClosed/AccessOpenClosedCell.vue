<template>
  <StatusAlert
      v-if="accessAttemptError || openClosedError"
      icon="mdi-cancel"
      :resource="accessAttemptError || openClosedError"/>

  <v-tooltip v-else left>
    <template #activator="{on}">
      <v-icon :class="[accessAttemptGrantState]" right size="20" v-on="on">
        {{ openClosedDoorState.icon }}
      </v-icon>
    </template>
    <span>
      <span class="text-capitalize">Access: {{ accessAttemptGrantState.split('_').join(' ') }}</span>
      {{ openClosedDoorState.text ? ` - ${openClosedDoorState.text}` : '' }}
    </span>
  </v-tooltip>
</template>

<script setup>
import StatusAlert from '@/components/StatusAlert.vue';
import useAccessTrait from '@/traits/access/useAccessTrait.js';
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

const {accessAttemptGrantState, error: accessAttemptError} = useAccessTrait(() => props.name, () => props.paused);
const {openClosedDoorState, error: openClosedError} = useOpenClosedTrait(() => props.name, () => props.paused);
</script>

<style scoped>
.granted {
  color: var(--v-success-base);
}

.denied {
  color: var(--v-warning-base);
}

.tailgate, .forced, .failed {
  color: var(--v-error-base);
}

.pending, .aborted {
  color: var(--v-info-base);
}

.grant_unknown {
  color: grey;
}
</style>
