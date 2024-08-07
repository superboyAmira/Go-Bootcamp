package detector

import (
	"log"
	"team00/internal/clientServices/receiver"
)

func AnomalyDetect(value *float64, k *float64, statistic *receiver.ReceivedData) bool {
	border_l := statistic.Mean - (*k * statistic.STD)
	border_h := statistic.Mean + (*k * statistic.STD)

	if !(*value >= border_l && *value <= border_h) {
		log.Printf("{ANOMALY DETECTED}: Value: %f\t lower border: %f higher border: %f\n", *value, border_l, border_h)
		return true
	} else {
		log.Printf("{ACCEPTED}: Value: %f\t lower border: %f higher border: %f\n", *value, border_l, border_h)
	}

	return false
}
