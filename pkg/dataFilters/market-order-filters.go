package dataFilters

import (
	"github.com/GabrielDCelery/eve-online-tools-cli/pkg/dataTransforms"
	"github.com/GabrielDCelery/eve-online-tools-cli/pkg/utils"
)

func doesMatchTypeID(m *dataTransforms.MarketOrder, typeID uint64) bool {
	return m.TypeID == typeID
}

func isOfOrderType(m *dataTransforms.MarketOrder, orderTypes *[]string) bool {
	returnSellOrders := utils.Contains(orderTypes, "sell")
	returnBuyOrders := utils.Contains(orderTypes, "buy")
	if returnBuyOrders && returnSellOrders {
		return true
	}
	if returnSellOrders {
		return !m.IsBuyOrder
	}
	if returnBuyOrders {
		return m.IsBuyOrder
	}
	return false
}

func isOfVolumeStatus(m *dataTransforms.MarketOrder, volumeStatuses *[]string) bool {
	returnFullOrders := utils.Contains(volumeStatuses, "full")
	returnDecreasingOrders := utils.Contains(volumeStatuses, "decreasing")
	if returnFullOrders && returnDecreasingOrders {
		return m.VolumeRemain <= m.VolumeTotal
	}
	if returnFullOrders {
		return m.VolumeRemain == m.VolumeTotal
	}
	if returnDecreasingOrders {
		return m.VolumeRemain < m.VolumeTotal
	}
	return false
}

func DoesMarketOrderMatchFilterConditions(marketOrder *dataTransforms.MarketOrder, typeID uint64, orderTypes *[]string, volumeStatuses *[]string) bool {
	return isOfVolumeStatus(marketOrder, volumeStatuses) && doesMatchTypeID(marketOrder, typeID) && isOfOrderType(marketOrder, orderTypes)
}
