package main

import (
	"gameon-twotwentyk-api/handlers"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
)

func main() {

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

	// r.Route("/user", func(r chi.Router) {
	// 	r.Group(func(r chi.Router) {
	// 		r.Use(handlers.RestrictAuth)
	// 		r.Post("/", handlers.NewUser)
	// 		r.Route("/{user_id}", func(r chi.Router) {
	// 			r.Get("/", handlers.GetUser)
	// 			r.Put("/", handlers.UpdateUser)
	// 			r.Delete("/", handlers.DeleteUser)
	// 		})
	// 	})
	// })

	r.Route("/celebrity", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(handlers.RestrictAdmin)
			r.Get("/", handlers.ListCelebrity)
		})
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
		r.Get("/", handlers.GetMyUserData)
		r.Get("/claim", handlers.ListClaimForUserById)
	})

	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", handlers.Login)
		r.Post("/register", handlers.Register)
	})

	r.Route("/feed", func(r chi.Router) {
		r.Get("/", handlers.SearchArticles)
		// r.Get("/new", handlers.GetArticlesNewest)
		r.Get("/personalised", handlers.GetArticlesPersonalised)
	})

	// blockchain
	//crafting

	r.Route("/nft", func(r chi.Router) {
		r.Post("/identity", handlers.CraftIdentity)
		r.Post("/prediction", handlers.CraftPrediction)
	})

	err := http.ListenAndServe(":3333", r)
	if err != nil {
		log.Fatalf("Error serving HTTP handlers: %v", err)
	}

}
