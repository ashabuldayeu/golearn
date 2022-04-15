package main

type Client interface {
	Calculate() float32
}

type Coefficient float32

type Defclient struct {
	Elcount int
	Elcost  int
}

type Limclient struct {
	Defclient
	Lim     int         `default:150`
	Addcoef Coefficient `default:1.33`
}

type Prefclient1 struct {
	Defclient
	Prefcoeff Coefficient `default:0.67`
}

type Prefclient2 struct {
	Defclient
	Lim       int         `default:50`
	Prefcoeff Coefficient `default:0.0`
}

func (dclient Defclient) Calculate() float32 {
	return float32(dclient.Elcost * dclient.Elcount)
}
func (lclient Limclient) Calculate() float32 {
	coeffCount := lclient.Elcount - lclient.Lim

	if coeffCount < 0 {
		coeffCount = 0
	}

	return (float32(coeffCount)*float32(lclient.Addcoef) + (float32(lclient.Elcount - coeffCount))) * float32(lclient.Elcost)
}

func (prClient1 Prefclient1) Calculate() float32 {
	return float32(prClient1.Prefcoeff) * float32(prClient1.Elcost*prClient1.Elcount)
}

func (prclient2 Prefclient2) Calculate() float32 {
	unLimcount := prclient2.Elcount - prclient2.Lim

	if unLimcount < 0 {
		unLimcount = 0
	}
	return float32(unLimcount*int(prclient2.Elcost)) + float32((prclient2.Elcount-unLimcount)*int(prclient2.Prefcoeff))
}
