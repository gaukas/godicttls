package godicttls

// source: https://www.iana.org/assignments/tls-extensiontype-values/tls-extensiontype-values.xhtml#tls-certificate-compression-algorithm-ids
// last updated: March 2023

const (
	CertCompAlg_zlib   uint8 = 1
	CertCompAlg_brotli uint8 = 2
	CertCompAlg_zstd   uint8 = 3
)

var DictCertificateCompressionAlgorithmValueIndexed = map[uint8]string{
	1: "zlib",
	2: "brotli",
	3: "zstd",
}

var DictCertificateCompressionAlgorithmNameIndexed = map[string]uint8{
	"zlib":   1,
	"brotli": 2,
	"zstd":   3,
}
