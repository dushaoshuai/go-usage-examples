https://github.com/google/wire/tree/main/_tutorial
https://github.com/google/wire/tree/main/docs

Wire has two basic concepts: providers and injectors.

Providers are ordinary Go functions that “provide” values given their dependencies, which are described simply as parameters to the function.

Injectors are generated functions that call providers in dependency order.
