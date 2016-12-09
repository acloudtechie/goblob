package crypto

type DigestProvider interface {
	CreateFromFile(path string, algorithm DigestAlgorithm) (Digest, error)
}

type Digest interface {
	VerifyingDigest

	Algorithm() DigestAlgorithm
	Digest() string
	String() string
	Compare(Digest) int
}

type VerifyingDigest interface {
	Verify(Digest) error
}
