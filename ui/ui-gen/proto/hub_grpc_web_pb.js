/**
 * @fileoverview gRPC-Web generated client stub for smartcore.bos
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.2
// 	protoc              v3.21.12
// source: hub.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var traits_metadata_pb = require('@smart-core-os/sc-api-grpc-web/traits/metadata_pb.js')
const proto = {};
proto.smartcore = {};
proto.smartcore.bos = require('./hub_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.bos.HubApiClient =
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
proto.smartcore.bos.HubApiPromiseClient =
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
 *   !proto.smartcore.bos.GetHubNodeRequest,
 *   !proto.smartcore.bos.HubNode>}
 */
const methodDescriptor_HubApi_GetHubNode = new grpc.web.MethodDescriptor(
  '/smartcore.bos.HubApi/GetHubNode',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.GetHubNodeRequest,
  proto.smartcore.bos.HubNode,
  /**
   * @param {!proto.smartcore.bos.GetHubNodeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.HubNode.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.GetHubNodeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.HubNode)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.HubNode>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.HubApiClient.prototype.getHubNode =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.HubApi/GetHubNode',
      request,
      metadata || {},
      methodDescriptor_HubApi_GetHubNode,
      callback);
};


/**
 * @param {!proto.smartcore.bos.GetHubNodeRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.HubNode>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.HubApiPromiseClient.prototype.getHubNode =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.HubApi/GetHubNode',
      request,
      metadata || {},
      methodDescriptor_HubApi_GetHubNode);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.ListHubNodesRequest,
 *   !proto.smartcore.bos.ListHubNodesResponse>}
 */
const methodDescriptor_HubApi_ListHubNodes = new grpc.web.MethodDescriptor(
  '/smartcore.bos.HubApi/ListHubNodes',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.ListHubNodesRequest,
  proto.smartcore.bos.ListHubNodesResponse,
  /**
   * @param {!proto.smartcore.bos.ListHubNodesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.ListHubNodesResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.ListHubNodesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.ListHubNodesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.ListHubNodesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.HubApiClient.prototype.listHubNodes =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.HubApi/ListHubNodes',
      request,
      metadata || {},
      methodDescriptor_HubApi_ListHubNodes,
      callback);
};


/**
 * @param {!proto.smartcore.bos.ListHubNodesRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.ListHubNodesResponse>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.HubApiPromiseClient.prototype.listHubNodes =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.HubApi/ListHubNodes',
      request,
      metadata || {},
      methodDescriptor_HubApi_ListHubNodes);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.InspectHubNodeRequest,
 *   !proto.smartcore.bos.HubNodeInspection>}
 */
const methodDescriptor_HubApi_InspectHubNode = new grpc.web.MethodDescriptor(
  '/smartcore.bos.HubApi/InspectHubNode',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.InspectHubNodeRequest,
  proto.smartcore.bos.HubNodeInspection,
  /**
   * @param {!proto.smartcore.bos.InspectHubNodeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.HubNodeInspection.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.InspectHubNodeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.HubNodeInspection)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.HubNodeInspection>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.HubApiClient.prototype.inspectHubNode =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.HubApi/InspectHubNode',
      request,
      metadata || {},
      methodDescriptor_HubApi_InspectHubNode,
      callback);
};


/**
 * @param {!proto.smartcore.bos.InspectHubNodeRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.HubNodeInspection>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.HubApiPromiseClient.prototype.inspectHubNode =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.HubApi/InspectHubNode',
      request,
      metadata || {},
      methodDescriptor_HubApi_InspectHubNode);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.EnrollHubNodeRequest,
 *   !proto.smartcore.bos.HubNode>}
 */
const methodDescriptor_HubApi_EnrollHubNode = new grpc.web.MethodDescriptor(
  '/smartcore.bos.HubApi/EnrollHubNode',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.EnrollHubNodeRequest,
  proto.smartcore.bos.HubNode,
  /**
   * @param {!proto.smartcore.bos.EnrollHubNodeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.HubNode.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.EnrollHubNodeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.HubNode)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.HubNode>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.HubApiClient.prototype.enrollHubNode =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.HubApi/EnrollHubNode',
      request,
      metadata || {},
      methodDescriptor_HubApi_EnrollHubNode,
      callback);
};


/**
 * @param {!proto.smartcore.bos.EnrollHubNodeRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.HubNode>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.HubApiPromiseClient.prototype.enrollHubNode =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.HubApi/EnrollHubNode',
      request,
      metadata || {},
      methodDescriptor_HubApi_EnrollHubNode);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.RenewHubNodeRequest,
 *   !proto.smartcore.bos.HubNode>}
 */
const methodDescriptor_HubApi_RenewHubNode = new grpc.web.MethodDescriptor(
  '/smartcore.bos.HubApi/RenewHubNode',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.RenewHubNodeRequest,
  proto.smartcore.bos.HubNode,
  /**
   * @param {!proto.smartcore.bos.RenewHubNodeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.HubNode.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.RenewHubNodeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.HubNode)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.HubNode>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.HubApiClient.prototype.renewHubNode =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.HubApi/RenewHubNode',
      request,
      metadata || {},
      methodDescriptor_HubApi_RenewHubNode,
      callback);
};


/**
 * @param {!proto.smartcore.bos.RenewHubNodeRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.HubNode>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.HubApiPromiseClient.prototype.renewHubNode =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.HubApi/RenewHubNode',
      request,
      metadata || {},
      methodDescriptor_HubApi_RenewHubNode);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.TestHubNodeRequest,
 *   !proto.smartcore.bos.TestHubNodeResponse>}
 */
const methodDescriptor_HubApi_TestHubNode = new grpc.web.MethodDescriptor(
  '/smartcore.bos.HubApi/TestHubNode',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.TestHubNodeRequest,
  proto.smartcore.bos.TestHubNodeResponse,
  /**
   * @param {!proto.smartcore.bos.TestHubNodeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.TestHubNodeResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.TestHubNodeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.TestHubNodeResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.TestHubNodeResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.HubApiClient.prototype.testHubNode =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.HubApi/TestHubNode',
      request,
      metadata || {},
      methodDescriptor_HubApi_TestHubNode,
      callback);
};


/**
 * @param {!proto.smartcore.bos.TestHubNodeRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.TestHubNodeResponse>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.HubApiPromiseClient.prototype.testHubNode =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.HubApi/TestHubNode',
      request,
      metadata || {},
      methodDescriptor_HubApi_TestHubNode);
};


module.exports = proto.smartcore.bos;

