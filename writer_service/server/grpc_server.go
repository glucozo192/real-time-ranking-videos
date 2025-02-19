package server

const (
	maxConnectionIdle = 5
	gRPCTimeout       = 15
	maxConnectionAge  = 5
	gRPCTime          = 10
)

//func (s *server) newWriterGrpcServer() (func() error, *grpc.Server, error) {
//	l, err := net.Listen("tcp", s.cfg.GRPC.Port)
//	if err != nil {
//		return nil, nil, errors.Wrap(err, "net.Listen")
//	}
//
//	grpcServer := grpc.NewServer(
//		grpc.KeepaliveParams(keepalive.ServerParameters{
//			MaxConnectionIdle: maxConnectionIdle * time.Minute,
//			Timeout:           gRPCTimeout * time.Second,
//			MaxConnectionAge:  maxConnectionAge * time.Minute,
//			Time:              gRPCTime * time.Minute,
//		}),
//		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
//			grpc_ctxtags.UnaryServerInterceptor(),
//			grpc_opentracing.UnaryServerInterceptor(),
//			grpc_prometheus.UnaryServerInterceptor,
//			grpc_recovery.UnaryServerInterceptor(),
//			s.im.Logger,
//		),
//		),
//	)
//
//	writerGrpcWriter := grpc2.NewWriterGrpcService(s.log, s.cfg, s.v, s.metrics)
//	proto_buf.writerService.RegisterWriterServiceServer(grpcServer, writerGrpcWriter)
//	grpc_prometheus.Register(grpcServer)
//
//	if s.cfg.GRPC.Development {
//		reflection.Register(grpcServer)
//	}
//
//	go func() {
//		s.log.Infof("Writer gRPC server is listening on port: %s", s.cfg.GRPC.Port)
//		s.log.Fatal(grpcServer.Serve(l))
//	}()
//
//	return l.Close, grpcServer, nil
//}
