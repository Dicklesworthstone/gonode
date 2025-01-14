// Code generated by goa v3.3.1, DO NOT EDIT.
//
// artworks HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/pastelnetwork/gonode/walletnode/api/design

package server

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"unicode/utf8"

	artworks "github.com/pastelnetwork/gonode/walletnode/api/gen/artworks"
	artworksviews "github.com/pastelnetwork/gonode/walletnode/api/gen/artworks/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeRegisterResponse returns an encoder for responses returned by the
// artworks register endpoint.
func EncodeRegisterResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*artworksviews.RegisterResult)
		enc := encoder(ctx, w)
		body := NewRegisterResponseBody(res.Projected)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeRegisterRequest returns a decoder for requests sent to the artworks
// register endpoint.
func DecodeRegisterRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body RegisterRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateRegisterRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewRegisterPayload(&body)

		return payload, nil
	}
}

// EncodeRegisterError returns an encoder for errors returned by the register
// artworks endpoint.
func EncodeRegisterError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "BadRequest":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRegisterBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "BadRequest")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "InternalServerError":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRegisterInternalServerErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "InternalServerError")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// DecodeRegisterTaskStateRequest returns a decoder for requests sent to the
// artworks registerTaskState endpoint.
func DecodeRegisterTaskStateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			taskID string
			err    error

			params = mux.Vars(r)
		)
		taskID = params["taskId"]
		if utf8.RuneCountInString(taskID) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("taskID", taskID, utf8.RuneCountInString(taskID), 8, true))
		}
		if utf8.RuneCountInString(taskID) > 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("taskID", taskID, utf8.RuneCountInString(taskID), 8, false))
		}
		if err != nil {
			return nil, err
		}
		payload := NewRegisterTaskStatePayload(taskID)

		return payload, nil
	}
}

// EncodeRegisterTaskStateError returns an encoder for errors returned by the
// registerTaskState artworks endpoint.
func EncodeRegisterTaskStateError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "NotFound":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRegisterTaskStateNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "NotFound")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "InternalServerError":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRegisterTaskStateInternalServerErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "InternalServerError")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeRegisterTaskResponse returns an encoder for responses returned by the
// artworks registerTask endpoint.
func EncodeRegisterTaskResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*artworksviews.Task)
		enc := encoder(ctx, w)
		body := NewRegisterTaskResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeRegisterTaskRequest returns a decoder for requests sent to the
// artworks registerTask endpoint.
func DecodeRegisterTaskRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			taskID string
			err    error

			params = mux.Vars(r)
		)
		taskID = params["taskId"]
		if utf8.RuneCountInString(taskID) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("taskID", taskID, utf8.RuneCountInString(taskID), 8, true))
		}
		if utf8.RuneCountInString(taskID) > 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("taskID", taskID, utf8.RuneCountInString(taskID), 8, false))
		}
		if err != nil {
			return nil, err
		}
		payload := NewRegisterTaskPayload(taskID)

		return payload, nil
	}
}

// EncodeRegisterTaskError returns an encoder for errors returned by the
// registerTask artworks endpoint.
func EncodeRegisterTaskError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "NotFound":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRegisterTaskNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "NotFound")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "InternalServerError":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRegisterTaskInternalServerErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "InternalServerError")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeRegisterTasksResponse returns an encoder for responses returned by the
// artworks registerTasks endpoint.
func EncodeRegisterTasksResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(artworksviews.TaskCollection)
		enc := encoder(ctx, w)
		body := NewTaskResponseTinyCollection(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeRegisterTasksError returns an encoder for errors returned by the
// registerTasks artworks endpoint.
func EncodeRegisterTasksError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "InternalServerError":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRegisterTasksInternalServerErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "InternalServerError")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeUploadImageResponse returns an encoder for responses returned by the
// artworks uploadImage endpoint.
func EncodeUploadImageResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*artworksviews.Image)
		enc := encoder(ctx, w)
		body := NewUploadImageResponseBody(res.Projected)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeUploadImageRequest returns a decoder for requests sent to the artworks
// uploadImage endpoint.
func DecodeUploadImageRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var payload *artworks.UploadImagePayload
		if err := decoder(r).Decode(&payload); err != nil {
			return nil, goa.DecodePayloadError(err.Error())
		}

		return payload, nil
	}
}

// NewArtworksUploadImageDecoder returns a decoder to decode the multipart
// request for the "artworks" service "uploadImage" endpoint.
func NewArtworksUploadImageDecoder(mux goahttp.Muxer, artworksUploadImageDecoderFn ArtworksUploadImageDecoderFunc) func(r *http.Request) goahttp.Decoder {
	return func(r *http.Request) goahttp.Decoder {
		return goahttp.EncodingFunc(func(v interface{}) error {
			mr, merr := r.MultipartReader()
			if merr != nil {
				return merr
			}
			p := v.(**artworks.UploadImagePayload)
			if err := artworksUploadImageDecoderFn(mr, p); err != nil {
				return err
			}
			return nil
		})
	}
}

// EncodeUploadImageError returns an encoder for errors returned by the
// uploadImage artworks endpoint.
func EncodeUploadImageError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "BadRequest":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUploadImageBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "BadRequest")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "InternalServerError":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUploadImageInternalServerErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "InternalServerError")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// DecodeArtSearchRequest returns a decoder for requests sent to the artworks
// artSearch endpoint.
func DecodeArtSearchRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			artist           *string
			limit            int
			query            string
			artistName       bool
			artTitle         bool
			series           bool
			descr            bool
			keyword          bool
			minCopies        *int
			maxCopies        *int
			minBlock         int
			maxBlock         *int
			minRarenessScore *int
			maxRarenessScore *int
			minNsfwScore     *int
			maxNsfwScore     *int
			err              error
		)
		artistRaw := r.URL.Query().Get("artist")
		if artistRaw != "" {
			artist = &artistRaw
		}
		if artist != nil {
			if utf8.RuneCountInString(*artist) > 256 {
				err = goa.MergeErrors(err, goa.InvalidLengthError("artist", *artist, utf8.RuneCountInString(*artist), 256, false))
			}
		}
		{
			limitRaw := r.URL.Query().Get("limit")
			if limitRaw == "" {
				limit = 10
			} else {
				v, err2 := strconv.ParseInt(limitRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("limit", limitRaw, "integer"))
				}
				limit = int(v)
			}
		}
		if limit < 10 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 10, true))
		}
		if limit > 200 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("limit", limit, 200, false))
		}
		query = r.URL.Query().Get("query")
		if query == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("query", "query string"))
		}
		{
			artistNameRaw := r.URL.Query().Get("artist_name")
			if artistNameRaw == "" {
				artistName = true
			} else {
				v, err2 := strconv.ParseBool(artistNameRaw)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("artistName", artistNameRaw, "boolean"))
				}
				artistName = v
			}
		}
		{
			artTitleRaw := r.URL.Query().Get("art_title")
			if artTitleRaw == "" {
				artTitle = true
			} else {
				v, err2 := strconv.ParseBool(artTitleRaw)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("artTitle", artTitleRaw, "boolean"))
				}
				artTitle = v
			}
		}
		{
			seriesRaw := r.URL.Query().Get("series")
			if seriesRaw == "" {
				series = true
			} else {
				v, err2 := strconv.ParseBool(seriesRaw)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("series", seriesRaw, "boolean"))
				}
				series = v
			}
		}
		{
			descrRaw := r.URL.Query().Get("descr")
			if descrRaw == "" {
				descr = true
			} else {
				v, err2 := strconv.ParseBool(descrRaw)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("descr", descrRaw, "boolean"))
				}
				descr = v
			}
		}
		{
			keywordRaw := r.URL.Query().Get("keyword")
			if keywordRaw == "" {
				keyword = true
			} else {
				v, err2 := strconv.ParseBool(keywordRaw)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("keyword", keywordRaw, "boolean"))
				}
				keyword = v
			}
		}
		{
			minCopiesRaw := r.URL.Query().Get("min_copies")
			if minCopiesRaw != "" {
				v, err2 := strconv.ParseInt(minCopiesRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("minCopies", minCopiesRaw, "integer"))
				}
				pv := int(v)
				minCopies = &pv
			}
		}
		if minCopies != nil {
			if *minCopies < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("minCopies", *minCopies, 1, true))
			}
		}
		if minCopies != nil {
			if *minCopies > 1000 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("minCopies", *minCopies, 1000, false))
			}
		}
		{
			maxCopiesRaw := r.URL.Query().Get("max_copies")
			if maxCopiesRaw != "" {
				v, err2 := strconv.ParseInt(maxCopiesRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("maxCopies", maxCopiesRaw, "integer"))
				}
				pv := int(v)
				maxCopies = &pv
			}
		}
		if maxCopies != nil {
			if *maxCopies < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("maxCopies", *maxCopies, 1, true))
			}
		}
		if maxCopies != nil {
			if *maxCopies > 1000 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("maxCopies", *maxCopies, 1000, false))
			}
		}
		{
			minBlockRaw := r.URL.Query().Get("min_block")
			if minBlockRaw == "" {
				minBlock = 1
			} else {
				v, err2 := strconv.ParseInt(minBlockRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("minBlock", minBlockRaw, "integer"))
				}
				minBlock = int(v)
			}
		}
		if minBlock < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("minBlock", minBlock, 1, true))
		}
		{
			maxBlockRaw := r.URL.Query().Get("max_block")
			if maxBlockRaw != "" {
				v, err2 := strconv.ParseInt(maxBlockRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("maxBlock", maxBlockRaw, "integer"))
				}
				pv := int(v)
				maxBlock = &pv
			}
		}
		if maxBlock != nil {
			if *maxBlock < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("maxBlock", *maxBlock, 1, true))
			}
		}
		{
			minRarenessScoreRaw := r.URL.Query().Get("min_rareness_score")
			if minRarenessScoreRaw != "" {
				v, err2 := strconv.ParseInt(minRarenessScoreRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("minRarenessScore", minRarenessScoreRaw, "integer"))
				}
				pv := int(v)
				minRarenessScore = &pv
			}
		}
		if minRarenessScore != nil {
			if *minRarenessScore < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("minRarenessScore", *minRarenessScore, 1, true))
			}
		}
		if minRarenessScore != nil {
			if *minRarenessScore > 1000 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("minRarenessScore", *minRarenessScore, 1000, false))
			}
		}
		{
			maxRarenessScoreRaw := r.URL.Query().Get("max_rareness_score")
			if maxRarenessScoreRaw != "" {
				v, err2 := strconv.ParseInt(maxRarenessScoreRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("maxRarenessScore", maxRarenessScoreRaw, "integer"))
				}
				pv := int(v)
				maxRarenessScore = &pv
			}
		}
		if maxRarenessScore != nil {
			if *maxRarenessScore < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("maxRarenessScore", *maxRarenessScore, 1, true))
			}
		}
		if maxRarenessScore != nil {
			if *maxRarenessScore > 1000 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("maxRarenessScore", *maxRarenessScore, 1000, false))
			}
		}
		{
			minNsfwScoreRaw := r.URL.Query().Get("min_nsfw_score")
			if minNsfwScoreRaw != "" {
				v, err2 := strconv.ParseInt(minNsfwScoreRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("minNsfwScore", minNsfwScoreRaw, "integer"))
				}
				pv := int(v)
				minNsfwScore = &pv
			}
		}
		if minNsfwScore != nil {
			if *minNsfwScore < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("minNsfwScore", *minNsfwScore, 1, true))
			}
		}
		if minNsfwScore != nil {
			if *minNsfwScore > 1000 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("minNsfwScore", *minNsfwScore, 1000, false))
			}
		}
		{
			maxNsfwScoreRaw := r.URL.Query().Get("max_nsfw_score")
			if maxNsfwScoreRaw != "" {
				v, err2 := strconv.ParseInt(maxNsfwScoreRaw, 10, strconv.IntSize)
				if err2 != nil {
					err = goa.MergeErrors(err, goa.InvalidFieldTypeError("maxNsfwScore", maxNsfwScoreRaw, "integer"))
				}
				pv := int(v)
				maxNsfwScore = &pv
			}
		}
		if maxNsfwScore != nil {
			if *maxNsfwScore < 1 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("maxNsfwScore", *maxNsfwScore, 1, true))
			}
		}
		if maxNsfwScore != nil {
			if *maxNsfwScore > 1000 {
				err = goa.MergeErrors(err, goa.InvalidRangeError("maxNsfwScore", *maxNsfwScore, 1000, false))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewArtSearchPayload(artist, limit, query, artistName, artTitle, series, descr, keyword, minCopies, maxCopies, minBlock, maxBlock, minRarenessScore, maxRarenessScore, minNsfwScore, maxNsfwScore)

		return payload, nil
	}
}

// EncodeArtSearchError returns an encoder for errors returned by the artSearch
// artworks endpoint.
func EncodeArtSearchError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "BadRequest":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewArtSearchBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "BadRequest")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "InternalServerError":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewArtSearchInternalServerErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "InternalServerError")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeArtworkGetResponse returns an encoder for responses returned by the
// artworks artworkGet endpoint.
func EncodeArtworkGetResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*artworks.ArtworkDetail)
		enc := encoder(ctx, w)
		body := NewArtworkGetResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeArtworkGetRequest returns a decoder for requests sent to the artworks
// artworkGet endpoint.
func DecodeArtworkGetRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			txid string
			err  error

			params = mux.Vars(r)
		)
		txid = params["txid"]
		if utf8.RuneCountInString(txid) < 64 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("txid", txid, utf8.RuneCountInString(txid), 64, true))
		}
		if utf8.RuneCountInString(txid) > 64 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("txid", txid, utf8.RuneCountInString(txid), 64, false))
		}
		if err != nil {
			return nil, err
		}
		payload := NewArtworkGetPayload(txid)

		return payload, nil
	}
}

// EncodeArtworkGetError returns an encoder for errors returned by the
// artworkGet artworks endpoint.
func EncodeArtworkGetError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "BadRequest":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewArtworkGetBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "BadRequest")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "NotFound":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewArtworkGetNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "NotFound")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "InternalServerError":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewArtworkGetInternalServerErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "InternalServerError")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalArtworksviewsTaskStateViewToTaskStateResponseBody builds a value of
// type *TaskStateResponseBody from a value of type
// *artworksviews.TaskStateView.
func marshalArtworksviewsTaskStateViewToTaskStateResponseBody(v *artworksviews.TaskStateView) *TaskStateResponseBody {
	if v == nil {
		return nil
	}
	res := &TaskStateResponseBody{
		Date:   *v.Date,
		Status: *v.Status,
	}

	return res
}

// marshalArtworksviewsArtworkTicketViewToArtworkTicketResponseBody builds a
// value of type *ArtworkTicketResponseBody from a value of type
// *artworksviews.ArtworkTicketView.
func marshalArtworksviewsArtworkTicketViewToArtworkTicketResponseBody(v *artworksviews.ArtworkTicketView) *ArtworkTicketResponseBody {
	res := &ArtworkTicketResponseBody{
		Name:                     *v.Name,
		Description:              v.Description,
		Keywords:                 v.Keywords,
		SeriesName:               v.SeriesName,
		IssuedCopies:             *v.IssuedCopies,
		YoutubeURL:               v.YoutubeURL,
		ArtistPastelID:           *v.ArtistPastelID,
		ArtistPastelIDPassphrase: *v.ArtistPastelIDPassphrase,
		ArtistName:               *v.ArtistName,
		ArtistWebsiteURL:         v.ArtistWebsiteURL,
		SpendableAddress:         *v.SpendableAddress,
		MaximumFee:               *v.MaximumFee,
	}

	return res
}

// marshalArtworksviewsTaskViewToTaskResponseTiny builds a value of type
// *TaskResponseTiny from a value of type *artworksviews.TaskView.
func marshalArtworksviewsTaskViewToTaskResponseTiny(v *artworksviews.TaskView) *TaskResponseTiny {
	res := &TaskResponseTiny{
		ID:     *v.ID,
		Status: *v.Status,
		Txid:   v.Txid,
	}
	if v.Ticket != nil {
		res.Ticket = marshalArtworksviewsArtworkTicketViewToArtworkTicketResponse(v.Ticket)
	}

	return res
}

// marshalArtworksviewsArtworkTicketViewToArtworkTicketResponse builds a value
// of type *ArtworkTicketResponse from a value of type
// *artworksviews.ArtworkTicketView.
func marshalArtworksviewsArtworkTicketViewToArtworkTicketResponse(v *artworksviews.ArtworkTicketView) *ArtworkTicketResponse {
	res := &ArtworkTicketResponse{
		Name:                     *v.Name,
		Description:              v.Description,
		Keywords:                 v.Keywords,
		SeriesName:               v.SeriesName,
		IssuedCopies:             *v.IssuedCopies,
		YoutubeURL:               v.YoutubeURL,
		ArtistPastelID:           *v.ArtistPastelID,
		ArtistPastelIDPassphrase: *v.ArtistPastelIDPassphrase,
		ArtistName:               *v.ArtistName,
		ArtistWebsiteURL:         v.ArtistWebsiteURL,
		SpendableAddress:         *v.SpendableAddress,
		MaximumFee:               *v.MaximumFee,
	}

	return res
}

// marshalArtworksArtworkSummaryToArtworkSummaryResponseBody builds a value of
// type *ArtworkSummaryResponseBody from a value of type
// *artworks.ArtworkSummary.
func marshalArtworksArtworkSummaryToArtworkSummaryResponseBody(v *artworks.ArtworkSummary) *ArtworkSummaryResponseBody {
	res := &ArtworkSummaryResponseBody{
		Thumbnail:        v.Thumbnail,
		Txid:             v.Txid,
		Title:            v.Title,
		Description:      v.Description,
		Keywords:         v.Keywords,
		SeriesName:       v.SeriesName,
		Copies:           v.Copies,
		YoutubeURL:       v.YoutubeURL,
		ArtistPastelID:   v.ArtistPastelID,
		ArtistName:       v.ArtistName,
		ArtistWebsiteURL: v.ArtistWebsiteURL,
	}

	return res
}

// marshalArtworksFuzzyMatchToFuzzyMatchResponseBody builds a value of type
// *FuzzyMatchResponseBody from a value of type *artworks.FuzzyMatch.
func marshalArtworksFuzzyMatchToFuzzyMatchResponseBody(v *artworks.FuzzyMatch) *FuzzyMatchResponseBody {
	res := &FuzzyMatchResponseBody{
		Str:       v.Str,
		FieldType: v.FieldType,
		Score:     v.Score,
	}
	if v.MatchedIndexes != nil {
		res.MatchedIndexes = make([]int, len(v.MatchedIndexes))
		for i, val := range v.MatchedIndexes {
			res.MatchedIndexes[i] = val
		}
	}

	return res
}
