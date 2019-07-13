package main

import (
	"strings"
)

const (
	FIELD_IDFA       = "idfa"
	FIELD_GPS_ADID   = "gps_adid"
	FIELD_ANDRIOD_ID = "android_id"
	FIELD_FIRE_ADID  = "fire_adid"
)

type template struct{}

// 获取目标字段的索引位置
func (t *template) GetTargetIndex(record []string) []int {
	res := []int{}

	for k, v := range record {
		if v == FIELD_IDFA || v == FIELD_GPS_ADID || v == FIELD_ANDRIOD_ID || v == FIELD_FIRE_ADID {
			res = append(res, k)
		} else if strings.Contains(v, "||") {
			if strings.Contains(v, FIELD_IDFA) || strings.Contains(v, FIELD_GPS_ADID) ||
				strings.Contains(v, FIELD_ANDRIOD_ID) || strings.Contains(v, FIELD_FIRE_ADID) {
				res = append(res, k)
			}
		}
	}

	return res
}
