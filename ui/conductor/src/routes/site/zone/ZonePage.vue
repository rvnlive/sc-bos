<template>
  <v-container fluid>
    <v-toolbar flat dense color="transparent" class="mb-3">
      <v-combobox
          v-if="configStore.config?.hub"
          v-model="node"
          :items="Object.values(hubStore.nodesList)"
          label="System Component"
          item-text="name"
          item-value="name"
          hide-details="auto"
          :loading="hubStore.nodesListCollection.loading ?? true"
          outlined/>
      <v-spacer/>
      <v-btn class="ml-6" v-if="editMode" @click="save" color="accent" :disabled="blockActions">
        <v-icon left>mdi-content-save</v-icon>
        Save
      </v-btn>
      <v-btn class="ml-6" v-else @click="editMode=true" :disabled="blockActions">
        <v-icon left>mdi-pencil</v-icon>
        Edit
      </v-btn>
    </v-toolbar>
    <div v-if="!zone"/>
    <device-table
        v-else-if="viewType === 'list'"
        :zone="zoneObj"
        :show-select="editMode"
        :row-select="false"
        :filter="zoneDevicesFilter"
        :selected-devices="deviceList"
        @update:selectedDevices="deviceList = $event"/>
  </v-container>
</template>

<script setup>
import {newActionTracker} from '@/api/resource';
import {ServiceNames} from '@/api/ui/services';
import useAuthSetup from '@/composables/useAuthSetup';
import DeviceTable from '@/routes/devices/components/DeviceTable.vue';
import {Zone} from '@/routes/site/zone/zone';
import {useHubStore} from '@/stores/hub';
import {useServicesStore} from '@/stores/services';
import {useUiConfigStore} from '@/stores/ui-config';
import {Service} from '@sc-bos/ui-gen/proto/services_pb';
import {computed, ref, watch} from 'vue';

const {blockActions} = useAuthSetup();

const servicesStore = useServicesStore();
const configStore = useUiConfigStore();
const hubStore = useHubStore();
const zoneCollection = ref();

const props = defineProps({
  zone: {
    type: String,
    default: ''
  }
});

const node = computed({
  get() {
    return servicesStore.node;
  },
  set(val) {
    servicesStore.node = val;
  }
});

watch(node, async () => {
  zoneCollection.value = servicesStore.getService(
      ServiceNames.Zones,
      await servicesStore.node?.commsAddress,
      await servicesStore.node?.commsName).servicesCollection;
}, {immediate: true});

const zoneObj = computed(() => {
  const z = zoneCollection?.value?.resources?.value[props.zone] ?? (new Service()).toObject();
  return new Zone(z);
});

const viewType = ref('list');
const editMode = ref(false);

const deviceList = computed({
  get() {
    return zoneObj.value.deviceIds;
  },
  set(value) {
    zoneObj.value.devices = value;
  }
});


/**
 * @param {Device.AsObject} device
 * @return {boolean}
 */
function zoneDevicesFilter(device) {
  return editMode.value || (zoneObj?.value?.deviceIds?.indexOf(device.name) >= 0 ?? true);
}

const saveTracker = newActionTracker();

/**
 * Save the new device list to the zone
 */
function save() {
  zoneObj.value.save(saveTracker);
  editMode.value = false;
}

</script>

<style scoped>

</style>
