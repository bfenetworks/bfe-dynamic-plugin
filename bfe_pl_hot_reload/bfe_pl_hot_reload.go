// Copyright (c) 2019 Baidu, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
)

import (
	"github.com/baidu/go-lib/web-monitor/web_monitor"
)

import (
	"github.com/bfenetworks/bfe/bfe_module"
)

var Name = "pl_reload"
var Version = "0.9.0-dev"
var Description = "plugin for bfe to hot reload config"

func Init(cbs *bfe_module.BfeCallbacks,
	whs *web_monitor.WebHandlers,
	cr string) error {

	// register handler
	err := cbs.AddExtendAsync(bfe_module.HandleExtendAsync, hotReloadHandler)
	if err != nil {
		return fmt.Errorf("%s.Init(): AddExtensionAsync(hotReloadHandler): %s", Name, err.Error())
	}

	return nil
}

func hotReloadHandler(em *bfe_module.ExtendModule) {
	// TODO step1: watch key from etcdv3/zk/other

	// TODO step2: hot reload base on watch event, eg: https://www.bfe-networks.net/zh_cn/operation/reload/
	// NameConfReload
	// TLSConfReload
	// GslbDataConfReload
	// ServerDataConfReload
	// SessionTicketKeyReload

	/*
		for example:

		err := em.NameConfReload(url.Values{})
		if err != nil {
			// do err
		}

	*/
}
