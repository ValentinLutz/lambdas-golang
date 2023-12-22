package core_test

import (
	"regexp"
	"root/services/order/lambda-v1-post-orders/core"
	"testing"

	"github.com/maxatome/go-testdeep/td"
)

func Benchmark_NewOrderId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		core.NewOrderId(core.RegionNone)
	}
}

func Test_NewOrderId(t *testing.T) {
	// given
	regions := []string{
		"REGION_NONE",
		"RegionEu",
		"RegionUs",
	}
	regex := regexp.MustCompile("^[A-Za-z0-9]{13}-[A-Z]{2,4}-[A-Za-z0-9]{13}$")

	for _, region := range regions {
		for i := 0; i < 100; i++ {
			t.Run(region, testNewOrderId(core.Region(region), regex))
		}
	}
}

func testNewOrderId(region core.Region, regex *regexp.Regexp) func(t *testing.T) {
	return func(t *testing.T) {
		t.Logf("Region: %v", region)

		// when
		orderId := core.NewOrderId(region)
		// then
		td.Re(t, orderId, regex)
	}
}
