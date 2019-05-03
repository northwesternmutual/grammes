package main

import (
	"flag"

	"github.com/northwesternmutual/grammes/logging"

	"github.com/northwesternmutual/grammes"
	"github.com/northwesternmutual/grammes/examples/exampleutil"
	"go.uber.org/zap"
)

var (
	// addr is used for holding the connection IP address.
	// for example this could be, "ws://127.0.0.1:8182"
	addr string
)

func main() {
	flag.StringVar(&addr, "h", "", "Connection IP")
	flag.Parse()

	logger := exampleutil.SetupLogger()
	defer logger.Sync()

	if addr == "" {
		logger.Fatal("No host address provided. Please run: go run main.go -h <host address>")
		return
	}

	// Create a new Grammes client with a standard websocket.
	client, err := grammes.DialWithWebSocket(addr, grammes.WithLogger(logging.NewDebugLogger()))
	if err != nil {
		logger.Fatal("Couldn't create client", zap.Error(err))
	}

	// DropAll will remove all vertices from the graph currently.
	// Essentially blank slating all of our data.
	// client.DropAll()

	// // Create a waitgroup to wait until all go routines are finished.
	// var wg sync.WaitGroup

	// // Add in 128 vertices to the graph.
	// for i := 0; i < 128; i++ {
	// 	// Add one new task to the waitgroup.
	// 	wg.Add(1)

	// 	// Launch the go routine.
	// 	go func(cli *grammes.Client, index int, logger *zap.Logger, wg *sync.WaitGroup) {
	// 		// Add a vertex to the graph concurrently.
	// 		_, err := cli.AddVertex(fmt.Sprintf("vert%d", index))
	// 		if err != nil {
	// 			logger.Fatal("Couldn't create vertex", zap.Error(err))
	// 		}
	// 		// Send a message that the waitgroup is finished.
	// 		wg.Done()
	// 	}(client, i, logger, &wg)
	// }

	// // Wait until all of the goroutines are finished.
	// wg.Wait()

	// // Log that the goroutines are done.
	// logger.Info("Done.")

	// Count the vertices on the graph
	count, err := client.VertexCount()
	if err != nil {
		logger.Fatal("Couldn't count vertices", zap.Error(err))
	}

	// Log the count of the vertices.
	logger.Info("Counted vertices", zap.Int64("count", count))

	vertices, err := client.AllVertices()
	if err != nil {
		logger.Fatal("Couldn't gather all vertices", zap.Error(err))
	}

	logger.Info("Gathered vertices", zap.Int("length", len(vertices)))
}
