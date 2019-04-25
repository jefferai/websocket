// Package websocket is a minimal and idiomatic implementation of the WebSocket protocol.
//
// See https://tools.ietf.org/html/rfc6455
//
// Please see https://nhooyr.io/websocket for overview docs and a
// comparison with existing implementations.
//
// Conn, Dial, and Accept are the main entrypoints into this package. Use Dial to dial
// a WebSocket server, Accept to accept a WebSocket client dial and then Conn to interact
// with the resulting WebSocket connections.
//
// The echo example is the best way to understand how to correctly use the library.
//
// The wsjson and wspb packages contain helpers for JSON and ProtoBuf messages.
package websocket
