import {grpcWebEndpoint} from '@/api/config';
import {closeResource, newActionTracker, newResourceCollection} from '@/api/resource';
import {getEnrollment} from '@/api/sc/traits/enrollment';
import {listHubNodes, pullHubNodes} from '@/api/sc/traits/hub';
import {useAppConfigStore} from '@/stores/app-config';
import {defineStore} from 'pinia';
import {computed, reactive, set, watch} from 'vue';

export const useHubStore = defineStore('hub', () => {
  const appConfig = useAppConfigStore();
  const nodesListCollection = reactive(newResourceCollection());

  watch(() => appConfig.config, async config => {
    closeResource(nodesListCollection);

    if (config?.hub) {
      pullHubNodes(nodesListCollection);
      try {
        if (!nodesListCollection.value) set(nodesListCollection, 'value', {});
        const hubNode = await getEnrollment(newActionTracker());
        set(nodesListCollection.value, hubNode.managerName, {
          name: hubNode.managerName,
          address: hubNode.managerAddress
        });
        const nodes = await listHubNodes(newActionTracker());
        for (const node of nodes.nodesList) {
          set(nodesListCollection.value, node.name, node);
        }
      } catch (e) {
        console.warn('Error fetching first page', e);
      }
    }
  }, {immediate: true});

  /**
   * @typedef {Object} Node
   * @property {string} name - the Smart Core name of the node
   * @property {string} address - the address of the node
   * @property {string} description - a human-readable description of the node
   * @property {string} commsAddress - the address to use to communicate with the node (based on proxy settings)
   * @property {string} commsName - the name to use to communicate with the node (based on proxy settings)
   */

  const nodesList = computed(() => {
    /** @type {Record<string, Node>} */
    const nodes = {};
    Object.values(nodesListCollection?.value || {}).forEach((node, name) => {
      console.log('node', node);
      nodes[node.name] = {
        ...node,
        commsAddress: proxiedAddress(node.address),
        commsName: proxiedName(node.name, node.address)
      };
    });
    return nodes;
  });


  /**
   * If we're communicating with named devices via a proxy, this will return prepend the controller name to the
   * resource, otherwise it will return the resource name as-is.
   *
   * @param {string} name
   * @param {string} address
   * @return {string}
   */
  async function proxiedName(name, address) {
    // check if running in proxy mode, and that the node address is not the same as our endpoint address
    if (appConfig.config?.proxy && (await grpcWebEndpoint()) !== address) {
      return name;
    }
    return '';
  }

  /**
   * If we're communicating with named devices via a proxy, this will return the proxy address, otherwise it will return
   * the address as-is.
   *
   * @param {string} address
   * @return {string}
   */
  async function proxiedAddress(address) {
    if (appConfig.config?.proxy) {
      return await grpcWebEndpoint();
    }
    return address;
  }

  return {
    nodesList
  };
});
