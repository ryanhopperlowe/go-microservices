package handlers

import (
	"context"
	"net/http"
	"product-api/data"
)

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := data.FromJSON(prod, r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)

			w.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, w)
			return
		}

		p.l.Println("[DEBUG] validating product in middleware")
		errs := p.v.Validate(prod)
		p.l.Println("[DEBUG] validation successful")

		if len(errs) != 0 {
			p.l.Println("[ERROR] validating product", errs)

			w.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, w)
			return

		}

		p.l.Println("[DEBUG] setting context")
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
