package simplejson

import (
	"testing"
)

func BenchmarkDemo1(b *testing.B) {
	b.ReportAllocs()

	for i:=0; i<10000; i++ {
	js, err := NewJson([]byte(`{
		"version": 3,
		"reqid": "6834dec7540fac9765a83de45828dd26",
		"seatIds": "100097",
		"requestType": "biddingad",
		"info": {
			"app": {
				"id": "5",
				"name": "7日天气预报-安卓辅助版2",
				"versionName": "9.0.0",
				"cat": [
					6501
				],
				"package_name": "com.tianqi2345",
				"ver": "90000",
				"keywords": "",
				"cid": "3",
				"appchannel": "cu1000490"
			},
			"device": {
				"hmsVersion": "",
				"asVersion": "86036",
				"ua": "Mozilla/5.0 (Linux; Android 10; PCAM10 Build/QP1A.190711.020; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/77.0.3865.92 Mobile Safari/537.36",
				"device_type": 1,
				"brand": "OPPO",
				"model": "PCAM10",
				"density": 480,
				"romOsVersion": "7.1",
				"oem": "OPPO",
				"os": 4,
				"osv": "10",
				"ip": "172.17.0.67",
				"ipv6": "",
				"carrier": 1,
				"network": 0,
				"width": 1080,
				"height": 2264,
				"orientation": 1,
				"impWidth": 1080,
				"impHeight": 2264,
				"imei": "869157023999228",
				"oaid": "FB4AECA3F7C44925901ABA140B0B4B8E441349265eeb4c098f70fd83f64e9f6e",
				"imei_md5": "",
				"dpid": "94766f392b637333",
				"dpid_md5": "8f703fb4a57516443a5fa780626d41e7",
				"idfa": "",
				"idfv": "",
				"idfa_md5": "",
				"mac": "70:dd:a8:28:12:6e",
				"mac_md5": "9de4c564b116e5b4c891875e4409380d",
				"geo": {
					"lat": 0,
					"lon": 0,
					"type": 2,
					"country": "",
					"province": "",
					"city": ""
				},
				"boot_mark": "8b50a2e8-9029-4647-83ef-a566a8cda3b6",
				"update_mark": "1568314639.113693966"
			},
			"user": {
				"id": "70dda828126e",
				"uid": "70dda828126e",
				"wuid": "A00000A1F17A14",
				"yob": "",
				"gender": "",
				"keywords": ""
			},
			"clientTime": 1632881674000,
			"version": {
				"apiversion": 0,
				"sdkjarversion": 100205
			},
			"isNewuser": 1,
			"clicksourceChannel": "",
			"thirdExt": {
				"buyerId111": "W8IBAqwCZzvPqU2HdvfTdbOTmlcJd4u0-q8HzhJ3UDmzSgvN3mTqCF8mGu7I6D_6XzYkvCCCcRAgt7FUrM67825O3VEZV3VWI4HOjEvGsobfhjTIwxs_uE4Q7ajcdewsGVX3oHqwMQQWpSAqSKdzdYIQyiVlqtAJXUv2JCBSE7mEfwRV",
				"wxInstalled": true,
				"opensdkVer": "5.5.8",
				"slotid": "1012421121533571",
				"sourceId": 10002,
				"yybId": "1106076180",
				"csjConfig": {
					"sourceId": 10019,
					"slotid": "887578722",
					"token": "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\n0tU=\n"
				}
			}
		}
	}`))

	if err != nil {
		b.Fatal(err)
	}
	if js != nil {

	}
}
}



