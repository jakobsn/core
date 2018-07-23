package connor

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sonm-io/core/connor_v2/price"
	"github.com/sonm-io/core/insonmnia/auth"
	"github.com/sonm-io/core/proto"
	"github.com/sonm-io/core/util"
	"github.com/sonm-io/core/util/xgrpc"
	"go.uber.org/zap"
)

type Connor struct {
	cfg *Config
	log *zap.Logger
	key *ecdsa.PrivateKey

	snmPriceProvider   price.Provider
	tokenPriceProvider price.Provider
	orderManager       *orderManager
	marketClient       sonm.MarketClient
	dealsClient        sonm.DealManagementClient
}

func New(ctx context.Context, cfg *Config, log *zap.Logger) (*Connor, error) {
	log.Debug("building new instance", zap.Any("config", *cfg))

	key, err := cfg.Eth.LoadKey()
	if err != nil {
		return nil, fmt.Errorf("cannot load eth keys: %v", err)
	}

	_, TLSConfig, err := util.NewHitlessCertRotator(context.Background(), key)
	if err != nil {
		return nil, fmt.Errorf("cannot create cert TLS config: %v", err)
	}

	creds := auth.NewWalletAuthenticator(util.NewTLS(TLSConfig), crypto.PubkeyToAddress(key.PublicKey))
	cc, err := xgrpc.NewClient(ctx, cfg.Node.Endpoint.String(), creds)
	if err != nil {
		return nil, fmt.Errorf("cannot create connection to node: %v", err)
	}

	return &Connor{
		key:                key,
		cfg:                cfg,
		log:                log,
		marketClient:       sonm.NewMarketClient(cc),
		dealsClient:        sonm.NewDealManagementClient(cc),
		snmPriceProvider:   price.NewSonmPriceProvider(),
		tokenPriceProvider: price.NewProvider(cfg.Mining.Token),
		orderManager:       NewOrderManager(ctx, log, nil),
	}, nil
}

func (c *Connor) Serve(ctx context.Context) error {
	c.log.Debug("starting Connor instance")

	if err := c.loadInitialData(ctx); err != nil {
		return fmt.Errorf("initializind failed: %v", err)
	}

	c.log.Debug("price", zap.String("SNM", c.snmPriceProvider.GetPrice().String()),
		zap.String(c.cfg.Mining.Token, c.tokenPriceProvider.GetPrice().String()))

	// todo: detach in background
	go c.orderManager.start(ctx)

	//newOrder := newBidTemplate(c.cfg.Mining.Token)
	//c.log.Debug("steps", zap.Uint64("from", c.cfg.Market.FromHashRate),
	//	zap.Uint64("to", c.cfg.Market.ToHashRate),
	//	zap.Uint64("step", c.cfg.Market.Step))

	for hr := c.cfg.Market.FromHashRate; hr <= c.cfg.Market.ToHashRate; hr += c.cfg.Market.Step {
		p := big.NewInt(0).Mul(c.tokenPriceProvider.GetPrice(), big.NewInt(int64(hr)))
		c.log.Debug("requesting initial order placement", zap.String("price", p.String()), zap.Uint64("hashrate", hr))
		NewCorderFromParams(c.cfg.Mining.Token, p, hr)
		// c.orderManager.Create(newOrder(p, hr))
	}

	<-ctx.Done()
	c.close()
	return nil
}

func (c *Connor) loadInitialData(ctx context.Context) error {
	if err := c.snmPriceProvider.Update(ctx); err != nil {
		return fmt.Errorf("cannot update SNM price: %v", err)
	}

	if err := c.tokenPriceProvider.Update(ctx); err != nil {
		return fmt.Errorf("cannot update %s price: %v", c.cfg.Mining.Token, err)
	}

	return nil
}

type ordersSets struct {
	orderToCreate   []*Corder
	ordersToRestore []*Corder
}

func (c *Connor) divideOrdersSets(ctx context.Context) (*ordersSets, error) {
	exitingOrders, err := c.marketClient.GetOrders(ctx, &sonm.Count{Count: 1000})
	if err != nil {
		return nil, fmt.Errorf("cannot load orders from market: %v", err)
	}

	exitingCorders := NewCordersSlice(exitingOrders.GetOrders(), c.cfg.Mining.Token)
	// todo: sdghdjfghdjkfg
	_ = exitingCorders
	return nil, nil
}

func (c *Connor) getTargetCorders() []*Corder {
	v := make([]*Corder, 0)

	for hr := c.cfg.Market.FromHashRate; hr <= c.cfg.Market.ToHashRate; hr += c.cfg.Market.Step {
		bigHashrate := big.NewInt(int64(hr))
		p := big.NewInt(0).Mul(bigHashrate, c.tokenPriceProvider.GetPrice())
		order, _ := NewCorderFromParams(c.cfg.Mining.Token, p, hr)
		v = append(v, order)
	}

	return v
}

func (c *Connor) close() {
	c.log.Debug("closing Connor")
}
