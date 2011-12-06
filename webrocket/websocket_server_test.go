// This package provides a hybrid of MQ and WebSockets server with
// support for horizontal scalability.
//
// Copyright (C) 2011 by Krzysztof Kowalik <chris@nu7hat.ch>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.
package webrocket

import (
	"bytes"
	"log"
	"testing"
)

func NewWsTestServer() *WebsocketServer {
	ctx := NewContext()
	ctx.Log = log.New(bytes.NewBuffer([]byte{}), "", log.LstdFlags)
	server := ctx.NewWebsocketServer(":9771")
	return server
}

func TestNewWsServer(t *testing.T) {
	server := NewWsTestServer()
	if server.Addr != ":9771" {
		t.Errorf("Expected server addr to be `:9771`, given %s", server.Addr)
	}
}