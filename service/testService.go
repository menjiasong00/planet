package service

import(
	pb "planet/pb"
	"golang.org/x/net/context"
	"fmt"
)


type TestServer struct{}
func (m *TestServer) GetTestMsg(c context.Context, s *pb.TestReq) (*pb.TestResp, error) {
	fmt.Printf("xxxxx(%q)\n", s.Value)
	details := &pb.TestDetail{Value:s.Value}
	return &pb.TestResp{Details:details}, nil
}