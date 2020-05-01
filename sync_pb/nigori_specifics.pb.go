// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nigori_specifics.proto

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

// The state of the passphrase required to decrypt |encryption_keybag|.
type NigoriSpecifics_PassphraseType int32

const (
	// Proto enums should begin with an 'unknown' value for switch statements,
	// unknown fields, etc.
	NigoriSpecifics_UNKNOWN NigoriSpecifics_PassphraseType = 0
	// Gaia-based encryption passphrase. Deprecated.
	NigoriSpecifics_IMPLICIT_PASSPHRASE NigoriSpecifics_PassphraseType = 1
	// Keystore key encryption passphrase. Uses |keystore_bootstrap| to
	// decrypt |encryption_keybag|.
	NigoriSpecifics_KEYSTORE_PASSPHRASE NigoriSpecifics_PassphraseType = 2
	// Previous Gaia-based passphrase frozen and treated as a custom passphrase.
	NigoriSpecifics_FROZEN_IMPLICIT_PASSPHRASE NigoriSpecifics_PassphraseType = 3
	// User provided custom passphrase.
	NigoriSpecifics_CUSTOM_PASSPHRASE NigoriSpecifics_PassphraseType = 4
	// Encryption key provided by a trusted vault.
	NigoriSpecifics_TRUSTED_VAULT_PASSPHRASE NigoriSpecifics_PassphraseType = 5
)

var NigoriSpecifics_PassphraseType_name = map[int32]string{
	0: "UNKNOWN",
	1: "IMPLICIT_PASSPHRASE",
	2: "KEYSTORE_PASSPHRASE",
	3: "FROZEN_IMPLICIT_PASSPHRASE",
	4: "CUSTOM_PASSPHRASE",
	5: "TRUSTED_VAULT_PASSPHRASE",
}

var NigoriSpecifics_PassphraseType_value = map[string]int32{
	"UNKNOWN":                    0,
	"IMPLICIT_PASSPHRASE":        1,
	"KEYSTORE_PASSPHRASE":        2,
	"FROZEN_IMPLICIT_PASSPHRASE": 3,
	"CUSTOM_PASSPHRASE":          4,
	"TRUSTED_VAULT_PASSPHRASE":   5,
}

func (x NigoriSpecifics_PassphraseType) Enum() *NigoriSpecifics_PassphraseType {
	p := new(NigoriSpecifics_PassphraseType)
	*p = x
	return p
}

func (x NigoriSpecifics_PassphraseType) String() string {
	return proto.EnumName(NigoriSpecifics_PassphraseType_name, int32(x))
}

func (x *NigoriSpecifics_PassphraseType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NigoriSpecifics_PassphraseType_value, data, "NigoriSpecifics_PassphraseType")
	if err != nil {
		return err
	}
	*x = NigoriSpecifics_PassphraseType(value)
	return nil
}

func (NigoriSpecifics_PassphraseType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_54e3415a76ee5ff6, []int{2, 0}
}

type NigoriSpecifics_KeyDerivationMethod int32

const (
	// This comes from a <= M69 client, who does not know about the field
	// (but implicitly uses PBKDF2_HMAC_SHA1_1003).
	NigoriSpecifics_UNSPECIFIED NigoriSpecifics_KeyDerivationMethod = 0
	// PBKDF2-HMAC-SHA1 with 1003 iterations and constant hardcoded salt. Was
	// implicitly used in <= M69.
	NigoriSpecifics_PBKDF2_HMAC_SHA1_1003 NigoriSpecifics_KeyDerivationMethod = 1
	// scrypt with N = 2^13, r = 8, p = 11 and random salt. Was added in M70.
	NigoriSpecifics_SCRYPT_8192_8_11 NigoriSpecifics_KeyDerivationMethod = 2
)

var NigoriSpecifics_KeyDerivationMethod_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "PBKDF2_HMAC_SHA1_1003",
	2: "SCRYPT_8192_8_11",
}

var NigoriSpecifics_KeyDerivationMethod_value = map[string]int32{
	"UNSPECIFIED":           0,
	"PBKDF2_HMAC_SHA1_1003": 1,
	"SCRYPT_8192_8_11":      2,
}

func (x NigoriSpecifics_KeyDerivationMethod) Enum() *NigoriSpecifics_KeyDerivationMethod {
	p := new(NigoriSpecifics_KeyDerivationMethod)
	*p = x
	return p
}

func (x NigoriSpecifics_KeyDerivationMethod) String() string {
	return proto.EnumName(NigoriSpecifics_KeyDerivationMethod_name, int32(x))
}

func (x *NigoriSpecifics_KeyDerivationMethod) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NigoriSpecifics_KeyDerivationMethod_value, data, "NigoriSpecifics_KeyDerivationMethod")
	if err != nil {
		return err
	}
	*x = NigoriSpecifics_KeyDerivationMethod(value)
	return nil
}

func (NigoriSpecifics_KeyDerivationMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_54e3415a76ee5ff6, []int{2, 1}
}

type NigoriKey struct {
	// Note that M78 and before rely on the name being populated, at least for
	// the main encrypted keybag within NigoriSpecifics.
	DeprecatedName       *string  `protobuf:"bytes,1,opt,name=deprecated_name,json=deprecatedName" json:"deprecated_name,omitempty"`            // Deprecated: Do not use.
	DeprecatedUserKey    []byte   `protobuf:"bytes,2,opt,name=deprecated_user_key,json=deprecatedUserKey" json:"deprecated_user_key,omitempty"` // Deprecated: Do not use.
	EncryptionKey        []byte   `protobuf:"bytes,3,opt,name=encryption_key,json=encryptionKey" json:"encryption_key,omitempty"`
	MacKey               []byte   `protobuf:"bytes,4,opt,name=mac_key,json=macKey" json:"mac_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NigoriKey) Reset()         { *m = NigoriKey{} }
func (m *NigoriKey) String() string { return proto.CompactTextString(m) }
func (*NigoriKey) ProtoMessage()    {}
func (*NigoriKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_54e3415a76ee5ff6, []int{0}
}

func (m *NigoriKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NigoriKey.Unmarshal(m, b)
}
func (m *NigoriKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NigoriKey.Marshal(b, m, deterministic)
}
func (m *NigoriKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NigoriKey.Merge(m, src)
}
func (m *NigoriKey) XXX_Size() int {
	return xxx_messageInfo_NigoriKey.Size(m)
}
func (m *NigoriKey) XXX_DiscardUnknown() {
	xxx_messageInfo_NigoriKey.DiscardUnknown(m)
}

var xxx_messageInfo_NigoriKey proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *NigoriKey) GetDeprecatedName() string {
	if m != nil && m.DeprecatedName != nil {
		return *m.DeprecatedName
	}
	return ""
}

// Deprecated: Do not use.
func (m *NigoriKey) GetDeprecatedUserKey() []byte {
	if m != nil {
		return m.DeprecatedUserKey
	}
	return nil
}

func (m *NigoriKey) GetEncryptionKey() []byte {
	if m != nil {
		return m.EncryptionKey
	}
	return nil
}

func (m *NigoriKey) GetMacKey() []byte {
	if m != nil {
		return m.MacKey
	}
	return nil
}

type NigoriKeyBag struct {
	Key                  []*NigoriKey `protobuf:"bytes,2,rep,name=key" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *NigoriKeyBag) Reset()         { *m = NigoriKeyBag{} }
func (m *NigoriKeyBag) String() string { return proto.CompactTextString(m) }
func (*NigoriKeyBag) ProtoMessage()    {}
func (*NigoriKeyBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_54e3415a76ee5ff6, []int{1}
}

func (m *NigoriKeyBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NigoriKeyBag.Unmarshal(m, b)
}
func (m *NigoriKeyBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NigoriKeyBag.Marshal(b, m, deterministic)
}
func (m *NigoriKeyBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NigoriKeyBag.Merge(m, src)
}
func (m *NigoriKeyBag) XXX_Size() int {
	return xxx_messageInfo_NigoriKeyBag.Size(m)
}
func (m *NigoriKeyBag) XXX_DiscardUnknown() {
	xxx_messageInfo_NigoriKeyBag.DiscardUnknown(m)
}

var xxx_messageInfo_NigoriKeyBag proto.InternalMessageInfo

func (m *NigoriKeyBag) GetKey() []*NigoriKey {
	if m != nil {
		return m.Key
	}
	return nil
}

// Properties of nigori sync object.
type NigoriSpecifics struct {
	EncryptionKeybag *EncryptedData `protobuf:"bytes,1,opt,name=encryption_keybag,json=encryptionKeybag" json:"encryption_keybag,omitempty"`
	// Once keystore migration is performed, we have to freeze the keybag so that
	// older clients (that don't support keystore encryption) do not attempt to
	// update the keybag.
	// Previously |using_explicit_passphrase|.
	KeybagIsFrozen *bool `protobuf:"varint,2,opt,name=keybag_is_frozen,json=keybagIsFrozen" json:"keybag_is_frozen,omitempty"`
	// Booleans corresponding to whether a datatype should be encrypted.
	// Passwords and Wi-Fi configurations are always encrypted, so we don't need
	// a field here.  History delete directives need to be consumable by the
	// server, and thus can't be encrypted.
	EncryptBookmarks       *bool `protobuf:"varint,13,opt,name=encrypt_bookmarks,json=encryptBookmarks" json:"encrypt_bookmarks,omitempty"`
	EncryptPreferences     *bool `protobuf:"varint,14,opt,name=encrypt_preferences,json=encryptPreferences" json:"encrypt_preferences,omitempty"`
	EncryptAutofillProfile *bool `protobuf:"varint,15,opt,name=encrypt_autofill_profile,json=encryptAutofillProfile" json:"encrypt_autofill_profile,omitempty"`
	EncryptAutofill        *bool `protobuf:"varint,16,opt,name=encrypt_autofill,json=encryptAutofill" json:"encrypt_autofill,omitempty"`
	EncryptThemes          *bool `protobuf:"varint,17,opt,name=encrypt_themes,json=encryptThemes" json:"encrypt_themes,omitempty"`
	EncryptTypedUrls       *bool `protobuf:"varint,18,opt,name=encrypt_typed_urls,json=encryptTypedUrls" json:"encrypt_typed_urls,omitempty"`
	EncryptExtensions      *bool `protobuf:"varint,19,opt,name=encrypt_extensions,json=encryptExtensions" json:"encrypt_extensions,omitempty"`
	EncryptSessions        *bool `protobuf:"varint,20,opt,name=encrypt_sessions,json=encryptSessions" json:"encrypt_sessions,omitempty"`
	EncryptApps            *bool `protobuf:"varint,21,opt,name=encrypt_apps,json=encryptApps" json:"encrypt_apps,omitempty"`
	EncryptSearchEngines   *bool `protobuf:"varint,22,opt,name=encrypt_search_engines,json=encryptSearchEngines" json:"encrypt_search_engines,omitempty"`
	// If true, all current and future datatypes will be encrypted.
	EncryptEverything        *bool `protobuf:"varint,24,opt,name=encrypt_everything,json=encryptEverything" json:"encrypt_everything,omitempty"`
	EncryptExtensionSettings *bool `protobuf:"varint,25,opt,name=encrypt_extension_settings,json=encryptExtensionSettings" json:"encrypt_extension_settings,omitempty"`
	EncryptAppNotifications  *bool `protobuf:"varint,26,opt,name=encrypt_app_notifications,json=encryptAppNotifications" json:"encrypt_app_notifications,omitempty"` // Deprecated: Do not use.
	EncryptAppSettings       *bool `protobuf:"varint,27,opt,name=encrypt_app_settings,json=encryptAppSettings" json:"encrypt_app_settings,omitempty"`
	// Enable syncing favicons as part of tab sync.
	SyncTabFavicons *bool `protobuf:"varint,29,opt,name=sync_tab_favicons,json=syncTabFavicons" json:"sync_tab_favicons,omitempty"`
	// An |int| field is used instead of enum PassphraseType so we can better
	// handle unknown values coming from later versions. Prior to M77, this was an
	// enum so old clients will assume IMPLICIT_PASSPHRASE for values greater than
	// 4.
	PassphraseType *int32 `protobuf:"varint,30,opt,name=passphrase_type,json=passphraseType,def=1" json:"passphrase_type,omitempty"`
	// The keystore decryptor token blob. Encrypted with the keystore key, and
	// contains the encryption key used to decrypt |encryption_keybag|.
	// Only set if passphrase_state == KEYSTORE_PASSPHRASE.
	KeystoreDecryptorToken *EncryptedData `protobuf:"bytes,31,opt,name=keystore_decryptor_token,json=keystoreDecryptorToken" json:"keystore_decryptor_token,omitempty"`
	// The time (in epoch milliseconds) at which the keystore migration was
	// performed.
	KeystoreMigrationTime *int64 `protobuf:"varint,32,opt,name=keystore_migration_time,json=keystoreMigrationTime" json:"keystore_migration_time,omitempty"`
	// The time (in epoch milliseconds) at which a custom passphrase was set.
	// Note: this field may not be set if the custom passphrase was applied before
	// this field was introduced.
	CustomPassphraseTime *int64 `protobuf:"varint,33,opt,name=custom_passphrase_time,json=customPassphraseTime" json:"custom_passphrase_time,omitempty"`
	// Boolean corresponding to whether custom spelling dictionary should be
	// encrypted.
	EncryptDictionary *bool `protobuf:"varint,34,opt,name=encrypt_dictionary,json=encryptDictionary" json:"encrypt_dictionary,omitempty"`
	// Boolean corresponding to Whether to encrypt favicons data or not.
	EncryptFaviconImages   *bool `protobuf:"varint,35,opt,name=encrypt_favicon_images,json=encryptFaviconImages" json:"encrypt_favicon_images,omitempty"`
	EncryptFaviconTracking *bool `protobuf:"varint,36,opt,name=encrypt_favicon_tracking,json=encryptFaviconTracking" json:"encrypt_favicon_tracking,omitempty"`
	// Boolean corresponding to whether app list items should be encrypted.
	EncryptAppList *bool `protobuf:"varint,38,opt,name=encrypt_app_list,json=encryptAppList" json:"encrypt_app_list,omitempty"`
	// Boolean corresponding to whether usage count and last use date of Wallet
	// data should be encrypted.
	EncryptAutofillWalletMetadata *bool `protobuf:"varint,39,opt,name=encrypt_autofill_wallet_metadata,json=encryptAutofillWalletMetadata" json:"encrypt_autofill_wallet_metadata,omitempty"`
	// Boolean indicating whether this node was originally missing a
	// |keystore_migration_time| field value, and was fixed on the server by
	// giving the field a value.
	// THIS FIELD SHOULD ONLY BE SET BY THE SERVER.
	ServerOnlyWasMissingKeystoreMigrationTime *bool `protobuf:"varint,40,opt,name=server_only_was_missing_keystore_migration_time,json=serverOnlyWasMissingKeystoreMigrationTime" json:"server_only_was_missing_keystore_migration_time,omitempty"`
	// Boolean corresponding to whether arc pakcage items should be encrypted.
	EncryptArcPackage *bool `protobuf:"varint,41,opt,name=encrypt_arc_package,json=encryptArcPackage" json:"encrypt_arc_package,omitempty"`
	// Boolean corresponding to whether printer items should be encrypted.
	EncryptPrinters *bool `protobuf:"varint,42,opt,name=encrypt_printers,json=encryptPrinters" json:"encrypt_printers,omitempty"`
	// Boolean corresponding to whether reading list items should be encrypted.
	EncryptReadingList *bool `protobuf:"varint,43,opt,name=encrypt_reading_list,json=encryptReadingList" json:"encrypt_reading_list,omitempty"`
	// ID of the method used to derive the encryption key from a custom
	// passphrase. Should be set only when |passphrase_type| is CUSTOM_PASSPHRASE
	// and only based on CustomPassphraseKeyDerivationMethod. This field has been
	// added in M70. All previous versions just ignore it, attempt to use
	// PBKDF2_HMAC_SHA1_1003 and, thus, reject any passphrase if a different
	// method has been used. The default corresponds to UNSPECIFIED. An |int|
	// field is used so we can detect unknown values coming from later versions.
	CustomPassphraseKeyDerivationMethod *int32 `protobuf:"varint,45,opt,name=custom_passphrase_key_derivation_method,json=customPassphraseKeyDerivationMethod,def=0" json:"custom_passphrase_key_derivation_method,omitempty"`
	// Base64-encoded salt used for the derivation of the key from the custom
	// passphrase. Valid only if custom_passphrase_key_derivation_method ==
	// SCRYPT_8192_8_11, ignored in other cases.
	CustomPassphraseKeyDerivationSalt *string `protobuf:"bytes,46,opt,name=custom_passphrase_key_derivation_salt,json=customPassphraseKeyDerivationSalt" json:"custom_passphrase_key_derivation_salt,omitempty"`
	// Boolean corresponding to whether send tab should be encrypted.
	EncryptSendTabToSelf *bool `protobuf:"varint,47,opt,name=encrypt_send_tab_to_self,json=encryptSendTabToSelf" json:"encrypt_send_tab_to_self,omitempty"`
	// Boolean corresponding to whether Web Apps data should be encrypted.
	EncryptWebApps *bool `protobuf:"varint,48,opt,name=encrypt_web_apps,json=encryptWebApps" json:"encrypt_web_apps,omitempty"`
	// Boolean corresponding to whether OS preferences should be encrypted.
	EncryptOsPreferences *bool    `protobuf:"varint,49,opt,name=encrypt_os_preferences,json=encryptOsPreferences" json:"encrypt_os_preferences,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NigoriSpecifics) Reset()         { *m = NigoriSpecifics{} }
func (m *NigoriSpecifics) String() string { return proto.CompactTextString(m) }
func (*NigoriSpecifics) ProtoMessage()    {}
func (*NigoriSpecifics) Descriptor() ([]byte, []int) {
	return fileDescriptor_54e3415a76ee5ff6, []int{2}
}

func (m *NigoriSpecifics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NigoriSpecifics.Unmarshal(m, b)
}
func (m *NigoriSpecifics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NigoriSpecifics.Marshal(b, m, deterministic)
}
func (m *NigoriSpecifics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NigoriSpecifics.Merge(m, src)
}
func (m *NigoriSpecifics) XXX_Size() int {
	return xxx_messageInfo_NigoriSpecifics.Size(m)
}
func (m *NigoriSpecifics) XXX_DiscardUnknown() {
	xxx_messageInfo_NigoriSpecifics.DiscardUnknown(m)
}

var xxx_messageInfo_NigoriSpecifics proto.InternalMessageInfo

const Default_NigoriSpecifics_PassphraseType int32 = 1
const Default_NigoriSpecifics_CustomPassphraseKeyDerivationMethod int32 = 0

func (m *NigoriSpecifics) GetEncryptionKeybag() *EncryptedData {
	if m != nil {
		return m.EncryptionKeybag
	}
	return nil
}

func (m *NigoriSpecifics) GetKeybagIsFrozen() bool {
	if m != nil && m.KeybagIsFrozen != nil {
		return *m.KeybagIsFrozen
	}
	return false
}

func (m *NigoriSpecifics) GetEncryptBookmarks() bool {
	if m != nil && m.EncryptBookmarks != nil {
		return *m.EncryptBookmarks
	}
	return false
}

func (m *NigoriSpecifics) GetEncryptPreferences() bool {
	if m != nil && m.EncryptPreferences != nil {
		return *m.EncryptPreferences
	}
	return false
}

func (m *NigoriSpecifics) GetEncryptAutofillProfile() bool {
	if m != nil && m.EncryptAutofillProfile != nil {
		return *m.EncryptAutofillProfile
	}
	return false
}

func (m *NigoriSpecifics) GetEncryptAutofill() bool {
	if m != nil && m.EncryptAutofill != nil {
		return *m.EncryptAutofill
	}
	return false
}

func (m *NigoriSpecifics) GetEncryptThemes() bool {
	if m != nil && m.EncryptThemes != nil {
		return *m.EncryptThemes
	}
	return false
}

func (m *NigoriSpecifics) GetEncryptTypedUrls() bool {
	if m != nil && m.EncryptTypedUrls != nil {
		return *m.EncryptTypedUrls
	}
	return false
}

func (m *NigoriSpecifics) GetEncryptExtensions() bool {
	if m != nil && m.EncryptExtensions != nil {
		return *m.EncryptExtensions
	}
	return false
}

func (m *NigoriSpecifics) GetEncryptSessions() bool {
	if m != nil && m.EncryptSessions != nil {
		return *m.EncryptSessions
	}
	return false
}

func (m *NigoriSpecifics) GetEncryptApps() bool {
	if m != nil && m.EncryptApps != nil {
		return *m.EncryptApps
	}
	return false
}

func (m *NigoriSpecifics) GetEncryptSearchEngines() bool {
	if m != nil && m.EncryptSearchEngines != nil {
		return *m.EncryptSearchEngines
	}
	return false
}

func (m *NigoriSpecifics) GetEncryptEverything() bool {
	if m != nil && m.EncryptEverything != nil {
		return *m.EncryptEverything
	}
	return false
}

func (m *NigoriSpecifics) GetEncryptExtensionSettings() bool {
	if m != nil && m.EncryptExtensionSettings != nil {
		return *m.EncryptExtensionSettings
	}
	return false
}

// Deprecated: Do not use.
func (m *NigoriSpecifics) GetEncryptAppNotifications() bool {
	if m != nil && m.EncryptAppNotifications != nil {
		return *m.EncryptAppNotifications
	}
	return false
}

func (m *NigoriSpecifics) GetEncryptAppSettings() bool {
	if m != nil && m.EncryptAppSettings != nil {
		return *m.EncryptAppSettings
	}
	return false
}

func (m *NigoriSpecifics) GetSyncTabFavicons() bool {
	if m != nil && m.SyncTabFavicons != nil {
		return *m.SyncTabFavicons
	}
	return false
}

func (m *NigoriSpecifics) GetPassphraseType() int32 {
	if m != nil && m.PassphraseType != nil {
		return *m.PassphraseType
	}
	return Default_NigoriSpecifics_PassphraseType
}

func (m *NigoriSpecifics) GetKeystoreDecryptorToken() *EncryptedData {
	if m != nil {
		return m.KeystoreDecryptorToken
	}
	return nil
}

func (m *NigoriSpecifics) GetKeystoreMigrationTime() int64 {
	if m != nil && m.KeystoreMigrationTime != nil {
		return *m.KeystoreMigrationTime
	}
	return 0
}

func (m *NigoriSpecifics) GetCustomPassphraseTime() int64 {
	if m != nil && m.CustomPassphraseTime != nil {
		return *m.CustomPassphraseTime
	}
	return 0
}

func (m *NigoriSpecifics) GetEncryptDictionary() bool {
	if m != nil && m.EncryptDictionary != nil {
		return *m.EncryptDictionary
	}
	return false
}

func (m *NigoriSpecifics) GetEncryptFaviconImages() bool {
	if m != nil && m.EncryptFaviconImages != nil {
		return *m.EncryptFaviconImages
	}
	return false
}

func (m *NigoriSpecifics) GetEncryptFaviconTracking() bool {
	if m != nil && m.EncryptFaviconTracking != nil {
		return *m.EncryptFaviconTracking
	}
	return false
}

func (m *NigoriSpecifics) GetEncryptAppList() bool {
	if m != nil && m.EncryptAppList != nil {
		return *m.EncryptAppList
	}
	return false
}

func (m *NigoriSpecifics) GetEncryptAutofillWalletMetadata() bool {
	if m != nil && m.EncryptAutofillWalletMetadata != nil {
		return *m.EncryptAutofillWalletMetadata
	}
	return false
}

func (m *NigoriSpecifics) GetServerOnlyWasMissingKeystoreMigrationTime() bool {
	if m != nil && m.ServerOnlyWasMissingKeystoreMigrationTime != nil {
		return *m.ServerOnlyWasMissingKeystoreMigrationTime
	}
	return false
}

func (m *NigoriSpecifics) GetEncryptArcPackage() bool {
	if m != nil && m.EncryptArcPackage != nil {
		return *m.EncryptArcPackage
	}
	return false
}

func (m *NigoriSpecifics) GetEncryptPrinters() bool {
	if m != nil && m.EncryptPrinters != nil {
		return *m.EncryptPrinters
	}
	return false
}

func (m *NigoriSpecifics) GetEncryptReadingList() bool {
	if m != nil && m.EncryptReadingList != nil {
		return *m.EncryptReadingList
	}
	return false
}

func (m *NigoriSpecifics) GetCustomPassphraseKeyDerivationMethod() int32 {
	if m != nil && m.CustomPassphraseKeyDerivationMethod != nil {
		return *m.CustomPassphraseKeyDerivationMethod
	}
	return Default_NigoriSpecifics_CustomPassphraseKeyDerivationMethod
}

func (m *NigoriSpecifics) GetCustomPassphraseKeyDerivationSalt() string {
	if m != nil && m.CustomPassphraseKeyDerivationSalt != nil {
		return *m.CustomPassphraseKeyDerivationSalt
	}
	return ""
}

func (m *NigoriSpecifics) GetEncryptSendTabToSelf() bool {
	if m != nil && m.EncryptSendTabToSelf != nil {
		return *m.EncryptSendTabToSelf
	}
	return false
}

func (m *NigoriSpecifics) GetEncryptWebApps() bool {
	if m != nil && m.EncryptWebApps != nil {
		return *m.EncryptWebApps
	}
	return false
}

func (m *NigoriSpecifics) GetEncryptOsPreferences() bool {
	if m != nil && m.EncryptOsPreferences != nil {
		return *m.EncryptOsPreferences
	}
	return false
}

func init() {
	proto.RegisterEnum("sync_pb.NigoriSpecifics_PassphraseType", NigoriSpecifics_PassphraseType_name, NigoriSpecifics_PassphraseType_value)
	proto.RegisterEnum("sync_pb.NigoriSpecifics_KeyDerivationMethod", NigoriSpecifics_KeyDerivationMethod_name, NigoriSpecifics_KeyDerivationMethod_value)
	proto.RegisterType((*NigoriKey)(nil), "sync_pb.NigoriKey")
	proto.RegisterType((*NigoriKeyBag)(nil), "sync_pb.NigoriKeyBag")
	proto.RegisterType((*NigoriSpecifics)(nil), "sync_pb.NigoriSpecifics")
}

func init() {
	proto.RegisterFile("nigori_specifics.proto", fileDescriptor_54e3415a76ee5ff6)
}

var fileDescriptor_54e3415a76ee5ff6 = []byte{
	// 1194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x56, 0xeb, 0x6e, 0x13, 0x47,
	0x14, 0xc6, 0x49, 0x28, 0x30, 0x81, 0xc4, 0x99, 0xdc, 0x86, 0xb4, 0xd0, 0x10, 0x08, 0x04, 0x02,
	0x26, 0x4e, 0x11, 0xa2, 0x55, 0x55, 0xc9, 0x89, 0x9d, 0xc6, 0x18, 0x3b, 0xab, 0xdd, 0x75, 0x23,
	0xfa, 0x67, 0x34, 0x5e, 0x1f, 0xdb, 0x23, 0xef, 0x4d, 0x33, 0x93, 0x50, 0xf7, 0x79, 0xfa, 0xb3,
	0xef, 0xd0, 0x57, 0xab, 0x76, 0xf6, 0x6e, 0x28, 0xfd, 0x65, 0xe9, 0x7c, 0xdf, 0x77, 0xe6, 0xdc,
	0xd7, 0x68, 0xcb, 0xe7, 0xe3, 0x40, 0x70, 0x2a, 0x43, 0x70, 0xf8, 0x88, 0x3b, 0xb2, 0x16, 0x8a,
	0x40, 0x05, 0xf8, 0x96, 0x9c, 0xf9, 0x0e, 0x0d, 0x07, 0x3b, 0x55, 0xf0, 0x1d, 0x31, 0x0b, 0x15,
	0x0f, 0xfc, 0x18, 0xda, 0xfb, 0xbb, 0x82, 0xee, 0xf4, 0xb4, 0xaa, 0x03, 0x33, 0x7c, 0x88, 0x56,
	0x87, 0x10, 0x0a, 0x70, 0x98, 0x82, 0x21, 0xf5, 0x99, 0x07, 0xa4, 0xb2, 0x5b, 0x39, 0xb8, 0x73,
	0xb2, 0x40, 0x2a, 0xe6, 0x4a, 0x0e, 0xf5, 0x98, 0x07, 0xf8, 0x18, 0xad, 0x17, 0xc8, 0x57, 0x12,
	0x04, 0x9d, 0xc2, 0x8c, 0x2c, 0xec, 0x56, 0x0e, 0xee, 0x6a, 0xc1, 0x5a, 0x0e, 0xf7, 0x25, 0x88,
	0xe8, 0x81, 0x7d, 0xb4, 0x92, 0x87, 0xa0, 0xe9, 0x8b, 0x11, 0xdd, 0xbc, 0x97, 0x5b, 0x23, 0xda,
	0x36, 0xba, 0xe5, 0x31, 0x47, 0xe3, 0x4b, 0x1a, 0xff, 0xc6, 0x63, 0x4e, 0x07, 0x66, 0x7b, 0x6f,
	0xd0, 0xdd, 0x2c, 0xda, 0x13, 0x36, 0xc6, 0x4f, 0xd0, 0x62, 0xfc, 0xe6, 0xe2, 0xc1, 0xf2, 0x31,
	0xae, 0x25, 0x79, 0xd6, 0x32, 0x8e, 0x19, 0xc1, 0x7b, 0xff, 0xac, 0xa1, 0xd5, 0xd8, 0x64, 0xa5,
	0x95, 0xc1, 0xa7, 0x68, 0xad, 0x1c, 0xc9, 0x80, 0x8d, 0x75, 0xb2, 0xcb, 0xc7, 0x5b, 0x99, 0x9f,
	0x56, 0xcc, 0x80, 0x61, 0x93, 0x29, 0x66, 0x56, 0x4b, 0x41, 0x0e, 0xd8, 0x18, 0x1f, 0xa0, 0x6a,
	0xac, 0xa4, 0x5c, 0xd2, 0x91, 0x08, 0xfe, 0x04, 0x5f, 0xe7, 0x7f, 0xdb, 0x5c, 0x89, 0xed, 0x6d,
	0x79, 0xa6, 0xad, 0xf8, 0x30, 0x7b, 0x8e, 0x0e, 0x82, 0x60, 0xea, 0x31, 0x31, 0x95, 0xe4, 0x9e,
	0xa6, 0xa6, 0x6e, 0x4f, 0x52, 0x3b, 0x7e, 0x8d, 0xd6, 0x53, 0x72, 0x28, 0x60, 0x04, 0x02, 0x7c,
	0x07, 0x24, 0x59, 0xd1, 0x74, 0x9c, 0x40, 0x46, 0x8e, 0xe0, 0x77, 0x88, 0xa4, 0x02, 0x76, 0xa5,
	0x82, 0x11, 0x77, 0x5d, 0x1a, 0x8a, 0xe8, 0x17, 0xc8, 0xaa, 0x56, 0x6d, 0x25, 0x78, 0x23, 0x81,
	0x8d, 0x18, 0xc5, 0xcf, 0x51, 0x75, 0x5e, 0x49, 0xaa, 0x5a, 0xb1, 0x3a, 0xa7, 0x28, 0xf4, 0x8e,
	0xaa, 0x09, 0x78, 0x20, 0xc9, 0x9a, 0x26, 0xa6, 0xbd, 0xb3, 0xb5, 0x11, 0xbf, 0x44, 0x38, 0xa3,
	0xcd, 0xc2, 0x68, 0x32, 0x84, 0x2b, 0x09, 0x2e, 0xa5, 0x6a, 0x47, 0x40, 0x5f, 0xb8, 0x12, 0xbf,
	0xca, 0xd9, 0xf0, 0x87, 0x02, 0x5f, 0xf2, 0xc0, 0x97, 0x64, 0x5d, 0xb3, 0xd3, 0x8a, 0xb5, 0x32,
	0xa0, 0x18, 0xae, 0x04, 0x19, 0x93, 0x37, 0x4a, 0xe1, 0x5a, 0x89, 0x19, 0x3f, 0x42, 0x77, 0xb3,
	0xcc, 0xc2, 0x50, 0x92, 0x4d, 0x4d, 0x5b, 0x4e, 0xb3, 0x0a, 0x43, 0x89, 0xdf, 0xa0, 0xad, 0xdc,
	0x1b, 0x13, 0xce, 0x84, 0x82, 0x3f, 0xe6, 0x3e, 0x48, 0xb2, 0xa5, 0xc9, 0x1b, 0x99, 0xcf, 0x08,
	0x6c, 0xc5, 0x58, 0x29, 0xe4, 0x6b, 0x10, 0x33, 0x35, 0xe1, 0xfe, 0x98, 0x90, 0x72, 0xc8, 0x19,
	0x80, 0x7f, 0x46, 0x3b, 0x9f, 0x65, 0x48, 0x25, 0x28, 0xc5, 0xfd, 0xb1, 0x24, 0xf7, 0xb5, 0x8c,
	0xcc, 0x67, 0x6a, 0x25, 0x38, 0xfe, 0x05, 0xdd, 0x2f, 0x64, 0x41, 0xfd, 0x40, 0x45, 0xe3, 0xcb,
	0x94, 0xce, 0x7c, 0x27, 0x12, 0xeb, 0x55, 0xdb, 0xce, 0xd3, 0xea, 0x15, 0x29, 0xf8, 0x08, 0x6d,
	0x14, 0xf5, 0xd9, 0xbb, 0xdf, 0x96, 0x66, 0xa9, 0x11, 0x86, 0xd9, 0x8b, 0x2f, 0xd0, 0x9a, 0x1e,
	0x7f, 0xc5, 0x06, 0x74, 0xc4, 0xae, 0xb9, 0x13, 0xbd, 0xf4, 0x20, 0xae, 0x71, 0x04, 0xd8, 0x6c,
	0x70, 0x96, 0x98, 0xf1, 0x0b, 0xb4, 0x1a, 0x32, 0x29, 0xc3, 0x89, 0x60, 0x12, 0x74, 0xbb, 0xc9,
	0xc3, 0xdd, 0xca, 0xc1, 0xcd, 0x9f, 0x2a, 0x75, 0x73, 0x25, 0x47, 0xa2, 0x76, 0x63, 0x03, 0x91,
	0x29, 0xcc, 0xa4, 0x0a, 0x04, 0xd0, 0x21, 0xe8, 0x67, 0x03, 0x41, 0x55, 0x30, 0x05, 0x9f, 0x7c,
	0xff, 0xd5, 0xbd, 0xdb, 0x4a, 0x75, 0xcd, 0x54, 0x66, 0x47, 0x2a, 0xfc, 0x16, 0x6d, 0x67, 0x1e,
	0x3d, 0x3e, 0x16, 0x3a, 0x65, 0xaa, 0xb8, 0x07, 0x64, 0x77, 0xb7, 0x72, 0xb0, 0x68, 0x6e, 0xa6,
	0x70, 0x37, 0x45, 0x6d, 0xee, 0x41, 0xd4, 0x76, 0xe7, 0x4a, 0xaa, 0xc0, 0xa3, 0xc5, 0xe0, 0x23,
	0xd9, 0x23, 0x2d, 0xdb, 0x88, 0x51, 0x23, 0x8f, 0x3f, 0x52, 0x15, 0xda, 0x3e, 0xe4, 0x4e, 0xe4,
	0x8c, 0x89, 0x19, 0xd9, 0x2b, 0xb5, 0xbd, 0x99, 0x01, 0xc5, 0xd9, 0x4a, 0xaa, 0x48, 0xb9, 0xc7,
	0xc6, 0x20, 0xc9, 0xe3, 0xd2, 0x6c, 0x25, 0xb5, 0x6c, 0x6b, 0xac, 0xb8, 0xc8, 0xa9, 0x4a, 0x09,
	0xe6, 0x4c, 0xa3, 0x09, 0x7b, 0x52, 0x5a, 0xe4, 0x44, 0x67, 0x27, 0x68, 0x74, 0x8a, 0x8a, 0x8d,
	0x76, 0xb9, 0x54, 0xe4, 0x69, 0x7c, 0x8a, 0xf2, 0x26, 0x7f, 0xe0, 0x52, 0xe1, 0x5f, 0xd1, 0xee,
	0x67, 0xc7, 0xe2, 0x13, 0x73, 0x5d, 0x50, 0xd4, 0x03, 0xc5, 0x86, 0x4c, 0x31, 0xf2, 0x4c, 0x2b,
	0x1f, 0xcc, 0x9d, 0x80, 0x4b, 0xcd, 0xea, 0x26, 0x24, 0x3c, 0x40, 0xaf, 0x25, 0x88, 0x6b, 0x10,
	0x34, 0xf0, 0xdd, 0x19, 0xfd, 0xc4, 0x24, 0xf5, 0xb8, 0x94, 0xdc, 0x1f, 0xd3, 0xff, 0xea, 0xcb,
	0x81, 0xf6, 0xfb, 0x3c, 0x96, 0x5d, 0xf8, 0xee, 0xec, 0x92, 0xc9, 0x6e, 0xac, 0xe9, 0x7c, 0xb1,
	0x57, 0xb5, 0xfc, 0x14, 0x32, 0xe1, 0xd0, 0x90, 0x39, 0x53, 0x36, 0x06, 0xf2, 0xbc, 0x54, 0xf6,
	0x86, 0x70, 0x8c, 0x18, 0x28, 0x1e, 0x88, 0x50, 0x70, 0x5f, 0x81, 0x90, 0xe4, 0x45, 0xe9, 0x40,
	0x18, 0x89, 0xb9, 0xb8, 0x1a, 0x02, 0xd8, 0x30, 0x0a, 0x5b, 0x57, 0xed, 0xb0, 0xb4, 0x1a, 0x66,
	0x0c, 0xe9, 0xca, 0x5d, 0xa2, 0x67, 0x9f, 0x0f, 0xce, 0x14, 0x66, 0x74, 0x08, 0x82, 0x5f, 0xc7,
	0x69, 0x7a, 0xa0, 0x26, 0xc1, 0x90, 0xbc, 0x8a, 0xd7, 0xe0, 0xc8, 0x7c, 0x3c, 0x3f, 0x4c, 0x1d,
	0x98, 0x35, 0x33, 0x7a, 0x57, 0xb3, 0xb1, 0x81, 0xf6, 0xff, 0xd7, 0xb1, 0x64, 0xae, 0x22, 0xb5,
	0xe8, 0x6b, 0x6c, 0x3e, 0xfa, 0xaa, 0x4f, 0x8b, 0xb9, 0x0a, 0xbf, 0xcd, 0x07, 0x49, 0x82, 0x3f,
	0xd4, 0xdb, 0xac, 0x02, 0x2a, 0xc1, 0x1d, 0x91, 0xd7, 0x73, 0xc7, 0xcd, 0x1f, 0xda, 0x6c, 0x60,
	0x07, 0x16, 0xb8, 0xa3, 0xe2, 0x18, 0x7d, 0x82, 0x41, 0x7c, 0x39, 0x8f, 0x4a, 0x63, 0x74, 0x09,
	0x83, 0xf9, 0xe3, 0x19, 0xc8, 0xd2, 0x77, 0xaa, 0x5e, 0xf2, 0x7f, 0x21, 0x0b, 0x5f, 0xaa, 0xbd,
	0xbf, 0x2a, 0x68, 0xc5, 0x28, 0x1f, 0x86, 0x65, 0x74, 0xab, 0xdf, 0xeb, 0xf4, 0x2e, 0x2e, 0x7b,
	0xd5, 0x1b, 0x78, 0x1b, 0xad, 0xb7, 0xbb, 0xc6, 0x87, 0xf6, 0x69, 0xdb, 0xa6, 0x46, 0xc3, 0xb2,
	0x8c, 0x73, 0xb3, 0x61, 0xb5, 0xaa, 0x95, 0x08, 0xe8, 0xb4, 0x3e, 0x5a, 0xf6, 0x85, 0xd9, 0x2a,
	0x02, 0x0b, 0xf8, 0x21, 0xda, 0x39, 0x33, 0x2f, 0x7e, 0x6f, 0xf5, 0xe8, 0x97, 0x84, 0x8b, 0x78,
	0x13, 0xad, 0x9d, 0xf6, 0x2d, 0xfb, 0xa2, 0x5b, 0x34, 0x2f, 0xe1, 0xef, 0x10, 0xb1, 0xcd, 0xbe,
	0x65, 0xb7, 0x9a, 0xf4, 0xb7, 0x46, 0xff, 0x43, 0x49, 0x74, 0x73, 0xef, 0x12, 0xad, 0x7f, 0xa9,
	0x4f, 0xab, 0x68, 0xb9, 0xdf, 0xb3, 0x8c, 0xd6, 0x69, 0xfb, 0xac, 0xdd, 0x6a, 0x56, 0x6f, 0xe0,
	0xfb, 0x68, 0xd3, 0x38, 0xe9, 0x34, 0xcf, 0x8e, 0xe9, 0x79, 0xb7, 0x71, 0x4a, 0xad, 0xf3, 0x46,
	0x9d, 0xd6, 0x8f, 0x8e, 0x7e, 0xa8, 0x56, 0xf0, 0x06, 0xaa, 0x5a, 0xa7, 0xe6, 0x47, 0xc3, 0xa6,
	0xef, 0xea, 0x3f, 0x1e, 0xd3, 0x77, 0xb4, 0x5e, 0xaf, 0x2e, 0xbc, 0x5f, 0xba, 0xbd, 0x5f, 0x7d,
	0xfa, 0x7e, 0xe9, 0xf6, 0xcb, 0xea, 0x2b, 0x33, 0x5f, 0x58, 0xa1, 0xb8, 0xe3, 0x82, 0xcc, 0xce,
	0x38, 0xf5, 0x82, 0x2b, 0x5f, 0x31, 0xee, 0x53, 0x39, 0x61, 0x02, 0xe4, 0xc9, 0x21, 0xda, 0x0f,
	0xc4, 0xb8, 0xe6, 0x4c, 0x44, 0xe0, 0xf1, 0x2b, 0xaf, 0xe6, 0x04, 0x5e, 0x18, 0xf8, 0xe0, 0x2b,
	0xa9, 0x6f, 0x66, 0xfc, 0x67, 0xce, 0x09, 0xdc, 0xf3, 0x45, 0xa3, 0xf2, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x88, 0x06, 0xef, 0x4a, 0x07, 0x0a, 0x00, 0x00,
}
