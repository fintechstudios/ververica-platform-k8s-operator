package appManager

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	appManagerApi "github.com/fintechstudios/ververica-platform-k8s-controller/appmanager-api-client"
)

var _ = Describe("AuthStore", func() {
	const (
		TestNsToken  = "test-ns-token"
		DefaultToken = "no-ns-token"
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

	BeforeEach(func() {
		authStore = NewAuthStore()
	})

	Describe("#getTokenForNamespace", func() {
		When("env tokens are present", func() {
			BeforeEach(setEnv)
			AfterEach(unsetEnv)

			It("should should find the token with an exact namespace match", func() {
				token, err := authStore.getTokenForNamespace("test")
				Expect(err).To(BeNil())
				Expect(token).To(Equal(TestNsToken))
			})

			It("should should find the token with a default match", func() {
				token, err := authStore.getTokenForNamespace("another-ns")
				Expect(err).To(BeNil())
				Expect(token).To(Equal(DefaultToken))
			})

			It("should cache tokens", func() {
				token, err := authStore.getTokenForNamespace("test")
				Expect(err).To(BeNil())
				Expect(token).To(Equal(TestNsToken))
				unsetEnv()
				token, err = authStore.getTokenForNamespace("test")
				Expect(err).To(BeNil())
				Expect(token).To(Equal(TestNsToken))
			})
		})

		It("should return an error if no token could be found", func() {
			token, err := authStore.getTokenForNamespace("test")
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
				ctx, err := authStore.ContextForNamespace("test")
				Expect(err).To(BeNil())
				Expect(ctx.Value(appManagerApi.ContextAccessToken)).To(Equal(TestNsToken))
			})
		})

		It("should return an error if no token could be found", func() {
			ctx, err := authStore.ContextForNamespace("test")
			Expect(err).ToNot(BeNil())
			Expect(err.Error()).To(ContainSubstring("namespace test"))
			Expect(ctx).To(BeNil())
		})
	})
})
