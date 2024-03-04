import {stringsMatch} from '@/util/string.js';
import {defineStore} from 'pinia';
import {computed, ref, watch} from 'vue';

export const useNotificationFilterStore = defineStore('notificationFilter', () => {
  const severityLevels = {
    9: 'Info',
    13: 'Warn',
    17: 'Alert',
    21: 'Danger'
  };
  const activeQuery = ref({
    acknowledged: undefined,
    createdNotBefore: undefined,
    createdNotAfter: undefined,
    severityNotAbove: undefined,
    severityNotBelow: undefined,
    floor: undefined,
    subsystem: undefined,
    source: undefined,
    resolved: false,
    resolvedNotBefore: undefined,
    resolvedNotAfter: undefined,
    zone: undefined
  });
  const initialQuery = ref({
    acknowledged: undefined,
    createdNotBefore: undefined,
    createdNotAfter: undefined,
    severityNotAbove: undefined,
    severityNotBelow: undefined,
    floor: undefined,
    subsystem: undefined,
    source: undefined,
    resolved: false,
    resolvedNotBefore: undefined,
    resolvedNotAfter: undefined,
    zone: undefined
  });
  const isLoading = ref(false);

  //
  //

  /**
   * Update the active query reference with the new value
   * This will only update the value if the key exists in the current context
   * Otherwise, it will be ignored
   * This is used to update the query on initial load - eg.: props are defined by the parent component
   *
   * @param {string} propName
   * @param {*} propValue
   * @return {void}
   */
  const updateActiveQuery = (propName, propValue) => {
    // Check if this key/value exists in the current context
    if (activeQuery.value.hasOwnProperty(propName)) {
      activeQuery.value[propName] = propValue; // Update the value
    }
  };

  /**
   * Store the initial query for reference and reset purposes
   *
   * @param {string} propName
   * @param {string|number|boolean|Date|undefined} propValue
   */
  const storeInitialQuery = (propName, propValue) => {
    // Check if this key/value exists in the current context
    if (initialQuery.value.hasOwnProperty(propName)) {
      activeQuery.value[propName] = propValue; // Update the value
    }
  };

  //
  // -------------------- Available Sources -------------------- //
  /**
   * Reference to the available sources
   * This is used to populate the filter menu's main filter option list
   *
   * @type {import('vue').Ref<Array<*>>}
   */
  const availableSources = ref([]);

  /**
   * Reset the available sources to an empty array
   *
   * @return {void}
   */
  const resetAvailableSources = () => {
    availableSources.value = [];
  };

  //
  // -------------------- Filter Setup -------------------- //
  const availableFilters = ref([]);

  /**
   * Computed property for structuring the filters and their states based on the initial query
   * This we can use to determine the default state of each filter on active query change.
   * Each 'undefined' value is considered a default filter and active
   *
   * @type {import('vue').ComputedRef<Array<*>>}
   */
  const defaultFilters = computed(() => {
    // Start by filtering and mapping the initialQuery to check if each is a default filter
    let filters = Object.entries(initialQuery.value).map(([key, initialValue]) => {
      const isActiveValueEqual = activeQuery.value[key] === initialValue;
      // Return each filter with its current state, and whether it's considered default
      return {
        key,
        value: activeQuery.value[key] !== undefined ? activeQuery.value[key] : initialValue,
        isActive: true,
        isDefault: isActiveValueEqual
      };
    });

    // Check if severityNotAbove and severityNotBelow are both undefined in the activeQuery
    // If so, add them as a combined entry under 'Severity'
    if (activeQuery.value.severityNotAbove === undefined && activeQuery.value.severityNotBelow === undefined) {
      filters = filters.filter(filter => filter.key !== 'severityNotBelow' && filter.key !== 'severityNotAbove');
      filters.push({
        key: 'Severity',
        value: [undefined, undefined],
        isActive: true,
        isDefault: true
      });
    } else {
      // If either severityNotAbove or severityNotBelow is defined, we adjust their representation
      const severityIndexAbove = filters.findIndex(filter => filter.key === 'severityNotAbove');
      const severityIndexBelow = filters.findIndex(filter => filter.key === 'severityNotBelow');
      if (severityIndexAbove !== -1 && severityIndexBelow !== -1) {
        const combinedSeverity = {
          key: 'Severity',
          value: [
            filters[severityIndexBelow].value,
            filters[severityIndexAbove].value
          ],
          isActive: true,
          isDefault: filters[severityIndexAbove].isDefault && filters[severityIndexBelow].isDefault

        };
        // Remove the individual severity filters from the array
        filters = filters.filter(filter => filter.key !== 'severityNotBelow' && filter.key !== 'severityNotAbove');
        // Add the combined severity filter
        filters.push(combinedSeverity);
      }
    }

    return filters;
  });


  /**
   * Computed property that returns an array of the active filters
   *
   * @param sources
   */
  const activeFilters = computed(() => {
    const query = activeQuery.value;

    // Directly use a structure that mirrors the approach taken with defaultFilters
    let filters = Object.entries(query)
        .filter(([key, value]) => value !== undefined && value !== null && key !== 'resolved')
        .map(([key, value]) => ({
          key: key,
          value: value,
          isActive: true, // Mark as active since it's part of the active query
          isDefault: defaultFilters.value.some(def => def.key === key && def.isDefault)
        }));

    // Handling Severity as a combined entry
    const severityNotAbove = filters.find(filter => filter.key === 'severityNotAbove');
    const severityNotBelow = filters.find(filter => filter.key === 'severityNotBelow');

    // Remove individual severity filters from the array
    filters = filters.filter(filter => filter.key !== 'severityNotAbove' && filter.key !== 'severityNotBelow');

    if (severityNotAbove || severityNotBelow) {
      // Add combined Severity entry if either severity filter is defined
      filters.push({
        key: 'Severity',
        value: [
          severityNotBelow ? severityNotBelow.value : undefined,
          severityNotAbove ? severityNotAbove.value : undefined
        ],
        isActive: true,
        isDefault: (severityNotAbove ? severityNotAbove.isDefault : true) &&
            (severityNotBelow ? severityNotBelow.isDefault : true)
      });
    }

    return filters;
  });


  const setUpAvailableFilters = (sources) => {
    if (!sources.length) return [];


    return sources.map(source => {
      const defaultFilter = defaultFilters.value.find(def => stringsMatch(def.key, source.title));

      return {
        isActive: false,
        isDefault: defaultFilter.isDefault,
        showChip: !defaultFilter.isDefault,
        showMenu: false,
        title: source.title,
        value: defaultFilter.isDefault ? defaultFilter.value : null
      };
    });
  };

  // Watch for changes in the availableSources and update the availableFilters value,
  // so that the filter menu and chip can display the correct values
  watch(availableSources, (sources) => {
    availableFilters.value = setUpAvailableFilters(sources);
  }, {immediate: true, deep: true});

  // Watch for changes in the activeFilters and update the availableFilters value,
  // so that the filter menu and chip can display the correct values
  watch(activeFilters, (filters) => {
    availableFilters.value.forEach(filter => {
      const activeFilter = filters.find(f => stringsMatch(f.key, filter.title));
      if (activeFilter) {
        filter.value = activeFilter.value;
        filter.isActive = true;
        filter.isDefault = false;
        filter.showChip = true;
      } else {
        filter.value = defaultFilters.value.find(def => stringsMatch(def.key, filter.title)).value;
        filter.isActive = false;
        filter.isDefault = true;
        filter.showChip = false;
      }
    });
  }, {immediate: true, deep: true});

  //
  // -------------------- Identifiers/Getters -------------------- //
  const getSeverityRangeLength = computed(() => {
    return Object.entries(severityLevels).length;
  });
  const getSeverityRange = computed({
    get: () => {
      const severityIndexes = Object.keys(severityLevels).map(Number);

      // Determine indexes for the low and high severity levels in the active query
      let lowIndex = severityIndexes.indexOf(activeQuery.value.severityNotBelow);
      let highIndex = severityIndexes.indexOf(activeQuery.value.severityNotAbove);

      // Default to the full range if both are unspecified
      if (lowIndex === -1 && highIndex === -1) {
        return [0, severityIndexes.length - 1];
      }

      // Adjust indexes if only one side is specified
      if (lowIndex === -1) lowIndex = 0;
      if (highIndex === -1) highIndex = severityIndexes.length - 1;

      return [lowIndex, highIndex];
    },
    set: ([lowIndex, highIndex]) => {
      const severityIndexes = Object.keys(severityLevels).map(Number);

      // Update activeQuery based on the provided indexes
      activeQuery.value.severityNotBelow = severityIndexes[lowIndex];
      activeQuery.value.severityNotAbove = severityIndexes[highIndex];

      // Handle boundary cases by setting to undefined if indexes are at the ends
      if (lowIndex === 0) activeQuery.value.severityNotBelow = undefined;
      if (highIndex === severityIndexes.length - 1) activeQuery.value.severityNotAbove = undefined;
    }
  });
  const getActiveSeverityRangeString = computed(() => {
    const {severityNotBelow: lowValue, severityNotAbove: highValue} = activeQuery.value;

    // Handle cases where both or neither of the values are undefined
    if (lowValue === undefined && highValue === undefined) {
      return 'All';
    }

    // Default labels (used if severityLevels does not contain the value)
    const defaultLabels = {low: 'Info', high: 'Danger'};

    // Labels from severityLevels with fallbacks
    const lowLabel = severityLevels[lowValue] ?? defaultLabels.low;
    const highLabel = severityLevels[highValue] ?? defaultLabels.high;

    // Special cases based on specific combinations of lowValue and highValue
    if (lowValue === undefined) {
      if (highValue === 9) return 'Info';
      return `${highLabel} and below`;
    }

    if (highValue === undefined) {
      if (lowValue === 21) return 'Danger';
      return `${lowLabel} and above`;
    }

    if (lowValue === highValue) {
      return lowLabel;
    }

    // General case for defined and unequal lowValue and highValue
    return `${lowLabel} to ${highLabel}`;
  });

  //
  // -------------------- Filter Modifiers -------------------- //
  /**
   * Update the filter value in the active query
   *
   * @param {string} filter
   * @param {string|boolean|undefined} value
   * @return {void}
   */
  const updateFilter = (filter, value) => {
    if (filter === 'acknowledged') {
      if (value === 'All') {
        activeQuery.value[filter] = undefined;
      } else {
        activeQuery.value[filter] = value === 'Yes';
      }
      return;
    }

    if (filter === 'severity') {
      activeQuery.value.severityNotBelow = value[0];
      activeQuery.value.severityNotAbove = value[1];
      return;
    }

    activeQuery.value[filter] = value;
  };

  /**
   * Remove the filter from the active query
   *
   * @param {string} title
   * @return {void}
   */
  const removeFilter = (title) => {
    const resetShowMenu = availableFilters.value.find(filter => stringsMatch(filter.title, title));
    resetShowMenu.showMenu = false;

    if (title === 'resolved') {
      activeQuery.value[title] = false;
      return;
    }

    if (title === 'acknowledged') {
      activeQuery.value[title] = undefined;
      return;
    }

    if (title === 'severity') {
      activeQuery.value.severityNotBelow = undefined;
      activeQuery.value.severityNotAbove = undefined;
      return;
    }

    activeQuery.value[title] = undefined;
  };


  return {
    severityLevels,
    updateActiveQuery,
    storeInitialQuery,

    query: activeQuery.value,
    initialQuery,
    isLoading,

    availableSources,
    resetAvailableSources,

    availableFilters,
    activeFilters,
    defaultFilters,

    getRangeLength: getSeverityRangeLength,
    getSetRange: getSeverityRange,
    getActiveRangeString: getActiveSeverityRangeString,

    updateFilter,
    removeFilter

  };
});
