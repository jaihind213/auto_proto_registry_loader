package pocket_protobuff_registry

import (
	_ "github.com/jaihind213/auto_proto_registry_loader/example_go_bindings/pocket_registry"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"testing"
)

func TestLookup(t *testing.T) {
	_, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName("foo.bar.Keeper"))
	if err != nil {
		t.Error("error in looking up type Foo. it should exist in registry but does not ! ," + err.Error())
	}
}
