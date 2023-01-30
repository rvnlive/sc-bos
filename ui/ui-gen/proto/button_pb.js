// source: proto/button.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global =
    (typeof globalThis !== 'undefined' && globalThis) ||
    (typeof window !== 'undefined' && window) ||
    (typeof global !== 'undefined' && global) ||
    (typeof self !== 'undefined' && self) ||
    (function () { return this; }).call(null) ||
    Function('return this')();

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
goog.exportSymbol('proto.smartcore.bos.ButtonEvent', null, global);
goog.exportSymbol('proto.smartcore.bos.ButtonState', null, global);
goog.exportSymbol('proto.smartcore.bos.ButtonState.PressState', null, global);
goog.exportSymbol('proto.smartcore.bos.GetButtonStateRequest', null, global);
goog.exportSymbol('proto.smartcore.bos.PullButtonEventsRequest', null, global);
goog.exportSymbol('proto.smartcore.bos.PullButtonEventsResponse', null, global);
goog.exportSymbol('proto.smartcore.bos.PullButtonEventsResponse.Change', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.smartcore.bos.ButtonState = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.smartcore.bos.ButtonState, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.smartcore.bos.ButtonState.displayName = 'proto.smartcore.bos.ButtonState';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.smartcore.bos.GetButtonStateRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.smartcore.bos.GetButtonStateRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.smartcore.bos.GetButtonStateRequest.displayName = 'proto.smartcore.bos.GetButtonStateRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.smartcore.bos.PullButtonEventsRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.smartcore.bos.PullButtonEventsRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.smartcore.bos.PullButtonEventsRequest.displayName = 'proto.smartcore.bos.PullButtonEventsRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.smartcore.bos.PullButtonEventsResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.smartcore.bos.PullButtonEventsResponse.repeatedFields_, null);
};
goog.inherits(proto.smartcore.bos.PullButtonEventsResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.smartcore.bos.PullButtonEventsResponse.displayName = 'proto.smartcore.bos.PullButtonEventsResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.smartcore.bos.PullButtonEventsResponse.Change = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.smartcore.bos.PullButtonEventsResponse.Change, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.smartcore.bos.PullButtonEventsResponse.Change.displayName = 'proto.smartcore.bos.PullButtonEventsResponse.Change';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.smartcore.bos.ButtonState.prototype.toObject = function(opt_includeInstance) {
  return proto.smartcore.bos.ButtonState.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.smartcore.bos.ButtonState} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.smartcore.bos.ButtonState.toObject = function(includeInstance, msg) {
  var f, obj = {
    pressState: jspb.Message.getFieldWithDefault(msg, 1, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.smartcore.bos.ButtonState}
 */
proto.smartcore.bos.ButtonState.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.smartcore.bos.ButtonState;
  return proto.smartcore.bos.ButtonState.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.smartcore.bos.ButtonState} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.smartcore.bos.ButtonState}
 */
proto.smartcore.bos.ButtonState.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.smartcore.bos.ButtonState.PressState} */ (reader.readEnum());
      msg.setPressState(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.smartcore.bos.ButtonState.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.smartcore.bos.ButtonState.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.smartcore.bos.ButtonState} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.smartcore.bos.ButtonState.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPressState();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.smartcore.bos.ButtonState.PressState = {
  BUTTON_STATE_UNSPECIFIED: 0,
  PRESSED: 1,
  RELEASED: 2
};

/**
 * optional PressState press_state = 1;
 * @return {!proto.smartcore.bos.ButtonState.PressState}
 */
proto.smartcore.bos.ButtonState.prototype.getPressState = function() {
  return /** @type {!proto.smartcore.bos.ButtonState.PressState} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.smartcore.bos.ButtonState.PressState} value
 * @return {!proto.smartcore.bos.ButtonState} returns this
 */
proto.smartcore.bos.ButtonState.prototype.setPressState = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.smartcore.bos.GetButtonStateRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.smartcore.bos.GetButtonStateRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.smartcore.bos.GetButtonStateRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.smartcore.bos.GetButtonStateRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.smartcore.bos.GetButtonStateRequest}
 */
proto.smartcore.bos.GetButtonStateRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.smartcore.bos.GetButtonStateRequest;
  return proto.smartcore.bos.GetButtonStateRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.smartcore.bos.GetButtonStateRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.smartcore.bos.GetButtonStateRequest}
 */
proto.smartcore.bos.GetButtonStateRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.smartcore.bos.GetButtonStateRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.smartcore.bos.GetButtonStateRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.smartcore.bos.GetButtonStateRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.smartcore.bos.GetButtonStateRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string name = 1;
 * @return {string}
 */
proto.smartcore.bos.GetButtonStateRequest.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.smartcore.bos.GetButtonStateRequest} returns this
 */
proto.smartcore.bos.GetButtonStateRequest.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.smartcore.bos.PullButtonEventsRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.smartcore.bos.PullButtonEventsRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.smartcore.bos.PullButtonEventsRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.smartcore.bos.PullButtonEventsRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.smartcore.bos.PullButtonEventsRequest}
 */
proto.smartcore.bos.PullButtonEventsRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.smartcore.bos.PullButtonEventsRequest;
  return proto.smartcore.bos.PullButtonEventsRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.smartcore.bos.PullButtonEventsRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.smartcore.bos.PullButtonEventsRequest}
 */
proto.smartcore.bos.PullButtonEventsRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.smartcore.bos.PullButtonEventsRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.smartcore.bos.PullButtonEventsRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.smartcore.bos.PullButtonEventsRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.smartcore.bos.PullButtonEventsRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string name = 1;
 * @return {string}
 */
proto.smartcore.bos.PullButtonEventsRequest.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.smartcore.bos.PullButtonEventsRequest} returns this
 */
proto.smartcore.bos.PullButtonEventsRequest.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.smartcore.bos.PullButtonEventsResponse.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.smartcore.bos.PullButtonEventsResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.smartcore.bos.PullButtonEventsResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.smartcore.bos.PullButtonEventsResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.smartcore.bos.PullButtonEventsResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    changesList: jspb.Message.toObjectList(msg.getChangesList(),
    proto.smartcore.bos.PullButtonEventsResponse.Change.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.smartcore.bos.PullButtonEventsResponse}
 */
proto.smartcore.bos.PullButtonEventsResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.smartcore.bos.PullButtonEventsResponse;
  return proto.smartcore.bos.PullButtonEventsResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.smartcore.bos.PullButtonEventsResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.smartcore.bos.PullButtonEventsResponse}
 */
proto.smartcore.bos.PullButtonEventsResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.smartcore.bos.PullButtonEventsResponse.Change;
      reader.readMessage(value,proto.smartcore.bos.PullButtonEventsResponse.Change.deserializeBinaryFromReader);
      msg.addChanges(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.smartcore.bos.PullButtonEventsResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.smartcore.bos.PullButtonEventsResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.smartcore.bos.PullButtonEventsResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.smartcore.bos.PullButtonEventsResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getChangesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.smartcore.bos.PullButtonEventsResponse.Change.serializeBinaryToWriter
    );
  }
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.smartcore.bos.PullButtonEventsResponse.Change.prototype.toObject = function(opt_includeInstance) {
  return proto.smartcore.bos.PullButtonEventsResponse.Change.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.smartcore.bos.PullButtonEventsResponse.Change} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.smartcore.bos.PullButtonEventsResponse.Change.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, ""),
    changeTime: (f = msg.getChangeTime()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    buttonEvent: jspb.Message.getFieldWithDefault(msg, 3, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.smartcore.bos.PullButtonEventsResponse.Change}
 */
proto.smartcore.bos.PullButtonEventsResponse.Change.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.smartcore.bos.PullButtonEventsResponse.Change;
  return proto.smartcore.bos.PullButtonEventsResponse.Change.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.smartcore.bos.PullButtonEventsResponse.Change} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.smartcore.bos.PullButtonEventsResponse.Change}
 */
proto.smartcore.bos.PullButtonEventsResponse.Change.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 2:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setChangeTime(value);
      break;
    case 3:
      var value = /** @type {!proto.smartcore.bos.ButtonEvent} */ (reader.readEnum());
      msg.setButtonEvent(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.smartcore.bos.PullButtonEventsResponse.Change.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.smartcore.bos.PullButtonEventsResponse.Change.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.smartcore.bos.PullButtonEventsResponse.Change} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.smartcore.bos.PullButtonEventsResponse.Change.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getChangeTime();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getButtonEvent();
  if (f !== 0.0) {
    writer.writeEnum(
      3,
      f
    );
  }
};


/**
 * optional string name = 1;
 * @return {string}
 */
proto.smartcore.bos.PullButtonEventsResponse.Change.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.smartcore.bos.PullButtonEventsResponse.Change} returns this
 */
proto.smartcore.bos.PullButtonEventsResponse.Change.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional google.protobuf.Timestamp change_time = 2;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.smartcore.bos.PullButtonEventsResponse.Change.prototype.getChangeTime = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 2));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.smartcore.bos.PullButtonEventsResponse.Change} returns this
*/
proto.smartcore.bos.PullButtonEventsResponse.Change.prototype.setChangeTime = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.smartcore.bos.PullButtonEventsResponse.Change} returns this
 */
proto.smartcore.bos.PullButtonEventsResponse.Change.prototype.clearChangeTime = function() {
  return this.setChangeTime(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.smartcore.bos.PullButtonEventsResponse.Change.prototype.hasChangeTime = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional ButtonEvent button_event = 3;
 * @return {!proto.smartcore.bos.ButtonEvent}
 */
proto.smartcore.bos.PullButtonEventsResponse.Change.prototype.getButtonEvent = function() {
  return /** @type {!proto.smartcore.bos.ButtonEvent} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {!proto.smartcore.bos.ButtonEvent} value
 * @return {!proto.smartcore.bos.PullButtonEventsResponse.Change} returns this
 */
proto.smartcore.bos.PullButtonEventsResponse.Change.prototype.setButtonEvent = function(value) {
  return jspb.Message.setProto3EnumField(this, 3, value);
};


/**
 * repeated Change changes = 1;
 * @return {!Array<!proto.smartcore.bos.PullButtonEventsResponse.Change>}
 */
proto.smartcore.bos.PullButtonEventsResponse.prototype.getChangesList = function() {
  return /** @type{!Array<!proto.smartcore.bos.PullButtonEventsResponse.Change>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.smartcore.bos.PullButtonEventsResponse.Change, 1));
};


/**
 * @param {!Array<!proto.smartcore.bos.PullButtonEventsResponse.Change>} value
 * @return {!proto.smartcore.bos.PullButtonEventsResponse} returns this
*/
proto.smartcore.bos.PullButtonEventsResponse.prototype.setChangesList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.smartcore.bos.PullButtonEventsResponse.Change=} opt_value
 * @param {number=} opt_index
 * @return {!proto.smartcore.bos.PullButtonEventsResponse.Change}
 */
proto.smartcore.bos.PullButtonEventsResponse.prototype.addChanges = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.smartcore.bos.PullButtonEventsResponse.Change, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.smartcore.bos.PullButtonEventsResponse} returns this
 */
proto.smartcore.bos.PullButtonEventsResponse.prototype.clearChangesList = function() {
  return this.setChangesList([]);
};


/**
 * @enum {number}
 */
proto.smartcore.bos.ButtonEvent = {
  BUTTON_EVENT_UNSPECIFIED: 0,
  PRESS: 1,
  RELEASE: 2,
  SHORT_PRESS: 3,
  DOUBLE_PRESS: 4,
  LONG_PRESS_START: 5,
  LONG_PRESS_REPEAT: 6,
  LONG_PRESS_END: 7
};

goog.object.extend(exports, proto.smartcore.bos);
