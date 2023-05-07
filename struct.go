package main

//-----------------------------------------------------------------
//---------------ESTRUCTURAS DE LA BASE DE DATOS ------------------
//-----------------------------------------------------------------

type Facultades struct {
	ID     uint   `json:"id" form:"id"`
	Name   string `json:"name" form:"name" binding:"required"`
	Activa bool   `json:"activa" form:"activa"`
}

type Planes struct {
	ID          uint   `json:"id" form:"id"`
	Facultad    string `json:"facultad" form:"facultad" binding:"required"`
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	Activa      bool   `json:"activa" form:"activa"`
}

type Dpto struct {
	ID       uint   `json:"id" form:"id"`
	Facultad string `json:"facultad" form:"facultad" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Activa   bool   `json:"activa" form:"activa"`
}

type Profesores struct {
	ID     uint   `json:"id" form:"id"`
	Dpto   string `json:"dpto" form:"dpto" binding:"required"`
	Name   string `json:"name" form:"name" binding:"required"`
	Activa bool   `json:"activa" form:"activa"`
}

type Anos struct {
	ID     uint   `json:"id" form:"id"`
	Plan   string `json:"plan" form:"plan" binding:"required"`
	Name   string `json:"name" form:"name" binding:"required"`
	Activa bool   `json:"activa" form:"activa"`
}

type Semestres struct {
	ID          uint   `json:"id" form:"id"`
	Ano         string `json:"ano" form:"ano" binding:"required"`
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description"`
	Activa      bool   `json:"activa" form:"activa"`
}

type Asignaturas struct {
	ID                uint   `json:"id" form:"id"`
	Dpto              string `json:"dpto" form:"dpto" binding:"required"`
	Plan              string `json:"plan" form:"plan" binding:"required"`
	Ano               string `json:"ano" form:"ano" binding:"required"`
	Semestre          string `json:"semestre" form:"semestre" binding:"required"`
	Name              string `json:"name" form:"name" binding:"required"`
	Prof              uint   `json:"profesor" form:"profesor" binding:"required"`
	E_bienvenida      bool   `json:"e_bienvenida" form:"e_bienvenida"`
	D_general         bool   `json:"d_general" form:"d_general"`
	G_didactica       bool   `json:"g_didactica" form:"g_didactica"`
	E_intercambio     bool   `json:"e_intercambio" form:"e_intercambioactiva"`
	Glosario          bool   `json:"glosario" form:"glosario"`
	B_general         bool   `json:"b_general" form:"b_general"`
	O_unidad_tema     bool   `json:"o_unidad_tema" form:"o_unidad_tema"`
	R_educativos      bool   `json:"r_educativos" form:"r_educativos"`
	A_aprendizaje     bool   `json:"a_aprendizaje" form:"a_aprendizaje"`
	A_evaluacion      bool   `json:"a_evaluacion" form:"a_evaluacion"`
	B_tema            bool   `json:"b_tema" form:"b_tema"`
	U_url             bool   `json:"U_url" form:"U_url"`
	M_complementarios bool   `json:"M_complementarios" form:"M_complementarios"`
	Calidad           string `json:"calidad" form:"calidad"`
	Activa            bool   `json:"activa" form:"activa"`
}

type Reporte struct {
	ID     uint   `json:"id" form:"id" binding:"required"`
	Plan   string `json:"plan" form:"plan" binding:"required"`
	Fecha  string `json:"fecha" form:"fecha"`
	C_Asig int    `json:"c_asig" form:"c_asig"`
	C_MB   int    `json:"c_mb" form:"c_mb"`
	C_B    int    `json:"c_b" form:"c_b"`
	C_R    int    `json:"c_r" form:"c_r"`
	C_M    int    `json:"c_m" form:"c_m"`
}
