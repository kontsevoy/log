### Ease of use

The top 3 feaures of the logger are:

- Ease of configuring the output (stderr, syslog, stdout)
- Ease of setting logging prefix: globally and per-output
- Ease of setting logging priority: globally and per-output

### Priorities

Using `log.Write` or `log.Println` style functions is not helpful.
Errors, warnings and infos must be explicit, like this:

```
err := not_important_func()
if err != nil {
    log.Info("something not important failed: %v", err)
}

err = important_func()
if err != nil {
    log.Error("something IMPORTANT failed: %v", err)
}
```

Then you can turn on/off errors or warnings by setting the priority
level via `log.SetPriority()`

### Compatibility

Reuse as many of existing utilities as possible. For example:

* syslog already has a defined table of priorities. use it.
* make sure standard `log.Setwriter()` works with this log

