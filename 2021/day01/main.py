from pipe_operator import elixir_pipe,then
import itertools
import operator
from collections.abc import Iterable
from typing import TypeVar

to_int = lambda x: map(int,x)
map_sum = lambda x: map(sum,x)
map_diff = lambda x: map(lambda y: operator.sub(operator.getitem(y, 1),operator.getitem(y, 0)),x)
filter_positive = lambda x: filter(lambda y: operator.gt(y, 0), x)

T = TypeVar("T")

def sliding_window(iterable:Iterable[T], n: int) -> Iterable[T]:
    iterables = itertools.tee(iterable, n)
    for i, it in enumerate(iterables):
        next(itertools.islice(it, i, i), None)
    return zip(*iterables)

@elixir_pipe
def part_1(fp: str) -> int:
    with open(fp) as f:
        (
                f 
                >> _.read() 
                >> _.splitlines() 
                >> then(to_int)
                >> itertools.pairwise
                >> map_diff
                >> filter_positive
                >> list
                >> len
                >> print
        )
    return 0

@elixir_pipe
def part_2(fp: str) -> int:
    with open(fp) as f:
        (
                f 
                >> _.read() 
                >> _.splitlines() 
                >> then(to_int)
                >> sliding_window(n=3)
                >> map_sum
                >> itertools.pairwise
                >> map_diff
                >> filter_positive
                >> list
                >> len
                >> print
        )
    return 0


part_1("input.txt")
part_2("input.txt")
