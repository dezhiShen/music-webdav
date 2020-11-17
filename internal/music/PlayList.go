package music

// Music 音乐信息
type Music struct {
	Id   string `json:id`
	Name string `json:name`
}

// PlayList 播放列表
type PlayList struct {
	Id    string   `json:id`
	Name  string   `json:name`
	Music []*Music `json:music`
}

// LoadAllPlayList 加载全部的播放列表
func LoadAllPlayListNoSync() ([]*PlayList, error) {
	return LoadAllPlayList(false)
}

func LoadAllPlayList(sync bool) ([]*PlayList, error) {
	return nil, nil
}

// LoadPlayList 加载单个播放列表
func LoadPlayList(id string) (*PlayList, error) {
	return nil, nil
}
