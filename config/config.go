package config

import (
	"encoding/json"
	"errors"
	"github.com/garyburd/redigo/redis"
	"github.com/golang/glog"
	"gopkg.in/mgo.v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

var CacheConfig *Config

const (
	HTTP_USER_AGENT  = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.90 Safari/537.36"
	HTTP_ACCEPT      = "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"
	MinPoolSize      = 5
	MaxPoolSize      = 50
	AcquireIncrement = 5
	SIMILARTIMEOUT   = 240
)

type Config struct {
	GrpcListen      string `yaml:"grpc_listen" json:"grpc_listen"`
	HTTPPort        string `yaml:"http_port" json:"http_port"`
	SyncHttpReq     string `yaml:"sync_http_req" json:"sync_http_req"`
	AsynHttpReq     string `yaml:"asyn_http_req" json:"asyn_http_req"`
	LogLevel        int    `yaml:"log_level" json:"log_level"`
	DeveloperPrefix string `yaml:"developer_prefix" json:"developer_prefix"`
	MongoDBUrl      string `yaml:"mongodb_url" json:"mongodb_url"`
	RedisDB         `yaml:"redis" json:"redis"`
}

type RedisDB struct {
	Url        string `yaml:"url" json:"url"`
	Expiration int64  `yaml:"expiration" json:"expiration"`
}

func init() {
	var err error
	if CacheConfig == nil {
		CacheConfig, err = LoadConf("./conf/app.yml")
		if err != nil {
			glog.V(0).Infoln("LoadConf", err)
		}
	}
}

func CreatDatabase() (db *mgo.Database, err error) {

	session, err := mgo.Dial(CacheConfig.MongoDBUrl)
	if err != nil {

		panic(err)
	}
	db = session.DB("")

	return db, err
}

func ConnRedis() (conn redis.Conn, err error) {
	//fmt.Println(config.CacheConfig.RedisDB.Url)
	c, err := redis.Dial("tcp", CacheConfig.RedisDB.Url)
	if err != nil {
		panic(err)
	}
	return c, nil
}

func LoadConf(filepath string) (*Config, error) {
	if filepath == "" {
		return nil, errors.New("filepath is empty, must use --config xxx.yml/json")
	}

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if strings.HasSuffix(filepath, ".json") {
		err = json.Unmarshal(data, &cfg)
	} else if strings.HasSuffix(filepath, ".yml") || strings.HasSuffix(filepath, ".yaml") {
		err = yaml.Unmarshal(data, &cfg)
	} else {
		return nil, errors.New("you config file must be json/yml")
	}

	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
