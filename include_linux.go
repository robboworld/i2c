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

/*
#include <linux/i2c-dev.h>
*/
import "C"

const (
	I2C_RETRIES     = C.I2C_RETRIES
	I2C_TIMEOUT     = C.I2C_TIMEOUT
	I2C_SLAVE       = C.I2C_SLAVE
	I2C_SLAVE_FORCE = C.I2C_SLAVE_FORCE
	I2C_TENBIT      = C.I2C_TENBIT
	I2C_FUNCS       = C.I2C_FUNCS
	I2C_RDWR        = C.I2C_RDWR
	I2C_PEC         = C.I2C_PEC
	I2C_SMBUS       = C.I2C_SMBUS
)
