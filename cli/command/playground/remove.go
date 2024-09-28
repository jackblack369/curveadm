/*
 *  Copyright (c) 2022 NetEase Inc.
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
 * Created Date: 2022-06-23
 * Author: Jingli Chen (Wine93)
 */

package playground

import (
	"github.com/dingodb/curveadm/cli/cli"
	"github.com/dingodb/curveadm/internal/errno"
	"github.com/dingodb/curveadm/internal/playbook"
	"github.com/dingodb/curveadm/internal/storage"
	cliutil "github.com/dingodb/curveadm/internal/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type removeOptions struct {
	id string
}

var REMOVE_PLAYGROUND_PLAYBOOK_STEPS = []int{
	playbook.REMOVE_PLAYGROUND,
}

func NewRemoveCommand(curveadm *cli.CurveAdm) *cobra.Command {
	var options removeOptions

	cmd := &cobra.Command{
		Use:     "rm ID",
		Aliases: []string{"delete"},
		Short:   "Remove playground",
		Args:    cliutil.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			options.id = args[0]
			return runRemove(curveadm, options)
		},
		DisableFlagsInUseLine: true,
	}

	return cmd
}

func genRemovePlaybook(curveadm *cli.CurveAdm,
	playgrounds []storage.Playground) (*playbook.Playbook, error) {
	configs := []interface{}{}
	for _, playground := range playgrounds {
		configs = append(configs, playground)
	}
	steps := REMOVE_PLAYGROUND_PLAYBOOK_STEPS
	pb := playbook.NewPlaybook(curveadm)
	for _, step := range steps {
		pb.AddStep(&playbook.PlaybookStep{
			Type:    step,
			Configs: configs,
		})
	}
	return pb, nil
}

func runRemove(curveadm *cli.CurveAdm, options removeOptions) error {
	// 1) get playground
	id := options.id
	playgrounds, err := curveadm.Storage().GetPlaygroundById(id)
	if err != nil {
		return errno.ERR_GET_PLAYGROUND_BY_NAME_FAILED.E(err)
	} else if len(playgrounds) == 0 {
		return errno.ERR_PLAYGROUND_NOT_FOUND.
			F("id=%s", id)
	}

	// 2) generate remove playground
	pb, err := genRemovePlaybook(curveadm, playgrounds)
	if err != nil {
		return err
	}

	// 3) run playground
	err = pb.Run()
	if err != nil {
		return err
	}

	// 4) print success prompt
	curveadm.WriteOutln("")
	curveadm.WriteOutln(color.GreenString("Playground '%s' removed.", playgrounds[0].Name))
	return nil
}
