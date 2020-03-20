/*
 * Copyright 2018-present Open Networking Foundation

 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at

 * http://www.apache.org/licenses/LICENSE-2.0

 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package model

import (
	"context"
	"runtime/debug"
	"sync"

	"github.com/opencord/voltha-protos/v3/go/voltha"
)

var callbackMutex sync.Mutex

func commonChanCallback(ctx context.Context, args ...interface{}) interface{} {
	logger.Infof("Running common callback - arg count: %d", len(args))

	//for i := 0; i < len(args); i++ {
	//	logger.Infof("ARG %d : %+v", i, args[i])
	//}

	callbackMutex.Lock()
	defer callbackMutex.Unlock()

	execDoneChan := args[1].(*chan struct{})

	// Inform the caller that the callback was executed
	if *execDoneChan != nil {
		logger.Infof("Sending completion indication - stack:%s", string(debug.Stack()))
		close(*execDoneChan)
		*execDoneChan = nil
	}

	return nil
}

func commonCallback2(ctx context.Context, args ...interface{}) interface{} {
	logger.Infof("Running common2 callback - arg count: %d %+v", len(args), args)

	return nil
}

func commonCallbackFunc(ctx context.Context, args ...interface{}) interface{} {
	logger.Infof("Running common callback - arg count: %d", len(args))

	for i := 0; i < len(args); i++ {
		logger.Infof("ARG %d : %+v", i, args[i])
	}
	execStatusFunc := args[1].(func(bool))

	// Inform the caller that the callback was executed
	execStatusFunc(true)

	return nil
}

func firstCallback(ctx context.Context, args ...interface{}) interface{} {
	name := args[0]
	id := args[1]
	logger.Infof("Running first callback - name: %s, id: %s\n", name, id)
	return nil
}

func secondCallback(ctx context.Context, args ...interface{}) interface{} {
	name := args[0].(map[string]string)
	id := args[1]
	logger.Infof("Running second callback - name: %s, id: %f\n", name["name"], id)
	// FIXME: the panic call seem to interfere with the logging mechanism
	//panic("Generating a panic in second callback")
	return nil
}

func thirdCallback(ctx context.Context, args ...interface{}) interface{} {
	name := args[0]
	id := args[1].(*voltha.Device)
	logger.Infof("Running third callback - name: %+v, id: %s\n", name, id.Id)
	return nil
}
