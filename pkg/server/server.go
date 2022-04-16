package server

import (
	"github.com/gin-gonic/gin"
	"kubook/pkg/cmdbsvc"
	"time"
)

type Config struct {
	CmdbHost string
	CmdbUsername string
	CmdbPassword string
}

type Server struct {
	cmdbtoken *cmdbsvc.CmdbClient
	processTime time.Time
}

type Result struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type GinGesult struct {
	Services []*cmdbsvc.Service  `json:"services"`
	RunningTime string           `json:"running_time"`
	Total int				     `json:"total"`
}

func New(conf Config) *Server{
	gin.SetMode(gin.DebugMode)
	return &Server{
		cmdbtoken: cmdbsvc.New(cmdbsvc.Config{
			CmdbHost: conf.CmdbHost,
			CmdbUsername: conf.CmdbUsername,
			CmdbPassword: conf.CmdbPassword,
		}),
	}
}

func (s *Server)S1() (services []*cmdbsvc.Service,err error) {
	bt := s.cmdbtoken.GetServices()
	for _,d := range bt {
		s := &cmdbsvc.Service{
			Service_name: d.Workload_name,
			Workload_type: d.Workload_type,
			Git_url: d.Service_giturl,
			Service_center: d.Service_project,
			Status: d.Metadata.Status,
		}
		services = append(services,s)
	}
	return
}

func (s *Server)S2()(ret *cmdbsvc.SvcResult,err error){

	svc ,err := s.S1()
	return &cmdbsvc.SvcResult{
		Services:    svc,
		RunningTime: s.processTime.Format(time.RFC3339),
	}, nil
}


func (s *Server) Run() error {
	r := gin.Default()
	r.GET("/health", func(context *gin.Context) {
		context.String(200,"%s","ok")
	})
	r.GET("/api/v1/service", func(c *gin.Context) {
		result , err := s.S2()
		if err != nil {
			c.JSON(500, Result{
				Status: 500,
				Msg:    err.Error(),
				Data:   nil,
			})
		}
		c.JSON(200,Result{
			Status: 0,
			Msg: "",
			Data: GinGesult{
				Services: result.Services,
				RunningTime: result.RunningTime,
				Total: len(result.Services),
			},
		})
	})
	return r.Run()
}




//func (s *Server) Run() err {
//	r := gin.Default()
//	r.GET("/health", s.Health)
//	r.GET("/api/v1/info", s.Calculate)
//
//	return r.Run()
//
//
//}