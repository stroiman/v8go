package v8go

// #include "inspector.h"
import "C"
import (
	"fmt"
	"sync"
)

type Inspector struct {
	ptr C.InspectorPtr
}

type InspectorClient struct {
	ptr     C.InspectorClientPtr
	handler ConsoleAPIMessageHandler
}

type MessageErrorLevel int

// registry is a simple map of int->something. Allows passing an int to C-code
// that can be used in a callback; then Go code can retrieve the right object
// afterwards.
// This can be generalised, e.g. Isolate has similar functionality handling
// callback functions
type registry[T any] struct {
	entries map[C.int]T
	id      C.int
	lock    sync.RWMutex
}

func newRegistry[T any]() *registry[T] {
	return &registry[T]{
		entries: make(map[C.int]T),
	}
}

var clientRegistry = newRegistry[*InspectorClient]()

func (r *registry[T]) register(instance T) C.int {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.id++
	r.entries[r.id] = instance
	return r.id
}

func (r *registry[T]) unregister(id C.int) {
	r.lock.Lock()
	defer r.lock.Unlock()
	delete(r.entries, id)
}

func (r *registry[T]) get(id C.int) T {
	r.lock.RLock()
	defer r.lock.RUnlock()
	// What if not found? Return two parameters? Panic? (my preference)
	return r.entries[id]
}

type ConsoleAPIMessage struct {
	contextGroupId int
	errorLevel     MessageErrorLevel
	Message        string
	url            string
	lineNumber     int
	columnNumber   int
	// stackTrace StackTrace
}

type ConsoleAPIMessageHandler interface {
	ConsoleAPIMessage(message ConsoleAPIMessage)
}

type ConsoleAPIMessageHandlerFunc func(msg ConsoleAPIMessage)

func (f ConsoleAPIMessageHandlerFunc) ConsoleAPIMessage(msg ConsoleAPIMessage) {
	f(msg)
}

func NewInspector(iso *Isolate, client *InspectorClient) *Inspector {
	ptr := C.CreateInspector(iso.ptr, client.ptr)
	return &Inspector{
		ptr: ptr,
	}
}

func (i *Inspector) Dispose() {
	C.DeleteInspector(i.ptr)
}

func (i *Inspector) ContextCreated(ctx *Context) {
	fmt.Println("Created from Go")
	C.InspectorContextCreated(i.ptr, ctx.ptr)
}

func (i *Inspector) ContextDestroyed(ctx *Context) {
	C.InspectorContextDestroyed(i.ptr, ctx.ptr)
}

func NewInspectorClient(handler ConsoleAPIMessageHandler) *InspectorClient {
	result := &InspectorClient{
		handler: handler,
	}
	ref := clientRegistry.register(result)
	result.ptr = C.NewInspectorClient(ref)
	return result
}

func (c *InspectorClient) Dispose() {
	C.DeleteInspectorClient(c.ptr)
}

//export goHandleConsoleAPIMessageCallback
func goHandleConsoleAPIMessageCallback(callbackRef C.int) {
	fmt.Println("Callback from Go")
	// Convert data to Go data
	client := clientRegistry.get(callbackRef)
	client.handler.ConsoleAPIMessage(ConsoleAPIMessage{
		Message: "Hello",
	})
	// client.handleConsoleAPIMessageCallback(data)
}
