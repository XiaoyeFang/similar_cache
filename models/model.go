package models

import (
	"similar_cache/config"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	APKDETAIL  = "apk"
	APKSIMILAR = "apk_similar"
)

type Similar struct {
	Id           bson.ObjectId `json:"_id" bson:"_id"`
	PackageName  string        `json:"package_name" bson:"package_name"`
	UpdateDate   time.Time     `json:"update_date" bson:"update_date"`
	CreateDate   time.Time     `json:"create_date" bson:"create_date"`
	Result       []string      `json:"result" bson:"result"`
	EditorResult []string      `json:"editor_result" bson:"editor_result"`
}

type Developer struct {
	Id          bson.ObjectId `json:"_id" bson:"_id"`
	IconFid     string        `json:"icon_fid" bson:"icon_fid"`
	BannerFid   string        `json:"banner_fid" bson:"banner_fid"`
	DeveloperId string        `json:"developer_id" bson:"developer_id"`
	Name        string        `json:"name" bson:"name"`
	CreateDate  time.Time     `json:"create_date" bson:"create_date"`
	TopIndex    int32         `json:"top_index" bson:"top_index"`
	ReviewStars float64       `json:"review_stars" bson:"review_stars"`
}

type SimDevResult struct {
	Similar   ApkResult `json:"similar"`
	Developer ApkResult `json:"developer"`
}

type ApkResult struct {
	List  []ApkDetail `json:"list"`
	Total int         `json:"total"`
}

type ApkDetail struct {
	Id                    bson.ObjectId `json:"_id" bson:"_id"`
	Name                  string        `json:"name" bson:"name"`
	LanguageName          interface{}   `json:"language_name" bson:"language_name"`
	ApkName               string        `json:"apk_name" bson:"apk_name"`
	PackageName           string        `json:"package_name" bson:"package_name"`
	Price                 string        `json:"price" bson:"price"`
	ApkType               int32         `json:"apk_type" bson:"apk_type"`
	Category              string        `json:"category" bson:"category"`
	Icon                  string        `json:"icon" bson:"icon"`
	IconFid               string        `json:"icon_fid" bson:"icon_fid"`
	ReviewStars           float32       `json:"review_stars" bson:"review_stars"`
	ReviewCount           int32         `json:"review_count" bson:"review_count"`
	UpdateDate            time.Time     `json:"update_date" bson:"update_date"`
	VersionCurrentVersion string        `json:"version_current_version" bson:"version_current_version"`
	Developer             string        `json:"developer" bson:"developer"`
	//DeveloperUserId         int32              `json:"developer_user_id" bson:"developer_user_id"`
	DeveloperLink   []interface{} `json:"developer_link" bson:"developer_link"`
	Updated         time.Time     `json:"updated" bson:"updated"`
	InstallTotal    int32         `json:"install_total" bson:"install_total"`
	CustomUrl       string        `json:"custom_url" bson:"custom_url"`
	CustomUrlOnly   bool          `json:"custom_url_only" bson:"custom_url_only"`
	CustomAttr      interface{}   `json:"custom_attr" bson:"custom_attr"`
	WishList        string        `json:"wish_list" bson:"wish_list"`
	PreRegister     string        `json:"pre_register" bson:"pre_register"`
	PreRegisterInfo interface{}   `json:"pre_register_info" bson:"pre_register_info"`
	//Published               bool               `json:"published" bson:"published"`
	IsShowStructured    bool        `json:"is_show_structured" bson:"is_show_structured"`
	Disable             []string    `json:"disable" bson:"disable"`
	DisableCountry      []string    `json:"disable_country" bson:"disable_country"`
	VersionXapk         bool        `json:"version_xapk" bson:"version_xapk"`
	EnableGoogleAdsense bool        `json:"enable_google_adsense" bson:"enable_google_adsense"`
	Banner              string      `json:"banner" bson:"banner"`
	BannerFid           string      `json:"banner_fid" bson:"banner_fid"`
	Keyword             []string    `json:"keyword" bson:"keyword"`
	Tags                []string    `json:"tags" bson:"tags"`
	Enable              bool        `json:"enable" bson:"enable"`
	Active              bool        `json:"active" bson:"active"`
	ActiveDate          time.Time   `json:"active_date" bson:"active_date"`
	Completed           bool        `json:"completed" bson:"completed"`
	ContentRating       []string    `json:"content_rating" bson:"content_rating"`
	InAppProducts       string      `json:"in_app_products" bson:"in_app_products"`
	ContentSourceType   int32       `json:"content_source_type" bson:"content_source_type"`
	SourceName          string      `json:"source_name" bson:"source_name"`
	InstallCount        string      `json:"install_count" bson:"install_count"`
	CreateDate          time.Time   `json:"create_date" bson:"create_date"`
	DescriptionShort    string      `json:"description_short" bson:"description_short"`
	LanguageComment     interface{} `json:"language_comment" bson:"language_comment"`
	FollowTotal         int32       `json:"follow_total" bson:"follow_total"`
	CommentScore1       int32       `json:"comment_score_1" bson:"comment_score_1"`
	CommentScore2       int32       `json:"comment_score_2" bson:"comment_score_2"`
	CommentScore3       int32       `json:"comment_score_3" bson:"comment_score_3"`
	CommentScore4       int32       `json:"comment_score_4" bson:"comment_score_4"`
	CommentScore5       int32       `json:"comment_score_5" bson:"comment_score_5"`
	CommentTotal        int32       `json:"comment_total" bson:"comment_total"`
	CommentScore        int32       `json:"comment_score" bson:"comment_score"`
	CommentScoreStars   float32     `json:"comment_score_stars" bson:"comment_score_stars"`
	CommentScoreTotal   int32       `json:"comment_score_total" bson:"comment_score_total"`
	IsShowBeta          bool        `json:"is_show_beta" bson:"is_show_beta"`
	UnCrawler           bool        `json:"un_crawler" bson:"un_crawler"`
	//Language []string             `json:"language" bson:"language"`
	LanguageApkName         interface{}   `json:"language_apk_name" bson:"language_apk_name"`
	LanguageRecommend       interface{}   `json:"language_recommend" bson:"language_recommend"`
	UrlName                 string        `json:"url_name" bson:"url_name"`
	LanguageTag             interface{}   `json:"language_tag" bson:"language_tag"`
	Version                 []interface{} `json:"version" bson:"version"`
	FastDownloadCountry     []string      `json:"fast_download_country" bson:"fast_download_country"`
	FastDownloadType        string        `json:"fast_download_type" bson:"fast_download_type"`
	FeatureDescriptionShort string        `json:"feature_description_short" bson:"feature_description_short"`
	FeatureBannerFid        string        `json:"feature_banner_fid" bson:"feature_banner_fid"`
}

type Tags []*Tag

type Tag struct {
	Id        int32  `json:"id" bson:"id"`
	Count     int32  `json:"count" bson:"count"`
	Name      string `json:"name" bson:"name"`
	Effective bool   `json:"effective" bson:"effective"`
}

type Comment struct {
	CommentScore1     int32 `json:"comment_score_1" bson:"comment_score_1"`
	CommentScore2     int32 `json:"comment_score_2" bson:"comment_score_2"`
	CommentScore3     int32 `json:"comment_score_3" bson:"comment_score_3"`
	CommentScore4     int32 `json:"comment_score_4" bson:"comment_score_4"`
	CommentScore5     int32 `json:"comment_score_5" bson:"comment_score_5"`
	CommentTotal      int32 `json:"comment_total" bson:"comment_total"`
	CommentScore      int32 `json:"comment_score" bson:"comment_score"`
	CommentScoreStars int32 `json:"comment_score_stars" bson:"comment_score_stars"`
	CommentScoreTotal int32 `json:"comment_score_total" bson:"comment_score_total"`
}

func CreatDatabase() (db *mgo.Database, err error) {

	session, err := mgo.Dial(config.CacheConfig.MongoDBUrl)
	if err != nil {

		panic(err)
	}
	db = session.DB("")

	return db, err
}

func ConnRedis() (conn redis.Conn, err error) {
	//fmt.Println(config.CacheConfig.RedisDB.Url)
	c, err := redis.Dial("tcp", config.CacheConfig.RedisDB.Url)
	if err != nil {
		panic(err)
	}
	return c, nil
}
