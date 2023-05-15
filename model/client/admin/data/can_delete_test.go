package data

import (
	"testing"

	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/test"
)

func TestCanDeleteRequest(t *testing.T) {
	s := CanDeleteRequest{
		CanDelete: common.CURRENT,
	}

	j := `{
	"can_delete": "current"
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestCanDeleteResponse(t *testing.T) {
	s := CanDeleteResponse{
		CanDelete: 54321,
	}

	j := `{
	"can_delete": 54321
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}
