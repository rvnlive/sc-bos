import {closeResource, newActionTracker, newResourceCollection} from '@/api/resource.js';
import {acknowledgeAlert, listAlerts, pullAlerts, unacknowledgeAlert} from '@/api/ui/alerts.js';
import {Collection} from '@/util/query.js';
import {Alert, ListAlertsResponse} from '@bsp-ew/ui-gen/proto/alerts_pb';
import {acceptHMRUpdate, defineStore} from 'pinia';
import {computed, reactive, set, watch} from 'vue';


const SeverityStrings = {
  [Alert.Severity.INFO]: 'INFO',
  [Alert.Severity.WARNING]: 'WARN',
  [Alert.Severity.SEVERE]: 'ALERT',
  [Alert.Severity.LIFE_SAFETY]: 'DANGER'
};
const SeverityColor = {
  [Alert.Severity.INFO]: 'info',
  [Alert.Severity.WARNING]: 'warning',
  [Alert.Severity.SEVERE]: 'error',
  [Alert.Severity.LIFE_SAFETY]: 'error'
};

export const useNotifications = defineStore('notifications', () => {
  // todo: get the name from somewhere
  const name = computed(() => 'test-ac');

  const alerts = reactive(/** @type {ResourceCollection<Alert.AsObject, Alert>} */newResourceCollection()); // holds all the alerts we can show
  const fetchingPage = reactive(/** @type {ActionTracker<ListAlertsResponse.AsObject>} */ newActionTracker()); // tracks the fetching of a single page

  watch(name, async name => {
    closeResource(alerts);
    pullAlerts({name}, alerts);
    try {
      const firstPage = await listAlerts({name, pageSize: 100, pageToken: undefined}, fetchingPage);
      for (let alert of firstPage.alertsList) {
        set(alerts.value, alert.id, alert);
      }
      fetchingPage.response = null;
    } catch (e) {
      console.warn('Error fetching first page', e);
    }
  }, {immediate: true});

  function severityData(severity) {
    for (let i = severity; i > 0; i--) {
      if (SeverityStrings[i]) {
        let str = SeverityStrings[i];
        if (i < severity) {
          str += '+' + (severity - i);
        }
        return {text: str, color: `${SeverityColor[i]}--text`};
      }
    }
    return {text: 'unspecified', color: 'gray--text'};
  }

  function setAcknowledged(e, alert) {
    if (e) {
      acknowledgeAlert({name: name.value, id: alert.id, allowAcknowledged: false, allowMissing: false})
          .catch(err => console.error(err));
    } else {
      unacknowledgeAlert({name: name.value, id: alert.id, allowAcknowledged: false, allowMissing: false})
          .catch(err => console.error(err));
    }
  }

  function isAcknowledged(alert) {
    return Boolean(alert.acknowledgement);
  }

  function newCollection() {
    const listFn = async (query, tracker, pageToken, recordFn) => {
      const page = await listAlerts({name: name.value, pageToken, query, pageSize: 100}, tracker);
      for (let alert of page.alertsList) {
        recordFn(alert, alert.id);
      }
      return page.nextPageToken
    }
    const pullFn = (query, resources) => {
      pullAlerts({name: name.value, query}, resources);
    }
    return new Collection(listFn, pullFn);
  }

  return {
    name,
    alerts,
    loading: computed(() => alerts.loading || fetchingPage.loading),
    newCollection,
    severityData,
    setAcknowledged,
    isAcknowledged
  }
});

// enable hot reload for this store
if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useNotifications, import.meta.hot))
}