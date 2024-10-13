// Copyright 2023 The GitBundle Inc. All rights reserved.
// Copyright 2016 The Gitea Authors. All rights reserved.
// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a CC BY-NC 4.0
// license that can be found in the LICENSE file.

package x

import (
	"context"
	"errors"

	"github.com/gitbundle/modules/graceful"
	"github.com/gitbundle/modules/log"
	"github.com/gitbundle/modules/setting"

	"github.com/nsqio/go-nsq"
	"golang.org/x/sync/errgroup"
)

type NsqConsumer struct {
	Topic       string
	Channel     string
	Concurrency int

	nsq.Handler
}

func MustInitNsqConsumer(ctx context.Context, c NsqConsumer) {
	// Instantiate a consumer that will subscribe to the provided channel.
	config := nsq.NewConfig()
	config.AuthSecret = setting.Nsq.AuthSecret
	config.MaxInFlight = 1

	consumer, err := nsq.NewConsumer(c.Topic, c.Channel, config)
	if err != nil {
		log.Fatal("%v", err)
	}

	// Set the Handler for messages received by this Consumer. Can be called multiple times.
	// See also AddConcurrentHandlers.
	consumer.AddConcurrentHandlers(c, c.Concurrency)

	// Use nsqlookupd to discover nsqd instances.
	// See also ConnectToNSQD, ConnectToNSQDs, ConnectToNSQLookupds.
	err = consumer.ConnectToNSQD(setting.Nsq.ClusterTcpAddr)
	if err != nil {
		log.Fatal("%v", err)
	}

	<-ctx.Done()

	// Gracefully stop the consumer.
	consumer.Stop()
}

func mustInitNsqProducer() *nsq.Producer {
	options := nsq.NewConfig()
	options.AuthSecret = setting.Nsq.AuthSecret

	producer, err := nsq.NewProducer(setting.Nsq.ClusterTcpAddr, options)
	if err != nil {
		log.Fatal("nsqProducer: %v", err)
	}

	err = producer.Ping()
	if err != nil {
		log.Fatal("nsqProducer: %v", err)
	}

	return producer
}

var NsqProducer *nsq.Producer

func Init() error {
	NsqProducer = mustInitNsqProducer()

	ctx := graceful.GetManager().ShutdownContext()

	// NOTE: need redis connection, but not need nsq when disable open
	if setting.Nsq.Disable && setting.Redis.Disable {
		return nil
	}
	if setting.Nsq.Disable != setting.Redis.Disable {
		return errors.New("nsq and redis must be enabled or disabled at the same time")
	}

	var (
		g errgroup.Group
	)

	g.Go(func() error {
		log.Trace("Starting nsq producer")
		<-ctx.Done()
		NsqProducer.Stop()
		log.Trace("Exited nsq producer")
		return nil
	})
	graceful.GetManager().RunAtTerminate(func() {
		log.Trace("Exiting nsq producer")
		err := g.Wait()
		if err != nil {
			log.Error("Exiting nsq producer error: %v", err)
		} else {
			log.Trace("Gracefully exited nsq producer")
		}
	})

	return nil
}
