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

package i2c

import (
	"errors"
	"fmt"
	"os"
)

func (dev *Slave) fileName() string {
	return fmt.Sprintf("/dev/i2c-%d", dev.Bus)
}

// Open opens the slave device with specified flag (e. g. os.O_RDWR).
// It returns a error if any.
func (dev *Slave) Open(flag int) (err error) {
	if dev.io != nil {
		return errors.New(dev.String() + " is already opened")
	}
	dev.io, err = os.OpenFile(dev.fileName(), flag, os.ModeDevice)
	if err != nil {
		return errors.New("cannot open " + dev.fileName() + ": " + err.Error())
	}
	err = Ioctl(dev.io.(*os.File).Fd(), I2C_SLAVE, uintptr(dev.Addr))
	if err != nil {
		dev.io.Close()
		return err
	}
	return nil
}
