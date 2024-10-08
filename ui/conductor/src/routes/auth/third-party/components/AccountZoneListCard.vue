<template>
  <v-card class="pa-0" flat tile>
    <v-card-title
        class="d-flex align-center text-subtitle-2 text-title-caps-large text-neutral-lighten-3 mb-2"
        style="height: 50px">
      Zones
      <v-spacer/>
      <v-btn v-if="hasZoneChanges" color="primary" :disabled="blockActions" @click="emitZoneChanges">
        Save
      </v-btn>
    </v-card-title>
    <v-card-actions>
      <account-zone-chooser v-model:zones="chooserZones" class="mx-2 flex-grow-1"/>
    </v-card-actions>
    <v-card-text v-if="chooserZones.length === 0" class="pt-1 font-italic">
      {{ hasZoneChanges ? 'If there are no' : 'No' }}
      zones associated with this account, they will not be able to access any devices.
    </v-card-text>
  </v-card>
</template>

<script setup>
import useAuthSetup from '@/composables/useAuthSetup';
import AccountZoneChooser from '@/routes/auth/third-party/components/AccountZoneChooser.vue';
import {computed, ref} from 'vue';

const props = defineProps({
  zoneList: {
    type: Array, // string[]
    default: () => []
  }
});

const emit = defineEmits(['update:zone-list']);

const zoneChanges = ref(null);
const chooserZones = computed({
  get() {
    if (zoneChanges.value !== null) {
      return zoneChanges.value;
    } else {
      return props.zoneList;
    }
  },
  set(v) {
    zoneChanges.value = v;
  }
});

const hasZoneChanges = computed(() => {
  if (zoneChanges.value === null) {
    return false;
  }
  if (zoneChanges.value.length !== props.zoneList.length) {
    return true;
  }

  const zoneNames = props.zoneList.reduce((acc, zone) => {
    acc[zone] = true;
    return acc;
  }, {});
  return zoneChanges.value?.some(zone => !zoneNames[zone.name]) ?? false;
});

const emitZoneChanges = () => {
  emit('update:zone-list', zoneChanges.value.map(zone => zone.name));
};

// ------------------------------ //
// ----- Authentication settings ----- //

const {blockActions} = useAuthSetup();
</script>

<style scoped></style>
