package structs

func CheckPesel(p *Person, hash string) {
	if hash != p.pesel {
		panic("error")
	}
}
