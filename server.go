package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"strings"
)

// ------------------------
const (
	userkey = "user"
)

// ------------------------
func engine() *gin.Engine {
	r := gin.New()
	//r.Delims("<owl", "owl>")
	//-- Cargar plantillas

	r.LoadHTMLGlob("./template/*")
	//-- Cargar archivos estaticos
	r.StaticFS("/static/", http.Dir("./static/"))
	r.Use(sessions.Sessions("mysession", cookie.NewStore([]byte("secret"))))
	r.GET("/", func(c *gin.Context) {
		http.Redirect(c.Writer, c.Request, "/private/home", http.StatusFound)
	})
	r.POST("/login", login)
	r.GET("/logout", logout)

	private := r.Group("/private")
	//private.Use(AuthRequired)
	{
		private.GET("/home", home)
		private.GET("/facultades", facultades)
		private.GET("/facultad/:id", planes)
		private.GET("/reporte/:id", reporteModaliad)
		private.GET("/reportedpto/:id", reporteDpto)
		private.GET("/profesores/:id", profesoresDpto)
		private.GET("/ano/:facultad/:modalidad", anos)
		private.GET("/semestre/:facultad/:modalidad/:ano", semestres)
		private.GET("/asignaturas/:facultad/:modalidad/:ano/:semestre", asignaturas)

		private.POST("/AddFacultad", AddFacultad)
		private.POST("/AddPlan", AddPlan)
		private.POST("/AddDpto", AddDpto)
		private.POST("/DElProfesor", DElProfesor)
		private.POST("/AddProfesor", AddProfesor)
		private.POST("/AddAno", AddAno)
		private.POST("/AddSemestre", AddSemestre)
		private.POST("/DelSemestre", DelSemestre)
		private.POST("/AddAsignatura", AddAsignatura)
		private.POST("/EditAsignatura", EditAsignatura)
		private.POST("/EditAno", EditAno)
		private.POST("/AddReporte", AddReporte)
		private.POST("/DelAsignatura", DelAsignatura)
		private.POST("/EditDpto", EditDpto)
		private.POST("/EditPlan", EditPlan)
	}
	return r
}

//-----------

// -----------
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}

// login is a handler that parses a form and checks for specific data
func login(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Validate form input
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check for username and password match, usually from a database
	if username != "admin" || password != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// Save the username in the session
	session.Set(userkey, username) // In real world usage you'd set this to the users ID
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		http.Redirect(c.Writer, c.Request, "/", http.StatusFound)
		return
	}
	session.Delete(userkey)
	if err := session.Save(); err != nil {
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		http.Redirect(c.Writer, c.Request, "/", http.StatusFound)
		return
	}
	//c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
	http.Redirect(c.Writer, c.Request, "/", http.StatusFound)
}

// ----------------GET
// ----Home
func home(c *gin.Context) {
	//db, err = gorm.Open("sqlite3", "./db/data.db")
	//defer db.Close()
	////--------------
	//var hosts []Host_Group
	//db.Find(&hosts)
	session := sessions.Default(c)
	user := session.Get(userkey)

	c.HTML(http.StatusOK, "index.html", gin.H{"session": user})
}

// ----facultades
func facultades(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./db/data.db")
	defer db.Close()
	////--------------
	var facultades []Facultades
	db.Find(&facultades)
	session := sessions.Default(c)
	user := session.Get(userkey)

	c.HTML(http.StatusOK, "facultades.html", gin.H{"session": user, "facultades": facultades})
}

// ----planes
func planes(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./db/data.db")
	defer db.Close()
	////--------------
	id := c.Params.ByName("id")
	var facultad Facultades
	db.Where("ID = ?", id).Find(&facultad)
	var planes []Planes
	db.Where("Facultad = ?", id).Find(&planes)

	var dptos []Dpto
	db.Where("Facultad = ?", id).Find(&dptos)
	session := sessions.Default(c)
	user := session.Get(userkey)
	c.HTML(http.StatusOK, "planes.html", gin.H{"session": user, "planes": planes, "dptos": dptos, "facultad": id, "nameFacultad": facultad.Name})
}

// ----planes
func reporteModaliad(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./db/data.db")
	defer db.Close()
	////--------------
	session := sessions.Default(c)
	user := session.Get(userkey)
	////--------------
	id := c.Params.ByName("id")
	var planes Planes
	db.Where("ID = ?", id).Find(&planes)

	var anos []Anos
	db.Where("Plan = ?", id).Find(&anos)
	var cant_cursos float64
	var mb float64
	var b float64
	var r float64
	var m float64
	for _, ano := range anos {
		var semestres []Semestres
		db.Where("Ano = ?", ano.ID).Find(&semestres)

		for _, semestre := range semestres {
			var asignaturas []Asignaturas
			db.Where("Semestre = ?", semestre.ID).Find(&asignaturas)
			cant_cursos += float64(len(asignaturas))
			for _, asig := range asignaturas {
				calidad := asig.EvaluarCalidad()
				if calidad == "Muy bien" {
					mb += 1
				} else if calidad == "Bien" {
					b += 1
				} else if calidad == "Regular" {
					r += 1
				} else {
					m += 1
				}
			}
		}
	}

	p_mb := Redondear(mb/cant_cursos*100, 2)
	p_b := Redondear(b/cant_cursos*100, 2)
	p_r := Redondear(r/cant_cursos*100, 2)
	p_m := Redondear(m/cant_cursos*100, 2)

	var reportes []Reporte
	db.Where("Plan = ?", id).Find(&reportes)

	c.HTML(http.StatusOK, "reporte.html", gin.H{"session": user, "planes": planes, "modalidad": id, "cant_curso": cant_cursos, "mb": mb, "b": b, "r": r, "m": m, "p_mb": p_mb, "p_b": p_b, "p_r": p_r, "p_m": p_m, "reportes": reportes})
}

func reporteDpto(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./db/data.db")
	defer db.Close()
	session := sessions.Default(c)
	user := session.Get(userkey)
	id := c.Params.ByName("id")

	var dpto Dpto
	db.Where("ID = ?", id).Find(&dpto)

	var planes []Planes
	db.Where("Facultad = ?", dpto.Facultad).Find(&planes)

	var asig []Asignaturas
	db.Where("Dpto = ?", dpto.ID).Find(&asig)

	//date := CurrentTime()
	//dateString := FormatDate(strings.Split(date, " ")[0])

	c.HTML(http.StatusOK, "reporteDPTO.html", gin.H{"title": "Reporte del departamento " + dpto.Name, "session": user, "dpto": dpto, "planes": planes, "asigs": asig})
}

func profesoresDpto(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./db/data.db")
	defer db.Close()

	session := sessions.Default(c)
	user := session.Get(userkey)

	id := c.Params.ByName("id")
	var dpto Dpto
	db.Where("ID = ?", id).Find(&dpto)

	var profs []Profesores
	db.Where("Dpto = ?", dpto.ID).Find(&profs)

	c.HTML(http.StatusOK, "profesores.html", gin.H{"title": "Profesores del departamento " + dpto.Name, "session": user, "dpto": dpto, "profs": profs})
}

// ----anos
func anos(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./db/data.db")
	defer db.Close()
	////--------------
	facultad := c.Params.ByName("facultad")
	modalidad := c.Params.ByName("modalidad")
	var fac Facultades
	db.Where("ID = ?", facultad).Find(&fac)
	var anos []Anos
	db.Where("Plan = ?", modalidad).Find(&anos)
	var plan Planes
	db.Where("ID = ?", modalidad).Find(&plan)
	session := sessions.Default(c)
	user := session.Get(userkey)
	c.HTML(http.StatusOK, "anos.html", gin.H{"session": user, "anos": anos, "facultad": facultad, "nameFacultad": fac.Name, "nameModalidad": plan.Name, "modalidad": modalidad})
}

// ----semestres
func semestres(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./db/data.db")
	defer db.Close()
	////--------------
	facultad := c.Params.ByName("facultad")
	modalidad := c.Params.ByName("modalidad")
	ano := c.Params.ByName("ano")
	var fac Facultades
	db.Where("ID = ?", facultad).Find(&fac)
	var plan Planes
	db.Where("ID = ?", modalidad).Find(&plan)
	var anos Anos
	db.Where("ID = ?", ano).Find(&anos)
	var semestre []Semestres
	db.Where("Ano = ?", ano).Find(&semestre)

	session := sessions.Default(c)
	user := session.Get(userkey)

	//title := plan.Name + "/" + anos.Name + "/periodos"

	c.HTML(http.StatusOK, "semestres.html", gin.H{"session": user, "ano": ano, "facultad": facultad, "modalidad": modalidad, "semestres": semestre, "nameFacultad": fac.Name, "nameModalidad": plan.Name, "nameAno": anos.Name})
}

// ----semestres
func asignaturas(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./db/data.db")
	defer db.Close()
	////--------------
	facultad := c.Params.ByName("facultad")
	modalidad := c.Params.ByName("modalidad")
	ano := c.Params.ByName("ano")
	semestre := c.Params.ByName("semestre")
	var fac Facultades
	db.Where("ID = ?", facultad).Find(&fac)
	var plan Planes
	db.Where("ID = ?", modalidad).Find(&plan)
	var anos Anos
	db.Where("ID = ?", ano).Find(&anos)
	var periodo Semestres
	db.Where("ID = ?", semestre).Find(&periodo)
	var asignaturas []Asignaturas
	db.Where("Semestre = ?", semestre).Find(&asignaturas)
	var profesor []Profesores
	db.Where("Activa = ?", true).Find(&profesor)

	session := sessions.Default(c)
	user := session.Get(userkey)

	c.HTML(http.StatusOK, "asignaturas.html", gin.H{"session": user, "ano": ano, "facultad": facultad, "modalidad": modalidad, "semestre": semestre, "asignaturas": asignaturas, "profesores": profesor, "nameFacultad": fac.Name, "nameModalidad": plan.Name, "nameAno": anos.Name, "namePeriodo": periodo.Name})
}

// ----add facultades
func AddFacultad(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./db/data.db")
	defer db.Close()
	////--------------
	name := c.PostForm("name")
	activa := c.PostForm("activa")

	var facultads []Facultades
	db.Where("Name = ?", name).Find(&facultads)
	if len(facultads) == 0 {
		var facultad Facultades

		facultad.Name = name
		if activa == "true" {
			facultad.Activa = true
		} else {
			facultad.Activa = false
		}
		db.Save(&facultad)

	}

	c.JSON(200, gin.H{"msg": "add facultad"})
	return
}

// ----add planes
func AddPlan(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./db/data.db")
	defer db.Close()
	////--------------
	id := c.PostForm("id")
	name := c.PostForm("name")
	descripcion := c.PostForm("descripcion")
	activa := c.PostForm("activa")

	var plans []Planes
	db.Where("Facultad = ? and Name = ?", id, name).Find(&plans)
	if len(plans) == 0 {
		var plan Planes
		plan.Facultad = id
		plan.Name = name
		plan.Description = descripcion
		if activa == "true" {
			plan.Activa = true
		} else {
			plan.Activa = false
		}
		db.Save(&plan)
	}
	c.JSON(200, gin.H{"msg": "add plan"})
	return
}

// ----add AddReporte
func AddReporte(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./db/data.db")
	defer db.Close()
	////--------------
	id := c.PostForm("id")
	fecha := c.PostForm("fecha")
	c_asig, _ := strconv.Atoi(c.PostForm("c_asig"))
	mb, _ := strconv.Atoi(c.PostForm("mb"))
	b, _ := strconv.Atoi(c.PostForm("b"))
	r, _ := strconv.Atoi(c.PostForm("r"))
	m, _ := strconv.Atoi(c.PostForm("m"))

	var reporte Reporte
	reporte.Plan = id
	reporte.Fecha = fecha
	reporte.C_Asig = c_asig
	reporte.C_MB = mb
	reporte.C_B = b
	reporte.C_R = r
	reporte.C_M = m
	db.Save(&reporte)

	c.JSON(200, gin.H{"msg": "add reporte"})
	return
}

// ----add dpto
func AddDpto(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./db/data.db")
	defer db.Close()
	////--------------
	id := c.PostForm("id")
	name := c.PostForm("name")
	activa := c.PostForm("activa")

	var dptos []Dpto
	db.Where("Facultad = ? and Name = ?", id, name).Find(&dptos)
	if len(dptos) == 0 {
		var dpto Dpto
		dpto.Facultad = id
		dpto.Name = name
		if activa == "true" {
			dpto.Activa = true
		} else {
			dpto.Activa = false
		}
		db.Save(&dpto)

	}

	c.JSON(200, gin.H{"msg": "add dpto"})
	return
}

// ----add profesor
func AddProfesor(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./db/data.db")
	defer db.Close()
	////--------------
	dpto := c.PostForm("dpto")
	name := c.PostForm("name")

	var prof Profesores
	prof.Name = name
	prof.Dpto = dpto
	prof.Activa = true

	db.Save(&prof)

	c.JSON(200, gin.H{"msg": "add profesor"})
	return
}

func DElProfesor(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./db/data.db")
	defer db.Close()
	////--------------
	id := c.PostForm("id")
	var prof Profesores
	db.Where("ID = ?", id).Delete(&prof)
	c.JSON(200, gin.H{"msg": "del profesor"})
	return
}

// ----edit dpto
func EditDpto(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./db/data.db")
	defer db.Close()
	////--------------
	id := c.PostForm("id")
	name := c.PostForm("name")
	activa := c.PostForm("activa")

	var dpto Dpto
	db.Where("ID = ?", id).Find(&dpto)

	dpto.Name = name
	if activa == "true" {
		dpto.Activa = true
	} else {
		dpto.Activa = false
	}
	db.Save(&dpto)

	c.JSON(200, gin.H{"msg": "add dpto"})
	return
}

// ----edit planes
func EditPlan(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./db/data.db")
	defer db.Close()
	////--------------
	id := c.PostForm("id")
	name := c.PostForm("name")
	activa := c.PostForm("activa")

	var plan Planes
	db.Where("ID = ?", id).Find(&plan)

	plan.Name = name
	if activa == "true" {
		plan.Activa = true
	} else {
		plan.Activa = false
	}
	db.Save(&plan)

	c.JSON(200, gin.H{"msg": "edit plan"})
	return
}

// ----add planes
func AddAno(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./db/data.db")
	defer db.Close()
	////--------------
	id := c.PostForm("id")
	name := c.PostForm("name")
	activa := c.PostForm("activa")
	fmt.Println(id)
	var ano []Anos
	db.Where("Name = ? and Plan = ?", name, id).Find(&ano)
	if len(ano) == 0 {
		var anos Anos
		anos.Plan = id
		anos.Name = name
		if activa == "true" {
			anos.Activa = true
		} else {
			anos.Activa = false
		}
		db.Save(&anos)

	}

	c.JSON(200, gin.H{"msg": "add anos"})
	return
}

// ----add planes
func EditAno(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./db/data.db")
	defer db.Close()
	////--------------
	id := c.PostForm("id")
	name := c.PostForm("name")
	activa := c.PostForm("activa")

	var anos Anos
	db.Where("ID = ?", id).Find(&anos)

	anos.Name = name
	if activa == "true" {
		anos.Activa = true
	} else {
		anos.Activa = false
	}
	db.Save(&anos)

	c.JSON(200, gin.H{"msg": "edit anos"})
	return
}

// ----add semestre
func AddSemestre(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./db/data.db")
	defer db.Close()
	////--------------
	id := c.PostForm("id")
	name := c.PostForm("name")
	descripcion := c.PostForm("descripcion")
	activa := c.PostForm("activa")

	var semestres []Semestres
	db.Where("Name = ? and Ano = ?", name, id).Find(&semestres)
	if len(semestres) == 0 {
		var semestre Semestres
		semestre.Ano = id
		semestre.Name = name
		semestre.Description = descripcion
		if activa == "true" {
			semestre.Activa = true
		} else {
			semestre.Activa = false
		}
		db.Save(&semestre)

	}

	c.JSON(200, gin.H{"msg": "add semestre"})
	return
}

// ----del semenstre
func DelSemestre(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./db/data.db")
	defer db.Close()
	////--------------
	id := c.PostForm("id")
	var semestres Semestres
	db.Where("ID = ?", id).Find(&semestres)

	var asigs []Asignaturas
	db.Where("Semestre = ?", semestres.ID).Delete(&asigs)

	db.Delete(&semestres)

	c.JSON(200, gin.H{"msg": "del semestre"})
	return
}

// ----add asignatura
func AddAsignatura(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./db/data.db")
	defer db.Close()
	////
	plan := c.PostForm("plan")
	ano := c.PostForm("ano")
	semestre := c.PostForm("semestre")
	name := c.PostForm("name")
	profesor := c.PostForm("profesor")
	E_bienvenida := c.PostForm("E_bienvenida")
	D_general := c.PostForm("D_general")
	G_didactica := c.PostForm("G_didactica")
	E_intercambio := c.PostForm("E_intercambio")
	Glosario := c.PostForm("Glosario")
	B_general := c.PostForm("B_general")
	O_unidad_tema := c.PostForm("O_unidad_tema")
	R_educativos := c.PostForm("R_educativos")
	A_aprendizaje := c.PostForm("A_aprendizaje")
	A_evaluacion := c.PostForm("A_evaluacion")
	B_tema := c.PostForm("B_tema")
	U_url := c.PostForm("U_url")
	M_complementarios := c.PostForm("M_complementarios")

	var prof Profesores
	db.Where("ID = ?", profesor).Find(&prof)

	var asigs []Asignaturas
	db.Where("Name = ? and Semestre = ?", name, semestre).Find(&asigs)
	if len(asigs) == 0 {
		var asig Asignaturas
		asig.Plan = plan
		asig.Ano = ano
		asig.Semestre = semestre
		asig.Dpto = prof.Dpto
		asig.Name = name
		asig.Prof = prof.ID
		//--------
		if E_bienvenida == "true" {
			asig.E_bienvenida = true
		} else {
			asig.E_bienvenida = false
		}
		//--------

		// --------
		if D_general == "true" {
			asig.D_general = true
		} else {
			asig.D_general = false
		}
		//--------
		// --------
		if G_didactica == "true" {
			asig.G_didactica = true
		} else {
			asig.G_didactica = false
		}
		//--------
		// --------
		if E_intercambio == "true" {
			asig.E_intercambio = true
		} else {
			asig.E_intercambio = false
		}
		//--------
		// --------
		if Glosario == "true" {
			asig.Glosario = true
		} else {
			asig.Glosario = false
		}
		//--------
		// --------
		if B_general == "true" {
			asig.B_general = true
		} else {
			asig.B_general = false
		}
		//--------
		// --------
		if O_unidad_tema == "true" {
			asig.O_unidad_tema = true
		} else {
			asig.O_unidad_tema = false
		}
		//--------
		// --------
		if R_educativos == "true" {
			asig.R_educativos = true
		} else {
			asig.R_educativos = false
		}
		//--------
		// --------
		if A_aprendizaje == "true" {
			asig.A_aprendizaje = true
		} else {
			asig.A_aprendizaje = false
		}
		//--------
		// --------
		if A_evaluacion == "true" {
			asig.A_evaluacion = true
		} else {
			asig.A_evaluacion = false
		}
		//--------
		// --------
		if B_tema == "true" {
			asig.B_tema = true
		} else {
			asig.B_tema = false
		}
		//--------
		// --------
		if U_url == "true" {
			asig.U_url = true
		} else {
			asig.U_url = false
		}
		//--------
		// --------
		if M_complementarios == "true" {
			asig.M_complementarios = true
		} else {
			asig.M_complementarios = false
		}
		//--------
		asig.Activa = true
		db.Save(&asig)
	}

	c.JSON(200, gin.H{"msg": "add semestre"})
	return
}

// ----add asignatura
func DelAsignatura(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./db/data.db")
	defer db.Close()
	////
	id := c.PostForm("id")

	var asigs Asignaturas

	db.Where("ID = ?", id).Delete(&asigs)

	c.JSON(200, gin.H{"msg": "add semestre"})
	return
}

// ----Edit asignatura
func EditAsignatura(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "./db/data.db")
	defer db.Close()
	id := c.PostForm("id")
	name := c.PostForm("name")
	plan := c.PostForm("plan")
	ano := c.PostForm("ano")
	profesor := c.PostForm("profesor")
	E_bienvenida := c.PostForm("E_bienvenida")
	D_general := c.PostForm("D_general")
	G_didactica := c.PostForm("G_didactica")
	E_intercambio := c.PostForm("E_intercambio")
	Glosario := c.PostForm("Glosario")
	B_general := c.PostForm("B_general")
	O_unidad_tema := c.PostForm("O_unidad_tema")
	R_educativos := c.PostForm("R_educativos")
	A_aprendizaje := c.PostForm("A_aprendizaje")
	A_evaluacion := c.PostForm("A_evaluacion")
	B_tema := c.PostForm("B_tema")
	U_url := c.PostForm("U_url")
	M_complementarios := c.PostForm("M_complementarios")

	var prof Profesores
	db.Where("ID = ?", profesor).Find(&prof)
	var asig Asignaturas
	db.Where("ID = ?", id).Find(&asig)
	asig.Plan = plan
	asig.Ano = ano
	asig.Name = name
	asig.Dpto = prof.Dpto
	asig.Prof = prof.ID
	//--------
	if E_bienvenida == "true" {
		asig.E_bienvenida = true
	} else {
		asig.E_bienvenida = false
	}
	//--------

	// --------
	if D_general == "true" {
		asig.D_general = true
	} else {
		asig.D_general = false
	}
	//--------
	// --------
	if G_didactica == "true" {
		asig.G_didactica = true
	} else {
		asig.G_didactica = false
	}
	//--------
	// --------
	if E_intercambio == "true" {
		asig.E_intercambio = true
	} else {
		asig.E_intercambio = false
	}
	//--------
	// --------
	if Glosario == "true" {
		asig.Glosario = true
	} else {
		asig.Glosario = false
	}
	//--------
	// --------
	if B_general == "true" {
		asig.B_general = true
	} else {
		asig.B_general = false
	}
	//--------
	// --------
	if O_unidad_tema == "true" {
		asig.O_unidad_tema = true
	} else {
		asig.O_unidad_tema = false
	}
	//--------
	// --------
	if R_educativos == "true" {
		asig.R_educativos = true
	} else {
		asig.R_educativos = false
	}
	//--------
	// --------
	if A_aprendizaje == "true" {
		asig.A_aprendizaje = true
	} else {
		asig.A_aprendizaje = false
	}
	//--------
	// --------
	if A_evaluacion == "true" {
		asig.A_evaluacion = true
	} else {
		asig.A_evaluacion = false
	}
	//--------
	// --------
	if B_tema == "true" {
		asig.B_tema = true
	} else {
		asig.B_tema = false
	}
	//--------
	// --------
	if U_url == "true" {
		asig.U_url = true
	} else {
		asig.U_url = false
	}
	//--------
	// --------
	if M_complementarios == "true" {
		asig.M_complementarios = true
	} else {
		asig.M_complementarios = false
	}
	//--------
	asig.Activa = true
	db.Save(&asig)

	c.JSON(200, gin.H{"msg": "edit asignatura"})
	return
}
