package rmq

import (
	"testing"

	. "github.com/iostrovok/check"
	. "github.com/golang/mock/mockgen"
)

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func TestTestSuite(t *testing.T) { TestingT(t) }
