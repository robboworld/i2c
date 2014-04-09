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

// Package i2c provides interface to I2C slave devices. Only linux systems are
// supported currently.
package i2c

import (
	"errors"
	"fmt"
	"io"
)

// Slave represent an I2C slave device.
type Slave struct {
	Bus  uint
	Addr uint
	io   io.ReadWriteCloser
}

// NewSlave creates a new Slave object representing device on bus with address
// addr.
func NewSlave(bus uint, addr uint) *Slave {
	d := Slave{bus, addr, nil}
	return &d
}

func (dev *Slave) String() string {
	return fmt.Sprintf("I2C-%d:0x%x", dev.Bus, dev.Addr)
}

// Close closes the corresponding device. It can not be used for reading
// nor writing until reopened.
// It returns a error if any.
func (dev *Slave) Close() error {
	if dev.io == nil {
		return errors.New("cannot close " + dev.String() + " that is not opened")
	}
	if err := dev.io.Close(); err != nil {
		return errors.New("cannot close " + dev.String() + ": " + err.Error())
	}
	dev.io = nil
	return nil
}

// Write writes len(b) bytes to the device.
// It returns number of bytes written and a error if any.
func (dev *Slave) Write(b []byte) (n int, err error) {
	if dev.io == nil {
		return 0, errors.New(dev.String() + " is not opened")
	}
	return dev.io.Write(b)
}

// WriteByte writes single byte to the device.
// It returns a error if any.
func (dev *Slave) WriteByte(b byte) (err error) {
	_, err = dev.Write([]byte{b})
	return err
}

// Read reads up to len(b) bytes from the device.
// It returns number of bytes read and a error if any.
func (dev *Slave) Read(b []byte) (n int, err error) {
	if dev.io == nil {
		return 0, errors.New(dev.String() + " is not opened")
	}
	return dev.io.Read(b)
}

// ReadByte reads single byte from the device.
// It returns a byte read and a error if any.
func (dev *Slave) ReadByte() (byte, error) {
	b := make([]byte, 1, 1)
	_, err := dev.Read(b)
	return b[0], err
}
