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
    "github.com/codenjoyme/codenjoy-go-client/engine"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestDirectionValue(t *testing.T) {
    assert.Equal(t, 0, engine.LEFT.Value())
    assert.Equal(t, 1, engine.RIGHT.Value())
    assert.Equal(t, 2, engine.UP.Value())
    assert.Equal(t, 3, engine.DOWN.Value())
    assert.Equal(t, 4, engine.ACT.Value())
    assert.Equal(t, 5, engine.STOP.Value())
}

func TestDirectionChangeX(t *testing.T) {
    assert.Equal(t, 0, engine.LEFT.ChangeX(1))
    assert.Equal(t, 2, engine.RIGHT.ChangeX(1))
    assert.Equal(t, 1, engine.UP.ChangeX(1))
    assert.Equal(t, 1, engine.DOWN.ChangeX(1))
}

func TestDirectionChangeY(t *testing.T) {
    assert.Equal(t, 1, engine.LEFT.ChangeY(1))
    assert.Equal(t, 1, engine.RIGHT.ChangeY(1))
    assert.Equal(t, 2, engine.UP.ChangeY(1))
    assert.Equal(t, 0, engine.DOWN.ChangeY(1))
}

func TestDirectionInverted(t *testing.T) {
    assert.Equal(t, engine.RIGHT, engine.LEFT.Inverted())
    assert.Equal(t, engine.LEFT, engine.RIGHT.Inverted())
    assert.Equal(t, engine.DOWN, engine.UP.Inverted())
    assert.Equal(t, engine.UP, engine.DOWN.Inverted())
}

func TestDirectionInvalidInverted(t *testing.T) {
    assert.Panics(t, func() { engine.ACT.Inverted() })
    assert.Panics(t, func() { engine.STOP.Inverted() })
}
