/**
 * @fileoverview gRPC-Web generated client stub for smartcore.bos
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v5.29.1
// source: mqtt.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.smartcore = {};
proto.smartcore.bos = require('./mqtt_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.bos.MqttServiceClient =
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
proto.smartcore.bos.MqttServicePromiseClient =
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
 *   !proto.smartcore.bos.PullMessagesRequest,
 *   !proto.smartcore.bos.PullMessagesResponse>}
 */
const methodDescriptor_MqttService_PullMessages = new grpc.web.MethodDescriptor(
  '/smartcore.bos.MqttService/PullMessages',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.bos.PullMessagesRequest,
  proto.smartcore.bos.PullMessagesResponse,
  /**
   * @param {!proto.smartcore.bos.PullMessagesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.PullMessagesResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.PullMessagesRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.PullMessagesResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.MqttServiceClient.prototype.pullMessages =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.bos.MqttService/PullMessages',
      request,
      metadata || {},
      methodDescriptor_MqttService_PullMessages);
};


/**
 * @param {!proto.smartcore.bos.PullMessagesRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.PullMessagesResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.MqttServicePromiseClient.prototype.pullMessages =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.bos.MqttService/PullMessages',
      request,
      metadata || {},
      methodDescriptor_MqttService_PullMessages);
};


module.exports = proto.smartcore.bos;

