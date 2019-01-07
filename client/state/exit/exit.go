/*
Copyright (C) 2018-2019 Andreas T Jonsson

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package exit

import "gitlab.com/phix/den/state"

const Name = "exit"

type Exit struct {
}

func New() *Exit {
	return &Exit{}
}

func (s *Exit) Name() string {
	return Name
}

func (s *Exit) Enter(m state.Switcher, from string, data ...interface{}) {}

func (s *Exit) Leave(to string) {}

func (s *Exit) Update() error {
	return nil
}
