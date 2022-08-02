We spawn external processes when we need an externl process accessible to a running Go process.
Sometimes we just want to completely replace the current Go process with another (perhaps non-Go) one.
To do this we use Go's implementation fo the classic exec function.
