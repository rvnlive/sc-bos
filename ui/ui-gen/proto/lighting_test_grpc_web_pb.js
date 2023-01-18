/**
 * @fileoverview gRPC-Web generated client stub for smartcore.bos
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')

var google_protobuf_duration_pb = require('google-protobuf/google/protobuf/duration_pb.js')
const proto = {};
proto.smartcore = {};
proto.smartcore.bos = require('./lighting_test_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.bos.LightingTestApiClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.bos.LightingTestApiPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.GetLightHealthRequest,
 *   !proto.smartcore.bos.LightHealth>}
 */
const methodDescriptor_LightingTestApi_GetLightHealth = new grpc.web.MethodDescriptor(
  '/smartcore.bos.LightingTestApi/GetLightHealth',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.GetLightHealthRequest,
  proto.smartcore.bos.LightHealth,
  /**
   * @param {!proto.smartcore.bos.GetLightHealthRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.LightHealth.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.smartcore.bos.GetLightHealthRequest,
 *   !proto.smartcore.bos.LightHealth>}
 */
const methodInfo_LightingTestApi_GetLightHealth = new grpc.web.AbstractClientBase.MethodInfo(
  proto.smartcore.bos.LightHealth,
  /**
   * @param {!proto.smartcore.bos.GetLightHealthRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.LightHealth.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.GetLightHealthRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.smartcore.bos.LightHealth)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.LightHealth>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.LightingTestApiClient.prototype.getLightHealth =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.LightingTestApi/GetLightHealth',
      request,
      metadata || {},
      methodDescriptor_LightingTestApi_GetLightHealth,
      callback);
};


/**
 * @param {!proto.smartcore.bos.GetLightHealthRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.LightHealth>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.LightingTestApiPromiseClient.prototype.getLightHealth =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.LightingTestApi/GetLightHealth',
      request,
      metadata || {},
      methodDescriptor_LightingTestApi_GetLightHealth);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.ListLightHealthRequest,
 *   !proto.smartcore.bos.ListLightHealthResponse>}
 */
const methodDescriptor_LightingTestApi_ListLightHealth = new grpc.web.MethodDescriptor(
  '/smartcore.bos.LightingTestApi/ListLightHealth',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.ListLightHealthRequest,
  proto.smartcore.bos.ListLightHealthResponse,
  /**
   * @param {!proto.smartcore.bos.ListLightHealthRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.ListLightHealthResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.smartcore.bos.ListLightHealthRequest,
 *   !proto.smartcore.bos.ListLightHealthResponse>}
 */
const methodInfo_LightingTestApi_ListLightHealth = new grpc.web.AbstractClientBase.MethodInfo(
  proto.smartcore.bos.ListLightHealthResponse,
  /**
   * @param {!proto.smartcore.bos.ListLightHealthRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.ListLightHealthResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.ListLightHealthRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.smartcore.bos.ListLightHealthResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.ListLightHealthResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.LightingTestApiClient.prototype.listLightHealth =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.LightingTestApi/ListLightHealth',
      request,
      metadata || {},
      methodDescriptor_LightingTestApi_ListLightHealth,
      callback);
};


/**
 * @param {!proto.smartcore.bos.ListLightHealthRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.ListLightHealthResponse>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.LightingTestApiPromiseClient.prototype.listLightHealth =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.LightingTestApi/ListLightHealth',
      request,
      metadata || {},
      methodDescriptor_LightingTestApi_ListLightHealth);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.ListLightEventsRequest,
 *   !proto.smartcore.bos.ListLightEventsResponse>}
 */
const methodDescriptor_LightingTestApi_ListLightEvents = new grpc.web.MethodDescriptor(
  '/smartcore.bos.LightingTestApi/ListLightEvents',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.ListLightEventsRequest,
  proto.smartcore.bos.ListLightEventsResponse,
  /**
   * @param {!proto.smartcore.bos.ListLightEventsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.ListLightEventsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.smartcore.bos.ListLightEventsRequest,
 *   !proto.smartcore.bos.ListLightEventsResponse>}
 */
const methodInfo_LightingTestApi_ListLightEvents = new grpc.web.AbstractClientBase.MethodInfo(
  proto.smartcore.bos.ListLightEventsResponse,
  /**
   * @param {!proto.smartcore.bos.ListLightEventsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.ListLightEventsResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.ListLightEventsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.smartcore.bos.ListLightEventsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.ListLightEventsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.LightingTestApiClient.prototype.listLightEvents =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.LightingTestApi/ListLightEvents',
      request,
      metadata || {},
      methodDescriptor_LightingTestApi_ListLightEvents,
      callback);
};


/**
 * @param {!proto.smartcore.bos.ListLightEventsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.ListLightEventsResponse>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.LightingTestApiPromiseClient.prototype.listLightEvents =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.LightingTestApi/ListLightEvents',
      request,
      metadata || {},
      methodDescriptor_LightingTestApi_ListLightEvents);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.GetReportCSVRequest,
 *   !proto.smartcore.bos.ReportCSV>}
 */
const methodDescriptor_LightingTestApi_GetReportCSV = new grpc.web.MethodDescriptor(
  '/smartcore.bos.LightingTestApi/GetReportCSV',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.GetReportCSVRequest,
  proto.smartcore.bos.ReportCSV,
  /**
   * @param {!proto.smartcore.bos.GetReportCSVRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.ReportCSV.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.smartcore.bos.GetReportCSVRequest,
 *   !proto.smartcore.bos.ReportCSV>}
 */
const methodInfo_LightingTestApi_GetReportCSV = new grpc.web.AbstractClientBase.MethodInfo(
  proto.smartcore.bos.ReportCSV,
  /**
   * @param {!proto.smartcore.bos.GetReportCSVRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.ReportCSV.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.GetReportCSVRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.smartcore.bos.ReportCSV)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.ReportCSV>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.LightingTestApiClient.prototype.getReportCSV =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.LightingTestApi/GetReportCSV',
      request,
      metadata || {},
      methodDescriptor_LightingTestApi_GetReportCSV,
      callback);
};


/**
 * @param {!proto.smartcore.bos.GetReportCSVRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.ReportCSV>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.LightingTestApiPromiseClient.prototype.getReportCSV =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.LightingTestApi/GetReportCSV',
      request,
      metadata || {},
      methodDescriptor_LightingTestApi_GetReportCSV);
};


module.exports = proto.smartcore.bos;
