package corev1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	proto "google.golang.org/protobuf/proto"
)

func TestVTProtoUnmarshalCompat(t *testing.T) {
	// Construct our usage of empty message as still being a valid oneof.
	m := &AllowedRelation{
		RelationOrWildcard: &AllowedRelation_PublicWildcard_{},
	}

	// Marshal it with protobuf.
	protobuf, err := proto.Marshal(m)
	if err != nil {
		t.Fatal(err)
	}

	// Unmarshal it with protobuf.
	pbm := &AllowedRelation{}
	if err := proto.Unmarshal(protobuf, m); err != nil {
		t.Fatal(err)
	}

	// Unmarshal it with vtprotobuf.
	vtm := &AllowedRelation{}
	if err := vtm.UnmarshalVT(protobuf); err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n\n", pbm)
	fmt.Printf("%#v\n", vtm)

	require.True(t, proto.Equal(pbm, vtm))
}
