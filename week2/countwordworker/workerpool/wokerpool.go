package workerpool

type Job interface {
	Run()
}

type Worker struct {
	index int
	job   chan Job
	pool  chan chan Job
}

func (w *Worker) start() {
	for {
		w.pool <- w.job
		job := <-w.job
		job.Run()
	}
}

type Pool struct {
	job     chan Job
	pool    chan chan Job
	workers []Worker
}

func (p *Pool) start() {
	for {
		job := <-p.job
		w := <-p.pool
		w <- job
	}
}

func (p *Pool) Dispatch(job Job) {
	p.job <- job
}

func NewWorkerPool(nWorkers int) *Pool {
	p := Pool{
		pool: make(chan chan Job, nWorkers),
		job:  make(chan Job),
	}
	for i := 0; i < nWorkers; i++ {
		p.workers = append(p.workers, Worker{
			pool:  p.pool,
			job:   make(chan Job),
			index: i + 1,
		})
		go p.workers[i].start()
	}
	go p.start()
	return &p
}
