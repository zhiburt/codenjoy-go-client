package engine

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
	assert.Equal(t, 0, LEFT.Value())
	assert.Equal(t, 1, RIGHT.Value())
	assert.Equal(t, 2, UP.Value())
	assert.Equal(t, 3, DOWN.Value())
	assert.Equal(t, 4, ACT.Value())
	assert.Equal(t, 5, STOP.Value())
}

func TestDirectionChangeX(t *testing.T) {
	assert.Equal(t, 0, LEFT.ChangeX(1))
	assert.Equal(t, 2, RIGHT.ChangeX(1))
	assert.Equal(t, 1, UP.ChangeX(1))
	assert.Equal(t, 1, DOWN.ChangeX(1))
}

func TestDirectionChangeY(t *testing.T) {
	assert.Equal(t, 1, LEFT.ChangeY(1))
	assert.Equal(t, 1, RIGHT.ChangeY(1))
	assert.Equal(t, 2, UP.ChangeY(1))
	assert.Equal(t, 0, DOWN.ChangeY(1))
}

func TestDirectionInverted(t *testing.T) {
	assert.Equal(t, RIGHT, LEFT.Inverted())
	assert.Equal(t, LEFT, RIGHT.Inverted())
	assert.Equal(t, DOWN, UP.Inverted())
	assert.Equal(t, UP, DOWN.Inverted())
}

func TestDirectionInvalidInverted(t *testing.T) {
	assert.Panics(t, func() { ACT.Inverted() })
	assert.Panics(t, func() { STOP.Inverted() })
}
