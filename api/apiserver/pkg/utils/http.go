/*
Copyright 2021 The PlanetPulse Authors.

Planet Pulse is an API designed to serve climate data pulled from NOAA's
Global Monitoring Laboratory FTP server. This API is based on the
OpenAPI v3 specification.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

A copy of the GNU General Public License can be found here:
https://www.gnu.org/licenses/

API version: 0.1.0
Contact: planetpulse.api@gmail.com
*/

package utils

import (
	"net/http"
	"net/url"
	"strings"
)

// ParseQuery expands the parameters passed to the endpoint
// to account for array-like paramters.
// This allows one to search for, say, the following:
//
//		example.com/v1/stuff?day=1,2,3&day=4
//
// The expansion will allow the day slice to become:
//
//		day := ["1", "2", "3". "4"]
func ParseQuery(r *http.Request) url.Values {
	params := r.URL.Query()
	for key, val := range params {
		var expanded []string
		for _, elem := range val {
			array := strings.Split(elem, ",")
			expanded = append(expanded, array...)
		}
		params[key] = expanded
	}
	return params
}

func SetJsonHeader(w http.ResponseWriter) {

}