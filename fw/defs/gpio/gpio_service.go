// Code generated by clubbygen.
// GENERATED FILE DO NOT EDIT
// +build !clubby_strict

package gpio

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"cesanta.com/common/go/mgrpc"
	"cesanta.com/common/go/mgrpc/frame"
	"cesanta.com/common/go/ourjson"
	"cesanta.com/common/go/ourtrace"
	"github.com/cesanta/errors"
	"golang.org/x/net/trace"
)

var _ = bytes.MinRead
var _ = fmt.Errorf
var emptyMessage = ourjson.RawMessage{}
var _ = ourtrace.New
var _ = trace.New

const ServiceID = "http://mongoose-iot.com/fwGPIO"

type ReadArgs struct {
	Pin *int64 `json:"pin,omitempty"`
}

type ReadResult struct {
	Value *int64 `json:"value,omitempty"`
}

type RemoveIntHandlerArgs struct {
	Pin *int64 `json:"pin,omitempty"`
}

type SetIntHandlerArgs struct {
	Debounce_ms *int64  `json:"debounce_ms,omitempty"`
	Dst         *string `json:"dst,omitempty"`
	Edge        *string `json:"edge,omitempty"`
	Method      *string `json:"method,omitempty"`
	Pin         *int64  `json:"pin,omitempty"`
	Pull        *string `json:"pull,omitempty"`
}

type SetIntHandlerResult struct {
	Value *int64 `json:"value,omitempty"`
}

type ToggleArgs struct {
	Pin *int64 `json:"pin,omitempty"`
}

type ToggleResult struct {
	Value *int64 `json:"value,omitempty"`
}

type WriteArgs struct {
	Pin   *int64 `json:"pin,omitempty"`
	Value *int64 `json:"value,omitempty"`
}

type Service interface {
	Read(ctx context.Context, args *ReadArgs) (*ReadResult, error)
	RemoveIntHandler(ctx context.Context, args *RemoveIntHandlerArgs) error
	SetIntHandler(ctx context.Context, args *SetIntHandlerArgs) (*SetIntHandlerResult, error)
	Toggle(ctx context.Context, args *ToggleArgs) (*ToggleResult, error)
	Write(ctx context.Context, args *WriteArgs) error
}

type Instance interface {
	Call(context.Context, string, *frame.Command) (*frame.Response, error)
	//TraceCall(context.Context, string, *frame.Command) (context.Context, trace.Trace, func(*error))
}

func NewClient(i Instance, addr string) Service {
	return &_Client{i: i, addr: addr}
}

type _Client struct {
	i    Instance
	addr string
}

func (c *_Client) Read(ctx context.Context, args *ReadArgs) (res *ReadResult, err error) {
	cmd := &frame.Command{
		Cmd: "GPIO.Read",
	}
	//ctx, tr, finish := c.i.TraceCall(pctx, c.addr, cmd)
	//defer finish(&err)
	//_ = tr

	//tr.LazyPrintf("args: %s", ourjson.LazyJSON(&args))
	cmd.Args = ourjson.DelayMarshaling(args)
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if resp.Status != 0 {
		return nil, errors.Trace(&mgrpc.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}

	//tr.LazyPrintf("res: %s", ourjson.LazyJSON(&resp))

	var r *ReadResult
	err = resp.Response.UnmarshalInto(&r)
	if err != nil {
		return nil, errors.Annotatef(err, "unmarshaling response")
	}
	return r, nil
}

func (c *_Client) RemoveIntHandler(ctx context.Context, args *RemoveIntHandlerArgs) (err error) {
	cmd := &frame.Command{
		Cmd: "GPIO.RemoveIntHandler",
	}
	//ctx, tr, finish := c.i.TraceCall(pctx, c.addr, cmd)
	//defer finish(&err)
	//_ = tr

	//tr.LazyPrintf("args: %s", ourjson.LazyJSON(&args))
	cmd.Args = ourjson.DelayMarshaling(args)
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return errors.Trace(err)
	}
	if resp.Status != 0 {
		return errors.Trace(&mgrpc.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}
	return nil
}

func (c *_Client) SetIntHandler(ctx context.Context, args *SetIntHandlerArgs) (res *SetIntHandlerResult, err error) {
	cmd := &frame.Command{
		Cmd: "GPIO.SetIntHandler",
	}
	//ctx, tr, finish := c.i.TraceCall(pctx, c.addr, cmd)
	//defer finish(&err)
	//_ = tr

	//tr.LazyPrintf("args: %s", ourjson.LazyJSON(&args))
	cmd.Args = ourjson.DelayMarshaling(args)
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if resp.Status != 0 {
		return nil, errors.Trace(&mgrpc.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}

	//tr.LazyPrintf("res: %s", ourjson.LazyJSON(&resp))

	var r *SetIntHandlerResult
	err = resp.Response.UnmarshalInto(&r)
	if err != nil {
		return nil, errors.Annotatef(err, "unmarshaling response")
	}
	return r, nil
}

func (c *_Client) Toggle(ctx context.Context, args *ToggleArgs) (res *ToggleResult, err error) {
	cmd := &frame.Command{
		Cmd: "GPIO.Toggle",
	}
	//ctx, tr, finish := c.i.TraceCall(pctx, c.addr, cmd)
	//defer finish(&err)
	//_ = tr

	//tr.LazyPrintf("args: %s", ourjson.LazyJSON(&args))
	cmd.Args = ourjson.DelayMarshaling(args)
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if resp.Status != 0 {
		return nil, errors.Trace(&mgrpc.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}

	//tr.LazyPrintf("res: %s", ourjson.LazyJSON(&resp))

	var r *ToggleResult
	err = resp.Response.UnmarshalInto(&r)
	if err != nil {
		return nil, errors.Annotatef(err, "unmarshaling response")
	}
	return r, nil
}

func (c *_Client) Write(ctx context.Context, args *WriteArgs) (err error) {
	cmd := &frame.Command{
		Cmd: "GPIO.Write",
	}
	//ctx, tr, finish := c.i.TraceCall(pctx, c.addr, cmd)
	//defer finish(&err)
	//_ = tr

	//tr.LazyPrintf("args: %s", ourjson.LazyJSON(&args))
	cmd.Args = ourjson.DelayMarshaling(args)
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return errors.Trace(err)
	}
	if resp.Status != 0 {
		return errors.Trace(&mgrpc.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}
	return nil
}

//func RegisterService(i *clubby.Instance, impl Service) error {
//s := &_Server{impl}
//i.RegisterCommandHandler("GPIO.Read", s.Read)
//i.RegisterCommandHandler("GPIO.RemoveIntHandler", s.RemoveIntHandler)
//i.RegisterCommandHandler("GPIO.SetIntHandler", s.SetIntHandler)
//i.RegisterCommandHandler("GPIO.Toggle", s.Toggle)
//i.RegisterCommandHandler("GPIO.Write", s.Write)
//i.RegisterService(ServiceID, _ServiceDefinition)
//return nil
//}

type _Server struct {
	impl Service
}

func (s *_Server) Read(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	var args ReadArgs
	if len(cmd.Args) > 0 {
		if err := cmd.Args.UnmarshalInto(&args); err != nil {
			return nil, errors.Annotatef(err, "unmarshaling args")
		}
	}
	return s.impl.Read(ctx, &args)
}

func (s *_Server) RemoveIntHandler(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	var args RemoveIntHandlerArgs
	if len(cmd.Args) > 0 {
		if err := cmd.Args.UnmarshalInto(&args); err != nil {
			return nil, errors.Annotatef(err, "unmarshaling args")
		}
	}
	return nil, s.impl.RemoveIntHandler(ctx, &args)
}

func (s *_Server) SetIntHandler(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	var args SetIntHandlerArgs
	if len(cmd.Args) > 0 {
		if err := cmd.Args.UnmarshalInto(&args); err != nil {
			return nil, errors.Annotatef(err, "unmarshaling args")
		}
	}
	return s.impl.SetIntHandler(ctx, &args)
}

func (s *_Server) Toggle(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	var args ToggleArgs
	if len(cmd.Args) > 0 {
		if err := cmd.Args.UnmarshalInto(&args); err != nil {
			return nil, errors.Annotatef(err, "unmarshaling args")
		}
	}
	return s.impl.Toggle(ctx, &args)
}

func (s *_Server) Write(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	var args WriteArgs
	if len(cmd.Args) > 0 {
		if err := cmd.Args.UnmarshalInto(&args); err != nil {
			return nil, errors.Annotatef(err, "unmarshaling args")
		}
	}
	return nil, s.impl.Write(ctx, &args)
}

var _ServiceDefinition = json.RawMessage([]byte(`{
  "methods": {
    "Read": {
      "args": {
        "pin": {
          "doc": "Pin number.",
          "type": "integer"
        }
      },
      "doc": "Read value of a pin. Switches the pin to input mode if needed.",
      "result": {
        "properties": {
          "value": {
            "doc": "Value of the pin, 0 or 1.",
            "type": "integer"
          }
        },
        "type": "object"
      }
    },
    "RemoveIntHandler": {
      "args": {
        "pin": {
          "doc": "Pin number.",
          "type": "integer"
        }
      },
      "doc": "Removes interrupt handler previosuly set."
    },
    "SetIntHandler": {
      "args": {
        "debounce_ms": {
          "doc": "Optional debouncing delay.",
          "type": "integer"
        },
        "dst": {
          "doc": "Destination address for the RPC. Defaults to source of the request.",
          "type": "string"
        },
        "edge": {
          "doc": "Interrupt trigger edge. \"pos\" (0 -\u003e 1), \"neg\" (1 -\u003e 0) or \"any\".",
          "type": "string"
        },
        "method": {
          "doc": "Method for the request. Defaults to GPIO.Int.",
          "type": "string"
        },
        "pin": {
          "doc": "Pin number.",
          "type": "integer"
        },
        "pull": {
          "doc": "Pull setting. \"up\", \"down\" or \"none\" (default).",
          "type": "string"
        }
      },
      "doc": "Configures an interrupt handler on a pin. Switches the pin to input mode if needed. An RPC with the specified method is sent to the specified address and method when an interrupt happens. Request comes with the following arguments: \"pin\" - pin, \"value\" - read before sending, \"ts\" - timestamp. Response to these requests is not expected.\n",
      "result": {
        "properties": {
          "value": {
            "doc": "Value of the pin after toggle, 0 or 1.",
            "type": "integer"
          }
        },
        "type": "object"
      }
    },
    "Toggle": {
      "args": {
        "pin": {
          "doc": "Pin number.",
          "type": "integer"
        }
      },
      "doc": "Toggles pin value. Switches the pin to output mode if needed.",
      "result": {
        "properties": {
          "value": {
            "doc": "Value of the pin after toggle, 0 or 1.",
            "type": "integer"
          }
        },
        "type": "object"
      }
    },
    "Write": {
      "args": {
        "pin": {
          "doc": "Pin number.",
          "type": "integer"
        },
        "value": {
          "doc": "Value to set, 0 or 1.",
          "type": "integer"
        }
      },
      "doc": "Set value of a pin. Switches the pin to output mode if needed."
    }
  },
  "name": "GPIO",
  "namespace": "http://mongoose-iot.com/fw"
}`))
