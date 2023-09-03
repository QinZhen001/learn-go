package register

import (
	"fmt"
	"learngo/internal"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

type IRegister interface {
	Register(name, id string, port int, tags []string) error
	DeRegister(serviceId string) error
}

type ConsulRegistry struct {
	Host string
	Port int
}

func NewConsulRegistry(host string, port int) ConsulRegistry {
	return ConsulRegistry{
		Host: host,
		Port: port,
	}
}

func (c ConsulRegistry) Register(name, id string, port int, tags []string) error {
	defaultConfig := api.DefaultConfig()
	h := internal.AppConf.ConsulConfig.Host
	p := internal.AppConf.ConsulConfig.Port
	defaultConfig.Address = fmt.Sprintf("%s:%d", h, p)
	client, err := api.NewClient(defaultConfig)
	if err != nil {
		zap.S().Error("new consul client failed", "msg", err.Error())
		return err
	}
	agentServiceReg := new(api.AgentServiceRegistration)
	agentServiceReg.Address = defaultConfig.Address
	agentServiceReg.ID = id
	agentServiceReg.Name = name
	agentServiceReg.Port = port
	agentServiceReg.Tags = tags
	serverAdder := fmt.Sprintf("http://%s:%d/health", internal.AppConf.ProductWebConfig.Host,
		internal.AppConf.ProductWebConfig.Port)
	check := api.AgentServiceCheck{
		HTTP:                           serverAdder,
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "30s",
	}
	agentServiceReg.Check = &check
	err = client.Agent().ServiceRegister(agentServiceReg)
	if err != nil {
		zap.S().Error("register server error", "msg", err.Error())
		return err
	}
	return nil
}

func (c ConsulRegistry) DeRegister(serviceId string) error {
	defaultConfig := api.DefaultConfig()
	h := internal.AppConf.ConsulConfig.Host
	p := internal.AppConf.ConsulConfig.Port
	defaultConfig.Address = fmt.Sprintf("%s:%d", h, p)
	client, err := api.NewClient(defaultConfig)
	if err != nil {
		zap.S().Error("new consul client failed", "msg", err.Error())
		return err
	}
	return client.Agent().ServiceDeregister(serviceId)
}
