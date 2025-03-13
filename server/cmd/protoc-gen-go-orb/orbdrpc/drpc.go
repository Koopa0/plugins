// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

// Package orbdrpc generates DRPC code for protobuf services.
//
//nolint:lll,funlen,wsl,dupword
package orbdrpc

import (
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

// Version will be set by orb/main.
var Version string //nolint:gochecknoglobals

// Config will be set by orb/main.
type Config struct {
	Protolib string
	JSON     bool
}

func protocVersion(gen *protogen.Plugin) string {
	v := gen.Request.GetCompilerVersion()
	if v == nil {
		return "(unknown)"
	}

	var suffix string

	if s := v.GetSuffix(); s != "" {
		suffix = "-" + s
	}

	return fmt.Sprintf("v%d.%d.%d%s", v.GetMajor(), v.GetMinor(), v.GetPatch(), suffix)
}

// GenerateFile writes the file.
func GenerateFile(plugin *protogen.Plugin, file *protogen.File, conf Config) {
	gf := plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+"_orb-drpc.pb.go", file.GoImportPath)
	d := &drpc{gf, file}

	// Generate header message.
	d.P("// Code generated by protoc-gen-go-orb. DO NOT EDIT.")
	d.P("//")
	d.P("// version:")
	d.P("// - protoc-gen-go-orb        v" + Version)
	d.P("// - protoc                   ", protocVersion(plugin))
	d.P("//")
	d.P("// source: ", file.Desc.Path())
	d.P()
	d.P("package ", file.GoPackageName)
	d.P()

	d.generateEncoding(conf)

	for _, service := range file.Services {
		d.generateService(service)
	}
}

type drpc struct {
	*protogen.GeneratedFile
	file *protogen.File
}

//
// name helpers
//

func (d *drpc) Ident(path, ident string) string {
	return d.QualifiedGoIdent(protogen.GoImportPath(path).Ident(ident))
}

func (d *drpc) EncodingName() string {
	return "drpcEncoding_" + d.file.GoDescriptorIdent.GoName
}

func (d *drpc) RPCGoString(method *protogen.Method) string {
	return strconv.Quote(fmt.Sprintf("/%s/%s", method.Parent.Desc.FullName(), method.Desc.Name()))
}

func (d *drpc) InputType(method *protogen.Method) string {
	return d.QualifiedGoIdent(method.Input.GoIdent)
}

func (d *drpc) OutputType(method *protogen.Method) string {
	return d.QualifiedGoIdent(method.Output.GoIdent)
}

func (d *drpc) ServerIface(service *protogen.Service) string {
	return "DRPC" + service.GoName + "Server"
}

func (d *drpc) ServerImpl(service *protogen.Service) string {
	return "drpc" + service.GoName + "Server"
}

func (d *drpc) ServerUnimpl(service *protogen.Service) string {
	return "DRPC" + service.GoName + "UnimplementedServer"
}

func (d *drpc) ServerDesc(service *protogen.Service) string {
	return "DRPC" + service.GoName + "Description"
}

func (d *drpc) ClientStreamIface(method *protogen.Method) string {
	return "DRPC" +
		strings.ReplaceAll(method.Parent.GoName, "_", "__") + "_" +
		strings.ReplaceAll(method.GoName, "_", "__") +
		"Client"
}

func (d *drpc) ClientStreamImpl(method *protogen.Method) string {
	return "drpc" +
		strings.ReplaceAll(method.Parent.GoName, "_", "__") + "_" +
		strings.ReplaceAll(method.GoName, "_", "__") +
		"Client"
}

func (d *drpc) ServerStreamIface(method *protogen.Method) string {
	return "DRPC" +
		strings.ReplaceAll(method.Parent.GoName, "_", "__") + "_" +
		strings.ReplaceAll(method.GoName, "_", "__") +
		"Stream"
}

func (d *drpc) ServerStreamImpl(method *protogen.Method) string {
	return "drpc" +
		strings.ReplaceAll(method.Parent.GoName, "_", "__") + "_" +
		strings.ReplaceAll(method.GoName, "_", "__") +
		"Stream"
}

//
// encoding generation
//

func (d *drpc) generateEncoding(conf Config) {
	d.P("type ", d.EncodingName(), " struct{}")
	d.P()

	switch conf.Protolib {
	case "google.golang.org/protobuf":
		d.P("func (", d.EncodingName(), ") Marshal(msg ", d.Ident("storj.io/drpc", "Message"), ") ([]byte, error) {")
		d.P("return ", d.Ident("google.golang.org/protobuf/proto", "Marshal"), "(msg.(", d.Ident("google.golang.org/protobuf/proto", "Message"), "))")
		d.P("}")
		d.P()

		d.P("func (", d.EncodingName(), ") MarshalAppend(buf []byte, msg ", d.Ident("storj.io/drpc", "Message"), ") ([]byte, error) {")
		d.P("return ", d.Ident("google.golang.org/protobuf/proto", "MarshalOptions"), "{}.MarshalAppend(buf, msg.(", d.Ident("google.golang.org/protobuf/proto", "Message"), "))")
		d.P("}")
		d.P()

		d.P("func (", d.EncodingName(), ") Unmarshal(buf []byte, msg ", d.Ident("storj.io/drpc", "Message"), ") error {")
		d.P("return ", d.Ident("google.golang.org/protobuf/proto", "Unmarshal"), "(buf, msg.(", d.Ident("google.golang.org/protobuf/proto", "Message"), "))")
		d.P("}")
		d.P()

		if conf.JSON {
			d.P("func (", d.EncodingName(), ") JSONMarshal(msg ", d.Ident("storj.io/drpc", "Message"), ") ([]byte, error) {")
			d.P("return ", d.Ident("google.golang.org/protobuf/encoding/protojson", "Marshal"), "(msg.(", d.Ident("google.golang.org/protobuf/proto", "Message"), "))")
			d.P("}")
			d.P()

			d.P("func (", d.EncodingName(), ") JSONUnmarshal(buf []byte, msg ", d.Ident("storj.io/drpc", "Message"), ") error {")
			d.P("return ", d.Ident("google.golang.org/protobuf/encoding/protojson", "Unmarshal"), "(buf, msg.(", d.Ident("google.golang.org/protobuf/proto", "Message"), "))")
			d.P("}")
			d.P()
		}

	case "github.com/gogo/protobuf":
		d.P("func (", d.EncodingName(), ") Marshal(msg ", d.Ident("storj.io/drpc", "Message"), ") ([]byte, error) {")
		d.P("return ", d.Ident("github.com/gogo/protobuf/proto", "Marshal"), "(msg.(", d.Ident("github.com/gogo/protobuf/proto", "Message"), "))")
		d.P("}")
		d.P()

		d.P("func (", d.EncodingName(), ") Unmarshal(buf []byte, msg ", d.Ident("storj.io/drpc", "Message"), ") error {")
		d.P("return ", d.Ident("github.com/gogo/protobuf/proto", "Unmarshal"), "(buf, msg.(", d.Ident("github.com/gogo/protobuf/proto", "Message"), "))")
		d.P("}")
		d.P()

		if conf.JSON {
			d.P("func (", d.EncodingName(), ") JSONMarshal(msg ", d.Ident("storj.io/drpc", "Message"), ") ([]byte, error) {")
			d.P("var buf ", d.Ident("bytes", "Buffer"))
			d.P("err := new(", d.Ident("github.com/gogo/protobuf/jsonpb", "Marshaler"), ").Marshal(&buf, msg.(", d.Ident("github.com/gogo/protobuf/proto", "Message"), "))")
			d.P("if err != nil {")
			d.P("return nil, err")
			d.P("}")
			d.P("return buf.Bytes(), nil")
			d.P("}")
			d.P()

			d.P("func (", d.EncodingName(), ") JSONUnmarshal(buf []byte, msg ", d.Ident("storj.io/drpc", "Message"), ") error {")
			d.P("return ", d.Ident("github.com/gogo/protobuf/jsonpb", "Unmarshal"), "(", d.Ident("bytes", "NewReader"), "(buf), msg.(", d.Ident("github.com/gogo/protobuf/proto", "Message"), "))")
			d.P("}")
			d.P()
		}

	default:
		d.P("func (", d.EncodingName(), ") Marshal(msg ", d.Ident("storj.io/drpc", "Message"), ") ([]byte, error) {")
		d.P("return ", d.Ident(conf.Protolib, "Marshal"), "(msg)")
		d.P("}")
		d.P()

		d.P("func (", d.EncodingName(), ") Unmarshal(buf []byte, msg ", d.Ident("storj.io/drpc", "Message"), ") error {")
		d.P("return ", d.Ident(conf.Protolib, "Unmarshal"), "(buf, msg)")
		d.P("}")
		d.P()

		if conf.JSON {
			d.P("func (", d.EncodingName(), ") JSONMarshal(msg ", d.Ident("storj.io/drpc", "Message"), ") ([]byte, error) {")
			d.P("return ", d.Ident(conf.Protolib, "JSONMarshal"), "(msg)")
			d.P("}")
			d.P()

			d.P("func (", d.EncodingName(), ") JSONUnmarshal(buf []byte, msg ", d.Ident("storj.io/drpc", "Message"), ") error {")
			d.P("return ", d.Ident(conf.Protolib, "JSONUnmarshal"), "(buf, msg)")
			d.P("}")
			d.P()
		}
	}
}

//
// service generation
//

func (d *drpc) generateService(service *protogen.Service) {
	// Server interface
	d.P("type ", d.ServerIface(service), " interface {")
	for _, method := range service.Methods {
		d.P(d.generateServerSignature(method))
	}
	d.P("}")
	d.P()

	// Server Unimplemented struct
	d.P("type ", d.ServerUnimpl(service), " struct {}")
	d.P()
	for _, method := range service.Methods {
		d.generateUnimplementedServerMethod(method)
	}
	d.P()

	// Server description.
	d.P("type ", d.ServerDesc(service), " struct{}")
	d.P()
	d.P("func (", d.ServerDesc(service), ") NumMethods() int { return ", len(service.Methods), " }")
	d.P()
	d.P("func (", d.ServerDesc(service), ") Method(n int) (string, ", d.Ident("storj.io/drpc", "Encoding"), ", ", d.Ident("storj.io/drpc", "Receiver"), ", interface{}, bool) {")
	d.P("switch n {")
	for i, method := range service.Methods {
		d.P("case ", i, ":")
		d.P("return ", d.RPCGoString(method), ", ", d.EncodingName(), "{}, ")
		d.generateServerReceiver(method)
		d.P("}, ", d.ServerIface(service), ".", method.GoName, ", true")
	}
	d.P("default:")
	d.P(`return "", nil, nil, nil, false`)
	d.P("}")
	d.P("}")
	d.P()

	// Server methods
	for _, method := range service.Methods {
		d.generateServerMethod(method)
	}
}

//
// server methods
//

func (d *drpc) generateServerSignature(method *protogen.Method) string {
	var reqArgs []string
	ret := "error"
	if !method.Desc.IsStreamingServer() && !method.Desc.IsStreamingClient() {
		reqArgs = append(reqArgs, d.Ident("context", "Context"))
		ret = "(*" + d.OutputType(method) + ", error)"
	}
	if !method.Desc.IsStreamingClient() {
		reqArgs = append(reqArgs, "*"+d.InputType(method))
	}
	if method.Desc.IsStreamingServer() || method.Desc.IsStreamingClient() {
		reqArgs = append(reqArgs, d.ServerStreamIface(method))
	}
	return method.GoName + "(" + strings.Join(reqArgs, ", ") + ") " + ret
}

func (d *drpc) generateUnimplementedServerMethod(method *protogen.Method) {
	d.P("func (s *", d.ServerUnimpl(method.Parent), ") ", d.generateServerSignature(method), " {")
	if !method.Desc.IsStreamingServer() && !method.Desc.IsStreamingClient() {
		d.P("return nil, ", d.Ident("storj.io/drpc/drpcerr", "WithCode"), "(", d.Ident("errors", "New"), "(\"Unimplemented\"), ", d.Ident("storj.io/drpc/drpcerr", "Unimplemented"), ")")
	} else {
		d.P("return ", d.Ident("storj.io/drpc/drpcerr", "WithCode"), "(", d.Ident("errors", "New"), "(\"Unimplemented\"), ", d.Ident("storj.io/drpc/drpcerr", "Unimplemented"), ")")
	}
	d.P("}")
	d.P()
}

func (d *drpc) generateServerReceiver(method *protogen.Method) {
	d.P("func (srv interface{}, ctx " + d.Ident("context", "Context") + ", in1, in2 interface{}) (" + d.Ident("storj.io/drpc", "Message") + ", error) {")
	if !method.Desc.IsStreamingServer() && !method.Desc.IsStreamingClient() {
		d.P("return srv.(", d.ServerIface(method.Parent), ").")
	} else {
		d.P("return nil, srv.(", d.ServerIface(method.Parent), ").")
	}
	d.P(method.GoName, "(")

	n := 1
	if !method.Desc.IsStreamingServer() && !method.Desc.IsStreamingClient() {
		d.P("ctx,")
	}
	if !method.Desc.IsStreamingClient() {
		d.P("in", n, ".(*", d.InputType(method), "),")
		n++
	}
	if method.Desc.IsStreamingServer() || method.Desc.IsStreamingClient() {
		d.P("&", d.ServerStreamImpl(method), "{in", n, ".(", d.Ident("storj.io/drpc", "Stream"), ")},")
	}
	d.P(")")
}

func (d *drpc) generateServerMethod(method *protogen.Method) {
	genSend := method.Desc.IsStreamingServer()
	genSendAndClose := !method.Desc.IsStreamingServer()
	genRecv := method.Desc.IsStreamingClient()

	// Stream auxiliary types and methods.
	d.P("type ", d.ServerStreamIface(method), " interface {")
	d.P(d.Ident("storj.io/drpc", "Stream"))
	if genSend {
		d.P("Send(*", d.OutputType(method), ") error")
	}
	if genSendAndClose {
		d.P("SendAndClose(*", d.OutputType(method), ") error")
	}
	if genRecv {
		d.P("Recv() (*", d.InputType(method), ", error)")
	}
	d.P("}")
	d.P()

	d.P("type ", d.ServerStreamImpl(method), " struct {")
	d.P(d.Ident("storj.io/drpc", "Stream"))
	d.P("}")
	d.P()

	d.P("func (x *", d.ServerStreamImpl(method), ") GetStream() drpc.Stream {")
	d.P("return x.Stream")
	d.P("}")
	d.P()

	if genSend {
		d.P("func (x *", d.ServerStreamImpl(method), ") Send(m *", d.OutputType(method), ") error {")
		d.P("return x.MsgSend(m, ", d.EncodingName(), "{})")
		d.P("}")
		d.P()
	}

	if genSendAndClose {
		d.P("func (x *", d.ServerStreamImpl(method), ") SendAndClose(m *", d.OutputType(method), ") error {")
		d.P("if err := x.MsgSend(m, ", d.EncodingName(), "{}); err != nil { return err }")
		d.P("return x.CloseSend()")
		d.P("}")
		d.P()
	}

	if genRecv {
		d.P("func (x *", d.ServerStreamImpl(method), ") Recv() (*", d.InputType(method), ", error) {")
		d.P("m := new(", d.InputType(method), ")")
		d.P("if err := x.MsgRecv(m, ", d.EncodingName(), "{}); err != nil { return nil, err }")
		d.P("return m, nil")
		d.P("}")
		d.P()

		d.P("func (x *", d.ServerStreamImpl(method), ") RecvMsg(m *", d.InputType(method), ") error {")
		d.P("return x.MsgRecv(m, ", d.EncodingName(), "{})")
		d.P("}")
		d.P()
	}
}
