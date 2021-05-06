module github.com/rancher/rancher/pkg/apis

go 1.14

replace k8s.io/client-go => k8s.io/client-go v0.21.0

require (
	github.com/pkg/errors v0.9.1
	github.com/rancher/eks-operator v1.0.6-rc1
	github.com/rancher/fleet/pkg/apis v0.0.0-20210428191153-f414eab0e4de
	github.com/rancher/gke-operator v1.0.1
	github.com/rancher/norman v0.0.0-20210504005327-7b74a9f308a7
	github.com/rancher/rke v1.3.0-rc1.0.20210503155726-c25848db1e86
	github.com/rancher/wrangler v0.8.1-0.20210506052526-673b7f8692d9
	github.com/sirupsen/logrus v1.7.0
	k8s.io/api v0.21.0
	k8s.io/apimachinery v0.21.0
	sigs.k8s.io/cluster-api v0.3.11-0.20210430180359-45b6080c2764
)
