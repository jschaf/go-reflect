package main

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"reflect"
)

func main() {
	msg, err := newProtoByName("main.Foo")
	if err != nil {
		log.Fatal(err)
	}
	foo := msg.(*Foo)
	foo.Name = "foo_name"
	fmt.Printf("Proto: %+v", foo)
}

// Reflectively creates a proto from the fully qualified name.
func newProtoByName(protoType string) (proto.Message, error) {
	// reflect.Type of a pointer to the Go struct representing protoType.
	msgPtr := proto.MessageType(protoType)
	if msgPtr == nil {
		return nil, errors.New("failed to reflectively find a proto message named " + protoType)
	}
	// Dereference the pointer and get a struct reflect.Type.
	msgType := msgPtr.Elem()
	newMsg := reflect.New(msgType)
	// Why is this proto.Message and not *proto.Message?
	// I thought reflect.New would create a pointer.
	msg, ok := newMsg.Interface().(proto.Message)
	if !ok {
		return nil, errors.New(
			fmt.Sprintf("the type found for proto '%s' does not implement proto.Message", protoType))
	}
	return msg, nil
}
