package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"kube-opt/protol/pb/kube_opt_pb"
	"os"
	"time"
)

var ctx = context.Background()

var conn *grpc.ClientConn

// Set your grpc server address
var addr string

var keepAliveCfg = keepalive.ClientParameters{
	Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
	Timeout:             30 * time.Second, // wait 1 second for ping ack before considering the connection dead
	PermitWithoutStream: true,             // send pings even without active streams
}

func GetPods(namespace string) (*kube_opt_pb.KubeOptResp, error) {

	kubeOptCli := kube_opt_pb.NewKubeOptServiceClient(conn)
	resp, err := kubeOptCli.Pods(ctx, &kube_opt_pb.KubeOptReq{
		Namespace: namespace,
	})
	return resp, err
}

func DelPod(namespace, pod string) (*kube_opt_pb.KubeOptResp, error) {
	kubeOptCli := kube_opt_pb.NewKubeOptServiceClient(conn)
	resp, err := kubeOptCli.DeletePod(ctx, &kube_opt_pb.KubeOptReq{
		PodId:     pod,
		Namespace: namespace,
	})
	return resp, err
}

func init() {
	addr = os.Getenv("INPUT_SERVER")
	if len(addr) == 0 {
		println("grpc server must be specific")
		os.Exit(1)
	}
	_conn, err := grpc.Dial(addr, grpc.WithKeepaliveParams(keepAliveCfg), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		println(fmt.Sprintf("Connect grpc server failed:%s", err.Error()))
		os.Exit(1)
	}
	conn = _conn
}
