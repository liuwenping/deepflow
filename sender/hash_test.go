package sender

import (
	"sort"
	"testing"

	"gitlab.x.lan/application/droplet-app/pkg/mapper/consolelog"
	"gitlab.x.lan/application/droplet-app/pkg/mapper/flow"
	"gitlab.x.lan/application/droplet-app/pkg/mapper/flowtype"
	"gitlab.x.lan/application/droplet-app/pkg/mapper/fps"
	"gitlab.x.lan/application/droplet-app/pkg/mapper/geo"
	"gitlab.x.lan/application/droplet-app/pkg/mapper/perf"
	"gitlab.x.lan/application/droplet-app/pkg/mapper/usage"
	"gitlab.x.lan/yunshan/droplet-libs/zerodoc"
)

var FLOW_APP_CODES = [][]zerodoc.Code{
	consolelog.CODES,

	flow.NODE_CODES,
	flow.NODE_PORT_CODES,
	flow.EDGE_CODES,
	flow.TOR_EDGE_CODES,
	flow.GROUP_NODE_CODES,
	flow.GROUP_NODE_PORT_CODES,
	flow.GROUP_EDGE_CODES,
	flow.GROUP_EDGE_PORT_CODES,
	flow.POLICY_NODE_CODES,
	flow.POLICY_ISP_EDGE_CODES,
	flow.POLICY_PORT_CODES,
	flow.POLICY_GROUP_EDGE_CODES,
	flow.WHITELIST_GROUP_EDGE_PORT_CODES,

	flowtype.NODE_CODES,

	geo.CHN_CODES,
	geo.NON_CHN_CODES,
	geo.POLICY_CHN_CODES,
	geo.POLICY_NON_CHN_CODES,

	perf.NODE_CODES,
	perf.GROUP_NODE_CODES,
	perf.GROUP_NODE_PORT_CODES,
	perf.GROUP_EDGE_CODES,
	perf.GROUP_EDGE_PORT_CODES,
	perf.POLICY_NODE_CODES,
}

var METERING_APP_CODES = [][]zerodoc.Code{
	fps.NODE_CODES,
	fps.POLICY_NODE_CODES,

	usage.NODE_CODES,
	usage.POLICY_NODE_CODES,
}

var FLOW_CODES []zerodoc.Code
var METERING_CODES []zerodoc.Code

func init() {
	set := make(map[zerodoc.Code]bool)
	for _, app := range FLOW_APP_CODES {
		for _, v := range app {
			set[v] = true
		}
	}
	for v, _ := range set {
		FLOW_CODES = append(FLOW_CODES, v)
	}

	set = make(map[zerodoc.Code]bool)
	for _, app := range METERING_APP_CODES {
		for _, v := range app {
			set[v] = true
		}
	}
	for v, _ := range set {
		METERING_CODES = append(METERING_CODES, v)
	}
}

type pair struct {
	k uint32
	v int
}

func showDist(t *testing.T, list []uint32) {
	counter := make(map[uint32]int)
	for _, v := range list {
		counter[v]++
	}
	pairs := make([]*pair, 0, len(counter))
	for k, v := range counter {
		pairs = append(pairs, &pair{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].v > pairs[j].v
	})
	for _, p := range pairs {
		if p.v > 1 {
			t.Error("存在hash相同的code")
		}
		t.Logf("\t%d: %d\n", p.k, p.v)
	}
}

func TestHashedCodes(t *testing.T) {
	flowHashed := make([]uint32, len(FLOW_CODES))
	for i, v := range FLOW_CODES {
		flowHashed[i] = codeHash(v)
	}
	t.Log("flow codes:")
	showDist(t, flowHashed)
	meteringHashed := make([]uint32, len(METERING_CODES))
	for i, v := range METERING_CODES {
		meteringHashed[i] = codeHash(v)
	}
	t.Log("metering codes:")
	showDist(t, meteringHashed)
}
