# Daily Coding Problem: Problem #560 [Easy]  

## Problem statement 

Given a list of numbers and a number k, return whether any two numbers from the
list add up to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

Bonus: Can you do this in one pass?

## Analysis

This is an OK interview question.
There's a more-or-less obvious O(n<sup>2</sup>) solution where the program loops
over the list of numbers.
Inside each loop, you loop over the list of numbers again,
checking each number in the list against every other number in the list
to see if they add up to `k`.
You can see if the interview candidate understands looping, loop indexing,
and maybe some idiosyncracies of whatever programming language you've chosen.
In this solution, the program needs to avoid adding the number under
consideration from the outer loop with itself,
when the program encounters that same number in the inner loop.
Other than that, even a junior level programmer should be able to write a function to do this.

It's probably completely legitimate to expect an experienced programmer to come up
with a single-pass solution when prompted and assured it's possible.
A one-pass solution doesn't require some Robert Tarjan-level obscure algorithm.
It only requires realizing that 
you can subtract the current number from the desired number,
yielding the only number that will sum to the desired number.
The algorithm can check to see if that number has appeared so far
by keeping all numbers seen so far in a map/dict/hashtable.
A single pass will encounter every number in the list,
and because `k - m = n` and `k - n = m`, you're guaranteed to eventually find `m`
(if you're given `n`) in the map.

A more advanced programmer can demonstrate understanding of "single pass",
hash tables, and algorithmic complexity.
