<template>
  <div class="d-flex flex-wrap">
    <v-card v-for="node in nodesList" :key="node.name" width="300px" class="ma-2">
      <v-card-title class="text-body-large font-weight-bold">{{ node.name }}</v-card-title>
      <v-card-subtitle v-if="node.description !== ''">{{ node.description }}</v-card-subtitle>
      <v-card-text>
        <v-list dense>
          <v-list-item
              class="pa-0"
              style="min-height: 20px">
            {{ node.address }}
          </v-list-item>
          <v-list-item
              class="pa-0 ma-0"
              style="min-height: 20px"
              v-for="(trackers, service) in nodeDetails[node.name]"
              :key="service">
            {{ service }}: {{ trackers.metadataTracker?.response?.totalCount }}
          </v-list-item>
        </v-list>
        <v-chip-group>
          <v-chip v-if="isProxy(node.name)" color="accent" small>proxy</v-chip>
          <v-chip v-if="isHub(node.name)" color="primary" small>hub</v-chip>
        </v-chip-group>
      </v-card-text>
    </v-card>
    <div/>
  </div>
</template>

<script setup>
import {ServiceNames} from '@/api/ui/services';
import {useErrorStore} from '@/components/ui-error/error';
import {useHubStore} from '@/stores/hub';
import {useServicesStore} from '@/stores/services';
import {computed, onUnmounted, reactive, set} from 'vue';

const hubStore = useHubStore();
const servicesStore = useServicesStore();
const errorStore = useErrorStore();

const nodeDetails = reactive({});

const unwatchTrackers = [];
onUnmounted(() => {
  unwatchTrackers.forEach(unwatch => unwatch());
});

const nodesList = computed(() => {
  return Object.values(hubStore.nodesList).map(node => {
    Promise.all([node.commsAddress, node.commsName])
        .then(([address, name]) => {
          console.debug('node', node);
          set(nodeDetails, node.name, {
            automations: servicesStore.getService(ServiceNames.Automations, address, name),
            drivers: servicesStore.getService(ServiceNames.Drivers, address, name),
            systems: servicesStore.getService(ServiceNames.Systems, address, name)
          });
          unwatchTrackers.push(errorStore.registerTracker(nodeDetails[node.name].automations.metadataTracker));
          unwatchTrackers.push(errorStore.registerTracker(nodeDetails[node.name].drivers.metadataTracker));
          unwatchTrackers.push(errorStore.registerTracker(nodeDetails[node.name].systems.metadataTracker));
          return Promise.all([
            servicesStore.refreshMetadata(ServiceNames.Automations, address, name),
            servicesStore.refreshMetadata(ServiceNames.Drivers, address, name),
            servicesStore.refreshMetadata(ServiceNames.Systems, address, name)
          ]);
        })
        .catch(e => {
          console.error(e);
        });
    return {
      ...node
    };
  });
});

/**
 * Check if the node has a proxy system service configured
 *
 * @param {string} nodeName
 * @return {boolean}
 */
function isProxy(nodeName) {
  return nodeDetails[nodeName]?.systems.metadataTracker?.response?.typeCountsMap?.some(
      ([name, count]) => name === 'proxy' && count > 0
  );
}
/**
 * Check if the node has a hub system service configured
 *
 * @param {string} nodeName
 * @return {boolean}
 */
function isHub(nodeName) {
  return nodeDetails[nodeName]?.systems.metadataTracker?.response?.typeCountsMap?.some(
      ([name, count]) => name === 'hub' && count > 0
  );
}

</script>

<style scoped>

</style>