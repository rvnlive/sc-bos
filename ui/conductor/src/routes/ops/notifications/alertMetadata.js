import {closeResource, newResourceValue} from '@/api/resource';
import {pullAlertMetadata} from '@/api/ui/alerts';
import {useErrorStore} from '@/components/ui-error/error';
import {useAppConfigStore} from '@/stores/app-config';
import {useHubStore} from '@/stores/hub';
import {defineStore} from 'pinia';
import {computed, onMounted, onUnmounted, reactive, watch} from 'vue';

export const useAlertMetadata = defineStore('alertMetadata', () => {
  const alertMetadata = reactive(/** @type {ResourceValue<AlertMetadata.AsObject, AlertMetadata>} */newResourceValue());
  const appConfig = useAppConfigStore();
  const hubStore = useHubStore();

  watch(() => appConfig.config, () => {
    init();
  }, {immediate: true});
  watch(() => hubStore.hubNode, () => {
    init();
  }, {immediate: true});

  /**
   *
   */
  function init() {
    // check config is loaded
    if (!appConfig.config) return;
    // check hubNode is loaded if proxy is enabled
    if (appConfig.config.proxy && !hubStore.hubNode) return;

    closeResource(alertMetadata);
    const name = appConfig.config.proxy? hubStore.hubNode.name : '';
    console.debug('Fetching alert metadata for', name);
    pullAlertMetadata({name, updatesOnly: false}, alertMetadata);
  }

  // Ui Error Handling
  const errorStore = useErrorStore();
  let unwatchErrors;
  onMounted(() => {
    unwatchErrors = errorStore.registerValue(alertMetadata);
  });
  onUnmounted(() => {
    if (unwatchErrors) unwatchErrors();
  });

  /**
   * Converts a proto map, which is an array of [k,v] into a js object.
   *
   * @param {Array<[K,V]>} arr
   * @return {Object<K,V>}
   * @template K,V
   */
  function convertProtoMap(arr) {
    if (!arr) return {};
    const dst = {};
    for (const [k, v] of arr || []) {
      dst[k] = v;
    }
    return dst;
  }

  const acknowledgedCountMap = computed(() => convertProtoMap(alertMetadata.value?.acknowledgedCountsMap));
  const floorCountsMap = computed(() => convertProtoMap(alertMetadata.value?.floorCountsMap));
  const zoneCountsMap = computed(() => convertProtoMap(alertMetadata.value?.zoneCountsMap));
  const severityCountsMap = computed(() => convertProtoMap(alertMetadata.value?.severityCountsMap));

  const unacknowledgedAlertCount = computed(() => acknowledgedCountMap.value[false]);

  return {
    alertMetadata,

    acknowledgedCountMap,
    floorCountsMap,
    zoneCountsMap,
    severityCountsMap,

    unacknowledgedAlertCount
  };
});
