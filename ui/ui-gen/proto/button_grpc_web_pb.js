/**
 * @fileoverview gRPC-Web generated client stub for smartcore.bos
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v5.27.3
// source: button.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')

var google_protobuf_field_mask_pb = require('google-protobuf/google/protobuf/field_mask_pb.js')
const proto = {};
proto.smartcore = {};
proto.smartcore.bos = require('./button_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.bos.ButtonApiClient =
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
proto.smartcore.bos.ButtonApiPromiseClient =
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
 *   !proto.smartcore.bos.GetButtonStateRequest,
 *   !proto.smartcore.bos.ButtonState>}
 */
const methodDescriptor_ButtonApi_GetButtonState = new grpc.web.MethodDescriptor(
  '/smartcore.bos.ButtonApi/GetButtonState',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.GetButtonStateRequest,
  proto.smartcore.bos.ButtonState,
  /**
   * @param {!proto.smartcore.bos.GetButtonStateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.ButtonState.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.GetButtonStateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.ButtonState)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.ButtonState>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.ButtonApiClient.prototype.getButtonState =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.ButtonApi/GetButtonState',
      request,
      metadata || {},
      methodDescriptor_ButtonApi_GetButtonState,
      callback);
};


/**
 * @param {!proto.smartcore.bos.GetButtonStateRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.ButtonState>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.ButtonApiPromiseClient.prototype.getButtonState =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.ButtonApi/GetButtonState',
      request,
      metadata || {},
      methodDescriptor_ButtonApi_GetButtonState);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.PullButtonStateRequest,
 *   !proto.smartcore.bos.PullButtonStateResponse>}
 */
const methodDescriptor_ButtonApi_PullButtonState = new grpc.web.MethodDescriptor(
  '/smartcore.bos.ButtonApi/PullButtonState',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.bos.PullButtonStateRequest,
  proto.smartcore.bos.PullButtonStateResponse,
  /**
   * @param {!proto.smartcore.bos.PullButtonStateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.PullButtonStateResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.PullButtonStateRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.PullButtonStateResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.ButtonApiClient.prototype.pullButtonState =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.bos.ButtonApi/PullButtonState',
      request,
      metadata || {},
      methodDescriptor_ButtonApi_PullButtonState);
};


/**
 * @param {!proto.smartcore.bos.PullButtonStateRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.PullButtonStateResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.ButtonApiPromiseClient.prototype.pullButtonState =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.bos.ButtonApi/PullButtonState',
      request,
      metadata || {},
      methodDescriptor_ButtonApi_PullButtonState);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.UpdateButtonStateRequest,
 *   !proto.smartcore.bos.ButtonState>}
 */
const methodDescriptor_ButtonApi_UpdateButtonState = new grpc.web.MethodDescriptor(
  '/smartcore.bos.ButtonApi/UpdateButtonState',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.UpdateButtonStateRequest,
  proto.smartcore.bos.ButtonState,
  /**
   * @param {!proto.smartcore.bos.UpdateButtonStateRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.ButtonState.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.UpdateButtonStateRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.ButtonState)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.ButtonState>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.ButtonApiClient.prototype.updateButtonState =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.ButtonApi/UpdateButtonState',
      request,
      metadata || {},
      methodDescriptor_ButtonApi_UpdateButtonState,
      callback);
};


/**
 * @param {!proto.smartcore.bos.UpdateButtonStateRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.ButtonState>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.ButtonApiPromiseClient.prototype.updateButtonState =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.ButtonApi/UpdateButtonState',
      request,
      metadata || {},
      methodDescriptor_ButtonApi_UpdateButtonState);
};


module.exports = proto.smartcore.bos;

