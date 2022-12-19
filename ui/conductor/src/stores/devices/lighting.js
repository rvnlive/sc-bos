import {defineStore} from 'pinia';
import {computed, ref} from 'vue';
import lightingData from './lighting.json';

export const useLightingStore = defineStore('lightingStore', () => {
  // state data for lighting
  const lights = ref(lightingData);

  const headers = ref([
    {
      text: 'Device ID',
      align: 'start',
      value: 'device_id'
    },
    {text: 'Location', value: 'location'},
    {text: 'Status', value: 'status'},
    {text: 'Battery Status', value: 'battery_status'},
    {text: 'Model', value: 'model'}
  ]);

  const lightData = ref([
    {
      title: 'Building',
      content: 'Upper Gough Street'
    },
    {
      title: 'Floor',
      content: 'LO1'
    },
    {
      title: 'Zone',
      content: 'L02_12'
    },
    {
      title: 'Manufacturer',
      content: 'Philips'
    },
    {
      title: 'Model',
      content: 'LED 1245812'
    },
    {
      title: 'Installed on',
      content: '12.09.22'
    },
    {
      title: 'Serial Number',
      content: '12348a7a595'
    },
    {
      title: 'DALI Address',
      content: '1234'
    },
    {
      title: 'DALI Controller',
      content: '1234'
    }
  ]);

  // selected rows in the data table
  const selected = ref([]);

  // the item we are displaying in the lighting sidebar
  const selectedItem = ref({});

  // state for filters for the table
  const status = ref('All');

  const model = ref('All');

  const search = ref('');

  const models = computed(() => {
    return [...new Set(lights.value.map((light) => light.model))];
  });

  models.value.unshift('All');

  const statuses = ref(['All', 'On', 'Off']);

  // state for the drawer for the light details
  const drawer = ref(false);

  const filteredLights = computed(() =>
    lights.value.filter((light) => {
      if (status.value === 'All' && model.value === 'All') {
        return true;
      } else if (status.value === 'All') {
        return light.model === model.value;
      } else if (model.value === 'All') {
        return light.status === status.value;
      } else {
        return light.status === status.value && light.model === model.value;
      }
    })
  );

  const bulkAction = (action) => {
    if (action === 'On') {
      selected.value.forEach((light) => {
        light.status = 'On';
      });
      selected.value = [];
    } else if (action === 'Off') {
      selected.value.forEach((light) => {
        light.status = 'Off';
      });
      selected.value = [];
    }
  };

  const toggleDrawer = () => {
    drawer.value = !drawer.value;
  };

  const setSelectedItem = (selected) => {
    selectedItem.value = selected;
  };

  const increaseBrightness = () => {
    if (selectedItem.value.brightness < 100) {
      selectedItem.value.brightness++;
    }
  };
  const decreaseBrightness = () => {
    if (selectedItem.value.brightness > 0) {
      selectedItem.value.brightness--;
    }
  };

  const turnOff = () => {
    selectedItem.value.status = 'Off';
  };

  const turnOn = () => {
    selectedItem.value.status = 'On';
  };

  return {
    lights,
    headers,
    selected,
    lightData,
    status,
    model,
    search,
    models,
    statuses,
    filteredLights,
    bulkAction,
    drawer,
    toggleDrawer,
    selectedItem,
    setSelectedItem,
    increaseBrightness,
    decreaseBrightness,
    turnOff,
    turnOn
  };
});
