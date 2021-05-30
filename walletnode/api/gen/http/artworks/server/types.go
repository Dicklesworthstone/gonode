// Code generated by goa v3.3.1, DO NOT EDIT.
//
// artworks HTTP server types
//
// Command:
// $ goa gen github.com/pastelnetwork/gonode/walletnode/api/design

package server

import (
	"unicode/utf8"

	artworks "github.com/pastelnetwork/gonode/walletnode/api/gen/artworks"
	artworksviews "github.com/pastelnetwork/gonode/walletnode/api/gen/artworks/views"
	goa "goa.design/goa/v3/pkg"
)

// RegisterRequestBody is the type of the "artworks" service "register"
// endpoint HTTP request body.
type RegisterRequestBody struct {
	// Uploaded image ID
	ImageID *string `form:"image_id,omitempty" json:"image_id,omitempty" xml:"image_id,omitempty"`
	// Name of the artwork
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Description of the artwork
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Keywords
	Keywords *string `form:"keywords,omitempty" json:"keywords,omitempty" xml:"keywords,omitempty"`
	// Series name
	SeriesName *string `form:"series_name,omitempty" json:"series_name,omitempty" xml:"series_name,omitempty"`
	// Number of copies issued
	IssuedCopies *int `form:"issued_copies,omitempty" json:"issued_copies,omitempty" xml:"issued_copies,omitempty"`
	// Artwork creation video youtube URL
	YoutubeURL *string `form:"youtube_url,omitempty" json:"youtube_url,omitempty" xml:"youtube_url,omitempty"`
	// Artist's PastelID
	ArtistPastelID *string `form:"artist_pastelid,omitempty" json:"artist_pastelid,omitempty" xml:"artist_pastelid,omitempty"`
	// Passphrase of the artist's PastelID
	ArtistPastelIDPassphrase *string `form:"artist_pastelid_passphrase,omitempty" json:"artist_pastelid_passphrase,omitempty" xml:"artist_pastelid_passphrase,omitempty"`
	// Name of the artist
	ArtistName *string `form:"artist_name,omitempty" json:"artist_name,omitempty" xml:"artist_name,omitempty"`
	// Artist website URL
	ArtistWebsiteURL *string `form:"artist_website_url,omitempty" json:"artist_website_url,omitempty" xml:"artist_website_url,omitempty"`
	// Spendable address
	SpendableAddress *string `form:"spendable_address,omitempty" json:"spendable_address,omitempty" xml:"spendable_address,omitempty"`
	// Used to find a suitable masternode with a fee equal or less
	MaximumFee *float64 `form:"maximum_fee,omitempty" json:"maximum_fee,omitempty" xml:"maximum_fee,omitempty"`
}

// UploadImageRequestBody is the type of the "artworks" service "uploadImage"
// endpoint HTTP request body.
type UploadImageRequestBody struct {
	// File to upload
	Bytes []byte `form:"file,omitempty" json:"file,omitempty" xml:"file,omitempty"`
	// For internal use
	Filename *string `form:"filename,omitempty" json:"filename,omitempty" xml:"filename,omitempty"`
}

// RegisterResponseBody is the type of the "artworks" service "register"
// endpoint HTTP response body.
type RegisterResponseBody struct {
	// Task ID of the registration process
	TaskID string `form:"task_id" json:"task_id" xml:"task_id"`
}

// RegisterTaskStateResponseBody is the type of the "artworks" service
// "registerTaskState" endpoint HTTP response body.
type RegisterTaskStateResponseBody struct {
	// Date of the status creation
	Date string `form:"date" json:"date" xml:"date"`
	// Status of the registration process
	Status string `form:"status" json:"status" xml:"status"`
}

// RegisterTaskResponseBody is the type of the "artworks" service
// "registerTask" endpoint HTTP response body.
type RegisterTaskResponseBody struct {
	// JOb ID of the registration process
	ID string `form:"id" json:"id" xml:"id"`
	// Status of the registration process
	Status string `form:"status" json:"status" xml:"status"`
	// List of states from the very beginning of the process
	States []*TaskStateResponseBody `form:"states,omitempty" json:"states,omitempty" xml:"states,omitempty"`
	// txid
	Txid   *string                    `form:"txid,omitempty" json:"txid,omitempty" xml:"txid,omitempty"`
	Ticket *ArtworkTicketResponseBody `form:"ticket" json:"ticket" xml:"ticket"`
}

// TaskResponseTinyCollection is the type of the "artworks" service
// "registerTasks" endpoint HTTP response body.
type TaskResponseTinyCollection []*TaskResponseTiny

// UploadImageResponseBody is the type of the "artworks" service "uploadImage"
// endpoint HTTP response body.
type UploadImageResponseBody struct {
	// Uploaded image ID
	ImageID string `form:"image_id" json:"image_id" xml:"image_id"`
	// Image expiration
	ExpiresIn string `form:"expires_in" json:"expires_in" xml:"expires_in"`
}

// SearchRequestResponseBody is the type of the "artworks" service
// "searchRequest" endpoint HTTP response body.
type SearchRequestResponseBody struct {
	Ticket *struct {
		Type         string  `form:"type" json:"type" xml:"type"`
		PastelID     *string `form:"pastelID" json:"pastelID" xml:"pastelID"`
		RegTxid      *string `form:"reg_txid" json:"reg_txid" xml:"reg_txid"`
		ArtistHeight *string `form:"artist_height" json:"artist_height" xml:"artist_height"`
		RegFee       *string `form:"reg_fee" json:"reg_fee" xml:"reg_fee"`
		Signature    *string `form:"signature" json:"signature" xml:"signature"`
	} `form:"ticket,omitempty" json:"ticket,omitempty" xml:"ticket,omitempty"`
}

// RegisterBadRequestResponseBody is the type of the "artworks" service
// "register" endpoint HTTP response body for the "BadRequest" error.
type RegisterBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// RegisterInternalServerErrorResponseBody is the type of the "artworks"
// service "register" endpoint HTTP response body for the "InternalServerError"
// error.
type RegisterInternalServerErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// RegisterTaskStateNotFoundResponseBody is the type of the "artworks" service
// "registerTaskState" endpoint HTTP response body for the "NotFound" error.
type RegisterTaskStateNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// RegisterTaskStateInternalServerErrorResponseBody is the type of the
// "artworks" service "registerTaskState" endpoint HTTP response body for the
// "InternalServerError" error.
type RegisterTaskStateInternalServerErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// RegisterTaskNotFoundResponseBody is the type of the "artworks" service
// "registerTask" endpoint HTTP response body for the "NotFound" error.
type RegisterTaskNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// RegisterTaskInternalServerErrorResponseBody is the type of the "artworks"
// service "registerTask" endpoint HTTP response body for the
// "InternalServerError" error.
type RegisterTaskInternalServerErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// RegisterTasksInternalServerErrorResponseBody is the type of the "artworks"
// service "registerTasks" endpoint HTTP response body for the
// "InternalServerError" error.
type RegisterTasksInternalServerErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UploadImageBadRequestResponseBody is the type of the "artworks" service
// "uploadImage" endpoint HTTP response body for the "BadRequest" error.
type UploadImageBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UploadImageInternalServerErrorResponseBody is the type of the "artworks"
// service "uploadImage" endpoint HTTP response body for the
// "InternalServerError" error.
type UploadImageInternalServerErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// SearchRequestBadRequestResponseBody is the type of the "artworks" service
// "searchRequest" endpoint HTTP response body for the "BadRequest" error.
type SearchRequestBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// SearchRequestInternalServerErrorResponseBody is the type of the "artworks"
// service "searchRequest" endpoint HTTP response body for the
// "InternalServerError" error.
type SearchRequestInternalServerErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// TaskStateResponseBody is used to define fields on response body types.
type TaskStateResponseBody struct {
	// Date of the status creation
	Date string `form:"date" json:"date" xml:"date"`
	// Status of the registration process
	Status string `form:"status" json:"status" xml:"status"`
}

// ArtworkTicketResponseBody is used to define fields on response body types.
type ArtworkTicketResponseBody struct {
	// Name of the artwork
	Name string `form:"name" json:"name" xml:"name"`
	// Description of the artwork
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Keywords
	Keywords *string `form:"keywords,omitempty" json:"keywords,omitempty" xml:"keywords,omitempty"`
	// Series name
	SeriesName *string `form:"series_name,omitempty" json:"series_name,omitempty" xml:"series_name,omitempty"`
	// Number of copies issued
	IssuedCopies int `form:"issued_copies" json:"issued_copies" xml:"issued_copies"`
	// Artwork creation video youtube URL
	YoutubeURL *string `form:"youtube_url,omitempty" json:"youtube_url,omitempty" xml:"youtube_url,omitempty"`
	// Artist's PastelID
	ArtistPastelID string `form:"artist_pastelid" json:"artist_pastelid" xml:"artist_pastelid"`
	// Passphrase of the artist's PastelID
	ArtistPastelIDPassphrase string `form:"artist_pastelid_passphrase" json:"artist_pastelid_passphrase" xml:"artist_pastelid_passphrase"`
	// Name of the artist
	ArtistName string `form:"artist_name" json:"artist_name" xml:"artist_name"`
	// Artist website URL
	ArtistWebsiteURL *string `form:"artist_website_url,omitempty" json:"artist_website_url,omitempty" xml:"artist_website_url,omitempty"`
	// Spendable address
	SpendableAddress string `form:"spendable_address" json:"spendable_address" xml:"spendable_address"`
	// Used to find a suitable masternode with a fee equal or less
	MaximumFee float64 `form:"maximum_fee" json:"maximum_fee" xml:"maximum_fee"`
}

// TaskResponseTiny is used to define fields on response body types.
type TaskResponseTiny struct {
	// JOb ID of the registration process
	ID string `form:"id" json:"id" xml:"id"`
	// Status of the registration process
	Status string `form:"status" json:"status" xml:"status"`
	// txid
	Txid   *string                `form:"txid,omitempty" json:"txid,omitempty" xml:"txid,omitempty"`
	Ticket *ArtworkTicketResponse `form:"ticket" json:"ticket" xml:"ticket"`
}

// ArtworkTicketResponse is used to define fields on response body types.
type ArtworkTicketResponse struct {
	// Name of the artwork
	Name string `form:"name" json:"name" xml:"name"`
	// Description of the artwork
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Keywords
	Keywords *string `form:"keywords,omitempty" json:"keywords,omitempty" xml:"keywords,omitempty"`
	// Series name
	SeriesName *string `form:"series_name,omitempty" json:"series_name,omitempty" xml:"series_name,omitempty"`
	// Number of copies issued
	IssuedCopies int `form:"issued_copies" json:"issued_copies" xml:"issued_copies"`
	// Artwork creation video youtube URL
	YoutubeURL *string `form:"youtube_url,omitempty" json:"youtube_url,omitempty" xml:"youtube_url,omitempty"`
	// Artist's PastelID
	ArtistPastelID string `form:"artist_pastelid" json:"artist_pastelid" xml:"artist_pastelid"`
	// Passphrase of the artist's PastelID
	ArtistPastelIDPassphrase string `form:"artist_pastelid_passphrase" json:"artist_pastelid_passphrase" xml:"artist_pastelid_passphrase"`
	// Name of the artist
	ArtistName string `form:"artist_name" json:"artist_name" xml:"artist_name"`
	// Artist website URL
	ArtistWebsiteURL *string `form:"artist_website_url,omitempty" json:"artist_website_url,omitempty" xml:"artist_website_url,omitempty"`
	// Spendable address
	SpendableAddress string `form:"spendable_address" json:"spendable_address" xml:"spendable_address"`
	// Used to find a suitable masternode with a fee equal or less
	MaximumFee float64 `form:"maximum_fee" json:"maximum_fee" xml:"maximum_fee"`
}

// NewRegisterResponseBody builds the HTTP response body from the result of the
// "register" endpoint of the "artworks" service.
func NewRegisterResponseBody(res *artworksviews.RegisterResultView) *RegisterResponseBody {
	body := &RegisterResponseBody{
		TaskID: *res.TaskID,
	}
	return body
}

// NewRegisterTaskStateResponseBody builds the HTTP response body from the
// result of the "registerTaskState" endpoint of the "artworks" service.
func NewRegisterTaskStateResponseBody(res *artworks.TaskState) *RegisterTaskStateResponseBody {
	body := &RegisterTaskStateResponseBody{
		Date:   res.Date,
		Status: res.Status,
	}
	return body
}

// NewRegisterTaskResponseBody builds the HTTP response body from the result of
// the "registerTask" endpoint of the "artworks" service.
func NewRegisterTaskResponseBody(res *artworksviews.TaskView) *RegisterTaskResponseBody {
	body := &RegisterTaskResponseBody{
		ID:     *res.ID,
		Status: *res.Status,
		Txid:   res.Txid,
	}
	if res.States != nil {
		body.States = make([]*TaskStateResponseBody, len(res.States))
		for i, val := range res.States {
			body.States[i] = marshalArtworksviewsTaskStateViewToTaskStateResponseBody(val)
		}
	}
	if res.Ticket != nil {
		body.Ticket = marshalArtworksviewsArtworkTicketViewToArtworkTicketResponseBody(res.Ticket)
	}
	return body
}

// NewTaskResponseTinyCollection builds the HTTP response body from the result
// of the "registerTasks" endpoint of the "artworks" service.
func NewTaskResponseTinyCollection(res artworksviews.TaskCollectionView) TaskResponseTinyCollection {
	body := make([]*TaskResponseTiny, len(res))
	for i, val := range res {
		body[i] = marshalArtworksviewsTaskViewToTaskResponseTiny(val)
	}
	return body
}

// NewUploadImageResponseBody builds the HTTP response body from the result of
// the "uploadImage" endpoint of the "artworks" service.
func NewUploadImageResponseBody(res *artworksviews.ImageView) *UploadImageResponseBody {
	body := &UploadImageResponseBody{
		ImageID:   *res.ImageID,
		ExpiresIn: *res.ExpiresIn,
	}
	return body
}

// NewSearchRequestResponseBody builds the HTTP response body from the result
// of the "searchRequest" endpoint of the "artworks" service.
func NewSearchRequestResponseBody(res *artworks.ArtworkSearchResult) *SearchRequestResponseBody {
	body := &SearchRequestResponseBody{}
	if res.Ticket != nil {
		body.Ticket = &struct {
			Type         string  `form:"type" json:"type" xml:"type"`
			PastelID     *string `form:"pastelID" json:"pastelID" xml:"pastelID"`
			RegTxid      *string `form:"reg_txid" json:"reg_txid" xml:"reg_txid"`
			ArtistHeight *string `form:"artist_height" json:"artist_height" xml:"artist_height"`
			RegFee       *string `form:"reg_fee" json:"reg_fee" xml:"reg_fee"`
			Signature    *string `form:"signature" json:"signature" xml:"signature"`
		}{
			Type:         res.Ticket.Type,
			PastelID:     res.Ticket.PastelID,
			RegTxid:      res.Ticket.RegTxid,
			ArtistHeight: res.Ticket.ArtistHeight,
			RegFee:       res.Ticket.RegFee,
			Signature:    res.Ticket.Signature,
		}
		{
			var zero string
			if body.Ticket.Type == zero {
				body.Ticket.Type = "art-act"
			}
		}
	}
	return body
}

// NewRegisterBadRequestResponseBody builds the HTTP response body from the
// result of the "register" endpoint of the "artworks" service.
func NewRegisterBadRequestResponseBody(res *goa.ServiceError) *RegisterBadRequestResponseBody {
	body := &RegisterBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRegisterInternalServerErrorResponseBody builds the HTTP response body
// from the result of the "register" endpoint of the "artworks" service.
func NewRegisterInternalServerErrorResponseBody(res *goa.ServiceError) *RegisterInternalServerErrorResponseBody {
	body := &RegisterInternalServerErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRegisterTaskStateNotFoundResponseBody builds the HTTP response body from
// the result of the "registerTaskState" endpoint of the "artworks" service.
func NewRegisterTaskStateNotFoundResponseBody(res *goa.ServiceError) *RegisterTaskStateNotFoundResponseBody {
	body := &RegisterTaskStateNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRegisterTaskStateInternalServerErrorResponseBody builds the HTTP response
// body from the result of the "registerTaskState" endpoint of the "artworks"
// service.
func NewRegisterTaskStateInternalServerErrorResponseBody(res *goa.ServiceError) *RegisterTaskStateInternalServerErrorResponseBody {
	body := &RegisterTaskStateInternalServerErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRegisterTaskNotFoundResponseBody builds the HTTP response body from the
// result of the "registerTask" endpoint of the "artworks" service.
func NewRegisterTaskNotFoundResponseBody(res *goa.ServiceError) *RegisterTaskNotFoundResponseBody {
	body := &RegisterTaskNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRegisterTaskInternalServerErrorResponseBody builds the HTTP response body
// from the result of the "registerTask" endpoint of the "artworks" service.
func NewRegisterTaskInternalServerErrorResponseBody(res *goa.ServiceError) *RegisterTaskInternalServerErrorResponseBody {
	body := &RegisterTaskInternalServerErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRegisterTasksInternalServerErrorResponseBody builds the HTTP response
// body from the result of the "registerTasks" endpoint of the "artworks"
// service.
func NewRegisterTasksInternalServerErrorResponseBody(res *goa.ServiceError) *RegisterTasksInternalServerErrorResponseBody {
	body := &RegisterTasksInternalServerErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUploadImageBadRequestResponseBody builds the HTTP response body from the
// result of the "uploadImage" endpoint of the "artworks" service.
func NewUploadImageBadRequestResponseBody(res *goa.ServiceError) *UploadImageBadRequestResponseBody {
	body := &UploadImageBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUploadImageInternalServerErrorResponseBody builds the HTTP response body
// from the result of the "uploadImage" endpoint of the "artworks" service.
func NewUploadImageInternalServerErrorResponseBody(res *goa.ServiceError) *UploadImageInternalServerErrorResponseBody {
	body := &UploadImageInternalServerErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewSearchRequestBadRequestResponseBody builds the HTTP response body from
// the result of the "searchRequest" endpoint of the "artworks" service.
func NewSearchRequestBadRequestResponseBody(res *goa.ServiceError) *SearchRequestBadRequestResponseBody {
	body := &SearchRequestBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewSearchRequestInternalServerErrorResponseBody builds the HTTP response
// body from the result of the "searchRequest" endpoint of the "artworks"
// service.
func NewSearchRequestInternalServerErrorResponseBody(res *goa.ServiceError) *SearchRequestInternalServerErrorResponseBody {
	body := &SearchRequestInternalServerErrorResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRegisterPayload builds a artworks service register endpoint payload.
func NewRegisterPayload(body *RegisterRequestBody) *artworks.RegisterPayload {
	v := &artworks.RegisterPayload{
		ImageID:                  *body.ImageID,
		Name:                     *body.Name,
		Description:              body.Description,
		Keywords:                 body.Keywords,
		SeriesName:               body.SeriesName,
		IssuedCopies:             *body.IssuedCopies,
		YoutubeURL:               body.YoutubeURL,
		ArtistPastelID:           *body.ArtistPastelID,
		ArtistPastelIDPassphrase: *body.ArtistPastelIDPassphrase,
		ArtistName:               *body.ArtistName,
		ArtistWebsiteURL:         body.ArtistWebsiteURL,
		SpendableAddress:         *body.SpendableAddress,
		MaximumFee:               *body.MaximumFee,
	}

	return v
}

// NewRegisterTaskStatePayload builds a artworks service registerTaskState
// endpoint payload.
func NewRegisterTaskStatePayload(taskID string) *artworks.RegisterTaskStatePayload {
	v := &artworks.RegisterTaskStatePayload{}
	v.TaskID = taskID

	return v
}

// NewRegisterTaskPayload builds a artworks service registerTask endpoint
// payload.
func NewRegisterTaskPayload(taskID string) *artworks.RegisterTaskPayload {
	v := &artworks.RegisterTaskPayload{}
	v.TaskID = taskID

	return v
}

// NewUploadImagePayload builds a artworks service uploadImage endpoint payload.
func NewUploadImagePayload(body *UploadImageRequestBody) *artworks.UploadImagePayload {
	v := &artworks.UploadImagePayload{
		Bytes:    body.Bytes,
		Filename: body.Filename,
	}

	return v
}

// NewSearchRequestArtworkSearchRequestPayload builds a artworks service
// searchRequest endpoint payload.
func NewSearchRequestArtworkSearchRequestPayload(term string) *artworks.ArtworkSearchRequestPayload {
	v := &artworks.ArtworkSearchRequestPayload{}
	v.Term = &term

	return v
}

// ValidateRegisterRequestBody runs the validations defined on
// RegisterRequestBody
func ValidateRegisterRequestBody(body *RegisterRequestBody) (err error) {
	if body.ImageID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("image_id", "body"))
	}
	if body.ArtistName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("artist_name", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.IssuedCopies == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("issued_copies", "body"))
	}
	if body.ArtistPastelID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("artist_pastelid", "body"))
	}
	if body.ArtistPastelIDPassphrase == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("artist_pastelid_passphrase", "body"))
	}
	if body.SpendableAddress == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("spendable_address", "body"))
	}
	if body.MaximumFee == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("maximum_fee", "body"))
	}
	if body.ImageID != nil {
		if utf8.RuneCountInString(*body.ImageID) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.image_id", *body.ImageID, utf8.RuneCountInString(*body.ImageID), 8, true))
		}
	}
	if body.ImageID != nil {
		if utf8.RuneCountInString(*body.ImageID) > 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.image_id", *body.ImageID, utf8.RuneCountInString(*body.ImageID), 8, false))
		}
	}
	if body.Name != nil {
		if utf8.RuneCountInString(*body.Name) > 256 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.name", *body.Name, utf8.RuneCountInString(*body.Name), 256, false))
		}
	}
	if body.Description != nil {
		if utf8.RuneCountInString(*body.Description) > 1024 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", *body.Description, utf8.RuneCountInString(*body.Description), 1024, false))
		}
	}
	if body.Keywords != nil {
		if utf8.RuneCountInString(*body.Keywords) > 256 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.keywords", *body.Keywords, utf8.RuneCountInString(*body.Keywords), 256, false))
		}
	}
	if body.SeriesName != nil {
		if utf8.RuneCountInString(*body.SeriesName) > 256 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.series_name", *body.SeriesName, utf8.RuneCountInString(*body.SeriesName), 256, false))
		}
	}
	if body.IssuedCopies != nil {
		if *body.IssuedCopies < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.issued_copies", *body.IssuedCopies, 1, true))
		}
	}
	if body.IssuedCopies != nil {
		if *body.IssuedCopies > 1000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.issued_copies", *body.IssuedCopies, 1000, false))
		}
	}
	if body.YoutubeURL != nil {
		if utf8.RuneCountInString(*body.YoutubeURL) > 128 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.youtube_url", *body.YoutubeURL, utf8.RuneCountInString(*body.YoutubeURL), 128, false))
		}
	}
	if body.ArtistPastelID != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.artist_pastelid", *body.ArtistPastelID, "^[a-zA-Z0-9]+$"))
	}
	if body.ArtistPastelID != nil {
		if utf8.RuneCountInString(*body.ArtistPastelID) < 86 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.artist_pastelid", *body.ArtistPastelID, utf8.RuneCountInString(*body.ArtistPastelID), 86, true))
		}
	}
	if body.ArtistPastelID != nil {
		if utf8.RuneCountInString(*body.ArtistPastelID) > 86 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.artist_pastelid", *body.ArtistPastelID, utf8.RuneCountInString(*body.ArtistPastelID), 86, false))
		}
	}
	if body.ArtistName != nil {
		if utf8.RuneCountInString(*body.ArtistName) > 256 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.artist_name", *body.ArtistName, utf8.RuneCountInString(*body.ArtistName), 256, false))
		}
	}
	if body.ArtistWebsiteURL != nil {
		if utf8.RuneCountInString(*body.ArtistWebsiteURL) > 256 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.artist_website_url", *body.ArtistWebsiteURL, utf8.RuneCountInString(*body.ArtistWebsiteURL), 256, false))
		}
	}
	if body.SpendableAddress != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.spendable_address", *body.SpendableAddress, "^[a-zA-Z0-9]+$"))
	}
	if body.SpendableAddress != nil {
		if utf8.RuneCountInString(*body.SpendableAddress) < 35 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.spendable_address", *body.SpendableAddress, utf8.RuneCountInString(*body.SpendableAddress), 35, true))
		}
	}
	if body.SpendableAddress != nil {
		if utf8.RuneCountInString(*body.SpendableAddress) > 35 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.spendable_address", *body.SpendableAddress, utf8.RuneCountInString(*body.SpendableAddress), 35, false))
		}
	}
	if body.MaximumFee != nil {
		if *body.MaximumFee < 1e-05 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.maximum_fee", *body.MaximumFee, 1e-05, true))
		}
	}
	return
}

// ValidateUploadImageRequestBody runs the validations defined on
// UploadImageRequestBody
func ValidateUploadImageRequestBody(body *UploadImageRequestBody) (err error) {
	if body.Bytes == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("file", "body"))
	}
	return
}
