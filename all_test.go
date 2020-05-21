package rmq

import (
	"testing"

	. "github.com/iostrovok/check"
	_ "github.com/golang/mock/mockgen/model"
)

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func TestTestSuite(t *testing.T) { TestingT(t) }
