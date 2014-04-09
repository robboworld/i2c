/*
   i2c go package, a go interface to I2C devices.
   Copyright (C) 2014  Dmitry Mikhirev

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

// Package ioctl provides ioctl system call missing in standard library.
package i2c

import (
	"syscall"
)

func Ioctl(fd uintptr, cmd int, ptr uintptr) error {
	_, _, en := syscall.Syscall(
		syscall.SYS_IOCTL,
		fd,
		uintptr(cmd),
		uintptr(ptr),
	)
	if en != 0 {
		return en
	}
	return nil
}
