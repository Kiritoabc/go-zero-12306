// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	regionStation "go-zero-12306/app/ticket-service/cmd/api/desc/internal/handler/regionStation"
	tempSeat "go-zero-12306/app/ticket-service/cmd/api/desc/internal/handler/tempSeat"
	ticket "go-zero-12306/app/ticket-service/cmd/api/desc/internal/handler/ticket"
	trainStation "go-zero-12306/app/ticket-service/cmd/api/desc/internal/handler/trainStation"
	"go-zero-12306/app/ticket-service/cmd/api/desc/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/region-station/query",
				Handler: regionStation.ListRegionStationHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/station/all",
				Handler: regionStation.ListAllStationHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/ticket-service"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/seat/reset",
				Handler: tempSeat.PurchaseTicketsHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/ticket-service/temp"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/query",
				Handler: ticket.PageListTicketQueryHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/purchase",
				Handler: ticket.PurchaseTicketsV1Handler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/purchase/v2",
				Handler: ticket.PurchaseTicketsV2Handler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/cancel",
				Handler: ticket.CancelTicketOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/pay/query",
				Handler: ticket.GetPayInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/refund",
				Handler: ticket.CommonTicketRefundHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/ticket-service/ticket"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/train-station/query",
				Handler: trainStation.ListTrainStationQueryHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/ticket-service"),
	)
}
