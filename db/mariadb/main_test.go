package main

import (
	"fmt"
	"learngo/db/mariadb/models"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	// "gorm.io/gorm" // new version
	"github.com/jinzhu/gorm" // deprecated version
)

func getDb() *gorm.DB {
	dsn := "root:arthur@tcp(127.0.0.1:13306)/metadata_inventory?charset=utf8mb4&parseTime=True&loc=Local"
	// session, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	driverName := fmt.Sprintf("mysql-dyn-auth-%d", time.Now().UnixNano())
	d, ok := gorm.GetDialect("mysql")
	if !ok {
		fmt.Println("Failed to get Gorm dialect \"mysql\"")
	}

	gorm.RegisterDialect(driverName, d)
	session, err := gorm.Open(driverName, dsn, "false")
	if err != nil {
		fmt.Printf("fail to open db. %v", err)
	}
	return session
}

func TestBatchInset(t *testing.T) {
	session := getDb()
	fmt.Println("Batch Inset")

	var tag string
	metaList := []models.MetadataInfo{
		{NeSwId: "5G19A_1", FragmentId: "FM", AdaptationId: "com.compname.nrbts", AdaptationVersion: "5G19A", Sha1: "123"},
		{NeSwId: "5G19A_1", FragmentId: "FM", AdaptationId: "com.compname.nrbts", AdaptationVersion: "5G19A", Sha1: "124", Tag: tag},
	}
	tx := session.Create(&metaList)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}
	fmt.Println(tx.RowsAffected)
}

func TestArtifacts(t *testing.T) {
	session := getDb()
	fmt.Println("complex search")

	var infos []models.MetadataInfo
	session.Where("ne_sw_id IN ?", []string{"SBTS20A_ENB_0000_000626_000000", "ISBC-20.2"}).Find(&infos)
	for _, i := range infos {
		fmt.Println(i)
	}
}

func TestDelete(t *testing.T) {
	session := getDb()
	fmt.Println("testDelete")

	var result []models.MetadataInfo

	// AND package_version IS NULL
	// tx := session.Where("fragment_id IN ?", []string{"CUDO"}).Find(&result)
	// tx := session.Where("fragment_id IN ?", []string{"5G19A_1", "CUDO"}).Find(&result)

	q := models.MetadataInfo{NeSwId: "ISBC-20.2"}
	tx := session.Where(&q).Find(&result)

	if tx.Error != nil {
		fmt.Println(tx.Error)
	}
	fmt.Println(tx.RowsAffected)
	for _, v := range result {
		fmt.Println(v)
	}

	// tx = session.Where("ne_sw_id IN ? AND package_version IS NULL", []string{"5G19A_1"}).Delete(&models.MetadataInfo{})
	// if tx.Error != nil {
	// 	fmt.Println(tx.Error)
	// }
	// fmt.Println(tx.RowsAffected)
}

func TestMetadataInfo(t *testing.T) {
	session := getDb()
	var infos []*models.MetadataInfo
	// r := session.Find(&infos)
	r := session.Model(models.MetadataInfo{}).Group("fragment_id").Find(&infos)
	fmt.Println(r.RowsAffected)
	for _, info := range infos {
		fmt.Println(*info)
		if strings.HasPrefix(info.AdaptationVersion, "XX") {
			info.AdaptationVersion = info.AdaptationVersion[2:]
		} else {
			info.AdaptationVersion = fmt.Sprintf("XX%s", info.AdaptationVersion)
		}
		// session.Save(info)
	}

	a := models.MetadataInfo{
		// Id:                250,
		NeSwId:            "SBTS20A_ENB_0000_000626_000000",
		FragmentId:        "FM",
		AdaptationId:      "com.compname.eqm.hw",
		AdaptationVersion: "EQMHW19A_1904_004",
	}
	// fmt.Println(a)
	sr := session.Create(&a)
	fmt.Println(sr.RowsAffected)
	if sr.Error != nil {
		fmt.Println(sr.Error)
	}

	fmt.Println("end")
}

func TestMappings(t *testing.T) {
	session := getDb()
	var a []*ApiMapping
	q := &Mapping{IntegratedType: "DA"}
	// session.Find(&a, q)
	session.Model(q).Where(q).Find(&a)
	fmt.Println(len(a))
	for i, v := range a {
		fmt.Printf("#%d [%s] [%s]\n", i, *v.NeSwID, v.IntegratedID)
	}
}

// metadata_info
// ArtifactNew
type Mapping struct {

	// 	ne_sw_id	fragment_id	adaptation_id	adaptation_version	package_id	package_version	fragment_type
	//	tag	artifact_build	source	sha1
	NeSwID         string // `gorm:"column:ne_sw_id"`
	IntegratedID   string // column name is `integrated_id`
	IntegratedType string
	// PackageID         string
	// PackageVersion    string
	// FragmentType      string
	// Tag               string
	// ArtifactBuild     string
	// Source            string
	// Sha1              string
}

type ApiMapping struct {

	// 	ne_sw_id	fragment_id	adaptation_id	adaptation_version	package_id	package_version	fragment_type
	//	tag	artifact_build	source	sha1
	NeSwID       *string // `gorm:"column:ne_sw_id"`
	IntegratedID string  // column name is `integrated_id`
}

/*
s3_store_client_test.go

func TestReindexOnline(t *testing.T) {

	serverURL := "localhost:9000"
	mc, err := miniowrapper.NewMinioWrapper(&miniowrapper.Config{Endpoint: serverURL, AccessKey: "minio", SecretKey: "miniostorage"})
	if err != nil {
		t.Fail()
	}

	storage, _ := NewFileStoreClient(mc, serverURL)

	// info, version, err := storage.Reindex("SBTS20A_ENB_0000_000626_000000")
	// assert.Equal(t, len(info), 0)
	// assert.Equal(t, version, "")
	// assert.Equal(t, err.Error(), "test package is empty")

	meta := storage.GetMetadata("5G_CLA_BTS_1234.100", "20200613T122748")
	fmt.Println(len(meta))
	fmt.Println(meta[0].Source)

}

*/

/*

var logger = logging.NewLogger()
var log = tracing.NewLogger(logger)

func TestOnline(t *testing.T) {

	session, err := mariadb.NewDBClientFromENV(log)
	if err != nil {
		fmt.Printf("fail to open db. %v", err)
		return
	}
	// session.Where("sha1 IS NULL AND package_version IS NULL").Delete(models.MetadataInfo{})
	var infos []models.MetadataInfo
	// mi := models.MetadataInfo{NeSwId: "5G19A_1"}
	// session.Where(&mi).Find(&infos)
	session.Where("fragment_id IN ?", []string{"5G19A_1", "FM"}).Find(&infos)
	// session.Where("fragment_id = ? AND id > ?", "FM", 1265).Find(&infos)
	fmt.Println("len(infos)", len(infos))
	// for _, i := range infos {
	// 	fmt.Println(i.NeSwId, i.FragmentId)
	// }
	// r := session.Where("artifact_build IS NULL AND source IS NULL").Delete(models.MetadataInfo{}) // works
	// fmt.Println(r.RowsAffected)

	t.Skip("1")
	var metaList []models.MetadataInfo

	// {NeSwId: "5G19A_1", FragmentId: "FM", AdaptationId: "com.compname.nrbts", AdaptationVersion: "5G19A", Sha1: "123"},
	// {NeSwId: "5G19A_1", FragmentId: "FM", AdaptationId: "com.compname.nrbts", AdaptationVersion: "5G19A", Sha1: "124"},

	meta := models.MetadataInfo{
		NeSwId:            "SBTS20A_ENB_0000_000626_000000",
		FragmentId:        "NEAC",
		AdaptationId:      "com.compname.srbts",
		AdaptationVersion: "SBTS20A",
		PackageVersion:    "20200506T152539",
		FragmentType:      "NADA",
		Tag:               "",
		ArtifactBuild:     "20200305T130258",
		Source:            "files/NEAC/com.compname.srbts.nada/SBTS20A/o2ml/20200305.130258/adaptation_com.compname.srbts.nada-SBTS20A-20200305T130258.zip",
		Sha1:              "d4ea84687afa4b27396493aa245cd8aa288aac74"}
	// result := session.Create(&meta)
	metaList = append(metaList, meta)
	fmt.Println("len(metaList)", len(metaList))
	result := session.Create(&metaList)

	fmt.Println(result.RowsAffected)

	t.Skip("x")

	mc, err := miniowrapper.NewMinioWrapper(&miniowrapper.Config{Endpoint: "localhost:9000", AccessKey: "minio", SecretKey: "miniostorage"})
	if err != nil {
		t.Fail()
	}
	fileStoreClient, _ := storage.NewFileStoreClient(mc, "")
	mdiStore, _ := storage.NewMdiStore("mariadb", 5, 3, log)
	var producer notifications.NotificationProducer

	h := NewPostV1JobsSyncPackages(log, producer, fileStoreClient, mdiStore)
	fmt.Printf("(%v, %T)\n", h, h)

	// syncPackages(&h, false)

}

*/
