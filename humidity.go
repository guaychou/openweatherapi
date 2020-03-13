package openweatherapi

func humidityRange(humidity int) string{
	var result string
	if humidity<20{
		result="Kering"
	}else if humidity<=60 && humidity>=20{
		result="Normal"
	}else if humidity>60{
		result="Lembab"
	}
	return result
}