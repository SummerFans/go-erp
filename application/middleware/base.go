package middleware

import (
	"fmt"
	"go-erp/application/handler"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

type baseMiddleware struct {
	handler.Handler
}

// Locale Middleware
func (s *baseMiddleware) Locale() gin.HandlerFunc {

	var matcher = language.NewMatcher([]language.Tag{
		language.English, // The first language is used as fallback.
		language.Chinese,
		language.MustParse("zh-CN"),
		language.MustParse("zh_tw"),
		language.MustParse("zh_HK"),
	})

	return func(c *gin.Context) {

		// lang, _ := c.Cookie("lang")
		lang := "zh-TW"
		fmt.Println(lang)
		accept := c.Request.Header.Get("Accept-Language")

		tag, index := language.MatchStrings(matcher, lang, accept)
		// tag

		fmt.Println(tag)
		fmt.Println(index)
		c.Set("locale", "zh-CN")
		c.Next()
	}
}

func NewBaseMiddleware() *baseMiddleware {
	return &baseMiddleware{}
}
