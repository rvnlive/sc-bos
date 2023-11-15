import {LightingTestApiPromiseClient} from '@sc-bos/ui-gen/proto/lighting_test_grpc_web_pb';
import {GetReportCSVRequest, ListLightHealthRequest} from '@sc-bos/ui-gen/proto/lighting_test_pb';
import {clientOptions} from '@/api/grpcweb';
import {trackAction} from '@/api/resource';
import {DaliApiPromiseClient} from '@sc-bos/ui-gen/proto/dali_grpc_web_pb';
import {StartTestRequest} from '@sc-bos/ui-gen/proto/dali_pb';
import {setProperties} from '@/api/convpb';

/**
 *
 * @param {ActionTracker<ReportCSV.AsObject>} [tracker]
 * @return {Promise<ReportCSV.AsObject>}
 */
export function getReportCSV(tracker) {
  return trackAction('LightingTest.getReportCSV', tracker ?? {}, endpoint => {
    const api = client(endpoint);
    const req = new GetReportCSVRequest();
    req.setIncludeHeader(true);
    return api.getReportCSV(req);
  });
}

/**
 *
 * @param {Partial<ListLightHealthRequest.AsObject>} request
 * @param {ActionTracker<ListLightHealthResponse.AsObject>} [tracker]
 * @return {Promise<ListLightHealthResponse.AsObject>}
 */
export function listLightHealth(request, tracker) {
  return trackAction('LightingTest.listLightHealth', tracker ?? {}, endpoint => {
    const api = client(endpoint);
    return api.listLightHealth(listLightHealthRequestFromObject(request));
  });
}

/**
 *
 * @param {StartTestRequest.AsObject} request
 * @param {ActionTracker<StartTestRequest.AsObject>} [tracker]
 * @return {Promise<StartTestResponse.AsObject>}
 */
export function runTest(request, tracker) {
  return trackAction('Dali.StartTest', tracker ?? {}, endpoint => {
    const api = new DaliApiPromiseClient(endpoint, null, clientOptions());
    return api.startTest(startTestRequestFromObject(request));
  });
}

/**
 *
 * @param {number} faultId
 * @return {string}
 */
export function faultToString(faultId) {
  switch (faultId) {
    case 1: return 'DURATION_TEST_FAILED';
    case 2: return 'FUNCTION_TEST_FAILED';
    case 3: return 'BATTERY_FAULT';
    case 4: return 'LAMP_FAULT';
    case 5: return 'COMMUNICATION_FAILURE';
    case 6: return 'OTHER_FAULT';
    default: return 'FAULT_UNSPECIFIED';
  }
}

/**
 *
 * @param {string} endpoint
 * @return {LightingTestApiPromiseClient}
 */
function client(endpoint) {
  return new LightingTestApiPromiseClient(endpoint, null, clientOptions());
}

/**
 * @param {StartTestRequest.AsObject} obj
 * @return {StartTestRequest}
 */
function startTestRequestFromObject(obj) {
  if (!obj) return undefined;

  const req = new StartTestRequest();
  setProperties(req, obj, 'name', 'test');
  return req;
}

/**
 * @param {ListLightHealthRequest.AsObject} obj
 * @return {undefined|ListLightHealthRequest}
 */
function listLightHealthRequestFromObject(obj) {
  if (!obj) return undefined;

  const dst = new ListLightHealthRequest();
  setProperties(dst, obj, 'pageSize', 'pageToken');
  return dst;
}
