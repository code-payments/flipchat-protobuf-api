// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file messaging/v1/model.proto (package flipchat.messaging.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message as Message$1, proto3, Timestamp } from "@bufbuild/protobuf";
import { PublicKey, UserId } from "../../common/v1/flipchat_pb";

/**
 * @generated from message flipchat.messaging.v1.MessageId
 */
export class MessageId extends Message$1<MessageId> {
  /**
   * A lexicographically sortable ID that can be used to sort source of
   * chat history.
   *
   * @generated from field: bytes value = 1;
   */
  value = new Uint8Array(0);

  constructor(data?: PartialMessage<MessageId>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.MessageId";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "value", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MessageId {
    return new MessageId().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MessageId {
    return new MessageId().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MessageId {
    return new MessageId().fromJsonString(jsonString, options);
  }

  static equals(a: MessageId | PlainMessage<MessageId> | undefined, b: MessageId | PlainMessage<MessageId> | undefined): boolean {
    return proto3.util.equals(MessageId, a, b);
  }
}

/**
 * A message in a chat
 *
 * @generated from message flipchat.messaging.v1.Message
 */
export class Message extends Message$1<Message> {
  /**
   * Globally unique ID for this message
   *
   * @generated from field: flipchat.messaging.v1.MessageId message_id = 1;
   */
  messageId?: MessageId;

  /**
   * The chat member that sent the message. For NOTIFICATION chats, this field
   * is omitted since the chat has exactly 1 member.
   *
   * @generated from field: flipchat.common.v1.UserId sender_id = 2;
   */
  senderId?: UserId;

  /**
   * Ordered message content. A message may have more than one piece of content.
   *
   * @generated from field: repeated flipchat.messaging.v1.Content content = 3;
   */
  content: Content[] = [];

  /**
   * Timestamp this message was generated at. This value is also encoded in
   * any time-based UUID message IDs.
   *
   * @generated from field: google.protobuf.Timestamp ts = 4;
   */
  ts?: Timestamp;

  constructor(data?: PartialMessage<Message>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.Message";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "message_id", kind: "message", T: MessageId },
    { no: 2, name: "sender_id", kind: "message", T: UserId },
    { no: 3, name: "content", kind: "message", T: Content, repeated: true },
    { no: 4, name: "ts", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Message {
    return new Message().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Message {
    return new Message().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Message {
    return new Message().fromJsonString(jsonString, options);
  }

  static equals(a: Message | PlainMessage<Message> | undefined, b: Message | PlainMessage<Message> | undefined): boolean {
    return proto3.util.equals(Message, a, b);
  }
}

/**
 * Pointer in a chat indicating a user's message history state in a chat.
 *
 * @generated from message flipchat.messaging.v1.Pointer
 */
export class Pointer extends Message$1<Pointer> {
  /**
   * The type of pointer indicates which user's message history state can be
   * inferred from the pointer value. It is also possible to infer cross-pointer
   * state. For example, if a chat member has a READ pointer for a message with
   * ID N, then the DELIVERED pointer must be at least N.
   *
   * @generated from field: flipchat.messaging.v1.Pointer.Type type = 1;
   */
  type = Pointer_Type.UNKNOWN;

  /**
   * Everything at or before this message ID is considered to have the state
   * inferred by the type of pointer.
   *
   * @generated from field: flipchat.messaging.v1.MessageId value = 2;
   */
  value?: MessageId;

  constructor(data?: PartialMessage<Pointer>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.Pointer";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "type", kind: "enum", T: proto3.getEnumType(Pointer_Type) },
    { no: 2, name: "value", kind: "message", T: MessageId },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Pointer {
    return new Pointer().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Pointer {
    return new Pointer().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Pointer {
    return new Pointer().fromJsonString(jsonString, options);
  }

  static equals(a: Pointer | PlainMessage<Pointer> | undefined, b: Pointer | PlainMessage<Pointer> | undefined): boolean {
    return proto3.util.equals(Pointer, a, b);
  }
}

/**
 * @generated from enum flipchat.messaging.v1.Pointer.Type
 */
export enum Pointer_Type {
  /**
   * @generated from enum value: UNKNOWN = 0;
   */
  UNKNOWN = 0,

  /**
   * Always inferred by OK result in SendMessageResponse or message presence in a chat
   *
   * @generated from enum value: SENT = 1;
   */
  SENT = 1,

  /**
   * @generated from enum value: DELIVERED = 2;
   */
  DELIVERED = 2,

  /**
   * @generated from enum value: READ = 3;
   */
  READ = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(Pointer_Type)
proto3.util.setEnumType(Pointer_Type, "flipchat.messaging.v1.Pointer.Type", [
  { no: 0, name: "UNKNOWN" },
  { no: 1, name: "SENT" },
  { no: 2, name: "DELIVERED" },
  { no: 3, name: "READ" },
]);

/**
 * @generated from message flipchat.messaging.v1.IsTyping
 */
export class IsTyping extends Message$1<IsTyping> {
  /**
   * @generated from field: flipchat.common.v1.UserId user_id = 1;
   */
  userId?: UserId;

  /**
   * is_typing indicates whether or not the user is typing.
   * If false, the user has explicitly stopped typing.
   *
   * @generated from field: bool is_typing = 2;
   */
  isTyping = false;

  constructor(data?: PartialMessage<IsTyping>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.IsTyping";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_id", kind: "message", T: UserId },
    { no: 2, name: "is_typing", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): IsTyping {
    return new IsTyping().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): IsTyping {
    return new IsTyping().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): IsTyping {
    return new IsTyping().fromJsonString(jsonString, options);
  }

  static equals(a: IsTyping | PlainMessage<IsTyping> | undefined, b: IsTyping | PlainMessage<IsTyping> | undefined): boolean {
    return proto3.util.equals(IsTyping, a, b);
  }
}

/**
 * Content for a chat message
 *
 * @generated from message flipchat.messaging.v1.Content
 */
export class Content extends Message$1<Content> {
  /**
   * @generated from oneof flipchat.messaging.v1.Content.type
   */
  type: {
    /**
     * @generated from field: flipchat.messaging.v1.TextContent text = 1;
     */
    value: TextContent;
    case: "text";
  } | {
    /**
     * @generated from field: flipchat.messaging.v1.LocalizedAnnouncementContent localized_announcement = 2;
     */
    value: LocalizedAnnouncementContent;
    case: "localizedAnnouncement";
  } | {
    /**
     * ExchangeDataContent         exchange_data     = 3;
     *
     * @generated from field: flipchat.messaging.v1.NaclBoxEncryptedContent nacl_box = 4;
     */
    value: NaclBoxEncryptedContent;
    case: "naclBox";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<Content>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.Content";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "text", kind: "message", T: TextContent, oneof: "type" },
    { no: 2, name: "localized_announcement", kind: "message", T: LocalizedAnnouncementContent, oneof: "type" },
    { no: 4, name: "nacl_box", kind: "message", T: NaclBoxEncryptedContent, oneof: "type" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Content {
    return new Content().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Content {
    return new Content().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Content {
    return new Content().fromJsonString(jsonString, options);
  }

  static equals(a: Content | PlainMessage<Content> | undefined, b: Content | PlainMessage<Content> | undefined): boolean {
    return proto3.util.equals(Content, a, b);
  }
}

/**
 * Raw text content
 *
 * @generated from message flipchat.messaging.v1.TextContent
 */
export class TextContent extends Message$1<TextContent> {
  /**
   * @generated from field: string text = 1;
   */
  text = "";

  constructor(data?: PartialMessage<TextContent>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.TextContent";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "text", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TextContent {
    return new TextContent().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TextContent {
    return new TextContent().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TextContent {
    return new TextContent().fromJsonString(jsonString, options);
  }

  static equals(a: TextContent | PlainMessage<TextContent> | undefined, b: TextContent | PlainMessage<TextContent> | undefined): boolean {
    return proto3.util.equals(TextContent, a, b);
  }
}

/**
 * LocalizedAnnouncementContent content is an annoucement that is either a
 * localization key that should be translated on client, or a server-side
 * translated piece of text.
 *
 * @generated from message flipchat.messaging.v1.LocalizedAnnouncementContent
 */
export class LocalizedAnnouncementContent extends Message$1<LocalizedAnnouncementContent> {
  /**
   * @generated from field: string key_or_text = 1;
   */
  keyOrText = "";

  constructor(data?: PartialMessage<LocalizedAnnouncementContent>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.LocalizedAnnouncementContent";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "key_or_text", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LocalizedAnnouncementContent {
    return new LocalizedAnnouncementContent().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LocalizedAnnouncementContent {
    return new LocalizedAnnouncementContent().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LocalizedAnnouncementContent {
    return new LocalizedAnnouncementContent().fromJsonString(jsonString, options);
  }

  static equals(a: LocalizedAnnouncementContent | PlainMessage<LocalizedAnnouncementContent> | undefined, b: LocalizedAnnouncementContent | PlainMessage<LocalizedAnnouncementContent> | undefined): boolean {
    return proto3.util.equals(LocalizedAnnouncementContent, a, b);
  }
}

/**
 * Encrypted piece of content using NaCl box encryption
 *
 * @generated from message flipchat.messaging.v1.NaclBoxEncryptedContent
 */
export class NaclBoxEncryptedContent extends Message$1<NaclBoxEncryptedContent> {
  /**
   * The sender's public key that is used to derive the shared private key for
   * decryption for message content.
   *
   * @generated from field: flipchat.common.v1.PublicKey peer_public_key = 1;
   */
  peerPublicKey?: PublicKey;

  /**
   * Globally random nonce that is unique to this encrypted piece of content
   *
   * @generated from field: bytes nonce = 2;
   */
  nonce = new Uint8Array(0);

  /**
   * The encrypted piece of message content
   *
   * @generated from field: bytes encrypted_payload = 3;
   */
  encryptedPayload = new Uint8Array(0);

  constructor(data?: PartialMessage<NaclBoxEncryptedContent>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.NaclBoxEncryptedContent";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "peer_public_key", kind: "message", T: PublicKey },
    { no: 2, name: "nonce", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 3, name: "encrypted_payload", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NaclBoxEncryptedContent {
    return new NaclBoxEncryptedContent().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NaclBoxEncryptedContent {
    return new NaclBoxEncryptedContent().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NaclBoxEncryptedContent {
    return new NaclBoxEncryptedContent().fromJsonString(jsonString, options);
  }

  static equals(a: NaclBoxEncryptedContent | PlainMessage<NaclBoxEncryptedContent> | undefined, b: NaclBoxEncryptedContent | PlainMessage<NaclBoxEncryptedContent> | undefined): boolean {
    return proto3.util.equals(NaclBoxEncryptedContent, a, b);
  }
}

