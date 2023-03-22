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

package huawei

import (
	"time"

	"hcm/pkg/adaptor/huawei"
	routetable "hcm/pkg/api/hc-service/route-table"
	dataservice "hcm/pkg/client/data-service"
	hcservice "hcm/pkg/client/hc-service"
	"hcm/pkg/criteria/enumor"
	"hcm/pkg/kit"
	"hcm/pkg/logs"
)

// SyncRouteTable 同步路由表
func SyncRouteTable(kt *kit.Kit, service *hcservice.Client, dataCli *dataservice.Client, accountID string) error {
	start := time.Now()
	logs.V(3).Infof("cloud-server-sync-%s account[%s] sync route table start, time: %v, rid: %s",
		enumor.HuaWei, accountID, start, kt.Rid)

	defer func() {
		logs.V(3).Infof("cloud-server-sync-%s account[%s] sync route table end, cost: %v, rid: %s",
			enumor.HuaWei, accountID, time.Since(start), kt.Rid)
	}()

	regions, err := ListRegionByService(kt, dataCli, huawei.Vpc)
	if err != nil {
		logs.Errorf("sync huawei list region failed, err: %v, rid: %s", err, kt.Rid)
		return err
	}

	for _, region := range regions {
		req := &routetable.HuaWeiRouteTableSyncReq{
			AccountID: accountID,
			Region:    region,
		}
		if err := service.HuaWei.RouteTable.SyncRouteTable(kt.Ctx, kt.Header(), req); Error(err) != nil {
			logs.Errorf("cloud-server-sync-%s account[%s] sync route table failed, req: %v, err: %v, rid: %s",
				enumor.HuaWei, accountID, req, err, kt.Rid)
			return err
		}
	}

	return nil
}
