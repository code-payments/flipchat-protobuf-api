// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file messaging/v1/model.proto (package flipchat.messaging.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message as Message$1, proto3, Timestamp } from "@bufbuild/protobuf";
import { PaymentAmount, UserId } from "../../common/v1/common_pb";

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
 * @generated from message flipchat.messaging.v1.MessageIdBatch
 */
export class MessageIdBatch extends Message$1<MessageIdBatch> {
  /**
   * @generated from field: repeated flipchat.messaging.v1.MessageId message_ids = 1;
   */
  messageIds: MessageId[] = [];

  constructor(data?: PartialMessage<MessageIdBatch>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.MessageIdBatch";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "message_ids", kind: "message", T: MessageId, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MessageIdBatch {
    return new MessageIdBatch().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MessageIdBatch {
    return new MessageIdBatch().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MessageIdBatch {
    return new MessageIdBatch().fromJsonString(jsonString, options);
  }

  static equals(a: MessageIdBatch | PlainMessage<MessageIdBatch> | undefined, b: MessageIdBatch | PlainMessage<MessageIdBatch> | undefined): boolean {
    return proto3.util.equals(MessageIdBatch, a, b);
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
   * Message content, which is currently guaranteed to have exactly one item.
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

  /**
   * If sender_id is provided, were they off stage at the time of sending
   * this message
   *
   * @generated from field: bool was_sender_off_stage = 5;
   */
  wasSenderOffStage = false;

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
    { no: 5, name: "was_sender_off_stage", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
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
 * @generated from message flipchat.messaging.v1.MessageBatch
 */
export class MessageBatch extends Message$1<MessageBatch> {
  /**
   * @generated from field: repeated flipchat.messaging.v1.Message messages = 1;
   */
  messages: Message[] = [];

  constructor(data?: PartialMessage<MessageBatch>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.MessageBatch";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "messages", kind: "message", T: Message, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MessageBatch {
    return new MessageBatch().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MessageBatch {
    return new MessageBatch().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MessageBatch {
    return new MessageBatch().fromJsonString(jsonString, options);
  }

  static equals(a: MessageBatch | PlainMessage<MessageBatch> | undefined, b: MessageBatch | PlainMessage<MessageBatch> | undefined): boolean {
    return proto3.util.equals(MessageBatch, a, b);
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
     * @generated from field: flipchat.messaging.v1.ReactionContent reaction = 5;
     */
    value: ReactionContent;
    case: "reaction";
  } | {
    /**
     * @generated from field: flipchat.messaging.v1.ReplyContent reply = 6;
     */
    value: ReplyContent;
    case: "reply";
  } | {
    /**
     * @generated from field: flipchat.messaging.v1.TipContent tip = 7;
     */
    value: TipContent;
    case: "tip";
  } | {
    /**
     * @generated from field: flipchat.messaging.v1.DeleteMessageContent deleted = 8;
     */
    value: DeleteMessageContent;
    case: "deleted";
  } | {
    /**
     * @generated from field: flipchat.messaging.v1.ReviewContent review = 9;
     */
    value: ReviewContent;
    case: "review";
  } | {
    /**
     * @generated from field: flipchat.messaging.v1.ActionableAnnouncementContent actionable_announcement = 10;
     */
    value: ActionableAnnouncementContent;
    case: "actionableAnnouncement";
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
    { no: 5, name: "reaction", kind: "message", T: ReactionContent, oneof: "type" },
    { no: 6, name: "reply", kind: "message", T: ReplyContent, oneof: "type" },
    { no: 7, name: "tip", kind: "message", T: TipContent, oneof: "type" },
    { no: 8, name: "deleted", kind: "message", T: DeleteMessageContent, oneof: "type" },
    { no: 9, name: "review", kind: "message", T: ReviewContent, oneof: "type" },
    { no: 10, name: "actionable_announcement", kind: "message", T: ActionableAnnouncementContent, oneof: "type" },
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
 * ActionableAnnouncementContent is like LocalizedAnnouncementContent, but
 * contains additional metadata for actions
 *
 * @generated from message flipchat.messaging.v1.ActionableAnnouncementContent
 */
export class ActionableAnnouncementContent extends Message$1<ActionableAnnouncementContent> {
  /**
   * @generated from field: string key_or_text = 1;
   */
  keyOrText = "";

  /**
   * An action that can be taken by a user
   *
   * @generated from field: flipchat.messaging.v1.ActionableAnnouncementContent.Action action = 3;
   */
  action?: ActionableAnnouncementContent_Action;

  constructor(data?: PartialMessage<ActionableAnnouncementContent>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.ActionableAnnouncementContent";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "key_or_text", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "action", kind: "message", T: ActionableAnnouncementContent_Action },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ActionableAnnouncementContent {
    return new ActionableAnnouncementContent().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ActionableAnnouncementContent {
    return new ActionableAnnouncementContent().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ActionableAnnouncementContent {
    return new ActionableAnnouncementContent().fromJsonString(jsonString, options);
  }

  static equals(a: ActionableAnnouncementContent | PlainMessage<ActionableAnnouncementContent> | undefined, b: ActionableAnnouncementContent | PlainMessage<ActionableAnnouncementContent> | undefined): boolean {
    return proto3.util.equals(ActionableAnnouncementContent, a, b);
  }
}

/**
 * @generated from message flipchat.messaging.v1.ActionableAnnouncementContent.Action
 */
export class ActionableAnnouncementContent_Action extends Message$1<ActionableAnnouncementContent_Action> {
  /**
   * @generated from oneof flipchat.messaging.v1.ActionableAnnouncementContent.Action.type
   */
  type: {
    /**
     * @generated from field: flipchat.messaging.v1.ActionableAnnouncementContent.Action.ShareRoomLink share_room_link = 1;
     */
    value: ActionableAnnouncementContent_Action_ShareRoomLink;
    case: "shareRoomLink";
  } | { case: undefined; value?: undefined } = { case: undefined };

  constructor(data?: PartialMessage<ActionableAnnouncementContent_Action>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.ActionableAnnouncementContent.Action";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "share_room_link", kind: "message", T: ActionableAnnouncementContent_Action_ShareRoomLink, oneof: "type" },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ActionableAnnouncementContent_Action {
    return new ActionableAnnouncementContent_Action().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ActionableAnnouncementContent_Action {
    return new ActionableAnnouncementContent_Action().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ActionableAnnouncementContent_Action {
    return new ActionableAnnouncementContent_Action().fromJsonString(jsonString, options);
  }

  static equals(a: ActionableAnnouncementContent_Action | PlainMessage<ActionableAnnouncementContent_Action> | undefined, b: ActionableAnnouncementContent_Action | PlainMessage<ActionableAnnouncementContent_Action> | undefined): boolean {
    return proto3.util.equals(ActionableAnnouncementContent_Action, a, b);
  }
}

/**
 * Displays a button to share a link to a room
 *
 * @generated from message flipchat.messaging.v1.ActionableAnnouncementContent.Action.ShareRoomLink
 */
export class ActionableAnnouncementContent_Action_ShareRoomLink extends Message$1<ActionableAnnouncementContent_Action_ShareRoomLink> {
  constructor(data?: PartialMessage<ActionableAnnouncementContent_Action_ShareRoomLink>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.ActionableAnnouncementContent.Action.ShareRoomLink";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ActionableAnnouncementContent_Action_ShareRoomLink {
    return new ActionableAnnouncementContent_Action_ShareRoomLink().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ActionableAnnouncementContent_Action_ShareRoomLink {
    return new ActionableAnnouncementContent_Action_ShareRoomLink().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ActionableAnnouncementContent_Action_ShareRoomLink {
    return new ActionableAnnouncementContent_Action_ShareRoomLink().fromJsonString(jsonString, options);
  }

  static equals(a: ActionableAnnouncementContent_Action_ShareRoomLink | PlainMessage<ActionableAnnouncementContent_Action_ShareRoomLink> | undefined, b: ActionableAnnouncementContent_Action_ShareRoomLink | PlainMessage<ActionableAnnouncementContent_Action_ShareRoomLink> | undefined): boolean {
    return proto3.util.equals(ActionableAnnouncementContent_Action_ShareRoomLink, a, b);
  }
}

/**
 * Emoji reaction to another message
 *
 * @generated from message flipchat.messaging.v1.ReactionContent
 */
export class ReactionContent extends Message$1<ReactionContent> {
  /**
   * The message ID of the message this reaction is associated with
   *
   * @generated from field: flipchat.messaging.v1.MessageId original_message_id = 1;
   */
  originalMessageId?: MessageId;

  /**
   * The emoji or reaction symbol
   *
   * @generated from field: string emoji = 2;
   */
  emoji = "";

  constructor(data?: PartialMessage<ReactionContent>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.ReactionContent";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "original_message_id", kind: "message", T: MessageId },
    { no: 2, name: "emoji", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ReactionContent {
    return new ReactionContent().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ReactionContent {
    return new ReactionContent().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ReactionContent {
    return new ReactionContent().fromJsonString(jsonString, options);
  }

  static equals(a: ReactionContent | PlainMessage<ReactionContent> | undefined, b: ReactionContent | PlainMessage<ReactionContent> | undefined): boolean {
    return proto3.util.equals(ReactionContent, a, b);
  }
}

/**
 * Text reply of another message
 *
 * @generated from message flipchat.messaging.v1.ReplyContent
 */
export class ReplyContent extends Message$1<ReplyContent> {
  /**
   * The message ID of the message this reply is referencing
   *
   * @generated from field: flipchat.messaging.v1.MessageId original_message_id = 1;
   */
  originalMessageId?: MessageId;

  /**
   * The reply text, which can be handled similarly to TextContent
   *
   * @generated from field: string reply_text = 2;
   */
  replyText = "";

  constructor(data?: PartialMessage<ReplyContent>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.ReplyContent";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "original_message_id", kind: "message", T: MessageId },
    { no: 2, name: "reply_text", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ReplyContent {
    return new ReplyContent().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ReplyContent {
    return new ReplyContent().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ReplyContent {
    return new ReplyContent().fromJsonString(jsonString, options);
  }

  static equals(a: ReplyContent | PlainMessage<ReplyContent> | undefined, b: ReplyContent | PlainMessage<ReplyContent> | undefined): boolean {
    return proto3.util.equals(ReplyContent, a, b);
  }
}

/**
 * @generated from message flipchat.messaging.v1.TipContent
 */
export class TipContent extends Message$1<TipContent> {
  /**
   * The message ID of the message this tip is referencing
   *
   * @generated from field: flipchat.messaging.v1.MessageId original_message_id = 1;
   */
  originalMessageId?: MessageId;

  /**
   * The amount tipped for the message
   *
   * @generated from field: flipchat.common.v1.PaymentAmount tip_amount = 2;
   */
  tipAmount?: PaymentAmount;

  constructor(data?: PartialMessage<TipContent>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.TipContent";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "original_message_id", kind: "message", T: MessageId },
    { no: 2, name: "tip_amount", kind: "message", T: PaymentAmount },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TipContent {
    return new TipContent().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TipContent {
    return new TipContent().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TipContent {
    return new TipContent().fromJsonString(jsonString, options);
  }

  static equals(a: TipContent | PlainMessage<TipContent> | undefined, b: TipContent | PlainMessage<TipContent> | undefined): boolean {
    return proto3.util.equals(TipContent, a, b);
  }
}

/**
 * @generated from message flipchat.messaging.v1.DeleteMessageContent
 */
export class DeleteMessageContent extends Message$1<DeleteMessageContent> {
  /**
   * The message ID of the message that was deleted
   *
   * @generated from field: flipchat.messaging.v1.MessageId original_message_id = 1;
   */
  originalMessageId?: MessageId;

  constructor(data?: PartialMessage<DeleteMessageContent>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.DeleteMessageContent";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "original_message_id", kind: "message", T: MessageId },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteMessageContent {
    return new DeleteMessageContent().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteMessageContent {
    return new DeleteMessageContent().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteMessageContent {
    return new DeleteMessageContent().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteMessageContent | PlainMessage<DeleteMessageContent> | undefined, b: DeleteMessageContent | PlainMessage<DeleteMessageContent> | undefined): boolean {
    return proto3.util.equals(DeleteMessageContent, a, b);
  }
}

/**
 * @generated from message flipchat.messaging.v1.ReviewContent
 */
export class ReviewContent extends Message$1<ReviewContent> {
  /**
   * The message ID of the message that is being reviewed. Currently, only
   * off stage messages can be reviewed
   *
   * @generated from field: flipchat.messaging.v1.MessageId original_message_id = 1;
   */
  originalMessageId?: MessageId;

  /**
   * Whether the message has been approved. In the event of multiple reviews,
   * the first message in the message log takes priority.
   *
   * @generated from field: bool is_approved = 2;
   */
  isApproved = false;

  constructor(data?: PartialMessage<ReviewContent>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "flipchat.messaging.v1.ReviewContent";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "original_message_id", kind: "message", T: MessageId },
    { no: 2, name: "is_approved", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ReviewContent {
    return new ReviewContent().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ReviewContent {
    return new ReviewContent().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ReviewContent {
    return new ReviewContent().fromJsonString(jsonString, options);
  }

  static equals(a: ReviewContent | PlainMessage<ReviewContent> | undefined, b: ReviewContent | PlainMessage<ReviewContent> | undefined): boolean {
    return proto3.util.equals(ReviewContent, a, b);
  }
}

