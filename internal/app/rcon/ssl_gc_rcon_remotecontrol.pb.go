// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssl_gc_rcon_remotecontrol.proto

package rcon

import (
	fmt "fmt"
	state "github.com/RoboCup-SSL/ssl-game-controller/internal/app/state"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// a registration that must be send by the remote control to the controller as the very first message
type RemoteControlRegistration struct {
	// the team to be controlled
	Team *state.Team `protobuf:"varint,1,req,name=team,enum=Team" json:"team,omitempty"`
	// signature can optionally be specified to enable secure communication
	Signature            *Signature `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RemoteControlRegistration) Reset()         { *m = RemoteControlRegistration{} }
func (m *RemoteControlRegistration) String() string { return proto.CompactTextString(m) }
func (*RemoteControlRegistration) ProtoMessage()    {}
func (*RemoteControlRegistration) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7030c3864b7b834, []int{0}
}

func (m *RemoteControlRegistration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteControlRegistration.Unmarshal(m, b)
}
func (m *RemoteControlRegistration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteControlRegistration.Marshal(b, m, deterministic)
}
func (m *RemoteControlRegistration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteControlRegistration.Merge(m, src)
}
func (m *RemoteControlRegistration) XXX_Size() int {
	return xxx_messageInfo_RemoteControlRegistration.Size(m)
}
func (m *RemoteControlRegistration) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteControlRegistration.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteControlRegistration proto.InternalMessageInfo

func (m *RemoteControlRegistration) GetTeam() state.Team {
	if m != nil && m.Team != nil {
		return *m.Team
	}
	return state.Team_UNKNOWN
}

func (m *RemoteControlRegistration) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

// wrapper for all messages from the remote control to the controller
type RemoteControlToController struct {
	// signature can optionally be specified to enable secure communication
	Signature *Signature `protobuf:"bytes,1,opt,name=signature" json:"signature,omitempty"`
	// Types that are valid to be assigned to Msg:
	//	*RemoteControlToController_Ping
	//	*RemoteControlToController_DesiredKeeper
	//	*RemoteControlToController_SubstituteBot
	//	*RemoteControlToController_EmergencyStop
	//	*RemoteControlToController_Timeout
	//	*RemoteControlToController_ChallengeFlag
	Msg                  isRemoteControlToController_Msg `protobuf_oneof:"msg"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *RemoteControlToController) Reset()         { *m = RemoteControlToController{} }
func (m *RemoteControlToController) String() string { return proto.CompactTextString(m) }
func (*RemoteControlToController) ProtoMessage()    {}
func (*RemoteControlToController) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7030c3864b7b834, []int{1}
}

func (m *RemoteControlToController) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteControlToController.Unmarshal(m, b)
}
func (m *RemoteControlToController) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteControlToController.Marshal(b, m, deterministic)
}
func (m *RemoteControlToController) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteControlToController.Merge(m, src)
}
func (m *RemoteControlToController) XXX_Size() int {
	return xxx_messageInfo_RemoteControlToController.Size(m)
}
func (m *RemoteControlToController) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteControlToController.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteControlToController proto.InternalMessageInfo

func (m *RemoteControlToController) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

type isRemoteControlToController_Msg interface {
	isRemoteControlToController_Msg()
}

type RemoteControlToController_Ping struct {
	Ping bool `protobuf:"varint,2,opt,name=ping,oneof"`
}

type RemoteControlToController_DesiredKeeper struct {
	DesiredKeeper int32 `protobuf:"varint,3,opt,name=desired_keeper,json=desiredKeeper,oneof"`
}

type RemoteControlToController_SubstituteBot struct {
	SubstituteBot bool `protobuf:"varint,4,opt,name=substitute_bot,json=substituteBot,oneof"`
}

type RemoteControlToController_EmergencyStop struct {
	EmergencyStop bool `protobuf:"varint,5,opt,name=emergency_stop,json=emergencyStop,oneof"`
}

type RemoteControlToController_Timeout struct {
	Timeout bool `protobuf:"varint,6,opt,name=timeout,oneof"`
}

type RemoteControlToController_ChallengeFlag struct {
	ChallengeFlag bool `protobuf:"varint,7,opt,name=challenge_flag,json=challengeFlag,oneof"`
}

func (*RemoteControlToController_Ping) isRemoteControlToController_Msg() {}

func (*RemoteControlToController_DesiredKeeper) isRemoteControlToController_Msg() {}

func (*RemoteControlToController_SubstituteBot) isRemoteControlToController_Msg() {}

func (*RemoteControlToController_EmergencyStop) isRemoteControlToController_Msg() {}

func (*RemoteControlToController_Timeout) isRemoteControlToController_Msg() {}

func (*RemoteControlToController_ChallengeFlag) isRemoteControlToController_Msg() {}

func (m *RemoteControlToController) GetMsg() isRemoteControlToController_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *RemoteControlToController) GetPing() bool {
	if x, ok := m.GetMsg().(*RemoteControlToController_Ping); ok {
		return x.Ping
	}
	return false
}

func (m *RemoteControlToController) GetDesiredKeeper() int32 {
	if x, ok := m.GetMsg().(*RemoteControlToController_DesiredKeeper); ok {
		return x.DesiredKeeper
	}
	return 0
}

func (m *RemoteControlToController) GetSubstituteBot() bool {
	if x, ok := m.GetMsg().(*RemoteControlToController_SubstituteBot); ok {
		return x.SubstituteBot
	}
	return false
}

func (m *RemoteControlToController) GetEmergencyStop() bool {
	if x, ok := m.GetMsg().(*RemoteControlToController_EmergencyStop); ok {
		return x.EmergencyStop
	}
	return false
}

func (m *RemoteControlToController) GetTimeout() bool {
	if x, ok := m.GetMsg().(*RemoteControlToController_Timeout); ok {
		return x.Timeout
	}
	return false
}

func (m *RemoteControlToController) GetChallengeFlag() bool {
	if x, ok := m.GetMsg().(*RemoteControlToController_ChallengeFlag); ok {
		return x.ChallengeFlag
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RemoteControlToController) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RemoteControlToController_Ping)(nil),
		(*RemoteControlToController_DesiredKeeper)(nil),
		(*RemoteControlToController_SubstituteBot)(nil),
		(*RemoteControlToController_EmergencyStop)(nil),
		(*RemoteControlToController_Timeout)(nil),
		(*RemoteControlToController_ChallengeFlag)(nil),
	}
}

// wrapper for all messages from controller to a team's computer
type ControllerToRemoteControl struct {
	// a reply from the controller
	ControllerReply *ControllerReply `protobuf:"bytes,1,opt,name=controller_reply,json=controllerReply" json:"controller_reply,omitempty"`
	// currently set keeper id
	Keeper *int32 `protobuf:"varint,2,opt,name=keeper" json:"keeper,omitempty"`
	// true, if substitution request pending
	SubstituteBot *bool `protobuf:"varint,3,opt,name=substitute_bot,json=substituteBot" json:"substitute_bot,omitempty"`
	// true, if emergency stop pending
	EmergencyStop *bool `protobuf:"varint,4,opt,name=emergency_stop,json=emergencyStop" json:"emergency_stop,omitempty"`
	// true, if timeout request pending
	Timeout *bool `protobuf:"varint,5,opt,name=timeout" json:"timeout,omitempty"`
	// true, if challenge flag pending
	ChallengeFlag        *bool    `protobuf:"varint,6,opt,name=challenge_flag,json=challengeFlag" json:"challenge_flag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ControllerToRemoteControl) Reset()         { *m = ControllerToRemoteControl{} }
func (m *ControllerToRemoteControl) String() string { return proto.CompactTextString(m) }
func (*ControllerToRemoteControl) ProtoMessage()    {}
func (*ControllerToRemoteControl) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7030c3864b7b834, []int{2}
}

func (m *ControllerToRemoteControl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControllerToRemoteControl.Unmarshal(m, b)
}
func (m *ControllerToRemoteControl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControllerToRemoteControl.Marshal(b, m, deterministic)
}
func (m *ControllerToRemoteControl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControllerToRemoteControl.Merge(m, src)
}
func (m *ControllerToRemoteControl) XXX_Size() int {
	return xxx_messageInfo_ControllerToRemoteControl.Size(m)
}
func (m *ControllerToRemoteControl) XXX_DiscardUnknown() {
	xxx_messageInfo_ControllerToRemoteControl.DiscardUnknown(m)
}

var xxx_messageInfo_ControllerToRemoteControl proto.InternalMessageInfo

func (m *ControllerToRemoteControl) GetControllerReply() *ControllerReply {
	if m != nil {
		return m.ControllerReply
	}
	return nil
}

func (m *ControllerToRemoteControl) GetKeeper() int32 {
	if m != nil && m.Keeper != nil {
		return *m.Keeper
	}
	return 0
}

func (m *ControllerToRemoteControl) GetSubstituteBot() bool {
	if m != nil && m.SubstituteBot != nil {
		return *m.SubstituteBot
	}
	return false
}

func (m *ControllerToRemoteControl) GetEmergencyStop() bool {
	if m != nil && m.EmergencyStop != nil {
		return *m.EmergencyStop
	}
	return false
}

func (m *ControllerToRemoteControl) GetTimeout() bool {
	if m != nil && m.Timeout != nil {
		return *m.Timeout
	}
	return false
}

func (m *ControllerToRemoteControl) GetChallengeFlag() bool {
	if m != nil && m.ChallengeFlag != nil {
		return *m.ChallengeFlag
	}
	return false
}

func init() {
	proto.RegisterType((*RemoteControlRegistration)(nil), "RemoteControlRegistration")
	proto.RegisterType((*RemoteControlToController)(nil), "RemoteControlToController")
	proto.RegisterType((*ControllerToRemoteControl)(nil), "ControllerToRemoteControl")
}

func init() {
	proto.RegisterFile("ssl_gc_rcon_remotecontrol.proto", fileDescriptor_f7030c3864b7b834)
}

var fileDescriptor_f7030c3864b7b834 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x97, 0xae, 0xe9, 0x86, 0xd1, 0xc6, 0x08, 0x08, 0xa5, 0xbb, 0x50, 0x55, 0xaa, 0xe8,
	0xa5, 0x89, 0xb4, 0x2b, 0x88, 0xc3, 0x26, 0xa1, 0x49, 0x70, 0x72, 0x7b, 0xe2, 0x12, 0xdc, 0xec,
	0xe1, 0x59, 0xd8, 0x7e, 0x96, 0xfd, 0x72, 0xd8, 0xd7, 0xe1, 0x8b, 0x82, 0x92, 0xba, 0x0d, 0x59,
	0x25, 0x6e, 0xf1, 0xcf, 0xff, 0xf7, 0x7f, 0xca, 0x4f, 0x66, 0xef, 0x43, 0xd0, 0x95, 0xac, 0x2b,
	0x5f, 0xa3, 0xad, 0x3c, 0x18, 0x24, 0xa8, 0xd1, 0x92, 0x47, 0x5d, 0x38, 0x8f, 0x84, 0xd7, 0x6f,
	0x62, 0xa0, 0x46, 0x63, 0xd0, 0x46, 0xf8, 0xfa, 0x9f, 0xa9, 0x1d, 0x9a, 0xff, 0x60, 0x53, 0xde,
	0x8d, 0xdf, 0xed, 0xc6, 0x39, 0x48, 0x15, 0xc8, 0x0b, 0x52, 0x68, 0xb3, 0x29, 0x1b, 0x13, 0x08,
	0x93, 0x27, 0xb3, 0xd1, 0xf2, 0xf2, 0x26, 0x2d, 0x36, 0x20, 0x0c, 0xef, 0x50, 0xb6, 0x64, 0x2f,
	0x82, 0x92, 0x56, 0x50, 0xe3, 0x21, 0x1f, 0xcd, 0x92, 0xe5, 0xcb, 0x1b, 0x56, 0xac, 0xf7, 0x84,
	0xf7, 0x97, 0xf3, 0xdf, 0xa3, 0x67, 0x2b, 0x36, 0x18, 0x3f, 0x34, 0xf8, 0x61, 0x4f, 0xf2, 0x9f,
	0x9e, 0xec, 0x2d, 0x1b, 0x3b, 0x65, 0x65, 0xb7, 0xec, 0xfc, 0xfe, 0x84, 0x77, 0xa7, 0xec, 0x03,
	0xbb, 0x7c, 0x80, 0xa0, 0x3c, 0x3c, 0x54, 0xbf, 0x00, 0x1c, 0xf8, 0xfc, 0x74, 0x96, 0x2c, 0xd3,
	0xfb, 0x13, 0x7e, 0x11, 0xf9, 0xd7, 0x0e, 0xb7, 0xc1, 0xd0, 0x6c, 0x03, 0x29, 0x6a, 0x08, 0xaa,
	0x2d, 0x52, 0x3e, 0x8e, 0x45, 0x17, 0x3d, 0xbf, 0x45, 0x6a, 0x83, 0x60, 0xc0, 0x4b, 0xb0, 0xf5,
	0x53, 0x15, 0x08, 0x5d, 0x9e, 0xee, 0x83, 0x07, 0xbe, 0x26, 0x74, 0xd9, 0x35, 0x3b, 0x23, 0x65,
	0x00, 0x1b, 0xca, 0x27, 0x31, 0xb1, 0x07, 0x6d, 0x49, 0xfd, 0x28, 0xb4, 0x06, 0x2b, 0xa1, 0xfa,
	0xa9, 0x85, 0xcc, 0xcf, 0xf6, 0x25, 0x07, 0xfe, 0x45, 0x0b, 0x79, 0x9b, 0xb2, 0x53, 0x13, 0xe4,
	0xfc, 0x4f, 0xc2, 0xa6, 0xbd, 0x95, 0x0d, 0x0e, 0x84, 0x65, 0x1f, 0xd9, 0x55, 0x7d, 0xb8, 0xac,
	0x3c, 0x38, 0xfd, 0x14, 0x5d, 0x5d, 0x15, 0xfd, 0x14, 0x6f, 0x39, 0x7f, 0x55, 0x0f, 0x41, 0xf6,
	0x8e, 0x4d, 0xa2, 0x99, 0xd6, 0x5c, 0xca, 0xe3, 0x29, 0x5b, 0x1c, 0x09, 0x69, 0xcd, 0x9d, 0x3f,
	0xd7, 0xb1, 0x38, 0xd2, 0x31, 0xde, 0xc5, 0x86, 0x32, 0xf2, 0x5e, 0x46, 0xa7, 0xab, 0x57, 0xb1,
	0x38, 0x52, 0x31, 0xd9, 0x15, 0x0c, 0x45, 0x7c, 0xfe, 0xfe, 0x49, 0x2a, 0x7a, 0x6c, 0xb6, 0x45,
	0x8d, 0xa6, 0xe4, 0xb8, 0xc5, 0xbb, 0xc6, 0xad, 0xd6, 0xeb, 0x6f, 0x65, 0x08, 0x7a, 0x25, 0x85,
	0x81, 0x55, 0xff, 0x67, 0xa5, 0xb2, 0x04, 0xde, 0x0a, 0x5d, 0x0a, 0xe7, 0xca, 0xf6, 0x39, 0xff,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x1d, 0x22, 0x84, 0x12, 0x03, 0x00, 0x00,
}