package main

import (
				"net"
				"io"
				"log/slog"
)


type Server struct {
				ln net.Listener
}

func (s *Server) Start() error {
				return nil
}

func (s* Server) listen() error {
				ln , err := net.Listen("tcp", ":9092")
				if err != nil {
								return err
				}
				s.ln = ln
				for {
								conn, err := ln.Accept()
								if err != nil {
												if err == io.EOF {
																return err
												}
												slog.Error("server accept error", "err", err)
								}
								go s.handleConn(conn)
				}
}
