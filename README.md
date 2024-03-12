# go-delay [![Documentation][doc-img]][doc] [![Build Status][ci-img]][ci]

Utility functions for calculating and performing linear backoff, exponential
backoff, rate limiting, and random jitter.

## Installation

```bash
go get github.com/cholland1989/go-delay
```

This library supports [version 1.18 and later][ver] of Go.

## Usage

```go
import "github.com/cholland1989/go-delay/pkg/delay"
```

The `delay` package provides utility functions for calculating linear backoff,
exponential backoff, rate limiting, and random jitter.

`delay.LinearBackoff` calculates the specified duration with linear backoff,
according to the formula `duration * multiplier * (attempt + 1)`. For example,
`delay.LinearBackoff(time.Second, 2.0, 3)` will return a `time.Duration`
representing 8 seconds.

`delay.ExponentialBackoff` calculates the specified duration with exponential
backoff, according to the formula `duration * pow(multiplier, attempt + 1)`.
For example, `delay.ExponentialBackoff(time.Second, 2.0, 3)` will return a
`time.Duration` representing 16 seconds.

`delay.RateLimit` calculates the minimum duration per action for the specified
actions per time period, according to the formula `period / actions`. For
example, `delay.RateLimit(2, time.Second)` will return a `time.Duration`
representing 0.5 seconds.

`delay.RandomJitter` calculates the specified duration plus or minus random
jitter. For example, `delay.RandomJitter(time.Second, 0.5)` will return a
`time.Duration` between 0.5 seconds and 1.5 seconds.

```go
import "github.com/cholland1989/go-delay/pkg/sleep"
```

The `sleep` package provides utility functions for performing linear backoff,
exponential backoff, rate limiting, and random jitter.

`sleep.LinearBackoff` pauses the current goroutine for the specified duration
with linear backoff and random jitter. For example,
`sleep.LinearBackoff(time.Second, 2.0, 0.5, 3)` will pause the current
goroutine for a duration between 4 seconds and 12 seconds.

`sleep.ExponentialBackoff` pauses the current goroutine for the specified
duration with exponential backoff and random jitter. For example,
`sleep.ExponentialBackoff(time.Second, 2.0, 0.5, 3)` will pause the current
goroutine for a duration between 8 seconds and 24 seconds.

`sleep.RateLimit` pauses the current goroutine for the minimum duration per
action with random jitter. For example, `sleep.RateLimit(2, time.Second, 0.5)`
will pause the current goroutine for a duration between 0.25 seconds and 0.75
seconds.

`sleep.RandomJitter` pauses the current goroutine for the specified duration
with random jitter. For example, `sleep.RandomJitter(time.Second, 0.5)` will
pause the current goroutine for a duration between 0.5 seconds and 1.5 seconds.

See the [documentation][doc] for more details.

## License

Released under the [MIT License](LICENSE).

[ci]: https://github.com/cholland1989/go-delay/actions/workflows/build.yml
[ci-img]: https://github.com/cholland1989/go-delay/actions/workflows/build.yml/badge.svg
[doc]: https://pkg.go.dev/github.com/cholland1989/go-delay/pkg/delay
[doc-img]: https://pkg.go.dev/badge/github.com/cholland1989/go-delay/pkg/delay
[ver]: https://go.dev/doc/devel/release
