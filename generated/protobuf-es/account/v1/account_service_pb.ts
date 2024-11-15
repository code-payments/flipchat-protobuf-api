// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file account/v1/account_service.proto (package flipchat.account.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";
import { Auth, PaymentAmount, PublicKey, Signature, UserId } from "../../common/v1/flipchat_pb";

/**
 * @generated from message flipchat.account.v1.RegisterRequest
 */
export class RegisterRequest extends Message<RegisterRequest> {
  /**
   * PublicKey the public key that is authorized to perform actions on the
   * registered users behalf.
   *
   * @generated from field: flipchat.common.v1.PublicKey public_key = 1;
   */
  publicKey?: PublicKey;

  /**
   * Signature of this message (without the signature), using the provided keypaid.
   *
   * @generated from field: flipchat.common.v1.Signature signature = 2;
   */
  signature?: Signature;

  /**
   * DisplayName of the new user.
   *
   * If the user already exists, this is not used.
   *
   * @generated from field: string display_name = 3;
   */
  displayName = "";

  constructor(data?: PartialMessage<RegisterRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.account.v1.RegisterRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "public_key", kind: "message", T: PublicKey },
    { no: 2, name: "signature", kind: "message", T: Signature },
    { no: 3, name: "display_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RegisterRequest {
    return new RegisterRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RegisterRequest {
    return new RegisterRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RegisterRequest {
    return new RegisterRequest().fromJsonString(jsonString, options);
  }

  static equals(a: RegisterRequest | PlainMessage<RegisterRequest> | undefined, b: RegisterRequest | PlainMessage<RegisterRequest> | undefined): boolean {
    return proto3.util.equals(RegisterRequest, a, b);
  }
}

/**
 * @generated from message flipchat.account.v1.RegisterResponse
 */
export class RegisterResponse extends Message<RegisterResponse> {
  /**
   * @generated from field: flipchat.account.v1.RegisterResponse.Result result = 1;
   */
  result = RegisterResponse_Result.OK;

  /**
   * Error reason contains the reason for the error, if the
   * result > 1. This allows for server to impose moderation restrictions
   * on user provided fields.
   *
   * @generated from field: string error_reason = 2;
   */
  errorReason = "";

  /**
   * The UserId associated with the account.
   *
   * @generated from field: flipchat.common.v1.UserId user_id = 3;
   */
  userId?: UserId;

  constructor(data?: PartialMessage<RegisterResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.account.v1.RegisterResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "result", kind: "enum", T: proto3.getEnumType(RegisterResponse_Result) },
    { no: 2, name: "error_reason", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "user_id", kind: "message", T: UserId },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RegisterResponse {
    return new RegisterResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RegisterResponse {
    return new RegisterResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RegisterResponse {
    return new RegisterResponse().fromJsonString(jsonString, options);
  }

  static equals(a: RegisterResponse | PlainMessage<RegisterResponse> | undefined, b: RegisterResponse | PlainMessage<RegisterResponse> | undefined): boolean {
    return proto3.util.equals(RegisterResponse, a, b);
  }
}

/**
 * @generated from enum flipchat.account.v1.RegisterResponse.Result
 */
export enum RegisterResponse_Result {
  /**
   * @generated from enum value: OK = 0;
   */
  OK = 0,

  /**
   * @generated from enum value: INVALID_SIGNATURE = 1;
   */
  INVALID_SIGNATURE = 1,

  /**
   * @generated from enum value: INVALID_DISPLAY_NAME = 2;
   */
  INVALID_DISPLAY_NAME = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(RegisterResponse_Result)
proto3.util.setEnumType(RegisterResponse_Result, "flipchat.account.v1.RegisterResponse.Result", [
  { no: 0, name: "OK" },
  { no: 1, name: "INVALID_SIGNATURE" },
  { no: 2, name: "INVALID_DISPLAY_NAME" },
]);

/**
 * @generated from message flipchat.account.v1.LoginRequest
 */
export class LoginRequest extends Message<LoginRequest> {
  /**
   * Timestamp is the timestamp the request was generated.
   *
   * The server may reject the request if the timestamp is too far off
   * the current (server) time. This is to prevent replay attacks.
   *
   * @generated from field: google.protobuf.Timestamp timestamp = 1;
   */
  timestamp?: Timestamp;

  /**
   * @generated from field: flipchat.common.v1.Auth auth = 2;
   */
  auth?: Auth;

  constructor(data?: PartialMessage<LoginRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.account.v1.LoginRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "timestamp", kind: "message", T: Timestamp },
    { no: 2, name: "auth", kind: "message", T: Auth },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LoginRequest {
    return new LoginRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LoginRequest {
    return new LoginRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LoginRequest {
    return new LoginRequest().fromJsonString(jsonString, options);
  }

  static equals(a: LoginRequest | PlainMessage<LoginRequest> | undefined, b: LoginRequest | PlainMessage<LoginRequest> | undefined): boolean {
    return proto3.util.equals(LoginRequest, a, b);
  }
}

/**
 * @generated from message flipchat.account.v1.LoginResponse
 */
export class LoginResponse extends Message<LoginResponse> {
  /**
   * @generated from field: flipchat.account.v1.LoginResponse.Result result = 1;
   */
  result = LoginResponse_Result.OK;

  /**
   * UserId is the user associated with the PubKey/Auth.
   *
   * @generated from field: flipchat.common.v1.UserId user_id = 2;
   */
  userId?: UserId;

  constructor(data?: PartialMessage<LoginResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.account.v1.LoginResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "result", kind: "enum", T: proto3.getEnumType(LoginResponse_Result) },
    { no: 2, name: "user_id", kind: "message", T: UserId },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LoginResponse {
    return new LoginResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LoginResponse {
    return new LoginResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LoginResponse {
    return new LoginResponse().fromJsonString(jsonString, options);
  }

  static equals(a: LoginResponse | PlainMessage<LoginResponse> | undefined, b: LoginResponse | PlainMessage<LoginResponse> | undefined): boolean {
    return proto3.util.equals(LoginResponse, a, b);
  }
}

/**
 * @generated from enum flipchat.account.v1.LoginResponse.Result
 */
export enum LoginResponse_Result {
  /**
   * @generated from enum value: OK = 0;
   */
  OK = 0,

  /**
   * @generated from enum value: INVALID_TIMESTAMP = 1;
   */
  INVALID_TIMESTAMP = 1,

  /**
   * @generated from enum value: DENIED = 2;
   */
  DENIED = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(LoginResponse_Result)
proto3.util.setEnumType(LoginResponse_Result, "flipchat.account.v1.LoginResponse.Result", [
  { no: 0, name: "OK" },
  { no: 1, name: "INVALID_TIMESTAMP" },
  { no: 2, name: "DENIED" },
]);

/**
 * @generated from message flipchat.account.v1.AuthorizePublicKeyRequest
 */
export class AuthorizePublicKeyRequest extends Message<AuthorizePublicKeyRequest> {
  /**
   * UserId to bound the new public key to.
   *
   * @generated from field: flipchat.common.v1.UserId user_id = 1;
   */
  userId?: UserId;

  /**
   * PublicKey of the account to be added.
   *
   * @generated from field: flipchat.common.v1.PublicKey public_key = 2;
   */
  publicKey?: PublicKey;

  /**
   * Signature of this message, not including auth or signature, using the
   * new public key.
   *
   * @generated from field: flipchat.common.v1.Signature signature = 3;
   */
  signature?: Signature;

  /**
   * @generated from field: flipchat.common.v1.Auth auth = 4;
   */
  auth?: Auth;

  constructor(data?: PartialMessage<AuthorizePublicKeyRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.account.v1.AuthorizePublicKeyRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_id", kind: "message", T: UserId },
    { no: 2, name: "public_key", kind: "message", T: PublicKey },
    { no: 3, name: "signature", kind: "message", T: Signature },
    { no: 4, name: "auth", kind: "message", T: Auth },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AuthorizePublicKeyRequest {
    return new AuthorizePublicKeyRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AuthorizePublicKeyRequest {
    return new AuthorizePublicKeyRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AuthorizePublicKeyRequest {
    return new AuthorizePublicKeyRequest().fromJsonString(jsonString, options);
  }

  static equals(a: AuthorizePublicKeyRequest | PlainMessage<AuthorizePublicKeyRequest> | undefined, b: AuthorizePublicKeyRequest | PlainMessage<AuthorizePublicKeyRequest> | undefined): boolean {
    return proto3.util.equals(AuthorizePublicKeyRequest, a, b);
  }
}

/**
 * @generated from message flipchat.account.v1.AuthorizePublicKeyResponse
 */
export class AuthorizePublicKeyResponse extends Message<AuthorizePublicKeyResponse> {
  /**
   * @generated from field: flipchat.account.v1.AuthorizePublicKeyResponse.Result result = 1;
   */
  result = AuthorizePublicKeyResponse_Result.OK;

  constructor(data?: PartialMessage<AuthorizePublicKeyResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.account.v1.AuthorizePublicKeyResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "result", kind: "enum", T: proto3.getEnumType(AuthorizePublicKeyResponse_Result) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AuthorizePublicKeyResponse {
    return new AuthorizePublicKeyResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AuthorizePublicKeyResponse {
    return new AuthorizePublicKeyResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AuthorizePublicKeyResponse {
    return new AuthorizePublicKeyResponse().fromJsonString(jsonString, options);
  }

  static equals(a: AuthorizePublicKeyResponse | PlainMessage<AuthorizePublicKeyResponse> | undefined, b: AuthorizePublicKeyResponse | PlainMessage<AuthorizePublicKeyResponse> | undefined): boolean {
    return proto3.util.equals(AuthorizePublicKeyResponse, a, b);
  }
}

/**
 * @generated from enum flipchat.account.v1.AuthorizePublicKeyResponse.Result
 */
export enum AuthorizePublicKeyResponse_Result {
  /**
   * @generated from enum value: OK = 0;
   */
  OK = 0,

  /**
   * @generated from enum value: DENIED = 1;
   */
  DENIED = 1,
}
// Retrieve enum metadata with: proto3.getEnumType(AuthorizePublicKeyResponse_Result)
proto3.util.setEnumType(AuthorizePublicKeyResponse_Result, "flipchat.account.v1.AuthorizePublicKeyResponse.Result", [
  { no: 0, name: "OK" },
  { no: 1, name: "DENIED" },
]);

/**
 * @generated from message flipchat.account.v1.RevokePublicKeyRequest
 */
export class RevokePublicKeyRequest extends Message<RevokePublicKeyRequest> {
  /**
   * UserId to remove the public key from.
   *
   * @generated from field: flipchat.common.v1.UserId user_id = 1;
   */
  userId?: UserId;

  /**
   * PublicKey to remove.
   *
   * @generated from field: flipchat.common.v1.PublicKey public_key = 2;
   */
  publicKey?: PublicKey;

  /**
   * @generated from field: flipchat.common.v1.Auth auth = 4;
   */
  auth?: Auth;

  constructor(data?: PartialMessage<RevokePublicKeyRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.account.v1.RevokePublicKeyRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_id", kind: "message", T: UserId },
    { no: 2, name: "public_key", kind: "message", T: PublicKey },
    { no: 4, name: "auth", kind: "message", T: Auth },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RevokePublicKeyRequest {
    return new RevokePublicKeyRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RevokePublicKeyRequest {
    return new RevokePublicKeyRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RevokePublicKeyRequest {
    return new RevokePublicKeyRequest().fromJsonString(jsonString, options);
  }

  static equals(a: RevokePublicKeyRequest | PlainMessage<RevokePublicKeyRequest> | undefined, b: RevokePublicKeyRequest | PlainMessage<RevokePublicKeyRequest> | undefined): boolean {
    return proto3.util.equals(RevokePublicKeyRequest, a, b);
  }
}

/**
 * @generated from message flipchat.account.v1.RevokePublicKeyResponse
 */
export class RevokePublicKeyResponse extends Message<RevokePublicKeyResponse> {
  /**
   * @generated from field: flipchat.account.v1.RevokePublicKeyResponse.Result result = 1;
   */
  result = RevokePublicKeyResponse_Result.OK;

  constructor(data?: PartialMessage<RevokePublicKeyResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.account.v1.RevokePublicKeyResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "result", kind: "enum", T: proto3.getEnumType(RevokePublicKeyResponse_Result) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RevokePublicKeyResponse {
    return new RevokePublicKeyResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RevokePublicKeyResponse {
    return new RevokePublicKeyResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RevokePublicKeyResponse {
    return new RevokePublicKeyResponse().fromJsonString(jsonString, options);
  }

  static equals(a: RevokePublicKeyResponse | PlainMessage<RevokePublicKeyResponse> | undefined, b: RevokePublicKeyResponse | PlainMessage<RevokePublicKeyResponse> | undefined): boolean {
    return proto3.util.equals(RevokePublicKeyResponse, a, b);
  }
}

/**
 * @generated from enum flipchat.account.v1.RevokePublicKeyResponse.Result
 */
export enum RevokePublicKeyResponse_Result {
  /**
   * @generated from enum value: OK = 0;
   */
  OK = 0,

  /**
   * @generated from enum value: DENIED = 1;
   */
  DENIED = 1,

  /**
   * @generated from enum value: LAST_PUB_KEY = 2;
   */
  LAST_PUB_KEY = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(RevokePublicKeyResponse_Result)
proto3.util.setEnumType(RevokePublicKeyResponse_Result, "flipchat.account.v1.RevokePublicKeyResponse.Result", [
  { no: 0, name: "OK" },
  { no: 1, name: "DENIED" },
  { no: 2, name: "LAST_PUB_KEY" },
]);

/**
 * @generated from message flipchat.account.v1.GetPaymentDestinationRequest
 */
export class GetPaymentDestinationRequest extends Message<GetPaymentDestinationRequest> {
  /**
   * UserId to get the payment destination from.
   *
   * @generated from field: flipchat.common.v1.UserId user_id = 1;
   */
  userId?: UserId;

  constructor(data?: PartialMessage<GetPaymentDestinationRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.account.v1.GetPaymentDestinationRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_id", kind: "message", T: UserId },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetPaymentDestinationRequest {
    return new GetPaymentDestinationRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetPaymentDestinationRequest {
    return new GetPaymentDestinationRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetPaymentDestinationRequest {
    return new GetPaymentDestinationRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetPaymentDestinationRequest | PlainMessage<GetPaymentDestinationRequest> | undefined, b: GetPaymentDestinationRequest | PlainMessage<GetPaymentDestinationRequest> | undefined): boolean {
    return proto3.util.equals(GetPaymentDestinationRequest, a, b);
  }
}

/**
 * @generated from message flipchat.account.v1.GetPaymentDestinationResponse
 */
export class GetPaymentDestinationResponse extends Message<GetPaymentDestinationResponse> {
  /**
   * @generated from field: flipchat.account.v1.GetPaymentDestinationResponse.Result result = 1;
   */
  result = GetPaymentDestinationResponse_Result.OK;

  /**
   * Payment destination for the UserId.
   *
   * @generated from field: flipchat.common.v1.PublicKey payment_destination = 2;
   */
  paymentDestination?: PublicKey;

  constructor(data?: PartialMessage<GetPaymentDestinationResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.account.v1.GetPaymentDestinationResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "result", kind: "enum", T: proto3.getEnumType(GetPaymentDestinationResponse_Result) },
    { no: 2, name: "payment_destination", kind: "message", T: PublicKey },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetPaymentDestinationResponse {
    return new GetPaymentDestinationResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetPaymentDestinationResponse {
    return new GetPaymentDestinationResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetPaymentDestinationResponse {
    return new GetPaymentDestinationResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetPaymentDestinationResponse | PlainMessage<GetPaymentDestinationResponse> | undefined, b: GetPaymentDestinationResponse | PlainMessage<GetPaymentDestinationResponse> | undefined): boolean {
    return proto3.util.equals(GetPaymentDestinationResponse, a, b);
  }
}

/**
 * @generated from enum flipchat.account.v1.GetPaymentDestinationResponse.Result
 */
export enum GetPaymentDestinationResponse_Result {
  /**
   * @generated from enum value: OK = 0;
   */
  OK = 0,

  /**
   * @generated from enum value: NOT_FOUND = 1;
   */
  NOT_FOUND = 1,
}
// Retrieve enum metadata with: proto3.getEnumType(GetPaymentDestinationResponse_Result)
proto3.util.setEnumType(GetPaymentDestinationResponse_Result, "flipchat.account.v1.GetPaymentDestinationResponse.Result", [
  { no: 0, name: "OK" },
  { no: 1, name: "NOT_FOUND" },
]);

/**
 * @generated from message flipchat.account.v1.GetFlagsRequest
 */
export class GetFlagsRequest extends Message<GetFlagsRequest> {
  /**
   * UserId to bound the new public key to.
   *
   * @generated from field: flipchat.common.v1.UserId user_id = 1;
   */
  userId?: UserId;

  /**
   * @generated from field: flipchat.common.v1.Auth auth = 2;
   */
  auth?: Auth;

  constructor(data?: PartialMessage<GetFlagsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.account.v1.GetFlagsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_id", kind: "message", T: UserId },
    { no: 2, name: "auth", kind: "message", T: Auth },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetFlagsRequest {
    return new GetFlagsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetFlagsRequest {
    return new GetFlagsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetFlagsRequest {
    return new GetFlagsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetFlagsRequest | PlainMessage<GetFlagsRequest> | undefined, b: GetFlagsRequest | PlainMessage<GetFlagsRequest> | undefined): boolean {
    return proto3.util.equals(GetFlagsRequest, a, b);
  }
}

/**
 * @generated from message flipchat.account.v1.GetFlagsResponse
 */
export class GetFlagsResponse extends Message<GetFlagsResponse> {
  /**
   * @generated from field: flipchat.account.v1.GetFlagsResponse.Result result = 1;
   */
  result = GetFlagsResponse_Result.OK;

  /**
   * @generated from field: flipchat.account.v1.Flags flags = 2;
   */
  flags?: Flags;

  constructor(data?: PartialMessage<GetFlagsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.account.v1.GetFlagsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "result", kind: "enum", T: proto3.getEnumType(GetFlagsResponse_Result) },
    { no: 2, name: "flags", kind: "message", T: Flags },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetFlagsResponse {
    return new GetFlagsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetFlagsResponse {
    return new GetFlagsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetFlagsResponse {
    return new GetFlagsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetFlagsResponse | PlainMessage<GetFlagsResponse> | undefined, b: GetFlagsResponse | PlainMessage<GetFlagsResponse> | undefined): boolean {
    return proto3.util.equals(GetFlagsResponse, a, b);
  }
}

/**
 * @generated from enum flipchat.account.v1.GetFlagsResponse.Result
 */
export enum GetFlagsResponse_Result {
  /**
   * @generated from enum value: OK = 0;
   */
  OK = 0,

  /**
   * @generated from enum value: DENIED = 1;
   */
  DENIED = 1,
}
// Retrieve enum metadata with: proto3.getEnumType(GetFlagsResponse_Result)
proto3.util.setEnumType(GetFlagsResponse_Result, "flipchat.account.v1.GetFlagsResponse.Result", [
  { no: 0, name: "OK" },
  { no: 1, name: "DENIED" },
]);

/**
 * @generated from message flipchat.account.v1.Flags
 */
export class Flags extends Message<Flags> {
  /**
   * @generated from field: bool is_staff = 1;
   */
  isStaff = false;

  /**
   * @generated from field: flipchat.common.v1.PaymentAmount start_group_cost = 2;
   */
  startGroupCost?: PaymentAmount;

  constructor(data?: PartialMessage<Flags>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.account.v1.Flags";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "is_staff", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "start_group_cost", kind: "message", T: PaymentAmount },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Flags {
    return new Flags().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Flags {
    return new Flags().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Flags {
    return new Flags().fromJsonString(jsonString, options);
  }

  static equals(a: Flags | PlainMessage<Flags> | undefined, b: Flags | PlainMessage<Flags> | undefined): boolean {
    return proto3.util.equals(Flags, a, b);
  }
}

