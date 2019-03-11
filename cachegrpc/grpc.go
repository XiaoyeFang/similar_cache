package cachegrpc

import (
	"encoding/json"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"similar_cache/cache"
	"similar_cache/protos"
)

type FragmentServer struct {
}

func (*FragmentServer) QuerySimilar(ctx context.Context, request *protos.SimilarRequest) (reply *protos.GeneralReply, err error) {
	reply = &protos.GeneralReply{}

	result, err := cache.QuerySimilar(request.NamePackage, request.Skip, request.Limit, request.Disable, request.EnableGoogleAdsense, request.NePackageName, request.NotNexting)

	result.List = cache.HandSimilar(result.List, request.Skip, request.Limit, request.Disable, request.EnableGoogleAdsense, request.NePackageName, request.NotNexting)
	result.Total = len(result.List)
	jsonStr, err := json.Marshal(result)
	if err != nil {
		glog.Errorf("json.Marshal err %s \n", err)
		return &protos.GeneralReply{}, nil
	}
	reply.Reply = string(jsonStr)
	return reply, err
}

func (*FragmentServer) QueryDeveloper(ctx context.Context, request *protos.DeveloperRequest) (reply *protos.GeneralReply, err error) {
	reply = &protos.GeneralReply{}

	result, err := cache.QueryDeveloper(request.DeveloperId, request.Skip, request.Limit, request.Disable, request.EnableGoogleAdsense, request.NePackageName, request.NotNexting)
	//处理分页
	result.List = cache.HandlerPaging(result.List, request.Skip, request.Limit, request.Disable, request.EnableGoogleAdsense, request.NePackageName, request.NotNexting)
	//result.Total = len(result.List)
	jsonStr, err := json.Marshal(result)
	if err != nil {
		glog.Errorf("json.Marshal err %s \n", err)
	}
	reply.Reply = string(jsonStr)
	//cache.DeleteRediskey()

	return reply, err
}

func (*FragmentServer) QuerySimAndDev(ctx context.Context, request *protos.SimDevRequest) (reply *protos.GeneralReply, err error) {
	reply = &protos.GeneralReply{}

	reply, err = cache.QuerySimiAndDev(request.NamePackage, request.DeveloperId, request.Skip1, request.Skip2, request.Limit1, request.Limit2, request.Disable1, request.Disable2,
		request.EnableGoogleAdsense1, request.EnableGoogleAdsense2, request.NePackageName1, request.NePackageName2, request.NotNexting1, request.NotNexting2)
	//处理分页
	glog.V(5).Infoln("QuerySimAndDev reply===", len(reply.Reply))
	return reply, err
}
