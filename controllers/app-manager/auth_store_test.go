package appManager

import (
	"context"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	appManagerApi "github.com/fintechstudios/ververica-platform-k8s-controller/appmanager-api-client"
)

type TestTokenManager struct {
}

func (p *TestTokenManager) TokenExists(ctx context.Context, name, namespace string) (bool, error) {
	return false, nil
}

func (p *TestTokenManager) CreateToken(ctx context.Context, name, role, namespace string) (string, error) {
	return "", nil
}

func (p *TestTokenManager) GetToken(ctx context.Context, name, namespace string) (string, error) {
	return "", nil
}

func (p *TestTokenManager) RemoveToken(ctx context.Context, name, namespace string) (bool, error) {
	return false, nil
}

func (p *TestTokenManager) ListAllTokens(context.Context) ([]string, error) {
	return nil, nil
}

var _ = Describe("AuthStore", func() {
	const (
		TestNsToken  = "test-ns-TokenData"
		DefaultToken = "no-ns-TokenData"
	)
	setEnv := func() {
		_ = os.Setenv("VP_API_TOKEN_TEST", TestNsToken)
		_ = os.Setenv("VP_API_TOKEN", DefaultToken)
	}
	unsetEnv := func() {
		_ = os.Unsetenv("VP_API_TOKEN_TEST")
		_ = os.Unsetenv("VP_API_TOKEN")
	}

	var authStore *AuthStore
	var ctx context.Context
	BeforeEach(func() {
		authStore = NewAuthStore(&TestTokenManager{})
		ctx = context.Background()
	})

	Describe("#getTokenForNamespace", func() {
		When("env tokens are present", func() {
			BeforeEach(setEnv)
			AfterEach(unsetEnv)

			It("should should find the TokenData with an exact namespace match", func() {
				token, err := authStore.getTokenForNamespace(ctx, "test")
				Expect(err).To(BeNil())
				Expect(token).To(Equal(TestNsToken))
			})

			It("should should find the TokenData with a default match", func() {
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

		It("should return an error if no TokenData could be found", func() {
			token, err := authStore.getTokenForNamespace(ctx, "test")
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(ContainSubstring("namespace test"))
			Expect(token).To(HaveLen(0))
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

		It("should return an error if no TokenData could be found", func() {
			ctx, err := authStore.ContextForNamespace(ctx, "test")
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(ContainSubstring("namespace test"))
			Expect(ctx).To(BeNil())
		})
	})
})
