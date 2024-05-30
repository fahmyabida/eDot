package app

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var CancelOrderWorkerCommand = &cobra.Command{
	Use:   "cancel-order-worker",
	Short: "Worker for checking order that exceeded payment time",
	Run:   runCancelOrderWorker,
}

func init() {
	rootCmd.AddCommand(CancelOrderWorkerCommand)
}

func runCancelOrderWorker(cmd *cobra.Command, args []string) {
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Done()
	go func() {
		for {
			log.Default().Printf("order worker to cancel order, start....")
			err := OrderUsecase.CancelExceededOrder(context.Background())
			if err != nil {
				log.Default().Printf("error occurred during cancel order worker: %v", err)
			}
			log.Default().Printf("order worker to cancel order, end....")
			time.Sleep(1 * time.Minute)
		}
	}()
	wg.Wait()
}
