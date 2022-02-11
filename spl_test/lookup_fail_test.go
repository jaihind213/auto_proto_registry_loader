package spl_test

import (
	//_ "github.com/jaihind213/pocket_protobuff_registry/go_bindings/pocket_registry"
	"fmt"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"testing"
)

func TestLookupWithNoImplicitImport(t *testing.T) {
	_, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName("foo.bar.Keeper"))
	if err == nil {
		t.Error("we expect an error while looking up type as the implicit import of pocket registry package has not been done")
		//_ "github.com/jaihind213/pocket_protobuff_registry/go_bindings/pocket_registry"
	} else {
		fmt.Println(err.Error())
	}
}
