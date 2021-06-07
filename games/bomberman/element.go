package bomberman

import "github.com/codenjoyme/codenjoy-go-client/engine"

const (
    HERO      engine.Element = '☺' // this is what he usually looks like
    BOMB_HERO                = '☻' // this is if he is sitting on own bomb
    DEAD_HERO                = 'Ѡ' // oops your Hero is dead (don`t worry he will appear somewhere in next move)
    // you`re getting -200 for each death

    OTHER_HERO      = '♥' // this is what other Heroes looks like
    OTHER_BOMB_HERO = '♠' // this is if player just set the bomb
    OTHER_DEAD_HERO = '♣' // enemy corpse (it will disappear shortly right on the next move)
    // if you`ve done it you`ll get +1000

    BOMB_TIMER_5 = '5' // after hero set the bomb the timer starts (5 tacts)
    BOMB_TIMER_4 = '4' // this will blow up after 4 tacts
    BOMB_TIMER_3 = '3' // this after 3
    BOMB_TIMER_2 = '2' // two
    BOMB_TIMER_1 = '1' // one
    BOOM         = '҉' // Boom! this is what is bomb does everything that is destroyable got destroyed

    WALL             = '☼' // indestructible wall - it will not fall from bomb
    DESTROYABLE_WALL = '#' // this wall could be blowed up
    DESTROYED_WALL   = 'H' // this is how broken wall looks like it will dissapear on next move
    // if it`s you did it - you`ll get +10 points.

    MEAT_CHOPPER = '&' // this guys runs over the board randomly and gets in the way all the time
    // if it will touch hero - it will die
    DEAD_MEAT_CHOPPER = 'x' // you`d better kill this piece of ... meat you`ll get +100 point for it
    // this is chopper corpse

    BOMB_BLAST_RADIUS_INCREASE = '+' // Bomb blast radius increase. Applicable only to new bombs. The perk is temporary.
    BOMB_COUNT_INCREASE        = 'c' // Increase available bombs count. Number of extra bombs can be set in settings. Temporary.
    BOMB_REMOTE_CONTROL        = 'r' // Bomb blast not by timer but by second act. Number of RC triggers is limited and can be set in settings.
    BOMB_IMMUNE                = 'i' // Do not die after bomb blast (own bombs and others as well). Temporary.

    NONE = ' ' // this is the only place where you can move your Hero
)
