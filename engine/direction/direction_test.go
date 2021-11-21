package direction

/*-
 * #%L
 * Codenjoy - it's a dojo-like platform from developers to developers.
 * %%
 * Copyright (C) 2021 Codenjoy
 * %%
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public
 * License along with this program.  If not, see
 * <http://www.gnu.org/licenses/gpl-3.0.html>.
 * #L%
 */

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirectionValue(t *testing.T) {
	assert.Equal(t, 0, New(0, 0, 0, Left).Value())
	assert.Equal(t, 1, New(1, 0, 0, Right).Value())
	assert.Equal(t, 2, New(2, 0, 0, Up).Value())
	assert.Equal(t, 3, New(3, 0, 0, Down).Value())
	assert.Equal(t, 4, New(4, 0, 0, Act).Value())
	assert.Equal(t, 5, New(5, 0, 0, Stop).Value())
}

func TestDirectionChangeX(t *testing.T) {
	assert.Equal(t, 0, New(0, -1, 0, Left).ChangeX(1))
	assert.Equal(t, 2, New(0, 1, 0, Right).ChangeX(1))
	assert.Equal(t, 1, New(0, 0, 0, Up).ChangeX(1))
	assert.Equal(t, 1, New(0, 0, 0, Down).ChangeX(1))
}

func TestDirectionChangeY(t *testing.T) {
	assert.Equal(t, 1, New(0, 0, 0, Left).ChangeY(1))
	assert.Equal(t, 1, New(0, 0, 0, Right).ChangeY(1))
	assert.Equal(t, 2, New(0, 0, 1, Up).ChangeY(1))
	assert.Equal(t, 0, New(0, 0, -1, Down).ChangeY(1))
}

func TestDirectionInverted(t *testing.T) {
	assert.Equal(t, Right, Direction{name: Left}.Inverted())
	assert.Equal(t, Left, Direction{name: Right}.Inverted())
	assert.Equal(t, Down, Direction{name: Up}.Inverted())
	assert.Equal(t, Up, Direction{name: Down}.Inverted())
}

func TestDirectionInvalidInverted(t *testing.T) {
	assert.Panics(t, func() { Direction{name: Act}.Inverted() })
	assert.Panics(t, func() { Direction{name: Stop}.Inverted() })
}
