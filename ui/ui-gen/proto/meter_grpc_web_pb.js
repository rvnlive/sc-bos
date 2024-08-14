/**
 * @fileoverview gRPC-Web generated client stub for smartcore.bos
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v5.27.3
// source: meter.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_field_mask_pb = require('google-protobuf/google/protobuf/field_mask_pb.js')

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')

var types_info_pb = require('@smart-core-os/sc-api-grpc-web/types/info_pb.js')
const proto = {};
proto.smartcore = {};
proto.smartcore.bos = require('./meter_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.bos.MeterApiClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.bos.MeterApiPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.GetMeterReadingRequest,
 *   !proto.smartcore.bos.MeterReading>}
 */
const methodDescriptor_MeterApi_GetMeterReading = new grpc.web.MethodDescriptor(
  '/smartcore.bos.MeterApi/GetMeterReading',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.GetMeterReadingRequest,
  proto.smartcore.bos.MeterReading,
  /**
   * @param {!proto.smartcore.bos.GetMeterReadingRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.MeterReading.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.GetMeterReadingRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.MeterReading)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.MeterReading>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.MeterApiClient.prototype.getMeterReading =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.MeterApi/GetMeterReading',
      request,
      metadata || {},
      methodDescriptor_MeterApi_GetMeterReading,
      callback);
};


/**
 * @param {!proto.smartcore.bos.GetMeterReadingRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.MeterReading>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.MeterApiPromiseClient.prototype.getMeterReading =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.MeterApi/GetMeterReading',
      request,
      metadata || {},
      methodDescriptor_MeterApi_GetMeterReading);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.PullMeterReadingsRequest,
 *   !proto.smartcore.bos.PullMeterReadingsResponse>}
 */
const methodDescriptor_MeterApi_PullMeterReadings = new grpc.web.MethodDescriptor(
  '/smartcore.bos.MeterApi/PullMeterReadings',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.bos.PullMeterReadingsRequest,
  proto.smartcore.bos.PullMeterReadingsResponse,
  /**
   * @param {!proto.smartcore.bos.PullMeterReadingsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.PullMeterReadingsResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.PullMeterReadingsRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.PullMeterReadingsResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.MeterApiClient.prototype.pullMeterReadings =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.bos.MeterApi/PullMeterReadings',
      request,
      metadata || {},
      methodDescriptor_MeterApi_PullMeterReadings);
};


/**
 * @param {!proto.smartcore.bos.PullMeterReadingsRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.PullMeterReadingsResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.MeterApiPromiseClient.prototype.pullMeterReadings =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.bos.MeterApi/PullMeterReadings',
      request,
      metadata || {},
      methodDescriptor_MeterApi_PullMeterReadings);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.bos.MeterInfoClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.bos.MeterInfoPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.DescribeMeterReadingRequest,
 *   !proto.smartcore.bos.MeterReadingSupport>}
 */
const methodDescriptor_MeterInfo_DescribeMeterReading = new grpc.web.MethodDescriptor(
  '/smartcore.bos.MeterInfo/DescribeMeterReading',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.DescribeMeterReadingRequest,
  proto.smartcore.bos.MeterReadingSupport,
  /**
   * @param {!proto.smartcore.bos.DescribeMeterReadingRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.MeterReadingSupport.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.DescribeMeterReadingRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.MeterReadingSupport)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.MeterReadingSupport>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.MeterInfoClient.prototype.describeMeterReading =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.MeterInfo/DescribeMeterReading',
      request,
      metadata || {},
      methodDescriptor_MeterInfo_DescribeMeterReading,
      callback);
};


/**
 * @param {!proto.smartcore.bos.DescribeMeterReadingRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.MeterReadingSupport>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.MeterInfoPromiseClient.prototype.describeMeterReading =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.MeterInfo/DescribeMeterReading',
      request,
      metadata || {},
      methodDescriptor_MeterInfo_DescribeMeterReading);
};


module.exports = proto.smartcore.bos;

