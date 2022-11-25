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

package errf

import (
	"encoding/json"
	"testing"
)

func TestWrap(t *testing.T) {
	ef := &ErrorF{
		Code:    Unknown,
		Message: "unknown error",
	}

	if Error(ef) == nil {
		t.Errorf("an error should occur, can not be nil")
		return
	}

	compare := new(ErrorF)
	if err := json.Unmarshal([]byte(ef.Error()), compare); err != nil {
		t.Errorf("unexpected formatted error format: %s", ef.Error())
		return
	}

	if compare.Code != Unknown || compare.Message != "unknown error" {
		t.Errorf("invalid error format, parse error info failed")
		return
	}
}

func TestAssignResp(t *testing.T) {
	ef := &ErrorF{
		Code:    Unknown,
		Message: "unknown error",
	}

	resp := ef.Resp()

	if !(resp.Code == ef.Code && resp.Message == ef.Message) {
		t.Error("errorF assign code message failed")
	}
}
