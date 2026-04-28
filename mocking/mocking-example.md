# How `c.sleep(c.duration)` Works

> **Question:** How does `c.sleep(c.duration)` work if we never defined what `sleep()` does — we just put it as a field in the struct?

This is the heart of the mocking lesson. `sleep` is not a method you define — it's a **function stored as a field**.

## The Struct

```go
type ConfigurableSleeper struct {
    duration time.Duration
    sleep    func(time.Duration)  // a function VALUE, not a method
}
```

`sleep` holds any function that accepts a `time.Duration` — like a callback. When you call `c.sleep(c.duration)`, Go calls whatever function was stored in that field, passing `c.duration` to it.

## How It Gets Wired Up

In the test, `ConfigurableSleeper` is constructed like this:

```go
sleeper := &ConfigurableSleeper{sleepTime, spyTime.SetSleepDuration}
```

`spyTime.SetSleepDuration` has the signature `func(time.Duration)`, so it matches the field type. In production you'd pass the real `time.Sleep`, which also has the signature `func(time.Duration)`.

## Why This Is Useful

| Context     | Function injected            | Behavior                              |
|-------------|------------------------------|---------------------------------------|
| Tests       | `spyTime.SetSleepDuration`   | Records the duration, doesn't sleep   |
| Production  | `time.Sleep`                 | Actually sleeps                        |

The struct doesn't care which function it receives — it just calls whatever it was given. This lets you verify the correct duration was used in tests without your test suite actually pausing.
