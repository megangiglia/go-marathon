/*
Copyright 2014 Rohith All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package marathon

import (
	"errors"
	"net"
	"strings"
)

func getInterfaceAddress(name string) (string, error) {
	if interfaces, err := net.Interfaces(); err != nil {
		return "", err
	} else {
		for _, iface := range interfaces {
			/* step: get only the interface we're interested in */
			if iface.Name == name {
				addrs, err := iface.Addrs()
				if err != nil {
					return "", err
				}
				/* step: return the first address */
				if len(addrs) > 0 {
					return strings.SplitN(addrs[0].String(), "/", 2)[0], nil
				}
			}
		}
	}
	return "", errors.New("Unable to determine or find the interface")
}

func contains(elements []string, value string) bool {
	for _, element := range elements {
		if element == value {
			return true
		}
	}
	return false
}
