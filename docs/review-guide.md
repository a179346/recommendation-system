This file discusses some decision-making aspects. I believe it will help you better understand the system.

- [Three-layer architecture](#three-layer-architecture)
- [Dependency Injection (DI)](#dependency-injection)
- [Packages](#packages)
- [Project layout](#project-layout)

## Three-layer architecture

The system is structured into three layers to achieve [separation of concerns](https://en.wikipedia.org/wiki/Separation_of_concerns):

- `Handler Layer` – **Manages input/output**, request validation, and response formatting.
- `Logic Layer` – **Implements business logic**
- `Provider Layer` – **Handles I/O operations** like database queries, Redis access, and JWT generation.

## Dependency Injection

[Dependency Injection](https://en.wikipedia.org/wiki/Dependency_injection) can improve testability.

## Packages

- [Echo](#echo): Go web framework
- [Jet](#jet): Type safe SQL builder

#### Echo

I chose [Echo](https://echo.labstack.com/) because I feel it has a similar design to [GoHF](https://github.com/gohf-http/gohf), another HTTP framework I built, and I wanted to explore something new.

#### Jet

##### Why I chose Query Builder over ORM:

I am not a fan of ORM as it introduces a black box to the project. When the project gets large, it would be challenging to optimize the performance.

##### Why I chose Query Builder over Raw SQL:

Query builder makes the application support different database systems. I don't need to worry about syntax differences between MySQL and PostgreSQL.

## Project Layout

TODO
