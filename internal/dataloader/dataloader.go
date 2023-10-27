package dataloader

// import graph gophers with your other imports
import (
	"context"
	"go-gql-sample/app/ent"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/dataloader"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)
type Loaders struct {
	UserLoader *dataloader.Loader
}

// NewLoaders instantiates data loaders for the middleware
func NewLoaders(client *ent.Client) *Loaders {
	// define the data loader
	userReader := &UserReader{Client: client}
	loaders := &Loaders{
		UserLoader: dataloader.NewBatchedLoader(userReader.GetUsers),
	}
	return loaders
}

func Middleware(loaders *Loaders) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "dataloaders", loaders)
		log.Print("Middleware")
		log.Print(ctx.Value("dataloaders"))
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// For returns the dataloader for a given context
func For(ctx context.Context) *Loaders {
	ginContext := ctx.Value("dataloaders")

	return ginContext.(*Loaders)
}