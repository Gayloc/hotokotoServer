package hitokoto

type Hitokoto struct {
	Id           int    `json:"id"`
	Hitokoto     string `json:"hitokoto"`
	HitokotoType string `json:"hitokoto_type"`
	Reviewer     int    `json:"reviewer"`
	From_who     string `json:"from_who"`
	Length       int    `json:"length"`
}

type DataBase struct {
	Content Hitokoto   `json:"content"`
	Data    []Hitokoto `json:"data"`
}

func (d DataBase) FindByType(hitokotoType string) *DataBase {
	if hitokotoType == "" {
		return &d
	} else {
		result := &DataBase{}
		result.Content.HitokotoType = hitokotoType
		for i := 0; i < len(d.Data); i++ {
			if d.Data[i].HitokotoType == hitokotoType {
				result.Data = append(result.Data, d.Data[i])
			}
		}

		return result

	}
}

func (d *DataBase) AddItem(hitokoto Hitokoto) bool {
	if hitokoto.Hitokoto == "" || hitokoto.HitokotoType == "" || hitokoto.From_who == "" {
		return false
	}
	if hitokoto.Id == 0 {
		if len(d.Data) > 0 {
			hitokoto.Id = d.Data[len(d.Data)-1].Id + 1
		} else {
			hitokoto.Id = 1
		}
	}
	if hitokoto.Length == 0 {
		hitokoto.Length = len(hitokoto.Hitokoto)
	}
	d.Data = append(d.Data, hitokoto)
	return true
}

func (d *DataBase) DelItem(id int) bool {
	for i := 0; i < len(d.Data); i++ {
		if d.Data[i].Id == id {
			d.Data = append(d.Data[:i], d.Data[i+1:]...)
			return true
		}
	}
	return false
}

func (d *DataBase) EditItem(id int, hitokoto Hitokoto) bool {
	for i := 0; i < len(d.Data); i++ {
		if d.Data[i].Id == id {
			if hitokoto.Hitokoto != "" {
				d.Data[i].Hitokoto = hitokoto.Hitokoto
			}
			if hitokoto.HitokotoType != "" {
				d.Data[i].HitokotoType = hitokoto.HitokotoType
			}
			if hitokoto.Reviewer != 0 {
				d.Data[i].Reviewer = hitokoto.Reviewer
			}
			if hitokoto.From_who != "" {
				d.Data[i].From_who = hitokoto.From_who
			}
			if hitokoto.Length != 0 {
				d.Data[i].Length = hitokoto.Length
			}
			return true
		}
	}
	return false
}
