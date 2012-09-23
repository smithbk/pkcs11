package pkcs11

import (
	"unsafe"
)

// Generated with:
// grep CKA *t.h|grep '#define' | sed 's/^#define //' | awk ' { print $1 "=" $2 } '

type CKA_MODULUS_BITS struct{ value uint32 }

func (c *CKA_MODULUS_BITS) Type() uint            { return cKA_MODULUS_BITS }
func (c *CKA_MODULUS_BITS) Value() unsafe.Pointer { return unsafe.Pointer(&c.value) }
func (c *CKA_MODULUS_BITS) Len() uint             { return 4 }

const (
	CKA_CLASS                      = 0x00000000
	CKA_TOKEN                      = 0x00000001
	CKA_PRIVATE                    = 0x00000002
	CKA_LABEL                      = 0x00000003
	CKA_APPLICATION                = 0x00000010
	CKA_VALUE                      = 0x00000011
	CKA_OBJECT_ID                  = 0x00000012
	CKA_CERTIFICATE_TYPE           = 0x00000080
	CKA_ISSUER                     = 0x00000081
	CKA_SERIAL_NUMBER              = 0x00000082
	CKA_AC_ISSUER                  = 0x00000083
	CKA_OWNER                      = 0x00000084
	CKA_ATTR_TYPES                 = 0x00000085
	CKA_TRUSTED                    = 0x00000086
	CKA_CERTIFICATE_CATEGORY       = 0x00000087
	CKA_JAVA_MIDP_SECURITY_DOMAIN  = 0x00000088
	CKA_URL                        = 0x00000089
	CKA_HASH_OF_SUBJECT_PUBLIC_KEY = 0x0000008A
	CKA_HASH_OF_ISSUER_PUBLIC_KEY  = 0x0000008B
	CKA_CHECK_VALUE                = 0x00000090
	CKA_KEY_TYPE                   = 0x00000100
	CKA_SUBJECT                    = 0x00000101
	CKA_ID                         = 0x00000102
	CKA_SENSITIVE                  = 0x00000103
	CKA_ENCRYPT                    = 0x00000104
	CKA_DECRYPT                    = 0x00000105
	CKA_WRAP                       = 0x00000106
	CKA_UNWRAP                     = 0x00000107
	CKA_SIGN                       = 0x00000108
	CKA_SIGN_RECOVER               = 0x00000109
	CKA_VERIFY                     = 0x0000010A
	CKA_VERIFY_RECOVER             = 0x0000010B
	CKA_DERIVE                     = 0x0000010C
	CKA_START_DATE                 = 0x00000110
	CKA_END_DATE                   = 0x00000111
	CKA_MODULUS                    = 0x00000120
	cKA_MODULUS_BITS               = 0x00000121
	CKA_PUBLIC_EXPONENT            = 0x00000122
	CKA_PRIVATE_EXPONENT           = 0x00000123
	CKA_PRIME_1                    = 0x00000124
	CKA_PRIME_2                    = 0x00000125
	CKA_EXPONENT_1                 = 0x00000126
	CKA_EXPONENT_2                 = 0x00000127
	CKA_COEFFICIENT                = 0x00000128
	CKA_PRIME                      = 0x00000130
	CKA_SUBPRIME                   = 0x00000131
	CKA_BASE                       = 0x00000132
	CKA_PRIME_BITS                 = 0x00000133
	CKA_SUBPRIME_BITS              = 0x00000134
	CKA_SUB_PRIME_BITS             = CKA_SUBPRIME_BITS
	CKA_VALUE_BITS                 = 0x00000160
	CKA_VALUE_LEN                  = 0x00000161
	CKA_EXTRACTABLE                = 0x00000162
	CKA_LOCAL                      = 0x00000163
	CKA_NEVER_EXTRACTABLE          = 0x00000164
	CKA_ALWAYS_SENSITIVE           = 0x00000165
	CKA_KEY_GEN_MECHANISM          = 0x00000166
	CKA_MODIFIABLE                 = 0x00000170
	CKA_ECDSA_PARAMS               = 0x00000180
	CKA_EC_PARAMS                  = 0x00000180
	CKA_EC_POINT                   = 0x00000181
	CKA_SECONDARY_AUTH             = 0x00000200
	CKA_AUTH_PIN_FLAGS             = 0x00000201
	CKA_ALWAYS_AUTHENTICATE        = 0x00000202
	CKA_WRAP_WITH_TRUSTED          = 0x00000210
	CKA_WRAP_TEMPLATE              = (CKF_ARRAY_ATTRIBUTE | 0x00000211)
	CKA_UNWRAP_TEMPLATE            = (CKF_ARRAY_ATTRIBUTE | 0x00000212)
	CKA_OTP_FORMAT                 = 0x00000220
	CKA_OTP_LENGTH                 = 0x00000221
	CKA_OTP_TIME_INTERVAL          = 0x00000222
	CKA_OTP_USER_FRIENDLY_MODE     = 0x00000223
	CKA_OTP_CHALLENGE_REQUIREMENT  = 0x00000224
	CKA_OTP_TIME_REQUIREMENT       = 0x00000225
	CKA_OTP_COUNTER_REQUIREMENT    = 0x00000226
	CKA_OTP_PIN_REQUIREMENT        = 0x00000227
	CKA_OTP_COUNTER                = 0x0000022E
	CKA_OTP_TIME                   = 0x0000022F
	CKA_OTP_USER_IDENTIFIER        = 0x0000022A
	CKA_OTP_SERVICE_IDENTIFIER     = 0x0000022B
	CKA_OTP_SERVICE_LOGO           = 0x0000022C
	CKA_OTP_SERVICE_LOGO_TYPE      = 0x0000022D
	CKA_HW_FEATURE_TYPE            = 0x00000300
	CKA_RESET_ON_INIT              = 0x00000301
	CKA_HAS_RESET                  = 0x00000302
	CKA_PIXEL_X                    = 0x00000400
	CKA_PIXEL_Y                    = 0x00000401
	CKA_RESOLUTION                 = 0x00000402
	CKA_CHAR_ROWS                  = 0x00000403
	CKA_CHAR_COLUMNS               = 0x00000404
	CKA_COLOR                      = 0x00000405
	CKA_BITS_PER_PIXEL             = 0x00000406
	CKA_CHAR_SETS                  = 0x00000480
	CKA_ENCODING_METHODS           = 0x00000481
	CKA_MIME_TYPES                 = 0x00000482
	CKA_MECHANISM_TYPE             = 0x00000500
	CKA_REQUIRED_CMS_ATTRIBUTES    = 0x00000501
	CKA_DEFAULT_CMS_ATTRIBUTES     = 0x00000502
	CKA_SUPPORTED_CMS_ATTRIBUTES   = 0x00000503
	CKA_ALLOWED_MECHANISMS         = (CKF_ARRAY_ATTRIBUTE | 0x00000600)
	CKA_VENDOR_DEFINED             = 0x80000000
)