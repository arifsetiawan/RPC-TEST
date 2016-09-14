package greeter

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

type Greeter struct{ Client capnp.Client }

func (c Greeter) SayHello(ctx context.Context, params func(Greeter_sayHello_Params) error, opts ...capnp.CallOption) Greeter_sayHello_Results_Promise {
	if c.Client == nil {
		return Greeter_sayHello_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xe63e3f25a21bee05,
			MethodID:      0,
			InterfaceName: "helloworld.capnp:Greeter",
			MethodName:    "sayHello",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Greeter_sayHello_Params{Struct: s}) }
	}
	return Greeter_sayHello_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Greeter_Server interface {
	SayHello(Greeter_sayHello) error
}

func Greeter_ServerToClient(s Greeter_Server) Greeter {
	c, _ := s.(server.Closer)
	return Greeter{Client: server.New(Greeter_Methods(nil, s), c)}
}

func Greeter_Methods(methods []server.Method, s Greeter_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe63e3f25a21bee05,
			MethodID:      0,
			InterfaceName: "helloworld.capnp:Greeter",
			MethodName:    "sayHello",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Greeter_sayHello{c, opts, Greeter_sayHello_Params{Struct: p}, Greeter_sayHello_Results{Struct: r}}
			return s.SayHello(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// Greeter_sayHello holds the arguments for a server call to Greeter.sayHello.
type Greeter_sayHello struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Greeter_sayHello_Params
	Results Greeter_sayHello_Results
}

type Greeter_sayHello_Params struct{ capnp.Struct }

// Greeter_sayHello_Params_TypeID is the unique identifier for the type Greeter_sayHello_Params.
const Greeter_sayHello_Params_TypeID = 0xd23278f1d59666e2

func NewGreeter_sayHello_Params(s *capnp.Segment) (Greeter_sayHello_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Greeter_sayHello_Params{st}, err
}

func NewRootGreeter_sayHello_Params(s *capnp.Segment) (Greeter_sayHello_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Greeter_sayHello_Params{st}, err
}

func ReadRootGreeter_sayHello_Params(msg *capnp.Message) (Greeter_sayHello_Params, error) {
	root, err := msg.RootPtr()
	return Greeter_sayHello_Params{root.Struct()}, err
}

func (s Greeter_sayHello_Params) String() string {
	str, _ := text.Marshal(0xd23278f1d59666e2, s.Struct)
	return str
}

func (s Greeter_sayHello_Params) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Greeter_sayHello_Params) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Greeter_sayHello_Params) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Greeter_sayHello_Params) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// Greeter_sayHello_Params_List is a list of Greeter_sayHello_Params.
type Greeter_sayHello_Params_List struct{ capnp.List }

// NewGreeter_sayHello_Params creates a new list of Greeter_sayHello_Params.
func NewGreeter_sayHello_Params_List(s *capnp.Segment, sz int32) (Greeter_sayHello_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Greeter_sayHello_Params_List{l}, err
}

func (s Greeter_sayHello_Params_List) At(i int) Greeter_sayHello_Params {
	return Greeter_sayHello_Params{s.List.Struct(i)}
}

func (s Greeter_sayHello_Params_List) Set(i int, v Greeter_sayHello_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Greeter_sayHello_Params_Promise is a wrapper for a Greeter_sayHello_Params promised by a client call.
type Greeter_sayHello_Params_Promise struct{ *capnp.Pipeline }

func (p Greeter_sayHello_Params_Promise) Struct() (Greeter_sayHello_Params, error) {
	s, err := p.Pipeline.Struct()
	return Greeter_sayHello_Params{s}, err
}

type Greeter_sayHello_Results struct{ capnp.Struct }

// Greeter_sayHello_Results_TypeID is the unique identifier for the type Greeter_sayHello_Results.
const Greeter_sayHello_Results_TypeID = 0xabd39b474be6e663

func NewGreeter_sayHello_Results(s *capnp.Segment) (Greeter_sayHello_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Greeter_sayHello_Results{st}, err
}

func NewRootGreeter_sayHello_Results(s *capnp.Segment) (Greeter_sayHello_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Greeter_sayHello_Results{st}, err
}

func ReadRootGreeter_sayHello_Results(msg *capnp.Message) (Greeter_sayHello_Results, error) {
	root, err := msg.RootPtr()
	return Greeter_sayHello_Results{root.Struct()}, err
}

func (s Greeter_sayHello_Results) String() string {
	str, _ := text.Marshal(0xabd39b474be6e663, s.Struct)
	return str
}

func (s Greeter_sayHello_Results) Rep() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Greeter_sayHello_Results) HasRep() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Greeter_sayHello_Results) RepBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Greeter_sayHello_Results) SetRep(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// Greeter_sayHello_Results_List is a list of Greeter_sayHello_Results.
type Greeter_sayHello_Results_List struct{ capnp.List }

// NewGreeter_sayHello_Results creates a new list of Greeter_sayHello_Results.
func NewGreeter_sayHello_Results_List(s *capnp.Segment, sz int32) (Greeter_sayHello_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Greeter_sayHello_Results_List{l}, err
}

func (s Greeter_sayHello_Results_List) At(i int) Greeter_sayHello_Results {
	return Greeter_sayHello_Results{s.List.Struct(i)}
}

func (s Greeter_sayHello_Results_List) Set(i int, v Greeter_sayHello_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Greeter_sayHello_Results_Promise is a wrapper for a Greeter_sayHello_Results promised by a client call.
type Greeter_sayHello_Results_Promise struct{ *capnp.Pipeline }

func (p Greeter_sayHello_Results_Promise) Struct() (Greeter_sayHello_Results, error) {
	s, err := p.Pipeline.Struct()
	return Greeter_sayHello_Results{s}, err
}

const schema_bbd09ee5a468f60c = "x\xda\x12\xe8w`2d\xcdgb`\x08\x94ae" +
	"\xfb\x9f\xfc\xec\x99\xb7\xfb\xec\xcb\xab\x19\x04%\x19\x19\x18" +
	"X\x19\xd9\x19\x18\x8ce\x19\x83\x18\x19\x18\x855\x19\xed" +
	"\x19\x18\xff?J\x9bv\xf5c\x85\xd1%d\x05\x9e\x8c" +
	"^ \x05\xa1`\x05\xac\xef\xa4\x17\xa9\xda\xdb=c\x10" +
	"\x14d\xfe\xcf\xf3-c\xc9\xd3y\x17v3\x00\xa5K" +
	"\x19O\x097\x82\xd4\x0b\xd72\xba\x0b\xaf\x04\xb2t\xfe" +
	"g\xa4\xe6\xe4\xe4\x97\xe7\x17\xb1\xe4\xa4\xe8%'\x16\xe4" +
	"\x15X\xb9\x17\xa5\xa6\x96\xa4\x16\xe9\x15'Vz\x80$" +
	"U\x82R\x8bKsJ\x18\x8b\x03Y\x98Y\x18\x18X" +
	"\x80\x96\x0a\xf2*\x01]\xcb\xc1\xcc\x18(\xc2\xc4\xc8^" +
	"\x94Z\xc0\xc8\xc3\xc0\x04\xc4\x8c\xc4\x18\x17\x90X\x94\x98" +
	"[\x0c\xd4\x8fd\x9c\x16\xc28\xfe\xbc\xc4\xdcT\x0c\xf3" +
	"\x98\xd0\xcdc`\x08`d\x04\x1a\xc1\xca\xc0\x00\x0f\x11" +
	"FX\xd8\x09\x0az10\x09r\xb2\xff\x87Y\xca\xc0" +
	"\xc0\xe0\xc0\x08\xd4\x00\x08\x00\x00\xff\xff\x1d2bS"

func init() {
	schemas.Register(schema_bbd09ee5a468f60c,
		0xabd39b474be6e663,
		0xd23278f1d59666e2,
		0xe63e3f25a21bee05)
}
