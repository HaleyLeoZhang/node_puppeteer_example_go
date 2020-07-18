package comic

func GetComicList(pageNum int, pageSize int, maps interface{}) ([]*Comics, error) {
	var ComicList []*Comics

	err := db.Where(maps).Offset(pageNum).Order("weight desc").Limit(pageSize).Find(&ComicList).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return ComicList, nil
}

func GetComicInfo(Channel int, SourceID int, maps interface{}) (*Comics, error) {
	var ComicInfo Comics

	query_maps := make(map[string]interface{})
	query_maps["channel"] = Channel
	query_maps["source_id"] = SourceID

	err := db.Where(query_maps).Where(maps).First(&ComicInfo).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &ComicInfo, nil
}

func GetComicTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Comics{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
