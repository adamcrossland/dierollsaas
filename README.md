# dierollsaas
## A simple web service written in Go that provides AD&D style die rolls

**This code base is intended to be deployed as a Google AppEngine application. It is not
structured like a normal Go application would, and some of the files in the codebase are
AppEngine-specific.**

## About dierollsaas
*DieRollSaaS* is a service that provides randomly-generated numbers. The numbers
to be generated are specified in a format common to fantasy role-playing
games: d6, 2d6, 1d4+1, and so on.

Specifically, __d__ is followed by a number that nominally represents the number
of sides on the die to be rolled; d6 would be a familiar, six-sided cube. While
fantasy role-playing games commonly use d4, d6, d8, d10, d12 and d20, this 
service will allow dice to have any number of sides. While a d31 would make for
a very strangely-shaped die, *DieRollSaaS* is perfectly happy to roll one for you.

The optional number that precedes the __d__
(like __2d6__) tells the number of dice to throw. A roll of __2d6__ means two
six-sided dice.

The number that results from adding together the values of all of the rolled
dice can be modified up or down by any number by appending a plus or minus.
For example, 2d8+1 means roll two eight-sided dice, add the two numbers together
and add one, resulting in a value between 3 and 17. Or, you might want __2d4-1__
which means roll two four-sided dice and
subtract one, giving a range of values from 1 to 7.

A novel addition to the classic format is the times specifier. To request that
the given set of rolls be performed repeatedly, simply append an
__x__ followed
by the number of times. For example: __3d6x6__ would roll 3d6 6 times.

A second addition is the 'best-of' modifer. Occasionally, one will need to
roll a certain number of dice and take only a subset of highest values. For
example, when generating charaacter attributes in the range of 3-18, one
might choose to roll four six-sided dice and use the highest 3 values.
That case is accounted for like this:
__3,4d6__. To get an entire set of 6 attributes,
one would use __3,4d6x6__.

This service is written in [Go](http://golang.org) to run on
[Google AppEngine](https://cloud.google.com/appengine/docs), but it can easily
be adapted to on any web platform.

## Try it out
This code is currently running at: [adamcrossland.net/rolldice](http://rolldice.adamcrossland.net/).

To make a roll, just make a __GET__ request to

    http://rolldice.adamcrossland.net/roll/spec
    
    or

    http://dierolsaas.appspot.com/roll/spec

where __spec__ is the die roll that you want, like this:

    http://rolldice.adamcrossland.net/roll/2d6+1x10

The rolls are returned in a simple JSON format. Here is a sample return for a roll of 2d6+1x10:

    {"count":10,
	    "rolls":[
		    {"total":7,"count":2,"dice":[4,2]},
		    {"total":7,"count":2,"dice":[4,2]},
		    {"total":4,"count":2,"dice":[1,2]},
		    {"total":8,"count":2,"dice":[2,5]},
		    {"total":5,"count":2,"dice":[1,3]},
		    {"total":5,"count":2,"dice":[2,2]},
		    {"total":5,"count":2,"dice":[2,2]},
		    {"total":9,"count":2,"dice":[2,6]},
		    {"total":11,"count":2,"dice":[4,6]},
		    {"total":9,"count":2,"dice":[4,4]}
	    ]
    }
