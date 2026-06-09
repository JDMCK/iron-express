# How to profile (with flamegraph!)

1. Check out the main code. You basically start up a "profiling server."
    - See https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/
2. Do `curl -o cpu.prof "http://localhost:6060/debug/pprof/profile"`
    - This gets you the profile timing. There are similar options for mem,
    mutexes, etc. etc.
3. Do `go tool pprof -http=:8080 cpu.prof` to run another server 
    - (for displaying the results)
4. Go to localhost:8080/ui/flamegraph
