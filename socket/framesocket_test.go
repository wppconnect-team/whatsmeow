package socket

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/coder/websocket"

	waLog "go.mau.fi/whatsmeow/util/log"
)

func TestSendFramePreservesWirePayloadAndOneTimeHeader(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	frames := make(chan []byte, 2)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Accept(w, r, nil)
		if err != nil {
			return
		}
		defer conn.CloseNow()
		conn.SetReadLimit(FrameMaxSize)
		for range 2 {
			kind, frame, err := conn.Read(ctx)
			if err != nil || kind != websocket.MessageBinary {
				return
			}
			frames <- frame
		}
	}))
	defer server.Close()
	conn, _, err := websocket.Dial(ctx, "ws"+strings.TrimPrefix(server.URL, "http"), nil)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.CloseNow()
	fs := NewFrameSocket(waLog.Noop, nil)
	fs.cancelCtx = ctx
	fs.conn.Store(conn)
	for i, payload := range [][]byte{bytes.Repeat([]byte{0x42, 0x71}, 131073), {0x01, 0x02, 0x03}} {
		if err := fs.SendFrame(payload); err != nil {
			t.Fatal(err)
		}
		select {
		case frame := <-frames:
			offset := 0
			if i == 0 {
				if !bytes.HasPrefix(frame, WAConnHeader) {
					t.Fatal("initial protocol header is missing")
				}
				offset = len(WAConnHeader)
			}
			if len(frame) < offset+3 {
				t.Fatal("length prefix is missing")
			}
			declared := int(frame[offset])<<16 | int(frame[offset+1])<<8 | int(frame[offset+2])
			if declared != len(payload) || !bytes.Equal(frame[offset+3:], payload) {
				t.Fatalf("frame %d lost its payload: declared=%d actual=%d want=%d", i, declared, len(frame)-offset-3, len(payload))
			}
		case <-ctx.Done():
			t.Fatal("peer did not receive the frame")
		}
	}
	fs.Header = make([]byte, FrameMaxSize)
	if err := fs.SendFrame([]byte{1}); !errors.Is(err, ErrFrameTooLarge) {
		t.Fatal("oversized header was accepted")
	}
}
