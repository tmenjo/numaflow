/*
Copyright 2022 The Numaproj Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package applier

import (
	"context"

	"github.com/numaproj/numaflow/pkg/isb"
)

// MapApplier applies the UDF on the read message and gives back a new message. Any UserError will be retried here, while
// InternalErr can be returned and could be retried by the callee.
type MapApplier interface {
	ApplyMap(ctx context.Context, message *isb.ReadMessage) ([]*isb.WriteMessage, error)
	ApplyMapStream(ctx context.Context, message *isb.ReadMessage, writeMessageCh chan<- isb.WriteMessage) error
}

// ApplyMapFunc utility function used to create an Applier implementation
type ApplyMapFunc struct {
	applyMap       func(context.Context, *isb.ReadMessage) ([]*isb.WriteMessage, error)
	applyMapStream func(context.Context, *isb.ReadMessage, chan<- isb.WriteMessage) error
}

func (a ApplyMapFunc) ApplyMap(ctx context.Context, message *isb.ReadMessage) ([]*isb.WriteMessage, error) {
	return a.applyMap(ctx, message)
}

func (a ApplyMapFunc) ApplyMapStream(ctx context.Context, message *isb.ReadMessage, writeMessageCh chan<- isb.WriteMessage) error {
	return a.applyMapStream(ctx, message, writeMessageCh)
}

var (
	// Terminal Applier do not make any change to the message
	Terminal = ApplyMapFunc{
		applyMap: func(ctx context.Context, msg *isb.ReadMessage) ([]*isb.WriteMessage, error) {
			return []*isb.WriteMessage{{
				Message: msg.Message,
			}}, nil
		},
		applyMapStream: func(ctx context.Context, msg *isb.ReadMessage, writeMessageCh chan<- isb.WriteMessage) error {
			defer close(writeMessageCh)
			writeMessage := &isb.WriteMessage{
				Message: msg.Message,
			}

			writeMessageCh <- *writeMessage
			return nil
		},
	}
)
