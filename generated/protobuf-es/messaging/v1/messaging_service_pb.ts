// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file messaging/v1/messaging_service.proto (package flipchat.messaging.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { Auth, ChatId, ClientPong, QueryOptions, ServerPing } from "../../common/v1/model_pb";
import { Content, Message as Message$1, MessageId, Pointer } from "./model_pb";

/**
 * @generated from message flipchat.messaging.v1.StreamMessagesRequest
 */
export class StreamMessagesRequest extends Message<StreamMessagesRequest> {
  /**
   * @generated from oneof flipchat.messaging.v1.StreamMessagesRequest.type
   */
  type: {
    /**
     * @generated from field: flipchat.messaging.v1.StreamMessagesRequest.Params params = 1;
     */
    value: StreamMessagesRequest_Params;
    case: "params";
  } | {
    /**
     * @generated from field: flipchat.common.v1.ClientPong pong = 2;
     */
    value: ClientPong;
    case: "pong";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<StreamMessagesRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.StreamMessagesRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "params", kind: "message", T: StreamMessagesRequest_Params, oneof: "type" },
    { no: 2, name: "pong", kind: "message", T: ClientPong, oneof: "type" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StreamMessagesRequest {
    return new StreamMessagesRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StreamMessagesRequest {
    return new StreamMessagesRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StreamMessagesRequest {
    return new StreamMessagesRequest().fromJsonString(jsonString, options);
  }

  static equals(a: StreamMessagesRequest | PlainMessage<StreamMessagesRequest> | undefined, b: StreamMessagesRequest | PlainMessage<StreamMessagesRequest> | undefined): boolean {
    return proto3.util.equals(StreamMessagesRequest, a, b);
  }
}

/**
 * @generated from message flipchat.messaging.v1.StreamMessagesRequest.Params
 */
export class StreamMessagesRequest_Params extends Message<StreamMessagesRequest_Params> {
  /**
   * @generated from field: flipchat.common.v1.ChatId chat_id = 1;
   */
  chatId?: ChatId;

  /**
   * Callers may optionally specify a resume mode other than last delivery pointer.
   *
   * @generated from oneof flipchat.messaging.v1.StreamMessagesRequest.Params.resume
   */
  resume: {
    /**
     * Server will return all messages newer than this message id.
     *
     * @generated from field: flipchat.messaging.v1.MessageId last_known_message_id = 2;
     */
    value: MessageId;
    case: "lastKnownMessageId";
  } | {
    /**
     * Server will not load any previous messages.
     *
     * @generated from field: bool latest_only = 3;
     */
    value: boolean;
    case: "latestOnly";
  } | { case: undefined; value?: undefined } = { case: undefined };

  /**
   * @generated from field: flipchat.common.v1.Auth auth = 4;
   */
  auth?: Auth;

  constructor(data?: PartialMessage<StreamMessagesRequest_Params>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.StreamMessagesRequest.Params";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "chat_id", kind: "message", T: ChatId },
    { no: 2, name: "last_known_message_id", kind: "message", T: MessageId, oneof: "resume" },
    { no: 3, name: "latest_only", kind: "scalar", T: 8 /* ScalarType.BOOL */, oneof: "resume" },
    { no: 4, name: "auth", kind: "message", T: Auth },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StreamMessagesRequest_Params {
    return new StreamMessagesRequest_Params().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StreamMessagesRequest_Params {
    return new StreamMessagesRequest_Params().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StreamMessagesRequest_Params {
    return new StreamMessagesRequest_Params().fromJsonString(jsonString, options);
  }

  static equals(a: StreamMessagesRequest_Params | PlainMessage<StreamMessagesRequest_Params> | undefined, b: StreamMessagesRequest_Params | PlainMessage<StreamMessagesRequest_Params> | undefined): boolean {
    return proto3.util.equals(StreamMessagesRequest_Params, a, b);
  }
}

/**
 * @generated from message flipchat.messaging.v1.StreamMessagesResponse
 */
export class StreamMessagesResponse extends Message<StreamMessagesResponse> {
  /**
   * @generated from oneof flipchat.messaging.v1.StreamMessagesResponse.type
   */
  type: {
    /**
     * @generated from field: flipchat.common.v1.ServerPing ping = 1;
     */
    value: ServerPing;
    case: "ping";
  } | {
    /**
     * @generated from field: flipchat.messaging.v1.StreamMessagesResponse.StreamError error = 2;
     */
    value: StreamMessagesResponse_StreamError;
    case: "error";
  } | {
    /**
     * @generated from field: flipchat.messaging.v1.StreamMessagesResponse.MessageBatch messages = 3;
     */
    value: StreamMessagesResponse_MessageBatch;
    case: "messages";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<StreamMessagesResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.StreamMessagesResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "ping", kind: "message", T: ServerPing, oneof: "type" },
    { no: 2, name: "error", kind: "message", T: StreamMessagesResponse_StreamError, oneof: "type" },
    { no: 3, name: "messages", kind: "message", T: StreamMessagesResponse_MessageBatch, oneof: "type" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StreamMessagesResponse {
    return new StreamMessagesResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StreamMessagesResponse {
    return new StreamMessagesResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StreamMessagesResponse {
    return new StreamMessagesResponse().fromJsonString(jsonString, options);
  }

  static equals(a: StreamMessagesResponse | PlainMessage<StreamMessagesResponse> | undefined, b: StreamMessagesResponse | PlainMessage<StreamMessagesResponse> | undefined): boolean {
    return proto3.util.equals(StreamMessagesResponse, a, b);
  }
}

/**
 * @generated from message flipchat.messaging.v1.StreamMessagesResponse.StreamError
 */
export class StreamMessagesResponse_StreamError extends Message<StreamMessagesResponse_StreamError> {
  /**
   * @generated from field: flipchat.messaging.v1.StreamMessagesResponse.StreamError.Code code = 1;
   */
  code = StreamMessagesResponse_StreamError_Code.DENIED;

  constructor(data?: PartialMessage<StreamMessagesResponse_StreamError>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.StreamMessagesResponse.StreamError";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "code", kind: "enum", T: proto3.getEnumType(StreamMessagesResponse_StreamError_Code) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StreamMessagesResponse_StreamError {
    return new StreamMessagesResponse_StreamError().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StreamMessagesResponse_StreamError {
    return new StreamMessagesResponse_StreamError().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StreamMessagesResponse_StreamError {
    return new StreamMessagesResponse_StreamError().fromJsonString(jsonString, options);
  }

  static equals(a: StreamMessagesResponse_StreamError | PlainMessage<StreamMessagesResponse_StreamError> | undefined, b: StreamMessagesResponse_StreamError | PlainMessage<StreamMessagesResponse_StreamError> | undefined): boolean {
    return proto3.util.equals(StreamMessagesResponse_StreamError, a, b);
  }
}

/**
 * @generated from enum flipchat.messaging.v1.StreamMessagesResponse.StreamError.Code
 */
export enum StreamMessagesResponse_StreamError_Code {
  /**
   * @generated from enum value: DENIED = 0;
   */
  DENIED = 0,
}
// Retrieve enum metadata with: proto3.getEnumType(StreamMessagesResponse_StreamError_Code)
proto3.util.setEnumType(StreamMessagesResponse_StreamError_Code, "flipchat.messaging.v1.StreamMessagesResponse.StreamError.Code", [
  { no: 0, name: "DENIED" },
]);

/**
 * @generated from message flipchat.messaging.v1.StreamMessagesResponse.MessageBatch
 */
export class StreamMessagesResponse_MessageBatch extends Message<StreamMessagesResponse_MessageBatch> {
  /**
   * @generated from field: repeated flipchat.messaging.v1.Message messages = 1;
   */
  messages: Message$1[] = [];

  constructor(data?: PartialMessage<StreamMessagesResponse_MessageBatch>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.StreamMessagesResponse.MessageBatch";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "messages", kind: "message", T: Message$1, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StreamMessagesResponse_MessageBatch {
    return new StreamMessagesResponse_MessageBatch().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StreamMessagesResponse_MessageBatch {
    return new StreamMessagesResponse_MessageBatch().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StreamMessagesResponse_MessageBatch {
    return new StreamMessagesResponse_MessageBatch().fromJsonString(jsonString, options);
  }

  static equals(a: StreamMessagesResponse_MessageBatch | PlainMessage<StreamMessagesResponse_MessageBatch> | undefined, b: StreamMessagesResponse_MessageBatch | PlainMessage<StreamMessagesResponse_MessageBatch> | undefined): boolean {
    return proto3.util.equals(StreamMessagesResponse_MessageBatch, a, b);
  }
}

/**
 * @generated from message flipchat.messaging.v1.GetMessagesRequest
 */
export class GetMessagesRequest extends Message<GetMessagesRequest> {
  /**
   * @generated from field: flipchat.common.v1.ChatId chat_id = 1;
   */
  chatId?: ChatId;

  /**
   * @generated from field: flipchat.common.v1.QueryOptions query_options = 2;
   */
  queryOptions?: QueryOptions;

  /**
   * @generated from field: flipchat.common.v1.Auth auth = 5;
   */
  auth?: Auth;

  constructor(data?: PartialMessage<GetMessagesRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.GetMessagesRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "chat_id", kind: "message", T: ChatId },
    { no: 2, name: "query_options", kind: "message", T: QueryOptions },
    { no: 5, name: "auth", kind: "message", T: Auth },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetMessagesRequest {
    return new GetMessagesRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetMessagesRequest {
    return new GetMessagesRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetMessagesRequest {
    return new GetMessagesRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetMessagesRequest | PlainMessage<GetMessagesRequest> | undefined, b: GetMessagesRequest | PlainMessage<GetMessagesRequest> | undefined): boolean {
    return proto3.util.equals(GetMessagesRequest, a, b);
  }
}

/**
 * @generated from message flipchat.messaging.v1.GetMessagesResponse
 */
export class GetMessagesResponse extends Message<GetMessagesResponse> {
  /**
   * @generated from field: flipchat.messaging.v1.GetMessagesResponse.Result result = 1;
   */
  result = GetMessagesResponse_Result.OK;

  /**
   * @generated from field: repeated flipchat.messaging.v1.Message messages = 2;
   */
  messages: Message$1[] = [];

  constructor(data?: PartialMessage<GetMessagesResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.GetMessagesResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "result", kind: "enum", T: proto3.getEnumType(GetMessagesResponse_Result) },
    { no: 2, name: "messages", kind: "message", T: Message$1, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetMessagesResponse {
    return new GetMessagesResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetMessagesResponse {
    return new GetMessagesResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetMessagesResponse {
    return new GetMessagesResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetMessagesResponse | PlainMessage<GetMessagesResponse> | undefined, b: GetMessagesResponse | PlainMessage<GetMessagesResponse> | undefined): boolean {
    return proto3.util.equals(GetMessagesResponse, a, b);
  }
}

/**
 * @generated from enum flipchat.messaging.v1.GetMessagesResponse.Result
 */
export enum GetMessagesResponse_Result {
  /**
   * @generated from enum value: OK = 0;
   */
  OK = 0,

  /**
   * @generated from enum value: DENIED = 1;
   */
  DENIED = 1,
}
// Retrieve enum metadata with: proto3.getEnumType(GetMessagesResponse_Result)
proto3.util.setEnumType(GetMessagesResponse_Result, "flipchat.messaging.v1.GetMessagesResponse.Result", [
  { no: 0, name: "OK" },
  { no: 1, name: "DENIED" },
]);

/**
 * @generated from message flipchat.messaging.v1.SendMessageRequest
 */
export class SendMessageRequest extends Message<SendMessageRequest> {
  /**
   * @generated from field: flipchat.common.v1.ChatId chat_id = 1;
   */
  chatId?: ChatId;

  /**
   * Allowed content types that can be sent by client:
   *  - TextContent
   *  - ThankYouContent
   *
   * @generated from field: repeated flipchat.messaging.v1.Content content = 2;
   */
  content: Content[] = [];

  /**
   * @generated from field: flipchat.common.v1.Auth auth = 3;
   */
  auth?: Auth;

  constructor(data?: PartialMessage<SendMessageRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.SendMessageRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "chat_id", kind: "message", T: ChatId },
    { no: 2, name: "content", kind: "message", T: Content, repeated: true },
    { no: 3, name: "auth", kind: "message", T: Auth },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SendMessageRequest {
    return new SendMessageRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SendMessageRequest {
    return new SendMessageRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SendMessageRequest {
    return new SendMessageRequest().fromJsonString(jsonString, options);
  }

  static equals(a: SendMessageRequest | PlainMessage<SendMessageRequest> | undefined, b: SendMessageRequest | PlainMessage<SendMessageRequest> | undefined): boolean {
    return proto3.util.equals(SendMessageRequest, a, b);
  }
}

/**
 * @generated from message flipchat.messaging.v1.SendMessageResponse
 */
export class SendMessageResponse extends Message<SendMessageResponse> {
  /**
   * @generated from field: flipchat.messaging.v1.SendMessageResponse.Result result = 1;
   */
  result = SendMessageResponse_Result.OK;

  /**
   * The chat message that was sent if the RPC was succesful, which includes
   * server-side metadata like the generated message ID and official timestamp
   *
   * @generated from field: flipchat.messaging.v1.Message message = 2;
   */
  message?: Message$1;

  constructor(data?: PartialMessage<SendMessageResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.SendMessageResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "result", kind: "enum", T: proto3.getEnumType(SendMessageResponse_Result) },
    { no: 2, name: "message", kind: "message", T: Message$1 },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SendMessageResponse {
    return new SendMessageResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SendMessageResponse {
    return new SendMessageResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SendMessageResponse {
    return new SendMessageResponse().fromJsonString(jsonString, options);
  }

  static equals(a: SendMessageResponse | PlainMessage<SendMessageResponse> | undefined, b: SendMessageResponse | PlainMessage<SendMessageResponse> | undefined): boolean {
    return proto3.util.equals(SendMessageResponse, a, b);
  }
}

/**
 * @generated from enum flipchat.messaging.v1.SendMessageResponse.Result
 */
export enum SendMessageResponse_Result {
  /**
   * @generated from enum value: OK = 0;
   */
  OK = 0,

  /**
   * @generated from enum value: DENIED = 1;
   */
  DENIED = 1,

  /**
   * @generated from enum value: INVALID_CONTENT_TYPE = 2;
   */
  INVALID_CONTENT_TYPE = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(SendMessageResponse_Result)
proto3.util.setEnumType(SendMessageResponse_Result, "flipchat.messaging.v1.SendMessageResponse.Result", [
  { no: 0, name: "OK" },
  { no: 1, name: "DENIED" },
  { no: 2, name: "INVALID_CONTENT_TYPE" },
]);

/**
 * @generated from message flipchat.messaging.v1.AdvancePointerRequest
 */
export class AdvancePointerRequest extends Message<AdvancePointerRequest> {
  /**
   * @generated from field: flipchat.common.v1.ChatId chat_id = 1;
   */
  chatId?: ChatId;

  /**
   * @generated from field: flipchat.messaging.v1.Pointer pointer = 2;
   */
  pointer?: Pointer;

  /**
   * @generated from field: flipchat.common.v1.Auth auth = 3;
   */
  auth?: Auth;

  constructor(data?: PartialMessage<AdvancePointerRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.AdvancePointerRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "chat_id", kind: "message", T: ChatId },
    { no: 2, name: "pointer", kind: "message", T: Pointer },
    { no: 3, name: "auth", kind: "message", T: Auth },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AdvancePointerRequest {
    return new AdvancePointerRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AdvancePointerRequest {
    return new AdvancePointerRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AdvancePointerRequest {
    return new AdvancePointerRequest().fromJsonString(jsonString, options);
  }

  static equals(a: AdvancePointerRequest | PlainMessage<AdvancePointerRequest> | undefined, b: AdvancePointerRequest | PlainMessage<AdvancePointerRequest> | undefined): boolean {
    return proto3.util.equals(AdvancePointerRequest, a, b);
  }
}

/**
 * @generated from message flipchat.messaging.v1.AdvancePointerResponse
 */
export class AdvancePointerResponse extends Message<AdvancePointerResponse> {
  /**
   * @generated from field: flipchat.messaging.v1.AdvancePointerResponse.Result result = 1;
   */
  result = AdvancePointerResponse_Result.OK;

  constructor(data?: PartialMessage<AdvancePointerResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.AdvancePointerResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "result", kind: "enum", T: proto3.getEnumType(AdvancePointerResponse_Result) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AdvancePointerResponse {
    return new AdvancePointerResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AdvancePointerResponse {
    return new AdvancePointerResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AdvancePointerResponse {
    return new AdvancePointerResponse().fromJsonString(jsonString, options);
  }

  static equals(a: AdvancePointerResponse | PlainMessage<AdvancePointerResponse> | undefined, b: AdvancePointerResponse | PlainMessage<AdvancePointerResponse> | undefined): boolean {
    return proto3.util.equals(AdvancePointerResponse, a, b);
  }
}

/**
 * @generated from enum flipchat.messaging.v1.AdvancePointerResponse.Result
 */
export enum AdvancePointerResponse_Result {
  /**
   * @generated from enum value: OK = 0;
   */
  OK = 0,

  /**
   * @generated from enum value: DENIED = 1;
   */
  DENIED = 1,

  /**
   * @generated from enum value: MESSAGE_NOT_FOUND = 2;
   */
  MESSAGE_NOT_FOUND = 2,
}
// Retrieve enum metadata with: proto3.getEnumType(AdvancePointerResponse_Result)
proto3.util.setEnumType(AdvancePointerResponse_Result, "flipchat.messaging.v1.AdvancePointerResponse.Result", [
  { no: 0, name: "OK" },
  { no: 1, name: "DENIED" },
  { no: 2, name: "MESSAGE_NOT_FOUND" },
]);

/**
 * @generated from message flipchat.messaging.v1.NotifyIsTypingRequest
 */
export class NotifyIsTypingRequest extends Message<NotifyIsTypingRequest> {
  /**
   * @generated from field: flipchat.common.v1.ChatId chat_id = 1;
   */
  chatId?: ChatId;

  /**
   * @generated from field: bool is_typing = 2;
   */
  isTyping = false;

  /**
   * @generated from field: flipchat.common.v1.Auth auth = 3;
   */
  auth?: Auth;

  constructor(data?: PartialMessage<NotifyIsTypingRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.NotifyIsTypingRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "chat_id", kind: "message", T: ChatId },
    { no: 2, name: "is_typing", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "auth", kind: "message", T: Auth },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NotifyIsTypingRequest {
    return new NotifyIsTypingRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NotifyIsTypingRequest {
    return new NotifyIsTypingRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NotifyIsTypingRequest {
    return new NotifyIsTypingRequest().fromJsonString(jsonString, options);
  }

  static equals(a: NotifyIsTypingRequest | PlainMessage<NotifyIsTypingRequest> | undefined, b: NotifyIsTypingRequest | PlainMessage<NotifyIsTypingRequest> | undefined): boolean {
    return proto3.util.equals(NotifyIsTypingRequest, a, b);
  }
}

/**
 * @generated from message flipchat.messaging.v1.NotifyIsTypingResponse
 */
export class NotifyIsTypingResponse extends Message<NotifyIsTypingResponse> {
  /**
   * @generated from field: flipchat.messaging.v1.NotifyIsTypingResponse.Result result = 1;
   */
  result = NotifyIsTypingResponse_Result.OK;

  constructor(data?: PartialMessage<NotifyIsTypingResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.NotifyIsTypingResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "result", kind: "enum", T: proto3.getEnumType(NotifyIsTypingResponse_Result) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NotifyIsTypingResponse {
    return new NotifyIsTypingResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NotifyIsTypingResponse {
    return new NotifyIsTypingResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NotifyIsTypingResponse {
    return new NotifyIsTypingResponse().fromJsonString(jsonString, options);
  }

  static equals(a: NotifyIsTypingResponse | PlainMessage<NotifyIsTypingResponse> | undefined, b: NotifyIsTypingResponse | PlainMessage<NotifyIsTypingResponse> | undefined): boolean {
    return proto3.util.equals(NotifyIsTypingResponse, a, b);
  }
}

/**
 * @generated from enum flipchat.messaging.v1.NotifyIsTypingResponse.Result
 */
export enum NotifyIsTypingResponse_Result {
  /**
   * @generated from enum value: OK = 0;
   */
  OK = 0,

  /**
   * @generated from enum value: DENIED = 1;
   */
  DENIED = 1,
}
// Retrieve enum metadata with: proto3.getEnumType(NotifyIsTypingResponse_Result)
proto3.util.setEnumType(NotifyIsTypingResponse_Result, "flipchat.messaging.v1.NotifyIsTypingResponse.Result", [
  { no: 0, name: "OK" },
  { no: 1, name: "DENIED" },
]);

