Sometimes we'd like our Go programs to intelligently handle Unix signals.
For example, we might want a server to gracefully shutdown when it receives a SIGTERM, or a command=line tool to stop processing input if it receives a SIGINT.
