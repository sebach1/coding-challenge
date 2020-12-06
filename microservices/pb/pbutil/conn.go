package pbutil



// func conn(ctx context.Context, host string) (*grpc.ClientConn, error) {
// 	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithBlock())
// 	if err != nil {
// 		return nil, status.Errorf(xerrors.Internal, "grpc.Dial: %v at %v", err, host)
// 	}
// 	return conn, nil
// }

// func ConnMessenger(ctx context.Context) (pbmessenger.MessagesClient, func() error, error) {
// 	host := pbconf.Messenger.GetHost()
// 	conn, err := conn(ctx, host)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	c := pbmessenger.NewMessagesClient(conn)
// 	return c, conn.Close, nil
// }

// func ConnAgreement(ctx context.Context) (pbagreement.AgreementClient, func() error, error) {
// 	host := pbconf.Agreement.GetHost()
// 	conn, err := conn(ctx, host)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	c := pbagreement.NewAgreementClient(conn)
// 	return c, conn.Close, nil
// }
