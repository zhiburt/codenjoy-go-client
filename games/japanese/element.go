package japanese

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

        // Игрок утверждает, что пиксель белый.

    "WHITE": '-',

        // Игрок утверждает, что пиксель черный.

    "BLACK": '*',

        // Игрок пока не определился, какого цвета этот пиксель.

    "UNSET": ' ',

        // Пустое место в полое для цифр.

    "NAN": '.',

        // Блок отсутствует.

    "_0": '0',

        // Блок длинной в 1 пиксель.

    "_1": '1',

        // Блок длинной в 2 пикселя.

    "_2": '2',

        // Блок длинной в 3 пикселя.

    "_3": '3',

        // Блок длинной в 4 пикселя.

    "_4": '4',

        // Блок длинной в 5 пикселей.

    "_5": '5',

        // Блок длинной в 6 пикселей.

    "_6": '6',

        // Блок длинной в 7 пикселей.

    "_7": '7',

        // Блок длинной в 8 пикселей.

    "_8": '8',

        // Блок длинной в 9 пикселей.

    "_9": '9',

        // Блок длинной в 10 пикселей.

    "_10": 'a',

        // Блок длинной в 11 пикселей.

    "_11": 'b',

        // Блок длинной в 12 пикселей.

    "_12": 'c',

        // Блок длинной в 13 пикселей.

    "_13": 'd',

        // Блок длинной в 14 пикселей.

    "_14": 'e',

        // Блок длинной в 15 пикселей.

    "_15": 'f',

        // Блок длинной в 16 пикселей.

    "_16": 'g',

        // Блок длинной в 17 пикселей.

    "_17": 'h',

        // Блок длинной в 18 пикселей.

    "_18": 'i',

        // Блок длинной в 19 пикселей.

    "_19": 'j',

        // Блок длинной в 20 пикселей.

    "_20": 'k',

        // Блок длинной в 21 пиксель.

    "_21": 'l',

        // Блок длинной в 22 пикселя.

    "_22": 'm',

        // Блок длинной в 23 пикселя.

    "_23": 'n',

        // Блок длинной в 24 пикселя.

    "_24": 'o',

        // Блок длинной в 25 пикселей.

    "_25": 'p',

        // Блок длинной в 26 пикселей.

    "_26": 'q',

        // Блок длинной в 27 пикселей.

    "_27": 'r',

        // Блок длинной в 28 пикселей.

    "_28": 's',

        // Блок длинной в 29 пикселей.

    "_29": 't',

        // Блок длинной в 30 пикселей.

    "_30": 'u',

        // Блок длинной в 31 пиксель.

    "_31": 'v',

        // Блок длинной в 32 пикселя.

    "_32": 'w',

        // Блок длинной в 33 пикселя.

    "_33": 'x',

        // Блок длинной в 34 пикселя.

    "_34": 'y',

        // Блок длинной в 35 пикселей.

    "_35": 'z',

        // Блок длинной в 36 пикселей.

    "_36": 'A',

        // Блок длинной в 37 пикселей.

    "_37": 'B',

        // Блок длинной в 38 пикселей.

    "_38": 'C',

        // Блок длинной в 39 пикселей.

    "_39": 'D',

        // Блок длинной в 40 пикселей.

    "_40": 'E',

        // Блок длинной в 41 пиксель.

    "_41": 'F',

        // Блок длинной в 42 пикселя.

    "_42": 'G',

        // Блок длинной в 43 пикселя.

    "_43": 'H',

        // Блок длинной в 44 пикселя.

    "_44": 'I',

        // Блок длинной в 45 пикселей.

    "_45": 'J',

        // Блок длинной в 46 пикселей.

    "_46": 'K',

        // Блок длинной в 47 пикселей.

    "_47": 'L',

        // Блок длинной в 48 пикселей.

    "_48": 'M',

        // Блок длинной в 49 пикселей.

    "_49": 'N',

        // Блок длинной в 50 пикселей.

    "_50": 'O',

        // Блок длинной в 51 пиксель.

    "_51": 'P',

        // Блок длинной в 52 пикселя.

    "_52": 'Q',

        // Блок длинной в 53 пикселя.

    "_53": 'R',

        // Блок длинной в 54 пикселя.

    "_54": 'S',

        // Блок длинной в 55 пикселей.

    "_55": 'T',

        // Блок длинной в 56 пикселей.

    "_56": 'U',

        // Блок длинной в 57 пикселей.

    "_57": 'V',

        // Блок длинной в 58 пикселей.

    "_58": 'W',

        // Блок длинной в 59 пикселей.

    "_59": 'X',

        // Блок длинной в 60 пикселей.

    "_60": 'Y',

        // Блок длинной в 61 пиксель.

    "_61": 'Z',

}
