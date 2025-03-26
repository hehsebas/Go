package models

import (
	"errors"
	"marketplace/config"
)

// User representa el modelo de usuario
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"` // No devolveremos la contraseña en las respuestas
}

// CrearUsuario inserta un nuevo usuario en la base de datos
func CrearUsuario(user *User) error {
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	err := config.DB.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return errors.New("error al crear usuario: " + err.Error())
	}
	return nil
}

// ObtenerUsuarios devuelve todos los usuarios
func ObtenerUsuarios() ([]User, error) {
	rows, err := config.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, errors.New("error al obtener usuarios: " + err.Error())
	}
	defer rows.Close()

	var usuarios []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, user)
	}
	return usuarios, nil
}
func ObtenerUsuarioPorID() ([]User, error) {
	rows, err := config.DB.Query("SELECT id FROM users")
	if err != nil {
		return nil, errors.New("ID no encontrado")
	}
	defer rows.Close()
	var userId []User
	for rows.Next() {
		if err := rows.Scan(&user.ID); err != nil {
			return nil, err
		}
		userId = append(userId, user)
	}
	return userId, nil
}
func ActualizarUsuario(id) ([]User, error) {

}

// Product representa el modelo de producto
type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

// CrearProducto inserta un nuevo producto en la base de datos
func CrearProducto(product *Product) error {
	query := `INSERT INTO products (name, price, category) VALUES ($1, $2, $3) RETURNING id`
	err := config.DB.QueryRow(query, product.Name, product.Price, product.Category).Scan(&product.ID)
	if err != nil {
		return errors.New("error al crear producto: " + err.Error())
	}
	return nil
}

// ObtenerProductos devuelve todos los productos
func ObtenerProductos() ([]Product, error) {
	rows, err := config.DB.Query("SELECT id, name, price, category FROM products")
	if err != nil {
		return nil, errors.New("error al obtener productos: " + err.Error())
	}
	defer rows.Close()

	var productos []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Category); err != nil {
			return nil, err
		}
		productos = append(productos, product)
	}
	return productos, nil
}

// ObtenerProductoPorID devuelve un producto por su ID
func ObtenerProductoPorID(id string) (Product, error) {
	var product Product
	query := `SELECT id, name, price, category FROM products WHERE id = $1`
	err := config.DB.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price, &product.Category)
	if err != nil {
		return product, errors.New("producto no encontrado: " + err.Error())
	}
	return product, nil
}

// ActualizarProducto actualiza un producto en la base de datos
func ActualizarProducto(id string, product *Product) error {
	query := `UPDATE products SET name = $1, price = $2, category = $3 WHERE id = $4`
	_, err := config.DB.Exec(query, product.Name, product.Price, product.Category, id)
	if err != nil {
		return errors.New("error al actualizar producto: " + err.Error())
	}
	return nil
}

// EliminarProducto elimina un producto por su ID
func EliminarProducto(id string) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := config.DB.Exec(query, id)
	if err != nil {
		return errors.New("error al eliminar producto: " + err.Error())
	}
	return nil
}
