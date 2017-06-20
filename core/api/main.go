// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"os"
	"time"

	"google.golang.org/grpc/grpclog"

	"github.com/franela/goreq"
	"github.com/ellcrys/cocoon/core/api/cmd"
	"github.com/ellcrys/cocoon/core/common"
	"github.com/ellcrys/cocoon/core/config"
	logging "github.com/op/go-logging"
)

var log *logging.Logger

func init() {
	config.ConfigureLogger()
	log = logging.MustGetLogger("main")
	goreq.SetConnectTimeout(5 * time.Second)
	if len(os.Getenv("ENABLE_GRPC_LOG")) == 0 {
		gl := common.GLogger{}
		gl.Disable(true, true)
		grpclog.SetLogger(&gl)
	}
}

func main() {
	// defer profile.Start(profile.MemProfile).Stop()

	// go func() {
	// 	go func() {
	// 		log.Info(http.ListenAndServe("localhost:6060", nil).Error())
	// 	}()

	// 	for {
	// 		var mem runtime.MemStats
	// 		runtime.ReadMemStats(&mem)
	// 		log.Infof("Alloc: %d, Total Alloc: %d, HeapAlloc: %d, HeapSys: %d", mem.Alloc, mem.TotalAlloc, mem.HeapAlloc, mem.HeapSys)
	// 		time.Sleep(10 * time.Second)
	// 	}
	// }()

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Error(err.Error())
		os.Exit(-1)
	}
}
