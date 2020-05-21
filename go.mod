module github.com/fintechstudios/ververica-platform-k8s-operator

go 1.14

require (
	github.com/antihax/optional v0.0.0-20180407024304-ca021399b1a6
	github.com/go-logr/logr v0.1.0
	github.com/joho/godotenv v1.3.0
	github.com/onsi/ginkgo v1.12.1
	github.com/onsi/gomega v1.10.1
	github.com/pkg/errors v0.9.0 // indirect
	github.com/stretchr/testify v1.5.1
	golang.org/x/net v0.0.0-20200520004742-59133d7f0dd7
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	k8s.io/api v0.18.3
	k8s.io/apimachinery v0.18.3
	k8s.io/client-go v0.18.3
	k8s.io/utils v0.0.0-20200324210504-a9aa75ae1b89
	sigs.k8s.io/controller-runtime v0.6.0
	sigs.k8s.io/structured-merge-diff v0.0.0-20190525122527-15d366b2352e // indirect
)
