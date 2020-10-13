module github.com/fintechstudios/ververica-platform-k8s-operator

go 1.14

require (
	github.com/antihax/optional v0.0.0-20180407024304-ca021399b1a6
	github.com/go-logr/logr v0.1.0
	github.com/joho/godotenv v1.3.0
	github.com/onsi/ginkgo v1.14.0
	github.com/onsi/gomega v1.10.3
	github.com/pkg/errors v0.9.0 // indirect
	github.com/stretchr/testify v1.6.1
	golang.org/x/net v0.0.0-20201006153459-a7d1128ccaa0
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	k8s.io/api v0.18.8
	k8s.io/apimachinery v0.18.8
	k8s.io/client-go v0.18.6
	k8s.io/utils v0.0.0-20200603063816-c1c6865ac451
	sigs.k8s.io/controller-runtime v0.6.3
	sigs.k8s.io/structured-merge-diff v0.0.0-20190525122527-15d366b2352e // indirect
)
