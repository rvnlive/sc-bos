<template>
  <v-menu
      content-class="mt-2"
      :close-on-content-click="false"
      left
      max-width="290"
      min-width="290"
      nudge-bottom="0"
      nudge-left="0"
      nudge-right="0"
      nudge-top="0"
      offset-y
      return-value="true"
      v-model="showMenu">
    <template #activator="{ on: menuOn, attrs: menuAttrs }">
      <v-tooltip bottom>
        <template #activator="{ on: tooltipOn, attrs: tooltipAttrs }">
          <v-btn
              :class="[
                'text-decoration-none text-capitalize text-truncate text-caption mr-2',
                {'non-interactive-chip': disableDefault(props.filter.title)}
              ]"
              color="neutral lighten-2"
              :disabled="disableDefault(props.filter.title)"
              elevation="0"
              small
              :ripple="false"
              style="max-width: 200px; min-height: 36px;"
              v-on="{ ...menuOn, ...tooltipOn }"
              v-bind="{ ...menuAttrs, ...tooltipAttrs }"
              @click="editFilter(props.filter)">
            {{ camelToSentence(formatObjectValues(props.filter).title) }}
            {{ !formatObjectValues(props.filter).showValue ? '' : `: ${formatObjectValues(props.filter).value}` }}
            <v-icon
                v-if="!disableDefault(props.filter.title)"
                right
                small
                class="ml-2"
                @click.stop="filterByStore.removeOneFilter(props.filter.title)">
              mdi-close
            </v-icon>
          </v-btn>
        </template>
        Edit filter
      </v-tooltip>
    </template>

    <!-- Menu Content -->
    <FilterMenu/>
  </v-menu>
</template>

<script setup>
import FilterMenu from '@/components/filterBy/FilterMenu.vue';
import {useFilterByStore} from '@/components/filterBy/useFilterByStore.js';
import {camelToSentence, includesSubstring, stringsMatch} from '@/util/string.js';
import {storeToRefs} from 'pinia';
import {computed} from 'vue';


const props = defineProps({
  filter: {
    type: Object,
    default: () => {
      return {
        title: '',
        value: null,
        showValue: false
      };
    }
  }
});

const filterByStore = useFilterByStore();
const {
  availableFilters,
  selectedFilterChip,
  selectedOption,
  disableDefault,
  activeRangeString
} = storeToRefs(filterByStore);

const showMenu = computed({
  get: () => {
    return availableFilters.value.find((filter) => stringsMatch(filter.title, props.filter.title)).showMenu;
  },
  // eslint-disable-next-line no-unused-vars
  set: (_) => {
    filterByStore.showHideChipMenu(props.filter.title);
  }
});

const editFilter = (filter) => {
  selectedFilterChip.value = filter;
  selectedOption.value = filterByStore.findSearchedSources
      .filter((source) => source !== 'All')
      .find((option) =>
        includesSubstring(filter.title, option.title) ||
          stringsMatch(option.title, filter.title)
      );
};

const formatObjectValues = (obj) => {
  let formattedObj;

  if (stringsMatch(obj.title, 'Severity')) {
    formattedObj = {
      title: activeRangeString.value,
      value: null,
      showValue: false
    };
  } else if (stringsMatch(obj.title, 'Acknowledged')) {
    formattedObj = {
      title: obj.value === true ? 'Acknowledged' : 'Unacknowledged',
      value: obj.value,
      showValue: false
    };
  } else {
    formattedObj = {
      title: obj.title,
      value: obj.value,
      showValue: true
    };
  }

  return formattedObj;
};
</script>

<style lang="scss" scoped>
.non-interactive-chip {
  pointer-events: none;
}


::v-deep.v-chip:not(.v-chip--disabled) .v-chip__close:hover {
  color: var(--v-error-base); /* Change the close button color on hover for visual effect */
}
</style>
