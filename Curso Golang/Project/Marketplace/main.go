package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"marketplace/products"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib" // Driver PostgreSQL
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var jwtSecret = []byte("mi_clave_secreta") // Clave secreta para firmar el token
// Función para parsear el token JWT
func parseToken(token string) (jwt.MapClaims, error) {
	tokenParts := strings.Split(token, "Bearer ")
	if len(tokenParts) != 2 {
		return nil, fmt.Errorf("token no válido")
	}

	jwtToken, err := jwt.Parse(tokenParts[1], func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !jwtToken.Valid {
		return nil, fmt.Errorf("token inválido")
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("no se pudieron obtener los claims del token")
	}

	return claims, nil
}

func main() {
	// DB Conexión
	db, err := sql.Open("pgx", "postgres://admin:admin123@localhost:5432/marketplace_db")
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	defer db.Close()

	// Verificar conexión a la base de datos
	if err := db.Ping(); err != nil {
		log.Fatal("No se pudo conectar a la base de datos:", err)
	}
	log.Println("Conexión exitosa a PostgreSQL")

	// Configuración de GIN
	r := gin.Default()

	// Ruta para crear un usuario
	r.POST("/users", func(c *gin.Context) {
		var user User
		// Validar que el JSON recibido sea correcto
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
			return
		}

		// Validar que el nombre no esté vacío
		if user.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El nombre es obligatorio"})
			return
		}

		// Validar que el correo tenga un formato válido
		emailRegex := `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`
		if match, _ := regexp.MatchString(emailRegex, user.Email); !match {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Correo electrónico inválido"})
			return
		}

		// Validar que la contraseña tenga al menos 8 caracteres
		if len(user.Password) < 8 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "La contraseña debe tener al menos 8 caracteres"})
			return
		}

		// Encriptar la contraseña
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al encriptar la contraseña"})
			return
		}
		user.Password = string(hashedPassword)

		// Insertar el usuario en la base de datos
		query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
		err = db.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el usuario", "details": err.Error()})
			return
		}

		// Responder con el usuario creado
		c.JSON(http.StatusCreated, user)
	})

	// Ruta para hacer login y obtener un token JWT
	r.POST("/login", func(c *gin.Context) {
		var loginData User
		// Recibir datos de login (email y password)
		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
			return
		}

		// Buscar el usuario por email
		var user User
		query := "SELECT id, name, email, password FROM users WHERE email = $1"
		err := db.QueryRow(query, loginData.Email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no encontrado"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al verificar el usuario"})
			return
		}

		// Verificar la contraseña
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Contraseña incorrecta"})
			return
		}

		// Crear un JWT token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":    user.ID,
			"email": user.Email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(), // Expira en 24 horas
		})

		// Firmar el token con la clave secreta
		tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar el token"})
			return
		}

		// Responder con el token
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	})

	// Ruta para listar todos los usuarios
	r.GET("/users", func(c *gin.Context) {
		rows, err := db.Query("SELECT id, name, email, password FROM users")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los usuarios"})
			return
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var user User
			if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al escanear los usuarios"})
				return
			}
			users = append(users, user)
		}

		c.JSON(http.StatusOK, users)
	})

	// Ruta para obtener un usuario por ID
	r.GET("/users/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		var user User
		query := "SELECT id, name, email, password FROM users WHERE id = $1"
		err = db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el usuario"})
			return
		}

		c.JSON(http.StatusOK, user)
	})
	// Ruta para actualizar un usuario
	r.PUT("/users/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		// Verificar si el usuario tiene permiso para actualizar (por ejemplo, si el ID del usuario coincide con el ID del token)
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autenticación requerido"})
			return
		}

		claims, err := parseToken(token)
		if err != nil || claims["id"] != float64(id) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No tienes permisos para actualizar este usuario"})
			return
		}

		var user User
		// Recibir los nuevos datos del usuario
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
			return
		}

		// Actualizar el usuario en la base de datos
		query := "UPDATE users SET name = $1, email = $2 WHERE id = $3"
		_, err = db.Exec(query, user.Name, user.Email, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el usuario"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado"})
	})
	// Ruta para eliminar un usuario
	r.DELETE("/users/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		// Verificar si el usuario tiene permiso para eliminar (por ejemplo, si el ID del usuario coincide con el ID del token)
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autenticación requerido"})
			return
		}

		claims, err := parseToken(token)
		if err != nil || claims["id"] != float64(id) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No tienes permisos para eliminar este usuario"})
			return
		}

		// Eliminar el usuario de la base de datos
		query := "DELETE FROM users WHERE id = $1"
		_, err = db.Exec(query, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el usuario"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado"})
	})

	// Registrar rutas para productos
	products.RegisterProductRoutes(r, db)
	// Levantar el servidor
	r.Run(":8080")
}
