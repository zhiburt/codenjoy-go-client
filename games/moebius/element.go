package moebius

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

var elements = map[string]rune{

        // Труба повернута слева вверх.

    "LEFT_UP": '╝',

        // Труба повернута сверху направо.

    "UP_RIGHT": '╚',

        // Труба повернута справа вниз.

    "RIGHT_DOWN": '╔',

        // Труба повернута снизу влево.

    "DOWN_LEFT": '╗',

        // Прямая труба слева направо.

    "LEFT_RIGHT": '═',

        // Прямая труба сверху вниз.

    "UP_DOWN": '║',

        // Две пересеченные прямые трубы, одна сверху вниз, другая
        // слева направо.

    "CROSS": '╬',

        // Пустое место на поле.

    "EMPTY": ' ',

}
