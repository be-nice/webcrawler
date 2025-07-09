package pkg

type Crawler struct {
	Pages   map[string]int
	BaseURL string
	Config
}

type Config struct {
	MaxWorkers int
	MaxPages   int
}

type KeyVal struct {
	k string
	v int
}
