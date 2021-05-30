// Code generated by goa v3.3.1, DO NOT EDIT.
//
// HTTP request path constructors for the artworks service.
//
// Command:
// $ goa gen github.com/pastelnetwork/gonode/walletnode/api/design

package server

import (
	"fmt"
)

// RegisterArtworksPath returns the URL path to the artworks service register HTTP endpoint.
func RegisterArtworksPath() string {
	return "/artworks/register"
}

// RegisterTaskStateArtworksPath returns the URL path to the artworks service registerTaskState HTTP endpoint.
func RegisterTaskStateArtworksPath(taskID string) string {
	return fmt.Sprintf("/artworks/register/%v/state", taskID)
}

// RegisterTaskArtworksPath returns the URL path to the artworks service registerTask HTTP endpoint.
func RegisterTaskArtworksPath(taskID string) string {
	return fmt.Sprintf("/artworks/register/%v", taskID)
}

// RegisterTasksArtworksPath returns the URL path to the artworks service registerTasks HTTP endpoint.
func RegisterTasksArtworksPath() string {
	return "/artworks/register"
}

// UploadImageArtworksPath returns the URL path to the artworks service uploadImage HTTP endpoint.
func UploadImageArtworksPath() string {
	return "/artworks/register/upload"
}

// SearchRequestArtworksPath returns the URL path to the artworks service searchRequest HTTP endpoint.
func SearchRequestArtworksPath(term string) string {
	return fmt.Sprintf("/artworks/search/%v", term)
}
