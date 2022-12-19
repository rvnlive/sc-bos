import * as jspb from 'google-protobuf'

import * as google_protobuf_field_mask_pb from 'google-protobuf/google/protobuf/field_mask_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as traits_metadata_pb from '@smart-core-os/sc-api-grpc-web/traits/metadata_pb';
import * as types_change_pb from '@smart-core-os/sc-api-grpc-web/types/change_pb';


export class Device extends jspb.Message {
  getName(): string;
  setName(value: string): Device;

  getMetadata(): traits_metadata_pb.Metadata | undefined;
  setMetadata(value?: traits_metadata_pb.Metadata): Device;
  hasMetadata(): boolean;
  clearMetadata(): Device;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Device.AsObject;
  static toObject(includeInstance: boolean, msg: Device): Device.AsObject;
  static serializeBinaryToWriter(message: Device, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Device;
  static deserializeBinaryFromReader(message: Device, reader: jspb.BinaryReader): Device;
}

export namespace Device {
  export type AsObject = {
    name: string,
    metadata?: traits_metadata_pb.Metadata.AsObject,
  }

  export class Query extends jspb.Message {
    getConditionsList(): Array<Device.Query.Condition>;
    setConditionsList(value: Array<Device.Query.Condition>): Query;
    clearConditionsList(): Query;
    addConditions(value?: Device.Query.Condition, index?: number): Device.Query.Condition;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Query.AsObject;
    static toObject(includeInstance: boolean, msg: Query): Query.AsObject;
    static serializeBinaryToWriter(message: Query, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Query;
    static deserializeBinaryFromReader(message: Query, reader: jspb.BinaryReader): Query;
  }

  export namespace Query {
    export type AsObject = {
      conditionsList: Array<Device.Query.Condition.AsObject>,
    }

    export class Condition extends jspb.Message {
      getField(): string;
      setField(value: string): Condition;

      getStringEqual(): string;
      setStringEqual(value: string): Condition;

      getValueCase(): Condition.ValueCase;

      serializeBinary(): Uint8Array;
      toObject(includeInstance?: boolean): Condition.AsObject;
      static toObject(includeInstance: boolean, msg: Condition): Condition.AsObject;
      static serializeBinaryToWriter(message: Condition, writer: jspb.BinaryWriter): void;
      static deserializeBinary(bytes: Uint8Array): Condition;
      static deserializeBinaryFromReader(message: Condition, reader: jspb.BinaryReader): Condition;
    }

    export namespace Condition {
      export type AsObject = {
        field: string,
        stringEqual: string,
      }

      export enum ValueCase { 
        VALUE_NOT_SET = 0,
        STRING_EQUAL = 2,
      }
    }

  }

}

export class ListDevicesRequest extends jspb.Message {
  getReadMask(): google_protobuf_field_mask_pb.FieldMask | undefined;
  setReadMask(value?: google_protobuf_field_mask_pb.FieldMask): ListDevicesRequest;
  hasReadMask(): boolean;
  clearReadMask(): ListDevicesRequest;

  getPageSize(): number;
  setPageSize(value: number): ListDevicesRequest;

  getPageToken(): string;
  setPageToken(value: string): ListDevicesRequest;

  getQuery(): Device.Query | undefined;
  setQuery(value?: Device.Query): ListDevicesRequest;
  hasQuery(): boolean;
  clearQuery(): ListDevicesRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListDevicesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListDevicesRequest): ListDevicesRequest.AsObject;
  static serializeBinaryToWriter(message: ListDevicesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListDevicesRequest;
  static deserializeBinaryFromReader(message: ListDevicesRequest, reader: jspb.BinaryReader): ListDevicesRequest;
}

export namespace ListDevicesRequest {
  export type AsObject = {
    readMask?: google_protobuf_field_mask_pb.FieldMask.AsObject,
    pageSize: number,
    pageToken: string,
    query?: Device.Query.AsObject,
  }
}

export class ListDevicesResponse extends jspb.Message {
  getDevicesList(): Array<Device>;
  setDevicesList(value: Array<Device>): ListDevicesResponse;
  clearDevicesList(): ListDevicesResponse;
  addDevices(value?: Device, index?: number): Device;

  getNextPageToken(): string;
  setNextPageToken(value: string): ListDevicesResponse;

  getTotalSize(): number;
  setTotalSize(value: number): ListDevicesResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListDevicesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListDevicesResponse): ListDevicesResponse.AsObject;
  static serializeBinaryToWriter(message: ListDevicesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListDevicesResponse;
  static deserializeBinaryFromReader(message: ListDevicesResponse, reader: jspb.BinaryReader): ListDevicesResponse;
}

export namespace ListDevicesResponse {
  export type AsObject = {
    devicesList: Array<Device.AsObject>,
    nextPageToken: string,
    totalSize: number,
  }
}

export class PullDevicesRequest extends jspb.Message {
  getReadMask(): google_protobuf_field_mask_pb.FieldMask | undefined;
  setReadMask(value?: google_protobuf_field_mask_pb.FieldMask): PullDevicesRequest;
  hasReadMask(): boolean;
  clearReadMask(): PullDevicesRequest;

  getUpdatesOnly(): boolean;
  setUpdatesOnly(value: boolean): PullDevicesRequest;

  getQuery(): Device.Query | undefined;
  setQuery(value?: Device.Query): PullDevicesRequest;
  hasQuery(): boolean;
  clearQuery(): PullDevicesRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PullDevicesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PullDevicesRequest): PullDevicesRequest.AsObject;
  static serializeBinaryToWriter(message: PullDevicesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PullDevicesRequest;
  static deserializeBinaryFromReader(message: PullDevicesRequest, reader: jspb.BinaryReader): PullDevicesRequest;
}

export namespace PullDevicesRequest {
  export type AsObject = {
    readMask?: google_protobuf_field_mask_pb.FieldMask.AsObject,
    updatesOnly: boolean,
    query?: Device.Query.AsObject,
  }
}

export class PullDevicesResponse extends jspb.Message {
  getChangesList(): Array<PullDevicesResponse.Change>;
  setChangesList(value: Array<PullDevicesResponse.Change>): PullDevicesResponse;
  clearChangesList(): PullDevicesResponse;
  addChanges(value?: PullDevicesResponse.Change, index?: number): PullDevicesResponse.Change;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PullDevicesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PullDevicesResponse): PullDevicesResponse.AsObject;
  static serializeBinaryToWriter(message: PullDevicesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PullDevicesResponse;
  static deserializeBinaryFromReader(message: PullDevicesResponse, reader: jspb.BinaryReader): PullDevicesResponse;
}

export namespace PullDevicesResponse {
  export type AsObject = {
    changesList: Array<PullDevicesResponse.Change.AsObject>,
  }

  export class Change extends jspb.Message {
    getName(): string;
    setName(value: string): Change;

    getType(): types_change_pb.ChangeType;
    setType(value: types_change_pb.ChangeType): Change;

    getNewValue(): Device | undefined;
    setNewValue(value?: Device): Change;
    hasNewValue(): boolean;
    clearNewValue(): Change;

    getOldValue(): Device | undefined;
    setOldValue(value?: Device): Change;
    hasOldValue(): boolean;
    clearOldValue(): Change;

    getChangeTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setChangeTime(value?: google_protobuf_timestamp_pb.Timestamp): Change;
    hasChangeTime(): boolean;
    clearChangeTime(): Change;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Change.AsObject;
    static toObject(includeInstance: boolean, msg: Change): Change.AsObject;
    static serializeBinaryToWriter(message: Change, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Change;
    static deserializeBinaryFromReader(message: Change, reader: jspb.BinaryReader): Change;
  }

  export namespace Change {
    export type AsObject = {
      name: string,
      type: types_change_pb.ChangeType,
      newValue?: Device.AsObject,
      oldValue?: Device.AsObject,
      changeTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    }
  }

}
