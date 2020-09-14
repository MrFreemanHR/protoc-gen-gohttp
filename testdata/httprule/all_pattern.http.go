// Code generated by protoc-gen-gohttp. DO NOT EDIT.
// source: httprule/all_pattern.proto

package httprulepb

import (
	bytes "bytes"
	context "context"
	base64 "encoding/base64"
	fmt "fmt"
	jsonpb "github.com/golang/protobuf/jsonpb"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	ioutil "io/ioutil"
	mime "mime"
	http "net/http"
	strconv "strconv"
	strings "strings"
)

// AllPatternHTTPService is the server API for AllPattern service.
type AllPatternHTTPService interface {
	AllPattern(context.Context, *AllPatternRequest) (*AllPatternResponse, error)
}

// AllPatternHTTPConverter has a function to convert AllPatternHTTPService interface to http.HandlerFunc.
type AllPatternHTTPConverter struct {
	srv AllPatternHTTPService
}

// NewAllPatternHTTPConverter returns AllPatternHTTPConverter.
func NewAllPatternHTTPConverter(srv AllPatternHTTPService) *AllPatternHTTPConverter {
	return &AllPatternHTTPConverter{
		srv: srv,
	}
}

// AllPattern returns AllPatternHTTPService interface's AllPattern converted to http.HandlerFunc.
func (h *AllPatternHTTPConverter) AllPattern(cb func(ctx context.Context, w http.ResponseWriter, r *http.Request, arg, ret proto.Message, err error), interceptors ...grpc.UnaryServerInterceptor) http.HandlerFunc {
	if cb == nil {
		cb = func(ctx context.Context, w http.ResponseWriter, r *http.Request, arg, ret proto.Message, err error) {
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				p := status.New(codes.Unknown, err.Error()).Proto()
				switch contentType, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type")); contentType {
				case "application/protobuf", "application/x-protobuf":
					buf, err := proto.Marshal(p)
					if err != nil {
						return
					}
					if _, err := io.Copy(w, bytes.NewBuffer(buf)); err != nil {
						return
					}
				case "application/json":
					m := jsonpb.Marshaler{
						EnumsAsInts:  true,
						EmitDefaults: true,
					}
					if err := m.Marshal(w, p); err != nil {
						return
					}
				default:
				}
			}
		}
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		contentType, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))

		accepts := strings.Split(r.Header.Get("Accept"), ",")
		accept := accepts[0]
		if accept == "*/*" || accept == "" {
			if contentType != "" {
				accept = contentType
			} else {
				accept = "application/json"
			}
		}

		w.Header().Set("Content-Type", accept)

		arg := &AllPatternRequest{}
		if r.Method != http.MethodGet {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				cb(ctx, w, r, nil, nil, err)
				return
			}

			switch contentType {
			case "application/protobuf", "application/x-protobuf":
				if err := proto.Unmarshal(body, arg); err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
			case "application/json":
				if err := jsonpb.Unmarshal(bytes.NewBuffer(body), arg); err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
			default:
				w.WriteHeader(http.StatusUnsupportedMediaType)
				_, err := fmt.Fprintf(w, "Unsupported Content-Type: %s", contentType)
				cb(ctx, w, r, nil, nil, err)
				return
			}
		}

		n := len(interceptors)
		chained := func(ctx context.Context, arg interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			chainer := func(currentInter grpc.UnaryServerInterceptor, currentHandler grpc.UnaryHandler) grpc.UnaryHandler {
				return func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
					return currentInter(currentCtx, currentReq, info, currentHandler)
				}
			}

			chainedHandler := handler
			for i := n - 1; i >= 0; i-- {
				chainedHandler = chainer(interceptors[i], chainedHandler)
			}
			return chainedHandler(ctx, arg)
		}

		info := &grpc.UnaryServerInfo{
			Server:     h.srv,
			FullMethod: "/httprule.AllPattern/AllPattern",
		}

		handler := func(c context.Context, req interface{}) (interface{}, error) {
			return h.srv.AllPattern(c, req.(*AllPatternRequest))
		}

		iret, err := chained(ctx, arg, info, handler)
		if err != nil {
			cb(ctx, w, r, arg, nil, err)
			return
		}

		ret, ok := iret.(*AllPatternResponse)
		if !ok {
			cb(ctx, w, r, arg, nil, fmt.Errorf("/httprule.AllPattern/AllPattern: interceptors have not return AllPatternResponse"))
			return
		}

		switch accept {
		case "application/protobuf", "application/x-protobuf":
			buf, err := proto.Marshal(ret)
			if err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
			if _, err := io.Copy(w, bytes.NewBuffer(buf)); err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
		case "application/json":
			m := jsonpb.Marshaler{
				EnumsAsInts:  true,
				EmitDefaults: true,
			}
			if err := m.Marshal(w, ret); err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
		default:
			w.WriteHeader(http.StatusUnsupportedMediaType)
			_, err := fmt.Fprintf(w, "Unsupported Accept: %s", accept)
			cb(ctx, w, r, arg, ret, err)
			return
		}
		cb(ctx, w, r, arg, ret, nil)
	})
}

// AllPatternWithName returns Service name, Method name and AllPatternHTTPService interface's AllPattern converted to http.HandlerFunc.
func (h *AllPatternHTTPConverter) AllPatternWithName(cb func(ctx context.Context, w http.ResponseWriter, r *http.Request, arg, ret proto.Message, err error), interceptors ...grpc.UnaryServerInterceptor) (string, string, http.HandlerFunc) {
	return "AllPattern", "AllPattern", h.AllPattern(cb, interceptors...)
}

// AllPatternHTTPRule returns HTTP method, path and AllPatternHTTPService interface's AllPattern converted to http.HandlerFunc.
func (h *AllPatternHTTPConverter) AllPatternHTTPRule(cb func(ctx context.Context, w http.ResponseWriter, r *http.Request, arg, ret proto.Message, err error), interceptors ...grpc.UnaryServerInterceptor) (string, string, http.HandlerFunc) {
	if cb == nil {
		cb = func(ctx context.Context, w http.ResponseWriter, r *http.Request, arg, ret proto.Message, err error) {
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				p := status.New(codes.Unknown, err.Error()).Proto()
				switch contentType, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type")); contentType {
				case "application/protobuf", "application/x-protobuf":
					buf, err := proto.Marshal(p)
					if err != nil {
						return
					}
					if _, err := io.Copy(w, bytes.NewBuffer(buf)); err != nil {
						return
					}
				case "application/json":
					m := jsonpb.Marshaler{
						EnumsAsInts:  true,
						EmitDefaults: true,
					}
					if err := m.Marshal(w, p); err != nil {
						return
					}
				default:
				}
			}
		}
	}
	return http.MethodGet, "/all/pattern", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		contentType, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))

		accepts := strings.Split(r.Header.Get("Accept"), ",")
		accept := accepts[0]
		if accept == "*/*" || accept == "" {
			if contentType != "" {
				accept = contentType
			} else {
				accept = "application/json"
			}
		}

		w.Header().Set("Content-Type", accept)

		arg := &AllPatternRequest{}
		if r.Method == http.MethodGet {
			if v := r.URL.Query().Get("double"); v != "" {
				c, err := strconv.ParseFloat(v, 64)
				if err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
				arg.Double = c
			}
			if v := r.URL.Query().Get("float"); v != "" {
				c, err := strconv.ParseFloat(v, 32)
				if err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
				arg.Float = float32(c)
			}
			if v := r.URL.Query().Get("int32"); v != "" {
				c, err := strconv.ParseInt(v, 10, 32)
				if err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
				arg.Int32 = int32(c)
			}
			if v := r.URL.Query().Get("int64"); v != "" {
				c, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
				arg.Int64 = c
			}
			if v := r.URL.Query().Get("uint32"); v != "" {
				c, err := strconv.ParseUint(v, 10, 32)
				if err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
				arg.Uint32 = uint32(c)
			}
			if v := r.URL.Query().Get("uint64"); v != "" {
				c, err := strconv.ParseUint(v, 10, 64)
				if err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
				arg.Uint64 = c
			}
			if v := r.URL.Query().Get("fixed32"); v != "" {
				c, err := strconv.ParseUint(v, 10, 32)
				if err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
				arg.Fixed32 = uint32(c)
			}
			if v := r.URL.Query().Get("fixed64"); v != "" {
				c, err := strconv.ParseUint(v, 10, 64)
				if err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
				arg.Fixed64 = c
			}
			if v := r.URL.Query().Get("sfixed32"); v != "" {
				c, err := strconv.ParseInt(v, 10, 32)
				if err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
				arg.Sfixed32 = int32(c)
			}
			if v := r.URL.Query().Get("sfixed64"); v != "" {
				c, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
				arg.Sfixed64 = c
			}
			if v := r.URL.Query().Get("bool"); v != "" {
				c, err := strconv.ParseBool(v)
				if err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
				arg.Bool = c
			}
			if v := r.URL.Query().Get("string"); v != "" {
				arg.String_ = v
			}
			if v := r.URL.Query().Get("bytes"); v != "" {
				c, err := base64.StdEncoding.DecodeString(v)
				if err != nil {
					cb(ctx, w, r, nil, nil, err)
					return
				}
				arg.Bytes = c
			}
			if repeated := r.URL.Query()["repeated_double"]; len(repeated) != 0 {
				arr := make([]float64, 0, len(repeated))
				for _, v := range repeated {
					c, err := strconv.ParseFloat(v, 64)
					if err != nil {
						cb(ctx, w, r, nil, nil, err)
						return
					}
					arr = append(arr, c)
				}
				arg.RepeatedDouble = arr
			}
			if repeated := r.URL.Query()["repeated_float"]; len(repeated) != 0 {
				arr := make([]float32, 0, len(repeated))
				for _, v := range repeated {
					c, err := strconv.ParseFloat(v, 32)
					if err != nil {
						cb(ctx, w, r, nil, nil, err)
						return
					}
					arr = append(arr, float32(c))
				}
				arg.RepeatedFloat = arr
			}
			if repeated := r.URL.Query()["repeated_int32"]; len(repeated) != 0 {
				arr := make([]int32, 0, len(repeated))
				for _, v := range repeated {
					c, err := strconv.ParseInt(v, 10, 32)
					if err != nil {
						cb(ctx, w, r, nil, nil, err)
						return
					}
					arr = append(arr, int32(c))
				}
				arg.RepeatedInt32 = arr
			}
			if repeated := r.URL.Query()["repeated_int64"]; len(repeated) != 0 {
				arr := make([]int64, 0, len(repeated))
				for _, v := range repeated {
					c, err := strconv.ParseInt(v, 10, 64)
					if err != nil {
						cb(ctx, w, r, nil, nil, err)
						return
					}
					arr = append(arr, c)
				}
				arg.RepeatedInt64 = arr
			}
			if repeated := r.URL.Query()["repeated_uint32"]; len(repeated) != 0 {
				arr := make([]uint32, 0, len(repeated))
				for _, v := range repeated {
					c, err := strconv.ParseUint(v, 10, 32)
					if err != nil {
						cb(ctx, w, r, nil, nil, err)
						return
					}
					arr = append(arr, uint32(c))
				}
				arg.RepeatedUint32 = arr
			}
			if repeated := r.URL.Query()["repeated_uint64"]; len(repeated) != 0 {
				arr := make([]uint64, 0, len(repeated))
				for _, v := range repeated {
					c, err := strconv.ParseUint(v, 10, 64)
					if err != nil {
						cb(ctx, w, r, nil, nil, err)
						return
					}
					arr = append(arr, c)
				}
				arg.RepeatedUint64 = arr
			}
			if repeated := r.URL.Query()["repeated_fixed32"]; len(repeated) != 0 {
				arr := make([]uint32, 0, len(repeated))
				for _, v := range repeated {
					c, err := strconv.ParseUint(v, 10, 32)
					if err != nil {
						cb(ctx, w, r, nil, nil, err)
						return
					}
					arr = append(arr, uint32(c))
				}
				arg.RepeatedFixed32 = arr
			}
			if repeated := r.URL.Query()["repeated_fixed64"]; len(repeated) != 0 {
				arr := make([]uint64, 0, len(repeated))
				for _, v := range repeated {
					c, err := strconv.ParseUint(v, 10, 64)
					if err != nil {
						cb(ctx, w, r, nil, nil, err)
						return
					}
					arr = append(arr, c)
				}
				arg.RepeatedFixed64 = arr
			}
			if repeated := r.URL.Query()["repeated_sfixed32"]; len(repeated) != 0 {
				arr := make([]int32, 0, len(repeated))
				for _, v := range repeated {
					c, err := strconv.ParseInt(v, 10, 32)
					if err != nil {
						cb(ctx, w, r, nil, nil, err)
						return
					}
					arr = append(arr, int32(c))
				}
				arg.RepeatedSfixed32 = arr
			}
			if repeated := r.URL.Query()["repeated_sfixed64"]; len(repeated) != 0 {
				arr := make([]int64, 0, len(repeated))
				for _, v := range repeated {
					c, err := strconv.ParseInt(v, 10, 64)
					if err != nil {
						cb(ctx, w, r, nil, nil, err)
						return
					}
					arr = append(arr, c)
				}
				arg.RepeatedSfixed64 = arr
			}
			if repeated := r.URL.Query()["repeated_bool"]; len(repeated) != 0 {
				arr := make([]bool, 0, len(repeated))
				for _, v := range repeated {
					c, err := strconv.ParseBool(v)
					if err != nil {
						cb(ctx, w, r, nil, nil, err)
						return
					}
					arr = append(arr, c)
				}
				arg.RepeatedBool = arr
			}
			if repeated := r.URL.Query()["repeated_string"]; len(repeated) != 0 {
				arr := make([]string, 0, len(repeated))
				for _, v := range repeated {
					arr = append(arr, v)
				}
				arg.RepeatedString = arr
			}
			if repeated := r.URL.Query()["repeated_bytes"]; len(repeated) != 0 {
				arr := make([][]byte, 0, len(repeated))
				for _, v := range repeated {
					c, err := base64.StdEncoding.DecodeString(v)
					if err != nil {
						cb(ctx, w, r, nil, nil, err)
						return
					}
					arr = append(arr, c)
				}
				arg.RepeatedBytes = arr
			}
		}

		n := len(interceptors)
		chained := func(ctx context.Context, arg interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			chainer := func(currentInter grpc.UnaryServerInterceptor, currentHandler grpc.UnaryHandler) grpc.UnaryHandler {
				return func(currentCtx context.Context, currentReq interface{}) (interface{}, error) {
					return currentInter(currentCtx, currentReq, info, currentHandler)
				}
			}

			chainedHandler := handler
			for i := n - 1; i >= 0; i-- {
				chainedHandler = chainer(interceptors[i], chainedHandler)
			}
			return chainedHandler(ctx, arg)
		}

		info := &grpc.UnaryServerInfo{
			Server:     h.srv,
			FullMethod: "/httprule.AllPattern/AllPattern",
		}

		handler := func(c context.Context, req interface{}) (interface{}, error) {
			return h.srv.AllPattern(c, req.(*AllPatternRequest))
		}

		iret, err := chained(ctx, arg, info, handler)
		if err != nil {
			cb(ctx, w, r, arg, nil, err)
			return
		}

		ret, ok := iret.(*AllPatternResponse)
		if !ok {
			cb(ctx, w, r, arg, nil, fmt.Errorf("/httprule.AllPattern/AllPattern: interceptors have not return AllPatternResponse"))
			return
		}

		switch accept {
		case "application/protobuf", "application/x-protobuf":
			buf, err := proto.Marshal(ret)
			if err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
			if _, err := io.Copy(w, bytes.NewBuffer(buf)); err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
		case "application/json":
			m := jsonpb.Marshaler{
				EnumsAsInts:  true,
				EmitDefaults: true,
			}
			if err := m.Marshal(w, ret); err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
		default:
			w.WriteHeader(http.StatusUnsupportedMediaType)
			_, err := fmt.Fprintf(w, "Unsupported Accept: %s", accept)
			cb(ctx, w, r, arg, ret, err)
			return
		}
		cb(ctx, w, r, arg, ret, nil)
	})
}
