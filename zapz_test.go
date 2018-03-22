// Copyright © 2018 Douglas Chimento <dchimento@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package zapz

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dougEfresh/logzio-go"
	"github.com/magiconair/properties/assert"
	"go.uber.org/zap/zapcore"
)

func TestNew(t *testing.T) {
	var sent []byte
	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		defer r.Body.Close()
		sent, _ = ioutil.ReadAll(r.Body)
	}))
	ts.Start()
	defer ts.Close()
	logz, _ := logzio.New("fake", logzio.SetUrl(ts.URL))
	z, _ := New("fake", SetLogz(logz))

	z.Info("test")
	z.Sync()
	var sentMap map[string]string
	err := json.Unmarshal(sent, &sentMap)
	if err != nil {
		t.Fatal("error on json ", err)
	}
	time.Sleep(time.Millisecond * 200)
	assert.Equal(t, sentMap["type"], "zap-logger")
	assert.Equal(t, sentMap["level"], "info")
	assert.Equal(t, sentMap["message"], "test")
	if len(sentMap["ts"]) == 0 {
		t.Fatal("wrong ts", sentMap["ts"])
	}
}

func TestLevel(t *testing.T) {
	var sent []byte
	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		defer r.Body.Close()
		sent, _ = ioutil.ReadAll(r.Body)
	}))
	ts.Start()
	defer ts.Close()
	logz, _ := logzio.New("fake", logzio.SetUrl(ts.URL))
	z, _ := New("fake", SetLogz(logz), SetLevel(zapcore.DebugLevel))

	z.Debug("test")
	z.Sync()
	var sentMap map[string]string
	err := json.Unmarshal(sent, &sentMap)
	if err != nil {
		t.Fatal("error on json ", err)
	}
	time.Sleep(time.Millisecond * 200)
	assert.Equal(t, sentMap["type"], "zap-logger")
	assert.Equal(t, sentMap["level"], "debug")
	assert.Equal(t, sentMap["message"], "test")
	if len(sentMap["ts"]) == 0 {
		t.Fatal("wrong ts", sentMap["ts"])
	}
}

func TestType(t *testing.T) {
	var sent []byte
	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		defer r.Body.Close()
		sent, _ = ioutil.ReadAll(r.Body)
	}))
	ts.Start()
	defer ts.Close()
	logz, _ := logzio.New("fake", logzio.SetUrl(ts.URL))
	z, _ := New("fake", SetLogz(logz), SetLevel(zapcore.DebugLevel), SetType("tester"))

	z.Debug("test")
	z.Sync()
	var sentMap map[string]string
	err := json.Unmarshal(sent, &sentMap)
	if err != nil {
		t.Fatal("error on json ", err)
	}
	time.Sleep(time.Millisecond * 200)
	assert.Equal(t, sentMap["type"], "tester")
	assert.Equal(t, sentMap["level"], "debug")
	assert.Equal(t, sentMap["message"], "test")
	if len(sentMap["ts"]) == 0 {
		t.Fatal("wrong ts", sentMap["ts"])
	}
}

func TestConfig(t *testing.T) {

	var cfg = DefaultConfig
	cfg.TimeKey = "timestamp"
	var sent []byte
	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		defer r.Body.Close()
		sent, _ = ioutil.ReadAll(r.Body)
	}))
	ts.Start()
	defer ts.Close()
	logz, _ := logzio.New("fake", logzio.SetUrl(ts.URL))
	z, _ := New("fake", SetLogz(logz), SetEncodeConfig(cfg))

	z.Info("test")
	z.Sync()
	var sentMap map[string]string
	err := json.Unmarshal(sent, &sentMap)
	if err != nil {
		t.Fatal("error on json ", err)
	}
	time.Sleep(time.Millisecond * 200)
	assert.Equal(t, sentMap["type"], "zap-logger")
	assert.Equal(t, sentMap["level"], "info")
	assert.Equal(t, sentMap["message"], "test")
	t.Log("timestamp", sentMap["timestamp"])
	if len(sentMap["ts"]) > 0 {
		t.Fatal("wrong ts", sentMap["ts"])
	}

	if len(sentMap["timestamp"]) == 0 {
		t.Fatal("wrong timestamp")
	}
}

type debugWriter struct {
	dbg []byte
}

func (w *debugWriter) Write(p []byte) (int, error) {
	w.dbg = p
	return len(p), nil
}

func TestDebug(t *testing.T) {
	var sent []byte
	ts := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		defer r.Body.Close()
		sent, _ = ioutil.ReadAll(r.Body)
	}))
	ts.Start()
	defer ts.Close()
	var d = &debugWriter{
		dbg: make([]byte, 1024),
	}
	logz, _ := logzio.New("fake", logzio.SetUrl(ts.URL))
	z, _ := New("fake", SetLogz(logz), WithDebug(d), SetLevel(zapcore.DebugLevel), SetType("tester"))

	z.Debug("test")
	z.Sync()
	var sentMap map[string]string
	err := json.Unmarshal(sent, &sentMap)
	if err != nil {
		t.Fatal("error on json ", err)
	}
	time.Sleep(time.Millisecond * 200)
	assert.Equal(t, sentMap["type"], "tester")
	assert.Equal(t, sentMap["level"], "debug")
	assert.Equal(t, sentMap["message"], "test")
	if len(sentMap["ts"]) == 0 {
		t.Fatal("wrong ts", sentMap["ts"])
	}
}
