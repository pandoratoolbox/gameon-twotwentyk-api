package main

import (
	"gameon-twotwentyk-api/connections"
	"gameon-twotwentyk-api/graphql"
	"gameon-twotwentyk-api/handlers"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
)

func main() {
	var err error

	r := chi.NewRouter()

	corsParams := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	r.Use(corsParams.Handler)
	r.Use(middleware.Logger)
	r.Use(jwtauth.Verifier(handlers.TokenAuth))
	r.Use(handlers.Authenticator)

	connections.InitPostgres()
	graphql.Init()

	// GenerateArticles(100)
	// err = ImportCelebrities()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	r.Route("/trigger", func(r chi.Router) {
		r.Get("/", handlers.ListTrigger)
	})

	r.Route("/category", func(r chi.Router) {
		r.Get("/", handlers.ListCategory)
	})

	r.Route("/celebrity", func(r chi.Router) {
		r.Get("/", handlers.ListCelebrity)
		r.Group(func(r chi.Router) {
			r.Use(handlers.RestrictAuth)
			r.Post("/", handlers.NewCelebrity)
			r.Route("/{celebrity_id}", func(r chi.Router) {
				r.Get("/", handlers.GetCelebrity)
				r.Put("/", handlers.UpdateCelebrity)
				r.Delete("/", handlers.DeleteCelebrity)
			})
		})
	})

	r.Route("/claim", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(handlers.RestrictAuth)
			r.Post("/", handlers.NewClaim)
			r.Route("/{claim_id}", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Use(handlers.RestrictAdmin)
					r.Post("/approve", handlers.ApproveClaim)
					r.Post("/reject", handlers.RejectClaim)
				})
				r.Get("/", handlers.GetClaim)
				r.Put("/", handlers.UpdateClaim)
				r.Delete("/", handlers.DeleteClaim)
			})
		})
	})

	r.Route("/article", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(handlers.RestrictAuth)
			r.Post("/", handlers.NewArticle)
			r.Route("/{article_id}", func(r chi.Router) {
				r.Get("/", handlers.GetArticle)
				r.Put("/", handlers.UpdateArticle)
				r.Delete("/", handlers.DeleteArticle)
			})
		})
	})

	r.Route("/article_source", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(handlers.RestrictAuth)
			r.Post("/", handlers.NewArticleSource)
			r.Route("/{article_source_id}", func(r chi.Router) {
				r.Get("/", handlers.GetArticleSource)
				r.Put("/", handlers.UpdateArticleSource)
				r.Delete("/", handlers.DeleteArticleSource)
			})
		})
	})

	r.Route("/me", func(r chi.Router) {
		r.Use(handlers.RestrictAuth)
		r.Put("/", handlers.UpdateUser)
		r.Get("/", handlers.GetMyUserData)
		r.Get("/claim", handlers.ListClaimForUserById)
		r.Get("/nft", handlers.GetMyNfts)
		r.Get("/marketplace_listing", handlers.ListMarketplaceListingForUserById)
		r.Get("/recipe/identity", handlers.GetAvailableCelebrityRecipes)
		// r.Get("/recipe/prediction", handlers.GetMyPredictionRecipes)
		r.Get("/nft_card_identity", handlers.ListNftCardIdentityForUserById)
		r.Get("/nft_card_prediction", handlers.ListNftCardPredictionForUserById)
		r.Get("/nft_card_trigger", handlers.ListNftCardTriggerForUserById)
		r.Get("/nft_card_category", handlers.ListNftCardCategoryForUserById)
		r.Get("/nft_card_day_month", handlers.ListNftCardDayMonthForUserById)
		r.Get("/nft_card_year", handlers.ListNftCardYearForUserById)
		r.Get("/nft_card_crafting", handlers.ListNftCardCraftingForUserById)
	})

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", handlers.Login)
		r.Post("/register", handlers.Register)
	})

	r.Route("/feed", func(r chi.Router) {
		r.Get("/", handlers.SearchArticles)
		r.Get("/personalised", handlers.GetArticlesPersonalised)
	})

	// blockchain
	//crafting

	r.Route("/nft", func(r chi.Router) {
		r.Post("/identity", handlers.CraftIdentity)
		r.Post("/prediction", handlers.CraftPrediction)
		r.Post("/trigger", handlers.NewNftCardTrigger)
		r.Post("/day_month", handlers.NewNftCardDayMonth)
		r.Post("/crafting", handlers.NewNftCardCrafting)
		r.Post("/category", handlers.NewNftCardCategory)
		r.Post("/year", handlers.NewNftCardYear)

	})

	r.Route("/marketplace_listing", func(r chi.Router) {
		r.Get("/", handlers.SearchMarketplaceListings)
		r.Group(func(r chi.Router) {
			r.Use(handlers.RestrictAuth)
			r.Post("/", handlers.NewMarketplaceListing)
			r.Route("/marketplace_listing_id}", func(r chi.Router) {
				r.Get("/", handlers.GetMarketplaceListing)
				r.Put("/", handlers.UpdateMarketplaceListing)
				r.Delete("/", handlers.DeleteMarketplaceListing)
			})
		})
	})

	err = http.ListenAndServe(":3333", r)
	if err != nil {
		log.Fatalf("Error serving HTTP handlers: %v", err)
	}

}
