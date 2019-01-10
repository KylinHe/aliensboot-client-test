module github.com/KylinHe/aliensboot-client-test

replace (
	cloud.google.com/go v0.26.0 => github.com/GoogleCloudPlatform/gcloud-golang v0.32.0
	github.com/KylinHe/aliensboot-core v0.0.1 => ../aliensboot-core
	github.com/KylinHe/aliensboot-server v0.0.1 => ../aliensboot-demo/src/github.com/KylinHe/aliensboot-server
	golang.org/x/crypto v0.0.0-20180904163835-0709b304e793 => github.com/golang/crypto v0.0.0-20181106171534-e4dc69e5b2fd
	golang.org/x/lint v0.0.0-20180702182130-06c8688daad7 => github.com/golang/lint v0.0.0-20181026193005-c67002cb31c3
	golang.org/x/net => github.com/golang/net v0.0.0-20181108082009-03003ca0c849
	golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be => github.com/golang/oauth2 v0.0.0-20181106182150-f42d05182288
	golang.org/x/sync => github.com/golang/sync v0.0.0-20181108010431-42b317875d0f
	golang.org/x/sys => github.com/golang/sys v0.0.0-20181107165924-66b7b1311ac8
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
	golang.org/x/tools v0.0.0-20180828015842-6cd1fcedba52 => github.com/golang/tools v0.0.0-20181111003725-6d71ab8aade0
	google.golang.org/appengine v1.1.0 => github.com/golang/appengine v1.3.0
	google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8 => github.com/google/go-genproto v0.0.0-20181109154231-b5d43981345b
	google.golang.org/grpc v1.16.0 => github.com/grpc/grpc-go v1.16.0
)

require (
	github.com/KylinHe/aliensboot-core v0.0.2
	github.com/KylinHe/aliensboot-server v0.0.1
	github.com/gogo/protobuf v1.2.0
	github.com/pkg/errors v0.8.0
	github.com/xtaci/kcp-go v5.0.4+incompatible
)
