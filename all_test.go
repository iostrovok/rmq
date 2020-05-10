package rmq

import (
	"testing"

	. "github.com/iostrovok/check"
)

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func TestTestSuite(t *testing.T) { TestingT(t) }
