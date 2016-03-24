// Copyright 2015 flannel authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package remote

import (
	glog "github.com/coreos/flannel/Godeps/_workspace/src/github.com/golang/glog"
	"net/http"
)

type httpResp struct {
	writer http.ResponseWriter
	status int
}

func (r *httpResp) Header() http.Header {
	return r.writer.Header()
}

func (r *httpResp) Write(d []byte) (int, error) {
	return r.writer.Write(d)
}

func (r *httpResp) WriteHeader(status int) {
	r.status = status
	r.writer.WriteHeader(status)
}

type httpLoggerHandler struct {
	h http.Handler
}

func (lh httpLoggerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp := &httpResp{w, 0}
	lh.h.ServeHTTP(resp, r)
	glog.Infof("%v %v - %v", r.Method, r.RequestURI, resp.status)
}

func httpLogger(h http.Handler) http.Handler {
	return httpLoggerHandler{h}
}
