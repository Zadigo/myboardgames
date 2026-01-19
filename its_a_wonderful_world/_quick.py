

from dataclasses import dataclass


@dataclass
class A:
    @property
    def hey(self):
        return "hi"


@dataclass
class B(A):
    @property
    def hey(self):
        return 'hoo'


b = B()
print(b.hey)
