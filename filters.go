package main

import (
	"fmt"
	"math"
	"strconv"
)

// --------------Filters
func Contar_true(a Asignaturas) float64 {
	var num = 0
	if a.E_bienvenida {
		num += 1
	}
	if a.D_general {
		num += 1
	}
	if a.G_didactica {
		num += 1
	}
	if a.E_intercambio {
		num += 1
	}
	if a.Glosario {
		num += 1
	}
	if a.B_general {
		num += 1
	}
	if a.O_unidad_tema {
		num += 1
	}
	if a.R_educativos {
		num += 1
	}
	if a.A_aprendizaje {
		num += 1
	}
	if a.A_evaluacion {
		num += 1
	}
	if a.B_tema {
		num += 1
	}
	if a.U_url {
		num += 1
	}
	if a.M_complementarios {
		num += 1
	}
	return float64(num)
}

func Redondear(result float64, lugares int) float64 {
	var mostrar float64

	if lugares == 1 {
		mostrar = float64(math.Round(float64(result)))
	} else if lugares == 2 {
		redondeo, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", result), 64)
		mostrar = float64(redondeo)
	} else if lugares == 3 {
		redondeo, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", result), 64)
		mostrar = float64(redondeo)
	} else {
		redondeo, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", result), 64)
		mostrar = float64(redondeo)
	}

	return float64(mostrar)
}

func (a Asignaturas) EvaluarCalidad() string {
	result := ""
	num := Contar_true(a)
	if num <= 5 {
		result = "Mal"
	} else if num > 5 && num <= 8 {
		result = "Regular"
	} else if num > 8 && num <= 11 {
		result = "Bien"
	} else {
		result = "Muy bien"
	}
	return result
}

func (a Asignaturas) PorcientoCalidad() float64 {
	return Redondear(Contar_true(a)/13*100, 2)
}

func (d Dpto) CantAsig(p Planes) int {
	var asigs []Asignaturas

	db.Where("Plan = ? and Dpto = ?", p.ID, d.ID).Find(&asigs)

	return len(asigs)
}

func (d Dpto) CalidadesM(p Planes) int {
	var asigs []Asignaturas
	db.Where("Plan = ? and Dpto = ?", p.ID, d.ID).Find(&asigs)
	mal := 0
	regular := 0
	bien := 0
	muybien := 0
	for _, asig := range asigs {
		if asig.EvaluarCalidad() == "Mal" {
			mal += 1
		} else if asig.EvaluarCalidad() == "Regular" {
			regular += 1
		} else if asig.EvaluarCalidad() == "Bien" {
			bien += 1
		} else {
			muybien += 1
		}
	}

	return mal

}

func (d Dpto) CalidadesR(p Planes) int {
	var asigs []Asignaturas
	db.Where("Plan = ? and Dpto = ?", p.ID, d.ID).Find(&asigs)
	mal := 0
	regular := 0
	bien := 0
	muybien := 0
	for _, asig := range asigs {
		if asig.EvaluarCalidad() == "Mal" {
			mal += 1
		} else if asig.EvaluarCalidad() == "Regular" {
			regular += 1
		} else if asig.EvaluarCalidad() == "Bien" {
			bien += 1
		} else {
			muybien += 1
		}
	}

	return regular

}

func (d Dpto) CalidadesB(p Planes) int {
	var asigs []Asignaturas
	db.Where("Plan = ? and Dpto = ?", p.ID, d.ID).Find(&asigs)
	mal := 0
	regular := 0
	bien := 0
	muybien := 0
	for _, asig := range asigs {
		if asig.EvaluarCalidad() == "Mal" {
			mal += 1
		} else if asig.EvaluarCalidad() == "Regular" {
			regular += 1
		} else if asig.EvaluarCalidad() == "Bien" {
			bien += 1
		} else {
			muybien += 1
		}
	}

	return bien

}

func (d Dpto) CalidadesMB(p Planes) int {
	var asigs []Asignaturas
	db.Where("Plan = ? and Dpto = ?", p.ID, d.ID).Find(&asigs)
	mal := 0
	regular := 0
	bien := 0
	muybien := 0
	for _, asig := range asigs {
		if asig.EvaluarCalidad() == "Mal" {
			mal += 1
		} else if asig.EvaluarCalidad() == "Regular" {
			regular += 1
		} else if asig.EvaluarCalidad() == "Bien" {
			bien += 1
		} else {
			muybien += 1
		}
	}

	return muybien

}

func (d Dpto) EvaluarDpto(p Planes) string {
	//total := d.CantAsig()
	var asigs []Asignaturas
	db.Where("Plan = ? and Dpto = ?", p.ID, d.ID).Find(&asigs)
	mal := 0
	regular := 0
	bien := 0
	muybien := 0
	for _, asig := range asigs {
		if asig.EvaluarCalidad() == "Mal" {
			mal += 1
		} else if asig.EvaluarCalidad() == "Regular" {
			regular += 1
		} else if asig.EvaluarCalidad() == "Bien" {
			bien += 1
		} else {
			muybien += 1
		}
	}

	if mal > regular && mal > bien && mal > muybien {
		return "Mal"
	} else if regular > mal && regular > bien && regular > muybien {
		return "Regular"
	} else if bien > mal && bien > regular && bien > muybien {
		return "Bien"
	} else if muybien > mal && muybien > regular && muybien > bien {
		return "Muy bien"
	} else {

		if mal == regular && mal > bien && mal > muybien {
			return "Regular"
		} else if mal == bien && mal > regular && mal > muybien {
			return "Regular"
		} else if mal == muybien && mal > regular && mal > bien {
			return "Bien"
		} else if regular == bien && regular > mal && regular > muybien {
			return "Bien"
		} else if regular == muybien && regular > mal && regular > bien {
			return "Bien"
		} else if bien == muybien && bien > mal && bien > regular {
			return "Bien"
		} else if bien == mal && bien == regular && bien == muybien {
			return "Bien"
		}
		return "Definir"
	}

}

func (p Profesores) ExtractFacultad() string {
	//total := d.CantAsig()
	var dpto Dpto
	db.Where("ID = ?", p.Dpto).Find(&dpto)
	return dpto.Facultad
}

func (a Asignaturas) ExtractName() string {
	//total := d.CantAsig()
	var prof Profesores
	db.Where("ID = ?", a.Prof).Find(&prof)
	return prof.Name

}

func (a Asignaturas) ConvertUint() uint {
	u64, _ := strconv.ParseUint(a.Plan, 10, 32)
	return uint(u64)
}

func (a Asignaturas) ProfName() string {
	var prof Profesores
	db.Where("ID = ?", a.Prof).Find(&prof)
	return prof.Name
}
