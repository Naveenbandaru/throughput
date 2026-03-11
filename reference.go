type Trace struct {
	ID    int
	Stage string
	Time  time.Duration
}
func stage(name string, in <-chan int, out chan<- int,
	metrics chan<- string, logs chan<- string, traces chan<- Trace) {
	for id := range in {
		start := time.Now()
		time.Sleep(20 * time.Millisecond)
		metrics <- name
		logs <- fmt.Sprintf("req %d at %s", id, name)
		traces <- Trace{ID: id, Stage: name, Time: time.Since(start)}
		out <- id
	}
}
func main() {
	in := make(chan int)
	a := make(chan int)
	b := make(chan int)
	out := make(chan int)
	metrics := make(chan string)
	logs := make(chan string)
	traces := make(chan Trace)
	go stage("StageA", in, a, metrics, logs, traces)
	go stage("StageB", a, b, metrics, logs, traces)
	go stage("IO", b, out, metrics, logs, traces)
	go func() {
		for {
			select {
			case m := <-metrics:
				fmt.Println("metric", m)
			case l := <-logs:
				fmt.Println("log", l)
			case t := <-traces:
				fmt.Println("trace", t.ID, t.Stage, t.Time)
			}
		}
	}()
	for i := 0; i < 5; i++ {
		in <- i
	}
	time.Sleep(time.Second)
}
