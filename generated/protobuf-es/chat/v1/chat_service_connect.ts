// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=ts"
// @generated from file chat/v1/chat_service.proto (package flipchat.chat.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CheckDisplayNameRequest, CheckDisplayNameResponse, CloseChatRequest, CloseChatResponse, DemoteUserRequest, DemoteUserResponse, GetChatRequest, GetChatResponse, GetChatsRequest, GetChatsResponse, GetMemberUpdatesRequest, GetMemberUpdatesResponse, JoinChatRequest, JoinChatResponse, LeaveChatRequest, LeaveChatResponse, MuteChatRequest, MuteChatResponse, MuteUserRequest, MuteUserResponse, OpenChatRequest, OpenChatResponse, PromoteUserRequest, PromoteUserResponse, RemoveUserRequest, RemoveUserResponse, ReportUserRequest, ReportUserResponse, SetCoverChargeRequest, SetCoverChargeResponse, SetDisplayNameRequest, SetDisplayNameResponse, SetMessagingFeeRequest, SetMessagingFeeResponse, StartChatRequest, StartChatResponse, StreamChatEventsRequest, StreamChatEventsResponse, UnmuteChatRequest, UnmuteChatResponse } from "./chat_service_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service flipchat.chat.v1.Chat
 */
export const Chat = {
  typeName: "flipchat.chat.v1.Chat",
  methods: {
    /**
     * StreamChatEvents streams all chat events for the requesting user.
     *
     * Chat events will include any update to a chat, including:
     *   1. Metadata changes.
     *   2. Membership changes.
     *   3. Latest messages.
     *
     * The server will optionally filter out some events depending on load
     * and chat type. For example, Broadcast chats may not receive latest
     * messages.
     *
     * Clients should use GetMessages to backfill in any historical messages
     * for a chat. It should be sufficient to rely on ChatEvents for some types
     * of chats, but using StreamMessages provides a guarentee of message events
     * for all chats.
     *
     * @generated from rpc flipchat.chat.v1.Chat.StreamChatEvents
     */
    streamChatEvents: {
      name: "StreamChatEvents",
      I: StreamChatEventsRequest,
      O: StreamChatEventsResponse,
      kind: MethodKind.BiDiStreaming,
    },
    /**
     * GetChats gets the set of chats for an owner account using a paged API.
     * This RPC is aware of all identities tied to the owner account.
     *
     * @generated from rpc flipchat.chat.v1.Chat.GetChats
     */
    getChats: {
      name: "GetChats",
      I: GetChatsRequest,
      O: GetChatsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * GetChat returns the metadata for a specific chat.
     *
     * @generated from rpc flipchat.chat.v1.Chat.GetChat
     */
    getChat: {
      name: "GetChat",
      I: GetChatRequest,
      O: GetChatResponse,
      kind: MethodKind.Unary,
    },
    /**
     * StartChat starts a chat. The RPC call is idempotent and will use existing
     * chats whenever applicable within the context of message routing.
     *
     * @generated from rpc flipchat.chat.v1.Chat.StartChat
     */
    startChat: {
      name: "StartChat",
      I: StartChatRequest,
      O: StartChatResponse,
      kind: MethodKind.Unary,
    },
    /**
     * JoinChat joins a given chat.
     *
     * @generated from rpc flipchat.chat.v1.Chat.JoinChat
     */
    joinChat: {
      name: "JoinChat",
      I: JoinChatRequest,
      O: JoinChatResponse,
      kind: MethodKind.Unary,
    },
    /**
     * LeaveChat leaves a given chat.
     *
     * @generated from rpc flipchat.chat.v1.Chat.LeaveChat
     */
    leaveChat: {
      name: "LeaveChat",
      I: LeaveChatRequest,
      O: LeaveChatResponse,
      kind: MethodKind.Unary,
    },
    /**
     * OpenChat opens a chat up for messaging across all members
     *
     * @generated from rpc flipchat.chat.v1.Chat.OpenChat
     */
    openChat: {
      name: "OpenChat",
      I: OpenChatRequest,
      O: OpenChatResponse,
      kind: MethodKind.Unary,
    },
    /**
     * CloseChat closes a chat up for messaging to just the chat owner
     *
     * @generated from rpc flipchat.chat.v1.Chat.CloseChat
     */
    closeChat: {
      name: "CloseChat",
      I: CloseChatRequest,
      O: CloseChatResponse,
      kind: MethodKind.Unary,
    },
    /**
     * CheckDisplayName checks whether a chat display name passes moderation
     *
     * @generated from rpc flipchat.chat.v1.Chat.CheckDisplayName
     */
    checkDisplayName: {
      name: "CheckDisplayName",
      I: CheckDisplayNameRequest,
      O: CheckDisplayNameResponse,
      kind: MethodKind.Unary,
    },
    /**
     * SetDisplayName sets a chat's display name. If the display name isn't allowed,
     * then a set of alternate suggestions may be provided
     *
     * @generated from rpc flipchat.chat.v1.Chat.SetDisplayName
     */
    setDisplayName: {
      name: "SetDisplayName",
      I: SetDisplayNameRequest,
      O: SetDisplayNameResponse,
      kind: MethodKind.Unary,
    },
    /**
     * SetCoverCharge sets a chat's cover charge
     *
     * Deprecated: Use SetMessagingFee instead
     *
     * @generated from rpc flipchat.chat.v1.Chat.SetCoverCharge
     */
    setCoverCharge: {
      name: "SetCoverCharge",
      I: SetCoverChargeRequest,
      O: SetCoverChargeResponse,
      kind: MethodKind.Unary,
    },
    /**
     * SetMessagingFee sets a chat's messaging fee
     *
     * @generated from rpc flipchat.chat.v1.Chat.SetMessagingFee
     */
    setMessagingFee: {
      name: "SetMessagingFee",
      I: SetMessagingFeeRequest,
      O: SetMessagingFeeResponse,
      kind: MethodKind.Unary,
    },
    /**
     * GetMemberUpdates gets member updates for a given chat
     *
     * @generated from rpc flipchat.chat.v1.Chat.GetMemberUpdates
     */
    getMemberUpdates: {
      name: "GetMemberUpdates",
      I: GetMemberUpdatesRequest,
      O: GetMemberUpdatesResponse,
      kind: MethodKind.Unary,
    },
    /**
     * PromoteUser promotes a user to an elevated permission state
     *
     * @generated from rpc flipchat.chat.v1.Chat.PromoteUser
     */
    promoteUser: {
      name: "PromoteUser",
      I: PromoteUserRequest,
      O: PromoteUserResponse,
      kind: MethodKind.Unary,
    },
    /**
     * DemoteUser demotes a user to a lower permission state
     *
     * @generated from rpc flipchat.chat.v1.Chat.DemoteUser
     */
    demoteUser: {
      name: "DemoteUser",
      I: DemoteUserRequest,
      O: DemoteUserResponse,
      kind: MethodKind.Unary,
    },
    /**
     * RemoveUser removes a user from a chat
     *
     * @generated from rpc flipchat.chat.v1.Chat.RemoveUser
     */
    removeUser: {
      name: "RemoveUser",
      I: RemoveUserRequest,
      O: RemoveUserResponse,
      kind: MethodKind.Unary,
    },
    /**
     * MuteUser mutes a user in the chat and removes their ability to send messages
     *
     * @generated from rpc flipchat.chat.v1.Chat.MuteUser
     */
    muteUser: {
      name: "MuteUser",
      I: MuteUserRequest,
      O: MuteUserResponse,
      kind: MethodKind.Unary,
    },
    /**
     * MuteChat mutes a chat and disables push notifications
     *
     * @generated from rpc flipchat.chat.v1.Chat.MuteChat
     */
    muteChat: {
      name: "MuteChat",
      I: MuteChatRequest,
      O: MuteChatResponse,
      kind: MethodKind.Unary,
    },
    /**
     * UnmuteChat unmutes a chat and enables push notifications
     *
     * @generated from rpc flipchat.chat.v1.Chat.UnmuteChat
     */
    unmuteChat: {
      name: "UnmuteChat",
      I: UnmuteChatRequest,
      O: UnmuteChatResponse,
      kind: MethodKind.Unary,
    },
    /**
     * ReportUser reports a user for a given message
     *
     * todo: might belong in a different service long-term
     *
     * @generated from rpc flipchat.chat.v1.Chat.ReportUser
     */
    reportUser: {
      name: "ReportUser",
      I: ReportUserRequest,
      O: ReportUserResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

