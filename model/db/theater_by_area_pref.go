package db

type TheaterByAreaPref []Theater

func (arr TheaterByAreaPref) Len() int {
	return len(arr)
}

func (arr TheaterByAreaPref) Less(i, j int) bool {

	flg := true
	if arr[i].AreaId > arr[j].AreaId {
		flg = false
	} else if arr[i].AreaId == arr[j].AreaId &&
		arr[i].PrefectureId > arr[j].PrefectureId {
		flg = false
	} else if arr[i].AreaId == arr[j].AreaId &&
		arr[i].PrefectureId == arr[j].PrefectureId &&
		arr[i].Id > arr[j].Id {
		flg = false
	}
	/*if arr[i].AreaId < arr[j].AreaId {
		flg = false
	} else if arr[i].AreaId == arr[j].AreaId &&
		arr[i].PrefectureId < arr[j].PrefectureId {
		flg = false
	} else if arr[i].AreaId == arr[j].AreaId &&
		arr[i].PrefectureId == arr[j].PrefectureId &&
		arr[i].Id < arr[j].Id {
		flg = false
	}*/
	return flg
}

func (arr TheaterByAreaPref) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
