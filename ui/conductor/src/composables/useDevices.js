import {reactive, ref, computed, onMounted, onUnmounted, watch} from 'vue';
import {closeResource, newResourceValue} from '@/api/resource';
import {pullDevicesMetadata} from '@/api/ui/devices';

import {useErrorStore} from '@/components/ui-error/error';
import {useDevicesStore} from '@/routes/devices/store';

const NO_FLOOR = '< no floor >';

/**
 *
 * @param {Object} props
 */
export default function(props) {
  const devicesStore = useDevicesStore();
  const errorStore = useErrorStore();

  // Create reactive resources and data
  const floorListResource = reactive(newResourceValue());
  const filterFloor = ref('All');
  const search = ref('');

  // Fetch floor list on component mount
  onMounted(() => {
    const req = {includes: {fieldsList: ['metadata.location.floor']}, updatesOnly: false};
    pullDevicesMetadata(req, floorListResource);
  });

  // Close resource on component unmount
  onUnmounted(() => {
    closeResource(floorListResource);
  });

  // Computed property for the floor list
  const floorList = computed(() => {
    const fieldCounts = floorListResource.value?.fieldCountsList || [];
    const floorFieldCounts = fieldCounts.find(v => v.field === 'metadata.location.floor');
    if (!floorFieldCounts || floorFieldCounts.countsMap.size <= 0) return [];

    const floors = floorFieldCounts.countsMap.map(([k]) => (k === '' ? NO_FLOOR : k));
    floors.unshift('All');
    return floors;
  });

  // Create reactive collection
  const collection = reactive(devicesStore.newCollection());
  collection.needsMorePages = true; // Todo: Connect with paging logic instead

  // Computed property for the query object
  const query = computed(() => {
    const q = {conditionsList: []};
    if (search.value) {
      const words = search.value.split(/\s+/);
      q.conditionsList.push(...words.map(word => ({stringContainsFold: word})));
    }
    if (props.subsystem.toLowerCase() !== 'all') {
      q.conditionsList.push({field: 'metadata.membership.subsystem', stringEqualFold: props.subsystem});
    }
    switch (filterFloor.value.toLowerCase()) {
      case 'all':
        // no filter
        break;
      case NO_FLOOR:
        q.conditionsList.push({field: 'metadata.location.floor', stringEqualFold: ''});
        break;
      default:
        q.conditionsList.push({field: 'metadata.location.floor', stringEqualFold: filterFloor.value});
        break;
    }
    return q;
  });

  // Watch for changes to the query object and fetch new device list
  watch(query, () => collection.query(query.value), {deep: true, immediate: true});

  // UI error handling
  let unwatchErrors;
  onMounted(() => {
    unwatchErrors = errorStore.registerCollection(collection);
  });

  onUnmounted(() => {
    if (unwatchErrors) unwatchErrors();
    collection.reset(); // stop listening when the component is unmounted
  });

  // Computed property for the filtered table data
  const devicesData = computed(() => {
    return Object.values(collection.resources.value).filter(props.filter);
  });

  return {
    floorList,
    filterFloor,
    search,
    query,
    devicesData
  };
}