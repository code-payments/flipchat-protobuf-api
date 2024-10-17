// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file push/v1/push_service.proto (package flipchat.push.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { AppInstallId, Auth, UserId } from "../../common/v1/model_pb";

/**
 * @generated from enum flipchat.push.v1.TokenType
 */
export enum TokenType {
  /**
   * @generated from enum value: UNKNOWN = 0;
   */
  UNKNOWN = 0,

  /**
   * FCM registration token for an Android device
   *
   * @generated from enum value: FCM_ANDROID = 1;
   */
  FCM_ANDROID = 1,

  /**
   * FCM registration token or an iOS device
   *
   * @generated from enum value: FCM_APNS = 2;
   */
  FCM_APNS = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(TokenType)
proto3.util.setEnumType(TokenType, "flipchat.push.v1.TokenType", [
  { no: 0, name: "UNKNOWN" },
  { no: 1, name: "FCM_ANDROID" },
  { no: 2, name: "FCM_APNS" },
]);

/**
 * @generated from message flipchat.push.v1.AddTokenRequest
 */
export class AddTokenRequest extends Message<AddTokenRequest> {
  /**
   * @generated from field: flipchat.common.v1.UserId user_id = 1;
   */
  userId?: UserId;

  /**
   * @generated from field: string push_token = 2;
   */
  pushToken = "";

  /**
   * @generated from field: flipchat.push.v1.TokenType token_type = 3;
   */
  tokenType = TokenType.UNKNOWN;

  /**
   * @generated from field: flipchat.common.v1.AppInstallId app_install = 4;
   */
  appInstall?: AppInstallId;

  /**
   * @generated from field: flipchat.common.v1.Auth auth = 5;
   */
  auth?: Auth;

  constructor(data?: PartialMessage<AddTokenRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.push.v1.AddTokenRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_id", kind: "message", T: UserId },
    { no: 2, name: "push_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "token_type", kind: "enum", T: proto3.getEnumType(TokenType) },
    { no: 4, name: "app_install", kind: "message", T: AppInstallId },
    { no: 5, name: "auth", kind: "message", T: Auth },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AddTokenRequest {
    return new AddTokenRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AddTokenRequest {
    return new AddTokenRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AddTokenRequest {
    return new AddTokenRequest().fromJsonString(jsonString, options);
  }

  static equals(a: AddTokenRequest | PlainMessage<AddTokenRequest> | undefined, b: AddTokenRequest | PlainMessage<AddTokenRequest> | undefined): boolean {
    return proto3.util.equals(AddTokenRequest, a, b);
  }
}

/**
 * @generated from message flipchat.push.v1.AddTokenResponse
 */
export class AddTokenResponse extends Message<AddTokenResponse> {
  /**
   * @generated from field: flipchat.push.v1.AddTokenResponse.Result result = 1;
   */
  result = AddTokenResponse_Result.OK;

  constructor(data?: PartialMessage<AddTokenResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.push.v1.AddTokenResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "result", kind: "enum", T: proto3.getEnumType(AddTokenResponse_Result) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AddTokenResponse {
    return new AddTokenResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AddTokenResponse {
    return new AddTokenResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AddTokenResponse {
    return new AddTokenResponse().fromJsonString(jsonString, options);
  }

  static equals(a: AddTokenResponse | PlainMessage<AddTokenResponse> | undefined, b: AddTokenResponse | PlainMessage<AddTokenResponse> | undefined): boolean {
    return proto3.util.equals(AddTokenResponse, a, b);
  }
}

/**
 * @generated from enum flipchat.push.v1.AddTokenResponse.Result
 */
export enum AddTokenResponse_Result {
  /**
   * @generated from enum value: OK = 0;
   */
  OK = 0,

  /**
   * @generated from enum value: INVALID_PUSH_TOKEN = 1;
   */
  INVALID_PUSH_TOKEN = 1,
}
// Retrieve enum metadata with: proto3.getEnumType(AddTokenResponse_Result)
proto3.util.setEnumType(AddTokenResponse_Result, "flipchat.push.v1.AddTokenResponse.Result", [
  { no: 0, name: "OK" },
  { no: 1, name: "INVALID_PUSH_TOKEN" },
]);

/**
 * @generated from message flipchat.push.v1.DeleteTokenRequest
 */
export class DeleteTokenRequest extends Message<DeleteTokenRequest> {
  /**
   * @generated from field: flipchat.common.v1.UserId user_id = 1;
   */
  userId?: UserId;

  /**
   * @generated from field: string push_token = 2;
   */
  pushToken = "";

  /**
   * @generated from field: flipchat.push.v1.TokenType token_type = 3;
   */
  tokenType = TokenType.UNKNOWN;

  /**
   * @generated from field: flipchat.common.v1.AppInstallId app_install = 4;
   */
  appInstall?: AppInstallId;

  /**
   * @generated from field: flipchat.common.v1.Auth auth = 5;
   */
  auth?: Auth;

  constructor(data?: PartialMessage<DeleteTokenRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.push.v1.DeleteTokenRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_id", kind: "message", T: UserId },
    { no: 2, name: "push_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "token_type", kind: "enum", T: proto3.getEnumType(TokenType) },
    { no: 4, name: "app_install", kind: "message", T: AppInstallId },
    { no: 5, name: "auth", kind: "message", T: Auth },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteTokenRequest {
    return new DeleteTokenRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteTokenRequest {
    return new DeleteTokenRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteTokenRequest {
    return new DeleteTokenRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteTokenRequest | PlainMessage<DeleteTokenRequest> | undefined, b: DeleteTokenRequest | PlainMessage<DeleteTokenRequest> | undefined): boolean {
    return proto3.util.equals(DeleteTokenRequest, a, b);
  }
}

/**
 * @generated from message flipchat.push.v1.DeleteTokenResponse
 */
export class DeleteTokenResponse extends Message<DeleteTokenResponse> {
  constructor(data?: PartialMessage<DeleteTokenResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.push.v1.DeleteTokenResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteTokenResponse {
    return new DeleteTokenResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteTokenResponse {
    return new DeleteTokenResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteTokenResponse {
    return new DeleteTokenResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteTokenResponse | PlainMessage<DeleteTokenResponse> | undefined, b: DeleteTokenResponse | PlainMessage<DeleteTokenResponse> | undefined): boolean {
    return proto3.util.equals(DeleteTokenResponse, a, b);
  }
}

