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
	"fmt"
	"bytes"
)

type backendRequest struct {
	endpoint *BackendEndpoint
	vhost    *Vhost
	id       []byte
	cmd      string
	msg      [][]byte
}

func newBackendRequest(endpoint *BackendEndpoint, vhost *Vhost, id []byte,
	cmd string, msg [][]byte) (r *backendRequest) {
	r = &backendRequest{
		endpoint: endpoint,
		vhost:    vhost,
		id:       id,
		cmd:      cmd,
		msg:      msg,
	}
	return r
}

func (r *backendRequest) Reply(cmd string, frames ...string) (err error) {
	return r.endpoint.SendTo(r.id, false, cmd, frames...)
}

func (r *backendRequest) String() string {
	s := string(bytes.Join(r.msg, []byte(", ")))
	return fmt.Sprintf("[%s, %s]", r.cmd, s)
}