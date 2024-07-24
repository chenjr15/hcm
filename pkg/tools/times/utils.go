/*
 * TencentBlueKing is pleased to support the open source community by making
 * 蓝鲸智云 - 混合云管理平台 (BlueKing - Hybrid Cloud Management System) available.
 * Copyright (C) 2022 THL A29 Limited,
 * a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 *
 * We undertake not to change the open source license (MIT license) applicable
 *
 * to the current version of the project delivered to anyone in the future.
 */

package times

import (
	"fmt"
	"time"
)

// GetLastMonth get last month year and month
func GetLastMonth(billYear, billMonth int) (int, int, error) {
	t, err := time.Parse(
		"2006-01-02T15:04:05.000+08:00", fmt.Sprintf("%d-%02d-02T15:04:05.000+08:00", billYear, billMonth))
	if err != nil {
		return 0, 0, err
	}
	lastT := t.AddDate(0, -1, 0)
	return lastT.Year(), int(lastT.Month()), nil
}

// IsLastDayOfMonth 判断给定的天是否是该月的最后一天
func IsLastDayOfMonth(month, day int) (bool, error) {
	if month < 1 || month > 12 {
		return false, fmt.Errorf("invalid month: %d", month)
	}

	// 获取当前年份
	year := time.Now().Year()

	// 创建当月的下个月的第一天
	var nextMonth time.Month
	if month == 12 {
		nextMonth = 1
		year++
	} else {
		nextMonth = time.Month(month + 1)
	}
	firstDayOfNextMonth := time.Date(year, nextMonth, 1, 0, 0, 0, 0, time.UTC)

	// 获取当月的最后一天
	lastDayOfCurrentMonth := firstDayOfNextMonth.AddDate(0, 0, -1).Day()

	// 比较提供的天和当月的最后一天
	if day == lastDayOfCurrentMonth {
		return true, nil
	}
	return false, nil
}
