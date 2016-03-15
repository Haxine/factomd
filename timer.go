// Copyright 2015 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/FactomProject/factomd/common/interfaces"
	s "github.com/FactomProject/factomd/state"
	"time"
)

var _ = (*s.State)(nil)

func Timer(state interfaces.IState) {

	//s := state.(*s.State)

	time.Sleep(2 * time.Second)

	billion := int64(1000000000)
	period := int64(state.GetDirectoryBlockInSeconds()) * billion
	tenthPeriod := period / 10

	now := time.Now().UnixNano() // Time in billionths of a second

	wait := tenthPeriod - (now % tenthPeriod)

	next := now + wait + tenthPeriod

	state.Print(fmt.Sprintf("Time: %v\r\n", time.Now()))
	time.Sleep(time.Duration(wait))
	for {
		for i := 0; i < 10; i++ {
			now = time.Now().UnixNano()
			wait := next - now
			next += tenthPeriod
			time.Sleep(time.Duration(wait))

			/**
			if len(s.ShutdownChan) == 0 {
				state.Print(fmt.Sprintf("\r%19s: %s %s",
					"Timer",
					state.String(),
					(string)((([]byte)("-\\|/-\\|/-="))[i])))
			}
			**/
			// End of the last period, and this is a server, send messages that
			// close off the minute.
			found, _ := state.GetFedServerIndex(state.GetLeaderHeight())
			if found {
				eom := state.NewEOM(i)
				state.InMsgQueue() <- eom
			}
		}
	}

}
