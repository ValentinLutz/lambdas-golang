package core_test

import (
	"regexp"
	"root/services/order/lambda-v1-post-orders/core"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func Benchmark_NewOrderId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		core.NewOrderId(core.RegionEu)
	}
}

func Test_NewOrderId(t *testing.T) {
	// given
	regions := []core.Region{
		core.RegionEu,
		core.RegionEu,
	}
	regex := regexp.MustCompile("^[A-Z0-9]{13}-[A-Z]{2}-[A-Z0-9]{13}$")

	for _, region := range regions {
		for i := 0; i < 100; i++ {
			t.Run(string(region), testNewOrderId(region, regex))
		}
	}
}

func testNewOrderId(region core.Region, regex *regexp.Regexp) func(t *testing.T) {
	return func(t *testing.T) {
		t.Logf("Region: %v", region)

		// when
		orderId := core.NewOrderId(region)

		// then
		td.CmpRe(t, orderId, regex, nil)
	}
}
