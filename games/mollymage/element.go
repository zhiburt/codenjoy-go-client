package mollymage

import "github.com/codenjoyme/codenjoy-go-client/engine"

const (
/// your Molly

    // This is what she usually looks like.
    HERO engine.Element = '☺'

    // This is if she is sitting on own potion.
    POTION_HERO engine.Element = '☻'

    // Oops, your Molly is dead (don't worry,
    // she will appear somewhere in next move).
    // You're getting penalty points for each death.
    DEAD_HERO engine.Element = 'Ѡ'

/// other players heroes

    // This is what other heroes looks like.
    OTHER_HERO engine.Element = '♥'

    // This is if player is sitting on own potion.
    OTHER_POTION_HERO engine.Element = '♠'

    // Enemy corpse (it will disappear shortly,
    // right on the next move).
    // If you've done it you'll get score points.
    OTHER_DEAD_HERO engine.Element = '♣'

/// the potions
    // After Molly set the potion, the timer starts (5 ticks).
    POTION_TIMER_5 engine.Element = '5'

    // This will blow up after 4 ticks.
    POTION_TIMER_4 engine.Element = '4'

    // This after 3...
    POTION_TIMER_3 engine.Element = '3'

    // Two..
    POTION_TIMER_2 engine.Element = '2'

    // One.
    POTION_TIMER_1 engine.Element = '1'

    // Boom! this is what is potion does,
    // everything that is destroyable got destroyed.
    BOOM engine.Element = '҉'

/// walls

    // Indestructible wall - it will not fall from potion.
    WALL engine.Element = '☼'

    // this is a treasure box, it opens with an explosion.
    TREASURE_BOX engine.Element = '#'

    // this is like a treasure box opens looks
    // like, it will disappear on next move.
    // if it's you did it - you'll get score
    // points. Perhaps a prize will appear.
    OPENING_TREASURE_BOX engine.Element = 'H'

/// soulless creatures

    // This guys runs over the board randomly
    // and gets in the way all the time.
    // If it will touch Molly - she will die.
    // You'd better kill this piece of ... soul,
    // you'll get score points for it.
    GHOST engine.Element = '&'

    // This is ghost corpse.
    DEAD_GHOST engine.Element = 'x'

/// perks

    // Potion blast radius increase.
    // Applicable only to new potions.
    // The perk is temporary.
    POTION_BLAST_RADIUS_INCREASE engine.Element = '+'

    // Increase available potions count.
    // Number of extra potions can be set
    // in settings. Temporary.
    POTION_COUNT_INCREASE engine.Element = 'c'

    // Potion blast not by timer but by second act.
    // Number of RC triggers is limited and c
    // an be set in settings.
    POTION_REMOTE_CONTROL engine.Element = 'r'

    // Do not die after potion blast
    // (own potion and others as well). Temporary.
    POTION_IMMUNE engine.Element = 'i'

/// a void
    // This is the only place where you can move your Molly.
    NONE engine.Element = ' '
)