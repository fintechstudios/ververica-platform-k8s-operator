package appManager

import (
	"context"
	"os"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	appManagerApi "github.com/fintechstudios/ververica-platform-k8s-controller/appmanager-api-client"
)

func removeTokenData(slice []testTokenData, s int) []testTokenData {
	return append(slice[:s], slice[s+1:]...)
}

func fakeTokenSecret(role, namespace string, id int) string {
	return tokenName + namespace + strconv.Itoa(id) + role
}

type testTokenData struct {
	namespace string
	name      string
	role      string
	id        int
}

type testTokenManager struct {
	tokens []testTokenData
}

func (m *testTokenManager) TokenExists(ctx context.Context, name, namespace string) (bool, error) {
	for _, tokenData := range m.tokens {
		if tokenData.namespace == namespace && tokenData.name == name {
			return true, nil
		}
	}
	return false, nil
}

func (m *testTokenManager) CreateToken(ctx context.Context, name, role, namespace string) (string, error) {
	lastTokenData := m.tokens[len(m.tokens)-1]
	id := lastTokenData.id + 1
	m.tokens = append(m.tokens, testTokenData{
		namespace,
		name,
		role,
		id,
	})
	return fakeTokenSecret(role, namespace, id), nil
}

func (m *testTokenManager) RemoveToken(ctx context.Context, name, namespace string) (bool, error) {
	var index *int
	for i, tokenData := range m.tokens {
		if tokenData.namespace == namespace && tokenData.name == name {
			index = &i
		}
	}

	if index != nil {
		m.tokens = removeTokenData(m.tokens, *index)
		return true, nil
	}

	return false, nil
}

var _ = Describe("AuthStore", func() {
	const (
		TestNsToken  = "test-ns-TokenData"
		DefaultToken = "no-ns-TokenData"
	)
	setEnv := func() {
		_ = os.Setenv(defaultTokenEnvVar + "_TEST", TestNsToken)
		_ = os.Setenv(defaultTokenEnvVar, DefaultToken)
	}
	unsetEnv := func() {
		_ = os.Unsetenv(defaultTokenEnvVar + "_TEST")
		_ = os.Unsetenv(defaultTokenEnvVar)
	}

	var authStore *AuthStore
	var ctx context.Context
	BeforeEach(func() {
		authStore = NewAuthStore(&PlatformTokenManager{})
		ctx = context.Background()
	})

	Describe("#getTokenForNamespace", func() {
		When("env tokens are present", func() {
			BeforeEach(setEnv)
			AfterEach(unsetEnv)

			It("should should find the token with an exact namespace match", func() {
				token, err := authStore.getTokenForNamespace(ctx, "test")
				Expect(err).To(BeNil())
				Expect(token).To(Equal(TestNsToken))
			})

			It("should should find the token with a default match", func() {
				token, err := authStore.getTokenForNamespace(ctx, "another-ns")
				Expect(err).To(BeNil())
				Expect(token).To(Equal(DefaultToken))
			})

			It("should cache tokens", func() {
				token, err := authStore.getTokenForNamespace(ctx, "test")
				Expect(err).To(BeNil())
				Expect(token).To(Equal(TestNsToken))
				unsetEnv()
				token, err = authStore.getTokenForNamespace(ctx, "test")
				Expect(err).To(BeNil())
				Expect(token).To(Equal(TestNsToken))
			})
		})

		Context("with mocked token manager", func() {
			var manager testTokenManager
			BeforeEach(func() {
				manager = testTokenManager{
					tokens: []testTokenData{
						{
							namespace: "test",
							name:      tokenName,
							role:      "owner",
							id:        1,
						},
					},
				}
				authStore = NewAuthStore(&manager)
			})

			It("should create a token if one doesn't exist", func() {
				token, err := authStore.getTokenForNamespace(ctx, "test-a")
				Expect(err).To(BeNil())
				Expect(token).To(Equal(fakeTokenSecret("owner", "test-a", 2)))
				Expect(len(manager.tokens)).To(Equal(2))
			})

			It("should create a new token and delete the old if one already exist", func() {
				token, err := authStore.getTokenForNamespace(ctx, "test")
				Expect(err).To(BeNil())
				Expect(token).To(Equal(fakeTokenSecret("owner", "test", 2)))
				Expect(len(manager.tokens)).To(Equal(1))
			})
		})
	})

	Describe("#ContextForNamespace", func() {
		When("env tokens are present", func() {
			BeforeEach(setEnv)
			AfterEach(unsetEnv)

			It("should get a context for a namespace", func() {
				ctx, err := authStore.ContextForNamespace(ctx, "test")
				Expect(err).To(BeNil())
				Expect(ctx.Value(appManagerApi.ContextAccessToken)).To(Equal(TestNsToken))
			})
		})
	})
})
