package worker

import (
	"github.com/Deansquirrel/goToolCommon"
	"github.com/Deansquirrel/goZ5ReportMdBaoZhShouR/object"
	"time"
)

type DummyData struct {
}

/**
虚拟数据替换
*/
func (dummyData *DummyData) DummyDataReplace(
	mdId int,
	begDate time.Time,
	endDate time.Time,
	list map[string]*object.MdBaoZhShouRData) {
	//新顺店数据替换
	if mdId == 10194 {
		if (goToolCommon.GetDateStr(begDate) >= "2019-09-03" && goToolCommon.GetDateStr(begDate) <= "2019-09-30") ||
			(goToolCommon.GetDateStr(endDate) >= "2019-09-03" && goToolCommon.GetDateStr(endDate) <= "2019-09-30") ||
			(goToolCommon.GetDateStr(begDate) < "2019-09-03" && goToolCommon.GetDateStr(endDate) > "2019-09-30") {
			dummyData.xs201909Replace(list)
		}
	}
}

/*
获取兴顺店（门店ID 10194）待替换数据（2019-09-03至 2019-09-30）
*/
func (dummyData *DummyData) getXs201909DummyData() map[string]*object.MdBaoZhShouRData {
	vData := make(map[string]*object.MdBaoZhShouRData)

	vZzData20190904 := make(map[string]float64)
	vZzData20190905 := make(map[string]float64)
	vZzData20190906 := make(map[string]float64)
	vZzData20190907 := make(map[string]float64)
	vZzData20190908 := make(map[string]float64)
	vZzData20190909 := make(map[string]float64)
	vZzData20190910 := make(map[string]float64)
	vZzData20190911 := make(map[string]float64)
	vZzData20190912 := make(map[string]float64)
	vZzData20190913 := make(map[string]float64)
	vZzData20190914 := make(map[string]float64)
	vZzData20190915 := make(map[string]float64)
	vZzData20190916 := make(map[string]float64)
	vZzData20190917 := make(map[string]float64)
	vZzData20190918 := make(map[string]float64)
	vZzData20190919 := make(map[string]float64)
	vZzData20190920 := make(map[string]float64)
	vZzData20190921 := make(map[string]float64)
	vZzData20190922 := make(map[string]float64)
	vZzData20190923 := make(map[string]float64)
	vZzData20190924 := make(map[string]float64)
	vZzData20190925 := make(map[string]float64)
	vZzData20190926 := make(map[string]float64)
	vZzData20190927 := make(map[string]float64)
	vZzData20190928 := make(map[string]float64)
	vZzData20190929 := make(map[string]float64)
	vZzData20190930 := make(map[string]float64)

	vZzData20190904["钛海"] = 162
	vZzData20190904["微信支付"] = 22633.9
	vZzData20190904["移动积分兑换"] = 0
	vZzData20190904["银联"] = 3738
	vZzData20190904["支付宝支付"] = 9297
	vZzData20190904["中欣银宝通"] = 84
	vZzData20190904["中影票务通"] = 174.38
	vZzData20190905["钛海"] = 621
	vZzData20190905["微信支付"] = 21651.6
	vZzData20190905["移动积分兑换"] = 0
	vZzData20190905["银联"] = 1046
	vZzData20190905["支付宝支付"] = 8095
	vZzData20190905["中欣银宝通"] = 89
	vZzData20190905["中影票务通"] = 149.52
	vZzData20190906["钛海"] = 0
	vZzData20190906["微信支付"] = 12397
	vZzData20190906["移动积分兑换"] = 9
	vZzData20190906["银联"] = 376.5
	vZzData20190906["支付宝支付"] = 7422
	vZzData20190906["中欣银宝通"] = 3
	vZzData20190906["中影票务通"] = 434.28
	vZzData20190907["钛海"] = 377.1
	vZzData20190907["微信支付"] = 14175.5
	vZzData20190907["移动积分兑换"] = 18
	vZzData20190907["银联"] = 1116
	vZzData20190907["支付宝支付"] = 7377.5
	vZzData20190907["中欣银宝通"] = 190
	vZzData20190907["中影票务通"] = 343.56
	vZzData20190908["钛海"] = 414
	vZzData20190908["微信支付"] = 13084.6
	vZzData20190908["移动积分兑换"] = 0
	vZzData20190908["银联"] = 1704
	vZzData20190908["支付宝支付"] = 5215.5
	vZzData20190908["中影票务通"] = 494.76
	vZzData20190909["钛海"] = 72
	vZzData20190909["微信支付"] = 9930
	vZzData20190909["移动积分兑换"] = 0
	vZzData20190909["银联"] = 53814.2
	vZzData20190909["支付宝支付"] = 3135
	vZzData20190909["中欣银宝通"] = 25
	vZzData20190909["中影票务通"] = 25.2
	vZzData20190910["钛海"] = 468
	vZzData20190910["微信支付"] = 10805.7
	vZzData20190910["移动积分兑换"] = 18
	vZzData20190910["银联"] = 781
	vZzData20190910["支付宝支付"] = 11718
	vZzData20190910["中欣银宝通"] = 325
	vZzData20190910["中影票务通"] = 58.8
	vZzData20190911["钛海"] = 396
	vZzData20190911["微信支付"] = 10999.7
	vZzData20190911["移动积分兑换"] = 0
	vZzData20190911["银联"] = 10094.1
	vZzData20190911["支付宝支付"] = 6641
	vZzData20190911["中欣银宝通"] = 49
	vZzData20190911["中影票务通"] = 75.6
	vZzData20190912["钛海"] = 503.1
	vZzData20190912["微信支付"] = 18678.5
	vZzData20190912["移动积分兑换"] = 27
	vZzData20190912["银联"] = 379
	vZzData20190912["支付宝支付"] = 7934.3
	vZzData20190912["中欣银宝通"] = 79
	vZzData20190912["中影票务通"] = 51.24
	vZzData20190913["钛海"] = 450
	vZzData20190913["微信支付"] = 7422.1
	vZzData20190913["移动积分兑换"] = 18
	vZzData20190913["银联"] = 1856
	vZzData20190913["支付宝支付"] = 5890.6
	vZzData20190913["中欣银宝通"] = 363
	vZzData20190913["中影票务通"] = 357
	vZzData20190914["钛海"] = 513
	vZzData20190914["微信支付"] = 11100.5
	vZzData20190914["移动积分兑换"] = 0
	vZzData20190914["银联"] = 1366
	vZzData20190914["支付宝支付"] = 4615.5
	vZzData20190914["中欣银宝通"] = 253
	vZzData20190914["中影票务通"] = 460.32
	vZzData20190915["钛海"] = 567
	vZzData20190915["微信支付"] = 8093
	vZzData20190915["移动积分兑换"] = 0
	vZzData20190915["银联"] = 419
	vZzData20190915["支付宝支付"] = 4543.5
	vZzData20190915["中影票务通"] = 517.44
	vZzData20190916["钛海"] = 369
	vZzData20190916["微信支付"] = 4397.1
	vZzData20190916["移动积分兑换"] = 0
	vZzData20190916["银联"] = 445
	vZzData20190916["支付宝支付"] = 3246
	vZzData20190916["中欣银宝通"] = 19
	vZzData20190916["中影票务通"] = 240.24
	vZzData20190917["钛海"] = 295.2
	vZzData20190917["微信支付"] = 5345.4
	vZzData20190917["移动积分兑换"] = 0
	vZzData20190917["银联"] = 421
	vZzData20190917["支付宝支付"] = 1953.5
	vZzData20190917["中欣银宝通"] = 44
	vZzData20190917["中影票务通"] = 285.6
	vZzData20190918["钛海"] = 432
	vZzData20190918["微信支付"] = 7724.1
	vZzData20190918["移动积分兑换"] = 9
	vZzData20190918["银联"] = 428
	vZzData20190918["支付宝支付"] = 3108
	vZzData20190918["中欣银宝通"] = 16
	vZzData20190918["中影票务通"] = 108.36
	vZzData20190919["钛海"] = 99
	vZzData20190919["微信支付"] = 6342.5
	vZzData20190919["移动积分兑换"] = 18
	vZzData20190919["银联"] = 201
	vZzData20190919["支付宝支付"] = 2049.5
	vZzData20190919["中欣银宝通"] = 171
	vZzData20190919["中影票务通"] = 599.76
	vZzData20190920["聚优福利"] = 231
	//========================================================
	//20191219【礼品卡】合并到【移动积分兑换】
	//vZzData20190920["礼品卡"] = 89
	vZzData20190920["移动积分兑换"] = 18 + 89
	//========================================================
	vZzData20190920["钛海"] = 99
	vZzData20190920["微信支付"] = 6750.3
	vZzData20190920["银联"] = 1791
	vZzData20190920["支付宝支付"] = 3040.3
	vZzData20190920["中欣银宝通"] = 86
	vZzData20190920["中影票务通"] = 47.04
	vZzData20190921["钛海"] = 504
	vZzData20190921["微信支付"] = 8914.7
	vZzData20190921["移动积分兑换"] = 9
	vZzData20190921["银联"] = 440
	vZzData20190921["支付宝支付"] = 4011.5
	vZzData20190921["中欣银宝通"] = 243
	vZzData20190921["中影票务通"] = 473.76
	vZzData20190922["钛海"] = 432
	vZzData20190922["微信支付"] = 10133.5
	vZzData20190922["移动积分兑换"] = 18
	vZzData20190922["银联"] = 1354
	vZzData20190922["支付宝支付"] = 4158.5
	vZzData20190922["中欣银宝通"] = 339
	vZzData20190922["中影票务通"] = 525.84
	vZzData20190923["钛海"] = 252
	vZzData20190923["微信支付"] = 5377.2
	vZzData20190923["移动积分兑换"] = 0
	vZzData20190923["银联"] = 818
	vZzData20190923["支付宝支付"] = 2037.6
	vZzData20190923["中欣银宝通"] = 58
	vZzData20190923["中影票务通"] = 161.28
	vZzData20190924["钛海"] = 666
	vZzData20190924["微信支付"] = 5980.5
	vZzData20190924["移动积分兑换"] = 18
	vZzData20190924["银联"] = 92
	vZzData20190924["支付宝支付"] = 3438.7
	vZzData20190924["中欣银宝通"] = 24
	vZzData20190924["中影票务通"] = 144.48
	vZzData20190925["钛海"] = 234
	vZzData20190925["微信支付"] = 7847
	vZzData20190925["移动积分兑换"] = 0
	vZzData20190925["支付宝支付"] = 2129.5
	vZzData20190925["中欣银宝通"] = 167.5
	vZzData20190925["中影票务通"] = 263.34
	vZzData20190926["钛海"] = 261
	vZzData20190926["微信支付"] = 5883.07
	vZzData20190926["移动积分兑换"] = 9
	vZzData20190926["银联"] = 494
	vZzData20190926["支付宝支付"] = 1551.5
	vZzData20190926["中欣银宝通"] = 139.43
	vZzData20190926["中影票务通"] = 179.76
	vZzData20190927["钛海"] = 261
	vZzData20190927["微信支付"] = 8655.5
	vZzData20190927["移动积分兑换"] = 0
	vZzData20190927["支付宝支付"] = 2844.5
	vZzData20190927["中欣银宝通"] = 78
	vZzData20190927["中影票务通"] = 168
	vZzData20190928["钛海"] = 610.2
	vZzData20190928["微信支付"] = 9063.5
	vZzData20190928["移动积分兑换"] = 0
	vZzData20190928["银联"] = 2599
	vZzData20190928["支付宝支付"] = 5423.2
	vZzData20190928["中欣银宝通"] = 43
	vZzData20190928["中影票务通"] = 1160.04
	vZzData20190929["聚优福利"] = 28
	vZzData20190929["钛海"] = 313.2
	vZzData20190929["微信支付"] = 6924
	vZzData20190929["移动积分兑换"] = 0
	vZzData20190929["银联"] = 660
	vZzData20190929["支付宝支付"] = 3258
	vZzData20190929["中欣银宝通"] = 207
	vZzData20190929["中影票务通"] = 593.04
	vZzData20190930["钛海"] = 296.1
	vZzData20190930["微信支付"] = 8419.3
	vZzData20190930["移动积分兑换"] = 45
	vZzData20190930["银联"] = 630
	vZzData20190930["支付宝支付"] = 3569
	vZzData20190930["中欣银宝通"] = 159
	vZzData20190930["中影票务通"] = 319.2

	vData["2019-09-04"] = &object.MdBaoZhShouRData{Yyr: "2019-09-04", TransferDetail: vZzData20190904, Cash: 5823.7}
	vData["2019-09-05"] = &object.MdBaoZhShouRData{Yyr: "2019-09-05", TransferDetail: vZzData20190905, Cash: 1130.5}
	vData["2019-09-06"] = &object.MdBaoZhShouRData{Yyr: "2019-09-06", TransferDetail: vZzData20190906, Cash: 4338}
	vData["2019-09-07"] = &object.MdBaoZhShouRData{Yyr: "2019-09-07", TransferDetail: vZzData20190907, Cash: 5635.5}
	vData["2019-09-08"] = &object.MdBaoZhShouRData{Yyr: "2019-09-08", TransferDetail: vZzData20190908, Cash: 2029.9}
	vData["2019-09-09"] = &object.MdBaoZhShouRData{Yyr: "2019-09-09", TransferDetail: vZzData20190909, Cash: 2251}
	vData["2019-09-10"] = &object.MdBaoZhShouRData{Yyr: "2019-09-10", TransferDetail: vZzData20190910, Cash: 4123}
	vData["2019-09-11"] = &object.MdBaoZhShouRData{Yyr: "2019-09-11", TransferDetail: vZzData20190911, Cash: 1548}
	vData["2019-09-12"] = &object.MdBaoZhShouRData{Yyr: "2019-09-12", TransferDetail: vZzData20190912, Cash: 2622.5}
	vData["2019-09-13"] = &object.MdBaoZhShouRData{Yyr: "2019-09-13", TransferDetail: vZzData20190913, Cash: 2434.4}
	vData["2019-09-14"] = &object.MdBaoZhShouRData{Yyr: "2019-09-14", TransferDetail: vZzData20190914, Cash: 2119}
	vData["2019-09-15"] = &object.MdBaoZhShouRData{Yyr: "2019-09-15", TransferDetail: vZzData20190915, Cash: 1808}
	vData["2019-09-16"] = &object.MdBaoZhShouRData{Yyr: "2019-09-16", TransferDetail: vZzData20190916, Cash: 766}
	vData["2019-09-17"] = &object.MdBaoZhShouRData{Yyr: "2019-09-17", TransferDetail: vZzData20190917, Cash: 1164}
	vData["2019-09-18"] = &object.MdBaoZhShouRData{Yyr: "2019-09-18", TransferDetail: vZzData20190918, Cash: 687.5}
	vData["2019-09-19"] = &object.MdBaoZhShouRData{Yyr: "2019-09-19", TransferDetail: vZzData20190919, Cash: 1584}
	vData["2019-09-20"] = &object.MdBaoZhShouRData{Yyr: "2019-09-20", TransferDetail: vZzData20190920, Cash: 787.1}
	vData["2019-09-21"] = &object.MdBaoZhShouRData{Yyr: "2019-09-21", TransferDetail: vZzData20190921, Cash: 3505.5}
	vData["2019-09-22"] = &object.MdBaoZhShouRData{Yyr: "2019-09-22", TransferDetail: vZzData20190922, Cash: 2449.5}
	vData["2019-09-23"] = &object.MdBaoZhShouRData{Yyr: "2019-09-23", TransferDetail: vZzData20190923, Cash: 1981.5}
	vData["2019-09-24"] = &object.MdBaoZhShouRData{Yyr: "2019-09-24", TransferDetail: vZzData20190924, Cash: 858}
	vData["2019-09-25"] = &object.MdBaoZhShouRData{Yyr: "2019-09-25", TransferDetail: vZzData20190925, Cash: 1151.5}
	vData["2019-09-26"] = &object.MdBaoZhShouRData{Yyr: "2019-09-26", TransferDetail: vZzData20190926, Cash: 1119}
	vData["2019-09-27"] = &object.MdBaoZhShouRData{Yyr: "2019-09-27", TransferDetail: vZzData20190927, Cash: 2317}
	vData["2019-09-28"] = &object.MdBaoZhShouRData{Yyr: "2019-09-28", TransferDetail: vZzData20190928, Cash: 1587}
	vData["2019-09-29"] = &object.MdBaoZhShouRData{Yyr: "2019-09-29", TransferDetail: vZzData20190929, Cash: 1129.5}
	vData["2019-09-30"] = &object.MdBaoZhShouRData{Yyr: "2019-09-30", TransferDetail: vZzData20190930, Cash: 1793}

	for _, v := range vData {
		total := 0.0
		for _, zV := range v.TransferDetail {
			total = total + zV
		}
		v.Transfer = total
		v.Total = v.Cash + v.Transfer
	}

	return vData
}

/*
兴顺店（门店ID 10194）数据替换
日期 2019-09-04 至 2019-09-30
*/
func (dummyData *DummyData) xs201909Replace(list map[string]*object.MdBaoZhShouRData) {
	totalCheck := 0
	for k, v := range dummyData.getXs201909DummyData() {
		if _, ok := list[k]; ok {
			totalCheck = list[k].TotalCheck
			list[k] = v
			list[k].TotalCheck = totalCheck
		}
	}
}
