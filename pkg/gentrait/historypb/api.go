package historypb

import (
	"github.com/vanti-dev/sc-bos/pkg/gen"
	"github.com/vanti-dev/sc-bos/pkg/node"
)

func AddSupport(n node.Supporter) {
	{
		r := gen.NewElectricHistoryRouter()
		n.Support(node.Routing(r), node.Clients(gen.WrapElectricHistory(r)))
	}
	{
		r := gen.NewMeterHistoryRouter()
		n.Support(node.Routing(r), node.Clients(gen.WrapMeterHistory(r)))
	}
	{
		r := gen.NewOccupancySensorHistoryRouter()
		n.Support(node.Routing(r), node.Clients(gen.WrapOccupancySensorHistory(r)))
	}
}
