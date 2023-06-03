// Copyright (C) 2018-2023, John Chadwick <john@jchw.io>
//
// Permission to use, copy, modify, and/or distribute this software for any purpose
// with or without fee is hereby granted, provided that the above copyright notice
// and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
// REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY AND
// FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
// INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS
// OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER
// TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE OF
// THIS SOFTWARE.
//
// SPDX-FileCopyrightText: Copyright (c) 2018-2023 John Chadwick
// SPDX-License-Identifier: ISC

package minibox

import (
	"context"
	"net/http"

	"github.com/pangbox/server/qa/authserv"
	log "github.com/sirupsen/logrus"
)

type QAAuthOptions struct {
	Addr string
}

type QAAuthServer struct {
	service *Service
}

func NewQAAuth(ctx context.Context) *QAAuthServer {
	qaauth := new(QAAuthServer)
	qaauth.service = NewService(ctx)
	return qaauth
}

func (q *QAAuthServer) Configure(opts QAAuthOptions) error {
	spawn := func(ctx context.Context, service *Service) {
		qaAuthServer := http.Server{Addr: opts.Addr, Handler: authserv.New()}

		service.SetShutdownFunc(func(shutdownCtx context.Context) error {
			return qaAuthServer.Shutdown(shutdownCtx)
		})

		if ctx.Err() != nil {
			log.Errorf("QA Auth service cancelled before server could start: %v", ctx.Err())
			return
		}

		err := qaAuthServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Errorf("Error serving QA Auth server: %v", err)
		}
	}

	return q.service.Configure(spawn)
}

func (q *QAAuthServer) Running() bool {
	return q.service.Running()
}

func (q *QAAuthServer) Start() error {
	return q.service.Start()
}

func (q *QAAuthServer) Stop() error {
	return q.service.Stop()
}