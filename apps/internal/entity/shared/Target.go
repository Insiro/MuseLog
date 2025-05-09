package shared

type TargetType string

const (
	TargetArtist TargetType = "ARTIST"
	TargetAlbum  TargetType = "ALBUM"
)

type TargetRelation struct {
	Target   *TargetType `gorm:"enum('ARTIST', 'ALBUM');default:null;composite:idx_target,unique" json:"target"`
	TargetId int         `gorm:"composite:idx_target,unique" json:"targetId"`
}
