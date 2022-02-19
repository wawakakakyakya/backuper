package task

// func RunAsDaemon(ctx context.Context, c *config.Config, l logger.LoggerInterface) {
// 	l.Info("Run As Daemon")
// 	// child, childCancel := context.WithCancel(ctx)
// 	// defer childCancel()
// 	timer1 := time.NewTimer(3600 * time.Second)
// 	for {
// 		select {
// 		case <-timer1.C:
// 			l.Info("run started")
// 			if err := run(c, l); err != nil {
// 				l.Error(err)
// 			}
// 			l.Info("run finished")
// 		case <-ctx.Done(): // called cancel by signal
// 			l.Info("RunAsDaemon ctx.Done called!")
// 			return
// 		default:
// 			l.Info("waiting 1sec...")
// 			time.Sleep(1 * time.Millisecond)
// 		}
// 	}
// }
