// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wifi_configuration_specifics.proto

package sync_pb

import (
	fmt "fmt"
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

type WifiConfigurationSpecificsData_SecurityType int32

const (
	WifiConfigurationSpecificsData_SECURITY_TYPE_UNSPECIFIED WifiConfigurationSpecificsData_SecurityType = 0
	WifiConfigurationSpecificsData_SECURITY_TYPE_NONE        WifiConfigurationSpecificsData_SecurityType = 1
	WifiConfigurationSpecificsData_SECURITY_TYPE_WEP         WifiConfigurationSpecificsData_SecurityType = 2
	WifiConfigurationSpecificsData_SECURITY_TYPE_PSK         WifiConfigurationSpecificsData_SecurityType = 3
)

var WifiConfigurationSpecificsData_SecurityType_name = map[int32]string{
	0: "SECURITY_TYPE_UNSPECIFIED",
	1: "SECURITY_TYPE_NONE",
	2: "SECURITY_TYPE_WEP",
	3: "SECURITY_TYPE_PSK",
}

var WifiConfigurationSpecificsData_SecurityType_value = map[string]int32{
	"SECURITY_TYPE_UNSPECIFIED": 0,
	"SECURITY_TYPE_NONE":        1,
	"SECURITY_TYPE_WEP":         2,
	"SECURITY_TYPE_PSK":         3,
}

func (x WifiConfigurationSpecificsData_SecurityType) Enum() *WifiConfigurationSpecificsData_SecurityType {
	p := new(WifiConfigurationSpecificsData_SecurityType)
	*p = x
	return p
}

func (x WifiConfigurationSpecificsData_SecurityType) String() string {
	return proto.EnumName(WifiConfigurationSpecificsData_SecurityType_name, int32(x))
}

func (x *WifiConfigurationSpecificsData_SecurityType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(WifiConfigurationSpecificsData_SecurityType_value, data, "WifiConfigurationSpecificsData_SecurityType")
	if err != nil {
		return err
	}
	*x = WifiConfigurationSpecificsData_SecurityType(value)
	return nil
}

func (WifiConfigurationSpecificsData_SecurityType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_532bc6a80c7070ba, []int{0, 0}
}

type WifiConfigurationSpecificsData_AutomaticallyConnectOption int32

const (
	WifiConfigurationSpecificsData_AUTOMATICALLY_CONNECT_UNSPECIFIED WifiConfigurationSpecificsData_AutomaticallyConnectOption = 0
	WifiConfigurationSpecificsData_AUTOMATICALLY_CONNECT_DISABLED    WifiConfigurationSpecificsData_AutomaticallyConnectOption = 1
	WifiConfigurationSpecificsData_AUTOMATICALLY_CONNECT_ENABLED     WifiConfigurationSpecificsData_AutomaticallyConnectOption = 2
)

var WifiConfigurationSpecificsData_AutomaticallyConnectOption_name = map[int32]string{
	0: "AUTOMATICALLY_CONNECT_UNSPECIFIED",
	1: "AUTOMATICALLY_CONNECT_DISABLED",
	2: "AUTOMATICALLY_CONNECT_ENABLED",
}

var WifiConfigurationSpecificsData_AutomaticallyConnectOption_value = map[string]int32{
	"AUTOMATICALLY_CONNECT_UNSPECIFIED": 0,
	"AUTOMATICALLY_CONNECT_DISABLED":    1,
	"AUTOMATICALLY_CONNECT_ENABLED":     2,
}

func (x WifiConfigurationSpecificsData_AutomaticallyConnectOption) Enum() *WifiConfigurationSpecificsData_AutomaticallyConnectOption {
	p := new(WifiConfigurationSpecificsData_AutomaticallyConnectOption)
	*p = x
	return p
}

func (x WifiConfigurationSpecificsData_AutomaticallyConnectOption) String() string {
	return proto.EnumName(WifiConfigurationSpecificsData_AutomaticallyConnectOption_name, int32(x))
}

func (x *WifiConfigurationSpecificsData_AutomaticallyConnectOption) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(WifiConfigurationSpecificsData_AutomaticallyConnectOption_value, data, "WifiConfigurationSpecificsData_AutomaticallyConnectOption")
	if err != nil {
		return err
	}
	*x = WifiConfigurationSpecificsData_AutomaticallyConnectOption(value)
	return nil
}

func (WifiConfigurationSpecificsData_AutomaticallyConnectOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_532bc6a80c7070ba, []int{0, 1}
}

type WifiConfigurationSpecificsData_IsPreferredOption int32

const (
	WifiConfigurationSpecificsData_IS_PREFERRED_UNSPECIFIED WifiConfigurationSpecificsData_IsPreferredOption = 0
	WifiConfigurationSpecificsData_IS_PREFERRED_DISABLED    WifiConfigurationSpecificsData_IsPreferredOption = 1
	WifiConfigurationSpecificsData_IS_PREFERRED_ENABLED     WifiConfigurationSpecificsData_IsPreferredOption = 2
)

var WifiConfigurationSpecificsData_IsPreferredOption_name = map[int32]string{
	0: "IS_PREFERRED_UNSPECIFIED",
	1: "IS_PREFERRED_DISABLED",
	2: "IS_PREFERRED_ENABLED",
}

var WifiConfigurationSpecificsData_IsPreferredOption_value = map[string]int32{
	"IS_PREFERRED_UNSPECIFIED": 0,
	"IS_PREFERRED_DISABLED":    1,
	"IS_PREFERRED_ENABLED":     2,
}

func (x WifiConfigurationSpecificsData_IsPreferredOption) Enum() *WifiConfigurationSpecificsData_IsPreferredOption {
	p := new(WifiConfigurationSpecificsData_IsPreferredOption)
	*p = x
	return p
}

func (x WifiConfigurationSpecificsData_IsPreferredOption) String() string {
	return proto.EnumName(WifiConfigurationSpecificsData_IsPreferredOption_name, int32(x))
}

func (x *WifiConfigurationSpecificsData_IsPreferredOption) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(WifiConfigurationSpecificsData_IsPreferredOption_value, data, "WifiConfigurationSpecificsData_IsPreferredOption")
	if err != nil {
		return err
	}
	*x = WifiConfigurationSpecificsData_IsPreferredOption(value)
	return nil
}

func (WifiConfigurationSpecificsData_IsPreferredOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_532bc6a80c7070ba, []int{0, 2}
}

type WifiConfigurationSpecificsData_MeteredOption int32

const (
	WifiConfigurationSpecificsData_METERED_OPTION_UNSPECIFIED WifiConfigurationSpecificsData_MeteredOption = 0
	WifiConfigurationSpecificsData_METERED_OPTION_NO          WifiConfigurationSpecificsData_MeteredOption = 1
	WifiConfigurationSpecificsData_METERED_OPTION_YES         WifiConfigurationSpecificsData_MeteredOption = 2
	// Allows the device to use heuristics to determine if network is metered.
	WifiConfigurationSpecificsData_METERED_OPTION_AUTO WifiConfigurationSpecificsData_MeteredOption = 3
)

var WifiConfigurationSpecificsData_MeteredOption_name = map[int32]string{
	0: "METERED_OPTION_UNSPECIFIED",
	1: "METERED_OPTION_NO",
	2: "METERED_OPTION_YES",
	3: "METERED_OPTION_AUTO",
}

var WifiConfigurationSpecificsData_MeteredOption_value = map[string]int32{
	"METERED_OPTION_UNSPECIFIED": 0,
	"METERED_OPTION_NO":          1,
	"METERED_OPTION_YES":         2,
	"METERED_OPTION_AUTO":        3,
}

func (x WifiConfigurationSpecificsData_MeteredOption) Enum() *WifiConfigurationSpecificsData_MeteredOption {
	p := new(WifiConfigurationSpecificsData_MeteredOption)
	*p = x
	return p
}

func (x WifiConfigurationSpecificsData_MeteredOption) String() string {
	return proto.EnumName(WifiConfigurationSpecificsData_MeteredOption_name, int32(x))
}

func (x *WifiConfigurationSpecificsData_MeteredOption) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(WifiConfigurationSpecificsData_MeteredOption_value, data, "WifiConfigurationSpecificsData_MeteredOption")
	if err != nil {
		return err
	}
	*x = WifiConfigurationSpecificsData_MeteredOption(value)
	return nil
}

func (WifiConfigurationSpecificsData_MeteredOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_532bc6a80c7070ba, []int{0, 3}
}

type WifiConfigurationSpecificsData_ProxyConfiguration_ProxyOption int32

const (
	WifiConfigurationSpecificsData_ProxyConfiguration_PROXY_OPTION_UNSPECIFIED WifiConfigurationSpecificsData_ProxyConfiguration_ProxyOption = 0
	WifiConfigurationSpecificsData_ProxyConfiguration_PROXY_OPTION_DISABLED    WifiConfigurationSpecificsData_ProxyConfiguration_ProxyOption = 1
	// Use a Proxy Auto-config(PAC) Url, set in proxy_url
	WifiConfigurationSpecificsData_ProxyConfiguration_PROXY_OPTION_AUTOMATIC WifiConfigurationSpecificsData_ProxyConfiguration_ProxyOption = 2
	// Uses Web Proxy Auto-Discovery Protocol (WPAD) to discover the proxy
	// settings using DHCP/DNS.
	WifiConfigurationSpecificsData_ProxyConfiguration_PROXY_OPTION_AUTODISCOVERY WifiConfigurationSpecificsData_ProxyConfiguration_ProxyOption = 3
	// User sets proxy_url, proxy_port, and whitelisted_domains manually.
	WifiConfigurationSpecificsData_ProxyConfiguration_PROXY_OPTION_MANUAL WifiConfigurationSpecificsData_ProxyConfiguration_ProxyOption = 4
)

var WifiConfigurationSpecificsData_ProxyConfiguration_ProxyOption_name = map[int32]string{
	0: "PROXY_OPTION_UNSPECIFIED",
	1: "PROXY_OPTION_DISABLED",
	2: "PROXY_OPTION_AUTOMATIC",
	3: "PROXY_OPTION_AUTODISCOVERY",
	4: "PROXY_OPTION_MANUAL",
}

var WifiConfigurationSpecificsData_ProxyConfiguration_ProxyOption_value = map[string]int32{
	"PROXY_OPTION_UNSPECIFIED":   0,
	"PROXY_OPTION_DISABLED":      1,
	"PROXY_OPTION_AUTOMATIC":     2,
	"PROXY_OPTION_AUTODISCOVERY": 3,
	"PROXY_OPTION_MANUAL":        4,
}

func (x WifiConfigurationSpecificsData_ProxyConfiguration_ProxyOption) Enum() *WifiConfigurationSpecificsData_ProxyConfiguration_ProxyOption {
	p := new(WifiConfigurationSpecificsData_ProxyConfiguration_ProxyOption)
	*p = x
	return p
}

func (x WifiConfigurationSpecificsData_ProxyConfiguration_ProxyOption) String() string {
	return proto.EnumName(WifiConfigurationSpecificsData_ProxyConfiguration_ProxyOption_name, int32(x))
}

func (x *WifiConfigurationSpecificsData_ProxyConfiguration_ProxyOption) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(WifiConfigurationSpecificsData_ProxyConfiguration_ProxyOption_value, data, "WifiConfigurationSpecificsData_ProxyConfiguration_ProxyOption")
	if err != nil {
		return err
	}
	*x = WifiConfigurationSpecificsData_ProxyConfiguration_ProxyOption(value)
	return nil
}

func (WifiConfigurationSpecificsData_ProxyConfiguration_ProxyOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_532bc6a80c7070ba, []int{0, 0, 0}
}

type WifiConfigurationSpecificsData struct {
	// SSID encoded to hex, letters should be upper case and 0x prefix should be
	// omitted. For example, ssid "network" would be provided as "6E6574776F726B".
	HexSsid      []byte                                       `protobuf:"bytes,1,opt,name=hex_ssid,json=hexSsid" json:"hex_ssid,omitempty"`
	SecurityType *WifiConfigurationSpecificsData_SecurityType `protobuf:"varint,2,opt,name=security_type,json=securityType,enum=sync_pb.WifiConfigurationSpecificsData_SecurityType" json:"security_type,omitempty"`
	// The passphrase can be ASCII, UTF-8, or a string of hex digits.
	Passphrase           []byte                                                     `protobuf:"bytes,3,opt,name=passphrase" json:"passphrase,omitempty"`
	AutomaticallyConnect *WifiConfigurationSpecificsData_AutomaticallyConnectOption `protobuf:"varint,4,opt,name=automatically_connect,json=automaticallyConnect,enum=sync_pb.WifiConfigurationSpecificsData_AutomaticallyConnectOption" json:"automatically_connect,omitempty"`
	IsPreferred          *WifiConfigurationSpecificsData_IsPreferredOption          `protobuf:"varint,5,opt,name=is_preferred,json=isPreferred,enum=sync_pb.WifiConfigurationSpecificsData_IsPreferredOption" json:"is_preferred,omitempty"`
	Metered              *WifiConfigurationSpecificsData_MeteredOption              `protobuf:"varint,6,opt,name=metered,enum=sync_pb.WifiConfigurationSpecificsData_MeteredOption" json:"metered,omitempty"`
	ProxyConfiguration   *WifiConfigurationSpecificsData_ProxyConfiguration         `protobuf:"bytes,7,opt,name=proxy_configuration,json=proxyConfiguration" json:"proxy_configuration,omitempty"`
	// List of DNS servers to be used.  Up to 4.
	CustomDns []string `protobuf:"bytes,8,rep,name=custom_dns,json=customDns" json:"custom_dns,omitempty"`
	// This is represented by the UNIX timestamp, ms since epoch.
	LastUpdateTimestamp  *int64   `protobuf:"varint,9,opt,name=last_update_timestamp,json=lastUpdateTimestamp" json:"last_update_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WifiConfigurationSpecificsData) Reset()         { *m = WifiConfigurationSpecificsData{} }
func (m *WifiConfigurationSpecificsData) String() string { return proto.CompactTextString(m) }
func (*WifiConfigurationSpecificsData) ProtoMessage()    {}
func (*WifiConfigurationSpecificsData) Descriptor() ([]byte, []int) {
	return fileDescriptor_532bc6a80c7070ba, []int{0}
}

func (m *WifiConfigurationSpecificsData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WifiConfigurationSpecificsData.Unmarshal(m, b)
}
func (m *WifiConfigurationSpecificsData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WifiConfigurationSpecificsData.Marshal(b, m, deterministic)
}
func (m *WifiConfigurationSpecificsData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WifiConfigurationSpecificsData.Merge(m, src)
}
func (m *WifiConfigurationSpecificsData) XXX_Size() int {
	return xxx_messageInfo_WifiConfigurationSpecificsData.Size(m)
}
func (m *WifiConfigurationSpecificsData) XXX_DiscardUnknown() {
	xxx_messageInfo_WifiConfigurationSpecificsData.DiscardUnknown(m)
}

var xxx_messageInfo_WifiConfigurationSpecificsData proto.InternalMessageInfo

func (m *WifiConfigurationSpecificsData) GetHexSsid() []byte {
	if m != nil {
		return m.HexSsid
	}
	return nil
}

func (m *WifiConfigurationSpecificsData) GetSecurityType() WifiConfigurationSpecificsData_SecurityType {
	if m != nil && m.SecurityType != nil {
		return *m.SecurityType
	}
	return WifiConfigurationSpecificsData_SECURITY_TYPE_UNSPECIFIED
}

func (m *WifiConfigurationSpecificsData) GetPassphrase() []byte {
	if m != nil {
		return m.Passphrase
	}
	return nil
}

func (m *WifiConfigurationSpecificsData) GetAutomaticallyConnect() WifiConfigurationSpecificsData_AutomaticallyConnectOption {
	if m != nil && m.AutomaticallyConnect != nil {
		return *m.AutomaticallyConnect
	}
	return WifiConfigurationSpecificsData_AUTOMATICALLY_CONNECT_UNSPECIFIED
}

func (m *WifiConfigurationSpecificsData) GetIsPreferred() WifiConfigurationSpecificsData_IsPreferredOption {
	if m != nil && m.IsPreferred != nil {
		return *m.IsPreferred
	}
	return WifiConfigurationSpecificsData_IS_PREFERRED_UNSPECIFIED
}

func (m *WifiConfigurationSpecificsData) GetMetered() WifiConfigurationSpecificsData_MeteredOption {
	if m != nil && m.Metered != nil {
		return *m.Metered
	}
	return WifiConfigurationSpecificsData_METERED_OPTION_UNSPECIFIED
}

func (m *WifiConfigurationSpecificsData) GetProxyConfiguration() *WifiConfigurationSpecificsData_ProxyConfiguration {
	if m != nil {
		return m.ProxyConfiguration
	}
	return nil
}

func (m *WifiConfigurationSpecificsData) GetCustomDns() []string {
	if m != nil {
		return m.CustomDns
	}
	return nil
}

func (m *WifiConfigurationSpecificsData) GetLastUpdateTimestamp() int64 {
	if m != nil && m.LastUpdateTimestamp != nil {
		return *m.LastUpdateTimestamp
	}
	return 0
}

type WifiConfigurationSpecificsData_ProxyConfiguration struct {
	ProxyOption *WifiConfigurationSpecificsData_ProxyConfiguration_ProxyOption `protobuf:"varint,1,opt,name=proxy_option,json=proxyOption,enum=sync_pb.WifiConfigurationSpecificsData_ProxyConfiguration_ProxyOption" json:"proxy_option,omitempty"`
	// Only set if PROXY_OPTION_AUTOMATIC or PROXY_OPTION_MANUAL.
	ProxyUrl *string `protobuf:"bytes,2,opt,name=proxy_url,json=proxyUrl" json:"proxy_url,omitempty"`
	// Only set if PROXY_OPTION_MANUAL.
	ProxyPort *int32 `protobuf:"varint,3,opt,name=proxy_port,json=proxyPort" json:"proxy_port,omitempty"`
	// Only set if PROXY_OPTION_MANUAL.
	WhitelistedDomains   []string `protobuf:"bytes,4,rep,name=whitelisted_domains,json=whitelistedDomains" json:"whitelisted_domains,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WifiConfigurationSpecificsData_ProxyConfiguration) Reset() {
	*m = WifiConfigurationSpecificsData_ProxyConfiguration{}
}
func (m *WifiConfigurationSpecificsData_ProxyConfiguration) String() string {
	return proto.CompactTextString(m)
}
func (*WifiConfigurationSpecificsData_ProxyConfiguration) ProtoMessage() {}
func (*WifiConfigurationSpecificsData_ProxyConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_532bc6a80c7070ba, []int{0, 0}
}

func (m *WifiConfigurationSpecificsData_ProxyConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WifiConfigurationSpecificsData_ProxyConfiguration.Unmarshal(m, b)
}
func (m *WifiConfigurationSpecificsData_ProxyConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WifiConfigurationSpecificsData_ProxyConfiguration.Marshal(b, m, deterministic)
}
func (m *WifiConfigurationSpecificsData_ProxyConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WifiConfigurationSpecificsData_ProxyConfiguration.Merge(m, src)
}
func (m *WifiConfigurationSpecificsData_ProxyConfiguration) XXX_Size() int {
	return xxx_messageInfo_WifiConfigurationSpecificsData_ProxyConfiguration.Size(m)
}
func (m *WifiConfigurationSpecificsData_ProxyConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_WifiConfigurationSpecificsData_ProxyConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_WifiConfigurationSpecificsData_ProxyConfiguration proto.InternalMessageInfo

func (m *WifiConfigurationSpecificsData_ProxyConfiguration) GetProxyOption() WifiConfigurationSpecificsData_ProxyConfiguration_ProxyOption {
	if m != nil && m.ProxyOption != nil {
		return *m.ProxyOption
	}
	return WifiConfigurationSpecificsData_ProxyConfiguration_PROXY_OPTION_UNSPECIFIED
}

func (m *WifiConfigurationSpecificsData_ProxyConfiguration) GetProxyUrl() string {
	if m != nil && m.ProxyUrl != nil {
		return *m.ProxyUrl
	}
	return ""
}

func (m *WifiConfigurationSpecificsData_ProxyConfiguration) GetProxyPort() int32 {
	if m != nil && m.ProxyPort != nil {
		return *m.ProxyPort
	}
	return 0
}

func (m *WifiConfigurationSpecificsData_ProxyConfiguration) GetWhitelistedDomains() []string {
	if m != nil {
		return m.WhitelistedDomains
	}
	return nil
}

type WifiConfigurationSpecifics struct {
	// The actual wifi configuration data. Contains an encrypted
	// WifiConfigurationSpecificsData message.
	Encrypted *EncryptedData `protobuf:"bytes,1,opt,name=encrypted" json:"encrypted,omitempty"`
	// An unsynced field for use internally on the client. This field should
	// never be set in any network-based communications because it contains
	// unencrypted material.
	ClientOnlyEncryptedData *WifiConfigurationSpecificsData `protobuf:"bytes,2,opt,name=client_only_encrypted_data,json=clientOnlyEncryptedData" json:"client_only_encrypted_data,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                        `json:"-"`
	XXX_unrecognized        []byte                          `json:"-"`
	XXX_sizecache           int32                           `json:"-"`
}

func (m *WifiConfigurationSpecifics) Reset()         { *m = WifiConfigurationSpecifics{} }
func (m *WifiConfigurationSpecifics) String() string { return proto.CompactTextString(m) }
func (*WifiConfigurationSpecifics) ProtoMessage()    {}
func (*WifiConfigurationSpecifics) Descriptor() ([]byte, []int) {
	return fileDescriptor_532bc6a80c7070ba, []int{1}
}

func (m *WifiConfigurationSpecifics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WifiConfigurationSpecifics.Unmarshal(m, b)
}
func (m *WifiConfigurationSpecifics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WifiConfigurationSpecifics.Marshal(b, m, deterministic)
}
func (m *WifiConfigurationSpecifics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WifiConfigurationSpecifics.Merge(m, src)
}
func (m *WifiConfigurationSpecifics) XXX_Size() int {
	return xxx_messageInfo_WifiConfigurationSpecifics.Size(m)
}
func (m *WifiConfigurationSpecifics) XXX_DiscardUnknown() {
	xxx_messageInfo_WifiConfigurationSpecifics.DiscardUnknown(m)
}

var xxx_messageInfo_WifiConfigurationSpecifics proto.InternalMessageInfo

func (m *WifiConfigurationSpecifics) GetEncrypted() *EncryptedData {
	if m != nil {
		return m.Encrypted
	}
	return nil
}

func (m *WifiConfigurationSpecifics) GetClientOnlyEncryptedData() *WifiConfigurationSpecificsData {
	if m != nil {
		return m.ClientOnlyEncryptedData
	}
	return nil
}

func init() {
	proto.RegisterEnum("sync_pb.WifiConfigurationSpecificsData_SecurityType", WifiConfigurationSpecificsData_SecurityType_name, WifiConfigurationSpecificsData_SecurityType_value)
	proto.RegisterEnum("sync_pb.WifiConfigurationSpecificsData_AutomaticallyConnectOption", WifiConfigurationSpecificsData_AutomaticallyConnectOption_name, WifiConfigurationSpecificsData_AutomaticallyConnectOption_value)
	proto.RegisterEnum("sync_pb.WifiConfigurationSpecificsData_IsPreferredOption", WifiConfigurationSpecificsData_IsPreferredOption_name, WifiConfigurationSpecificsData_IsPreferredOption_value)
	proto.RegisterEnum("sync_pb.WifiConfigurationSpecificsData_MeteredOption", WifiConfigurationSpecificsData_MeteredOption_name, WifiConfigurationSpecificsData_MeteredOption_value)
	proto.RegisterEnum("sync_pb.WifiConfigurationSpecificsData_ProxyConfiguration_ProxyOption", WifiConfigurationSpecificsData_ProxyConfiguration_ProxyOption_name, WifiConfigurationSpecificsData_ProxyConfiguration_ProxyOption_value)
	proto.RegisterType((*WifiConfigurationSpecificsData)(nil), "sync_pb.WifiConfigurationSpecificsData")
	proto.RegisterType((*WifiConfigurationSpecificsData_ProxyConfiguration)(nil), "sync_pb.WifiConfigurationSpecificsData.ProxyConfiguration")
	proto.RegisterType((*WifiConfigurationSpecifics)(nil), "sync_pb.WifiConfigurationSpecifics")
}

func init() {
	proto.RegisterFile("wifi_configuration_specifics.proto", fileDescriptor_532bc6a80c7070ba)
}

var fileDescriptor_532bc6a80c7070ba = []byte{
	// 803 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xd1, 0x6a, 0xe3, 0x46,
	0x14, 0xad, 0xec, 0x6c, 0x13, 0x5f, 0x7b, 0x17, 0x65, 0xb2, 0xce, 0x2a, 0x6e, 0x13, 0xbc, 0x86,
	0xa5, 0x86, 0x82, 0x0b, 0x66, 0xfb, 0xd0, 0xbe, 0x39, 0xf6, 0x84, 0x8a, 0x26, 0x92, 0x90, 0xe4,
	0x6e, 0x0d, 0x85, 0x41, 0x95, 0xc6, 0xeb, 0xa1, 0x92, 0x46, 0x68, 0xc6, 0x24, 0xfa, 0x85, 0x7d,
	0xe8, 0x0f, 0xf4, 0x47, 0xfa, 0x79, 0x45, 0x92, 0x9d, 0x48, 0x76, 0xb6, 0x98, 0x7d, 0x9c, 0x73,
	0xee, 0x9c, 0x73, 0xef, 0xdc, 0x7b, 0x19, 0x18, 0xdc, 0xb3, 0x25, 0x23, 0x3e, 0x8f, 0x97, 0xec,
	0xe3, 0x3a, 0xf5, 0x24, 0xe3, 0x31, 0x11, 0x09, 0xf5, 0xd9, 0x92, 0xf9, 0x62, 0x94, 0xa4, 0x5c,
	0x72, 0x74, 0x2c, 0xb2, 0xd8, 0x27, 0xc9, 0x9f, 0x3d, 0x95, 0xc6, 0x7e, 0x9a, 0x25, 0x79, 0x50,
	0x49, 0x0d, 0xfe, 0xee, 0xc0, 0xd5, 0x07, 0xb6, 0x64, 0xd3, 0xaa, 0x80, 0xb3, 0xbd, 0x3f, 0xf3,
	0xa4, 0x87, 0x2e, 0xe0, 0x64, 0x45, 0x1f, 0x88, 0x10, 0x2c, 0xd0, 0x94, 0xbe, 0x32, 0xec, 0xd8,
	0xc7, 0x2b, 0xfa, 0xe0, 0x08, 0x16, 0xa0, 0x05, 0xbc, 0x14, 0xd4, 0x5f, 0xa7, 0x4c, 0x66, 0x44,
	0x66, 0x09, 0xd5, 0x1a, 0x7d, 0x65, 0xf8, 0x6a, 0xfc, 0x7e, 0xb4, 0x31, 0x1c, 0xfd, 0xbf, 0xf4,
	0xc8, 0xd9, 0x5c, 0x76, 0xb3, 0x84, 0xda, 0x1d, 0x51, 0x39, 0xa1, 0x2b, 0x80, 0xc4, 0x13, 0x22,
	0x59, 0xa5, 0x9e, 0xa0, 0x5a, 0xb3, 0xf0, 0xad, 0x20, 0xe8, 0x1e, 0xba, 0xde, 0x5a, 0xf2, 0xc8,
	0x93, 0xcc, 0xf7, 0xc2, 0x30, 0xcb, 0x9f, 0x20, 0xa6, 0xbe, 0xd4, 0x8e, 0x8a, 0x14, 0xae, 0x0f,
	0x4d, 0x61, 0x52, 0x15, 0x99, 0x96, 0x1a, 0x66, 0xf1, 0x42, 0xf6, 0x6b, 0xef, 0x19, 0x0e, 0xfd,
	0x01, 0x1d, 0x26, 0x48, 0x92, 0xd2, 0x25, 0x4d, 0x53, 0x1a, 0x68, 0x2f, 0x0a, 0xbf, 0x9f, 0x0e,
	0xf5, 0xd3, 0x85, 0xb5, 0xbd, 0xba, 0xb1, 0x69, 0xb3, 0x27, 0x08, 0x99, 0x70, 0x1c, 0x51, 0x49,
	0x73, 0xe1, 0xaf, 0x0b, 0xe1, 0x1f, 0x0f, 0x15, 0xbe, 0x2b, 0xaf, 0x6d, 0x44, 0xb7, 0x2a, 0xe8,
	0x2f, 0x38, 0x4b, 0x52, 0xfe, 0x90, 0xd5, 0x47, 0x44, 0x3b, 0xee, 0x2b, 0xc3, 0xf6, 0xf8, 0xe7,
	0x43, 0xc5, 0xad, 0x5c, 0xa2, 0xc6, 0xdb, 0x28, 0xd9, 0xc3, 0xd0, 0x25, 0x80, 0xbf, 0x16, 0x92,
	0x47, 0x24, 0x88, 0x85, 0x76, 0xd2, 0x6f, 0x0e, 0x5b, 0x76, 0xab, 0x44, 0x66, 0xb1, 0x40, 0x63,
	0xe8, 0x86, 0x9e, 0x90, 0x64, 0x9d, 0x04, 0x9e, 0xa4, 0x44, 0xb2, 0x88, 0x0a, 0xe9, 0x45, 0x89,
	0xd6, 0xea, 0x2b, 0xc3, 0xa6, 0x7d, 0x96, 0x93, 0xf3, 0x82, 0x73, 0xb7, 0x54, 0xef, 0x53, 0x13,
	0xd0, 0xbe, 0x3b, 0x62, 0xd0, 0x29, 0xcb, 0xe2, 0x45, 0xbd, 0xc5, 0x60, 0xbe, 0x1a, 0xdf, 0x7c,
	0x79, 0x3d, 0x25, 0xb4, 0x6d, 0x49, 0xf2, 0x74, 0x40, 0xdf, 0x40, 0xab, 0xb4, 0x5a, 0xa7, 0x61,
	0x31, 0xe0, 0x2d, 0xfb, 0xa4, 0x00, 0xe6, 0x69, 0x98, 0x57, 0x5c, 0x92, 0x09, 0x4f, 0x65, 0x31,
	0xa6, 0x2f, 0xec, 0x32, 0xdc, 0xe2, 0xa9, 0x44, 0x3f, 0xc0, 0xd9, 0xfd, 0x8a, 0x49, 0x1a, 0x32,
	0x21, 0x69, 0x40, 0x02, 0x1e, 0x79, 0x2c, 0x16, 0xda, 0x51, 0xf1, 0x32, 0xa8, 0x42, 0xcd, 0x4a,
	0x66, 0xf0, 0x8f, 0x02, 0xed, 0x4a, 0x26, 0xe8, 0x5b, 0xd0, 0x2c, 0xdb, 0xfc, 0x7d, 0x41, 0x4c,
	0xcb, 0xd5, 0x4d, 0x83, 0xcc, 0x0d, 0xc7, 0xc2, 0x53, 0xfd, 0x46, 0xc7, 0x33, 0xf5, 0x2b, 0x74,
	0x01, 0xdd, 0x1a, 0x3b, 0xd3, 0x9d, 0xc9, 0xf5, 0x2d, 0x9e, 0xa9, 0x0a, 0xea, 0xc1, 0x79, 0x8d,
	0x9a, 0xcc, 0x5d, 0xf3, 0x6e, 0xe2, 0xea, 0x53, 0xb5, 0x81, 0xae, 0xa0, 0xb7, 0xc7, 0xcd, 0x74,
	0x67, 0x6a, 0xfe, 0x86, 0xed, 0x85, 0xda, 0x44, 0x6f, 0xe0, 0xac, 0xc6, 0xdf, 0x4d, 0x8c, 0xf9,
	0xe4, 0x56, 0x3d, 0x1a, 0x08, 0xe8, 0x54, 0x57, 0x16, 0x5d, 0xc2, 0x85, 0x83, 0xa7, 0x73, 0x5b,
	0x77, 0x17, 0xc4, 0x5d, 0x58, 0x78, 0x27, 0xbd, 0x73, 0x40, 0x75, 0xda, 0x30, 0x0d, 0xac, 0x2a,
	0xa8, 0x0b, 0xa7, 0x75, 0xfc, 0x03, 0xb6, 0xd4, 0xc6, 0x3e, 0x6c, 0x39, 0xbf, 0xaa, 0xcd, 0xc1,
	0x27, 0x05, 0x7a, 0x9f, 0xdf, 0x52, 0xf4, 0x0e, 0xde, 0x3e, 0xd6, 0x36, 0xb9, 0xbd, 0x5d, 0x90,
	0xa9, 0x69, 0x18, 0x78, 0xea, 0xee, 0xe4, 0x32, 0x80, 0xab, 0xe7, 0xc3, 0x2a, 0x6f, 0xf6, 0x16,
	0x2e, 0x9f, 0x8f, 0xc1, 0x46, 0x19, 0xd2, 0x18, 0x2c, 0xe1, 0x74, 0x6f, 0x83, 0xf3, 0x26, 0xe9,
	0x0e, 0xb1, 0x6c, 0x7c, 0x83, 0x6d, 0x1b, 0xcf, 0xf6, 0x9b, 0x54, 0x63, 0x2b, 0x86, 0x1a, 0xbc,
	0xae, 0x51, 0x4f, 0x3e, 0xf7, 0xf0, 0xb2, 0xb6, 0xd0, 0x79, 0xcf, 0xee, 0xb0, 0x8b, 0xf3, 0xa8,
	0x67, 0x47, 0xa1, 0x0b, 0xa7, 0x3b, 0xbc, 0x61, 0xaa, 0x4a, 0xde, 0x82, 0x1d, 0x78, 0x81, 0x1d,
	0xb5, 0x91, 0xb7, 0x78, 0x07, 0xcf, 0x2b, 0x57, 0x9b, 0x83, 0x7f, 0x15, 0xe8, 0x7d, 0x7e, 0x79,
	0xd0, 0x7b, 0x68, 0x6d, 0xfe, 0x10, 0x5a, 0xfe, 0x06, 0xed, 0xf1, 0xf9, 0xe3, 0xd2, 0xe1, 0x2d,
	0x93, 0xef, 0x98, 0xfd, 0x14, 0x88, 0x02, 0xe8, 0xf9, 0x21, 0xa3, 0xb1, 0x24, 0x3c, 0x0e, 0x33,
	0xf2, 0x48, 0x90, 0xc0, 0x93, 0x5e, 0xb1, 0x53, 0xed, 0xf1, 0x77, 0x07, 0xee, 0xae, 0xfd, 0xa6,
	0x94, 0x32, 0xe3, 0x30, 0xab, 0x19, 0x5e, 0x7f, 0x0f, 0xef, 0x78, 0xfa, 0x71, 0xe4, 0xaf, 0x52,
	0x1e, 0xb1, 0x75, 0x34, 0xf2, 0x79, 0x94, 0xf0, 0x98, 0xc6, 0x52, 0x14, 0xd2, 0xe5, 0x8f, 0xe7,
	0xf3, 0xf0, 0x97, 0xa6, 0xa5, 0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0x34, 0x52, 0x89, 0xb9, 0x38,
	0x07, 0x00, 0x00,
}
