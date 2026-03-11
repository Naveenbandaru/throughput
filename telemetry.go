type Telemetry struct {
	ID        int
	Stage     string
	StartTime time.Time
	EndTime   time.Time
	Node      int
}

func stage(stage string, node int, in <-chan int, out chan<- int, uni chan<- Telemetry, wg *sync.WaitGroup) {
	defer wg.Done()
	for id := range in {
		start := time.Now()
		time.Sleep(time.Duration(20+rand.Intn(30)) * time.Millisecond)
		uni <- Telemetry{
			ID:        id,
			Stage:     stage,
			StartTime: start,
			EndTime:   time.Now(),
			Node:      node,
		}
		out <- id
	}
}

func collector(uni <-chan Telemetry, done chan<- bool) {
	count := 0
	for t := range uni {
		fmt.Println(
			"req", t.ID,
			"stage", t.Stage,
			"node", t.Node,
			"latency", t.EndTime.Sub(t.StartTime),
		)
		count++
		if count == 15 {
			done <- true
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	in := make(chan int)
	a := make(chan int)
	b := make(chan int)
	out := make(chan int)
	unified := make(chan Telemetry)
	done := make(chan bool)

	var wg sync.WaitGroup

	wg.Add(3)
	go stage("StageA", 1, in, a, unified, &wg)
	go stage("StageB", 2, a, b, unified, &wg)
	go stage("IO", 3, b, out, unified, &wg)

	go collector(unified, done)

	go func() {
		for i := 0; i < 5; i++ {
			in <- i
		}
		close(in)
	}()

	go func() {
		wg.Wait()
		close(unified)
	}()

	<-done
	time.Sleep(200 * time.Millisecond)
}
