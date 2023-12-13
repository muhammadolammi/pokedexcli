package pokeapi

type MyPokeMons map[string]PokeMonStrucResp

func NewPokeMonsData() MyPokeMons {
	pokeMonData := make(map[string]PokeMonStrucResp)
	return pokeMonData

}

func (pData MyPokeMons) AddPokeMon(name string, data PokeMonStrucResp) {
	pData[name] = data
}
