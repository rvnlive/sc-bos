import {useNotificationFilterStore} from '@/routes/ops/notifications/useNotificationFilterStore.js';
import {includesSubstring, stringsMatch} from '@/util/string.js';
import {toValue} from '@/util/vue.js';
import {defineStore} from 'pinia';
import {computed, ref} from 'vue';

export const useFilterByStore = defineStore('filterBy', () => {
  // ----------------------------- Store Imports ----------------------------- //
  const notificationStore = useNotificationFilterStore();
  // ------------------------------------------------------------------------- //

  const storeType = ref('');

  /**
   * Computed property that returns the store in use, based on the storeType value
   *
   * @type {import('vue').ComputedRef<T|null>}
   */
  const filterStore = computed(() => {
    let store = null;

    if (toValue(storeType.value) === 'notification') {
      store = notificationStore;
    }

    return store;
  });

  const isLoading = computed(() => {
    const store = toValue(filterStore);

    if (store) {
      return store.isLoading;
    }

    return false;
  });

  // ----------------------------- Filters & Sources ----------------------------- //
  /**
   * Computed property that returns an array of the available sources
   *
   * @type {import('vue').ComputedRef<Array<*>>}
   */
  const availableSources = computed(() => {
    const store = toValue(filterStore);

    if (store) {
      return store.availableSources.sort((a, b) => a.title.localeCompare(b.title));
    }

    return [];
  });

  /**
   * Computed property that returns an array of the active filters
   *
   * @type {import('vue').ComputedRef<*|[]>}
   */
  const availableFilters = computed(() => {
    const store = toValue(filterStore);

    if (store) {
      return store.availableFilters.sort((a, b) => a.title.localeCompare(b.title));
    }

    return [];
  });

  const defaultFilters = computed(() => {
    const store = toValue(filterStore);

    if (store) {
      return store.defaultFilters.sort((a, b) => a.key.localeCompare(b.key));
    }

    return [];
  });

  // ----------------------------- Filter Finders ----------------------------- //
  const filterInput = ref('');
  const selectedOption = ref(null);

  /**
   * Computed property that returns an array of the available sources when the filter menu is open
   * This is also being used to filter the sources based on the search input
   *
   * @type {import('vue').ComputedRef<Array<*>>}
   */
  const findSearchedSources = computed(() => {
    if (!availableSources.value) return [];

    if (filterInput.value === '') return availableSources.value;

    return ['All', ...availableSources.value.filter((source) => {
      return stringsMatch(source.title, filterInput.value) ||
          includesSubstring(source.title, filterInput.value);
    })];
  });

  /**
   * Computed property that returns an array of the available source options when a source is selected
   * This is used to filter the options available based on the search input
   *
   * @type {import('vue').ComputedRef<Array<*>>}
   */
  const findSearchedSourceOptions = computed(() => {
    if (!selectedOption.value || !selectedOption.value.value) return [];

    if (filterInput.value === '') return selectedOption.value.value;


    return selectedOption.value.value.filter((source) => {
      return stringsMatch(source, filterInput.value) ||
          includesSubstring(source, filterInput.value);
    });
  });

  /**
   * Checks if the passed value is an active filter
   *
   * @param {string} title
   * @return {boolean}
   */
  const findActiveFilter = (title) => {
    const findFilter = availableFilters.value.find((filter) => {
      return stringsMatch(filter.title, title);
    });

    if (findFilter?.isActive) {
      return findFilter.isActive;
    } else {
      return false;
    }
  };

  /**
   * Checks if the passed value is an active filter option
   *
   * @param {string} value
   * @return {boolean}
   */
  const findActiveFilterOption = (value) => {
    const filter = availableFilters.value.find((filter) => {
      return stringsMatch(filter.title, selectedOption.value.title);
    });

    const filterOption = availableFilters.value.find((filter) => {
      return stringsMatch(String(filter.value), value);
    });

    if (filter?.isActive && filterOption?.isActive) {
      return filterOption.isActive;
    } else {
      return false;
    }
  };

  const activeFilter = computed(() => {
    return (source, title) => source ? findActiveFilter(title) : findActiveFilterOption(title);
  });

  // ----------------------------- Filter Modifiers ----------------------------- //
  /**
   * Disable remove action if filter is set by default and source not available
   *
   * @param {string} title
   * @return {boolean}
   */
  const disableDefault = (title) => {
    return availableFilters.value.some((filter) => {
      return filter.title === title && filter.isDefault;
    });
  };


  /**
   * Update the filter value in the active query
   *
   * @param {string} title
   * @param {string|boolean|undefined} value
   * @return {void}
   */
  const updateFilter = (title, value) => {
    const store = toValue(filterStore);

    if (store) {
      store.updateFilter(title.toLowerCase(), value);
    }
  };

  /**
   * Remove the filter from the active query
   *
   * @param {string} title
   * @return {void}
   */
  const removeOneFilter = (title) => {
    const store = toValue(filterStore);

    if (store) {
      store.removeFilter(title.toLowerCase());
    }
  };

  /**
   * Remove all filters from the active query
   *
   * @return {void}
   */
  const removeAllFilters = () => {
    const store = toValue(filterStore);

    if (store) {
      availableFilters.value.forEach((filter) => {
        store.removeFilter(filter.title.toLowerCase());
      });
    }
  };

  const doRemoveFilter = () => {
    // If the selected option is null, remove all filters
    if (selectedOption.value === null) {
      removeAllFilters();
      //
    } else {
      // Otherwise just remove the selected option's filter
      const activeFilterTitle = availableFilters.value.find(
          (filter) => stringsMatch(filter.title, selectedOption.value.title)
      );

      removeOneFilter(activeFilterTitle.title);
    }
  };

  // ----------------------------- Range ----------------------------- //
  /**
   * Computed property that returns the length of the range.
   * Length of the range represents the number of options in the range.
   * This is used to determine the number of steps in the range slider.
   *
   * @type {import('vue').ComputedRef<number>}
   */
  const identifyRangeLength = computed(() => {
    const store = toValue(filterStore);

    if (store) {
      return store.getRangeLength;
    }

    return 0;
  });

  /**
   * Mutatable (writable) computed property that returns and/or sets the range values.
   * The range uses 2 values, the start and end of the range.
   * By default, the range is set with both values undefined.
   *
   * @type {import('vue').WritableComputedRef<Array<undefined | number>>}
   */
  const identifyRange = computed({
    get: () => {
      const store = toValue(filterStore);

      if (store) {
        return store.getSetRange;
      }

      return [];
    },
    set: (value) => {
      const store = toValue(filterStore);

      if (store) {
        store.getSetRange = value;
      }
    }
  });

  /**
   * Computed property that returns the active range as string.
   * This is used to display the range values in the UI.
   * The range values are displayed as a string in a small variety of formats:
   * - `All` - when no range is set (both values are undefined OR range set to first and last options)
   * - `${value 1}` - both values of the range are the same
   * - `${value 1} to ${value 2}` - both values of the range
   * - `${value 1} and above` - when only one value is set (modified start of the range)
   * - `${value 2} and below` - when only one value is set (modified end of the range)
   *
   * @type {import('vue').ComputedRef<string>}
   */
  const activeRangeString = computed(() => {
    const store = toValue(filterStore);

    if (store) {
      return store.getActiveRangeString;
    }

    return '';
  });

  //
  // ----------------------------- Helpers ----------------------------- //
  const selectedFilterChip = ref(null);

  /**
   * Clear the filter selection(s) and input
   */
  const clearSelection = () => {
    filterInput.value = '';
    selectedFilterChip.value = null;
    selectedOption.value = null;
  };

  const clearDisabled = computed(() => {
    const findActiveFilters = availableFilters.value.filter((filter) => filter.isActive);

    // Disable clear button if there are no active filters
    if (!findActiveFilters.length) {
      return true;
    }

    // Disable clear button if the selected option has no active filter
    if (selectedOption.value !== null) {
      return !findActiveFilter(selectedOption.value.title);
    }

    return false;
  });

  const clearResetButtonLabels = computed(() => {
    // If the selected option is null, we can assume that the filter menu is open
    // and the source list (not the options) are visible, so show 'Clear all filters'
    if (selectedOption.value === null) {
      return 'Clear all filters';
      //
      // If there is a selected option, check what type it is
      // If it is a range slider or radio buttons , show 'Reset'
    } else if (['radio', 'range'].includes(selectedOption.value?.type)) {
      return 'Reset';
    }

    // Otherwise, show 'Clear'
    return 'Clear';
  });

  //
  // ----------------------------- Menus ----------------------------- //
  const showFilterMenu = ref(false);


  /**
   * Show or hide the filter menu
   *
   * This accepts a string which is the title of the filter,
   * and it toggles the showMenu property of the filter.
   * Also clears the selection if the menu is not open, so we have a clean slate
   * when the menu is opened again (no leftover selections)
   *
   * @param {string} title
   * @return {void}
   */
  const showHideChipMenu = (title) => {
    const store = toValue(filterStore);

    if (store) {
      store.availableFilters = store.availableFilters.map((filter) => {
        if (stringsMatch(filter.title, title)) {
          filter.showMenu = !filter.showMenu;

          if (filter.showMenu === false) {
            clearSelection();
          }
        }
        return filter;
      });
    }
  };

  return {
    storeType,
    showFilterMenu,
    filterInput,
    selectedFilterChip,
    selectedOption,
    isLoading,

    availableSources,
    availableFilters,
    defaultFilters,
    findSearchedSources,
    findSearchedSourceOptions,
    findActiveFilter,
    findActiveFilterOption,
    activeFilter,

    disableDefault,
    updateFilter,
    removeOneFilter,
    removeAllFilters,
    doRemoveFilter,

    identifyRangeLength,
    identifyRange,
    activeRangeString,

    showHideChipMenu,

    clearSelection,
    clearDisabled,
    clearResetButtonLabels
  };
});
