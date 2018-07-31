package antifraud

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sonm-io/core/connor_v2/price"
	"github.com/sonm-io/core/proto"
	"github.com/sonm-io/core/util"
	"go.uber.org/zap"
)

type nanoPoolWatcher struct {
	log      *zap.Logger
	wallet   common.Address
	workerID string
}

func NewNanopoolWatcher(log *zap.Logger, deal *sonm.Deal) *nanoPoolWatcher {
	workerID := fmt.Sprintf("CONNOR_%s", deal.GetId().Unwrap().String())
	l := log.Named("nanopool").With(
		zap.String("deal_id", deal.GetId().Unwrap().String()),
		zap.String("worker_id", workerID))

	return &nanoPoolWatcher{
		log:      l,
		workerID: workerID,
		wallet:   deal.GetConsumerID().Unwrap(),
	}
}

func (w *nanoPoolWatcher) Run(ctx context.Context) error {
	t := util.NewImmediateTicker(1 * time.Minute)
	defer t.Stop()

	w.log.Info("starting pool watcher", zap.String("wallet", w.wallet.Hex()))

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-t.C:
			if err := w.watch(); err != nil {
				w.log.Warn("failed to load nanopool's data", zap.Error(err))
			}
		}
	}

}

func (w *nanoPoolWatcher) watch() error {
	url := fmt.Sprintf("https://api.nanopool.org/v1/eth/user/%s", w.wallet.Hex())
	data, err := price.FetchURLWithRetry(url)
	if err != nil {
		return fmt.Errorf("failed to fetch nanopool data: %v", err)
	}

	resp := &nanoPoolResponse{}
	if err := json.Unmarshal(data, resp); err != nil {
		return fmt.Errorf("failed to parse nanopool response: %v", err)
	}

	var nanoWorker *nanopoolWorker
	for _, worker := range resp.Data.Workers {
		if worker.ID == w.workerID {
			nanoWorker = worker
		}
	}

	if nanoWorker == nil {
		return fmt.Errorf("cannot find worker %s in reponse data", w.workerID)
	}

	w.log.Info("reported hashrate", zap.Any("worker", *nanoWorker))
	return nil
}

type nanopoolWorker struct {
	ID  string `json:"id"`
	H1  string `json:"h1"`
	H3  string `json:"h3"`
	H6  string `json:"h6"`
	H12 string `json:"h12"`
	H24 string `json:"h24"`
}

type nanoPoolResponse struct {
	Status bool `json:"status"`
	Data   struct {
		Workers []*nanopoolWorker `json:"workers"`
	} `json:"data"`
}
