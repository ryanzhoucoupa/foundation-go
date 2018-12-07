package health

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"encoding/json"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	r  *gin.Engine
	ts *httptest.Server
)

var _ = Describe("health", func() {
	Describe("trivialHealthCheck", func() {
		BeforeEach(func() {
			gin.SetMode(gin.TestMode)
			r = gin.New()
		})

		Context("get health check", func() {
			It("can get a health check", func() {
				HealthInformation.Revision = "x.xx"
				HealthInformation.Version = "y.yy"
				ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(200)
				}))
				r.GET("/health", HealthCheckHandler)
				defer ts.Close()
				request, _ := http.NewRequest("GET", "/health", nil)
				response := httptest.NewRecorder()
				r.ServeHTTP(response, request)
				fmt.Println("response code", response.Code)
				Expect(response.Code).To(Equal(http.StatusOK))
				var result HealthInfo
				json.NewDecoder(response.Body).Decode(&result)
				Expect(response.Body).ShouldNot(BeNil())
				Expect(result.Revision).To(Equal("x.xx"))
				Expect(result.Version).To(Equal("y.yy"))
			})
		})
	})
})

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Health check test Suite")
}
