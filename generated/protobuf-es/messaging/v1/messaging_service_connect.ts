// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=ts"
// @generated from file messaging/v1/messaging_service.proto (package flipchat.messaging.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AdvancePointerRequest, AdvancePointerResponse, GetMessageRequest, GetMessageResponse, GetMessagesRequest, GetMessagesResponse, NotifyIsTypingRequest, NotifyIsTypingResponse, SendMessageRequest, SendMessageResponse, StreamMessagesRequest, StreamMessagesResponse } from "./messaging_service_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service flipchat.messaging.v1.Messaging
 */
export const Messaging = {
  typeName: "flipchat.messaging.v1.Messaging",
  methods: {
    /**
     * StreamMessages streams all messages/message states (eg. pointers, typing, etc)
     * for the requested chat.
     *
     * @generated from rpc flipchat.messaging.v1.Messaging.StreamMessages
     */
    streamMessages: {
      name: "StreamMessages",
      I: StreamMessagesRequest,
      O: StreamMessagesResponse,
      kind: MethodKind.BiDiStreaming,
    },
    /**
     * GetMessage gets a single message in a chat
     *
     * @generated from rpc flipchat.messaging.v1.Messaging.GetMessage
     */
    getMessage: {
      name: "GetMessage",
      I: GetMessageRequest,
      O: GetMessageResponse,
      kind: MethodKind.Unary,
    },
    /**
     * GetMessages gets the set of messages for a chat using a paged and batched APIs
     *
     * @generated from rpc flipchat.messaging.v1.Messaging.GetMessages
     */
    getMessages: {
      name: "GetMessages",
      I: GetMessagesRequest,
      O: GetMessagesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * SendMessage sends a message to a chat.
     *
     * @generated from rpc flipchat.messaging.v1.Messaging.SendMessage
     */
    sendMessage: {
      name: "SendMessage",
      I: SendMessageRequest,
      O: SendMessageResponse,
      kind: MethodKind.Unary,
    },
    /**
     * AdvancePointer advances a pointer in message history for a chat member.
     *
     * @generated from rpc flipchat.messaging.v1.Messaging.AdvancePointer
     */
    advancePointer: {
      name: "AdvancePointer",
      I: AdvancePointerRequest,
      O: AdvancePointerResponse,
      kind: MethodKind.Unary,
    },
    /**
     * NotifyIsTypingRequest notifies a chat that the sending member is typing.
     *
     * These requests are transient, and may be dropped at any point.
     *
     * @generated from rpc flipchat.messaging.v1.Messaging.NotifyIsTyping
     */
    notifyIsTyping: {
      name: "NotifyIsTyping",
      I: NotifyIsTypingRequest,
      O: NotifyIsTypingResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

