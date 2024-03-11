package files

import (
	"github.com/go-chi/chi/v5"
)

func GetFileRouter(h *Handler) *chi.Mux {
	filesRouter := chi.NewRouter()

	filesRouter.Post("/upload", h.UploadHandler)
	filesRouter.Delete("/delete/{id}", h.DeleteHandler)
	filesRouter.Post("/batch-delete", h.MultipleDeleteHandler)
	filesRouter.Get("/download/{id}", h.DownloadHandler)
	filesRouter.Get("/overview/{id}", h.OverviewHandler)

	filesRouter.Post("/read-articles-price", h.ReadArticlesPriceHandler)
	filesRouter.Post("/read-articles", h.ReadArticlesHandler)
	filesRouter.Post("/read-articles-inventory", h.ReadArticlesInventoryHandler)
	filesRouter.Post("/read-articles-donation", h.ReadArticlesDonationHandler)
	filesRouter.Post("/read-articles-simple-procurement", h.ReadArticlesSimpleProcurementHandler)
	filesRouter.Post("/read-expire-inventories", h.ReadExpireInventoriesHandler)
	filesRouter.Post("/read-expire-imovable-inventories", h.ReadExpireImovableInventoriesHandler)
	filesRouter.Post("/import-inventories-excel", h.ImportExcelOrgUnitInventoriesHandler)
	//filesRouter.Post("/import-ps1-excel", h.ImportExcelPS1)
	filesRouter.Post("/import-vacations-excel", h.ImportUserProfileVacationsHandler)
	filesRouter.Post("/import-expirience-excel", h.ImportUserExpirienceHandler)

	return filesRouter
}
