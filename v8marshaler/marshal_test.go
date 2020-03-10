package v8marshaler

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type marshalTestSuite struct {
	suite.Suite
}

func TestMarshaler(t *testing.T) {
	suite.Run(t, new(marshalTestSuite))
}

func (t *marshalTestSuite) TestUnmarshalRepository() {

	object := &RepositoryCreateOptions{
		Designer: newDefaultDesigner(),
		Repository: &Repository{
			Path: "/tem/",
			User: "Администратор",
		},
		Extension:                 "temp",
		AllowConfigurationChanges: false,
		//ChangesAllowedRule:        REPOSITORY_SUPPORT_IS_EDITABLE,
		//ChangesNotRecommendedRule: REPOSITORY_SUPPORT_NOT_EDITABLE,
		NoBind: true,
	}

	args, err := v8Marshal(object)

	t.r().NoError(err)

	t.r().Equal(len(args), 1, "len must be equal")
	//t.r().Equal(codes[0].PromocodeID, "START", "Промокод должен быть START")

}
