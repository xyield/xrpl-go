package key

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/test"
)

func TestValidationCreateRequest(t *testing.T) {
	s := ValidationCreateRequest{
		Secret: "abc",
	}

	j := `{
	"secret": "abc"
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestValidationCreateResponse(t *testing.T) {
	s := ValidationCreateResponse{
		ValidationKey:       "FAWN JAVA JADE HEAL VARY HER REEL SHAW GAIL ARCH BEN IRMA",
		ValidationPublicKey: "n9Mxf6qD4J55XeLSCEpqaePW4GjoCR5U1ZeGZGJUCNe3bQa4yQbG",
		ValidationSeed:      "ssZkdwURFMBXenJPbrpE14b6noJSu",
	}

	j := `{
	"validation_key": "FAWN JAVA JADE HEAL VARY HER REEL SHAW GAIL ARCH BEN IRMA",
	"validation_public_key": "n9Mxf6qD4J55XeLSCEpqaePW4GjoCR5U1ZeGZGJUCNe3bQa4yQbG",
	"validation_seed": "ssZkdwURFMBXenJPbrpE14b6noJSu"
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}

}
