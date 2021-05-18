package models

//Metadata Info model to store the data in db
type MetadataInfo struct {
	ID                int    `gorm:"column:id;primary_key;auto_increment"`
	NeSwId            string `gorm:"column:ne_sw_id"`
	FragmentId        string `gorm:"column:fragment_id"`
	AdaptationId      string `gorm:"column:adaptation_id"`
	AdaptationVersion string `gorm:"column:adaptation_version"`
	PackageVersion    string `gorm:"column:package_version"`
	FragmentType      string `gorm:"column:fragment_type"`
	Tag               string `gorm:"column:tag"`
	// ArtifactBuild     string `gorm:"column:artifact_build"`
	// Source            string `gorm:"column:source"`
	Sha1 string `gorm:"column:sha1"`
}

//TableName tableName
func (MetadataInfo) TableName() string {
	return "metadata_info"
}
