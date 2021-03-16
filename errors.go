package httpsign

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func newPublicError(msg string) *gin.Error {
	return &gin.Error{
		Err:  errors.New(msg),
		Type: gin.ErrorTypePublic,
	}
}

var (
	// ErrInvalidAuthorizationHeader error when get invalid format of Authorization header
	ErrInvalidAuthorizationHeader = newPublicError("Authorization header format is incorrect")
	// ErrInvalidKeyID error when KeyID in header does not provided
	ErrInvalidKeyID = newPublicError("Invalid keyId")
	// ErrIncorrectAlgorithm error when Algorithm in header does not match with secret key
	ErrIncorrectAlgorithm = newPublicError("Algorithm does not match")
	// ErrHeaderNotEnough error when requiremts header do not appear on heder field
	ErrHeaderNotEnough = newPublicError("Header field is not match requirement")
	// ErrNoSignature error when no Signature not found in header
	ErrNoSignature = newPublicError("No Signature header found in request")
	// ErrInvalidSign error when signing string do not match
	ErrInvalidSign = newPublicError("Invalid sign")
	// ErrMissingKeyID error when keyId not in header
	ErrMissingKeyID = newPublicError("keyId must be on header")
	// ErrMissingSignature error when signature not in header
	ErrMissingSignature = newPublicError("signature must be on header")

	// ErrUnterminatedParameter err when could not parse value
	ErrUnterminatedParameter = newPublicError("Unterminated parameter")
	// ErrMissingDoubleQuote err when after character = not have double quote
	ErrMissingDoubleQuote = newPublicError(`Missing " after = character`)
	// ErrMissingEqualCharacter err when there is no character = before " or , character
	ErrMissingEqualCharacter = newPublicError(`Missing = character =`)
	// ErrEmptyHeader err when one of the required headers are empty
	ErrEmptyHeader = newPublicError(`Empty required header`)
)
