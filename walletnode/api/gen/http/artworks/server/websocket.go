// Code generated by goa v3.4.2, DO NOT EDIT.
//
// artworks WebSocket server streaming
//
// Command:
// $ goa gen github.com/pastelnetwork/gonode/walletnode/api/design

package server

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	artworks "github.com/pastelnetwork/gonode/walletnode/api/gen/artworks"
	goahttp "goa.design/goa/v3/http"
)

// ConnConfigurer holds the websocket connection configurer functions for the
// streaming endpoints in "artworks" service.
type ConnConfigurer struct {
	RegisterTaskStateFn goahttp.ConnConfigureFunc
	SearchFn            goahttp.ConnConfigureFunc
}

// RegisterTaskStateServerStream implements the
// artworks.RegisterTaskStateServerStream interface.
type RegisterTaskStateServerStream struct {
	once sync.Once
	// upgrader is the websocket connection upgrader.
	upgrader goahttp.Upgrader
	// configurer is the websocket connection configurer.
	configurer goahttp.ConnConfigureFunc
	// cancel is the context cancellation function which cancels the request
	// context when invoked.
	cancel context.CancelFunc
	// w is the HTTP response writer used in upgrading the connection.
	w http.ResponseWriter
	// r is the HTTP request.
	r *http.Request
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// SearchServerStream implements the artworks.SearchServerStream interface.
type SearchServerStream struct {
	once sync.Once
	// upgrader is the websocket connection upgrader.
	upgrader goahttp.Upgrader
	// configurer is the websocket connection configurer.
	configurer goahttp.ConnConfigureFunc
	// cancel is the context cancellation function which cancels the request
	// context when invoked.
	cancel context.CancelFunc
	// w is the HTTP response writer used in upgrading the connection.
	w http.ResponseWriter
	// r is the HTTP request.
	r *http.Request
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// NewConnConfigurer initializes the websocket connection configurer function
// with fn for all the streaming endpoints in "artworks" service.
func NewConnConfigurer(fn goahttp.ConnConfigureFunc) *ConnConfigurer {
	return &ConnConfigurer{
		RegisterTaskStateFn: fn,
		SearchFn:            fn,
	}
}

// Send streams instances of "artworks.TaskState" to the "registerTaskState"
// endpoint websocket connection.
func (s *RegisterTaskStateServerStream) Send(v *artworks.TaskState) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := v
	body := NewRegisterTaskStateResponseBody(res)
	return s.conn.WriteJSON(body)
}

// Close closes the "registerTaskState" endpoint websocket connection.
func (s *RegisterTaskStateServerStream) Close() error {
	var err error
	if s.conn == nil {
		return nil
	}
	if err = s.conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "server closing connection"),
		time.Now().Add(time.Second),
	); err != nil {
		return err
	}
	return s.conn.Close()
}

// Send streams instances of "artworks.ArtworkTicket" to the "search" endpoint
// websocket connection.
func (s *SearchServerStream) Send(v *artworks.ArtworkTicket) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.configurer != nil {
			conn = s.configurer(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := v
	body := NewSearchResponseBody(res)
	return s.conn.WriteJSON(body)
}

// Close closes the "search" endpoint websocket connection.
func (s *SearchServerStream) Close() error {
	var err error
	if s.conn == nil {
		return nil
	}
	if err = s.conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "server closing connection"),
		time.Now().Add(time.Second),
	); err != nil {
		return err
	}
	return s.conn.Close()
}
