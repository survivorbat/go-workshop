# Patterns

Hey welcome to the last section!
By now you should have a working API that allows you to retrieve, filter
and add people to an in-memory dataset.
There are however, a few issues with this code, mainly
its testability.

- [Patterns](#patterns)
  - [Mocking](#mocking)

---

## Dependency Inversion

Function `GetPeopleRoute` makes a call to `getPeople`, that means
that during a test, I'm also implicitly testing getPeople right?

What happens if I make calls to external libraries? Or 
if I make use of a database, what happens if I have
a route connect to a database, does that mean I must always
be connected to a database?

In other languages, you'd probably be using a Mocking framework
to prevent making any real calls to a database.
In Go this is not really an option, so we have to
be smart instead.
How do we mock this function? Interfaces.

The `exercises.md` will take you through this journey.


