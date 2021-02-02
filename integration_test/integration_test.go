package integration

import (
	msg_sdk "github.com/kaifei-bianjie/msg-parser"
	"github.com/stretchr/testify/suite"
	"testing"
)

type IntegrationTestSuite struct {
	msg_sdk.MsgClient
	suite.Suite
}

type SubTest struct {
	testName string
	testCase func(s IntegrationTestSuite)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.MsgClient = msg_sdk.NewMsgClient()
}
