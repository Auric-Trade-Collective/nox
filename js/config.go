package js

import (
	// "net/http"
)

type Config struct {
	Root string `json:"root"`
	Ip string   `json:"ip"`
	Port string `json:"port"`
	// Forward map[string]func(req *HttpRequest) `json:"forward"`
	// Return map[string]func(req *HttpResponse) `json:"return"`
}

// type HttpRequest struct {
// 	request *http.Request
// 	additional []byte
// 	ContentLength int64
// }

// func (r *HttpRequest) ReadChunk(size int) ([]byte, error) {
// 	buf := make([]byte, size)
// 	n, err := r.request.Body.Read(buf)
// 	return buf[:n], err
// }
//
// func (r *HttpRequest) ReadString(size int) (string, error) {
// 	buf := make([]byte, size)
// 	n, err := r.request.Body.Read(buf)
// 	return string(buf[:n]), err
// }

// func (r *HttpRequest) GetBody() ([]byte, error) {
// 	buf := make([]byte, r.ContentLength)
// 	n, err := r.request.Body.Read(buf)
// 	return buf[:n], err
// }

// func (r *HttpRequest) GetHeader(key string) string {
// 	val := r.request.Header.Get(key)
// 	return val
// }
//
// func (r *HttpRequest) SetHeader(key string, val string) {
// 	r.request.Header.Set(key, val)
// }
//
// func (r *HttpRequest) RemoveHeader(key string) {
// 	r.request.Header.Del(key)
// }
//
// func (r *HttpRequest) WriteString(buf string) {
// 	r.additional = append(r.additional, []byte(buf)...)
// }
//
// type HttpResponse struct {
// 	wr http.ResponseWriter
// }
//
// func (r *HttpResponse) SetHeader(key string, val string) {
// 	r.wr.Header().Set(key, val)
// }
//
// func (r *HttpResponse) RemoveHeader(key string) {
// 	r.wr.Header().Del(key)
// }
//
// func (r *HttpResponse) WriteString(val string) {
// 	r.wr.Write([]byte(val))
// }
//
// func (r *HttpResponse) Write(val []byte) {
// 	r.wr.Write(val)
// }
