/*
 *  Copyright (c) 2021 NetEase Inc.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

/*
 * Project: CurveAdm
 * Created Date: 2021-10-15
 * Author: Jingli Chen (Wine93)
 */

// __SIGN_BY_WINE93__

package context

import (
	"github.com/dingodb/curveadm/pkg/module"
)

type Context struct {
	sshClient *module.SSHClient
	module    *module.Module
	register  *Register
}

func NewContext(sshClient *module.SSHClient) (*Context, error) {
	return &Context{
		sshClient: sshClient,
		module:    module.NewModule(sshClient),
		register:  NewRegister(),
	}, nil
}

func (ctx *Context) Close() {
	if ctx.sshClient != nil {
		ctx.sshClient.Client().Close()
	}
}

func (ctx *Context) SSHClient() *module.SSHClient {
	return ctx.sshClient
}

func (ctx *Context) Module() *module.Module {
	return ctx.module
}

func (ctx *Context) Register() *Register {
	return ctx.register
}
