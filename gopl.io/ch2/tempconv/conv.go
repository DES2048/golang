package tempconv

func CtoF(c Celcius) Fahrengeit {
	// явно приводим значения к нужному типу
	return Fahrengeit(c*9/5 + 32)
}

func FtoC(f Fahrengeit) Celcius {
	return Celcius((f-32)*5/9)
}

func CtoK(c Celcius) Kalvin {
	return Kalvin(c + 273.15)
}

func KtoC(k Kalvin) Celcius {
	return Celcius(k - 273.15)
}

func FtoK(f Fahrengeit) Kalvin  {
	return CtoK(FtoC(f))
}

func KtoF(k Kalvin) Fahrengeit {
	return CtoF(KtoC(k))
}
