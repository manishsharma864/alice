// Copyright © 2021 AMIS Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: github.com/getamis/alice/crypto/ot/message.proto

package ot

import (
	binaryfield "github.com/manishsharma864/alice/crypto/binaryfield"
	ecpointgrouplaw "github.com/manishsharma864/alice/crypto/ecpointgrouplaw"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OtChallengeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Challenge [][]byte `protobuf:"bytes,1,rep,name=challenge,proto3" json:"challenge,omitempty"`
}

func (x *OtChallengeMessage) Reset() {
	*x = OtChallengeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_getamis_alice_crypto_ot_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtChallengeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtChallengeMessage) ProtoMessage() {}

func (x *OtChallengeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_getamis_alice_crypto_ot_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtChallengeMessage.ProtoReflect.Descriptor instead.
func (*OtChallengeMessage) Descriptor() ([]byte, []int) {
	return file_github_com_getamis_alice_crypto_ot_message_proto_rawDescGZIP(), []int{0}
}

func (x *OtChallengeMessage) GetChallenge() [][]byte {
	if x != nil {
		return x.Challenge
	}
	return nil
}

type OtReceiverMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seed []byte                            `protobuf:"bytes,1,opt,name=seed,proto3" json:"seed,omitempty"`
	Bi   []*ecpointgrouplaw.EcPointMessage `protobuf:"bytes,2,rep,name=bi,proto3" json:"bi,omitempty"`
}

func (x *OtReceiverMessage) Reset() {
	*x = OtReceiverMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_getamis_alice_crypto_ot_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtReceiverMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtReceiverMessage) ProtoMessage() {}

func (x *OtReceiverMessage) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_getamis_alice_crypto_ot_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtReceiverMessage.ProtoReflect.Descriptor instead.
func (*OtReceiverMessage) Descriptor() ([]byte, []int) {
	return file_github_com_getamis_alice_crypto_ot_message_proto_rawDescGZIP(), []int{1}
}

func (x *OtReceiverMessage) GetSeed() []byte {
	if x != nil {
		return x.Seed
	}
	return nil
}

func (x *OtReceiverMessage) GetBi() []*ecpointgrouplaw.EcPointMessage {
	if x != nil {
		return x.Bi
	}
	return nil
}

type OtSenderMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Z     *ecpointgrouplaw.EcPointMessage `protobuf:"bytes,1,opt,name=z,proto3" json:"z,omitempty"`
	Chall [][]byte                        `protobuf:"bytes,2,rep,name=chall,proto3" json:"chall,omitempty"`
	Gamma []byte                          `protobuf:"bytes,3,opt,name=gamma,proto3" json:"gamma,omitempty"`
}

func (x *OtSenderMessage) Reset() {
	*x = OtSenderMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_getamis_alice_crypto_ot_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtSenderMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtSenderMessage) ProtoMessage() {}

func (x *OtSenderMessage) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_getamis_alice_crypto_ot_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtSenderMessage.ProtoReflect.Descriptor instead.
func (*OtSenderMessage) Descriptor() ([]byte, []int) {
	return file_github_com_getamis_alice_crypto_ot_message_proto_rawDescGZIP(), []int{2}
}

func (x *OtSenderMessage) GetZ() *ecpointgrouplaw.EcPointMessage {
	if x != nil {
		return x.Z
	}
	return nil
}

func (x *OtSenderMessage) GetChall() [][]byte {
	if x != nil {
		return x.Chall
	}
	return nil
}

func (x *OtSenderMessage) GetGamma() []byte {
	if x != nil {
		return x.Gamma
	}
	return nil
}

type OtReceiverVerifyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ans []byte `protobuf:"bytes,1,opt,name=ans,proto3" json:"ans,omitempty"`
}

func (x *OtReceiverVerifyMessage) Reset() {
	*x = OtReceiverVerifyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_getamis_alice_crypto_ot_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtReceiverVerifyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtReceiverVerifyMessage) ProtoMessage() {}

func (x *OtReceiverVerifyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_getamis_alice_crypto_ot_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtReceiverVerifyMessage.ProtoReflect.Descriptor instead.
func (*OtReceiverVerifyMessage) Descriptor() ([]byte, []int) {
	return file_github_com_getamis_alice_crypto_ot_message_proto_rawDescGZIP(), []int{3}
}

func (x *OtReceiverVerifyMessage) GetAns() []byte {
	if x != nil {
		return x.Ans
	}
	return nil
}

type OtDMatrixMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	D [][]byte `protobuf:"bytes,1,rep,name=D,proto3" json:"D,omitempty"`
}

func (x *OtDMatrixMessage) Reset() {
	*x = OtDMatrixMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_getamis_alice_crypto_ot_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtDMatrixMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtDMatrixMessage) ProtoMessage() {}

func (x *OtDMatrixMessage) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_getamis_alice_crypto_ot_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtDMatrixMessage.ProtoReflect.Descriptor instead.
func (*OtDMatrixMessage) Descriptor() ([]byte, []int) {
	return file_github_com_getamis_alice_crypto_ot_message_proto_rawDescGZIP(), []int{4}
}

func (x *OtDMatrixMessage) GetD() [][]byte {
	if x != nil {
		return x.D
	}
	return nil
}

type OtDMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	D []byte `protobuf:"bytes,1,opt,name=D,proto3" json:"D,omitempty"`
}

func (x *OtDMessage) Reset() {
	*x = OtDMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_getamis_alice_crypto_ot_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtDMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtDMessage) ProtoMessage() {}

func (x *OtDMessage) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_getamis_alice_crypto_ot_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtDMessage.ProtoReflect.Descriptor instead.
func (*OtDMessage) Descriptor() ([]byte, []int) {
	return file_github_com_getamis_alice_crypto_ot_message_proto_rawDescGZIP(), []int{5}
}

func (x *OtDMessage) GetD() []byte {
	if x != nil {
		return x.D
	}
	return nil
}

type OtExtReceiveMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OtSendMsg *OtSenderMessage             `protobuf:"bytes,1,opt,name=otSendMsg,proto3" json:"otSendMsg,omitempty"`
	D         [][]byte                     `protobuf:"bytes,2,rep,name=D,proto3" json:"D,omitempty"`
	U         []*binaryfield.BinaryMessage `protobuf:"bytes,3,rep,name=u,proto3" json:"u,omitempty"`
	V         []*binaryfield.BinaryMessage `protobuf:"bytes,4,rep,name=v,proto3" json:"v,omitempty"`
}

func (x *OtExtReceiveMessage) Reset() {
	*x = OtExtReceiveMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_getamis_alice_crypto_ot_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtExtReceiveMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtExtReceiveMessage) ProtoMessage() {}

func (x *OtExtReceiveMessage) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_getamis_alice_crypto_ot_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtExtReceiveMessage.ProtoReflect.Descriptor instead.
func (*OtExtReceiveMessage) Descriptor() ([]byte, []int) {
	return file_github_com_getamis_alice_crypto_ot_message_proto_rawDescGZIP(), []int{6}
}

func (x *OtExtReceiveMessage) GetOtSendMsg() *OtSenderMessage {
	if x != nil {
		return x.OtSendMsg
	}
	return nil
}

func (x *OtExtReceiveMessage) GetD() [][]byte {
	if x != nil {
		return x.D
	}
	return nil
}

func (x *OtExtReceiveMessage) GetU() []*binaryfield.BinaryMessage {
	if x != nil {
		return x.U
	}
	return nil
}

func (x *OtExtReceiveMessage) GetV() []*binaryfield.BinaryMessage {
	if x != nil {
		return x.V
	}
	return nil
}

type OtExtSendResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A0             [][]byte                 `protobuf:"bytes,1,rep,name=a0,proto3" json:"a0,omitempty"`
	A1             [][]byte                 `protobuf:"bytes,2,rep,name=a1,proto3" json:"a1,omitempty"`
	OtRecVerifyMsg *OtReceiverVerifyMessage `protobuf:"bytes,3,opt,name=otRecVerifyMsg,proto3" json:"otRecVerifyMsg,omitempty"`
}

func (x *OtExtSendResponseMessage) Reset() {
	*x = OtExtSendResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_getamis_alice_crypto_ot_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OtExtSendResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtExtSendResponseMessage) ProtoMessage() {}

func (x *OtExtSendResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_getamis_alice_crypto_ot_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtExtSendResponseMessage.ProtoReflect.Descriptor instead.
func (*OtExtSendResponseMessage) Descriptor() ([]byte, []int) {
	return file_github_com_getamis_alice_crypto_ot_message_proto_rawDescGZIP(), []int{7}
}

func (x *OtExtSendResponseMessage) GetA0() [][]byte {
	if x != nil {
		return x.A0
	}
	return nil
}

func (x *OtExtSendResponseMessage) GetA1() [][]byte {
	if x != nil {
		return x.A1
	}
	return nil
}

func (x *OtExtSendResponseMessage) GetOtRecVerifyMsg() *OtReceiverVerifyMessage {
	if x != nil {
		return x.OtRecVerifyMsg
	}
	return nil
}

var File_github_com_getamis_alice_crypto_ot_message_proto protoreflect.FileDescriptor

var file_github_com_getamis_alice_crypto_ot_message_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x74,
	0x61, 0x6d, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2f, 0x6f, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x6f, 0x74, 0x1a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x65, 0x74, 0x61, 0x6d, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x69, 0x63, 0x65,
	0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x65, 0x63, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x6c, 0x61, 0x77, 0x2f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x65, 0x74, 0x61, 0x6d, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32,
	0x0a, 0x12, 0x4f, 0x74, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x22, 0x58, 0x0a, 0x11, 0x4f, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x73, 0x65, 0x65, 0x64, 0x12, 0x2f, 0x0a, 0x02, 0x62,
	0x69, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x63, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x6c, 0x61, 0x77, 0x2e, 0x45, 0x63, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x02, 0x62, 0x69, 0x22, 0x6c, 0x0a, 0x0f,
	0x4f, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x2d, 0x0a, 0x01, 0x7a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x63, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x6c, 0x61, 0x77, 0x2e, 0x45, 0x63, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x01, 0x7a, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x05, 0x63,
	0x68, 0x61, 0x6c, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x61, 0x6d, 0x6d, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x67, 0x61, 0x6d, 0x6d, 0x61, 0x22, 0x2b, 0x0a, 0x17, 0x4f, 0x74,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x03, 0x61, 0x6e, 0x73, 0x22, 0x20, 0x0a, 0x10, 0x4f, 0x74, 0x44, 0x4d, 0x61,
	0x74, 0x72, 0x69, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x44,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x01, 0x44, 0x22, 0x1a, 0x0a, 0x0a, 0x4f, 0x74, 0x44,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x01, 0x44, 0x22, 0xaa, 0x01, 0x0a, 0x13, 0x4f, 0x74, 0x45, 0x78, 0x74, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x31, 0x0a,
	0x09, 0x6f, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x6f, 0x74, 0x2e, 0x4f, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x09, 0x6f, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67,
	0x12, 0x0c, 0x0a, 0x01, 0x44, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x01, 0x44, 0x12, 0x28,
	0x0a, 0x01, 0x75, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x69, 0x6e, 0x61,
	0x72, 0x79, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x01, 0x75, 0x12, 0x28, 0x0a, 0x01, 0x76, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x01, 0x76, 0x22, 0x7f, 0x0a, 0x18, 0x4f, 0x74, 0x45, 0x78, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x61, 0x30, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x02, 0x61, 0x30, 0x12, 0x0e,
	0x0a, 0x02, 0x61, 0x31, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x02, 0x61, 0x31, 0x12, 0x43,
	0x0a, 0x0e, 0x6f, 0x74, 0x52, 0x65, 0x63, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4d, 0x73, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6f, 0x74, 0x2e, 0x4f, 0x74, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x0e, 0x6f, 0x74, 0x52, 0x65, 0x63, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x4d, 0x73, 0x67, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x65, 0x74, 0x61, 0x6d, 0x69, 0x73, 0x2f, 0x61, 0x6c, 0x69, 0x63, 0x65, 0x2f,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x6f, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_github_com_getamis_alice_crypto_ot_message_proto_rawDescOnce sync.Once
	file_github_com_getamis_alice_crypto_ot_message_proto_rawDescData = file_github_com_getamis_alice_crypto_ot_message_proto_rawDesc
)

func file_github_com_getamis_alice_crypto_ot_message_proto_rawDescGZIP() []byte {
	file_github_com_getamis_alice_crypto_ot_message_proto_rawDescOnce.Do(func() {
		file_github_com_getamis_alice_crypto_ot_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_getamis_alice_crypto_ot_message_proto_rawDescData)
	})
	return file_github_com_getamis_alice_crypto_ot_message_proto_rawDescData
}

var file_github_com_getamis_alice_crypto_ot_message_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_github_com_getamis_alice_crypto_ot_message_proto_goTypes = []interface{}{
	(*OtChallengeMessage)(nil),             // 0: ot.OtChallengeMessage
	(*OtReceiverMessage)(nil),              // 1: ot.OtReceiverMessage
	(*OtSenderMessage)(nil),                // 2: ot.OtSenderMessage
	(*OtReceiverVerifyMessage)(nil),        // 3: ot.OtReceiverVerifyMessage
	(*OtDMatrixMessage)(nil),               // 4: ot.OtDMatrixMessage
	(*OtDMessage)(nil),                     // 5: ot.OtDMessage
	(*OtExtReceiveMessage)(nil),            // 6: ot.OtExtReceiveMessage
	(*OtExtSendResponseMessage)(nil),       // 7: ot.OtExtSendResponseMessage
	(*ecpointgrouplaw.EcPointMessage)(nil), // 8: ecpointgrouplaw.EcPointMessage
	(*binaryfield.BinaryMessage)(nil),      // 9: binaryfield.binaryMessage
}
var file_github_com_getamis_alice_crypto_ot_message_proto_depIdxs = []int32{
	8, // 0: ot.OtReceiverMessage.bi:type_name -> ecpointgrouplaw.EcPointMessage
	8, // 1: ot.OtSenderMessage.z:type_name -> ecpointgrouplaw.EcPointMessage
	2, // 2: ot.OtExtReceiveMessage.otSendMsg:type_name -> ot.OtSenderMessage
	9, // 3: ot.OtExtReceiveMessage.u:type_name -> binaryfield.binaryMessage
	9, // 4: ot.OtExtReceiveMessage.v:type_name -> binaryfield.binaryMessage
	3, // 5: ot.OtExtSendResponseMessage.otRecVerifyMsg:type_name -> ot.OtReceiverVerifyMessage
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_github_com_getamis_alice_crypto_ot_message_proto_init() }
func file_github_com_getamis_alice_crypto_ot_message_proto_init() {
	if File_github_com_getamis_alice_crypto_ot_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_getamis_alice_crypto_ot_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtChallengeMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_getamis_alice_crypto_ot_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtReceiverMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_getamis_alice_crypto_ot_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtSenderMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_getamis_alice_crypto_ot_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtReceiverVerifyMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_getamis_alice_crypto_ot_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtDMatrixMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_getamis_alice_crypto_ot_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtDMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_getamis_alice_crypto_ot_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtExtReceiveMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_getamis_alice_crypto_ot_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OtExtSendResponseMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_getamis_alice_crypto_ot_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_getamis_alice_crypto_ot_message_proto_goTypes,
		DependencyIndexes: file_github_com_getamis_alice_crypto_ot_message_proto_depIdxs,
		MessageInfos:      file_github_com_getamis_alice_crypto_ot_message_proto_msgTypes,
	}.Build()
	File_github_com_getamis_alice_crypto_ot_message_proto = out.File
	file_github_com_getamis_alice_crypto_ot_message_proto_rawDesc = nil
	file_github_com_getamis_alice_crypto_ot_message_proto_goTypes = nil
	file_github_com_getamis_alice_crypto_ot_message_proto_depIdxs = nil
}
