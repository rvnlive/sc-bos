/**
 * @fileoverview gRPC-Web generated client stub for smartcore.bos
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.2
// 	protoc              v3.21.12
// source: services.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_field_mask_pb = require('google-protobuf/google/protobuf/field_mask_pb.js')

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')

var types_change_pb = require('@smart-core-os/sc-api-grpc-web/types/change_pb.js')
const proto = {};
proto.smartcore = {};
proto.smartcore.bos = require('./services_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.smartcore.bos.ServicesApiClient =
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
proto.smartcore.bos.ServicesApiPromiseClient =
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
 *   !proto.smartcore.bos.GetServiceRequest,
 *   !proto.smartcore.bos.Service>}
 */
const methodDescriptor_ServicesApi_GetService = new grpc.web.MethodDescriptor(
  '/smartcore.bos.ServicesApi/GetService',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.GetServiceRequest,
  proto.smartcore.bos.Service,
  /**
   * @param {!proto.smartcore.bos.GetServiceRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.Service.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.GetServiceRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.Service)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.Service>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.ServicesApiClient.prototype.getService =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.ServicesApi/GetService',
      request,
      metadata || {},
      methodDescriptor_ServicesApi_GetService,
      callback);
};


/**
 * @param {!proto.smartcore.bos.GetServiceRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.Service>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.ServicesApiPromiseClient.prototype.getService =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.ServicesApi/GetService',
      request,
      metadata || {},
      methodDescriptor_ServicesApi_GetService);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.PullServiceRequest,
 *   !proto.smartcore.bos.PullServiceResponse>}
 */
const methodDescriptor_ServicesApi_PullService = new grpc.web.MethodDescriptor(
  '/smartcore.bos.ServicesApi/PullService',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.bos.PullServiceRequest,
  proto.smartcore.bos.PullServiceResponse,
  /**
   * @param {!proto.smartcore.bos.PullServiceRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.PullServiceResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.PullServiceRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.PullServiceResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.ServicesApiClient.prototype.pullService =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.bos.ServicesApi/PullService',
      request,
      metadata || {},
      methodDescriptor_ServicesApi_PullService);
};


/**
 * @param {!proto.smartcore.bos.PullServiceRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.PullServiceResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.ServicesApiPromiseClient.prototype.pullService =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.bos.ServicesApi/PullService',
      request,
      metadata || {},
      methodDescriptor_ServicesApi_PullService);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.CreateServiceRequest,
 *   !proto.smartcore.bos.Service>}
 */
const methodDescriptor_ServicesApi_CreateService = new grpc.web.MethodDescriptor(
  '/smartcore.bos.ServicesApi/CreateService',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.CreateServiceRequest,
  proto.smartcore.bos.Service,
  /**
   * @param {!proto.smartcore.bos.CreateServiceRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.Service.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.CreateServiceRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.Service)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.Service>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.ServicesApiClient.prototype.createService =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.ServicesApi/CreateService',
      request,
      metadata || {},
      methodDescriptor_ServicesApi_CreateService,
      callback);
};


/**
 * @param {!proto.smartcore.bos.CreateServiceRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.Service>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.ServicesApiPromiseClient.prototype.createService =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.ServicesApi/CreateService',
      request,
      metadata || {},
      methodDescriptor_ServicesApi_CreateService);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.DeleteServiceRequest,
 *   !proto.smartcore.bos.Service>}
 */
const methodDescriptor_ServicesApi_DeleteService = new grpc.web.MethodDescriptor(
  '/smartcore.bos.ServicesApi/DeleteService',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.DeleteServiceRequest,
  proto.smartcore.bos.Service,
  /**
   * @param {!proto.smartcore.bos.DeleteServiceRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.Service.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.DeleteServiceRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.Service)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.Service>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.ServicesApiClient.prototype.deleteService =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.ServicesApi/DeleteService',
      request,
      metadata || {},
      methodDescriptor_ServicesApi_DeleteService,
      callback);
};


/**
 * @param {!proto.smartcore.bos.DeleteServiceRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.Service>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.ServicesApiPromiseClient.prototype.deleteService =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.ServicesApi/DeleteService',
      request,
      metadata || {},
      methodDescriptor_ServicesApi_DeleteService);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.ListServicesRequest,
 *   !proto.smartcore.bos.ListServicesResponse>}
 */
const methodDescriptor_ServicesApi_ListServices = new grpc.web.MethodDescriptor(
  '/smartcore.bos.ServicesApi/ListServices',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.ListServicesRequest,
  proto.smartcore.bos.ListServicesResponse,
  /**
   * @param {!proto.smartcore.bos.ListServicesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.ListServicesResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.ListServicesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.ListServicesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.ListServicesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.ServicesApiClient.prototype.listServices =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.ServicesApi/ListServices',
      request,
      metadata || {},
      methodDescriptor_ServicesApi_ListServices,
      callback);
};


/**
 * @param {!proto.smartcore.bos.ListServicesRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.ListServicesResponse>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.ServicesApiPromiseClient.prototype.listServices =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.ServicesApi/ListServices',
      request,
      metadata || {},
      methodDescriptor_ServicesApi_ListServices);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.PullServicesRequest,
 *   !proto.smartcore.bos.PullServicesResponse>}
 */
const methodDescriptor_ServicesApi_PullServices = new grpc.web.MethodDescriptor(
  '/smartcore.bos.ServicesApi/PullServices',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.bos.PullServicesRequest,
  proto.smartcore.bos.PullServicesResponse,
  /**
   * @param {!proto.smartcore.bos.PullServicesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.PullServicesResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.PullServicesRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.PullServicesResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.ServicesApiClient.prototype.pullServices =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.bos.ServicesApi/PullServices',
      request,
      metadata || {},
      methodDescriptor_ServicesApi_PullServices);
};


/**
 * @param {!proto.smartcore.bos.PullServicesRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.PullServicesResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.ServicesApiPromiseClient.prototype.pullServices =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.bos.ServicesApi/PullServices',
      request,
      metadata || {},
      methodDescriptor_ServicesApi_PullServices);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.StartServiceRequest,
 *   !proto.smartcore.bos.Service>}
 */
const methodDescriptor_ServicesApi_StartService = new grpc.web.MethodDescriptor(
  '/smartcore.bos.ServicesApi/StartService',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.StartServiceRequest,
  proto.smartcore.bos.Service,
  /**
   * @param {!proto.smartcore.bos.StartServiceRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.Service.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.StartServiceRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.Service)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.Service>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.ServicesApiClient.prototype.startService =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.ServicesApi/StartService',
      request,
      metadata || {},
      methodDescriptor_ServicesApi_StartService,
      callback);
};


/**
 * @param {!proto.smartcore.bos.StartServiceRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.Service>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.ServicesApiPromiseClient.prototype.startService =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.ServicesApi/StartService',
      request,
      metadata || {},
      methodDescriptor_ServicesApi_StartService);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.ConfigureServiceRequest,
 *   !proto.smartcore.bos.Service>}
 */
const methodDescriptor_ServicesApi_ConfigureService = new grpc.web.MethodDescriptor(
  '/smartcore.bos.ServicesApi/ConfigureService',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.ConfigureServiceRequest,
  proto.smartcore.bos.Service,
  /**
   * @param {!proto.smartcore.bos.ConfigureServiceRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.Service.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.ConfigureServiceRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.Service)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.Service>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.ServicesApiClient.prototype.configureService =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.ServicesApi/ConfigureService',
      request,
      metadata || {},
      methodDescriptor_ServicesApi_ConfigureService,
      callback);
};


/**
 * @param {!proto.smartcore.bos.ConfigureServiceRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.Service>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.ServicesApiPromiseClient.prototype.configureService =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.ServicesApi/ConfigureService',
      request,
      metadata || {},
      methodDescriptor_ServicesApi_ConfigureService);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.StopServiceRequest,
 *   !proto.smartcore.bos.Service>}
 */
const methodDescriptor_ServicesApi_StopService = new grpc.web.MethodDescriptor(
  '/smartcore.bos.ServicesApi/StopService',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.StopServiceRequest,
  proto.smartcore.bos.Service,
  /**
   * @param {!proto.smartcore.bos.StopServiceRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.Service.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.StopServiceRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.Service)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.Service>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.ServicesApiClient.prototype.stopService =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.ServicesApi/StopService',
      request,
      metadata || {},
      methodDescriptor_ServicesApi_StopService,
      callback);
};


/**
 * @param {!proto.smartcore.bos.StopServiceRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.Service>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.ServicesApiPromiseClient.prototype.stopService =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.ServicesApi/StopService',
      request,
      metadata || {},
      methodDescriptor_ServicesApi_StopService);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.GetServiceMetadataRequest,
 *   !proto.smartcore.bos.ServiceMetadata>}
 */
const methodDescriptor_ServicesApi_GetServiceMetadata = new grpc.web.MethodDescriptor(
  '/smartcore.bos.ServicesApi/GetServiceMetadata',
  grpc.web.MethodType.UNARY,
  proto.smartcore.bos.GetServiceMetadataRequest,
  proto.smartcore.bos.ServiceMetadata,
  /**
   * @param {!proto.smartcore.bos.GetServiceMetadataRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.ServiceMetadata.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.GetServiceMetadataRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.smartcore.bos.ServiceMetadata)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.ServiceMetadata>|undefined}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.ServicesApiClient.prototype.getServiceMetadata =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/smartcore.bos.ServicesApi/GetServiceMetadata',
      request,
      metadata || {},
      methodDescriptor_ServicesApi_GetServiceMetadata,
      callback);
};


/**
 * @param {!proto.smartcore.bos.GetServiceMetadataRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.smartcore.bos.ServiceMetadata>}
 *     Promise that resolves to the response
 */
proto.smartcore.bos.ServicesApiPromiseClient.prototype.getServiceMetadata =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/smartcore.bos.ServicesApi/GetServiceMetadata',
      request,
      metadata || {},
      methodDescriptor_ServicesApi_GetServiceMetadata);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.smartcore.bos.PullServiceMetadataRequest,
 *   !proto.smartcore.bos.PullServiceMetadataResponse>}
 */
const methodDescriptor_ServicesApi_PullServiceMetadata = new grpc.web.MethodDescriptor(
  '/smartcore.bos.ServicesApi/PullServiceMetadata',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.smartcore.bos.PullServiceMetadataRequest,
  proto.smartcore.bos.PullServiceMetadataResponse,
  /**
   * @param {!proto.smartcore.bos.PullServiceMetadataRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.smartcore.bos.PullServiceMetadataResponse.deserializeBinary
);


/**
 * @param {!proto.smartcore.bos.PullServiceMetadataRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.PullServiceMetadataResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.ServicesApiClient.prototype.pullServiceMetadata =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.bos.ServicesApi/PullServiceMetadata',
      request,
      metadata || {},
      methodDescriptor_ServicesApi_PullServiceMetadata);
};


/**
 * @param {!proto.smartcore.bos.PullServiceMetadataRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.smartcore.bos.PullServiceMetadataResponse>}
 *     The XHR Node Readable Stream
 */
proto.smartcore.bos.ServicesApiPromiseClient.prototype.pullServiceMetadata =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/smartcore.bos.ServicesApi/PullServiceMetadata',
      request,
      metadata || {},
      methodDescriptor_ServicesApi_PullServiceMetadata);
};


module.exports = proto.smartcore.bos;

