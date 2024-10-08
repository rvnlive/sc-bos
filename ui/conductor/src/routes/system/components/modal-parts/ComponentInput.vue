<template>
  <div v-if="!confirmForget">
    <v-form @submit.prevent="onEnroll">
      <v-text-field
          v-model="_address"
          class="mx-8"
          :clearable="_address !== null"
          density="compact"
          hide-details
          label="Component Address"
          variant="outlined"
          @click:clear="_address = null"/>
      <!-- Error label if the address is already enrolled -->
      <v-alert
          v-if="errorText"
          class="mx-8 mt-2 text-capitalize"
          color="error"
          density="compact"
          max-width="400px"
          variant="outlined"
          type="error"
          :text="errorText.message">
        <template #append>
          <status-alert
              v-if="errorText?.error"
              :resource="errorText.error"
              icon="mdi-alert-circle-outline"
              icon-size="24"
              style="margin-top: 2px"/>
        </template>
      </v-alert>

      <v-card-actions class="d-flex flex-row justify-space-around mt-10">
        <v-btn
            class="mr-4 px-4"
            color="primary"
            :disabled="!address || isEnrolled"
            variant="text"
            type="submit"
            @click="onEnroll">
          Enroll Node
        </v-btn>
      </v-card-actions>
    </v-form>
  </div>
  <div v-else style="max-width: 600px">
    <v-card-text class="px-7 text-left text-subtitle-1 font-weight-regular">
      <p>
        Forgetting a node means
        <span class="font-weight-bold text-warning">it can no longer interact with other Smart Core nodes,
          and those nodes cannot interact with it.</span>
      </p>
      <p>
        Any automations that rely on inter-node communication with or from this node
        <span class="font-weight-bold text-error">will stop working!</span>
        This includes managing the node centrally via this app. You can re-enrol this node at any time.
      </p>
    </v-card-text>
    <v-card-actions class="d-flex flex-row justify-space-around mt-10">
      <v-btn
          class="mr-4 px-4"
          variant="text"
          @click="cancelAction">
        Cancel
      </v-btn>
      <v-btn
          class="px-4"
          color="error"
          variant="text"
          @click="forgetHubNode">
        Forget Node
      </v-btn>
    </v-card-actions>
  </div>
</template>

<script setup>
import StatusAlert from '@/components/StatusAlert.vue';
import {formatErrorMessage} from '@/util/error';
import {computed, ref, watch} from 'vue';

const emits = defineEmits([
  'inspectHubNodeAction',
  'resetInspectHubNodeValue',
  'forgetHubNodeAction'
]);
const props = defineProps({
  inspectHubNodeValue: {
    type: Object,
    default: () => ({})
  },
  listItems: {
    type: /** @type {(typeof CohortNode)[]} */ Array,
    default: () => []
  },
  nodeQuery: {
    type: Object,
    default: () => ({})
  }
});
const _address = defineModel('address', {
  type: String,
  default: null
});
const _dialogState = defineModel('dialogState', {
  type: Boolean,
  default: false
});

const isEnrolled = ref(null);
const confirmForget = ref(false);

// Enroll the hub node
const onEnroll = () => {
  if (!_address.value) {
    return;
  }

  emits('inspectHubNodeAction', _address.value);
};

const forgetHubNode = () => {
  if (!props.nodeQuery.address) {
    return;
  }

  emits('forgetHubNodeAction', props.nodeQuery.address);
  _dialogState.value = false;
  confirmForget.value = false;
};

const cancelAction = () => {
  return props.nodeQuery.isToForget ? _dialogState.value = false : confirmForget.value = false;
};

// Display the correct dialog content depending on the confirmForget value
// If confirmForget is true, display the forget dialog
// If confirmForget is false, display the enroll/forget dialog
watch(() => props.nodeQuery, (newQuery) => {
  confirmForget.value = newQuery.isToForget;
}, {immediate: true, deep: true});

// Depending on input, check if the address is enrolled
// and update the isEnrolled value to enable the correct button
watch(_address, (newAddress, oldAddress) => {
  if (newAddress !== oldAddress) {
    emits('resetInspectHubNodeValue');
  }

  // If the address is empty, reset the isEnrolled value to disable all buttons
  if (!newAddress) {
    isEnrolled.value = null;
    return;
  }

  // If the address is not empty, check if it is enrolled
  const matchAddress = props.listItems.find(node => node.grpcAddress === newAddress);

  // If the address is enrolled, enable the forget button
  if (matchAddress) {
    isEnrolled.value = true;

    // If the address is not enrolled, enable the enroll button
  } else if (!matchAddress) {
    isEnrolled.value = false;

    // Otherwise, disable all buttons
  } else {
    isEnrolled.value = null;
  }
}, {immediate: true, deep: true});

// Reset the address value when the dialog is closed
watch(() => props.dialogState, (newState) => {
  if (!newState) {
    _address.value = null;
  }
}, {immediate: true, deep: true});

const errorText = computed(() => {
  if (props.inspectHubNodeValue.error) {
    return {
      message: formatErrorMessage(props.inspectHubNodeValue.error.error.message),
      error: props.inspectHubNodeValue.error
    };
  } else if (!props.inspectHubNodeValue.error && isEnrolled.value) {
    return {
      message: 'This node is already enrolled',
      error: null
    };
  }

  return null;
});
</script>

